// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/guacsec/guac/pkg/assembler/backends/ent/certifylegal"
	"github.com/guacsec/guac/pkg/assembler/backends/ent/license"
	"github.com/guacsec/guac/pkg/assembler/backends/ent/predicate"
)

// LicenseUpdate is the builder for updating License entities.
type LicenseUpdate struct {
	config
	hooks    []Hook
	mutation *LicenseMutation
}

// Where appends a list predicates to the LicenseUpdate builder.
func (lu *LicenseUpdate) Where(ps ...predicate.License) *LicenseUpdate {
	lu.mutation.Where(ps...)
	return lu
}

// SetName sets the "name" field.
func (lu *LicenseUpdate) SetName(s string) *LicenseUpdate {
	lu.mutation.SetName(s)
	return lu
}

// SetNillableName sets the "name" field if the given value is not nil.
func (lu *LicenseUpdate) SetNillableName(s *string) *LicenseUpdate {
	if s != nil {
		lu.SetName(*s)
	}
	return lu
}

// SetInline sets the "inline" field.
func (lu *LicenseUpdate) SetInline(s string) *LicenseUpdate {
	lu.mutation.SetInline(s)
	return lu
}

// SetNillableInline sets the "inline" field if the given value is not nil.
func (lu *LicenseUpdate) SetNillableInline(s *string) *LicenseUpdate {
	if s != nil {
		lu.SetInline(*s)
	}
	return lu
}

// ClearInline clears the value of the "inline" field.
func (lu *LicenseUpdate) ClearInline() *LicenseUpdate {
	lu.mutation.ClearInline()
	return lu
}

// SetListVersion sets the "list_version" field.
func (lu *LicenseUpdate) SetListVersion(s string) *LicenseUpdate {
	lu.mutation.SetListVersion(s)
	return lu
}

// SetNillableListVersion sets the "list_version" field if the given value is not nil.
func (lu *LicenseUpdate) SetNillableListVersion(s *string) *LicenseUpdate {
	if s != nil {
		lu.SetListVersion(*s)
	}
	return lu
}

// ClearListVersion clears the value of the "list_version" field.
func (lu *LicenseUpdate) ClearListVersion() *LicenseUpdate {
	lu.mutation.ClearListVersion()
	return lu
}

// AddDeclaredInCertifyLegalIDs adds the "declared_in_certify_legals" edge to the CertifyLegal entity by IDs.
func (lu *LicenseUpdate) AddDeclaredInCertifyLegalIDs(ids ...int) *LicenseUpdate {
	lu.mutation.AddDeclaredInCertifyLegalIDs(ids...)
	return lu
}

// AddDeclaredInCertifyLegals adds the "declared_in_certify_legals" edges to the CertifyLegal entity.
func (lu *LicenseUpdate) AddDeclaredInCertifyLegals(c ...*CertifyLegal) *LicenseUpdate {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return lu.AddDeclaredInCertifyLegalIDs(ids...)
}

// AddDiscoveredInCertifyLegalIDs adds the "discovered_in_certify_legals" edge to the CertifyLegal entity by IDs.
func (lu *LicenseUpdate) AddDiscoveredInCertifyLegalIDs(ids ...int) *LicenseUpdate {
	lu.mutation.AddDiscoveredInCertifyLegalIDs(ids...)
	return lu
}

// AddDiscoveredInCertifyLegals adds the "discovered_in_certify_legals" edges to the CertifyLegal entity.
func (lu *LicenseUpdate) AddDiscoveredInCertifyLegals(c ...*CertifyLegal) *LicenseUpdate {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return lu.AddDiscoveredInCertifyLegalIDs(ids...)
}

// Mutation returns the LicenseMutation object of the builder.
func (lu *LicenseUpdate) Mutation() *LicenseMutation {
	return lu.mutation
}

// ClearDeclaredInCertifyLegals clears all "declared_in_certify_legals" edges to the CertifyLegal entity.
func (lu *LicenseUpdate) ClearDeclaredInCertifyLegals() *LicenseUpdate {
	lu.mutation.ClearDeclaredInCertifyLegals()
	return lu
}

// RemoveDeclaredInCertifyLegalIDs removes the "declared_in_certify_legals" edge to CertifyLegal entities by IDs.
func (lu *LicenseUpdate) RemoveDeclaredInCertifyLegalIDs(ids ...int) *LicenseUpdate {
	lu.mutation.RemoveDeclaredInCertifyLegalIDs(ids...)
	return lu
}

