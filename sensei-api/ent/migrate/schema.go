// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// ActivitiesColumns holds the columns for the "activities" table.
	ActivitiesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "creation_date", Type: field.TypeTime},
		{Name: "name", Type: field.TypeString},
		{Name: "description", Type: field.TypeString, Nullable: true},
		{Name: "icon", Type: field.TypeString, Nullable: true},
		{Name: "size", Type: field.TypeInt, Default: 1},
		{Name: "user_activities", Type: field.TypeUUID},
	}
	// ActivitiesTable holds the schema information for the "activities" table.
	ActivitiesTable = &schema.Table{
		Name:       "activities",
		Columns:    ActivitiesColumns,
		PrimaryKey: []*schema.Column{ActivitiesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "activities_users_activities",
				Columns:    []*schema.Column{ActivitiesColumns[6]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// TasksColumns holds the columns for the "tasks" table.
	TasksColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "creation_date", Type: field.TypeTime},
		{Name: "due_date", Type: field.TypeTime},
		{Name: "completed", Type: field.TypeBool, Default: false},
		{Name: "activity_tasks", Type: field.TypeUUID},
		{Name: "user_tasks", Type: field.TypeUUID},
	}
	// TasksTable holds the schema information for the "tasks" table.
	TasksTable = &schema.Table{
		Name:       "tasks",
		Columns:    TasksColumns,
		PrimaryKey: []*schema.Column{TasksColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "tasks_activities_tasks",
				Columns:    []*schema.Column{TasksColumns[4]},
				RefColumns: []*schema.Column{ActivitiesColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "tasks_users_tasks",
				Columns:    []*schema.Column{TasksColumns[5]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// TemplatesColumns holds the columns for the "templates" table.
	TemplatesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "creation_date", Type: field.TypeTime},
		{Name: "name", Type: field.TypeString},
		{Name: "user_templates", Type: field.TypeUUID},
	}
	// TemplatesTable holds the schema information for the "templates" table.
	TemplatesTable = &schema.Table{
		Name:       "templates",
		Columns:    TemplatesColumns,
		PrimaryKey: []*schema.Column{TemplatesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "templates_users_templates",
				Columns:    []*schema.Column{TemplatesColumns[3]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// TemplateTasksColumns holds the columns for the "template_tasks" table.
	TemplateTasksColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "creation_date", Type: field.TypeTime},
		{Name: "week_day", Type: field.TypeInt},
		{Name: "activity_template_tasks", Type: field.TypeUUID},
		{Name: "template_template_tasks", Type: field.TypeUUID},
	}
	// TemplateTasksTable holds the schema information for the "template_tasks" table.
	TemplateTasksTable = &schema.Table{
		Name:       "template_tasks",
		Columns:    TemplateTasksColumns,
		PrimaryKey: []*schema.Column{TemplateTasksColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "template_tasks_activities_templateTasks",
				Columns:    []*schema.Column{TemplateTasksColumns[3]},
				RefColumns: []*schema.Column{ActivitiesColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "template_tasks_templates_templateTasks",
				Columns:    []*schema.Column{TemplateTasksColumns[4]},
				RefColumns: []*schema.Column{TemplatesColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// UsersColumns holds the columns for the "users" table.
	UsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "creation_date", Type: field.TypeTime},
		{Name: "username", Type: field.TypeString, Unique: true},
		{Name: "mail", Type: field.TypeString, Unique: true},
		{Name: "password", Type: field.TypeString},
		{Name: "dans", Type: field.TypeInt, Default: 0},
		{Name: "mail_valid", Type: field.TypeBool, Default: false},
		{Name: "tutorial_completed", Type: field.TypeBool, Default: false},
	}
	// UsersTable holds the schema information for the "users" table.
	UsersTable = &schema.Table{
		Name:       "users",
		Columns:    UsersColumns,
		PrimaryKey: []*schema.Column{UsersColumns[0]},
	}
	// VerificationCodesColumns holds the columns for the "verification_codes" table.
	VerificationCodesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "creation_date", Type: field.TypeTime},
		{Name: "type", Type: field.TypeString},
		{Name: "code", Type: field.TypeString},
		{Name: "expire_date", Type: field.TypeTime},
		{Name: "valid", Type: field.TypeBool},
		{Name: "user_codes", Type: field.TypeUUID},
	}
	// VerificationCodesTable holds the schema information for the "verification_codes" table.
	VerificationCodesTable = &schema.Table{
		Name:       "verification_codes",
		Columns:    VerificationCodesColumns,
		PrimaryKey: []*schema.Column{VerificationCodesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "verification_codes_users_codes",
				Columns:    []*schema.Column{VerificationCodesColumns[6]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		ActivitiesTable,
		TasksTable,
		TemplatesTable,
		TemplateTasksTable,
		UsersTable,
		VerificationCodesTable,
	}
)

func init() {
	ActivitiesTable.ForeignKeys[0].RefTable = UsersTable
	TasksTable.ForeignKeys[0].RefTable = ActivitiesTable
	TasksTable.ForeignKeys[1].RefTable = UsersTable
	TemplatesTable.ForeignKeys[0].RefTable = UsersTable
	TemplateTasksTable.ForeignKeys[0].RefTable = ActivitiesTable
	TemplateTasksTable.ForeignKeys[1].RefTable = TemplatesTable
	VerificationCodesTable.ForeignKeys[0].RefTable = UsersTable
	VerificationCodesTable.Annotation = &entsql.Annotation{
		Table: "verification_codes",
	}
}
