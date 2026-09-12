import { i18n } from '~/langs';

export const COMPUTER_RESOURCES_TITLE = [{ k: 'GPU', v: i18n.t('computeResourceTitle.GPU') }, { k: 'NPU', v: i18n.t('computeResourceTitle.NPU') }, { k: 'GCU', v: i18n.t('computeResourceTitle.GCU') }, { k: 'MLU', v: i18n.t('computeResourceTitle.MLU') }, { k: 'DCU', v: i18n.t('computeResourceTitle.DCU') }, { k: 'ILUVATAR-GPGPU', v: i18n.t('computeResourceTitle.ILUVATAR-GPGPU') }, { k: 'METAX-GPGPU', v: i18n.t('computeResourceTitle.METAX-GPGPU') }, { k: 'BIREN-GPU', v: i18n.t('computeResourceTitle.BIREN-GPU') }];

export const CreatePageConfigs = {
  // 调试任务
  'DEBUG': [{
    'C2Net': [{
      'GPU': [{
        form: {
          taskName: { required: true, },
          taskDescr: { required: false, },
          branchName: { required: false, },
          model: { required: false, multiple: true, useExceedSize: true },
          imagev1: { required: true, type: -1 },
          dataset: { required: false, useExceedSize: true },
          networkType: { required: true },
          spec: {},
          runTimeLimit: { required: true },
          repo: { required: false },
        },
        showMindTorchHelper: true,
        hideCluster: true,
      }],
      'NPU': [{
        form: {
          taskName: { required: true, },
          taskDescr: { required: false, },
          branchName: { required: false, },
          // imagev2: { required: true, relatedSpec: true },
          imagev1: { required: true, type: 2, useId: true },
          model: { required: false, multiple: true, useExceedSize: true },
          dataset: { required: false, useExceedSize: true },
          networkType: { required: true },
          spec: { required: true },
          runTimeLimit: { required: true },
          repo: { required: false },
        },
        hideCluster: true,
      }],
      'GCU': [{
        form: {
          taskName: { required: true, },
          taskDescr: { required: false, },
          branchName: { required: false, },
          model: { required: false, multiple: true, useExceedSize: true },
          imagev1: { required: true, type: -1 },
          dataset: { required: false, useExceedSize: true },
          networkType: { required: true },
          spec: { required: true },
          runTimeLimit: { required: true },
          repo: { required: false },
        },
        hideCluster: true,
      }],
      'MLU': [{
        form: {
          taskName: { required: true, },
          taskDescr: { required: false, },
          branchName: { required: false, },
          model: { required: false, multiple: true, useExceedSize: true },
          imagev1: { required: true, type: -1 },
          dataset: { required: false, useExceedSize: true },
          networkType: { required: true },
          spec: { required: true },
          runTimeLimit: { required: true },
          repo: { required: false },
        },
        hideCluster: true,
      }],
      'DCU': [{
        form: {
          taskName: { required: true, },
          taskDescr: { required: false, },
          branchName: { required: false, },
          model: { required: false, multiple: true, useExceedSize: true },
          imagev1: { required: true, type: -1 },
          dataset: { required: false, useExceedSize: true },
          networkType: { required: true },
          spec: { required: true },
          runTimeLimit: { required: true },
          repo: { required: false },
        },
        hideCluster: true,
      }],
      'ILUVATAR-GPGPU': [{
        form: {
          taskName: { required: true, },
          taskDescr: { required: false, },
          branchName: { required: false, },
          model: { required: false, multiple: true, useExceedSize: true },
          imagev1: { required: true, type: -1 },
          dataset: { required: false, useExceedSize: true },
          networkType: { required: true },
          spec: { required: true },
          runTimeLimit: { required: true },
          repo: { required: false },
        },
        hideCluster: true,
      }],
      'METAX-GPGPU': [{
        form: {
          taskName: { required: true, },
          taskDescr: { required: false, },
          branchName: { required: false, },
          model: { required: false, multiple: true, useExceedSize: true },
          imagev1: { required: true, type: -1 },
          dataset: { required: false, useExceedSize: true },
          networkType: { required: true },
          spec: { required: true },
          runTimeLimit: { required: true },
          repo: { required: false },
        },
        hideCluster: true,
      }],
      'BIREN-GPU': [{
        form: {
          taskName: { required: true, },
          taskDescr: { required: false, },
          branchName: { required: false, },
          model: { required: false, multiple: true, useExceedSize: true },
          imagev1: { required: true, type: -1 },
          dataset: { required: false, useExceedSize: true },
          networkType: { required: true },
          spec: { required: true },
          runTimeLimit: { required: true },
          repo: { required: false },
        },
        hideCluster: true,
      }],
    }]
  }],
  // 训练任务
  'TRAIN': [{
    'C2Net': [{
      'GPU': [{
        form: {
          taskName: { required: true, },
          taskDescr: { required: false, },
          branchName: { required: true, },
          model: { multiple: true },
          imagev1: { required: true, type: -1 },
          bootFile: { required: true, sampleUrl: 'https://openi.pcl.ac.cn/OpenIOSSG/OpenI_Cloudbrain_Example/src/branch/master/gpu_mnist_example/train.py' },
          dataset: { required: true },
          runParameters: { required: false },
          networkType: { required: true },
          visualization: { required: false },
          spec: { required: true },
          workServerNum: { required: true },
          repo: { required: true },
        },
        showMindTorchHelper: true,
        hideCluster: true,
      }],
      'NPU': [{
        form: {
          taskName: { required: true, },
          taskDescr: { required: false, },
          branchName: { required: true, },
          model: { required: false, multiple: true },
          // imagev2: { required: true, relatedSpec: true },
          imagev1: { required: true, type: 2, useId: true },
          bootFile: { required: true, sampleUrl: 'https://openi.pcl.ac.cn/OpenIOSSG/OpenI_Cloudbrain_Example/src/branch/master/npu_mnist_example/train.py' },
          dataset: { required: true },
          runParameters: { required: false },
          networkType: { required: true },
          visualization: { required: false },
          spec: { required: true },
          workServerNum: { required: true },
          repo: { required: true },
        },
        hideCluster: true,
      }],
      'GCU': [{
        form: {
          taskName: { required: true, },
          taskDescr: { required: false, },
          branchName: { required: true, },
          model: { required: false, multiple: true },
          imagev1: { required: true, type: -1 },
          bootFile: { required: true, sampleUrl: 'https://openi.pcl.ac.cn/OpenIOSSG/OpenI_Cloudbrain_Example/src/branch/master/gcu_mnist_example/train.py' },
          dataset: { required: true },
          runParameters: { required: false },
          networkType: { required: true },
          visualization: { required: false },
          spec: { required: true },
          workServerNum: { required: true },
          repo: { required: true },
        },
        hideCluster: true,
      }],
      'DCU': [{
        form: {
          taskName: { required: true, },
          taskDescr: { required: false, },
          branchName: { required: true, },
          model: { required: false, multiple: true },
          imagev1: { required: true, type: -1 },
          //imagev2: { required: true, relatedSpec: true },
          bootFile: { required: true, sampleUrl: 'https://openi.pcl.ac.cn/OpenIOSSG/OpenI_Cloudbrain_Example/src/branch/master/dcu_mnist_example/train.py' },
          dataset: { required: true },
          runParameters: { required: false },
          networkType: { required: true },
          visualization: { required: false },
          spec: { required: true },
          workServerNum: { required: true },
          repo: { required: true },
        },
        hideCluster: true,
      }],
      'ILUVATAR-GPGPU': [{
        form: {
          taskName: { required: true, },
          taskDescr: { required: false, },
          branchName: { required: true, },
          model: { required: false, multiple: true },
          imagev1: { required: true, type: -1 },
          bootFile: { required: true, sampleUrl: 'https://openi.pcl.ac.cn/OpenIOSSG/OpenI_Cloudbrain_Example/src/branch/master/gpgpu_mnist_example/train.py' },
          dataset: { required: true },
          runParameters: { required: false },
          networkType: { required: true },
          visualization: { required: false },
          spec: { required: true },
          workServerNum: { required: true },
          repo: { required: true },
        },
        hideCluster: true,
      }],
      'METAX-GPGPU': [{
        form: {
          taskName: { required: true, },
          taskDescr: { required: false, },
          branchName: { required: true, },
          model: { required: false, multiple: true },
          imagev1: { required: true, type: -1 },
          bootFile: { required: true, sampleUrl: 'https://openi.pcl.ac.cn/OpenIOSSG/OpenI_Cloudbrain_Example/src/branch/master/gpgpu_mnist_example/train.py' },
          dataset: { required: true },
          runParameters: { required: false },
          networkType: { required: true },
          visualization: { required: false },
          spec: { required: true },
          workServerNum: { required: true },
          repo: { required: true },
        },
        hideCluster: true,
      }],
      'BIREN-GPU': [{
        form: {
          taskName: { required: true, },
          taskDescr: { required: false, },
          branchName: { required: true, },
          model: { required: false, multiple: true },
          imagev1: { required: true, type: -1 },
          bootFile: { required: true, sampleUrl: '' },
          dataset: { required: true },
          runParameters: { required: false },
          networkType: { required: true },
          visualization: { required: false },
          spec: { required: true },
          workServerNum: { required: true },
          repo: { required: true },
        },
        hideCluster: true,
      }],
    }],
  }],
  //在线推理
  'ONLINEINFERENCE': [{
    'C2Net': [{
      'GPU': [{
        form: {
          taskName: { required: true, },
          taskDescr: { required: false, },
          branchName: {},
          bootFile: { required: true, sampleUrl: 'https://openi.pcl.ac.cn/OpenIOSSG/Online-Inference_Example' },
          model: { required: false, multiple: true },
          imagev1: { required: true, type: -1 },
          dataset: { required: false, useExceedSize: true },
          networkType: { required: true },
          spec: {},
          selfSshAddress: { required: false },
          repo: { required: true },
        },
        hideCluster: true,
      }],
      'NPU': [{
        form: {
          taskName: { required: true, },
          taskDescr: { required: false, },
          branchName: { required: true, },
          model: { required: false, multiple: true, useExceedSize: true },
          bootFile: { required: true, sampleUrl: 'https://openi.pcl.ac.cn/OpenIOSSG/Online-Inference_Example' },
          // imagev2: { required: true, relatedSpec: true },
          imagev1: { required: true, type: 2, useId: true },
          dataset: { required: false, useExceedSize: true },
          networkType: { required: true },
          spec: { required: true },
          // selfSshAddress: { required: false },
          repo: { required: true },
        },
        hideCluster: true,
      }],
      'GCU': [{
        form: {
          taskName: { required: true, },
          taskDescr: { required: false, },
          branchName: { required: true, },
          model: { required: false, multiple: true, useExceedSize: true },
          bootFile: { required: true, sampleUrl: 'https://openi.pcl.ac.cn/OpenIOSSG/Online-Inference_Example' },
          imagev1: { required: true, type: -1 },
          dataset: { required: false, useExceedSize: true },
          networkType: { required: true },
          spec: { required: true },
          selfSshAddress: { required: false },
          repo: { required: true },
        },
        hideCluster: true,
      }],
      'MLU': [{
        form: {
          taskName: { required: true, },
          taskDescr: { required: false, },
          branchName: {},
          bootFile: { required: true, sampleUrl: 'https://openi.pcl.ac.cn/OpenIOSSG/Online-Inference_Example' },
          model: { required: false, multiple: true },
          imagev1: { required: true, type: -1 },
          dataset: { required: false, useExceedSize: true },
          networkType: { required: true },
          spec: {},
          // selfSshAddress: { required: false },
          repo: { required: true },
        },
        hideCluster: true,
      }],
      'DCU': [{
        form: {
          taskName: { required: true, },
          taskDescr: { required: false, },
          branchName: {},
          bootFile: { required: true, sampleUrl: 'https://openi.pcl.ac.cn/OpenIOSSG/Online-Inference_Example' },
          model: { required: false, multiple: true },
          imagev1: { required: true, type: -1 },
          dataset: { required: false, useExceedSize: true },
          networkType: { required: true },
          spec: {},
          // selfSshAddress: { required: false },
          repo: { required: true },
        },
        hideCluster: true,
      }],
      'ILUVATAR-GPGPU': [{
        form: {
          taskName: { required: true, },
          taskDescr: { required: false, },
          branchName: {},
          bootFile: { required: true, sampleUrl: 'https://openi.pcl.ac.cn/OpenIOSSG/Online-Inference_Example' },
          model: { required: false, multiple: true },
          imagev1: { required: true, type: -1 },
          dataset: { required: false, useExceedSize: true },
          networkType: { required: true },
          spec: {},
          // selfSshAddress: { required: false },
          repo: { required: true },
        },
        hideCluster: true,
      }],
      'METAX-GPGPU': [{
        form: {
          taskName: { required: true, },
          taskDescr: { required: false, },
          branchName: {},
          bootFile: { required: true, sampleUrl: 'https://openi.pcl.ac.cn/OpenIOSSG/Online-Inference_Example' },
          model: { required: false, multiple: true },
          imagev1: { required: true, type: -1 },
          dataset: { required: false, useExceedSize: true },
          networkType: { required: true },
          spec: {},
          // selfSshAddress: { required: false },
          repo: { required: true },
        },
        hideCluster: true,
      }],
      'BIREN-GPU': [{
        form: {
          taskName: { required: true, },
          taskDescr: { required: false, },
          branchName: {},
          bootFile: { required: true, sampleUrl: 'https://openi.pcl.ac.cn/OpenIOSSG/Online-Inference_Example' },
          model: { required: false, multiple: true },
          imagev1: { required: true, type: -1 },
          dataset: { required: false, useExceedSize: true },
          networkType: { required: true },
          spec: {},
          // selfSshAddress: { required: false },
          repo: { required: true },
        },
        hideCluster: true,
      }],
    }]
  }],
  // 通用任务
  'GENERAL': [{
    'C2Net': [{
      'GPU': [{
        form: {
          taskName: { required: true, },
          taskDescr: { required: false, },
          branchName: {required: false},
          model: { required: false, multiple: true, useExceedSize: true },
          imagev1: { required: true, type: -1 },
          dataset: { required: false, useExceedSize: true },
          networkType: { required: true },
          spec: {},
          repo: { required: false },
        },
        hideCluster: true,
      }],
    }]
  }],
  'HPC': [{
    'C2Net': [{
      'CPU': [{
        appName: 'MMLSpark',
        form: {
          taskName: { required: true, },
          taskDescr: { required: false, },
          branchName: { required: false, },
          model: { required: false, multiple: true },
          imagev1: { required: true, type: -1, useId: true },
          // imagev2: { required: true },
          dataset: { required: false },
          spec: { required: true },
          repo: { required: false },
        },
        hideCluster: true,
      }]
    }]
  }],
};
// 基础配置类
class CreatePageConfigManager {
  constructor(config) {
    this.config = config;
    this.cache = new Map(); // 缓存查询结果
  }

