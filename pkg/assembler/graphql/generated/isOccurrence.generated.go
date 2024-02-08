// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package generated

import (
	"context"
	"errors"
	"fmt"
	"strconv"
	"sync"
	"sync/atomic"

	"github.com/99designs/gqlgen/graphql"
	"github.com/guacsec/guac/pkg/assembler/graphql/model"
	"github.com/vektah/gqlparser/v2/ast"
)

// region    ************************** generated!.gotpl **************************

// endregion ************************** generated!.gotpl **************************

// region    ***************************** args.gotpl *****************************

// endregion ***************************** args.gotpl *****************************

// region    ************************** directives.gotpl **************************

// endregion ************************** directives.gotpl **************************

// region    **************************** field.gotpl *****************************

func (ec *executionContext) _IsOccurrence_id(ctx context.Context, field graphql.CollectedField, obj *model.IsOccurrence) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_IsOccurrence_id(ctx, field)
	if err != nil {
		return graphql.Null
	}
	ctx = graphql.WithFieldContext(ctx, fc)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	resTmp := ec._fieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.ID, nil
	})

	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(string)
	fc.Result = res
	return ec.marshalNID2string(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext_IsOccurrence_id(ctx context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "IsOccurrence",
		Field:      field,
		IsMethod:   false,
		IsResolver: false,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			return nil, errors.New("field of type ID does not have child fields")
		},
	}
	return fc, nil
}

func (ec *executionContext) _IsOccurrence_subject(ctx context.Context, field graphql.CollectedField, obj *model.IsOccurrence) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_IsOccurrence_subject(ctx, field)
	if err != nil {
		return graphql.Null
	}
	ctx = graphql.WithFieldContext(ctx, fc)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	resTmp := ec._fieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.Subject, nil
	})

	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(model.PackageOrSource)
	fc.Result = res
	return ec.marshalNPackageOrSource2githubᚗcomᚋguacsecᚋguacᚋpkgᚋassemblerᚋgraphqlᚋmodelᚐPackageOrSource(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext_IsOccurrence_subject(ctx context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "IsOccurrence",
		Field:      field,
		IsMethod:   false,
		IsResolver: false,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			return nil, errors.New("field of type PackageOrSource does not have child fields")
		},
	}
	return fc, nil
}

func (ec *executionContext) _IsOccurrence_artifact(ctx context.Context, field graphql.CollectedField, obj *model.IsOccurrence) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_IsOccurrence_artifact(ctx, field)
	if err != nil {
		return graphql.Null
	}
	ctx = graphql.WithFieldContext(ctx, fc)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	resTmp := ec._fieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.Artifact, nil
	})

	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(*model.Artifact)
	fc.Result = res
	return ec.marshalNArtifact2ᚖgithubᚗcomᚋguacsecᚋguacᚋpkgᚋassemblerᚋgraphqlᚋmodelᚐArtifact(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext_IsOccurrence_artifact(ctx context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "IsOccurrence",
		Field:      field,
		IsMethod:   false,
		IsResolver: false,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			switch field.Name {
			case "id":
				return ec.fieldContext_Artifact_id(ctx, field)
			case "algorithm":
				return ec.fieldContext_Artifact_algorithm(ctx, field)
			case "digest":
				return ec.fieldContext_Artifact_digest(ctx, field)
			}
			return nil, fmt.Errorf("no field named %q was found under type Artifact", field.Name)
		},
	}
	return fc, nil
}

func (ec *executionContext) _IsOccurrence_justification(ctx context.Context, field graphql.CollectedField, obj *model.IsOccurrence) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_IsOccurrence_justification(ctx, field)
	if err != nil {
		return graphql.Null
	}
	ctx = graphql.WithFieldContext(ctx, fc)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	resTmp := ec._fieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.Justification, nil
	})

	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(string)
	fc.Result = res
	return ec.marshalNString2string(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext_IsOccurrence_justification(ctx context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "IsOccurrence",
		Field:      field,
		IsMethod:   false,
		IsResolver: false,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			return nil, errors.New("field of type String does not have child fields")
		},
	}
	return fc, nil
}

