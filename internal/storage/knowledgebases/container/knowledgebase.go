package container

import (
	kbEnt "github.com/anondigriz/mogan-core/pkg/entities/containers/knowledgebase"
	"github.com/pelletier/go-toml/v2"
	"go.uber.org/zap"

	"github.com/anondigriz/mogan-mini/internal/storage/errors"
	errMsgs "github.com/anondigriz/mogan-mini/internal/storage/errors/messages"
)

func (c Container) WriteKnowledgeBase(kb kbEnt.KnowledgeBase) error {
	data, err := toml.Marshal(kb)
	if err != nil {
		c.lg.Error(errMsgs.TomlMarshalFail, zap.Error(err))
		return errors.NewTomlMarshalFailErr(err)
	}

	err = c.writeFile(data, kb.UUID, "")
	if err != nil {
		c.lg.Error(errMsgs.WriteFileFail, zap.Error(err))
		return err
	}

	return nil
}

func (c Container) ReadKnowledgeBase(uuid string) (kbEnt.KnowledgeBase, error) {
	data, err := c.readFile(uuid, "")
	if err != nil {
		c.lg.Error(errMsgs.ReadFileFail, zap.Error(err))
		return kbEnt.KnowledgeBase{}, err
	}
	var k kbEnt.KnowledgeBase
	err = toml.Unmarshal(data, &k)
	if err != nil {
		c.lg.Error(errMsgs.TomlUnmarshalFail, zap.Error(err))
		return kbEnt.KnowledgeBase{}, errors.NewTomlUnmarshalFailErr(err)
	}

	return k, nil
}
