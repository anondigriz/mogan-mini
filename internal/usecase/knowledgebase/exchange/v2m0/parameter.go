package v2m0

import (
	"fmt"
	"time"

	uuidGen "github.com/google/uuid"

	entV2M0 "github.com/anondigriz/mogan-core/pkg/knowledgebases/exchange/v2m0"
	kbEnt "github.com/anondigriz/mogan-mini/internal/entity/knowledgebase"
	"github.com/anondigriz/mogan-mini/internal/usecase/errors"
)

func (vm V2M0) parseParameter(parameter entV2M0.Parameter, gr kbEnt.Group, cont *kbEnt.Container, ids *ids) error {
	pr := kbEnt.Parameter{
		BaseInfo: kbEnt.BaseInfo{
			UUID:        vm.getOrCreateParameterUUID(parameter.ID, ids),
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

	t, err := vm.convertTypeParameter(parameter.Type)
	if err != nil {
		return err
	}
	pr.Type = t
	pr.GroupUUID = gr.UUID

	ids.Parameters[pr.ID] = pr.UUID
	cont.Parameters[pr.UUID] = pr
	return nil
}

func (vm V2M0) convertTypeParameter(base string) (kbEnt.TypeParameter, error) {
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

func (vm V2M0) getOrCreateParameterUUID(id string, ids *ids) string {
	prUUID, ok := ids.Parameters[id]
	if !ok {
		prUUID = uuidGen.New().String()
	}
	return prUUID
}
