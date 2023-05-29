package knowledgebases

import (
	kbEnt "github.com/anondigriz/mogan-core/pkg/entities/containers/knowledgebase"
	"go.uber.org/zap"

	errMsgs "github.com/anondigriz/mogan-mini/internal/storage/errors/messages"
	"github.com/anondigriz/mogan-mini/internal/storage/knowledgebases/container"
)

func (st Storage) CreateRule(knowledgeBaseUUID string, rule kbEnt.Rule) error {
	cb := container.New(st.lg, st.KnowledgeBasesDir, knowledgeBaseUUID)
	err := cb.WriteRule(rule)
	if err != nil {
		st.lg.Error(errMsgs.CreateRuleFail, zap.Error(err))
		return err
	}
	return nil
}

func (st Storage) GetRule(knowledgeBaseUUID string, uuid string) (kbEnt.Rule, error) {
	cb := container.New(st.lg, st.KnowledgeBasesDir, knowledgeBaseUUID)
	gr, err := cb.ReadRule(uuid)
	if err != nil {
		st.lg.Error(errMsgs.GetRuleFail, zap.Error(err))
		return kbEnt.Rule{}, err
	}
	return gr, nil
}

func (st Storage) UpdateRule(ent kbEnt.Rule) error {
	cb := container.New(st.lg, st.KnowledgeBasesDir, ent.UUID)
	err := cb.WriteRule(ent)
	if err != nil {
		st.lg.Error(errMsgs.UpdateRuleFail, zap.Error(err))
		return err
	}
	return nil
}

func (st Storage) RemoveRule(knowledgeBaseUUID string, uuid string) error {
	cb := container.New(st.lg, st.KnowledgeBasesDir, knowledgeBaseUUID)
	err := cb.RemoveRule(uuid)
	if err != nil {
		st.lg.Error(errMsgs.RemoveRuleFail, zap.Error(err))
		return err
	}
	return nil
}

func (st Storage) GetAllRules(knowledgeBaseUUID string) (map[string]kbEnt.Rule, error) {
	cb := container.New(st.lg, st.KnowledgeBasesDir, knowledgeBaseUUID)
	result, err := cb.ReadRules()
	if err != nil {
		st.lg.Error(errMsgs.GetRulesFail, zap.Error(err))
		return nil, err
	}
	return result, nil
}
