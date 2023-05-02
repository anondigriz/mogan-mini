package v2m0

import (
	"fmt"
	"time"

	uuidGen "github.com/google/uuid"

	entV2M0 "github.com/anondigriz/mogan-core/pkg/knowledgebases/exchange/v2m0"
	kbEnt "github.com/anondigriz/mogan-mini/internal/entity/knowledgebase"
	"github.com/anondigriz/mogan-mini/internal/usecase/errors"
)

func (vm V2M0) parseRule(rule entV2M0.Rule, cont *kbEnt.Container, ids *ids) error {
	re := kbEnt.Rule{
		BaseInfo: kbEnt.BaseInfo{
			UUID:        uuidGen.New().String(),
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

	initIDs, err := vm.splitParameters(rule.InitIDs)
	if err != nil {
		return err
	}

	for k, v := range initIDs {
		re.ExtraData.InputParameters = append(re.ExtraData.InputParameters, kbEnt.ParameterRule{
			ShortName:     k,
			ParameterUUID: vm.getOrCreateParameterUUID(v, ids),
		})
	}

	resultIDs, err := vm.splitParameters(rule.ResultIDs)
	if err != nil {
		return err
	}

	for k, v := range resultIDs {
		re.ExtraData.OutputParameters = append(re.ExtraData.OutputParameters, kbEnt.ParameterRule{
			ShortName:     k,
			ParameterUUID: vm.getOrCreateParameterUUID(v, ids),
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