  /**
   * 获取任务类型配置
   * @param {string} taskType - 任务类型: DEBUG, TRAIN, ONLINEINFERENCE, GENERAL, HPC
   * @returns {Array|null} 任务类型配置
   */
  getTaskTypeConfig(taskType) {
    const cacheKey = `taskType_${taskType}`;
    if (this.cache.has(cacheKey)) {
      return this.cache.get(cacheKey);
    }

    const config = this.config[taskType];
    this.cache.set(cacheKey, config);
    return config;
  }

  /**
   * 获取集群类型配置
   * @param {string} taskType - 任务类型
   * @param {string} clusterType - 集群类型 (目前主要是 C2Net)
   * @returns {Array|null} 集群类型配置
   */
  getClusterConfig(taskType, clusterType = 'C2Net') {
    const cacheKey = `network_${taskType}_${clusterType}`;
    if (this.cache.has(cacheKey)) {
      console.log("cache cacheKey")
      return this.cache.get(cacheKey);
    }

    const taskConfig = this.getTaskTypeConfig(taskType);
    if (!taskConfig) return null;

    const clusterConfig = taskConfig.find(item => item[clusterType]);
    const result = clusterConfig ? clusterConfig[clusterType] : null;
    this.cache.set(cacheKey, result);
    return result;
  }
  /**
   * 获取具体资源配置
   * @param {string} taskType - 任务类型
   * @param {string} resourceType - 资源类型: GPU, NPU, GCU, etc.
   * @param {string} clusterType - 网络类型
   * @returns {Object|null} 资源配置
   */
  getResourceConfig(taskType, resourceType, clusterType = 'C2Net') {
    const cacheKey = `resource_${taskType}_${clusterType}_${resourceType}`;
    if (this.cache.has(cacheKey)) {
      return this.cache.get(cacheKey);
    }

    const clusterConfig = this.getClusterConfig(taskType, clusterType);
    if (!clusterConfig || !clusterConfig[0]) return null;

    const resourceConfig = clusterConfig[0][resourceType];
    const result = resourceConfig && resourceConfig[0] ? resourceConfig[0] : null;
    this.cache.set(cacheKey, result);
    return result;
  }
  /**
   * 获取计算资源列表
   * @param {string} taskType - 任务类型
   * @param {string} networkType - 网络类型
   * @returns {Array|null} 可用的计算资源列表
   */
  getComputerResources(taskType, clusterType = 'C2Net') {
    const cacheKey = `resources_${taskType}_${clusterType}`;
    if (this.cache.has(cacheKey)) {
      return this.cache.get(cacheKey);
    }

    const clusterConfig = this.getClusterConfig(taskType, clusterType);
    if (!clusterConfig || !clusterConfig[0]) return null;

    const result = Object.keys(clusterConfig[0]) || null;
    this.cache.set(cacheKey, result);
    return result;
  }

  

