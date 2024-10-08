package auth

import (
	"context"
	"fmt"
	"math/rand"
	"sensei/conf"
	"sensei/ent"
	"sensei/ent/user"
	"sensei/mail"
	"sensei/utils"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type Svc interface {
	Login(ctx context.Context, form LoginForm) (utils.HttpResponse, *string)
	SignUp(ctx context.Context, form SignUpForm) utils.HttpResponse
}

type Store struct {
	DB *ent.Client
}

func hashPassword(password string) (result []byte, err error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return bytes, err
}

func checkPassword(hashPassword, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashPassword), []byte(password))
	if err != nil {
		return false
	}
	return true
}

func (s *Store) Login(ctx context.Context, form LoginForm) (utils.HttpResponse, *string) {
	if form.Mail == "" || form.Password == "" {
		res := utils.BadRequest(form, nil)
		return res, nil
	}
	loginUser, err := s.DB.User.Query().Where(user.MailEQ(form.Mail)).Only(ctx)
	if err != nil {
		res := utils.NotFoundEntity(form.Mail)
		return res, nil
	}

	if !checkPassword(loginUser.Password, form.Password) {
		res := utils.Unauthorized("Incorrect password", "001")
		return res, nil

	}
	if loginUser.MailValid == false {
		res := utils.Unauthorized("Mail not verified", "002")
		return res, nil
	}
	csrfToken, err := utils.GenerateRandomToken(64)
	if err != nil {
		res := utils.InternalError(err)
		return res, nil
	}
	// hash csrf
	hashedCsrfToken, err := utils.HashAndSalt(csrfToken)
	if err != nil {
		res := utils.InternalError(err)
		return res, nil
	}

	tokenString, err := utils.GenerateJWT(loginUser.ID.String(), loginUser.Mail, loginUser.Username, csrfToken)
	if err != nil {
		res := utils.InternalError(err)
		return res, nil
	}
	result := utils.LoginResult{Csrf: hashedCsrfToken}
	res := utils.OkLogged(result)
	return res, &tokenString
}

func (s *Store) SignUp(ctx context.Context, form SignUpForm) utils.HttpResponse {
	if form.UserName == "" || form.Mail == "" || form.Password == "" {
		res := utils.BadRequest(form, nil)
		return res
	}
	bytesPass, err := bcrypt.GenerateFromPassword([]byte(form.Password), 14)

	alreadyExistUser, err := s.DB.User.Query().Where(user.UsernameEQ(form.UserName)).First(ctx)
	if alreadyExistUser != nil {
		res := utils.DuplicatedKey("user")
		return res
	}
	alreadyExistMail, err := s.DB.User.Query().Where(user.MailEQ(form.Mail)).First(ctx)
	if alreadyExistMail != nil {
		res := utils.DuplicatedKey("mail")
		return res

	}

	user, err := s.DB.User.Create().SetUsername(form.UserName).SetPassword(string(bytesPass[:])).SetMail(form.Mail).Save(ctx)
	if err != nil {
		res := utils.InternalError(err)
		return res
	}
	expireDate := time.Now().Add(time.Minute * 15)
	code := rand.Intn((1000000))
	codeStr := fmt.Sprintf("%06d", code)

	_, err = s.DB.VerificationCode.Create().SetUserID(user.ID).SetExpireDate(expireDate).SetType(utils.VALIDATION_TYPE).SetCode(codeStr).SetValid(true).Save(ctx)
	fmt.Println(err)
	conf := conf.Get()
	if err == nil && conf.Mail.Enabled {
		err = mail.SendMail(form.Mail, "Validation code", fmt.Sprintf("You can use the code %s to validate your account", codeStr))
	}
	res := utils.OkCreated(user)
	return res

}
