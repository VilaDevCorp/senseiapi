// Code generated by ent, DO NOT EDIT.

package ent

import (
	"sensei/ent/activity"
	"sensei/ent/task"
	"sensei/ent/template"
	"sensei/ent/templatetask"
	"sensei/ent/user"
	"sensei/ent/verificationcode"
	"sensei/schema"
	"time"

	"github.com/google/uuid"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	activityMixin := schema.Activity{}.Mixin()
	activityMixinFields0 := activityMixin[0].Fields()
	_ = activityMixinFields0
	activityFields := schema.Activity{}.Fields()
	_ = activityFields
	// activityDescCreationDate is the schema descriptor for creationDate field.
	activityDescCreationDate := activityMixinFields0[1].Descriptor()
	// activity.DefaultCreationDate holds the default value on creation for the creationDate field.
	activity.DefaultCreationDate = activityDescCreationDate.Default.(func() time.Time)
	// activityDescName is the schema descriptor for name field.
	activityDescName := activityFields[0].Descriptor()
	// activity.NameValidator is a validator for the "name" field. It is called by the builders before save.
	activity.NameValidator = activityDescName.Validators[0].(func(string) error)
	// activityDescSize is the schema descriptor for size field.
	activityDescSize := activityFields[3].Descriptor()
	// activity.DefaultSize holds the default value on creation for the size field.
	activity.DefaultSize = activityDescSize.Default.(int)
	// activityDescID is the schema descriptor for id field.
	activityDescID := activityMixinFields0[0].Descriptor()
	// activity.DefaultID holds the default value on creation for the id field.
	activity.DefaultID = activityDescID.Default.(func() uuid.UUID)
	taskMixin := schema.Task{}.Mixin()
	taskMixinFields0 := taskMixin[0].Fields()
	_ = taskMixinFields0
	taskFields := schema.Task{}.Fields()
	_ = taskFields
	// taskDescCreationDate is the schema descriptor for creationDate field.
	taskDescCreationDate := taskMixinFields0[1].Descriptor()
	// task.DefaultCreationDate holds the default value on creation for the creationDate field.
	task.DefaultCreationDate = taskDescCreationDate.Default.(func() time.Time)
	// taskDescCompleted is the schema descriptor for completed field.
	taskDescCompleted := taskFields[1].Descriptor()
	// task.DefaultCompleted holds the default value on creation for the completed field.
	task.DefaultCompleted = taskDescCompleted.Default.(bool)
	// taskDescID is the schema descriptor for id field.
	taskDescID := taskMixinFields0[0].Descriptor()
	// task.DefaultID holds the default value on creation for the id field.
	task.DefaultID = taskDescID.Default.(func() uuid.UUID)
	templateMixin := schema.Template{}.Mixin()
	templateMixinFields0 := templateMixin[0].Fields()
	_ = templateMixinFields0
	templateFields := schema.Template{}.Fields()
	_ = templateFields
	// templateDescCreationDate is the schema descriptor for creationDate field.
	templateDescCreationDate := templateMixinFields0[1].Descriptor()
	// template.DefaultCreationDate holds the default value on creation for the creationDate field.
	template.DefaultCreationDate = templateDescCreationDate.Default.(func() time.Time)
	// templateDescID is the schema descriptor for id field.
	templateDescID := templateMixinFields0[0].Descriptor()
	// template.DefaultID holds the default value on creation for the id field.
	template.DefaultID = templateDescID.Default.(func() uuid.UUID)
	templatetaskMixin := schema.TemplateTask{}.Mixin()
	templatetaskMixinFields0 := templatetaskMixin[0].Fields()
	_ = templatetaskMixinFields0
	templatetaskFields := schema.TemplateTask{}.Fields()
	_ = templatetaskFields
	// templatetaskDescCreationDate is the schema descriptor for creationDate field.
	templatetaskDescCreationDate := templatetaskMixinFields0[1].Descriptor()
	// templatetask.DefaultCreationDate holds the default value on creation for the creationDate field.
	templatetask.DefaultCreationDate = templatetaskDescCreationDate.Default.(func() time.Time)
	// templatetaskDescID is the schema descriptor for id field.
	templatetaskDescID := templatetaskMixinFields0[0].Descriptor()
	// templatetask.DefaultID holds the default value on creation for the id field.
	templatetask.DefaultID = templatetaskDescID.Default.(func() uuid.UUID)
	userMixin := schema.User{}.Mixin()
	userMixinFields0 := userMixin[0].Fields()
	_ = userMixinFields0
	userFields := schema.User{}.Fields()
	_ = userFields
	// userDescCreationDate is the schema descriptor for creationDate field.
	userDescCreationDate := userMixinFields0[1].Descriptor()
	// user.DefaultCreationDate holds the default value on creation for the creationDate field.
	user.DefaultCreationDate = userDescCreationDate.Default.(func() time.Time)
	// userDescUsername is the schema descriptor for Username field.
	userDescUsername := userFields[0].Descriptor()
	// user.UsernameValidator is a validator for the "Username" field. It is called by the builders before save.
	user.UsernameValidator = userDescUsername.Validators[0].(func(string) error)
	// userDescMail is the schema descriptor for Mail field.
	userDescMail := userFields[1].Descriptor()
	// user.MailValidator is a validator for the "Mail" field. It is called by the builders before save.
	user.MailValidator = userDescMail.Validators[0].(func(string) error)
	// userDescPassword is the schema descriptor for Password field.
	userDescPassword := userFields[2].Descriptor()
	// user.PasswordValidator is a validator for the "Password" field. It is called by the builders before save.
	user.PasswordValidator = userDescPassword.Validators[0].(func(string) error)
	// userDescDans is the schema descriptor for Dans field.
	userDescDans := userFields[3].Descriptor()
	// user.DefaultDans holds the default value on creation for the Dans field.
	user.DefaultDans = userDescDans.Default.(int)
	// userDescMailValid is the schema descriptor for MailValid field.
	userDescMailValid := userFields[4].Descriptor()
	// user.DefaultMailValid holds the default value on creation for the MailValid field.
	user.DefaultMailValid = userDescMailValid.Default.(bool)
	// userDescID is the schema descriptor for id field.
	userDescID := userMixinFields0[0].Descriptor()
	// user.DefaultID holds the default value on creation for the id field.
	user.DefaultID = userDescID.Default.(func() uuid.UUID)
	verificationcodeMixin := schema.VerificationCode{}.Mixin()
	verificationcodeMixinFields0 := verificationcodeMixin[0].Fields()
	_ = verificationcodeMixinFields0
	verificationcodeFields := schema.VerificationCode{}.Fields()
	_ = verificationcodeFields
	// verificationcodeDescCreationDate is the schema descriptor for creationDate field.
	verificationcodeDescCreationDate := verificationcodeMixinFields0[1].Descriptor()
	// verificationcode.DefaultCreationDate holds the default value on creation for the creationDate field.
	verificationcode.DefaultCreationDate = verificationcodeDescCreationDate.Default.(func() time.Time)
	// verificationcodeDescType is the schema descriptor for type field.
	verificationcodeDescType := verificationcodeFields[0].Descriptor()
	// verificationcode.TypeValidator is a validator for the "type" field. It is called by the builders before save.
	verificationcode.TypeValidator = verificationcodeDescType.Validators[0].(func(string) error)
	// verificationcodeDescCode is the schema descriptor for code field.
	verificationcodeDescCode := verificationcodeFields[1].Descriptor()
	// verificationcode.CodeValidator is a validator for the "code" field. It is called by the builders before save.
	verificationcode.CodeValidator = verificationcodeDescCode.Validators[0].(func(string) error)
	// verificationcodeDescID is the schema descriptor for id field.
	verificationcodeDescID := verificationcodeMixinFields0[0].Descriptor()
	// verificationcode.DefaultID holds the default value on creation for the id field.
	verificationcode.DefaultID = verificationcodeDescID.Default.(func() uuid.UUID)
}