  /**
   * 获取所有支持的任务类型
   * @returns {string[]} 任务类型列表
   */
  getAllTaskTypes() {
    return Object.keys(this.config);
  }
  /**
   * 获取任务支持的集群类型
   * @param {string} taskType - 任务类型
   * @returns {string[]} 任务类型列表
   */
  getTaskTypeAllClusters(taskType) {
    const taskConfig = this.getTaskTypeConfig(taskType);
    console.log(taskConfig)
    if (!taskConfig) return null;
    return Object.keys(taskConfig[0]);
  }
  /**
   * 清空缓存
   */
  clearCache() {
    this.cache.clear();
  }
}
// 创建配置管理器实例
export const configCreateManager = new CreatePageConfigManager(CreatePageConfigs);


export const FieldTemplates = {
  // 基础字段组
  basicFields: ['taskName', 'creator', 'descr'],
  resourceFields: ['aiCenter', 'computerRes', 'spec'],
  paramsFields1: ['imagev1', 'repo', 'branch'],
  paramsFields2: ['datasetList','modelList'],
  statusFields: ['status', 'createTime', 'startTime', 'endTime', 'duration'],
  // 特殊字段组
};
export const DetailPageConfigs = {
  // 调试任务
  'DEBUG': [{
    listUrl: 'debugjob?debugListType=all',
    clusters: ['OpenI', 'C2Net'],
    'OpenI': [{
      'GPU': [{
        detailUrl: 'cloudbrain/',
        summary: [],
        operations: ['debug', 'redebug', 'stop', 'saveTaskTmpl'],
        tabs: [{
          name: 'configInfo',
          fields: {
            basicFiled: [...FieldTemplates.basicFields],
            resourceFiled: [...FieldTemplates.resourceFields,'hasInternet','timeLimit'],
            paramsFiled: [...FieldTemplates.paramsFields1,...FieldTemplates.paramsFields2],
            runstatusFiled: [...FieldTemplates.statusFields],
            showSdkCode: [true],
          }
        }, {
          name: 'operationProfile'
        },{
          name: 'resultDownload'
        }],
      }],
      'NPU': [{
        detailUrl: 'modelarts/notebook/',
        summary: [],
        operations: ['debug', 'redebug', 'stop', 'saveTaskTmpl'],
        tabs: [{
          name: 'configInfo',
          fields: {
            basicFiled: [...FieldTemplates.basicFields],
            resourceFiled: [...FieldTemplates.resourceFields,'hasInternet','timeLimit'],
            paramsFiled: [...FieldTemplates.paramsFields1,...FieldTemplates.paramsFields2],
            runstatusFiled: [...FieldTemplates.statusFields],
            showSdkCode: [true],
          },
          showDatasetDownload: false,
          showModelFileDownload: false,
        },{
          name: 'resultDownload'
        }],
      }],
    }],
    'C2Net': [{
      'GPU': [{
        detailUrl: 'grampus/notebook/',
        summary: [],
        operations: ['debug', 'redebug', 'stop', 'saveTaskTmpl'],
        tabs: [{
          name: 'configInfo',
          fields: {
            basicFiled: [...FieldTemplates.basicFields],
            resourceFiled: [...FieldTemplates.resourceFields,'hasInternet','timeLimit'],
            paramsFiled: [...FieldTemplates.paramsFields1,...FieldTemplates.paramsFields2],
            runstatusFiled: [...FieldTemplates.statusFields],
            showSdkCode: [true],
          }
        },{
          name: 'resourceUseage',
          multiNodes: true,
        },{
          name: 'operationProfile'
        }, {
          name: 'resultDownload'
        }],
      }],
      'NPU': [{
        detailUrl: 'grampus/notebook/',
        summary: [],
        operations: ['debug', 'redebug', 'stop', 'saveTaskTmpl'],
        tabs: [{
          name: 'configInfo',
          fields: {
            basicFiled: [...FieldTemplates.basicFields],
            resourceFiled: [...FieldTemplates.resourceFields,'hasInternet','timeLimit'],
            paramsFiled: [...FieldTemplates.paramsFields1,...FieldTemplates.paramsFields2],
            runstatusFiled: [...FieldTemplates.statusFields],
            showSdkCode: [true],
          }
        }, {
          name: 'resourceUseage',
          multiNodes: true,
        }, {
          name: 'operationProfile'
        }, {
          name: 'resultDownload'
        }],
      }],
      'GCU': [{
        detailUrl: 'grampus/notebook/',
        summary: [],
        operations: ['debug', 'redebug', 'stop', 'saveTaskTmpl'],
        tabs: [{
          name: 'configInfo',
          fields: {
            basicFiled: [...FieldTemplates.basicFields],
            resourceFiled: [...FieldTemplates.resourceFields,'hasInternet','timeLimit'],
            paramsFiled: [...FieldTemplates.paramsFields1,...FieldTemplates.paramsFields2],
            runstatusFiled: [...FieldTemplates.statusFields],
            showSdkCode: [true],
          }
        }, {
          name: 'resourceUseage',
          multiNodes: true,
        }, {
          name: 'operationProfile'
        }, {
          name: 'resultDownload'
        }],
      }],
      'MLU': [{
        detailUrl: 'grampus/notebook/',
        summary: [],
        operations: ['debug', 'redebug', 'stop', 'saveTaskTmpl'],
        tabs: [{
          name: 'configInfo',
          fields: {
            basicFiled: [...FieldTemplates.basicFields],
            resourceFiled: [...FieldTemplates.resourceFields,'hasInternet','timeLimit'],
            paramsFiled: [...FieldTemplates.paramsFields1,...FieldTemplates.paramsFields2],
            runstatusFiled: [...FieldTemplates.statusFields],
            showSdkCode: [true],
          }
        }, {
          name: 'resourceUseage',
          multiNodes: true,
        }, {
          name: 'operationProfile'
        }, {
          name: 'resultDownload'
        }],
      }],
      'DCU': [{
        detailUrl: 'grampus/notebook/',
        summary: [],
        operations: ['debug', 'redebug', 'stop', 'saveTaskTmpl'],
        tabs: [{
          name: 'configInfo',
          fields: {
            basicFiled: [...FieldTemplates.basicFields],
            resourceFiled: [...FieldTemplates.resourceFields,'hasInternet','timeLimit'],
            paramsFiled: [...FieldTemplates.paramsFields1,...FieldTemplates.paramsFields2],
            runstatusFiled: [...FieldTemplates.statusFields],
            showSdkCode: [true],
          }
        }, {
          name: 'resourceUseage',
          multiNodes: true,
        }, {
          name: 'operationProfile'
        }, {
          name: 'resultDownload'
        }],
      }],
      'ILUVATAR-GPGPU': [{
        detailUrl: 'grampus/notebook/',
        summary: [],
        operations: ['debug', 'redebug', 'stop', 'saveTaskTmpl'],
        tabs: [{
          name: 'configInfo',
          fields: {
            basicFiled: [...FieldTemplates.basicFields],
            resourceFiled: [...FieldTemplates.resourceFields,'hasInternet','timeLimit'],
            paramsFiled: [...FieldTemplates.paramsFields1,...FieldTemplates.paramsFields2],
            runstatusFiled: [...FieldTemplates.statusFields],
            showSdkCode: [true],
          }
        }, {
          name: 'resourceUseage',
          multiNodes: true,
        }, {
          name: 'operationProfile'
        }, {
          name: 'resultDownload'
        }],
      }],
      'METAX-GPGPU': [{
        detailUrl: 'grampus/notebook/',
        summary: [],
        operations: ['debug', 'redebug', 'stop', 'saveTaskTmpl'],
        tabs: [{
          name: 'configInfo',
          fields: {
            basicFiled: [...FieldTemplates.basicFields],
            resourceFiled: [...FieldTemplates.resourceFields,'hasInternet','timeLimit'],
            paramsFiled: [...FieldTemplates.paramsFields1,...FieldTemplates.paramsFields2],
            runstatusFiled: [...FieldTemplates.statusFields],
            showSdkCode: [true],
          }
        }, {
          name: 'resourceUseage',
          multiNodes: true,
        }, {
          name: 'operationProfile'
        }, {
          name: 'resultDownload'
        }],
      }],
      'BIREN-GPU': [{
        detailUrl: 'grampus/notebook/',
        summary: [],
        operations: ['debug', 'redebug', 'stop', 'saveTaskTmpl'],
        tabs: [{
          name: 'configInfo',
          fields: {
            basicFiled: [...FieldTemplates.basicFields],
            resourceFiled: [...FieldTemplates.resourceFields,'hasInternet','timeLimit'],
            paramsFiled: [...FieldTemplates.paramsFields1,...FieldTemplates.paramsFields2],
            runstatusFiled: [...FieldTemplates.statusFields],
            showSdkCode: [true],
          }
        }, {
          name: 'resourceUseage',
          multiNodes: true,
        }, {
          name: 'operationProfile'
        }, {
          name: 'resultDownload'
        }],
      }],
    }]
  }],
  // 训练任务
  'TRAIN': [{
    listUrl: 'modelarts/train-job?listType=all',
    clusters: ['OpenI', 'C2Net'],
    'OpenI': [{
      'GPU': [{
        detailUrl: 'cloudbrain/train-job/',
        summary: [],
        operations: ['stop', 'tensorBoard', 'saveTaskTmpl', 'saveModel', 'exportDataset'],
        tabs: [{
          name: 'configInfo',
          fields: {
            basicFiled: [...FieldTemplates.basicFields],
            resourceFiled: [...FieldTemplates.resourceFields,'hasInternet','workServerNum','visualization'],
            paramsFiled: [...FieldTemplates.paramsFields1,'bootFile','runVersion','runParameters',...FieldTemplates.paramsFields2],
            runstatusFiled: [...FieldTemplates.statusFields],
            showSdkCode: [true],
          }
        }, {
          name: 'operationProfile'
        }, {
          name: 'logs'
        }, {
          name: 'resultDownload'
        }],
      }],
      'NPU': [{
        detailUrl: 'modelarts/train-job/',
        summary: [],
        operations: ['stop', 'tensorBoard', 'saveTaskTmpl', 'saveModel', 'exportDataset'],
        tabs: [{
          name: 'configInfo',
          fields: {
            basicFiled: [...FieldTemplates.basicFields],
            resourceFiled: [...FieldTemplates.resourceFields,'hasInternet','workServerNum','visualization'],
            paramsFiled: [...FieldTemplates.paramsFields1,'bootFile','runVersion','runParameters',...FieldTemplates.paramsFields2],
            runstatusFiled: [...FieldTemplates.statusFields],
            showSdkCode: [true],
          }
        }, {
          name: 'logs',
          multiNodes: true,
        }, {
          name: 'resourceUseage',
          multiNodes: true,
        }, {
          name: 'resultDownload'
        }],
      }],
    }],
    'C2Net': [{
      'GPU': [{
        detailUrl: 'grampus/train-job/',
        summary: [],
        operations: ['stop', 'tensorBoard', 'saveTaskTmpl', 'saveModel', 'exportDataset'],
        tabs: [{
          name: 'configInfo',
          fields: {
            basicFiled: [...FieldTemplates.basicFields],
            resourceFiled: [...FieldTemplates.resourceFields,'hasInternet','workServerNum','visualization'],
            paramsFiled: [...FieldTemplates.paramsFields1,'bootFile','runVersion','runParameters',...FieldTemplates.paramsFields2],
            runstatusFiled: [...FieldTemplates.statusFields],
            showSdkCode: [true],
          }
        }, {
          name: 'operationProfile'
        }, {
          name: 'logs',
          noScroll: true,
          multiNodes: true,
        }, {
          name: 'resourceUseage',
          multiNodes: true,
        }, {
          name: 'resultDownload'
        }],
      }],
      'NPU': [{
        detailUrl: 'grampus/train-job/',
        summary: [],
        operations: ['stop', 'tensorBoard', 'saveTaskTmpl', 'saveModel', 'exportDataset'],
        tabs: [{
          name: 'configInfo',
          fields: {
            basicFiled: [...FieldTemplates.basicFields],
            resourceFiled: [...FieldTemplates.resourceFields,'hasInternet','workServerNum','visualization'],
            paramsFiled: [...FieldTemplates.paramsFields1,'bootFile','runVersion','runParameters',...FieldTemplates.paramsFields2],
            runstatusFiled: [...FieldTemplates.statusFields],
            showSdkCode: [true],
          }
        }, {
          name: 'operationProfile'
        }, {
          name: 'logs',
          noScroll: true,
          multiNodes: true,
        }, {
          name: 'resourceUseage',
          multiNodes: true,
        }, {
          name: 'resultDownload'
        }],
      }],
      'GCU': [{
        detailUrl: 'grampus/train-job/',
        summary: [],
        operations: ['stop', 'tensorBoard', 'saveTaskTmpl', 'saveModel', 'exportDataset'],
        tabs: [{
          name: 'configInfo',
          fields: {
            basicFiled: [...FieldTemplates.basicFields],
            resourceFiled: [...FieldTemplates.resourceFields,'hasInternet','workServerNum','visualization'],
            paramsFiled: [...FieldTemplates.paramsFields1,'bootFile','runVersion','runParameters',...FieldTemplates.paramsFields2],
            runstatusFiled: [...FieldTemplates.statusFields],
            showSdkCode: [true],
          }
        }, {
          name: 'operationProfile'
        }, {
          name: 'logs',
          noScroll: true,
          multiNodes: true,
        }, {
          name: 'resourceUseage',
          multiNodes: true,
        }, {
          name: 'resultDownload'
        }],
      }],
      'DCU': [{
        detailUrl: 'grampus/train-job/',
        summary: [],
        operations: ['stop', 'tensorBoard', 'saveTaskTmpl', 'saveModel', 'exportDataset'],
        tabs: [{
          name: 'configInfo',
          fields: {
            basicFiled: [...FieldTemplates.basicFields],
            resourceFiled: [...FieldTemplates.resourceFields,'hasInternet','workServerNum','visualization'],
            paramsFiled: [...FieldTemplates.paramsFields1,'bootFile','runVersion','runParameters',...FieldTemplates.paramsFields2],
            runstatusFiled: [...FieldTemplates.statusFields],
            showSdkCode: [true],
          }
        }, {
          name: 'operationProfile'
        }, {
          name: 'logs',
          noScroll: true,
          multiNodes: true,
        }, {
          name: 'resourceUseage',
          multiNodes: true,
        }, {
          name: 'resultDownload'
        }],
      }],
      'MLU': [{
        detailUrl: 'grampus/train-job/',
        summary: [],
        operations: ['stop', 'tensorBoard', 'saveTaskTmpl', 'saveModel', 'exportDataset'],
        tabs: [{
          name: 'configInfo',
          fields: {
            basicFiled: [...FieldTemplates.basicFields],
            resourceFiled: [...FieldTemplates.resourceFields,'hasInternet','workServerNum','visualization'],
            paramsFiled: [...FieldTemplates.paramsFields1,'bootFile','runVersion','runParameters',...FieldTemplates.paramsFields2],
            runstatusFiled: [...FieldTemplates.statusFields],
            showSdkCode: [true],
          }
        }, {
          name: 'operationProfile'
        }, {
          name: 'logs',
          noScroll: true,
          multiNodes: true,
        }, {
          name: 'resourceUseage',
          multiNodes: true,
        }, {
          name: 'resultDownload'
        }],
      }],
      'ILUVATAR-GPGPU': [{
        detailUrl: 'grampus/train-job/',
        summary: [],
        operations: ['stop', 'tensorBoard', 'saveTaskTmpl', 'saveModel', 'exportDataset'],
        tabs: [{
          name: 'configInfo',
          fields: {
            basicFiled: [...FieldTemplates.basicFields],
            resourceFiled: [...FieldTemplates.resourceFields,'hasInternet','workServerNum','visualization'],
            paramsFiled: [...FieldTemplates.paramsFields1,'bootFile','runVersion','runParameters',...FieldTemplates.paramsFields2],
            runstatusFiled: [...FieldTemplates.statusFields],
            showSdkCode: [true],
          }
        }, {
          name: 'operationProfile'
        }, {
          name: 'logs',
          noScroll: true,
          multiNodes: true,
        }, {
          name: 'resourceUseage',
          multiNodes: true,
        }, {
          name: 'resultDownload'
        }],
      }],
      'METAX-GPGPU': [{
        detailUrl: 'grampus/train-job/',
        summary: [],
        operations: ['stop', 'tensorBoard', 'saveTaskTmpl', 'saveModel', 'exportDataset'],
        tabs: [{
          name: 'configInfo',
          fields: {
            basicFiled: [...FieldTemplates.basicFields],
            resourceFiled: [...FieldTemplates.resourceFields,'hasInternet','workServerNum','visualization'],
            paramsFiled: [...FieldTemplates.paramsFields1,'bootFile','runVersion','runParameters',...FieldTemplates.paramsFields2],
            runstatusFiled: [...FieldTemplates.statusFields],
            showSdkCode: [true],
          }
        }, {
          name: 'operationProfile'
        }, {
          name: 'logs',
          noScroll: true,
          multiNodes: true,
        }, {
          name: 'resourceUseage',
          multiNodes: true,
        }, {
          name: 'resultDownload'
        }],
      }],
    }]
  }],
  // 推理任务
  'INFERENCE': [{
    listUrl: 'modelarts/inference-job',
    clusters: ['OpenI', 'C2Net'],
    'OpenI': [{
      'NPU': [{
        detailUrl: 'modelarts/inference-job/',
        summary: [],
        operations: [],
        tabs: [{
          name: 'configInfo',
          fields: {
            basicFiled: [...FieldTemplates.basicFields],
            resourceFiled: [...FieldTemplates.resourceFields,'hasInternet'],
            paramsFiled: [...FieldTemplates.paramsFields1,'bootFile',...FieldTemplates.paramsFields2],
            runstatusFiled: [...FieldTemplates.statusFields],
          }
        }, {
          name: 'logs',
          multiNodes: true,
        }, {
          name: 'resourceUseage',
          multiNodes: true,
        }, {
          name: 'resultDownload'
        }],
      }],
    }],
    'C2Net': [{
      'GPU': [{
        detailUrl: 'grampus/inference-job/',
        summary: [],
        operations: [],
        tabs: [{
          name: 'configInfo',
          fields: {
            basicFiled: [...FieldTemplates.basicFields],
            resourceFiled: [...FieldTemplates.resourceFields,'hasInternet'],
            paramsFiled: [...FieldTemplates.paramsFields1,'bootFile',...FieldTemplates.paramsFields2],
            runstatusFiled: [...FieldTemplates.statusFields],
          }
        }, {
          name: 'operationProfile'
        }, {
          name: 'logs',
          noScroll: true,
          multiNodes: true,
        }, {
          name: 'resourceUseage',
          multiNodes: true,
        }, {
          name: 'resultDownload'
        }],
      }],
      'NPU': [{
        detailUrl: 'grampus/inference-job/',
        summary: [],
        operations: [],
        tabs: [{
          name: 'configInfo',
          fields: {
            basicFiled: [...FieldTemplates.basicFields],
            resourceFiled: [...FieldTemplates.resourceFields,'hasInternet'],
            paramsFiled: [...FieldTemplates.paramsFields1,'bootFile',...FieldTemplates.paramsFields2],
            runstatusFiled: [...FieldTemplates.statusFields],
          }
        }, {
          name: 'operationProfile'
        }, {
          name: 'logs',
          noScroll: true,
          multiNodes: true,
        }, {
          name: 'resourceUseage',
          multiNodes: true,
        }, {
          name: 'resultDownload'
        }],
      }],
      'ILUVATAR-GPGPU': [{
        detailUrl: 'grampus/inference-job/',
        summary: [],
        operations: [],
        tabs: [{
          name: 'configInfo',
          fields: {
            basicFiled: [...FieldTemplates.basicFields],
            resourceFiled: [...FieldTemplates.resourceFields,'hasInternet'],
            paramsFiled: [...FieldTemplates.paramsFields1,'bootFile',...FieldTemplates.paramsFields2],
            runstatusFiled: [...FieldTemplates.statusFields],
          }
        }, {
          name: 'operationProfile'
        }, {
          name: 'logs',
          noScroll: true,
          multiNodes: true,
        }, {
          name: 'resourceUseage',
          multiNodes: true,
        }, {
          name: 'resultDownload'
        }],
      }],
    }]
  }],
  'ONLINEINFERENCE': [{
    listUrl: 'grampus/onlineinfer',
    clusters: ['C2Net'],
    'C2Net': [{
      'GPU': [{
        detailUrl: 'grampus/onlineinfer',
        summary: [],
        operations: ['consoleHome', 'onlineInfer', 'reinfer', 'stop'],
        tabs: [{
          name: 'configInfo',
          fields: {
            basicFiled: [...FieldTemplates.basicFields],
            resourceFiled: [...FieldTemplates.resourceFields,'hasInternet'],
            paramsFiled: [...FieldTemplates.paramsFields1,'bootFile',...FieldTemplates.paramsFields2],
            runstatusFiled: [...FieldTemplates.statusFields],
          }
        }, {
          name: 'operationProfile'
        }, {
          name: 'logs',
          noScroll: true,
        }, {
          name: 'resultDownload'
        }],
      }],
      'GCU': [{
        detailUrl: 'grampus/onlineinfer',
        summary: [],
        operations: ['consoleHome', 'onlineInfer', 'reinfer', 'stop'],
        tabs: [{
          name: 'configInfo',
          fields: {
            basicFiled: [...FieldTemplates.basicFields],
            resourceFiled: [...FieldTemplates.resourceFields,'hasInternet'],
            paramsFiled: [...FieldTemplates.paramsFields1,'bootFile',...FieldTemplates.paramsFields2],
            runstatusFiled: [...FieldTemplates.statusFields],
          }
        }, {
          name: 'operationProfile'
        }, {
          name: 'logs',
          noScroll: true,
        }, {
          name: 'resultDownload'
        }],
      }],
      'NPU': [{
        detailUrl: 'grampus/onlineinfer',
        summary: [],
        operations: ['consoleHome', 'onlineInfer', 'reinfer', 'stop'],
        tabs: [{
          name: 'configInfo',
          fields: {
            basicFiled: [...FieldTemplates.basicFields],
            resourceFiled: [...FieldTemplates.resourceFields,'hasInternet'],
            paramsFiled: [...FieldTemplates.paramsFields1,'bootFile',...FieldTemplates.paramsFields2],
            runstatusFiled: [...FieldTemplates.statusFields],
          }
        }, {
          name: 'operationProfile'
        }, {
          name: 'logs',
          noScroll: true,
        }, {
          name: 'resultDownload'
        }],
      }],
      'MLU': [{
        detailUrl: 'grampus/onlineinfer',
        summary: [],
        operations: ['consoleHome', 'onlineInfer', 'reinfer', 'stop'],
        tabs: [{
          name: 'configInfo',
          fields: {
            basicFiled: [...FieldTemplates.basicFields],
            resourceFiled: [...FieldTemplates.resourceFields,'hasInternet'],
            paramsFiled: [...FieldTemplates.paramsFields1,'bootFile',...FieldTemplates.paramsFields2],
            runstatusFiled: [...FieldTemplates.statusFields],
          }
        }, {
          name: 'operationProfile'
        }, {
          name: 'logs',
          noScroll: true,
        }, {
          name: 'resultDownload'
        }],
      }],
      'DCU': [{
        detailUrl: 'grampus/onlineinfer',
        summary: [],
        operations: ['consoleHome', 'onlineInfer', 'reinfer', 'stop'],
        tabs: [{
          name: 'configInfo',
          fields: {
            basicFiled: [...FieldTemplates.basicFields],
            resourceFiled: [...FieldTemplates.resourceFields,'hasInternet'],
            paramsFiled: [...FieldTemplates.paramsFields1,'bootFile',...FieldTemplates.paramsFields2],
            runstatusFiled: [...FieldTemplates.statusFields],
          }
        }, {
          name: 'operationProfile'
        }, {
          name: 'logs',
          noScroll: true,
        }, {
          name: 'resultDownload'
        }],
      }],
      'ILUVATAR-GPGPU': [{
        detailUrl: 'grampus/onlineinfer',
        summary: [],
        operations: ['consoleHome', 'onlineInfer', 'reinfer', 'stop'],
        tabs: [{
          name: 'configInfo',
          fields: {
            basicFiled: [...FieldTemplates.basicFields],
            resourceFiled: [...FieldTemplates.resourceFields,'hasInternet'],
            paramsFiled: [...FieldTemplates.paramsFields1,'bootFile',...FieldTemplates.paramsFields2],
            runstatusFiled: [...FieldTemplates.statusFields],
          }
        }, {
          name: 'operationProfile'
        }, {
          name: 'logs',
          noScroll: true,
        }, {
          name: 'resultDownload'
        }],
      }],
      'METAX-GPGPU': [{
        detailUrl: 'grampus/onlineinfer',
        summary: [],
        operations: ['consoleHome', 'onlineInfer', 'reinfer', 'stop'],
        tabs: [{
          name: 'configInfo',
          fields: {
            basicFiled: [...FieldTemplates.basicFields],
            resourceFiled: [...FieldTemplates.resourceFields,'hasInternet'],
            paramsFiled: [...FieldTemplates.paramsFields1,'bootFile',...FieldTemplates.paramsFields2],
            runstatusFiled: [...FieldTemplates.statusFields],
          }
        }, {
          name: 'operationProfile'
        }, {
          name: 'logs',
          noScroll: true,
        }, {
          name: 'resultDownload'
        }],
      }],
      'BIREN-GPU': [{
        detailUrl: 'grampus/onlineinfer',
        summary: [],
        operations: ['consoleHome', 'onlineInfer', 'reinfer', 'stop'],
        tabs: [{
          name: 'configInfo',
          fields: {
            basicFiled: [...FieldTemplates.basicFields],
            resourceFiled: [...FieldTemplates.resourceFields,'hasInternet'],
            paramsFiled: [...FieldTemplates.paramsFields1,'bootFile',...FieldTemplates.paramsFields2],
            runstatusFiled: [...FieldTemplates.statusFields],
          }
        }, {
          name: 'operationProfile'
        }, {
          name: 'logs',
          noScroll: true,
        }, {
          name: 'resultDownload'
        }],
      }],
    }]
  }],
  'MODELEXPERIENCE': [{
    listUrl: '',
    clusters: ['C2Net'],
    'C2Net': [{
      'GPU': [{
        detailUrl: 'grampus/onlineinfer',
        summary: [],
        operations: ['onlinexperience', 'stop'],
        tabs: [{
          name: 'configInfo',
          fields: {
            basicFiled: [...FieldTemplates.basicFields],
            resourceFiled: [...FieldTemplates.resourceFields,'hasInternet'],
            paramsFiled: [...FieldTemplates.paramsFields1,'bootFile',...FieldTemplates.paramsFields2,'sourceFtName'],
            runstatusFiled: [...FieldTemplates.statusFields],
          }
        }, {
          name: 'operationProfile'
        }],
      }],
      'ILUVATAR-GPGPU': [{
        detailUrl: 'grampus/onlineinfer',
        summary: [],
        operations: ['onlinexperience', 'stop'],
        tabs: [{
          name: 'configInfo',
          fields: {
            basicFiled: [...FieldTemplates.basicFields],
            resourceFiled: [...FieldTemplates.resourceFields,'hasInternet'],
            paramsFiled: [...FieldTemplates.paramsFields1,'bootFile',...FieldTemplates.paramsFields2,'sourceFtName'],
            runstatusFiled: [...FieldTemplates.statusFields],
          }
        }, {
          name: 'operationProfile'
        }],
      }],
      'GCU': [{
        detailUrl: 'grampus/onlineinfer',
        summary: [],
        operations: ['onlinexperience', 'stop'],
        tabs: [{
          name: 'configInfo',
          fields: {
            basicFiled: [...FieldTemplates.basicFields],
            resourceFiled: [...FieldTemplates.resourceFields,'hasInternet'],
            paramsFiled: [...FieldTemplates.paramsFields1,'bootFile',...FieldTemplates.paramsFields2,'sourceFtName'],
            runstatusFiled: [...FieldTemplates.statusFields],
          }
        }, {
          name: 'operationProfile'
        }],
      }],
      'NPU': [{
        detailUrl: 'grampus/onlineinfer',
        summary: [],
        operations: ['onlinexperience', 'stop'],
        tabs: [{
          name: 'configInfo',
          fields: {
            basicFiled: [...FieldTemplates.basicFields],
            resourceFiled: [...FieldTemplates.resourceFields,'hasInternet'],
            paramsFiled: [...FieldTemplates.paramsFields1,'bootFile',...FieldTemplates.paramsFields2,'sourceFtName'],
            runstatusFiled: [...FieldTemplates.statusFields],
          }
        }],
      }],
    }],
  }],
  'FINETUNE': [{
    listUrl: '',
    clusters: ['C2Net'],
    'C2Net': [{
      'GPU': [{
        detailUrl: 'grampus/sftfinetune',
        summary: [],
        operations: ['stop', 'saveModel', "deployModel"],
        tabs: [{
          name: 'configInfo',
          fields: {
            basicFiled: [...FieldTemplates.basicFields],
            resourceFiled: [...FieldTemplates.resourceFields,'hasInternet'],
            paramsFiled: [...FieldTemplates.paramsFields1,'bootFile',...FieldTemplates.paramsFields2],
            runstatusFiled: [...FieldTemplates.statusFields],
          }
        }, {
          name: 'operationProfile'
        }, {
          name: 'logs',
          noScroll: true,
        }, {
          name: 'resourceUseage',
          multiNodes: true,
        }, {
          name: 'resultDownload'
        }, {
          name: 'loss',
        },],
      }],
      'NPU': [{
        detailUrl: 'grampus/sftfinetune',
        summary: [],
        operations: ['stop', 'saveModel', 'deployModel'],
        tabs: [{
          name: 'configInfo',
          fields: {
            basicFiled: [...FieldTemplates.basicFields],
            resourceFiled: [...FieldTemplates.resourceFields,'hasInternet'],
            paramsFiled: [...FieldTemplates.paramsFields1,'bootFile',...FieldTemplates.paramsFields2],
            runstatusFiled: [...FieldTemplates.statusFields],
          }
        }, {
          name: 'operationProfile'
        }, {
          name: 'logs',
          noScroll: true,
          multiNodes: true,
        }, {
          name: 'resourceUseage',
          multiNodes: true,
        }, {
          name: 'resultDownload'
        }, {
          name: 'loss',
        }],
      }],
      'ILUVATAR-GPGPU': [{
        detailUrl: 'grampus/sftfinetune',
        summary: [],
        operations: ['stop', 'saveModel'],
        tabs: [{
          name: 'configInfo',
          fields: {
            basicFiled: [...FieldTemplates.basicFields],
            resourceFiled: [...FieldTemplates.resourceFields,'hasInternet'],
            paramsFiled: [...FieldTemplates.paramsFields1,'bootFile',...FieldTemplates.paramsFields2],
            runstatusFiled: [...FieldTemplates.statusFields],
          }
        }, {
          name: 'operationProfile'
        }, {
          name: 'logs',
          noScroll: true,
          multiNodes: true,
        }, {
          name: 'resourceUseage',
          multiNodes: true,
        }, {
          name: 'resultDownload'
        }, {
          name: 'loss',
        }],
      }],
    }],
   
  }],
  'SDFINETUNE': [{
    listUrl: '',
    clusters: ['C2Net'],
    'C2Net': [{
      'GPU': [{
        detailUrl: 'grampus/sftfinetune',
        summary: [],
        operations: ['onlineLoraTrain', 'stop'],
        tabs: [{
          name: 'configInfo',
          fields: {
            basicFiled: [...FieldTemplates.basicFields],
            resourceFiled: [...FieldTemplates.resourceFields,'hasInternet'],
            paramsFiled: [...FieldTemplates.paramsFields1,'bootFile',...FieldTemplates.paramsFields2],
            runstatusFiled: [...FieldTemplates.statusFields],
          }
        }, {
          name: 'operationProfile'
        }, {
          name: 'logs',
          noScroll: true,
        }, {
          name: 'resultDownload'
        }],
      }],
      'ILUVATAR-GPGPU': [{
        detailUrl: 'grampus/sftfinetune',
        summary: [],
        operations: ['onlinexperience', 'stop'],
        tabs: [{
          name: 'configInfo',
          fields: {
            basicFiled: [...FieldTemplates.basicFields],
            resourceFiled: [...FieldTemplates.resourceFields,'hasInternet'],
            paramsFiled: [...FieldTemplates.paramsFields1,'bootFile',...FieldTemplates.paramsFields2],
            runstatusFiled: [...FieldTemplates.statusFields],
          }
        }, {
          name: 'operationProfile'
        }, {
          name: 'logs',
          noScroll: true,
          multiNodes: true,
        }],
      }],
      'GCU': [{
        detailUrl: 'grampus/sftfinetune',
        summary: [],
        operations: ['onlinexperience', 'stop'],
        tabs: [{
          name: 'configInfo',
          fields: {
            basicFiled: [...FieldTemplates.basicFields],
            resourceFiled: [...FieldTemplates.resourceFields,'hasInternet'],
            paramsFiled: [...FieldTemplates.paramsFields1,'bootFile',...FieldTemplates.paramsFields2],
            runstatusFiled: [...FieldTemplates.statusFields],
          }
        }, {
          name: 'operationProfile'
        }, {
          name: 'logs',
          noScroll: true,
          multiNodes: true,
        },],
      }],
      
    }],
    
  }],
  'ComfyuiExperience': [{
    listUrl: '',
    clusters: ['C2Net'],
    'C2Net': [{
      'GPU': [{
        detailUrl: 'grampus/sftfinetune',
        summary: [],
        operations: ['onlineWorkflow', 'stop'],
        tabs: [{
          name: 'configInfo',
          fields: {
            basicFiled: [...FieldTemplates.basicFields],
            resourceFiled: [...FieldTemplates.resourceFields,'hasInternet'],
            paramsFiled: [...FieldTemplates.paramsFields1,'bootFile',...FieldTemplates.paramsFields2],
            runstatusFiled: [...FieldTemplates.statusFields],
          }
        }, {
          name: 'operationProfile'
        }],
      }],
      'ILUVATAR-GPGPU': [{
        detailUrl: 'grampus/sftfinetune',
        summary: [],
        operations: ['onlinexperience', 'stop'],
        tabs: [{
          name: 'configInfo',
          fields: {
            basicFiled: [...FieldTemplates.basicFields],
            resourceFiled: [...FieldTemplates.resourceFields,'hasInternet'],
            paramsFiled: [...FieldTemplates.paramsFields1,'bootFile',...FieldTemplates.paramsFields2],
            runstatusFiled: [...FieldTemplates.statusFields],
          }
        }, {
          name: 'operationProfile'
        }],
      }],
      'GCU': [{
        detailUrl: 'grampus/sftfinetune',
        summary: [],
        operations: ['onlinexperience', 'stop'],
        tabs: [{
          name: 'configInfo',
          fields: {
            basicFiled: [...FieldTemplates.basicFields],
            resourceFiled: [...FieldTemplates.resourceFields,'hasInternet'],
            paramsFiled: [...FieldTemplates.paramsFields1,'bootFile',...FieldTemplates.paramsFields2],
            runstatusFiled: [...FieldTemplates.statusFields],
          }
        }, {
          name: 'operationProfile'
        }],
      }],
    }],
    
  }],
  'EVAL': [{
    listUrl: '',
    clusters: ['C2Net'],
    'C2Net': [{
      'GPU': [{
        detailUrl: 'grampus/evaluate',
        summary: [],
        operations: ['stop'],
        tabs: [{
          name: 'configInfo',
          fields: {
            basicFiled: [...FieldTemplates.basicFields],
            resourceFiled: [...FieldTemplates.resourceFields,'hasInternet'],
            paramsFiled: [...FieldTemplates.paramsFields1,'bootFile',...FieldTemplates.paramsFields2,'sourceFtName'],
            runstatusFiled: [...FieldTemplates.statusFields],
          }
        }, {
          name: 'operationProfile'
        }, {
          name: 'logs',
          noScroll: true,
        }, {
          name: 'evalOverview'
        }, {
          name: 'evalDetail'
        }],
      }],
      'ILUVATAR-GPGPU': [{
        detailUrl: 'grampus/evaluate',
        summary: [],
        operations: ['stop'],
        tabs: [{
          name: 'configInfo',
          fields: {
            basicFiled: [...FieldTemplates.basicFields],
            resourceFiled: [...FieldTemplates.resourceFields,'hasInternet'],
            paramsFiled: [...FieldTemplates.paramsFields1,'bootFile',...FieldTemplates.paramsFields2,'sourceFtName'],
            runstatusFiled: [...FieldTemplates.statusFields],
          }
        }, {
          name: 'operationProfile'
        }, {
          name: 'logs',
          noScroll: true,
        }, {
          name: 'evalOverview'
        }, {
          name: 'evalDetail'
        }],
      }],
      'GCU': [{
        detailUrl: 'grampus/evaluate',
        summary: [],
        operations: ['stop'],
        tabs: [{
          name: 'configInfo',
          fields: {
            basicFiled: [...FieldTemplates.basicFields],
            resourceFiled: [...FieldTemplates.resourceFields,'hasInternet'],
            paramsFiled: [...FieldTemplates.paramsFields1,'bootFile',...FieldTemplates.paramsFields2,'sourceFtName'],
            runstatusFiled: [...FieldTemplates.statusFields],
          }
        }, {
          name: 'logs',
          noScroll: true,
        }, {
          name: 'evalOverview'
        }, {
          name: 'evalDetail'
        }],
      }],
      'NPU': [{
        detailUrl: 'grampus/evaluate',
        summary: [],
        operations: ['stop'],
        tabs: [{
          name: 'configInfo',
          fields: {
            basicFiled: [...FieldTemplates.basicFields],
            resourceFiled: [...FieldTemplates.resourceFields,'hasInternet'],
            paramsFiled: [...FieldTemplates.paramsFields1,'bootFile',...FieldTemplates.paramsFields2,'sourceFtName'],
            runstatusFiled: [...FieldTemplates.statusFields],
          }
        }, {
          name: 'operationProfile'
        }, {
          name: 'logs',
          noScroll: true,
          multiNodes: true,
        }, {
          name: 'evalOverview'
        }, {
          name: 'evalDetail',
        }],
      }],
    }],
    
  }],
  // 通用任务
  'GENERAL': [{
    listUrl: 'grampus/general',
    clusters: ['C2Net'],
    'C2Net': [{
      'GPU': [{
        detailUrl: 'grampus/general/',
        summary: [],
        operations: ['debug', 'stop'],
        tabs: [{
          name: 'configInfo',
          fields: {
            basicFiled: [...FieldTemplates.basicFields],
            resourceFiled: [...FieldTemplates.resourceFields,'hasInternet'],
            paramsFiled: [...FieldTemplates.paramsFields1,...FieldTemplates.paramsFields2],
            runstatusFiled: [...FieldTemplates.statusFields],
            generalTaskCodeTips: [true],
          }
        }, {
          name: 'operationProfile'
        }],
      }],
    }]
  }],
  'HPC': [{
    listUrl: 'supercompute/job',
    clusters: ['C2Net'],
    'C2Net': [{
      'CPU': [{
        detailUrl: 'supercompute/job',
        summary: [],
        operations: ['stop'],
        tabs: [{
          name: 'configInfo',
          fields: {
            basicFiled: ['appName',...FieldTemplates.basicFields],
            resourceFiled: [...FieldTemplates.resourceFields],
            paramsFiled: [...FieldTemplates.paramsFields1,...FieldTemplates.paramsFields2],
            runstatusFiled: [...FieldTemplates.statusFields],
          }
        }, {
          name: 'operationProfile'
        }],
      }],
    }]
  }],
};
class DetailPageConfigManager {
  constructor(configs) {
    this.configMap = new Map();
    this.initialize(configs);
  }

