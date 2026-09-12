import { i18n } from '~/langs';
import { COMPUTER_RESOURCES_TITLE } from '~/pages/cloudbrain/configs';
import SparkMD5 from "spark-md5";

export class TaskTmplTools {
  constructor() { }
  transformDataTmplToForm(obj) {
    const runParameters = JSON.parse(obj.Parameters || '[]').map(item => {
      return {
        label: item.Label || item.label,
        value: item.Value || item.value,
      }
    });
    const models = (obj.ModelLists || [])
      .filter(item => !item.IsDeleted)
      .map(item => ({
        id: item.ModelID,
        name: item.ModelName,
        owner_name: item.OwnerName,
        alias: item.ModelAlias
      }));
    const datasets = (obj.DatasetLists || [])
      .filter(item => !item.IsDeleted)
      .map(item => ({
        id: item.DatasetID,
        name: item.DatasetName,
        owner_name: item.OwnerName,
        alias: item.DatasetAlias
      }));
    const data = {
      id: obj.ID,
      ownerId: obj.OwnerId,
      name: obj.Name || '',
      descr: obj.Description || '',
      tags: obj.Tags,
      isPrivate: obj.IsPrivate,
      taskType: obj.JobType || '',
      cluster: obj.Cluster || '',
      computeResource: obj.ComputeSource || '',
      networkType: obj.HasInternet == 1 ? 'no_internet' : obj.HasInternet == 2 ? 'has_internet' : 'has_internet',
      visualizeRequired: !!obj.VisualizeRequired,
      spec: '',
      acc_cards_num: obj.AccCardsNum || 0,
      acc_card_type: obj.AccCardType || '',
      cpu_cores: obj.CpuCores || 0,
      mem_gi_b: obj.MemGiB || 0,
      gpu_mem_gi_b: obj.GPUMemGiB || 0,
      share_mem_gi_b: obj.ShareMemGiB || 0,
      image: {
        image_id: obj.ImageID || '',
        image_name: obj.ImageName || '',
        image_url: obj.ImageUrl || '',
      },
      model: models,
      dataset: datasets,
      repoID: obj.RepoID,
      repoOwnerName: obj.RepoOwnerName,
      repoName: obj.RepoName,
      branchName: obj.BranchName || '',
      bootFile: obj.BootFile || '',
      runParameters: runParameters,
    };
    return data;
  }
  transformDataFormToTmpl(obj) {
    const modelList = (obj.model || []).map(item => {
      return {
        ID: item.id,
        ModelName: item.name,
        ModelAlias: item.alias,
        OwnerName: item.owner_name,
      }
    });
    const datasetList = (obj.dataset || []).map(item => {
      return {
        ID: item.id,
        DatasetName: item.name,
        DatasetAlias: item.alias,
        OwnerName: item.owner_name,
      }
    });
    const runParameterList = (obj.runParameters || []).map(item => {
      return {
        Label: item.label,
        Value: item.value,
      }
    });
    const data = {
      Name: obj.name,
      Description: obj.descr,
      Tags: obj.tags,
      IsPrivate: obj.isPrivate,
      JobType: obj.taskType,
      Cluster: obj.cluster,
      ComputeSource: obj.computeResource,
      HasInternet: obj.networkType == 'no_internet' ? 1 : obj.networkType == 'has_internet' ? 2 : 0,
      VisualizeRequired: obj.visualizeRequired,
      AccCardsNum: obj.acc_cards_num || 0,
      AccCardType: obj.acc_card_type || '',
      CpuCores: obj.cpu_cores || 0,
      MemGiB: obj.mem_gi_b || 0,
      GPUMemGiB: obj.gpu_mem_gi_b || 0,
      ShareMemGiB: obj.share_mem_gi_b || 0,
      ImageID: obj.image.image_id || '',
      ImageName: obj.image.image_name || '',
      ImageUrl: obj.image.image_url,
      RepoOwnerName: obj.repoOwnerName,
      RepoName: obj.repoName,
      BranchName: obj.branchName,
      BootFile: obj.bootFile,
      ModelIDs: modelList.map(item => item.ID),
      DatasetIDs: datasetList.map(item => item.ID),
      Parameters: JSON.stringify(runParameterList),
    };
    return data;
  }
  transformDataTaskToForm(obj) {
    const runParameters = (obj.parameters?.parameter || []).map(item => {
      return {
        label: item.Label || item.label,
        value: item.Value || item.value,
      }
    });
    const models = (obj.pretrain_model_list || [])
      .filter(item => !item.is_delete)
      .map(item => ({
        id: item.id,
        name: item.name,
        owner_name: item.owner_name,
        alias: item.alias
      }));
    const datasets = (obj.dataset_list || [])
      .filter(item => !item.IsDeleted)
      .map(item => ({
        id: item.uuid,
        name: item.dataset_name,
        owner_name: item.owner_name,
        alias: item.dataset_alias
      }));
    const spec = obj.spec || {};
    const data = {
      taskId: obj.id,
      name: obj.display_job_name || '',
      descr: obj.description || '',
      tags: [],
      isPrivate: false,
      taskType: obj.job_type || '',
      cluster: obj.cluster || '',
      computeResource: obj.compute_source || '',
      networkType: obj.has_internet == 1 ? 'no_internet' : obj.has_internet == 2 ? 'has_internet' : 'has_internet',
      visualizeRequired: !!obj.visualize_required,
      spec: '',
      acc_cards_num: spec.acc_cards_num || 0,
      acc_card_type: spec.acc_card_type || '',
      cpu_cores: spec.cpu_cores || 0,
      mem_gi_b: spec.mem_gi_b || 0,
      gpu_mem_gi_b: spec.gpu_mem_gi_b || 0,
      share_mem_gi_b: spec.share_mem_gi_b || 0,
      image: {
        image_id: obj.image_id || '',
        image_name: obj.image_name || '',
        image_url: obj.image_url || '',
      },
      model: models,
      dataset: datasets,
      repoID: obj.repo_id,
      repoOwnerName: obj.repo_owner_name,
      repoName: obj.repo_name,
      branchName: obj.branch_name || '',
      bootFile: obj.boot_file || '',
      runParameters: runParameters,
    };
    return data;
  }
}

