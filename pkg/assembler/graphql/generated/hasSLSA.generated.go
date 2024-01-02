// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package generated

import (
	"context"
	"errors"
	"fmt"
	"strconv"
	"sync"
	"sync/atomic"
	"time"

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

func (ec *executionContext) _HasSLSA_id(ctx context.Context, field graphql.CollectedField, obj *model.HasSlsa) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_HasSLSA_id(ctx, field)
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

func (ec *executionContext) fieldContext_HasSLSA_id(ctx context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "HasSLSA",
		Field:      field,
		IsMethod:   false,
		IsResolver: false,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			return nil, errors.New("field of type ID does not have child fields")
		},
	}
	return fc, nil
}

func (ec *executionContext) _HasSLSA_subject(ctx context.Context, field graphql.CollectedField, obj *model.HasSlsa) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_HasSLSA_subject(ctx, field)
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
	res := resTmp.(*model.Artifact)
	fc.Result = res
	return ec.marshalNArtifact2ᚖgithubᚗcomᚋguacsecᚋguacᚋpkgᚋassemblerᚋgraphqlᚋmodelᚐArtifact(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext_HasSLSA_subject(ctx context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "HasSLSA",
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

func (ec *executionContext) _HasSLSA_slsa(ctx context.Context, field graphql.CollectedField, obj *model.HasSlsa) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_HasSLSA_slsa(ctx, field)
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
		return obj.Slsa, nil
	})

	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(*model.Slsa)
	fc.Result = res
	return ec.marshalNSLSA2ᚖgithubᚗcomᚋguacsecᚋguacᚋpkgᚋassemblerᚋgraphqlᚋmodelᚐSlsa(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext_HasSLSA_slsa(ctx context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "HasSLSA",
		Field:      field,
		IsMethod:   false,
		IsResolver: false,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			switch field.Name {
			case "builtFrom":
				return ec.fieldContext_SLSA_builtFrom(ctx, field)
			case "builtBy":
				return ec.fieldContext_SLSA_builtBy(ctx, field)
			case "buildType":
				return ec.fieldContext_SLSA_buildType(ctx, field)
			case "slsaPredicate":
				return ec.fieldContext_SLSA_slsaPredicate(ctx, field)
			case "slsaVersion":
				return ec.fieldContext_SLSA_slsaVersion(ctx, field)
			case "startedOn":
				return ec.fieldContext_SLSA_startedOn(ctx, field)
			case "finishedOn":
				return ec.fieldContext_SLSA_finishedOn(ctx, field)
			case "origin":
				return ec.fieldContext_SLSA_origin(ctx, field)
			case "collector":
				return ec.fieldContext_SLSA_collector(ctx, field)
			}
			return nil, fmt.Errorf("no field named %q was found under type SLSA", field.Name)
		},
	}
	return fc, nil
}

func (ec *executionContext) _SLSA_builtFrom(ctx context.Context, field graphql.CollectedField, obj *model.Slsa) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_SLSA_builtFrom(ctx, field)
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
		return obj.BuiltFrom, nil
	})

	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.([]*model.Artifact)
	fc.Result = res
	return ec.marshalNArtifact2ᚕᚖgithubᚗcomᚋguacsecᚋguacᚋpkgᚋassemblerᚋgraphqlᚋmodelᚐArtifactᚄ(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext_SLSA_builtFrom(ctx context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "SLSA",
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

func (ec *executionContext) _SLSA_builtBy(ctx context.Context, field graphql.CollectedField, obj *model.Slsa) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_SLSA_builtBy(ctx, field)
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
		return obj.BuiltBy, nil
	})

	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(*model.Builder)
	fc.Result = res
	return ec.marshalNBuilder2ᚖgithubᚗcomᚋguacsecᚋguacᚋpkgᚋassemblerᚋgraphqlᚋmodelᚐBuilder(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext_SLSA_builtBy(ctx context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "SLSA",
		Field:      field,
		IsMethod:   false,
		IsResolver: false,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			switch field.Name {
			case "id":
				return ec.fieldContext_Builder_id(ctx, field)
			case "uri":
				return ec.fieldContext_Builder_uri(ctx, field)
			}
			return nil, fmt.Errorf("no field named %q was found under type Builder", field.Name)
		},
	}
	return fc, nil
}