// RemoveDeclaredInCertifyLegals removes "declared_in_certify_legals" edges to CertifyLegal entities.
func (lu *LicenseUpdate) RemoveDeclaredInCertifyLegals(c ...*CertifyLegal) *LicenseUpdate {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return lu.RemoveDeclaredInCertifyLegalIDs(ids...)
}

// ClearDiscoveredInCertifyLegals clears all "discovered_in_certify_legals" edges to the CertifyLegal entity.
func (lu *LicenseUpdate) ClearDiscoveredInCertifyLegals() *LicenseUpdate {
	lu.mutation.ClearDiscoveredInCertifyLegals()
	return lu
}

// RemoveDiscoveredInCertifyLegalIDs removes the "discovered_in_certify_legals" edge to CertifyLegal entities by IDs.
func (lu *LicenseUpdate) RemoveDiscoveredInCertifyLegalIDs(ids ...int) *LicenseUpdate {
	lu.mutation.RemoveDiscoveredInCertifyLegalIDs(ids...)
	return lu
}

// RemoveDiscoveredInCertifyLegals removes "discovered_in_certify_legals" edges to CertifyLegal entities.
func (lu *LicenseUpdate) RemoveDiscoveredInCertifyLegals(c ...*CertifyLegal) *LicenseUpdate {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return lu.RemoveDiscoveredInCertifyLegalIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (lu *LicenseUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, lu.sqlSave, lu.mutation, lu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (lu *LicenseUpdate) SaveX(ctx context.Context) int {
	affected, err := lu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (lu *LicenseUpdate) Exec(ctx context.Context) error {
	_, err := lu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (lu *LicenseUpdate) ExecX(ctx context.Context) {
	if err := lu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (lu *LicenseUpdate) check() error {
	if v, ok := lu.mutation.Name(); ok {
		if err := license.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "License.name": %w`, err)}
		}
	}
	return nil
}

func (lu *LicenseUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := lu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(license.Table, license.Columns, sqlgraph.NewFieldSpec(license.FieldID, field.TypeInt))
	if ps := lu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := lu.mutation.Name(); ok {
		_spec.SetField(license.FieldName, field.TypeString, value)
	}
	if value, ok := lu.mutation.Inline(); ok {
		_spec.SetField(license.FieldInline, field.TypeString, value)
	}
	if lu.mutation.InlineCleared() {
		_spec.ClearField(license.FieldInline, field.TypeString)
	}
	if value, ok := lu.mutation.ListVersion(); ok {
		_spec.SetField(license.FieldListVersion, field.TypeString, value)
	}
	if lu.mutation.ListVersionCleared() {
		_spec.ClearField(license.FieldListVersion, field.TypeString)
	}
	if lu.mutation.DeclaredInCertifyLegalsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   license.DeclaredInCertifyLegalsTable,
			Columns: license.DeclaredInCertifyLegalsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(certifylegal.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := lu.mutation.RemovedDeclaredInCertifyLegalsIDs(); len(nodes) > 0 && !lu.mutation.DeclaredInCertifyLegalsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   license.DeclaredInCertifyLegalsTable,
			Columns: license.DeclaredInCertifyLegalsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(certifylegal.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := lu.mutation.DeclaredInCertifyLegalsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   license.DeclaredInCertifyLegalsTable,
			Columns: license.DeclaredInCertifyLegalsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(certifylegal.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if lu.mutation.DiscoveredInCertifyLegalsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   license.DiscoveredInCertifyLegalsTable,
			Columns: license.DiscoveredInCertifyLegalsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(certifylegal.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := lu.mutation.RemovedDiscoveredInCertifyLegalsIDs(); len(nodes) > 0 && !lu.mutation.DiscoveredInCertifyLegalsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   license.DiscoveredInCertifyLegalsTable,
			Columns: license.DiscoveredInCertifyLegalsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(certifylegal.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := lu.mutation.DiscoveredInCertifyLegalsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   license.DiscoveredInCertifyLegalsTable,
			Columns: license.DiscoveredInCertifyLegalsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(certifylegal.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, lu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{license.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	lu.mutation.done = true
	return n, nil
}

// LicenseUpdateOne is the builder for updating a single License entity.
type LicenseUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *LicenseMutation
}

// SetName sets the "name" field.
func (luo *LicenseUpdateOne) SetName(s string) *LicenseUpdateOne {
	luo.mutation.SetName(s)
	return luo
}

// SetNillableName sets the "name" field if the given value is not nil.
func (luo *LicenseUpdateOne) SetNillableName(s *string) *LicenseUpdateOne {
	if s != nil {
		luo.SetName(*s)
	}
	return luo
}

// SetInline sets the "inline" field.
func (luo *LicenseUpdateOne) SetInline(s string) *LicenseUpdateOne {
	luo.mutation.SetInline(s)
	return luo
}

// SetNillableInline sets the "inline" field if the given value is not nil.
func (luo *LicenseUpdateOne) SetNillableInline(s *string) *LicenseUpdateOne {
	if s != nil {
		luo.SetInline(*s)
	}
	return luo
}

// ClearInline clears the value of the "inline" field.
func (luo *LicenseUpdateOne) ClearInline() *LicenseUpdateOne {
	luo.mutation.ClearInline()
	return luo
}

// SetListVersion sets the "list_version" field.
func (luo *LicenseUpdateOne) SetListVersion(s string) *LicenseUpdateOne {
	luo.mutation.SetListVersion(s)
	return luo
}

// SetNillableListVersion sets the "list_version" field if the given value is not nil.
func (luo *LicenseUpdateOne) SetNillableListVersion(s *string) *LicenseUpdateOne {
	if s != nil {
		luo.SetListVersion(*s)
	}
	return luo
}

// ClearListVersion clears the value of the "list_version" field.
func (luo *LicenseUpdateOne) ClearListVersion() *LicenseUpdateOne {
	luo.mutation.ClearListVersion()
	return luo
}

// AddDeclaredInCertifyLegalIDs adds the "declared_in_certify_legals" edge to the CertifyLegal entity by IDs.
func (luo *LicenseUpdateOne) AddDeclaredInCertifyLegalIDs(ids ...int) *LicenseUpdateOne {
	luo.mutation.AddDeclaredInCertifyLegalIDs(ids...)
	return luo
}

// AddDeclaredInCertifyLegals adds the "declared_in_certify_legals" edges to the CertifyLegal entity.
func (luo *LicenseUpdateOne) AddDeclaredInCertifyLegals(c ...*CertifyLegal) *LicenseUpdateOne {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return luo.AddDeclaredInCertifyLegalIDs(ids...)
}

// AddDiscoveredInCertifyLegalIDs adds the "discovered_in_certify_legals" edge to the CertifyLegal entity by IDs.
func (luo *LicenseUpdateOne) AddDiscoveredInCertifyLegalIDs(ids ...int) *LicenseUpdateOne {
	luo.mutation.AddDiscoveredInCertifyLegalIDs(ids...)
	return luo
}

// AddDiscoveredInCertifyLegals adds the "discovered_in_certify_legals" edges to the CertifyLegal entity.
func (luo *LicenseUpdateOne) AddDiscoveredInCertifyLegals(c ...*CertifyLegal) *LicenseUpdateOne {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return luo.AddDiscoveredInCertifyLegalIDs(ids...)
}

// Mutation returns the LicenseMutation object of the builder.
func (luo *LicenseUpdateOne) Mutation() *LicenseMutation {
	return luo.mutation
}

// ClearDeclaredInCertifyLegals clears all "declared_in_certify_legals" edges to the CertifyLegal entity.
func (luo *LicenseUpdateOne) ClearDeclaredInCertifyLegals() *LicenseUpdateOne {
	luo.mutation.ClearDeclaredInCertifyLegals()
	return luo
}

// RemoveDeclaredInCertifyLegalIDs removes the "declared_in_certify_legals" edge to CertifyLegal entities by IDs.
func (luo *LicenseUpdateOne) RemoveDeclaredInCertifyLegalIDs(ids ...int) *LicenseUpdateOne {
	luo.mutation.RemoveDeclaredInCertifyLegalIDs(ids...)
	return luo
}

// RemoveDeclaredInCertifyLegals removes "declared_in_certify_legals" edges to CertifyLegal entities.
func (luo *LicenseUpdateOne) RemoveDeclaredInCertifyLegals(c ...*CertifyLegal) *LicenseUpdateOne {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return luo.RemoveDeclaredInCertifyLegalIDs(ids...)
}

// ClearDiscoveredInCertifyLegals clears all "discovered_in_certify_legals" edges to the CertifyLegal entity.
func (luo *LicenseUpdateOne) ClearDiscoveredInCertifyLegals() *LicenseUpdateOne {
	luo.mutation.ClearDiscoveredInCertifyLegals()
	return luo
}

// RemoveDiscoveredInCertifyLegalIDs removes the "discovered_in_certify_legals" edge to CertifyLegal entities by IDs.
func (luo *LicenseUpdateOne) RemoveDiscoveredInCertifyLegalIDs(ids ...int) *LicenseUpdateOne {
	luo.mutation.RemoveDiscoveredInCertifyLegalIDs(ids...)
	return luo
}

// RemoveDiscoveredInCertifyLegals removes "discovered_in_certify_legals" edges to CertifyLegal entities.
func (luo *LicenseUpdateOne) RemoveDiscoveredInCertifyLegals(c ...*CertifyLegal) *LicenseUpdateOne {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return luo.RemoveDiscoveredInCertifyLegalIDs(ids...)
}

// Where appends a list predicates to the LicenseUpdate builder.
func (luo *LicenseUpdateOne) Where(ps ...predicate.License) *LicenseUpdateOne {
	luo.mutation.Where(ps...)
	return luo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (luo *LicenseUpdateOne) Select(field string, fields ...string) *LicenseUpdateOne {
	luo.fields = append([]string{field}, fields...)
	return luo
}

// Save executes the query and returns the updated License entity.
func (luo *LicenseUpdateOne) Save(ctx context.Context) (*License, error) {
	return withHooks(ctx, luo.sqlSave, luo.mutation, luo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (luo *LicenseUpdateOne) SaveX(ctx context.Context) *License {
	node, err := luo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (luo *LicenseUpdateOne) Exec(ctx context.Context) error {
	_, err := luo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (luo *LicenseUpdateOne) ExecX(ctx context.Context) {
	if err := luo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (luo *LicenseUpdateOne) check() error {
	if v, ok := luo.mutation.Name(); ok {
		if err := license.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "License.name": %w`, err)}
		}
	}
	return nil
}

func (luo *LicenseUpdateOne) sqlSave(ctx context.Context) (_node *License, err error) {
	if err := luo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(license.Table, license.Columns, sqlgraph.NewFieldSpec(license.FieldID, field.TypeInt))
	id, ok := luo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "License.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := luo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, license.FieldID)
		for _, f := range fields {
			if !license.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != license.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := luo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := luo.mutation.Name(); ok {
		_spec.SetField(license.FieldName, field.TypeString, value)
	}
	if value, ok := luo.mutation.Inline(); ok {
		_spec.SetField(license.FieldInline, field.TypeString, value)
	}
	if luo.mutation.InlineCleared() {
		_spec.ClearField(license.FieldInline, field.TypeString)
	}
	if value, ok := luo.mutation.ListVersion(); ok {
		_spec.SetField(license.FieldListVersion, field.TypeString, value)
	}
	if luo.mutation.ListVersionCleared() {
		_spec.ClearField(license.FieldListVersion, field.TypeString)
	}
	if luo.mutation.DeclaredInCertifyLegalsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   license.DeclaredInCertifyLegalsTable,
			Columns: license.DeclaredInCertifyLegalsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(certifylegal.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := luo.mutation.RemovedDeclaredInCertifyLegalsIDs(); len(nodes) > 0 && !luo.mutation.DeclaredInCertifyLegalsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   license.DeclaredInCertifyLegalsTable,
			Columns: license.DeclaredInCertifyLegalsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(certifylegal.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := luo.mutation.DeclaredInCertifyLegalsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   license.DeclaredInCertifyLegalsTable,
			Columns: license.DeclaredInCertifyLegalsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(certifylegal.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if luo.mutation.DiscoveredInCertifyLegalsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   license.DiscoveredInCertifyLegalsTable,
			Columns: license.DiscoveredInCertifyLegalsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(certifylegal.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := luo.mutation.RemovedDiscoveredInCertifyLegalsIDs(); len(nodes) > 0 && !luo.mutation.DiscoveredInCertifyLegalsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   license.DiscoveredInCertifyLegalsTable,
			Columns: license.DiscoveredInCertifyLegalsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(certifylegal.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := luo.mutation.DiscoveredInCertifyLegalsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   license.DiscoveredInCertifyLegalsTable,
			Columns: license.DiscoveredInCertifyLegalsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(certifylegal.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &License{config: luo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, luo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{license.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	luo.mutation.done = true
	return _node, nil
}