func (ec *executionContext) _IsOccurrence_origin(ctx context.Context, field graphql.CollectedField, obj *model.IsOccurrence) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_IsOccurrence_origin(ctx, field)
	if err != nil {
		return graphql.Null
	}
	ctx = graphql.WithFieldContext(ctx, fc)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	resTmp := ec._fieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.Origin, nil
	})

	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(string)
	fc.Result = res
	return ec.marshalNString2string(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext_IsOccurrence_origin(ctx context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "IsOccurrence",
		Field:      field,
		IsMethod:   false,
		IsResolver: false,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			return nil, errors.New("field of type String does not have child fields")
		},
	}
	return fc, nil
}

func (ec *executionContext) _IsOccurrence_collector(ctx context.Context, field graphql.CollectedField, obj *model.IsOccurrence) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_IsOccurrence_collector(ctx, field)
	if err != nil {
		return graphql.Null
	}
	ctx = graphql.WithFieldContext(ctx, fc)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	resTmp := ec._fieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.Collector, nil
	})

	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(string)
	fc.Result = res
	return ec.marshalNString2string(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext_IsOccurrence_collector(ctx context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "IsOccurrence",
		Field:      field,
		IsMethod:   false,
		IsResolver: false,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			return nil, errors.New("field of type String does not have child fields")
		},
	}
	return fc, nil
}

// endregion **************************** field.gotpl *****************************

// region    **************************** input.gotpl *****************************

func (ec *executionContext) unmarshalInputIsOccurrenceInputSpec(ctx context.Context, obj interface{}) (model.IsOccurrenceInputSpec, error) {
	var it model.IsOccurrenceInputSpec
	asMap := map[string]interface{}{}
	for k, v := range obj.(map[string]interface{}) {
		asMap[k] = v
	}

	fieldsInOrder := [...]string{"justification", "origin", "collector"}
	for _, k := range fieldsInOrder {
		v, ok := asMap[k]
		if !ok {
			continue
		}
		switch k {
		case "justification":
			ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("justification"))
			data, err := ec.unmarshalNString2string(ctx, v)
			if err != nil {
				return it, err
			}
			it.Justification = data
		case "origin":
			ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("origin"))
			data, err := ec.unmarshalNString2string(ctx, v)
			if err != nil {
				return it, err
			}
			it.Origin = data
		case "collector":
			ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("collector"))
			data, err := ec.unmarshalNString2string(ctx, v)
			if err != nil {
				return it, err
			}
			it.Collector = data
		}
	}

	return it, nil
}

func (ec *executionContext) unmarshalInputIsOccurrenceSpec(ctx context.Context, obj interface{}) (model.IsOccurrenceSpec, error) {
	var it model.IsOccurrenceSpec
	asMap := map[string]interface{}{}
	for k, v := range obj.(map[string]interface{}) {
		asMap[k] = v
	}

	fieldsInOrder := [...]string{"id", "subject", "artifact", "justification", "origin", "collector"}
	for _, k := range fieldsInOrder {
		v, ok := asMap[k]
		if !ok {
			continue
		}
		switch k {
		case "id":
			ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("id"))
			data, err := ec.unmarshalOID2ᚖstring(ctx, v)
			if err != nil {
				return it, err
			}
			it.ID = data
		case "subject":
			ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("subject"))
			data, err := ec.unmarshalOPackageOrSourceSpec2ᚖgithubᚗcomᚋguacsecᚋguacᚋpkgᚋassemblerᚋgraphqlᚋmodelᚐPackageOrSourceSpec(ctx, v)
			if err != nil {
				return it, err
			}
			it.Subject = data
		case "artifact":
			ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("artifact"))
			data, err := ec.unmarshalOArtifactSpec2ᚖgithubᚗcomᚋguacsecᚋguacᚋpkgᚋassemblerᚋgraphqlᚋmodelᚐArtifactSpec(ctx, v)
			if err != nil {
				return it, err
			}
			it.Artifact = data
		case "justification":
			ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("justification"))
			data, err := ec.unmarshalOString2ᚖstring(ctx, v)
			if err != nil {
				return it, err
			}
			it.Justification = data
		case "origin":
			ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("origin"))
			data, err := ec.unmarshalOString2ᚖstring(ctx, v)
			if err != nil {
				return it, err
			}
			it.Origin = data
		case "collector":
			ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("collector"))
			data, err := ec.unmarshalOString2ᚖstring(ctx, v)
			if err != nil {
				return it, err
			}
			it.Collector = data
		}
	}

	return it, nil
}

