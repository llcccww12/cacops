export class TaskTmplTools {
  constructor() { }
  transformData(obj) {
    const runParameters = (obj.Parameters || []).map(item => {
      return {
        label: item.Label || item.label,
        value: item.Value || item.value,
      }
    });
    const models = (obj.PretrainModelList || [])
      .filter(item => item.ID && item.ModelName && item.OwnerName && item.ModelAlias)
      .map(item => ({
        id: item.ID,
        name: item.ModelName,
        owner_name: item.OwnerName,
        alias: item.ModelAlias
      }));
    const datasets = (obj.DatasetList || [])
      .filter(item => item.ID && item.DatasetName && item.OwnerName && item.DatasetAlias)
      .map(item => ({
        id: item.ID,
        name: item.DatasetName,
        owner_name: item.OwnerName,
        alias: item.DatasetAlias
      }));
    // const datasets = (obj.DatasetList || [])
    //   .map((item) => {
    //   if (item.ID && item.DatasetName && item.OwnerName && item.DatasetAlias) {
    //     return {
    //       id: item.ID,
    //       name: item.DatasetName,
    //       owner_name:item.OwnerName,
    //       alias: item.DatasetAlias,
    //     };
    //   }
      
    // });
    const data = {
      name: obj.Name || '',
      descr: obj.Description || '',
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
        image_id: obj.Image?.ImageID || '',
        image_name: obj.Image?.ImageName || '',
        image_url: obj.Image?.ImageUrl || obj.ImageUrl || '',
      },
      model: models,
      dataset: datasets,
      branchName: obj.BranchName || '',
      bootFile: obj.BootFile || '',
      runParameters: runParameters,
    };
    return data;
  }
  transformDataReverse(obj) {
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
      Image: {
        ImageID: obj.image.image_id || '',
        ImageName: obj.image.image_name || '',
        ImageUrl: obj.image.image_url,
      },
      PretrainModelList: modelList,
      DatasetList: datasetList,
      BranchName: obj.branchName,
      BootFile: obj.bootFile,
      Parameters: runParameterList,
    };
    return data;
  }
  transformDataTaskToTmpl(obj) {
    const modelList = (obj.pretrain_model_list || []).map(item => {
      return {
        ID: item.id,
        ModelName: item.name,
        ModelAlias: item.alias,
        OwnerName: item.owner_name,
      }
    });
    const datasetList = (obj.dataset_list || []).map(item => {
      return {
        ID: item.uuid,
        DatasetName: item.dataset_name,
        DatasetAlias: item.dataset_alias,
        OwnerName: item.owner_name,
      }
    });
    const runParameterList = (obj?.parameters?.parameter || []).map(item => {
      return {
        Label: item.label,
        Value: item.value,
      }
    });
    const specObj = obj.spec || {};
    const data = {
      Name: obj.name || '',
      Description: obj.descr || '',
      JobType: obj.job_type,
      Cluster: obj.cluster == 'OpenICloudbrainOne' || obj.cluster == 'OpenICloudbrainTwo' ? 'OpenI' : obj.cluster,
      ComputeSource: obj.compute_source,
      HasInternet: obj.has_internet || 2,
      VisualizeRequired: !!obj.visualize_required,
      AccCardsNum: specObj.acc_cards_num || 0,
      AccCardType: specObj.acc_card_type || '',
      CpuCores: specObj.cpu_cores || 0,
      MemGiB: specObj.mem_gi_b || 0,
      GPUMemGiB: specObj.gpu_mem_gi_b || 0,
      ShareMemGiB: specObj.share_mem_gi_b || 0,
      Image: {
        ImageID: obj.image_id || '',
        ImageName: obj.image_id ? obj.image_name : '',
        ImageUrl: obj.image_url,
      },
      PretrainModelList: modelList,
      DatasetList: datasetList,
      BranchName: obj.branch_name,
      BootFile: obj.boot_file,
      Parameters: runParameterList,
    };
    return data;
  }
}
