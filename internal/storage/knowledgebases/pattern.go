package knowledgebases

import (
	kbEnt "github.com/anondigriz/mogan-core/pkg/entities/containers/knowledgebase"
	"go.uber.org/zap"

	errMsgs "github.com/anondigriz/mogan-mini/internal/storage/errors/messages"
	"github.com/anondigriz/mogan-mini/internal/storage/knowledgebases/container"
)

func (st Storage) CreatePattern(knowledgeBaseUUID string, pattern kbEnt.Pattern) error {
	cb := container.New(st.lg, st.KnowledgeBasesDir, knowledgeBaseUUID)
	err := cb.WritePattern(pattern)
	if err != nil {
		st.lg.Error(errMsgs.CreatePatternFail, zap.Error(err))
		return err
	}
	return nil
}

func (st Storage) GetPattern(knowledgeBaseUUID string, uuid string) (kbEnt.Pattern, error) {
	cb := container.New(st.lg, st.KnowledgeBasesDir, knowledgeBaseUUID)
	gr, err := cb.ReadPattern(uuid)
	if err != nil {
		st.lg.Error(errMsgs.GetPatternFail, zap.Error(err))
		return kbEnt.Pattern{}, err
	}
	return gr, nil
}

func (st Storage) UpdatePattern(ent kbEnt.Pattern) error {
	cb := container.New(st.lg, st.KnowledgeBasesDir, ent.UUID)
	err := cb.WritePattern(ent)
	if err != nil {
		st.lg.Error(errMsgs.UpdatePatternFail, zap.Error(err))
		return err
	}
	return nil
}

func (st Storage) RemovePattern(knowledgeBaseUUID string, uuid string) error {
	cb := container.New(st.lg, st.KnowledgeBasesDir, knowledgeBaseUUID)
	err := cb.RemovePattern(uuid)
	if err != nil {
		st.lg.Error(errMsgs.RemovePatternFail, zap.Error(err))
		return err
	}
	return nil
}

func (st Storage) GetAllPatterns(knowledgeBaseUUID string) (map[string]kbEnt.Pattern, error) {
	cb := container.New(st.lg, st.KnowledgeBasesDir, knowledgeBaseUUID)
	result, err := cb.ReadPatterns()
	if err != nil {
		st.lg.Error(errMsgs.GetPatternsFail, zap.Error(err))
		return nil, err
	}
	return result, nil
}