func (ec *executionContext) unmarshalInputPackageOrSourceInput(ctx context.Context, obj interface{}) (model.PackageOrSourceInput, error) {
	var it model.PackageOrSourceInput
	asMap := map[string]interface{}{}
	for k, v := range obj.(map[string]interface{}) {
		asMap[k] = v
	}

	fieldsInOrder := [...]string{"package", "source"}
	for _, k := range fieldsInOrder {
		v, ok := asMap[k]
		if !ok {
			continue
		}
		switch k {
		case "package":
			ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("package"))
			data, err := ec.unmarshalOPkgInputSpec2ᚖgithubᚗcomᚋguacsecᚋguacᚋpkgᚋassemblerᚋgraphqlᚋmodelᚐPkgInputSpec(ctx, v)
			if err != nil {
				return it, err
			}
			it.Package = data
		case "source":
			ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("source"))
			data, err := ec.unmarshalOSourceInputSpec2ᚖgithubᚗcomᚋguacsecᚋguacᚋpkgᚋassemblerᚋgraphqlᚋmodelᚐSourceInputSpec(ctx, v)
			if err != nil {
				return it, err
			}
			it.Source = data
		}
	}

	return it, nil
}

func (ec *executionContext) unmarshalInputPackageOrSourceInputs(ctx context.Context, obj interface{}) (model.PackageOrSourceInputs, error) {
	var it model.PackageOrSourceInputs
	asMap := map[string]interface{}{}
	for k, v := range obj.(map[string]interface{}) {
		asMap[k] = v
	}

	fieldsInOrder := [...]string{"packages", "sources"}
	for _, k := range fieldsInOrder {
		v, ok := asMap[k]
		if !ok {
			continue
		}
		switch k {
		case "packages":
			ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("packages"))
			data, err := ec.unmarshalOPkgInputSpec2ᚕᚖgithubᚗcomᚋguacsecᚋguacᚋpkgᚋassemblerᚋgraphqlᚋmodelᚐPkgInputSpecᚄ(ctx, v)
			if err != nil {
				return it, err
			}
			it.Packages = data
		case "sources":
			ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("sources"))
			data, err := ec.unmarshalOSourceInputSpec2ᚕᚖgithubᚗcomᚋguacsecᚋguacᚋpkgᚋassemblerᚋgraphqlᚋmodelᚐSourceInputSpecᚄ(ctx, v)
			if err != nil {
				return it, err
			}
			it.Sources = data
		}
	}

	return it, nil
}

func (ec *executionContext) unmarshalInputPackageOrSourceSpec(ctx context.Context, obj interface{}) (model.PackageOrSourceSpec, error) {
	var it model.PackageOrSourceSpec
	asMap := map[string]interface{}{}
	for k, v := range obj.(map[string]interface{}) {
		asMap[k] = v
	}

	fieldsInOrder := [...]string{"package", "source"}
	for _, k := range fieldsInOrder {
		v, ok := asMap[k]
		if !ok {
			continue
		}
		switch k {
		case "package":
			ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("package"))
			data, err := ec.unmarshalOPkgSpec2ᚖgithubᚗcomᚋguacsecᚋguacᚋpkgᚋassemblerᚋgraphqlᚋmodelᚐPkgSpec(ctx, v)
			if err != nil {
				return it, err
			}
			it.Package = data
		case "source":
			ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("source"))
			data, err := ec.unmarshalOSourceSpec2ᚖgithubᚗcomᚋguacsecᚋguacᚋpkgᚋassemblerᚋgraphqlᚋmodelᚐSourceSpec(ctx, v)
			if err != nil {
				return it, err
			}
			it.Source = data
		}
	}

	return it, nil
}

// endregion **************************** input.gotpl *****************************

// region    ************************** interface.gotpl ***************************

func (ec *executionContext) _PackageOrSource(ctx context.Context, sel ast.SelectionSet, obj model.PackageOrSource) graphql.Marshaler {
	switch obj := (obj).(type) {
	case nil:
		return graphql.Null
	case model.Package:
		return ec._Package(ctx, sel, &obj)
	case *model.Package:
		if obj == nil {
			return graphql.Null
		}
		return ec._Package(ctx, sel, obj)
	case model.Source:
		return ec._Source(ctx, sel, &obj)
	case *model.Source:
		if obj == nil {
			return graphql.Null
		}
		return ec._Source(ctx, sel, obj)
	default:
		panic(fmt.Errorf("unexpected type %T", obj))
	}
}

// endregion ************************** interface.gotpl ***************************

// region    **************************** object.gotpl ****************************

