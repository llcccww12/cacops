package container_builder

import (
	"code.gitea.io/gitea/routers/response"
	"code.gitea.io/gitea/services/ai_task_service/context"
)

type BuilderChain struct {
	builderList []ContainerBuilder
}

func NewBuilderChain() *BuilderChain {
	return &BuilderChain{builderList: make([]ContainerBuilder, 0)}
}

func (c *BuilderChain) Next(b ContainerBuilder) *BuilderChain {
	c.builderList = append(c.builderList, b)
	return c
}

func (c *BuilderChain) Run(ctx *context.CreationContext) *response.BizError {
	for _, builder := range c.builderList {
		current := ctx.GetContainerDataArray(builder.GetContainerType())
		//如果已经存在则不需要再构建
		if current != nil && len(current) > 0 {
			continue
		}
		d, err := builder.Build(ctx)
		if err != nil {
			return err
		}
		ctx.AddContainerData(builder.GetContainerType(), d)
	}
	return nil
}
