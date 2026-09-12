import { i18n } from '~/langs';

export const SOURCE_TYPE = [{ k: 'ACCOMPLISH_TASK', v: i18n.t('accomplishTask') }, { k: 'ADMIN_OPERATE', v: i18n.t('adminOperate') }, { k: 'RUN_CLOUDBRAIN_TASK', v: i18n.t('runCloudBrainTask') }];
export const CONSUME_STATUS = [{ k: 'OPERATING', v: i18n.t('operating') }, { k: 'SUCCEEDED', v: i18n.t('succeeded') }];
export const POINT_ACTIONS = [
  { k: 'CreatePublicRepo', v: i18n.t('createPublicProject') }, { k: 'CreateIssue', v: i18n.t('dailyPutforwardTasks') }, { k: 'CreatePullRequest', v: i18n.t('dailyPR') }, { k: 'CommentIssue', v: i18n.t('comment') }, { k: 'UploadAttachment', v: i18n.t('uploadDatasetFile') }, { k: 'CreateNewModelTask', v: i18n.t('importNewModel') }, { k: 'BindWechat', v: i18n.t('completeWechatCodeScanningVerification') },
  { k: 'CreateCloudbrainTask', v: i18n.t('dailyRunCloudbrainTasks') }, { k: 'DatasetRecommended', v: i18n.t('datasetRecommendedByThePlatform') }, { k: 'CreateImage', v: i18n.t('submitNewPublicImage') }, { k: 'ImageRecommend', v: i18n.t('imageRecommendedByThePlatform') }, { k: 'ChangeUserAvatar', v: i18n.t('firstChangeofAvatar') }, { k: 'PushCommits', v: i18n.t('dailyCommit') },
  { k: 'TaskInviteFriendRegister', v: i18n.t('user.inviteFriends') }, { k: 'TaskCreateDataset', v: i18n.t('dailyCreateDataset') }, { k: 'TaskCreateAimodel', v: i18n.t('dailyCreateAimodel1') }, { k: 'TaskAimodelRecommended', v: i18n.t('aimodelRecommendedByThePlatform') }
];
export const JOB_TYPE = [
  { k: 'DEBUG', v: i18n.t('debugTask'), train_type: 'Notebook', alias: i18n.t('debugTask') },
  { k: 'TRAIN', v: i18n.t('trainTask'), train_type: 'TrainJob', alias: i18n.t('trainTask') },
  { k: 'INFERENCE', v: i18n.t('inferenceTask'), train_type: 'TrainJob', alias: i18n.t('inferenceTask') },
  { k: 'BENCHMARK', v: i18n.t('benchmarkTask'), alias: i18n.t('benchmarkTask') },
  { k: 'ONLINEINFERENCE', v: i18n.t('onlineinfer'), alias: i18n.t('onlineinferTask'), train_type: 'Notebook' },
  { k: 'HPC', v: i18n.t('superComputeTask'), alias: i18n.t('superComputeTask') },
  { k: 'GENERAL', v: i18n.t('generalTask'), train_type: 'Notebook', alias: i18n.t('generalTask') },
  { k: 'MODELEXPERIENCE', v: i18n.t('modelManage.onlineInference'), alias: i18n.t('modelManage.onlineInferenceTask') },
  { k: 'FINETUNE', v: i18n.t('modelSquare.sftFinetune'), alias: i18n.t('modelSquare.sftFinetuneTask') },
  { k: 'SDFINETUNE', v: i18n.t('modelSquare.sdModelFinetuen'), alias: i18n.t('modelSquare.sdModelFinetuenTask') },
  { k: 'ComfyuiExperience', v: i18n.t('modelSquare.cvComfyui'), alias: i18n.t('modelSquare.cvComfyuiTask') },
  { k: 'EVAL', v: i18n.t('modelSquare.modelPerEvaluate'), alias: i18n.t('modelSquare.modelEvaluateTask') },
];
export const BenchmarkTypeList = ['BENCHMARK', 'SIM2BRAIN_SNN', 'SNN4ECOSET', 'SNN4IMAGENET', 'BRAINSCORE', 'MODELSAFETY'];
// 资源管理
export const CLUSTERS = [{ k: 'OpenI', v: i18n.t('resourcesManagement.OpenI') }, { k: 'C2Net', v: i18n.t('resourcesManagement.C2Net') }, { k: 'IFLYTEKTraining', v: i18n.t('resourcesManagement.IFLYTEKTraining') }];
export const AI_CENTER = [{ k: 'OpenIOne', v: i18n.t('resourcesManagement.OpenIOne') }, { k: 'OpenITwo', v: i18n.t('resourcesManagement.OpenITwo') }, { k: 'OpenIChengdu', v: i18n.t('resourcesManagement.OpenIChengdu') }, { k: 'pclcci', v: i18n.t('resourcesManagement.pclcci') }, { k: 'hefei', v: i18n.t('resourcesManagement.hefeiCenter') }, { k: 'xuchang', v: i18n.t('resourcesManagement.xuchangCenter') }, { k: 'OpenI-huoshi', v: i18n.t('resourcesManagement.huoshi') }];
export const COMPUTER_RESOURCES = [{ k: 'CPU', v: 'CPU' }, { k: 'GPU', v: 'GPU' }, { k: 'NPU', v: 'NPU' }, { k: 'GCU', v: 'GCU' }, { k: 'MLU', v: 'MLU' }, { k: 'DCU', v: 'DCU' }, { k: 'ILUVATAR-GPGPU', v: 'ILUVATAR-GPGPU' }, { k: 'METAX-GPGPU', v: 'METAX-GPGPU' }, { k: 'BIREN-GPU', v: 'BIREN-GPU' },];
export const COMPUTER_RESOURCES_COLORS = {
  'GPU': '#4fb62f',
  'NPU': '#c31d20',
  'GCU': '#e73828',
  'DCU': '#b01f24',
  'MLU': '#0077ed',
  'ILUVATAR-GPGPU': '#0038bd',
  'METAX-GPGPU': '#5c246a',
  'BIREN-GPU': '#50c878',
};
export const ACC_CARD_TYPE = [{ k: 'T4', v: 'T4' }, { k: 'A100', v: 'A100' }, { k: 'V100', v: 'V100' },
{ k: 'ASCEND910', v: 'Ascend 910' }, { k: 'ASCEND-D910B', v: 'Ascend 910B' }, { k: 'ASCEND-910C', v: 'Ascend 910C' },
{ k: 'MLU270', v: 'MLU270' }, { k: 'MLU290', v: 'MLU290' }, { k: 'RTX3080', v: 'RTX3080' }, { k: '3090', v: '3090' }, { k: '4090', v: '4090' },
{ k: 'ENFLAME-T20', v: 'ENFLAME-T20' }, { k: 'ENFLAME-I20', v: 'ENFLAME-I20' }, { k: 'DCU', v: 'DCU' }, { k: 'Z100L', v: 'Z100L' }, { k: 'K100_AI', v: 'K100_AI' }, { k: 'BI-V100', v: 'BI-V100' }, { k: 'MR-V100', v: 'MR-V100' }, { k: 'N100', v: 'N100' }, { k: 'N260', v: 'N260' }, { k: 'C500', v: 'C500' }, { k: 'L20', v: 'L20' }, { k: 'S60', v: 'S60' }, { k: 'BIREN106M', v: 'BIREN106M' }, { k: 'BW1000', v: 'BW1000' }];

