package container

import (
	kbEnt "github.com/anondigriz/mogan-core/pkg/entities/containers/knowledgebase"
	"github.com/pelletier/go-toml/v2"
	"go.uber.org/zap"

	"github.com/anondigriz/mogan-mini/internal/storage/errors"
	errMsgs "github.com/anondigriz/mogan-mini/internal/storage/errors/messages"
)

func (c Container) WriteRules(ps map[string]kbEnt.Rule) error {
	for _, v := range ps {
		err := c.WriteRule(v)
		if err != nil {
			c.lg.Error(errMsgs.WriteFileFail, zap.Error(err))
			return err
		}
	}
	return nil
}

func (c Container) WriteRule(r kbEnt.Rule) error {
	data, err := toml.Marshal(r)
	if err != nil {
		c.lg.Error(errMsgs.TomlMarshalFail, zap.Error(err))
		return errors.NewTomlMarshalFailErr(err)
	}

	err = c.writeFile(data, r.UUID, RulesSubDir)
	if err != nil {
		c.lg.Error(errMsgs.WriteFileFail, zap.Error(err))
		return err
	}

	return nil
}

func (c Container) ReadRules() (map[string]kbEnt.Rule, error) {
	uuids := c.getFilesUUIDsInDir(RulesSubDir)
	result := make(map[string]kbEnt.Rule, len(uuids))
	for _, uuid := range uuids {
		r, err := c.ReadRule(uuid)
		if err != nil {
			c.lg.Error(errMsgs.ReadFileFail, zap.Error(err))
			return nil, err
		}
		result[uuid] = r
	}

	return result, nil
}

func (c Container) ReadRule(uuid string) (kbEnt.Rule, error) {
	data, err := c.readFile(uuid, RulesSubDir)
	if err != nil {
		c.lg.Error(errMsgs.ReadFileFail, zap.Error(err))
		return kbEnt.Rule{}, err
	}
	var r kbEnt.Rule
	err = toml.Unmarshal(data, &r)
	if err != nil {
		c.lg.Error(errMsgs.TomlUnmarshalFail, zap.Error(err))
		return kbEnt.Rule{}, errors.NewTomlUnmarshalFailErr(err)
	}

	return r, nil
}

func (c Container) RemoveRule(uuid string) error {
	return c.removeFile(uuid, RulesSubDir)
}