func (ec *executionContext) _SLSA_buildType(ctx context.Context, field graphql.CollectedField, obj *model.Slsa) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_SLSA_buildType(ctx, field)
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
		return obj.BuildType, nil
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

func (ec *executionContext) fieldContext_SLSA_buildType(ctx context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "SLSA",
		Field:      field,
		IsMethod:   false,
		IsResolver: false,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			return nil, errors.New("field of type String does not have child fields")
		},
	}
	return fc, nil
}

func (ec *executionContext) _SLSA_slsaPredicate(ctx context.Context, field graphql.CollectedField, obj *model.Slsa) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_SLSA_slsaPredicate(ctx, field)
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
		return obj.SlsaPredicate, nil
	})

	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.([]*model.SLSAPredicate)
	fc.Result = res
	return ec.marshalNSLSAPredicate2ᚕᚖgithubᚗcomᚋguacsecᚋguacᚋpkgᚋassemblerᚋgraphqlᚋmodelᚐSLSAPredicateᚄ(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext_SLSA_slsaPredicate(ctx context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "SLSA",
		Field:      field,
		IsMethod:   false,
		IsResolver: false,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			switch field.Name {
			case "key":
				return ec.fieldContext_SLSAPredicate_key(ctx, field)
			case "value":
				return ec.fieldContext_SLSAPredicate_value(ctx, field)
			}
			return nil, fmt.Errorf("no field named %q was found under type SLSAPredicate", field.Name)
		},
	}
	return fc, nil
}

func (ec *executionContext) _SLSA_slsaVersion(ctx context.Context, field graphql.CollectedField, obj *model.Slsa) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_SLSA_slsaVersion(ctx, field)
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
		return obj.SlsaVersion, nil
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

func (ec *executionContext) fieldContext_SLSA_slsaVersion(ctx context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "SLSA",
		Field:      field,
		IsMethod:   false,
		IsResolver: false,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			return nil, errors.New("field of type String does not have child fields")
		},
	}
	return fc, nil
}

func (ec *executionContext) _SLSA_startedOn(ctx context.Context, field graphql.CollectedField, obj *model.Slsa) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_SLSA_startedOn(ctx, field)
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
		return obj.StartedOn, nil
	})

	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(*time.Time)
	fc.Result = res
	return ec.marshalOTime2ᚖtimeᚐTime(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext_SLSA_startedOn(ctx context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "SLSA",
		Field:      field,
		IsMethod:   false,
		IsResolver: false,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			return nil, errors.New("field of type Time does not have child fields")
		},
	}
	return fc, nil
}

func (ec *executionContext) _SLSA_finishedOn(ctx context.Context, field graphql.CollectedField, obj *model.Slsa) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_SLSA_finishedOn(ctx, field)
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
		return obj.FinishedOn, nil
	})

	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(*time.Time)
	fc.Result = res
	return ec.marshalOTime2ᚖtimeᚐTime(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext_SLSA_finishedOn(ctx context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "SLSA",
		Field:      field,
		IsMethod:   false,
		IsResolver: false,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			return nil, errors.New("field of type Time does not have child fields")
		},
	}
	return fc, nil
}

func (ec *executionContext) _SLSA_origin(ctx context.Context, field graphql.CollectedField, obj *model.Slsa) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_SLSA_origin(ctx, field)
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

func (ec *executionContext) fieldContext_SLSA_origin(ctx context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "SLSA",
		Field:      field,
		IsMethod:   false,
		IsResolver: false,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			return nil, errors.New("field of type String does not have child fields")
		},
	}
	return fc, nil
}

func (ec *executionContext) _SLSA_collector(ctx context.Context, field graphql.CollectedField, obj *model.Slsa) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_SLSA_collector(ctx, field)
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

func (ec *executionContext) fieldContext_SLSA_collector(ctx context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "SLSA",
		Field:      field,
		IsMethod:   false,
		IsResolver: false,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			return nil, errors.New("field of type String does not have child fields")
		},
	}
	return fc, nil
}