export const SPECIFICATION_STATUS = [{ k: '1', v: i18n.t('resourcesManagement.willOnShelf') }, { k: '2', v: i18n.t('resourcesManagement.onShelf') }, { k: '3', v: i18n.t('resourcesManagement.offShelf') }];
export const NETWORK_TYPE = [{ k: 1, v: `${i18n.t('cloudbrainObj.networkType')}(${i18n.t('cloudbrainObj.noInternet')})` }, { k: 2, v: `${i18n.t('cloudbrainObj.networkType')}(${i18n.t('cloudbrainObj.hasInternet')})` }];
export const NETWORK_TYPE_VALUE = [{ k: 1, v: i18n.t('cloudbrainObj.noInternet') }, { k: 2, v: i18n.t('cloudbrainObj.hasInternet') }];
export const OPERATION_TYPE = [{ k: 'edit', v: i18n.t('logManagement.resourceSpecificationEdit') }, { k: 'on-shelf', v: i18n.t('logManagement.resourceSpecificationOnshelf') }, { k: 'off-shelf', v: i18n.t('logManagement.resourceSpecificationOffshelf') }, { k: 'create', v: i18n.t('logManagement.resourceSpecificationAddition') }, { k: 'auto-update', v: i18n.t('logManagement.resourceSpecificationAutoUpdate') }]

// 模型
export const MODEL_ENGINES = [{ k: 0, v: 'PyTorch' }, { k: 1, v: 'TensorFlow' }, { k: 2, v: 'MindSpore' }, { k: 4, v: 'PaddlePaddle' }, { k: 5, v: 'OneFlow' }, { k: 6, v: 'MXNet' }, { k: 3, v: 'Other' }];


export const NEW_JOB_TYPE = [
  { k: 'DEBUG', v: i18n.t('debugTask') },
  { k: 'TRAIN', v: i18n.t('trainTask') },
  { k: 'ONLINEINFERENCE', v: i18n.t('onlineinfer') },
  { k: 'HPC', v: i18n.t('superComputeTask') },
  { k: 'GENERAL', v: i18n.t('generalTask') },
  { k: 'MODELEXPERIENCE', v: i18n.t('modelManage.onlineInference') },
  { k: 'FINETUNE', v: i18n.t('modelSquare.sftFinetune') },
  { k: 'SDFINETUNE', v: i18n.t('modelSquare.sdModelFinetuen') },
  { k: 'ComfyuiExperience', v: i18n.t('modelSquare.cvComfyui') },
  { k: 'EVAL', v: i18n.t('modelSquare.modelPerEvaluate') },
];

export const NEW_JOB_TYPE_OBJ = {
  'DEBUG': [],
  'TRAIN': [],
  'ONLINEINFERENCE': [],
  'HPC': [],
  'GENERAL': [],
  'MODELEXPERIENCE': [],
  'FINETUNE': [],
  'SDFINETUNE': [],
  "ComfyuiExperience": [],
  "EVAL": []
}
