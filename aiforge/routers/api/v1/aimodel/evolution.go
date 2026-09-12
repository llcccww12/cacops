package aimodel

import (
	"encoding/json"
	"fmt"
	"regexp"
	"strings"

	"code.gitea.io/gitea/entity"
	"code.gitea.io/gitea/models"
	"code.gitea.io/gitea/modules/log"
	"code.gitea.io/gitea/routers/response"
	"code.gitea.io/gitea/services/ai_task_service/task"
	"code.gitea.io/gitea/services/subject_service"
)

const (
	NODE_TYPE_TASK  = 0
	NODE_TYPE_MODEL = 1
)

func GetCloudbrain(aimodel *models.AiModelManage) (*models.Cloudbrain, error) {
	if aimodel.TrainTaskInfo == "" {
		log.Info("AimodelEvolution: no TrainTaskInfo; aimodelID=" + aimodel.ID)
		return nil, nil
	}

	if aimodel.ModelType != models.MODEL_ONLINE_TYPE {
		log.Info("AimodelEvolution: not online model; aimodelID=" + aimodel.ID)
		return nil, nil
	}

	removeIpInfo(aimodel)

	var job models.Cloudbrain
	if err := json.Unmarshal([]byte(aimodel.TrainTaskInfo), &job); err != nil {
		log.Error(fmt.Sprintf("AimodelEvolution: unmarshal TrainTaskInfo failed; aimodelID=%v; err=%v", aimodel.ID, err))
		return nil, err
	}

	return &job, nil
}

func GetTaskInfo(aimodel *models.AiModelManage) (*entity.AITaskDetailInfo, error) {
	job, err := GetCloudbrain(aimodel)
	if err != nil {
		log.Error(fmt.Sprintf("AimodelEvolution failed GetJobFromAimodel; aimodelID=%v; err=%v;", aimodel.ID, err.Error()))
		return nil, err
	}
	if job == nil {
		log.Info("AimodelEvolution: no Cloudbrain job for aimodelID=" + aimodel.ID)
		return nil, nil
	}

	taskDetailedInfo, err := task.BuildAITaskByCloudbrain(job)
	if err != nil {
		log.Error(fmt.Sprintf("AimodelEvolution failed query taskDetailedInfo; aimodelID=%v; err=%v;", aimodel.ID, err.Error()))
		return nil, err
	}

	job.GetRepository()
	taskDetailedInfo.RepoID = job.RepoID
	if job.Repo == nil {
		taskDetailedInfo.RepoName = ""
		taskDetailedInfo.RepoOwnerName = ""
	} else {
		taskDetailedInfo.RepoName = job.Repo.Name
		taskDetailedInfo.RepoOwnerName = job.Repo.OwnerName
	}

	log.Info(fmt.Sprintf("AimodelEvolution Cloudbrain GetRepository success; aimodelID=%v; jobID=%v; RepoID=%v", aimodel.ID, job.ID, job.RepoID))

	return taskDetailedInfo, nil
}

func buildTaskInfoNode(taskInfo *entity.AITaskDetailInfo) *entity.TaskInfo4AimodelGraph {
	return entity.BuildTaskInfo4AimodelGraph(taskInfo)
}

func buildAimodelInfoNode(aimodel *models.AiModelManage) *entity.AimodelInfo {
	if err := aimodel.GetOwner(); err != nil {
		log.Error(fmt.Sprintf("AimodelEvolution buildAimodelInfo GetOwner err; aimodelID=%v; err=%v", aimodel.ID, err.Error()))
	}
	return entity.BuildAimodelInfo(aimodel)
}

