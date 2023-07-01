// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"
	"sensei/ent/activity"
	"sensei/ent/predicate"
	"sensei/ent/template"
	"sensei/ent/templatetask"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// TemplateTaskQuery is the builder for querying TemplateTask entities.
type TemplateTaskQuery struct {
	config
	ctx          *QueryContext
	order        []templatetask.OrderOption
	inters       []Interceptor
	predicates   []predicate.TemplateTask
	withActivity *ActivityQuery
	withTemplate *TemplateQuery
	withFKs      bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the TemplateTaskQuery builder.
func (ttq *TemplateTaskQuery) Where(ps ...predicate.TemplateTask) *TemplateTaskQuery {
	ttq.predicates = append(ttq.predicates, ps...)
	return ttq
}

// Limit the number of records to be returned by this query.
func (ttq *TemplateTaskQuery) Limit(limit int) *TemplateTaskQuery {
	ttq.ctx.Limit = &limit
	return ttq
}

// Offset to start from.
func (ttq *TemplateTaskQuery) Offset(offset int) *TemplateTaskQuery {
	ttq.ctx.Offset = &offset
	return ttq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (ttq *TemplateTaskQuery) Unique(unique bool) *TemplateTaskQuery {
	ttq.ctx.Unique = &unique
	return ttq
}

// Order specifies how the records should be ordered.
func (ttq *TemplateTaskQuery) Order(o ...templatetask.OrderOption) *TemplateTaskQuery {
	ttq.order = append(ttq.order, o...)
	return ttq
}

// QueryActivity chains the current query on the "activity" edge.
func (ttq *TemplateTaskQuery) QueryActivity() *ActivityQuery {
	query := (&ActivityClient{config: ttq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := ttq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := ttq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(templatetask.Table, templatetask.FieldID, selector),
			sqlgraph.To(activity.Table, activity.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, templatetask.ActivityTable, templatetask.ActivityColumn),
		)
		fromU = sqlgraph.SetNeighbors(ttq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryTemplate chains the current query on the "template" edge.
func (ttq *TemplateTaskQuery) QueryTemplate() *TemplateQuery {
	query := (&TemplateClient{config: ttq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := ttq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := ttq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(templatetask.Table, templatetask.FieldID, selector),
			sqlgraph.To(template.Table, template.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, templatetask.TemplateTable, templatetask.TemplateColumn),
		)
		fromU = sqlgraph.SetNeighbors(ttq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first TemplateTask entity from the query.
// Returns a *NotFoundError when no TemplateTask was found.
func (ttq *TemplateTaskQuery) First(ctx context.Context) (*TemplateTask, error) {
	nodes, err := ttq.Limit(1).All(setContextOp(ctx, ttq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{templatetask.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (ttq *TemplateTaskQuery) FirstX(ctx context.Context) *TemplateTask {
	node, err := ttq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first TemplateTask ID from the query.
// Returns a *NotFoundError when no TemplateTask ID was found.
func (ttq *TemplateTaskQuery) FirstID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = ttq.Limit(1).IDs(setContextOp(ctx, ttq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{templatetask.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (ttq *TemplateTaskQuery) FirstIDX(ctx context.Context) uuid.UUID {
	id, err := ttq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single TemplateTask entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one TemplateTask entity is found.
// Returns a *NotFoundError when no TemplateTask entities are found.
func (ttq *TemplateTaskQuery) Only(ctx context.Context) (*TemplateTask, error) {
	nodes, err := ttq.Limit(2).All(setContextOp(ctx, ttq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{templatetask.Label}
	default:
		return nil, &NotSingularError{templatetask.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (ttq *TemplateTaskQuery) OnlyX(ctx context.Context) *TemplateTask {
	node, err := ttq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only TemplateTask ID in the query.
// Returns a *NotSingularError when more than one TemplateTask ID is found.
// Returns a *NotFoundError when no entities are found.
func (ttq *TemplateTaskQuery) OnlyID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = ttq.Limit(2).IDs(setContextOp(ctx, ttq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{templatetask.Label}
	default:
		err = &NotSingularError{templatetask.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (ttq *TemplateTaskQuery) OnlyIDX(ctx context.Context) uuid.UUID {
	id, err := ttq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of TemplateTasks.
func (ttq *TemplateTaskQuery) All(ctx context.Context) ([]*TemplateTask, error) {
	ctx = setContextOp(ctx, ttq.ctx, "All")
	if err := ttq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*TemplateTask, *TemplateTaskQuery]()
	return withInterceptors[[]*TemplateTask](ctx, ttq, qr, ttq.inters)
}

// AllX is like All, but panics if an error occurs.
func (ttq *TemplateTaskQuery) AllX(ctx context.Context) []*TemplateTask {
	nodes, err := ttq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of TemplateTask IDs.
func (ttq *TemplateTaskQuery) IDs(ctx context.Context) (ids []uuid.UUID, err error) {
	if ttq.ctx.Unique == nil && ttq.path != nil {
		ttq.Unique(true)
	}
	ctx = setContextOp(ctx, ttq.ctx, "IDs")
	if err = ttq.Select(templatetask.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (ttq *TemplateTaskQuery) IDsX(ctx context.Context) []uuid.UUID {
	ids, err := ttq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (ttq *TemplateTaskQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, ttq.ctx, "Count")
	if err := ttq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, ttq, querierCount[*TemplateTaskQuery](), ttq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (ttq *TemplateTaskQuery) CountX(ctx context.Context) int {
	count, err := ttq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (ttq *TemplateTaskQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, ttq.ctx, "Exist")
	switch _, err := ttq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (ttq *TemplateTaskQuery) ExistX(ctx context.Context) bool {
	exist, err := ttq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the TemplateTaskQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (ttq *TemplateTaskQuery) Clone() *TemplateTaskQuery {
	if ttq == nil {
		return nil
	}
	return &TemplateTaskQuery{
		config:       ttq.config,
		ctx:          ttq.ctx.Clone(),
		order:        append([]templatetask.OrderOption{}, ttq.order...),
		inters:       append([]Interceptor{}, ttq.inters...),
		predicates:   append([]predicate.TemplateTask{}, ttq.predicates...),
		withActivity: ttq.withActivity.Clone(),
		withTemplate: ttq.withTemplate.Clone(),
		// clone intermediate query.
		sql:  ttq.sql.Clone(),
		path: ttq.path,
	}
}

// WithActivity tells the query-builder to eager-load the nodes that are connected to
// the "activity" edge. The optional arguments are used to configure the query builder of the edge.
func (ttq *TemplateTaskQuery) WithActivity(opts ...func(*ActivityQuery)) *TemplateTaskQuery {
	query := (&ActivityClient{config: ttq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	ttq.withActivity = query
	return ttq
}

// WithTemplate tells the query-builder to eager-load the nodes that are connected to
// the "template" edge. The optional arguments are used to configure the query builder of the edge.
func (ttq *TemplateTaskQuery) WithTemplate(opts ...func(*TemplateQuery)) *TemplateTaskQuery {
	query := (&TemplateClient{config: ttq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	ttq.withTemplate = query
	return ttq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		CreationDate time.Time `json:"creationDate,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.TemplateTask.Query().
//		GroupBy(templatetask.FieldCreationDate).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (ttq *TemplateTaskQuery) GroupBy(field string, fields ...string) *TemplateTaskGroupBy {
	ttq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &TemplateTaskGroupBy{build: ttq}
	grbuild.flds = &ttq.ctx.Fields
	grbuild.label = templatetask.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		CreationDate time.Time `json:"creationDate,omitempty"`
//	}
//
//	client.TemplateTask.Query().
//		Select(templatetask.FieldCreationDate).
//		Scan(ctx, &v)
func (ttq *TemplateTaskQuery) Select(fields ...string) *TemplateTaskSelect {
	ttq.ctx.Fields = append(ttq.ctx.Fields, fields...)
	sbuild := &TemplateTaskSelect{TemplateTaskQuery: ttq}
	sbuild.label = templatetask.Label
	sbuild.flds, sbuild.scan = &ttq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a TemplateTaskSelect configured with the given aggregations.
func (ttq *TemplateTaskQuery) Aggregate(fns ...AggregateFunc) *TemplateTaskSelect {
	return ttq.Select().Aggregate(fns...)
}

func (ttq *TemplateTaskQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range ttq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, ttq); err != nil {
				return err
			}
		}
	}
	for _, f := range ttq.ctx.Fields {
		if !templatetask.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if ttq.path != nil {
		prev, err := ttq.path(ctx)
		if err != nil {
			return err
		}
		ttq.sql = prev
	}
	return nil
}

func (ttq *TemplateTaskQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*TemplateTask, error) {
	var (
		nodes       = []*TemplateTask{}
		withFKs     = ttq.withFKs
		_spec       = ttq.querySpec()
		loadedTypes = [2]bool{
			ttq.withActivity != nil,
			ttq.withTemplate != nil,
		}
	)
	if ttq.withActivity != nil || ttq.withTemplate != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, templatetask.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*TemplateTask).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &TemplateTask{config: ttq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, ttq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := ttq.withActivity; query != nil {
		if err := ttq.loadActivity(ctx, query, nodes, nil,
			func(n *TemplateTask, e *Activity) { n.Edges.Activity = e }); err != nil {
			return nil, err
		}
	}
	if query := ttq.withTemplate; query != nil {
		if err := ttq.loadTemplate(ctx, query, nodes, nil,
			func(n *TemplateTask, e *Template) { n.Edges.Template = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (ttq *TemplateTaskQuery) loadActivity(ctx context.Context, query *ActivityQuery, nodes []*TemplateTask, init func(*TemplateTask), assign func(*TemplateTask, *Activity)) error {
	ids := make([]uuid.UUID, 0, len(nodes))
	nodeids := make(map[uuid.UUID][]*TemplateTask)
	for i := range nodes {
		if nodes[i].activity_template_tasks == nil {
			continue
		}
		fk := *nodes[i].activity_template_tasks
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(activity.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "activity_template_tasks" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (ttq *TemplateTaskQuery) loadTemplate(ctx context.Context, query *TemplateQuery, nodes []*TemplateTask, init func(*TemplateTask), assign func(*TemplateTask, *Template)) error {
	ids := make([]uuid.UUID, 0, len(nodes))
	nodeids := make(map[uuid.UUID][]*TemplateTask)
	for i := range nodes {
		if nodes[i].template_template_tasks == nil {
			continue
		}
		fk := *nodes[i].template_template_tasks
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(template.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "template_template_tasks" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (ttq *TemplateTaskQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := ttq.querySpec()
	_spec.Node.Columns = ttq.ctx.Fields
	if len(ttq.ctx.Fields) > 0 {
		_spec.Unique = ttq.ctx.Unique != nil && *ttq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, ttq.driver, _spec)
}

func (ttq *TemplateTaskQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(templatetask.Table, templatetask.Columns, sqlgraph.NewFieldSpec(templatetask.FieldID, field.TypeUUID))
	_spec.From = ttq.sql
	if unique := ttq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if ttq.path != nil {
		_spec.Unique = true
	}
	if fields := ttq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, templatetask.FieldID)
		for i := range fields {
			if fields[i] != templatetask.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := ttq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := ttq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := ttq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := ttq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (ttq *TemplateTaskQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(ttq.driver.Dialect())
	t1 := builder.Table(templatetask.Table)
	columns := ttq.ctx.Fields
	if len(columns) == 0 {
		columns = templatetask.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if ttq.sql != nil {
		selector = ttq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if ttq.ctx.Unique != nil && *ttq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range ttq.predicates {
		p(selector)
	}
	for _, p := range ttq.order {
		p(selector)
	}
	if offset := ttq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := ttq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// TemplateTaskGroupBy is the group-by builder for TemplateTask entities.
type TemplateTaskGroupBy struct {
	selector
	build *TemplateTaskQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (ttgb *TemplateTaskGroupBy) Aggregate(fns ...AggregateFunc) *TemplateTaskGroupBy {
	ttgb.fns = append(ttgb.fns, fns...)
	return ttgb
}

// Scan applies the selector query and scans the result into the given value.
func (ttgb *TemplateTaskGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ttgb.build.ctx, "GroupBy")
	if err := ttgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*TemplateTaskQuery, *TemplateTaskGroupBy](ctx, ttgb.build, ttgb, ttgb.build.inters, v)
}

func (ttgb *TemplateTaskGroupBy) sqlScan(ctx context.Context, root *TemplateTaskQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(ttgb.fns))
	for _, fn := range ttgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*ttgb.flds)+len(ttgb.fns))
		for _, f := range *ttgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*ttgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ttgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// TemplateTaskSelect is the builder for selecting fields of TemplateTask entities.
type TemplateTaskSelect struct {
	*TemplateTaskQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (tts *TemplateTaskSelect) Aggregate(fns ...AggregateFunc) *TemplateTaskSelect {
	tts.fns = append(tts.fns, fns...)
	return tts
}

// Scan applies the selector query and scans the result into the given value.
func (tts *TemplateTaskSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, tts.ctx, "Select")
	if err := tts.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*TemplateTaskQuery, *TemplateTaskSelect](ctx, tts.TemplateTaskQuery, tts, tts.inters, v)
}

func (tts *TemplateTaskSelect) sqlScan(ctx context.Context, root *TemplateTaskQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(tts.fns))
	for _, fn := range tts.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*tts.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := tts.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
