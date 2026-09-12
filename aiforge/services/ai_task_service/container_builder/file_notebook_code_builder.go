package container_builder

import (
	"code.gitea.io/gitea/entity"
	"code.gitea.io/gitea/models"
	"code.gitea.io/gitea/modules/log"
	"code.gitea.io/gitea/modules/setting"
	"code.gitea.io/gitea/routers/response"
	"code.gitea.io/gitea/services/ai_task_service/context"
)

type FileNoteBookCodeBuilder struct {
	Opts *entity.ContainerBuildOpts
}

func init() {
	o := &FileNoteBookCodeBuilder{}
	RegisterContainerBuilder(o)
}

func (b *FileNoteBookCodeBuilder) SetOpts(opts *entity.ContainerBuildOpts) {
	b.Opts = opts
}

func (b *FileNoteBookCodeBuilder) GetContainerType() entity.ContainerDataType {
	return entity.ContainerFileNoteBookCode
}

func (b *FileNoteBookCodeBuilder) Build(ctx *context.CreationContext) ([]entity.ContainerData, *response.BizError) {
	if b.Opts.Disable {
		return nil, nil
	}
	repo := ctx.Request.FileRepository
	if repo == nil {
		return nil, nil
	}
	//在线运行notebook不需要代码挂载或者调度，只需要把对对应分支的代码仓下载到指定目录。上传目标分支的逻辑在其他地方（继承原有逻辑）
	err := DownloadBranch(repo, getCodePath(ctx.Request.JobName, repo, ctx.Request.FileBranchName), ctx.Request.FileBranchName)
	if err != nil {
		log.Error("download code failed", err)
		return nil, response.LOAD_CODE_FAILED
	}
	return nil, nil
}

func getCodePath(jobName string, repo *models.Repository, branchName string) string {
	return setting.JobPath + jobName + "/code" + "/" + repo.OwnerName + "/" + repo.Name + "/" + branchName
}
