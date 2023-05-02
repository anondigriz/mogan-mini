package v2m0

import (
	"time"

	uuidGen "github.com/google/uuid"

	entV2M0 "github.com/anondigriz/mogan-core/pkg/knowledgebases/exchange/v2m0"
	kbEnt "github.com/anondigriz/mogan-mini/internal/entity/knowledgebase"
)

func (vm V2M0) parseClass(class entV2M0.Class, cont *kbEnt.Container, ids *ids) (kbEnt.GroupHierarchy, error) {
	gr := kbEnt.Group{
		BaseInfo: kbEnt.BaseInfo{
			UUID:        uuidGen.New().String(),
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
		err := vm.parseParameter(v, gr, cont, ids)
		if err != nil {
			return kbEnt.GroupHierarchy{}, err
		}
	}

	for _, v := range class.Classes.Classes {
		g, err := vm.parseClass(v, cont, ids)
		if err != nil {
			return kbEnt.GroupHierarchy{}, err
		}
		gh.Contains = append(gh.Contains, g)
	}

	for _, v := range class.Rules.Rules {
		err := vm.parseRule(v, cont, ids)
		if err != nil {
			return kbEnt.GroupHierarchy{}, err
		}
	}

	for _, v := range class.Constraints.Constraints {
		err := vm.parseRule(v, cont, ids)
		if err != nil {
			return kbEnt.GroupHierarchy{}, err
		}
	}

	return gh, nil
}