func (ec *executionContext) _SLSAPredicate_key(ctx context.Context, field graphql.CollectedField, obj *model.SLSAPredicate) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_SLSAPredicate_key(ctx, field)
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
		return obj.Key, nil
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

func (ec *executionContext) fieldContext_SLSAPredicate_key(ctx context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "SLSAPredicate",
		Field:      field,
		IsMethod:   false,
		IsResolver: false,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			return nil, errors.New("field of type String does not have child fields")
		},
	}
	return fc, nil
}

func (ec *executionContext) _SLSAPredicate_value(ctx context.Context, field graphql.CollectedField, obj *model.SLSAPredicate) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_SLSAPredicate_value(ctx, field)
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
		return obj.Value, nil
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

func (ec *executionContext) fieldContext_SLSAPredicate_value(ctx context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "SLSAPredicate",
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

func (ec *executionContext) unmarshalInputHasSLSASpec(ctx context.Context, obj interface{}) (model.HasSLSASpec, error) {
	var it model.HasSLSASpec
	asMap := map[string]interface{}{}
	for k, v := range obj.(map[string]interface{}) {
		asMap[k] = v
	}

	if _, present := asMap["predicate"]; !present {
		asMap["predicate"] = []interface{}{}
	}

	fieldsInOrder := [...]string{"id", "subject", "builtFrom", "builtBy", "buildType", "predicate", "slsaVersion", "startedOn", "finishedOn", "origin", "collector"}
	for _, k := range fieldsInOrder {
		v, ok := asMap[k]
		if !ok {
			continue
		}
		switch k {
		case "id":
			var err error

			ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("id"))
			data, err := ec.unmarshalOID2ᚖstring(ctx, v)
			if err != nil {
				return it, err
			}
			it.ID = data
		case "subject":
			var err error

			ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("subject"))
			data, err := ec.unmarshalOArtifactSpec2ᚖgithubᚗcomᚋguacsecᚋguacᚋpkgᚋassemblerᚋgraphqlᚋmodelᚐArtifactSpec(ctx, v)
			if err != nil {
				return it, err
			}
			it.Subject = data
		case "builtFrom":
			var err error

			ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("builtFrom"))
			data, err := ec.unmarshalOArtifactSpec2ᚕᚖgithubᚗcomᚋguacsecᚋguacᚋpkgᚋassemblerᚋgraphqlᚋmodelᚐArtifactSpecᚄ(ctx, v)
			if err != nil {
				return it, err
			}
			it.BuiltFrom = data
		case "builtBy":
			var err error

			ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("builtBy"))
			data, err := ec.unmarshalOBuilderSpec2ᚖgithubᚗcomᚋguacsecᚋguacᚋpkgᚋassemblerᚋgraphqlᚋmodelᚐBuilderSpec(ctx, v)
			if err != nil {
				return it, err
			}
			it.BuiltBy = data
		case "buildType":
			var err error

			ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("buildType"))
			data, err := ec.unmarshalOString2ᚖstring(ctx, v)
			if err != nil {
				return it, err
			}
			it.BuildType = data
		case "predicate":
			var err error

			ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("predicate"))
			data, err := ec.unmarshalOSLSAPredicateSpec2ᚕᚖgithubᚗcomᚋguacsecᚋguacᚋpkgᚋassemblerᚋgraphqlᚋmodelᚐSLSAPredicateSpecᚄ(ctx, v)
			if err != nil {
				return it, err
			}
			it.Predicate = data
		case "slsaVersion":
			var err error

			ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("slsaVersion"))
			data, err := ec.unmarshalOString2ᚖstring(ctx, v)
			if err != nil {
				return it, err
			}
			it.SlsaVersion = data
		case "startedOn":
			var err error

			ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("startedOn"))
			data, err := ec.unmarshalOTime2ᚖtimeᚐTime(ctx, v)
			if err != nil {
				return it, err
			}
			it.StartedOn = data
		case "finishedOn":
			var err error

			ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("finishedOn"))
			data, err := ec.unmarshalOTime2ᚖtimeᚐTime(ctx, v)
			if err != nil {
				return it, err
			}
			it.FinishedOn = data
		case "origin":
			var err error

			ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("origin"))
			data, err := ec.unmarshalOString2ᚖstring(ctx, v)
			if err != nil {
				return it, err
			}
			it.Origin = data
		case "collector":
			var err error

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

