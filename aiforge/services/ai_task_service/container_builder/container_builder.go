package container_builder

import (
	"code.gitea.io/gitea/entity"
	"code.gitea.io/gitea/modules/log"
	"code.gitea.io/gitea/routers/response"
	"code.gitea.io/gitea/services/ai_task_service/context"
	"fmt"
	"reflect"
)

type ContainerBuilder interface {
	Build(ctx *context.CreationContext) ([]entity.ContainerData, *response.BizError)
	GetContainerType() entity.ContainerDataType
	SetOpts(opts *entity.ContainerBuildOpts)
}

var containerBuilderMap = map[entity.ContainerDataType]reflect.Type{}

func RegisterContainerBuilder(builder ContainerBuilder) {
	containerBuilderMap[builder.GetContainerType()] = reflect.TypeOf(builder)
}

func CreateContainerBuilder(containerType entity.ContainerDataType, opts *entity.ContainerBuildOpts) ContainerBuilder {
	defer func() {
		if err := recover(); err != nil {
			combinedErr := fmt.Errorf("%s\n%s", err, log.Stack(2))
			log.Error("PANIC:%v", combinedErr)
		}
	}()
	t := containerBuilderMap[containerType]
	if t == nil {
		return nil
	}
	b := reflect.New(t.Elem()).Interface().(ContainerBuilder)
	b.SetOpts(opts)
	return b
}

func BuildContainerDataChain(configMap map[entity.ContainerDataType]*entity.ContainerBuildOpts) *BuilderChain {
	c := NewBuilderChain()
	for k, v := range configMap {
		b := CreateContainerBuilder(k, v)
		if b == nil {
			continue
		}
		c.Next(b)
	}
	return c
}
