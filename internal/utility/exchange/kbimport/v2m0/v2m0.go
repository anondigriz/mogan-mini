package v2m0

import (
	"encoding/xml"
	"fmt"
	"strings"
	"time"

	entV2M0 "github.com/anondigriz/mogan-core/pkg/knowledgebases/exchange/v2m0"
	kbEnt "github.com/anondigriz/mogan-mini/internal/entity/knowledgebase"
	"github.com/anondigriz/mogan-mini/internal/utility/errors"
	"github.com/google/uuid"
	"go.uber.org/zap"
)

type V2M0 struct {
	lg *zap.Logger
}

type ids struct {
	Parameters map[string]string
	Patterns   map[string]string
}

func New(lg *zap.Logger) *V2M0 {
	vm := &V2M0{
		lg: lg,
	}
	return vm
}

func (vm *V2M0) ParseXML(kbUUID string, content []byte) (kbEnt.Container, error) {
	var model entV2M0.Model
	err := xml.Unmarshal(content, &model)
	if err != nil {
		vm.lg.Error("fail to unmarshal the xml file", zap.Error(err))
		return kbEnt.Container{}, errors.NewXMLUnmarshalFailErr(err)
	}
	cont := &kbEnt.Container{
		Groups:     map[string]kbEnt.Group{},
		Parameters: map[string]kbEnt.Parameter{},
		Patterns:   map[string]kbEnt.Pattern{},
		Rules:      map[string]kbEnt.Rule{},
	}

	cont.KnowledgeBase = kbEnt.KnowledgeBase{
		BaseInfo: kbEnt.BaseInfo{
			UUID:        kbUUID,
			ID:          model.ID,
			ShortName:   model.ShortName,
			CreatedDate: time.Now(),
		},
		ExtraData: kbEnt.ExtraDataKnowledgeBase{
			Description: model.Description,
		},
	}
	cont.KnowledgeBase.ModifiedDate = cont.KnowledgeBase.CreatedDate

	mapIDs := &ids{
		Parameters: map[string]string{},
		Patterns:   map[string]string{},
	}

	for _, v := range model.Relations.Relations {
		err = parseRelation(v, cont, mapIDs)
		if err != nil {
			vm.lg.Error("parsing of the rule ended with an error", zap.Error(err))
			return kbEnt.Container{}, err
		}
	}

	gh, err := parseClass(model.Class, cont, mapIDs)
	if err != nil {
		vm.lg.Error("parsing of the main class ended with an error", zap.Error(err))
		return kbEnt.Container{}, err
	}
	cont.KnowledgeBase.ExtraData.Groups = gh

	return *cont, nil
}

func parseClass(class entV2M0.Class, cont *kbEnt.Container, ids *ids) (kbEnt.GroupHierarchy, error) {
	gr := kbEnt.Group{
		BaseInfo: kbEnt.BaseInfo{
			UUID:        uuid.New().String(),
			ID:          class.ID,
			ShortName:   class.ShortName,
			CreatedDate: time.Now(),
		},
		ExtraData: kbEnt.ExtraDataGroup{
			Description: class.Description,
		},
	}
	gr.ModifiedDate = gr.CreatedDate
	gh := kbEnt.GroupHierarchy{
		GroupUUID: gr.UUID,
		Contains:  []kbEnt.GroupHierarchy{},
	}
	cont.Groups[gr.UUID] = gr

	for _, v := range class.Parameters.Parameters {
		err := parseParameter(v, gr, cont, ids)
		if err != nil {
			return kbEnt.GroupHierarchy{}, err
		}
	}

	for _, v := range class.Classes.Classes {
		g, err := parseClass(v, cont, ids)
		if err != nil {
			return kbEnt.GroupHierarchy{}, err
		}
		gh.Contains = append(gh.Contains, g)
	}

	for _, v := range class.Rules.Rules {
		err := parseRule(v, cont, ids)
		if err != nil {
			return kbEnt.GroupHierarchy{}, err
		}
	}

	for _, v := range class.Constraints.Constraints {
		err := parseRule(v, cont, ids)
		if err != nil {
			return kbEnt.GroupHierarchy{}, err
		}
	}

	return gh, nil
}

