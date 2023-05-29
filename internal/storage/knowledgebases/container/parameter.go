package container

import (
	kbEnt "github.com/anondigriz/mogan-core/pkg/entities/containers/knowledgebase"
	"github.com/pelletier/go-toml/v2"
	"go.uber.org/zap"

	"github.com/anondigriz/mogan-mini/internal/storage/errors"
	errMsgs "github.com/anondigriz/mogan-mini/internal/storage/errors/messages"
)

func (c Container) WriteParameters(ps map[string]kbEnt.Parameter) error {
	for _, v := range ps {
		err := c.WriteParameter(v)
		if err != nil {
			c.lg.Error(errMsgs.WriteFileFail, zap.Error(err))
			return err
		}
	}
	return nil
}

func (c Container) WriteParameter(p kbEnt.Parameter) error {
	data, err := toml.Marshal(p)
	if err != nil {
		c.lg.Error(errMsgs.TomlMarshalFail, zap.Error(err))
		return errors.NewTomlMarshalFailErr(err)
	}

	err = c.writeFile(data, p.UUID, ParametersSubDir)
	if err != nil {
		c.lg.Error(errMsgs.WriteFileFail, zap.Error(err))
		return err
	}

	return nil
}

func (c Container) ReadParameters() (map[string]kbEnt.Parameter, error) {
	uuids := c.getFilesUUIDsInDir(ParametersSubDir)
	result := make(map[string]kbEnt.Parameter, len(uuids))
	for _, uuid := range uuids {
		p, err := c.ReadParameter(uuid)
		if err != nil {
			c.lg.Error(errMsgs.ReadFileFail, zap.Error(err))
			return nil, err
		}
		result[uuid] = p
	}

	return result, nil
}

func (c Container) ReadParameter(uuid string) (kbEnt.Parameter, error) {
	data, err := c.readFile(uuid, ParametersSubDir)
	if err != nil {
		c.lg.Error(errMsgs.ReadFileFail, zap.Error(err))
		return kbEnt.Parameter{}, err
	}
	var p kbEnt.Parameter
	err = toml.Unmarshal(data, &p)
	if err != nil {
		c.lg.Error(errMsgs.TomlUnmarshalFail, zap.Error(err))
		return kbEnt.Parameter{}, errors.NewTomlUnmarshalFailErr(err)
	}

	return p, nil
}

func (c Container) RemoveParameter(uuid string) error {
	return c.removeFile(uuid, ParametersSubDir)
}
