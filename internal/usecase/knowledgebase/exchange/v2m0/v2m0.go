package v2m0

import (
	"encoding/xml"
	"time"

	"go.uber.org/zap"

	entV2M0 "github.com/anondigriz/mogan-core/pkg/knowledgebases/exchange/v2m0"
	kbEnt "github.com/anondigriz/mogan-mini/internal/entity/knowledgebase"
	"github.com/anondigriz/mogan-mini/internal/usecase/errors"
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

func (vm V2M0) ParseXML(kbUUID string, content []byte) (kbEnt.Container, error) {
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
		err = vm.parseRelation(v, cont, mapIDs)
		if err != nil {
			vm.lg.Error("parsing of the rule ended with an error", zap.Error(err))
			return kbEnt.Container{}, err
		}
	}

	gh, err := vm.parseClass(model.Class, cont, mapIDs)
	if err != nil {
		vm.lg.Error("parsing of the main class ended with an error", zap.Error(err))
		return kbEnt.Container{}, err
	}
	cont.KnowledgeBase.ExtraData.Groups = gh

	return *cont, nil
}
