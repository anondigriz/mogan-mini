package container

import (
	kbEnt "github.com/anondigriz/mogan-core/pkg/entities/containers/knowledgebase"
	"github.com/pelletier/go-toml/v2"
	"go.uber.org/zap"

	"github.com/anondigriz/mogan-mini/internal/storage/errors"
	errMsgs "github.com/anondigriz/mogan-mini/internal/storage/errors/messages"
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
	uuids := c.getFilesUUIDsInDir(GroupsSubDir)
	result := make(map[string]kbEnt.Group, len(uuids))
	for _, uuid := range uuids {
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

func (c Container) RemoveGroup(uuid string) error {
	return c.removeFile(uuid, GroupsSubDir)
}
