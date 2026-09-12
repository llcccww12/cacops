
import { formatDate } from 'element-ui/lib/utils/date-util';
import { SOURCE_TYPE, CONSUME_STATUS, POINT_ACTIONS, JOB_TYPE, ACC_CARD_TYPE } from '~/const';
import { i18n } from '~/langs';
import { getListValueWithKey } from '~/utils';



const getSourceType = (key) => {
  const find = SOURCE_TYPE.filter(item => item.k === key);
  return find.length ? find[0].v : key;
};

const getConsumeStatus = (key) => {
  const find = CONSUME_STATUS.filter(item => item.k === key);
  return find.length ? find[0].v : key;
};

const getPointAction = (key) => {
  const find = POINT_ACTIONS.filter(item => item.k === key);
  return find.length ? find[0].v : key;
};

const getJobType = (key) => {
  const find = JOB_TYPE.filter(item => item.k === key);
  return find.length ? find[0].alias : key;
};

const jobLinkGenerators = {
  MODELEXPERIENCE: (id) => `/modelbase/experience/detail/${id}`,
  FINETUNE: (id) => `/modelbase/nlp/sft/detail/${id}`,
  SDFINETUNE: (id) => `/modelbase/cv/sft/detail//${id}`,
  ComfyuiExperience: (id) => `/modelbase/cv/comfyui/detail/${id}`,
  EVAL: (id) => `/modelbase/eval/evaluate/detail/${id}`,
  DEFAULT: (id) => `/cloudbrains/detail/${id}`
};

const getJobTypeLink = (record, type) => {
  const cloudbrain = type === 'INCREASE' ? record.Action?.Cloudbrain : record.Cloudbrain;
  console.log("cloudbrain",cloudbrain)
  const generateJobLink = jobLinkGenerators[cloudbrain?.JobType] || jobLinkGenerators.DEFAULT
  return generateJobLink(cloudbrain?.ID);
};

const renderSpecStr = (spec, showPoint) => {
  var ngpu = `${spec.ComputeResource}: ${spec.AccCardsNum + '*' + getListValueWithKey(ACC_CARD_TYPE, spec.AccCardType)}`;
  var gpuMemStr = spec.GPUMemGiB != 0 ? `(${i18n.t('resourcesManagement.gpuMem')}: ${spec.GPUMemGiB}GB)` : '';
  var memStr = spec.MemGiB != 0 ? `, ${i18n.t('resourcesManagement.mem')}: ${spec.MemGiB}GB` : '';
  var sharedMemStr = spec.ShareMemGiB != 0 ? `, ${i18n.t('resourcesManagement.shareMem')}: ${spec.ShareMemGiB}GB` : '';
  var workServerNum = spec.workServerNumber;
  var workServerNumStr = showPoint && workServerNum != 1 && spec.UnitPrice != 0 ? '*' + workServerNum + i18n.t('resourcesManagement.node') : '';
  var pointStr = showPoint ? `, ${spec.UnitPrice == 0 ? i18n.t('resourcesManagement.free') : spec.UnitPrice.toFixed(2) + i18n.t('resourcesManagement.point_hr') + workServerNumStr}` : '';
  var specStr = `${ngpu}${gpuMemStr}, CPU: ${spec.CpuCores}${memStr}${sharedMemStr}${pointStr}`;
  return specStr;
};