  initialize(configs) {
    // 遍历所有配置，建立多层索引
    for (const [taskType, clustersConfig] of Object.entries(configs)) {
      const taskMap = new Map();
      
      for (const clusterConfig of clustersConfig) {
        for (const [clusterName, hardwareConfigs] of Object.entries(clusterConfig)) {
          if (clusterName === 'listUrl' || clusterName === 'clusters') continue;
          
          const clusterMap = new Map();
          
          for (const hardwareConfig of hardwareConfigs) {
            for (const [hardwareType, configArray] of Object.entries(hardwareConfig)) {
              // 每个硬件类型对应一个配置对象
              clusterMap.set(hardwareType, configArray[0]);
            }
          }
          
          taskMap.set(clusterName, clusterMap);
        }
      }
      
      this.configMap.set(taskType, {
        taskMap,
        listUrl: clustersConfig[0].listUrl,
        clusters: clustersConfig[0].clusters
      });
    }
  }

  // 快速获取配置
  getConfig(taskType, cluster, hardware) {
    const taskConfig = this.configMap.get(taskType);
    if (!taskConfig) return null;
    
    const clusterConfig = taskConfig.taskMap.get(cluster);
    if (!clusterConfig) return null;
    
    return clusterConfig.get(hardware);
  }

  // 获取列表URL
  getListUrl(taskType) {
    return this.configMap.get(taskType)?.listUrl;
  }

  // 获取支持的集群
  getClusters(taskType) {
    return this.configMap.get(taskType)?.clusters;
  }

  // 获取特定集群支持的硬件类型
  getHardwareTypes(taskType, cluster) {
    const taskConfig = this.configMap.get(taskType);
    if (!taskConfig) return [];
    
    const clusterConfig = taskConfig.taskMap.get(cluster);
    if (!clusterConfig) return [];
    
    return Array.from(clusterConfig.keys());
  }
}
export const configDetailManager = new DetailPageConfigManager(DetailPageConfigs);
// 快速访问
// const debugGPUConfig = configManager.getConfig('DEBUG', 'OpenI', 'GPU');