var isOccurrenceImplementors = []string{"IsOccurrence", "Node"}

func (ec *executionContext) _IsOccurrence(ctx context.Context, sel ast.SelectionSet, obj *model.IsOccurrence) graphql.Marshaler {
	fields := graphql.CollectFields(ec.OperationContext, sel, isOccurrenceImplementors)

	out := graphql.NewFieldSet(fields)
	deferred := make(map[string]*graphql.FieldSet)
	for i, field := range fields {
		switch field.Name {
		case "__typename":
			out.Values[i] = graphql.MarshalString("IsOccurrence")
		case "id":
			out.Values[i] = ec._IsOccurrence_id(ctx, field, obj)
			if out.Values[i] == graphql.Null {
				out.Invalids++
			}
		case "subject":
			out.Values[i] = ec._IsOccurrence_subject(ctx, field, obj)
			if out.Values[i] == graphql.Null {
				out.Invalids++
			}
		case "artifact":
			out.Values[i] = ec._IsOccurrence_artifact(ctx, field, obj)
			if out.Values[i] == graphql.Null {
				out.Invalids++
			}
		case "justification":
			out.Values[i] = ec._IsOccurrence_justification(ctx, field, obj)
			if out.Values[i] == graphql.Null {
				out.Invalids++
			}
		case "origin":
			out.Values[i] = ec._IsOccurrence_origin(ctx, field, obj)
			if out.Values[i] == graphql.Null {
				out.Invalids++
			}
		case "collector":
			out.Values[i] = ec._IsOccurrence_collector(ctx, field, obj)
			if out.Values[i] == graphql.Null {
				out.Invalids++
			}
		default:
			panic("unknown field " + strconv.Quote(field.Name))
		}
	}
	out.Dispatch(ctx)
	if out.Invalids > 0 {
		return graphql.Null
	}

	atomic.AddInt32(&ec.deferred, int32(len(deferred)))

	for label, dfs := range deferred {
		ec.processDeferredGroup(graphql.DeferredGroup{
			Label:    label,
			Path:     graphql.GetPath(ctx),
			FieldSet: dfs,
			Context:  ctx,
		})
	}

	return out
}

// endregion **************************** object.gotpl ****************************

// region    ***************************** type.gotpl *****************************

func (ec *executionContext) marshalNIsOccurrence2ᚕᚖgithubᚗcomᚋguacsecᚋguacᚋpkgᚋassemblerᚋgraphqlᚋmodelᚐIsOccurrenceᚄ(ctx context.Context, sel ast.SelectionSet, v []*model.IsOccurrence) graphql.Marshaler {
	ret := make(graphql.Array, len(v))
	var wg sync.WaitGroup
	isLen1 := len(v) == 1
	if !isLen1 {
		wg.Add(len(v))
	}
	for i := range v {
		i := i
		fc := &graphql.FieldContext{
			Index:  &i,
			Result: &v[i],
		}
		ctx := graphql.WithFieldContext(ctx, fc)
		f := func(i int) {
			defer func() {
				if r := recover(); r != nil {
					ec.Error(ctx, ec.Recover(ctx, r))
					ret = nil
				}
			}()
			if !isLen1 {
				defer wg.Done()
			}
			ret[i] = ec.marshalNIsOccurrence2ᚖgithubᚗcomᚋguacsecᚋguacᚋpkgᚋassemblerᚋgraphqlᚋmodelᚐIsOccurrence(ctx, sel, v[i])
		}
		if isLen1 {
			f(i)
		} else {
			go f(i)
		}

	}
	wg.Wait()

	for _, e := range ret {
		if e == graphql.Null {
			return graphql.Null
		}
	}

	return ret
}

func (ec *executionContext) marshalNIsOccurrence2ᚖgithubᚗcomᚋguacsecᚋguacᚋpkgᚋassemblerᚋgraphqlᚋmodelᚐIsOccurrence(ctx context.Context, sel ast.SelectionSet, v *model.IsOccurrence) graphql.Marshaler {
	if v == nil {
		if !graphql.HasFieldError(ctx, graphql.GetFieldContext(ctx)) {
			ec.Errorf(ctx, "the requested element is null which the schema does not allow")
		}
		return graphql.Null
	}
	return ec._IsOccurrence(ctx, sel, v)
}