export const getRewardPointRecordInfo = (record) => {
  const out = {
    sn: record.SerialNo,
    date: formatDate(new Date(record.LastOperateDate * 1000), 'yyyy-MM-dd HH:mm:ss'),
    _status: record.Status,
    status: getConsumeStatus(record.Status) || '--',
    statusColor: record.Status === 'OPERATING' ? 'rgb(33, 186, 69)' : '',
    _sourceType: record.SourceType,
    sourceType: getSourceType(record.SourceType),
    duration: record?.Cloudbrain?.Duration || '--',
    taskName: record?.Cloudbrain?.DisplayJobName || '--',
    taskId: record?.Cloudbrain?.ID,
    action: record?.SourceTemplateId ? getPointAction(record.SourceTemplateId) : '--',
    remark: record.Remark,
    amount: record.Amount,
  };
  if (record.Action?.Cloudbrain?.JobType === 'BRAINSCORE' || record.Action?.Cloudbrain?.JobType === 'SNN4IMAGENET'
    || record.Action?.Cloudbrain?.JobType === 'SNN4ECOSET' || record.Action?.Cloudbrain?.JobType === 'SIM2BRAIN_SNN'
    || record.Action?.Cloudbrain?.JobType === 'MODELSAFETY') {
    record.Action.Cloudbrain.oJobType = record.Action?.Cloudbrain?.JobType;
    record.Action.Cloudbrain.JobType = 'BENCHMARK';
  }
  if (record?.Cloudbrain?.JobType === 'BRAINSCORE' || record?.Cloudbrain?.JobType === 'SNN4IMAGENET'
    || record?.Cloudbrain?.JobType === 'SNN4ECOSET' || record?.Cloudbrain?.JobType === 'SIM2BRAIN_SNN'
    || record?.Cloudbrain?.JobType === 'MODELSAFETY') {
    record.Cloudbrain.oJobType = record?.Cloudbrain?.JobType;
    record.Cloudbrain.JobType = 'BENCHMARK';
  }
  if (record.OperateType === 'INCREASE') {
    if (record.SourceType === 'ADMIN_OPERATE') {
      out.remark = record.Remark;
    } else if (record.SourceType === 'ACCOMPLISH_TASK') {
      switch (record?.SourceTemplateId) {
        case 'CreatePublicRepo': // 创建公开项目 - 创建了项目OpenI/aiforge
          out.remark = record.Action ? `${i18n.t('createdRepository')}<a href="${record.Action.RepoLink}" rel="nofollow">${record.Action.ShortRepoFullDisplayName}</a>`
            : `${getPointAction(record.SourceTemplateId)}(${i18n.t('repositoryWasDel')})`;
          break;
        case 'CreateIssue': // 每日提出任务 - 创建了任务PCL-Platform.Intelligence/AISynergy#19
          out.remark = record.Action ? `${i18n.t('openedIssue')}<a href="${record.Action.RepoLink}/issues/${record.Action.IssueInfos[0]}" rel="nofollow">${record.Action.ShortRepoFullDisplayName}#${record.Action.IssueInfos[0]}</a>`
            : `${getPointAction(record.SourceTemplateId)}(${i18n.t('repositoryWasDel')})`;
          break;
        case 'CreatePullRequest': // 每日提出PR - 创建了合并请求OpenI/aiforge#1
          out.remark = record.Action ? `${i18n.t('createdPullRequest')}<a href="${record.Action.RepoLink}/pulls/${record.Action.IssueInfos[0]}" rel="nofollow">${record.Action.ShortRepoFullDisplayName}#${record.Action.IssueInfos[0]}</a>`
            : `${getPointAction(record.SourceTemplateId)}(${i18n.t('repositoryWasDel')})`;
          break;
        case 'CommentIssue': // 发表评论 - 评论了任务PCL-Platform.Intelligence/AISynergy#19
          out.remark = record.Action ? `${i18n.t('commentedOnIssue')}<a href="${record.Action.CommentLink}" rel="nofollow">${record.Action.ShortRepoFullDisplayName}#${record.Action.IssueInfos[0]}</a>`
            : `${getPointAction(record.SourceTemplateId)}(${i18n.t('repositoryWasDel')})`;
          break;
        case 'UploadAttachment': // 上传数据集文件 - 上传了数据集文件MMISTData.zip
          out.remark = record.Action ? `${i18n.t('uploadDataset')}<a href="${record.Action.RepoLink}/datasets" rel="nofollow">${record.Action?.RefName}</a>`
            : `${getPointAction(record.SourceTemplateId)}(${i18n.t('repositoryWasDel')})`;
          break;
        case 'CreateNewModelTask': // 导入新模型 - 导入了新模型resnet50_qx7l
          out.remark = record.Action ? `${i18n.t('createdNewModel')}<a href="${record.Action.RepoLink}/modelmanage/model_readme_tmpl?name=${record.Action?.RefName}" rel="nofollow">${record.Action?.RefName}</a>`
            : `${getPointAction(record.SourceTemplateId)}(${i18n.t('repositoryWasDel')})`;
          break;
        case 'BindWechat': // 完成微信扫码验证 - 首次绑定微信奖励
          out.remark = record.Action ? `${i18n.t('firstBindingWechatRewards')}` : `${getPointAction(record.SourceTemplateId)}(${i18n.t('userAccountWasDel')})`;
          break;
        case 'CreateCloudbrainTask': // 每日运行云脑任务 - 创建了(CPU/GPU/NPU)类型(调试/训练/推理/评测)任务tangl202204131431995
          out.remark = record.Action ? `${i18n.t('created')}${record.Action?.Cloudbrain?.ComputeResource}${i18n.t('type')}${getJobType(record.Action?.Cloudbrain?.JobType)} <a href="${getJobTypeLink(record, 'INCREASE')}" rel="nofollow">${record.Action.RefName}</a>`
            : `${getPointAction(record.SourceTemplateId)}(${i18n.t('repositoryWasDel')})`;
          break;
        case 'DatasetRecommended': // 数据集被平台推荐 - 数据集XXX被设置为推荐数据集
          if (record.Action.Dataset) {
            out.remark = record.Action ? `${i18n.t('dataset')}<a href="/datasets/detail/${record.Action.Dataset.OwnerName}/${record.Action.Dataset.Name}" rel="nofollow"> ${record.Action.Dataset?.Alias ? record.Action.Dataset?.Alias : record.Action.Dataset.Name} </a>${i18n.t('setAsRecommendedDataset')}`
              : `${getPointAction(record.SourceTemplateId)}(${i18n.t('repositoryWasDel')})`;
            break;
          } else {
            out.remark = record.Action ? `${i18n.t('dataset')}<a href="javascript:void(0)" rel="nofollow">${record.Action.Content && record.Action.Content.split('|')[1]}</a>${i18n.t('setAsRecommendedDataset')}`
              : `${getPointAction(record.SourceTemplateId)}(${i18n.t('repositoryWasDel')})`;
            break;
          }
        case 'TaskCreateDataset': // 创建数据集 - 创建了数据集xxxxx
          out.remark = record.Action ? `${i18n.t('dailyCreateDataset1')} <a href="/datasets/detail/${record.Action.Dataset.OwnerName}/${record.Action.Dataset.Name}" rel="nofollow">${record.Action.Dataset?.Alias ? record.Action.Dataset?.Alias : record.Action.Dataset.Name}</a>`
            : `${getPointAction(record.SourceTemplateId)}(${i18n.t('repositoryWasDel')})`;
          break;
        //
        case 'TaskAimodelRecommended': // 模型被平台推荐 - 模型XXX被设置为推荐数据集
          out.remark = record.Action ? `${i18n.t('repos.model')}<a href="/models/detail/${record.Action.Aimodel.OwnerName}/${record.Action.Aimodel.Name}" rel="nofollow"> ${record.Action.Aimodel?.Alias ? record.Action.Aimodel?.Alias : record.Action.Aimodel.Name} </a>${i18n.t('setAsRecommendedAimodel')}`
            : `${getPointAction(record.SourceTemplateId)}(${i18n.t('repositoryWasDel')})`;
          break;
        case 'TaskCreateAimodel': // 创建模型 - 创建了模型xxxxx
          out.remark = record.Action ? `${i18n.t('dailyCreateAimodel1')} <a href="/models/detail/${record.Action.Aimodel.OwnerName}/${record.Action.Aimodel.Name}" rel="nofollow">${record.Action.Aimodel?.Alias ? record.Action.Aimodel?.Alias : record.Action.Aimodel.Name}</a>`
            : `${getPointAction(record.SourceTemplateId)}(${i18n.t('repositoryWasDel')})`;
          break;
        case 'CreateImage': // 提交新公开镜像 - 提交了镜像jiangxiang_ceshi_tang03
          out.remark = record.Action ? `${i18n.t('committedImage')}<span style="font-weight:bold;">${record.Action.Content && record.Action.Content.split('|')[1]}</span>`
            : `${getPointAction(record.SourceTemplateId)}(${i18n.t('userAccountWasDel')})`;
          break;
        case 'ImageRecommend': // 镜像被平台推荐 - 镜像XXX被设置为推荐镜像
          out.remark = record.Action ? `${i18n.t('image')}<span style="font-weight:bold;">${record.Action.Content && record.Action.Content.split('|')[1]}</span>${i18n.t('setAsRecommendedImage')}`
            : `${getPointAction(record.SourceTemplateId)}(${i18n.t('userAccountWasDel')})`;
          break;
        case 'ChangeUserAvatar': // 首次更换头像 - 更新了头像
          out.remark = record.Action ? `${i18n.t('updatedAvatar')}` : `${getPointAction(record.SourceTemplateId)}(${i18n.t('userAccountWasDel')})`;
          break;
        case 'TaskInviteFriendRegister': // 邀请好友
          if (record.Action) {
            out.remark = record.Action.Data && record.Action.Data.InvitedUserNotExists == false
              ? `${i18n.t('invitedFriend')}<a href="/${record.Action.Data.InvitedUserName}" rel="nofollow">${record.Action.Data.InvitedUserName}</a>`
              : `${i18n.t('invitedFriend')}${record.Action.Data?.InvitedUserName}(${i18n.t('userAccountWasDel')})`;
          } else {
            out.remark = `${getPointAction(record.SourceTemplateId)}(${i18n.t('userAccountWasDel')})`
          }
          break;
        case 'PushCommits': // 每日commit - 推送了xxxx分支的代码到OpenI/aiforge
          if (record?.Action) {
            const opType = record.Action.OpType;
            if (opType == 5) {
              const words = record.Action.RefName.split('/');
              const branch = words[words.length - 1];
              out.remark = `${i18n.t('pushedBranch', {
                branch: `<a href="${record.Action.RepoLink}/src/branch/${branch}" rel="nofollow">${branch}</a>`
              })}<a href="${record.Action.RepoLink}" rel="nofollow">${record.Action.ShortRepoFullDisplayName}</a>`;
            } else if (opType == 9) {
              const words = record.Action.RefName.split('/');
              const tag = words[words.length - 1];
              out.remark = `${i18n.t('pushedTag', {
                tag: `<a href="${record.Action.RepoLink}/src/tag/${tag}" rel="nofollow">${tag}</a>`
              })}<a href="${record.Action.RepoLink}" rel="nofollow">${record.Action.ShortRepoFullDisplayName}</a>`;
            } else if (opType == 16) {
              const words = record.Action.RefName.split('/');
              const tag = words[words.length - 1];
              out.remark = `${i18n.t('deleteTag', {
                repo: `<a href="${record.Action.RepoLink}" rel="nofollow">${record.Action.ShortRepoFullDisplayName}</a>`,
                tag: tag,
              })}`;
            } else if (opType == 17) {
              const words = record.Action.RefName.split('/');
              const branch = words[words.length - 1];
              out.remark = `${i18n.t('deleteBranch', {
                repo: `<a href="${record.Action.RepoLink}" rel="nofollow">${record.Action.ShortRepoFullDisplayName}</a>`,
                branch: branch,
              })}`;
            }
          } else {
            out.remark = `${getPointAction(record.SourceTemplateId)}(${i18n.t('repositoryWasDel')})`;
          }
          break;
        default:
          break;
      }
    } else if (record.SourceType === 'RUN_CLOUDBRAIN_TASK') {
      //
    }
    if (record.LossAmount !== 0) {
      out.amount = record.Amount;
      out.remark += `${out.remark ? i18n.t(';') : ''}${i18n.t('dailyMaxTips')}`;
    }
  } else if (record.OperateType === 'DECREASE') {
    if (record.SourceType === 'ADMIN_OPERATE') {
      out.remark = record.Remark;
    } else if (record.SourceType === 'ACCOMPLISH_TASK') {
      //
    } else if (record.SourceType === 'RUN_CLOUDBRAIN_TASK') {
      out.taskName = `<a href="${getJobTypeLink(record, 'DECREASE')}" rel="nofollow">${record?.Cloudbrain?.DisplayJobName}</a>`
      // if (!record.Cloudbrain?.RepoFullName && record.Cloudbrain?.RepoID) {
      //   out.taskName = `${record?.Cloudbrain?.DisplayJobName}(${i18n.t('repositoryWasDel')})`;
      // } else if(record.Cloudbrain?.RepoFullName && record.Cloudbrain?.RepoID){
        
      // } else {
      //   out.taskName = record?.Cloudbrain?.DisplayJobName
      // }
      const resourceSpec = record?.Cloudbrain?.ResourceSpec;
      if (resourceSpec) {
        resourceSpec.workServerNumber = record?.Cloudbrain?.WorkServerNumber || 1;
        out.remark = `【${getJobType(record?.Cloudbrain?.JobType)}】【${renderSpecStr(resourceSpec, true)}】`;
      }
    }
  }
  return out;
};
