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

func (c Container) WriteGroups(gs map[string]kbEnt.Group) error {
	for _, v := range gs {
		err := c.WriteGroup(v)
		if err != nil {
			c.lg.Error(errMsgs.WriteFileFail, zap.Error(err))
			return err
		}
	}
	return nil
}

func (c Container) WriteGroup(g kbEnt.Group) error {
	data, err := toml.Marshal(g)
	if err != nil {
		c.lg.Error(errMsgs.TomlMarshalFail, zap.Error(err))
		return errors.NewTomlMarshalFailErr(err)
	}

	err = c.writeFile(data, g.UUID, GroupsSubDir)
	if err != nil {
		c.lg.Error(errMsgs.WriteFileFail, zap.Error(err))
		return err
	}

	return nil
}

func (c Container) ReadGroups() (map[string]kbEnt.Group, error) {
	fb := filesbroker.New(c.lg, path.Join(c.knowledgeBaseDir, GroupsSubDir), fileExtension)
	paths := fb.GetAllFilesPaths()
	result := make(map[string]kbEnt.Group, len(paths))
	for _, v := range paths {
		uuid := fb.GetFileUUID(v)
		g, err := c.ReadGroup(uuid)
		if err != nil {
			c.lg.Error(errMsgs.ReadFileFail, zap.Error(err))
			return nil, err
		}
		result[uuid] = g
	}

	return result, nil
}

func (c Container) ReadGroup(uuid string) (kbEnt.Group, error) {
	data, err := c.readFile(uuid, GroupsSubDir)
	if err != nil {
		c.lg.Error(errMsgs.ReadFileFail, zap.Error(err))
		return kbEnt.Group{}, err
	}
	var g kbEnt.Group
	err = toml.Unmarshal(data, &g)
	if err != nil {
		c.lg.Error(errMsgs.TomlUnmarshalFail, zap.Error(err))
		return kbEnt.Group{}, errors.NewTomlUnmarshalFailErr(err)
	}

	return g, nil
}