func parseParameter(parameter entV2M0.Parameter, gr kbEnt.Group, cont *kbEnt.Container, ids *ids) error {
	pr := kbEnt.Parameter{
		BaseInfo: kbEnt.BaseInfo{
			UUID:        getOrCreateParameterUUID(parameter.ID, ids),
			ID:          parameter.ID,
			ShortName:   parameter.ShortName,
			CreatedDate: time.Now(),
		},
		ExtraData: kbEnt.ExtraDataParameter{
			Description:  parameter.Description,
			DefaultValue: parameter.DefaultValue,
		},
	}
	pr.ModifiedDate = pr.CreatedDate

	t, err := convertTypeParameter(parameter.Type)
	if err != nil {
		return err
	}
	pr.Type = t
	pr.GroupUUID = gr.UUID

	ids.Parameters[pr.ID] = pr.UUID
	cont.Parameters[pr.UUID] = pr
	return nil
}

func parseRelation(relation entV2M0.Relation, cont *kbEnt.Container, ids *ids) error {
	rn := kbEnt.Pattern{
		BaseInfo: kbEnt.BaseInfo{
			UUID:        uuid.New().String(),
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

	t, err := convertTypePattern(relation.RelationType)
	if err != nil {
		return err
	}
	rn.Type = t

	inObj, err := splitParameters(relation.InObjects)
	if err != nil {
		return err
	}

	for k, v := range inObj {
		paramType, err := convertTypeParameter(v)
		if err != nil {
			return err
		}

		rn.ExtraData.InputParameters = append(rn.ExtraData.InputParameters, kbEnt.ParameterPattern{
			ShortName: k,
			Type:      paramType,
		})
	}

	outObj, err := splitParameters(relation.OutObjects)
	if err != nil {
		return err
	}

	for k, v := range outObj {
		paramType, err := convertTypeParameter(v)
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

func parseRule(rule entV2M0.Rule, cont *kbEnt.Container, ids *ids) error {
	re := kbEnt.Rule{
		BaseInfo: kbEnt.BaseInfo{
			UUID:        uuid.New().String(),
			ID:          rule.ID,
			ShortName:   rule.ShortName,
			CreatedDate: time.Now(),
		},
		ExtraData: kbEnt.ExtraDataRule{
			Description:      rule.Description,
			InputParameters:  []kbEnt.ParameterRule{},
			OutputParameters: []kbEnt.ParameterRule{},
		},
	}
	re.ModifiedDate = re.CreatedDate

	initIDs, err := splitParameters(rule.InitIDs)
	if err != nil {
		return err
	}

	for k, v := range initIDs {
		re.ExtraData.InputParameters = append(re.ExtraData.InputParameters, kbEnt.ParameterRule{
			ShortName:     k,
			ParameterUUID: getOrCreateParameterUUID(v, ids),
		})
	}

	resultIDs, err := splitParameters(rule.ResultIDs)
	if err != nil {
		return err
	}

	for k, v := range resultIDs {
		re.ExtraData.OutputParameters = append(re.ExtraData.OutputParameters, kbEnt.ParameterRule{
			ShortName:     k,
			ParameterUUID: getOrCreateParameterUUID(v, ids),
		})
	}

	pnUUID, ok := ids.Patterns[rule.RelationID]
	if !ok {
		return errors.NewParsingXMLFailErr(
			fmt.Sprintf("no relation with id '%s' was found for the rule with id '%s'", rule.RelationID, rule.ID),
			nil)
	}
	re.PatternUUID = pnUUID
	cont.Rules[re.UUID] = re

	return nil

}

func getOrCreateParameterUUID(id string, ids *ids) string {
	prUUID, ok := ids.Parameters[id]
	if !ok {
		prUUID = uuid.New().String()
	}
	return prUUID
}

func splitParameters(base string) (map[string]string, error) {
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

func convertTypeParameter(base string) (kbEnt.TypeParameter, error) {
	switch base {
	case "double":
		return kbEnt.Double, nil
	case "string":
		return kbEnt.String, nil
	default:
		return kbEnt.String, errors.NewParsingXMLFailErr(
			fmt.Sprintf("unknown parameter type from the XML file %s", base),
			nil)
	}
}

func convertTypePattern(base string) (kbEnt.TypePattern, error) {
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