func (ec *executionContext) unmarshalInputSLSAInputSpec(ctx context.Context, obj interface{}) (model.SLSAInputSpec, error) {
	var it model.SLSAInputSpec
	asMap := map[string]interface{}{}
	for k, v := range obj.(map[string]interface{}) {
		asMap[k] = v
	}

	fieldsInOrder := [...]string{"buildType", "slsaPredicate", "slsaVersion", "startedOn", "finishedOn", "origin", "collector"}
	for _, k := range fieldsInOrder {
		v, ok := asMap[k]
		if !ok {
			continue
		}
		switch k {
		case "buildType":
			var err error

			ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("buildType"))
			data, err := ec.unmarshalNString2string(ctx, v)
			if err != nil {
				return it, err
			}
			it.BuildType = data
		case "slsaPredicate":
			var err error

			ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("slsaPredicate"))
			data, err := ec.unmarshalNSLSAPredicateInputSpec2ᚕᚖgithubᚗcomᚋguacsecᚋguacᚋpkgᚋassemblerᚋgraphqlᚋmodelᚐSLSAPredicateInputSpecᚄ(ctx, v)
			if err != nil {
				return it, err
			}
			it.SlsaPredicate = data
		case "slsaVersion":
			var err error

			ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("slsaVersion"))
			data, err := ec.unmarshalNString2string(ctx, v)
			if err != nil {
				return it, err
			}
			it.SlsaVersion = data
		case "startedOn":
			var err error

			ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("startedOn"))
			data, err := ec.unmarshalOTime2ᚖtimeᚐTime(ctx, v)
			if err != nil {
				return it, err
			}
			it.StartedOn = data
		case "finishedOn":
			var err error

			ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("finishedOn"))
			data, err := ec.unmarshalOTime2ᚖtimeᚐTime(ctx, v)
			if err != nil {
				return it, err
			}
			it.FinishedOn = data
		case "origin":
			var err error

			ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("origin"))
			data, err := ec.unmarshalNString2string(ctx, v)
			if err != nil {
				return it, err
			}
			it.Origin = data
		case "collector":
			var err error

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

func (ec *executionContext) unmarshalInputSLSAPredicateInputSpec(ctx context.Context, obj interface{}) (model.SLSAPredicateInputSpec, error) {
	var it model.SLSAPredicateInputSpec
	asMap := map[string]interface{}{}
	for k, v := range obj.(map[string]interface{}) {
		asMap[k] = v
	}

	fieldsInOrder := [...]string{"key", "value"}
	for _, k := range fieldsInOrder {
		v, ok := asMap[k]
		if !ok {
			continue
		}
		switch k {
		case "key":
			var err error

			ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("key"))
			data, err := ec.unmarshalNString2string(ctx, v)
			if err != nil {
				return it, err
			}
			it.Key = data
		case "value":
			var err error

			ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("value"))
			data, err := ec.unmarshalNString2string(ctx, v)
			if err != nil {
				return it, err
			}
			it.Value = data
		}
	}

	return it, nil
}

func (ec *executionContext) unmarshalInputSLSAPredicateSpec(ctx context.Context, obj interface{}) (model.SLSAPredicateSpec, error) {
	var it model.SLSAPredicateSpec
	asMap := map[string]interface{}{}
	for k, v := range obj.(map[string]interface{}) {
		asMap[k] = v
	}

	fieldsInOrder := [...]string{"key", "value"}
	for _, k := range fieldsInOrder {
		v, ok := asMap[k]
		if !ok {
			continue
		}
		switch k {
		case "key":
			var err error

			ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("key"))
			data, err := ec.unmarshalNString2string(ctx, v)
			if err != nil {
				return it, err
			}
			it.Key = data
		case "value":
			var err error

			ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("value"))
			data, err := ec.unmarshalNString2string(ctx, v)
			if err != nil {
				return it, err
			}
			it.Value = data
		}
	}

	return it, nil
}