export const TmplTaskTypes = [{
  k: 'DEBUG',
  v: i18n.t('cloudbrainObj.tabTitDebug'),
  cluster: 'C2Net',
  computerResouce: 'NPU',
  desc: i18n.t('cloudbrainObj.debugTaskDesc'),
  descLong: i18n.t('cloudbrainObj.debugTaskDescLong'),
}, {
  k: 'TRAIN',
  v: i18n.t('cloudbrainObj.tabTitTrain'),
  cluster: 'C2Net',
  computerResouce: 'NPU',
  desc: i18n.t('cloudbrainObj.trainTaskDesc'),
  descLong: i18n.t('cloudbrainObj.trainTaskDescLong'),
}, {
  k: 'ONLINEINFERENCE',
  v: i18n.t('cloudbrainObj.tabTitOnlineInference'),
  cluster: 'C2Net',
  computerResouce: 'GPU',
  desc: i18n.t('cloudbrainObj.onlineinferTaskDesc'),
  descLong: i18n.t('cloudbrainObj.onlineinferTaskDescLong'),
}, {
  k: 'GENERAL',
  v: i18n.t('cloudbrainObj.tabTitGeneral'),
  cluster: 'C2Net',
  computerResouce: 'GPU',
  desc: i18n.t('cloudbrainObj.generalTaskDesc'),
  descLong: i18n.t('cloudbrainObj.generalTaskDescLong'),
}, {
  k: 'HPC',
  v: i18n.t('superComputeTask'),
  cluster: 'C2Net',
  computerResouce: 'CPU',
  desc: i18n.t('cloudbrainObj.inferenceTaskDesc'),
  descLong: i18n.t('cloudbrainObj.inferenceTaskDescLong'),
}];

export const TmplComputerResouces = [...COMPUTER_RESOURCES_TITLE];

const GradientColors = [
  ['rgba(178,195,255,0.85)', 'rgba(160,246,246,1)'],
  ['rgba(189,231,255,0.85)', 'rgba(200,202,255,1)'],
  ['rgba(255,237,189,0.85)', 'rgba(214,255,200,1)'],
  ['rgba(217,186,255,0.85)', 'rgba(255,233,201,1)'],
  ['rgba(200,255,232,1)', 'rgba(189,235,255,0.85)'],
  ['rgba(242,255,189,0.85)', 'rgba(200,205,255,1)'],
  ['rgba(203,178,255,0.85)', 'rgba(203,246,160,1)'],
  ['rgba(255,213,189,0.85)', 'rgba(200,255,232,1)'],
  ['rgba(255,189,237,0.85)', 'rgba(200,229,255,1)'],
  ['rgba(255,186,186,0.85)', 'rgba(255,233,201,1)'],
  ['rgba(255,250,200,1)', 'rgba(189,215,255,0.85)'],
  ['rgba(189,255,196,0.85)', 'rgba(200,255,239,1)']
]

export const getGradientColor = (name) => {
  const nameHash = SparkMD5.hash(name);
  const colourIndex = (nameHash.charCodeAt(0) + nameHash.charCodeAt(1) +
    nameHash.charCodeAt(2) + nameHash.charCodeAt(3)) % 12;
  return GradientColors[colourIndex];
};
