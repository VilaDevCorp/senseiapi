// Code generated by ent, DO NOT EDIT.

package templatetask

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the templatetask type in the database.
	Label = "template_task"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreationDate holds the string denoting the creationdate field in the database.
	FieldCreationDate = "creation_date"
	// FieldWeekDay holds the string denoting the weekday field in the database.
	FieldWeekDay = "week_day"
	// EdgeActivity holds the string denoting the activity edge name in mutations.
	EdgeActivity = "activity"
	// EdgeTemplate holds the string denoting the template edge name in mutations.
	EdgeTemplate = "template"
	// Table holds the table name of the templatetask in the database.
	Table = "template_tasks"
	// ActivityTable is the table that holds the activity relation/edge.
	ActivityTable = "template_tasks"
	// ActivityInverseTable is the table name for the Activity entity.
	// It exists in this package in order to avoid circular dependency with the "activity" package.
	ActivityInverseTable = "activities"
	// ActivityColumn is the table column denoting the activity relation/edge.
	ActivityColumn = "activity_template_tasks"
	// TemplateTable is the table that holds the template relation/edge.
	TemplateTable = "template_tasks"
	// TemplateInverseTable is the table name for the Template entity.
	// It exists in this package in order to avoid circular dependency with the "template" package.
	TemplateInverseTable = "templates"
	// TemplateColumn is the table column denoting the template relation/edge.
	TemplateColumn = "template_template_tasks"
)

// Columns holds all SQL columns for templatetask fields.
var Columns = []string{
	FieldID,
	FieldCreationDate,
	FieldWeekDay,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "template_tasks"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"activity_template_tasks",
	"template_template_tasks",
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultCreationDate holds the default value on creation for the "creationDate" field.
	DefaultCreationDate func() time.Time
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() uuid.UUID
)

// OrderOption defines the ordering options for the TemplateTask queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByCreationDate orders the results by the creationDate field.
func ByCreationDate(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreationDate, opts...).ToFunc()
}

// ByWeekDay orders the results by the weekDay field.
func ByWeekDay(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldWeekDay, opts...).ToFunc()
}

// ByActivityField orders the results by activity field.
func ByActivityField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newActivityStep(), sql.OrderByField(field, opts...))
	}
}

// ByTemplateField orders the results by template field.
func ByTemplateField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newTemplateStep(), sql.OrderByField(field, opts...))
	}
}
func newActivityStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(ActivityInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, ActivityTable, ActivityColumn),
	)
}
func newTemplateStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(TemplateInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, TemplateTable, TemplateColumn),
	)
}