// endregion **************************** input.gotpl *****************************

// region    ************************** interface.gotpl ***************************

// endregion ************************** interface.gotpl ***************************

// region    **************************** object.gotpl ****************************

var hasSLSAImplementors = []string{"HasSLSA", "Node"}

func (ec *executionContext) _HasSLSA(ctx context.Context, sel ast.SelectionSet, obj *model.HasSlsa) graphql.Marshaler {
	fields := graphql.CollectFields(ec.OperationContext, sel, hasSLSAImplementors)

	out := graphql.NewFieldSet(fields)
	deferred := make(map[string]*graphql.FieldSet)
	for i, field := range fields {
		switch field.Name {
		case "__typename":
			out.Values[i] = graphql.MarshalString("HasSLSA")
		case "id":
			out.Values[i] = ec._HasSLSA_id(ctx, field, obj)
			if out.Values[i] == graphql.Null {
				out.Invalids++
			}
		case "subject":
			out.Values[i] = ec._HasSLSA_subject(ctx, field, obj)
			if out.Values[i] == graphql.Null {
				out.Invalids++
			}
		case "slsa":
			out.Values[i] = ec._HasSLSA_slsa(ctx, field, obj)
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

var sLSAImplementors = []string{"SLSA"}

func (ec *executionContext) _SLSA(ctx context.Context, sel ast.SelectionSet, obj *model.Slsa) graphql.Marshaler {
	fields := graphql.CollectFields(ec.OperationContext, sel, sLSAImplementors)

	out := graphql.NewFieldSet(fields)
	deferred := make(map[string]*graphql.FieldSet)
	for i, field := range fields {
		switch field.Name {
		case "__typename":
			out.Values[i] = graphql.MarshalString("SLSA")
		case "builtFrom":
			out.Values[i] = ec._SLSA_builtFrom(ctx, field, obj)
			if out.Values[i] == graphql.Null {
				out.Invalids++
			}
		case "builtBy":
			out.Values[i] = ec._SLSA_builtBy(ctx, field, obj)
			if out.Values[i] == graphql.Null {
				out.Invalids++
			}
		case "buildType":
			out.Values[i] = ec._SLSA_buildType(ctx, field, obj)
			if out.Values[i] == graphql.Null {
				out.Invalids++
			}
		case "slsaPredicate":
			out.Values[i] = ec._SLSA_slsaPredicate(ctx, field, obj)
			if out.Values[i] == graphql.Null {
				out.Invalids++
			}
		case "slsaVersion":
			out.Values[i] = ec._SLSA_slsaVersion(ctx, field, obj)
			if out.Values[i] == graphql.Null {
				out.Invalids++
			}
		case "startedOn":
			out.Values[i] = ec._SLSA_startedOn(ctx, field, obj)
		case "finishedOn":
			out.Values[i] = ec._SLSA_finishedOn(ctx, field, obj)
		case "origin":
			out.Values[i] = ec._SLSA_origin(ctx, field, obj)
			if out.Values[i] == graphql.Null {
				out.Invalids++
			}
		case "collector":
			out.Values[i] = ec._SLSA_collector(ctx, field, obj)
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

var sLSAPredicateImplementors = []string{"SLSAPredicate"}

func (ec *executionContext) _SLSAPredicate(ctx context.Context, sel ast.SelectionSet, obj *model.SLSAPredicate) graphql.Marshaler {
	fields := graphql.CollectFields(ec.OperationContext, sel, sLSAPredicateImplementors)

	out := graphql.NewFieldSet(fields)
	deferred := make(map[string]*graphql.FieldSet)
	for i, field := range fields {
		switch field.Name {
		case "__typename":
			out.Values[i] = graphql.MarshalString("SLSAPredicate")
		case "key":
			out.Values[i] = ec._SLSAPredicate_key(ctx, field, obj)
			if out.Values[i] == graphql.Null {
				out.Invalids++
			}
		case "value":
			out.Values[i] = ec._SLSAPredicate_value(ctx, field, obj)
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

func (ec *executionContext) marshalNHasSLSA2ᚕᚖgithubᚗcomᚋguacsecᚋguacᚋpkgᚋassemblerᚋgraphqlᚋmodelᚐHasSlsaᚄ(ctx context.Context, sel ast.SelectionSet, v []*model.HasSlsa) graphql.Marshaler {
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
			ret[i] = ec.marshalNHasSLSA2ᚖgithubᚗcomᚋguacsecᚋguacᚋpkgᚋassemblerᚋgraphqlᚋmodelᚐHasSlsa(ctx, sel, v[i])
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

func (ec *executionContext) marshalNHasSLSA2ᚖgithubᚗcomᚋguacsecᚋguacᚋpkgᚋassemblerᚋgraphqlᚋmodelᚐHasSlsa(ctx context.Context, sel ast.SelectionSet, v *model.HasSlsa) graphql.Marshaler {
	if v == nil {
		if !graphql.HasFieldError(ctx, graphql.GetFieldContext(ctx)) {
			ec.Errorf(ctx, "the requested element is null which the schema does not allow")
		}
		return graphql.Null
	}
	return ec._HasSLSA(ctx, sel, v)
}

func (ec *executionContext) unmarshalNHasSLSASpec2githubᚗcomᚋguacsecᚋguacᚋpkgᚋassemblerᚋgraphqlᚋmodelᚐHasSLSASpec(ctx context.Context, v interface{}) (model.HasSLSASpec, error) {
	res, err := ec.unmarshalInputHasSLSASpec(ctx, v)
	return res, graphql.ErrorOnPath(ctx, err)
}

func (ec *executionContext) marshalNSLSA2ᚖgithubᚗcomᚋguacsecᚋguacᚋpkgᚋassemblerᚋgraphqlᚋmodelᚐSlsa(ctx context.Context, sel ast.SelectionSet, v *model.Slsa) graphql.Marshaler {
	if v == nil {
		if !graphql.HasFieldError(ctx, graphql.GetFieldContext(ctx)) {
			ec.Errorf(ctx, "the requested element is null which the schema does not allow")
		}
		return graphql.Null
	}
	return ec._SLSA(ctx, sel, v)
}

func (ec *executionContext) unmarshalNSLSAInputSpec2githubᚗcomᚋguacsecᚋguacᚋpkgᚋassemblerᚋgraphqlᚋmodelᚐSLSAInputSpec(ctx context.Context, v interface{}) (model.SLSAInputSpec, error) {
	res, err := ec.unmarshalInputSLSAInputSpec(ctx, v)
	return res, graphql.ErrorOnPath(ctx, err)
}

func (ec *executionContext) unmarshalNSLSAInputSpec2ᚕᚖgithubᚗcomᚋguacsecᚋguacᚋpkgᚋassemblerᚋgraphqlᚋmodelᚐSLSAInputSpecᚄ(ctx context.Context, v interface{}) ([]*model.SLSAInputSpec, error) {
	var vSlice []interface{}
	if v != nil {
		vSlice = graphql.CoerceList(v)
	}
	var err error
	res := make([]*model.SLSAInputSpec, len(vSlice))
	for i := range vSlice {
		ctx := graphql.WithPathContext(ctx, graphql.NewPathWithIndex(i))
		res[i], err = ec.unmarshalNSLSAInputSpec2ᚖgithubᚗcomᚋguacsecᚋguacᚋpkgᚋassemblerᚋgraphqlᚋmodelᚐSLSAInputSpec(ctx, vSlice[i])
		if err != nil {
			return nil, err
		}
	}
	return res, nil
}

func (ec *executionContext) unmarshalNSLSAInputSpec2ᚖgithubᚗcomᚋguacsecᚋguacᚋpkgᚋassemblerᚋgraphqlᚋmodelᚐSLSAInputSpec(ctx context.Context, v interface{}) (*model.SLSAInputSpec, error) {
	res, err := ec.unmarshalInputSLSAInputSpec(ctx, v)
	return &res, graphql.ErrorOnPath(ctx, err)
}

func (ec *executionContext) marshalNSLSAPredicate2ᚕᚖgithubᚗcomᚋguacsecᚋguacᚋpkgᚋassemblerᚋgraphqlᚋmodelᚐSLSAPredicateᚄ(ctx context.Context, sel ast.SelectionSet, v []*model.SLSAPredicate) graphql.Marshaler {
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
			ret[i] = ec.marshalNSLSAPredicate2ᚖgithubᚗcomᚋguacsecᚋguacᚋpkgᚋassemblerᚋgraphqlᚋmodelᚐSLSAPredicate(ctx, sel, v[i])
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

func (ec *executionContext) marshalNSLSAPredicate2ᚖgithubᚗcomᚋguacsecᚋguacᚋpkgᚋassemblerᚋgraphqlᚋmodelᚐSLSAPredicate(ctx context.Context, sel ast.SelectionSet, v *model.SLSAPredicate) graphql.Marshaler {
	if v == nil {
		if !graphql.HasFieldError(ctx, graphql.GetFieldContext(ctx)) {
			ec.Errorf(ctx, "the requested element is null which the schema does not allow")
		}
		return graphql.Null
	}
	return ec._SLSAPredicate(ctx, sel, v)
}

func (ec *executionContext) unmarshalNSLSAPredicateInputSpec2ᚕᚖgithubᚗcomᚋguacsecᚋguacᚋpkgᚋassemblerᚋgraphqlᚋmodelᚐSLSAPredicateInputSpecᚄ(ctx context.Context, v interface{}) ([]*model.SLSAPredicateInputSpec, error) {
	var vSlice []interface{}
	if v != nil {
		vSlice = graphql.CoerceList(v)
	}
	var err error
	res := make([]*model.SLSAPredicateInputSpec, len(vSlice))
	for i := range vSlice {
		ctx := graphql.WithPathContext(ctx, graphql.NewPathWithIndex(i))
		res[i], err = ec.unmarshalNSLSAPredicateInputSpec2ᚖgithubᚗcomᚋguacsecᚋguacᚋpkgᚋassemblerᚋgraphqlᚋmodelᚐSLSAPredicateInputSpec(ctx, vSlice[i])
		if err != nil {
			return nil, err
		}
	}
	return res, nil
}

func (ec *executionContext) unmarshalNSLSAPredicateInputSpec2ᚖgithubᚗcomᚋguacsecᚋguacᚋpkgᚋassemblerᚋgraphqlᚋmodelᚐSLSAPredicateInputSpec(ctx context.Context, v interface{}) (*model.SLSAPredicateInputSpec, error) {
	res, err := ec.unmarshalInputSLSAPredicateInputSpec(ctx, v)
	return &res, graphql.ErrorOnPath(ctx, err)
}

func (ec *executionContext) unmarshalNSLSAPredicateSpec2ᚖgithubᚗcomᚋguacsecᚋguacᚋpkgᚋassemblerᚋgraphqlᚋmodelᚐSLSAPredicateSpec(ctx context.Context, v interface{}) (*model.SLSAPredicateSpec, error) {
	res, err := ec.unmarshalInputSLSAPredicateSpec(ctx, v)
	return &res, graphql.ErrorOnPath(ctx, err)
}

func (ec *executionContext) unmarshalOSLSAPredicateSpec2ᚕᚖgithubᚗcomᚋguacsecᚋguacᚋpkgᚋassemblerᚋgraphqlᚋmodelᚐSLSAPredicateSpecᚄ(ctx context.Context, v interface{}) ([]*model.SLSAPredicateSpec, error) {
	if v == nil {
		return nil, nil
	}
	var vSlice []interface{}
	if v != nil {
		vSlice = graphql.CoerceList(v)
	}
	var err error
	res := make([]*model.SLSAPredicateSpec, len(vSlice))
	for i := range vSlice {
		ctx := graphql.WithPathContext(ctx, graphql.NewPathWithIndex(i))
		res[i], err = ec.unmarshalNSLSAPredicateSpec2ᚖgithubᚗcomᚋguacsecᚋguacᚋpkgᚋassemblerᚋgraphqlᚋmodelᚐSLSAPredicateSpec(ctx, vSlice[i])
		if err != nil {
			return nil, err
		}
	}
	return res, nil
}

// endregion ***************************** type.gotpl *****************************
