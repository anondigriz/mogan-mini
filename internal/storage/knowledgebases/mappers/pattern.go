package knowledgebase

import (
	"strconv"

	kbEnt "github.com/anondigriz/mogan-core/pkg/entities/containers/knowledgebase"
	"github.com/anondigriz/mogan-core/pkg/entities/types"
	"github.com/anondigriz/mogan-mini/internal/storage/errors"
)

const (
	Program    = "Program"
	Constraint = "Constraint"

	JS  = "JS"
	Lua = "Lua"
)

type Pattern struct {
	BaseInfo
	Type             string
	ScriptLanguage   string
	Script           string
	InputParameters  []ParameterPattern
	OutputParameters []ParameterPattern
}

type ParameterPattern struct {
	ShortName string
	Type      string
}

func mapFromPatternType(base types.PatternType) (string, error) {
	switch base {
	case types.Program:
		return Program, nil
	case types.Constraint:
		return Constraint, nil
	default:
		return "", errors.NewTypeIsNotSupportedByStorageErr(strconv.Itoa(int(base)))
	}
}

func mapToPatternType(base string) (types.PatternType, error) {
	switch base {
	case Program:
		return types.Program, nil
	case Constraint:
		return types.Constraint, nil
	default:
		return 0, errors.NewTypeIsNotSupportedByStorageErr(base)
	}
}

func mapFromScriptLanguageType(base types.ScriptLanguageType) (string, error) {
	switch base {
	case types.JS:
		return JS, nil
	case types.Lua:
		return Lua, nil
	default:
		return "", errors.NewTypeIsNotSupportedByStorageErr(strconv.Itoa(int(base)))
	}
}

func mapToScriptLanguageType(base string) (types.ScriptLanguageType, error) {
	switch base {
	case JS:
		return types.JS, nil
	case Lua:
		return types.Lua, nil
	default:
		return 0, errors.NewTypeIsNotSupportedByStorageErr(base)
	}
}

func (p *Pattern) Fill(base kbEnt.Pattern) error {
	p.BaseInfo.Fill(base.BaseInfo)

	t, err := mapFromPatternType(base.Type)
	if err != nil {
		return err
	}
	p.Type = t

	s, err := mapFromScriptLanguageType(base.ScriptLanguage)
	if err != nil {
		return err
	}
	p.ScriptLanguage = s

	p.Script = base.Script

	p.InputParameters = make([]ParameterPattern, len(base.InputParameters))
	for _, v := range base.InputParameters {
		var pp ParameterPattern
		err = pp.Fill(v)
		if err != nil {
			return err
		}
		p.InputParameters = append(p.InputParameters, pp)
	}

	p.OutputParameters = make([]ParameterPattern, len(base.OutputParameters))
	for _, v := range base.OutputParameters {
		var pp ParameterPattern
		err = pp.Fill(v)
		if err != nil {
			return err
		}
		p.OutputParameters = append(p.OutputParameters, pp)
	}

	return nil
}

func (p Pattern) Extract() (kbEnt.Pattern, error) {
	result := kbEnt.Pattern{
		BaseInfo: p.BaseInfo.Extract(),
	}

	t, err := mapToPatternType(p.Type)
	if err != nil {
		return kbEnt.Pattern{}, err
	}
	result.Type = t

	s, err := mapToScriptLanguageType(p.ScriptLanguage)
	if err != nil {
		return kbEnt.Pattern{}, err
	}
	result.ScriptLanguage = s

	result.Script = p.Script

	result.InputParameters = make([]kbEnt.ParameterPattern, len(p.InputParameters))
	for _, v := range p.InputParameters {
		pp, err := v.Extract()
		if err != nil {
			return kbEnt.Pattern{}, err
		}
		result.InputParameters = append(result.InputParameters, pp)
	}

	result.OutputParameters = make([]kbEnt.ParameterPattern, len(p.OutputParameters))
	for _, v := range p.OutputParameters {
		pp, err := v.Extract()
		if err != nil {
			return kbEnt.Pattern{}, err
		}
		result.OutputParameters = append(result.OutputParameters, pp)
	}

	return result, nil
}

func (p *ParameterPattern) Fill(base kbEnt.ParameterPattern) error {
	p.ShortName = base.ShortName
	t, err := mapFromParameterType(base.Type)
	if err != nil {
		return err
	}
	p.Type = t
	return nil
}

func (p ParameterPattern) Extract() (kbEnt.ParameterPattern, error) {
	result := kbEnt.ParameterPattern{
		ShortName: p.ShortName,
	}
	t, err := mapToParameterType(p.Type)
	if err != nil {
		return kbEnt.ParameterPattern{}, err
	}
	result.Type = t
	return result, nil
}