func BuildAimodelEvolutionGraph(doer *models.User, aimodel *models.AiModelManage) (*entity.AimodelGraphNode, *response.BizError) {
	log.Info("AimodelEvolution START building graph; aimodelID=" + aimodel.ID)
	currentNode := &entity.AimodelGraphNode{
		Type:        NODE_TYPE_MODEL,
		IsParent:    false,
		IsCurrent:   true,
		IsPrivate:   aimodel.IsPrivate,
		Visible:     subject_service.QueryAimodelVisibility(doer, aimodel),
		AimodelInfo: buildAimodelInfoNode(aimodel),
	}

	rootNode, errAn := BuildAntecedents(currentNode, aimodel, doer)
	if errAn != nil {
		log.Error("AimodelEvolution BuildAntecedents err; err=" + errAn.Error())
		return nil, response.NewBizError(errAn)
	}

	if errDe := BuildDescendent(currentNode, aimodel, doer); errDe != nil {
		return nil, response.NewBizError(fmt.Errorf("BuildDescendent failed: %w", errDe))
	}

	return rootNode, nil
}

func BuildAntecedents(curNode *entity.AimodelGraphNode, aimodel *models.AiModelManage, doer *models.User) (*entity.AimodelGraphNode, error) {
	log.Info("AimodelEvolution BuildAntecedents starts; aimodelID=" + aimodel.ID)
	taskInfo, err := GetTaskInfo(aimodel)
	if err != nil {
		log.Error("AimodelEvolution Antecedents failed getting taskInfo for curNode; err=" + err.Error())
		return curNode, nil
	}
	if taskInfo == nil {
		log.Info("AimodelEvolution BuildAntecedents: no taskInfo; aimodelID=" + aimodel.ID)
		return curNode, nil
	}

	if len(taskInfo.PretrainModelList) == 0 {
		log.Info("AimodelEvolution BuildAntecedents: empty PretrainModelList; aimodelID=" + aimodel.ID)
		return curNode, nil
	}

	parentsAimodelNodes := make([]*entity.AimodelGraphNode, 0)
	for _, pretrainModel := range taskInfo.PretrainModelList {
		modelId := pretrainModel.ID
		parentModel, err := models.GetAimodelByID(modelId)
		if err != nil {
			log.Error(fmt.Sprintf("AimodelEvolution Antecedents failed get aimodel, skipped; aimodelID=%v; err=%v", aimodel.ID, err.Error()))
			continue
		}
		parentModelNode := &entity.AimodelGraphNode{
			Type:        NODE_TYPE_MODEL,
			IsParent:    true,
			IsPrivate:   parentModel.IsPrivate,
			Visible:     subject_service.QueryAimodelVisibility(doer, parentModel),
			AimodelInfo: buildAimodelInfoNode(parentModel),
		}
		parentsAimodelNodes = append(parentsAimodelNodes, parentModelNode)
		log.Info(fmt.Sprintf("AimodelEvolution Antecedents success added parent aimodel; aimodelID=%v; parentModelID=%v", aimodel.ID, parentModel.ID))
		continue
	}

	if len(parentsAimodelNodes) == 0 {
		return curNode, nil
	}

	nexts := []*entity.AimodelGraphNode{curNode}
	rootNode := &entity.AimodelGraphNode{
		Type:      NODE_TYPE_MODEL,
		IsParent:  true,
		IsPrivate: false,
		Visible:   true,
		Parents:   parentsAimodelNodes,
		Next:      nexts,
	}
	log.Info(fmt.Sprintf("AimodelEvolution BuildAntecedents finish; aimodelID=%v; numParents=%v", aimodel.ID, len(parentsAimodelNodes)))

	return rootNode, nil
}

func isRealDescendent(taskInfo *entity.AITaskDetailInfo, aimodelID string) bool {
	splitModelIds := strings.Split(taskInfo.PretrainModelId, ";")
	for _, id := range splitModelIds {
		if id == aimodelID {
			return true
		}
	}
	return false
}

