import Vue from 'vue';
import Router from 'vue-router';
import { getMlopsRight } from '~/apis/modules/common';

const originalPush = Router.prototype.push;
Router.prototype.push = function push(location) {
  return originalPush.call(this, location).catch(err => err);
}

Vue.use(Router);

import Layout from '../layout/Layout.vue';
import LayoutEmpty from '../layout/LayoutEmpty.vue';

const router = new Router({
  base: 'modelbase',
  mode: 'history',
  routes: [{
    path: '/',
    component: Layout,
    redirect: 'list',
    children: [{
      path: 'list',
      name: 'ModelList',
      meta: { menu: true, icon: 'ri-apps-line', label: 'modelSquare.largeModelsList', ignoreAuth: true, },
      component: () => import(/* webpackChunkName: "modelbase-modellist" */ '../views/list/ModelList.vue'),
    }, {
      path: 'mind',
      name: 'Mind',
      meta: { menu: true, icon: 'ri-bookmark-3-line', label: 'modelSquare.pcMind', ignoreAuth: true, },
      component: () => import(/* webpackChunkName: "modelbase-mind" */ '../views/mind/Mind.vue'),
    }, {
      path: 'experience',
      name: 'Experience',
      meta: {
        menu: true, icon: 'ri-bilibili-line', label: 'modelSquare.modelExperience',
        descr: 'modelSquare.modelExperienceTips',
      },
      component: () => import(/* webpackChunkName: "modelbase-experience" */ '../views/experience/Experience.vue'),
    }, {
      path: 'experience/create',
      name: 'ExperienceCreate',
      meta: { label: 'modelSquare.createChatBtns1', },
      component: () => import(/* webpackChunkName: "modelbase-experience-create" */ '../views/experience/ExpCreate.vue'),
    }, {
      path: 'experience/detail/:taskid',
      name: 'ExperienceDetail',
      meta: { label: 'modelSquare.ModelExperienceDetail', },
      component: () => import(/* webpackChunkName: "modelbase-experience-detail" */ '../views/experience/ExpDetail.vue'),
    }, {
      path: 'arena',
      name: 'Arena',
      meta: {
        menu: true, icon: 'arena', iconType: 'svg', label: 'modelSquare.modelArena',
        descr: '',
      },
      component: () => import(/* webpackChunkName: "modelbase-arena" */ '../views/arena/Arena.vue'),
    }, {
      path: 'eval',
      name: 'Evaluate',
      meta: { menu: true, label: 'modelSquare.modelEvaluate' },
      redirect: 'eval/evaluate',
      component: LayoutEmpty,
      children: [{
        path: 'evaluate',
        name: 'Lndex',
        meta: {
          menu: true, icon: 'evaluate' , iconType: 'svg',label: 'modelSquare.modelPerEvaluate',
          descr: 'modelSquare.modelEvaluateTips',
        },
        component: () => import(/* webpackChunkName: "modelbase-evaluate" */ '../views/evaluate/Evaluate.vue'),
      },
      {
        path: 'list1',
        name: 'Lndex1',
        outLink: 'http://aisafety.openi.org.cn:30990/atp-web/index',
        meta: {
          menu: true, icon: 'evallist', iconType: 'svg', label: 'modelSquare.modelSafeEvaluate',
          descr: 'modelSquare.modelEvaluateTips',
        },
        component: () => import(/* webpackChunkName: "modelbase-evaluate" */ '../views/evaluate/Evaluate.vue'),
      },{
          path: 'evaluate/create',
          name: 'EvaluateCreate',
          meta: { label: 'modelSquare.createEvalBtns1', },
          component: () => import(/* webpackChunkName: "modelbase-evaluate-create" */ '../views/evaluate/EvalCreate.vue'),
        }, {
          path: 'evaluate/detail/:taskid',
          name: 'EvaluateDetail',
          meta: { label: 'modelSquare.ModelEvaluateDetail', },
          component: () => import(/* webpackChunkName: "modelbase-evaluate-detail" */ '../views/evaluate/EvalDetail.vue'),
        },
      ]
    },
      {
      path: 'nlp',
      name: 'nlp',
      meta: { menu: true, label: 'datasets.natural_language_processing' },
      redirect: 'nlp/sft',
      component: LayoutEmpty,
      children: [{
        path: 'sft',
        name: 'NlpSFT',
        meta: {
          menu: true, icon: 'ri-equalizer-line', label: 'modelSquare.sftFinetune',
          descr: 'modelSquare.sftFinetuneTips',
          hideAuths: { mlops_right: true }
        },
        component: () => import(/* webpackChunkName: "nlp-sft" */ '../views/finetuning/SFT.vue'),
      }, {
        path: 'sft/create',
        name: 'NlpSFTCreate',
        meta: {
          label: 'modelSquare.newSftFinetune',
          hideAuths: { mlops_right: true }
        },
        component: () => import(/* webpackChunkName: "nlp-sft-create" */ '../views/finetuning/SFTCreate.vue'),
      }, {
        path: 'sft/detail/:taskid',
        name: 'NlpSFTDetail',
        meta: {
          label: 'modelSquare.sftFinetuneDetail',
          hideAuths: { mlops_right: true }
        },
        component: () => import(/* webpackChunkName: "nlp-sft-detail" */ '../views/finetuning/SFTDetail.vue'),
      }, {
        path: 'iflyaicloud',
        name: 'Iflyaicloud',
        meta: { menu: true, icon: 'ri-cloud-line', label: '大模型微调', auths: { mlops_right: true } },
        redirect: 'iflyaicloud/modelsquare',
        component: LayoutEmpty,
        children: [
          {
            path: 'modelsquare',
            name: 'IflyaicloudModelsquare',
            meta: {
              menu: true, label: '模型集市',
              auths: { mlops_right: true },
              iframeLink: 'https://experience.pro.iflyaicloud.com/maas-finetune/modelSquare',
            },
            component: () => import(/* webpackChunkName: "IflyaicloudModelsquare" */ '../views/iframe/Iframe.vue'),
          }, {
            path: 'text2text',
            name: 'IflyaicloudText2text',
            meta: {
              menu: true, label: '文本体验',
              auths: { mlops_right: true },
              iframeLink: 'https://experience.pro.iflyaicloud.com/maas-finetune/experience/text2text',
            },
            component: () => import(/* webpackChunkName: "IflyaicloudText2text" */ '../views/iframe/Iframe.vue'),
          }, {
            path: 'model',
            name: 'IflyaicloudModel',
            meta: {
              menu: true, label: '我的模型',
              auths: { mlops_right: true },
              iframeLink: 'https://experience.pro.iflyaicloud.com/maas-finetune/model',
            },
            component: () => import(/* webpackChunkName: "IflyaicloudModel" */ '../views/iframe/Iframe.vue'),
          }, {
            path: 'dataset',
            name: 'IflyaicloudDataset',
            meta: {
              menu: true, label: '数据集',
              auths: { mlops_right: true },
              iframeLink: 'https://experience.pro.iflyaicloud.com/maas-finetune/dataset/index',
            },
            component: () => import(/* webpackChunkName: "IflyaicloudDataset" */ '../views/iframe/Iframe.vue'),
          }, {
            path: 'modelservice',
            name: 'IflyaicloudModelService',
            meta: {
              menu: true, label: '推理服务',
              auths: { mlops_right: true },
              iframeLink: 'https://experience.pro.iflyaicloud.com/maas-finetune/modelService',
            },
            component: () => import(/* webpackChunkName: "IflyaicloudModelService" */ '../views/iframe/Iframe.vue'),
          }, {
            path: 'batchinference',
            name: 'IflyaicloudBatchInference',
            meta: {
              menu: true, label: '批量推理',
              auths: { mlops_right: true },
              iframeLink: 'https://experience.pro.iflyaicloud.com/maas-finetune/batchInference',
            },
            component: () => import(/* webpackChunkName: "IflyaicloudBatchInference" */ '../views/iframe/Iframe.vue'),
          }, {
            path: 'modelevaluate',
            name: 'IflyaicloudModelEvaluate',
            meta: {
              menu: true, label: '模型评估',
              auths: { mlops_right: true },
              iframeLink: 'https://experience.pro.iflyaicloud.com/maas-finetune/modelEvaluate',
            },
            component: () => import(/* webpackChunkName: "IflyaicloudModelEvaluate" */ '../views/iframe/Iframe.vue'),
          }
        ]
      }]
    }, {
      path: 'cv',
      name: 'cv',
      meta: { menu: true, label: 'datasets.computer_vision' },
      redirect: 'cv/sft',
      component: LayoutEmpty,
      children: [{
        path: 'sft',
        name: 'CvLoraTrain',
        meta: {
          menu: true, icon: 'ri-gallery-line', label: 'modelSquare.cvLoraTrain',
          descr: 'modelSquare.cvloraTips',
        },
        component: () => import(/* webpackChunkName: "cv-lora-list" */ '../views/finetuning/cvlora/Lora.vue'),
      }, {
        path: 'sft/create',
        name: 'CvLoraCreate',
        meta: { label: 'modelSquare.newLoraFinetune', },
        component: () => import(/* webpackChunkName: "cv-lora-create" */ '../views/finetuning/cvlora/LoraCreate.vue'),
      }, {
        path: 'sft/detail/:taskid',
        name: 'CvLoraDetail',
        meta: { label: 'modelSquare.sftFinetuneDetail', },
        component: () => import(/* webpackChunkName: "cv-lora-detail" */ '../views/finetuning/SFTDetail.vue'),
      }, {
        path: 'sft/lora',
        name: 'FinetuneSFTLora',
        meta: { label: '', },
        component: () => import(/* webpackChunkName: "cv-lora-train" */ '../views/finetuning/cvlora/TrainIndex.vue'),
      },
      {
        path: 'comfyui',
        name: 'CvComfyui',
        meta: {
          menu: true, icon: 'ri-git-pull-request-line', label: 'modelSquare.cvComfyui', descr: 'modelSquare.ComfyUiTips',
        },
        component: () => import(/* webpackChunkName: "cv-comfyui" */ '../views/comfyui/ComfyUi.vue'),
      },
      {
        path: 'comfyui/create',
        name: 'CvComfyuiCreate',
        meta: { label: 'modelSquare.newComfyUi', },
        component: () => import(/* webpackChunkName: "cv-comfyui-create" */ '../views/comfyui/ComfyuiCreate.vue'),
      }, {
        path: 'comfyui/detail/:taskid',
        name: 'CvComfyuiDetail',
        meta: { label: 'modelSquare.sftFinetuneDetail', },
        component: () => import(/* webpackChunkName: "cv-comfyui-detail" */ '../views/finetuning/SFTDetail.vue'),
      },
      ]
    }, {
      path: 'appdev',
      name: 'AppDev',
      meta: { menu: true, icon: 'ri-app-store-line', label: 'modelSquare.appDev', ignoreAuth: true, topSplitLine: true, },
      component: () => import(/* webpackChunkName: "modelbase-appdev" */ '../views/appdev/AppDev.vue'),
    }, {
      path: 'relatedtools',
      name: 'Relatedtools',
      meta: {
        menu: true, icon: 'ri-outlet-2-line', label: 'modelSquare.relatedTools', ignoreAuth: true, topSplitLine: true,
        hideAuths: { mlops_right: true }
      },
      component: () => import(/* webpackChunkName: "modelbase-relatedtools" */ '../views/relatedtools/RelatedTools.vue'),
    }],
  }, {
    path: '*',
    redirect: 'list',
  }]
});

