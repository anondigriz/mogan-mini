package v2m0

import (
	"fmt"
	"strings"
	"time"

	uuidGen "github.com/google/uuid"

	entV2M0 "github.com/anondigriz/mogan-core/pkg/knowledgebases/exchange/v2m0"
	kbEnt "github.com/anondigriz/mogan-mini/internal/entity/knowledgebase"
	"github.com/anondigriz/mogan-mini/internal/usecase/errors"
)

func (vm V2M0) parseRelation(relation entV2M0.Relation, cont *kbEnt.Container, ids *ids) error {
	rn := kbEnt.Pattern{
		BaseInfo: kbEnt.BaseInfo{
			UUID:        uuidGen.New().String(),
			ID:          relation.ID,
			ShortName:   relation.ShortName,
			CreatedDate: time.Now(),
		},
		ExtraData: kbEnt.ExtraDataPattern{
			Description:      relation.Description,
			Language:         kbEnt.JS,
			Script:           relation.Script,
			InputParameters:  []kbEnt.ParameterPattern{},
			OutputParameters: []kbEnt.ParameterPattern{},
		},
	}
	rn.ModifiedDate = rn.CreatedDate

	t, err := vm.convertTypePattern(relation.RelationType)
	if err != nil {
		return err
	}
	rn.Type = t

	inObj, err := vm.splitParameters(relation.InObjects)
	if err != nil {
		return err
	}

	for k, v := range inObj {
		paramType, err := vm.convertTypeParameter(v)
		if err != nil {
			return err
		}

		rn.ExtraData.InputParameters = append(rn.ExtraData.InputParameters, kbEnt.ParameterPattern{
			ShortName: k,
			Type:      paramType,
		})
	}

	outObj, err := vm.splitParameters(relation.OutObjects)
	if err != nil {
		return err
	}

	for k, v := range outObj {
		paramType, err := vm.convertTypeParameter(v)
		if err != nil {
			return err
		}

		rn.ExtraData.OutputParameters = append(rn.ExtraData.OutputParameters, kbEnt.ParameterPattern{
			ShortName: k,
			Type:      paramType,
		})
	}

	ids.Patterns[rn.ID] = rn.UUID
	cont.Patterns[rn.UUID] = rn
	return nil

}

func (vm V2M0) convertTypePattern(base string) (kbEnt.TypePattern, error) {
	switch base {
	case "constr":
		return kbEnt.Constraint, nil
	case "ifclause":
		return kbEnt.IfThenElse, nil
	case "prog":
		return kbEnt.Program, nil
	case "simple":
		return kbEnt.Formula, nil
	default:
		return kbEnt.Program, errors.NewParsingXMLFailErr(
			fmt.Sprintf("unknown pattern type from the XML file %s", base),
			nil)
	}
}

func (vm V2M0) splitParameters(base string) (map[string]string, error) {
	params := map[string]string{}
	if base == "" {
		return params, nil
	}

	pair := strings.Split(base, ";")
	for _, v := range pair {
		keyValue := strings.Split(v, ":")
		if len(keyValue) != 2 {
			return nil, errors.NewParsingXMLFailErr(
				fmt.Sprintf("'%s' is not a key-value pair", v),
				nil)
		}
		params[keyValue[0]] = keyValue[1]
	}
	return params, nil
}