func (ec *executionContext) unmarshalNIsOccurrenceInputSpec2githubᚗcomᚋguacsecᚋguacᚋpkgᚋassemblerᚋgraphqlᚋmodelᚐIsOccurrenceInputSpec(ctx context.Context, v interface{}) (model.IsOccurrenceInputSpec, error) {
	res, err := ec.unmarshalInputIsOccurrenceInputSpec(ctx, v)
	return res, graphql.ErrorOnPath(ctx, err)
}

func (ec *executionContext) unmarshalNIsOccurrenceInputSpec2ᚕᚖgithubᚗcomᚋguacsecᚋguacᚋpkgᚋassemblerᚋgraphqlᚋmodelᚐIsOccurrenceInputSpecᚄ(ctx context.Context, v interface{}) ([]*model.IsOccurrenceInputSpec, error) {
	var vSlice []interface{}
	if v != nil {
		vSlice = graphql.CoerceList(v)
	}
	var err error
	res := make([]*model.IsOccurrenceInputSpec, len(vSlice))
	for i := range vSlice {
		ctx := graphql.WithPathContext(ctx, graphql.NewPathWithIndex(i))
		res[i], err = ec.unmarshalNIsOccurrenceInputSpec2ᚖgithubᚗcomᚋguacsecᚋguacᚋpkgᚋassemblerᚋgraphqlᚋmodelᚐIsOccurrenceInputSpec(ctx, vSlice[i])
		if err != nil {
			return nil, err
		}
	}
	return res, nil
}

func (ec *executionContext) unmarshalNIsOccurrenceInputSpec2ᚖgithubᚗcomᚋguacsecᚋguacᚋpkgᚋassemblerᚋgraphqlᚋmodelᚐIsOccurrenceInputSpec(ctx context.Context, v interface{}) (*model.IsOccurrenceInputSpec, error) {
	res, err := ec.unmarshalInputIsOccurrenceInputSpec(ctx, v)
	return &res, graphql.ErrorOnPath(ctx, err)
}

func (ec *executionContext) unmarshalNIsOccurrenceSpec2githubᚗcomᚋguacsecᚋguacᚋpkgᚋassemblerᚋgraphqlᚋmodelᚐIsOccurrenceSpec(ctx context.Context, v interface{}) (model.IsOccurrenceSpec, error) {
	res, err := ec.unmarshalInputIsOccurrenceSpec(ctx, v)
	return res, graphql.ErrorOnPath(ctx, err)
}

func (ec *executionContext) marshalNPackageOrSource2githubᚗcomᚋguacsecᚋguacᚋpkgᚋassemblerᚋgraphqlᚋmodelᚐPackageOrSource(ctx context.Context, sel ast.SelectionSet, v model.PackageOrSource) graphql.Marshaler {
	if v == nil {
		if !graphql.HasFieldError(ctx, graphql.GetFieldContext(ctx)) {
			ec.Errorf(ctx, "the requested element is null which the schema does not allow")
		}
		return graphql.Null
	}
	return ec._PackageOrSource(ctx, sel, v)
}

func (ec *executionContext) unmarshalNPackageOrSourceInput2githubᚗcomᚋguacsecᚋguacᚋpkgᚋassemblerᚋgraphqlᚋmodelᚐPackageOrSourceInput(ctx context.Context, v interface{}) (model.PackageOrSourceInput, error) {
	res, err := ec.unmarshalInputPackageOrSourceInput(ctx, v)
	return res, graphql.ErrorOnPath(ctx, err)
}

func (ec *executionContext) unmarshalNPackageOrSourceInputs2githubᚗcomᚋguacsecᚋguacᚋpkgᚋassemblerᚋgraphqlᚋmodelᚐPackageOrSourceInputs(ctx context.Context, v interface{}) (model.PackageOrSourceInputs, error) {
	res, err := ec.unmarshalInputPackageOrSourceInputs(ctx, v)
	return res, graphql.ErrorOnPath(ctx, err)
}

func (ec *executionContext) unmarshalOPackageOrSourceSpec2ᚖgithubᚗcomᚋguacsecᚋguacᚋpkgᚋassemblerᚋgraphqlᚋmodelᚐPackageOrSourceSpec(ctx context.Context, v interface{}) (*model.PackageOrSourceSpec, error) {
	if v == nil {
		return nil, nil
	}
	res, err := ec.unmarshalInputPackageOrSourceSpec(ctx, v)
	return &res, graphql.ErrorOnPath(ctx, err)
}

// endregion ***************************** type.gotpl *****************************
