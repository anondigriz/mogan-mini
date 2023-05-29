package container

import (
	kbEnt "github.com/anondigriz/mogan-core/pkg/entities/containers/knowledgebase"
	"github.com/pelletier/go-toml/v2"
	"go.uber.org/zap"

	"github.com/anondigriz/mogan-mini/internal/storage/errors"
	errMsgs "github.com/anondigriz/mogan-mini/internal/storage/errors/messages"
)

func (c Container) WritePatterns(ps map[string]kbEnt.Pattern) error {
	for _, v := range ps {
		err := c.WritePattern(v)
		if err != nil {
			c.lg.Error(errMsgs.WriteFileFail, zap.Error(err))
			return err
		}
	}
	return nil
}

func (c Container) WritePattern(p kbEnt.Pattern) error {
	data, err := toml.Marshal(p)
	if err != nil {
		c.lg.Error(errMsgs.TomlMarshalFail, zap.Error(err))
		return errors.NewTomlMarshalFailErr(err)
	}

	err = c.writeFile(data, p.UUID, PatternsSubDir)
	if err != nil {
		c.lg.Error(errMsgs.WriteFileFail, zap.Error(err))
		return err
	}

	return nil
}

func (c Container) ReadPatterns() (map[string]kbEnt.Pattern, error) {
	uuids := c.getFilesUUIDsInDir(PatternsSubDir)
	result := make(map[string]kbEnt.Pattern, len(uuids))
	for _, uuid := range uuids {
		p, err := c.ReadPattern(uuid)
		if err != nil {
			c.lg.Error(errMsgs.ReadFileFail, zap.Error(err))
			return nil, err
		}
		result[uuid] = p
	}

	return result, nil
}

func (c Container) ReadPattern(uuid string) (kbEnt.Pattern, error) {
	data, err := c.readFile(uuid, PatternsSubDir)
	if err != nil {
		c.lg.Error(errMsgs.ReadFileFail, zap.Error(err))
		return kbEnt.Pattern{}, err
	}
	var p kbEnt.Pattern
	err = toml.Unmarshal(data, &p)
	if err != nil {
		c.lg.Error(errMsgs.TomlUnmarshalFail, zap.Error(err))
		return kbEnt.Pattern{}, errors.NewTomlUnmarshalFailErr(err)
	}

	return p, nil
}

func (c Container) RemovePattern(uuid string) error {
	return c.removeFile(uuid, PatternsSubDir)
}
