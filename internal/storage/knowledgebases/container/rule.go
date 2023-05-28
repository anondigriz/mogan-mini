package container

import (
	"path"

	kbEnt "github.com/anondigriz/mogan-core/pkg/entities/containers/knowledgebase"
	"github.com/anondigriz/mogan-mini/internal/storage/errors"
	errMsgs "github.com/anondigriz/mogan-mini/internal/storage/errors/messages"
	"github.com/anondigriz/mogan-mini/internal/storage/knowledgebases/filesbroker"
	"github.com/pelletier/go-toml/v2"
	"go.uber.org/zap"
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
	fb := filesbroker.New(c.lg, path.Join(c.knowledgeBaseDir, RulesSubDir), fileExtension)
	paths := fb.GetAllFilesPaths()
	result := make(map[string]kbEnt.Rule, len(paths))
	for _, v := range paths {
		uuid := fb.GetFileUUID(v)
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
