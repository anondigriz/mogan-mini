package knowledgebases

import (
	kbEnt "github.com/anondigriz/mogan-core/pkg/entities/containers/knowledgebase"
	"go.uber.org/zap"

	errMsgs "github.com/anondigriz/mogan-mini/internal/storage/errors/messages"
	"github.com/anondigriz/mogan-mini/internal/storage/knowledgebases/container"
)

func (st Storage) GetContainerByUUID(uuid string) (*kbEnt.Container, error) {
	cb := container.New(st.lg, st.KnowledgeBasesSubDir, uuid)
	cont, err := cb.ReadContainer()

	if err != nil {
		st.lg.Error(errMsgs.GetKnowledgeFail, zap.Error(err))
		return nil, err
	}
	return cont, nil
}

func (st Storage) RemoveContainerByUUID(uuid string) error {
	return st.fb.RemoveFileByUUID(uuid)
}