func addTaskNodeDescendent(curNode *entity.AimodelGraphNode, taskInfo *entity.AITaskDetailInfo) (*entity.AimodelGraphNode, error) {
	if taskInfo == nil {
		log.Error(fmt.Sprintf("AimodelEvolution Descendent addTaskNode taskInfo is nil; taskInfo=%v", curNode))
		return nil, nil
	}
	if curNode.Type != NODE_TYPE_MODEL {
		log.Error(fmt.Sprintf("AimodelEvolution Descendent addTaskNode curNode type is not model; curNode=%v", curNode))
		return nil, nil
	}

	for _, node := range curNode.Next {
		if node.TaskInfo != nil && node.TaskInfo.ID == taskInfo.ID {
			log.Info(fmt.Sprintf("AimodelEvolution Descendent addTaskNode task node already exists; aimodelID=%v; taskID=%v", curNode.AimodelInfo.ID, taskInfo.ID))
			return node, nil
		}
	}

	taskNode := &entity.AimodelGraphNode{
		Type:      NODE_TYPE_TASK,
		IsParent:  false,
		IsCurrent: false,
		IsPrivate: false,
		Visible:   true,
		TaskInfo:  buildTaskInfoNode(taskInfo),
	}
	curNode.Next = append(curNode.Next, taskNode)
	log.Info(fmt.Sprintf("AimodelEvolution Descendent add new task node; aimodelID=%v; taskID=%v", curNode.AimodelInfo.ID, taskInfo.ID))
	return taskNode, nil
}

func BuildDescendent(curNode *entity.AimodelGraphNode, aimodel *models.AiModelManage, doer *models.User) error {
	log.Info("AimodelEvolution BuildDescendent start; aimodelID=" + aimodel.ID)

	if curNode == nil {
		log.Error("AimodelEvolution Descendent curNode is nil;")
		return nil
	}

	childrens, err := models.QueryChildAiModelVague(aimodel.ID)
	if err != nil {
		log.Error("AimodelEvolution Descendent QueryChildAiModelVague err; err=" + err.Error())
		return err
	}
	if len(childrens) == 0 {
		log.Info("AimodelEvolution Descendent no child models; aimodelID=" + aimodel.ID)
		return nil
	}

	for _, c := range childrens {
		log.Info(fmt.Sprintf("AimodelEvolution Descendent start adding child; aimodelID=%v; childModelID=%v", aimodel.ID, c.ID))

		taskInfo, err0 := GetTaskInfo(c)
		if err0 != nil || taskInfo == nil {
			log.Error("AimodelEvolution Descendent GetTaskInfo err; err=" + err0.Error())
			continue
		}
		if isReal := isRealDescendent(taskInfo, aimodel.ID); !isReal {
			log.Info("AimodelEvolution Descendent child model is not a real descendent; child aimodelID=" + c.ID)
			continue
		}

		taskNode, err1 := addTaskNodeDescendent(curNode, taskInfo)
		if err1 != nil {
			log.Error("AimodelEvolution Descendent addTaskNodeDescendent err; child aimodelID=" + c.ID + "; err=" + err1.Error())
			continue
		}
		modelNode := &entity.AimodelGraphNode{
			Type:        NODE_TYPE_MODEL,
			IsParent:    false,
			IsCurrent:   false,
			IsPrivate:   c.IsPrivate,
			Visible:     subject_service.QueryAimodelVisibility(doer, c),
			AimodelInfo: buildAimodelInfoNode(c),
		}
		taskNode.Next = append(taskNode.Next, modelNode)
		err2 := BuildDescendent(modelNode, c, doer)
		if err2 != nil {
			log.Error("AimodelEvolution BuildDescendent recursion call err; child aimodelID=" + c.ID + "; err=" + err2.Error())
			continue
		}
	}
	return nil
}

func removeIpInfo(aimodel *models.AiModelManage) {
	reg, _ := regexp.Compile(`[[:digit:]]{1,3}\.[[:digit:]]{1,3}\.[[:digit:]]{1,3}\.[[:digit:]]{1,3}`)
	taskInfo := aimodel.TrainTaskInfo
	taskInfo = reg.ReplaceAllString(taskInfo, "")
	aimodel.TrainTaskInfo = taskInfo
}
