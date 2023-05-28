package knowledgebase

import (
	"strconv"

	kbEnt "github.com/anondigriz/mogan-core/pkg/entities/containers/knowledgebase"
	"github.com/anondigriz/mogan-core/pkg/entities/types"

	"github.com/anondigriz/mogan-mini/internal/storage/errors"
)

const (
	Double  = "Double"
	String  = "String"
	Boolean = "Boolean"
)

type Parameter struct {
	BaseInfo
	Type         string
	DefaultValue string
}

func mapFromParameterType(base types.ParameterType) (string, error) {
	switch base {
	case types.Double:
		return Double, nil
	case types.String:
		return String, nil
	case types.Boolean:
		return Boolean, nil
	default:
		return "", errors.NewTypeIsNotSupportedByStorageErr(strconv.Itoa(int(base)))
	}
}

func mapToParameterType(base string) (types.ParameterType, error) {
	switch base {
	case String:
		return types.String, nil
	case Double:
		return types.Double, nil
	case Boolean:
		return types.Boolean, nil
	default:
		return 0, errors.NewTypeIsNotSupportedByStorageErr(base)
	}
}

func (p *Parameter) Fill(base kbEnt.Parameter) error {
	p.BaseInfo.Fill(base.BaseInfo)
	p.DefaultValue = base.DefaultValue
	t, err := mapFromParameterType(base.Type)
	if err != nil {
		return err
	}
	p.Type = t
	return nil
}

func (p Parameter) Extract() (kbEnt.Parameter, error) {
	result := kbEnt.Parameter{
		BaseInfo: p.BaseInfo.Extract(),
	}
	result.DefaultValue = p.DefaultValue
	t, err := mapToParameterType(p.Type)
	if err != nil {
		return kbEnt.Parameter{}, err
	}
	result.Type = t
	return result, nil
}