router.auth = {
  isInit: false,
}

const initAuth = async () => {
  router.auth.isLogin = !!document.querySelector('meta[name="_uid"]');
  const res = await getMlopsRight()
  router.auth = {
    ...router.auth,
    ...res.data,
    isInit: true,
  }
}

router.beforeEach(async (to, from, next) => {
  if (!router.auth.isInit) {
    try {
      await initAuth()
    } catch (err) {
      router.auth = {
        ...router.auth,
        isInit: true,
      }
      console.log(err)
    }
  }
  const allAuth = router.auth;
  let check = true, hideCheck = false;
  if (to.meta.auths) {
    for (let auth in to.meta.auths) {
      if (!allAuth[auth]) {
        check = false;
      }
    }
  }
  if (to.meta.hideAuths) {
    for (let auth in to.meta.hideAuths) {
      if (allAuth[auth]) {
        hideCheck = true;
      }
    }
  }
  if (to.meta.ignoreAuth) {
    if (check && !hideCheck) {
      next();
    } else {
      next({ path: '/' });
    }
  } else {
    if (allAuth.isLogin) {
      if (check && !hideCheck) {
        
        next();
      } else {
        next({ path: '/' });
      }
    } else {
      window.location.href = `/user/login?redirect_to=${encodeURIComponent(window.location.origin + '/' + router.options.base + to.fullPath)}`;
    }
  }
});

export default router;