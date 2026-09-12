<template>
  <div>
    <div v-if="emptyPage" style="padding-top:50px">
      <NotFound></NotFound>
    </div>
    <div v-else v-loading="loading">
      <Header :data="data" />
      <div class="content" :class="isZh ? 'zh' : ''">
        <div class="ui container" style="flex-wrap: wrap;">
          <div class="left">
            <div class="box">
              <div class="field-tit">{{ $t('cloudbrainObj.basicInfo') }}</div>
              <div class="fields">
                <div class="item">
                  <div class="item-l">{{ $t('resourcesManagement.jobType') }}：</div>
                  <div class="item-r">{{ data.JobTypeStr }}</div>
                </div>
                <div class="item">
                  <div class="item-l">{{ $t('resourcesManagement.computeResource') }}：</div>
                  <div class="item-r">{{ data.ComputeSourceStr }}</div>
                </div>
                <div class="item" v-if="formCfg.spec">
                  <div class="item-l">{{ $t('cloudbrainObj.resourceSpec') }}：</div>
                  <div class="item-r" v-html="renderContent('spec')"></div>
                </div>
                <div class="item" v-if="formCfg.networkType">
                  <div class="item-l">{{ $t('cloudbrainObj.networkType') }}：</div>
                  <div class="item-r" v-html="renderContent('networkType')"></div>
                </div>
                <div class="item" v-if="formCfg.visualization">
                  <div class="item-l">{{ $t('cloudbrainObj.visualization') }}：</div>
                  <div class="item-r" v-html="renderContent('visualization')"></div>
                </div>
              </div>
              <div class="field-tit">{{ $t('cloudbrainObj.paramsSetting') }}</div>
              <div class="fields">
                <div class="item" v-if="formCfg.imagev1">
                  <div class="item-l">{{ $t('cloudbrainObj.image') }}：</div>
                  <div class="item-r" v-html="renderContent('imagev1')"></div>
                </div>
                <div class="item" v-if="formCfg.imagev2">
                  <div class="item-l">{{ $t('cloudbrainObj.image') }}：</div>
                  <div class="item-r" v-html="renderContent('imagev2')"></div>
                </div>
                <div class="item" v-if="formCfg.repo">
                  <div class="item-l">{{ $t('cloudbrainObj.repo') }}：</div>
                  <div class="item-r" v-html="renderContent('repo')"></div>
                </div>
                <div class="item" v-if="formCfg.branchName">
                  <div class="item-l">{{ $t('cloudbrainObj.codeBranch') }}：</div>
                  <div class="item-r" v-html="renderContent('branchName')"></div>
                </div>
                <div class="item" v-if="formCfg.bootFile">
                  <div class="item-l">{{ $t('modelManage.bootFile') }}：</div>
                  <div class="item-r" v-html="renderContent('bootFile')"></div>
                </div>
                <div class="item" v-if="formCfg.runParameters">
                  <div class="item-l">{{ $t('modelManage.runParameters') }}：</div>
                  <div class="item-r" v-html="renderContent('runParameters')"></div>
                </div>
                <div class="item" v-if="formCfg.dataset">
                  <div class="item-l">{{ $t('cloudbrainObj.dataset') }}：</div>
                  <div class="item-r" v-html="renderContent('dataset')"></div>
                </div>
                <div class="item" v-if="formCfg.model">
                  <div class="item-l">{{ $t('repos.model') }}：</div>
                  <div class="item-r" v-html="renderContent('model')"></div>
                </div>
              </div>
            </div>
          </div>
          <div class="right">
            <div class="row">
              <div class="row-l" style="padding-top:2px;">{{ $t('modelManage.creator') }}：</div>
              <div class="row-r"><a v-if="data.Owner" :href="`/${data.Owner.Name}`" class="avatar-c">
                  <img class="avatar" :src="data.Owner.RelAvatarLink">{{ data.Owner.Name }}
                </a></div>
            </div>
            <div class="row">
              <div class="row-l">{{ $t('taskTmplObj.tmplAccessRight') }}：</div>
              <div class="row-r">
                <div v-if="data.IsPrivate" class="value" style="height: 21px;">
                  <span class="private">
                    <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 32 32" width="16" height="16">
                      <defs></defs>
                      <g>
                        <path
                          d="M25.333 13.333h1.333c0.736 0 1.333 0.597 1.333 1.333v0 13.333c0 0.736-0.597 1.333-1.333 1.333v0h-21.333c-0.736 0-1.333-0.597-1.333-1.333v0-13.333c0-0.736 0.597-1.333 1.333-1.333v0h1.333v-1.333c0-5.155 4.179-9.333 9.333-9.333s9.333 4.179 9.333 9.333v0 1.333zM6.667 16v10.667h18.667v-10.667h-18.667zM14.667 18.667h2.667v5.333h-2.667v-5.333zM22.667 13.333v-1.333c0-3.682-2.985-6.667-6.667-6.667s-6.667 2.985-6.667 6.667v0 1.333h13.333z">
                        </path>
                      </g>
                    </svg>
                    <span>{{ $t('modelManage.modelAccessPrivate') }}</span>
                  </span>
                </div>
                <div v-else>{{ $t('modelManage.modelAccessPublic') }}</div>
              </div>
            </div>
            <div class="row">
              <div class="row-l">{{ $t('taskTmplObj.tmplDescr') }}：</div>
              <div class="row-r descr" v-html="data.DescriptionStr || '--'"></div>
            </div>
            <div class="row">
              <div class="row-l">{{ $t('cloudbrainObj.createTime') }}：</div>
              <div class="row-r" v-html="data.CreatedUnixStr || '--'"></div>
            </div>
            <div class="row">
              <div class="row-l">{{ $t('modelManage.updateTime') }}：</div>
              <div class="row-r" v-html="data.UpdatedUnixStr || '--'"></div>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>
<script>
import NotFound from '~/components/NotFound.vue';
import Header from './components/Header.vue';
import { getAiTaskTmpl } from '~/apis/modules/aitasktmpl';
import { configCreateManager } from '~/pages/cloudbrain/configs';
import { renderSpecObject, initClipboard, getListValueWithKey, escapeHTML } from '~/utils';
import { TmplTaskTypes, TmplComputerResouces } from '~/pages/aitasktmpl/tools';
import { i18n, lang } from '~/langs';
import SparkMD5 from "spark-md5";
import { formatDate } from 'element-ui/lib/utils/date-util';

export default {
  data() {
    return {
      isZh: lang == 'zh-CN',
      emptyPage: false,
      loading: false,
      formCfg: {},
      data: {},
      clipboardHandler: null,
    };
  },
  components: { NotFound, Header },
  methods: {
    renderContent(key) {
      let result = '--';
      const data = this.data;
      switch (key) {
        case 'networkType':
          result = '--'
          if (data.HasInternet == 1) {
            result = i18n.t('cloudbrainObj.noInternet');
          } else if (data.HasInternet == 2) {
            result = i18n.t('cloudbrainObj.hasInternet');
          }
          break;
        case 'visualization':
          result = !!data.VisualizeRequired ? i18n.t('cloudbrainObj.tensorBoardVisualization') : i18n.t('none');
          break;
        case 'spec':
          result = '--';
          if (data.AccCardType) {
            const specObj = renderSpecObject({
              id: '',
              compute_resource: data.ComputeSource,
              acc_cards_num: data.AccCardsNum,
              acc_card_type: data.AccCardType,
              gpu_mem_gi_b: data.GPUMemGiB,
              mem_gi_b: data.MemGiB,
              share_mem_gi_b: data.ShareMemGiB,
              cpu_cores: data.CpuCores,
            }, false);
            result = specObj.specStr;
          }
          break;
        case 'imagev1':
          if (data.ImageName || data.ImageUrl) {
            result = `<span class="ui poping up clipboard"
              id="clipboard-${SparkMD5.hash(data.ImageName || data.ImageUrl)}"
              data-position="top center"
              data-variation="inverted tiny"
              data-success="${this.$t('copyedLink')}"
              data-content="${this.$t('copy')}" 
              data-original="${this.$t('copy')}"
              data-clipboard-text="${data.ImageName || data.ImageUrl}"><span title="${data.ImageName || data.ImageUrl}">${data.ImageName || data.ImageUrl}</span></span>`;
          }
          break;
        case 'imagev2':
          if (data.ImageName) {
            result = `<span class="ui poping up clipboard"
              id="clipboard-${SparkMD5.hash(data.ImageName)}"
              data-position="top center"
              data-variation="inverted tiny"
              data-success="${this.$t('copySuccess')}"
              data-content="${this.$t('copy')}" 
              data-original="${this.$t('copy')}"
              data-clipboard-text="${data.ImageName}"><span title="${data.ImageName}">${data.ImageName}</span></span>`;
          }
          break;
        case 'repo':
          if (data.RepoOwnerName && data.RepoName) {
            result = `<a target="_blank" href="/${data.RepoOwnerName}/${data.RepoName}">${data.RepoOwnerName} / ${data.RepoName}</a>`;
          }
          break;
        case 'branchName':
          result = data.BranchName || '--';
          break;
        case 'bootFile':
          result = data.BootFile || '--';
          break;
        case 'runParameters':
          const parametersStr = data.Parameters || '[]';
          const paramList = JSON.parse(parametersStr);
          const paramsStr = paramList.map(item => (`${item.Label} = ${item.Value}`)).join('; ');
          result = paramsStr ? `<span class="ui poping up clipboard"
          id="clipboard-${SparkMD5.hash(paramsStr)}"
          data-position="top center"
          data-variation="inverted tiny"
          data-success="${this.$t('copySuccess')}"
          data-content="${this.$t('copy')}" 
          data-original="${this.$t('copy')}"
          data-clipboard-text="${paramsStr}"><span title="${paramsStr}">${paramsStr}</span></span>` : '--';
          break;
        case 'dataset':
          const datasetList = (data.DatasetLists || []);
          if (datasetList.length) {
            result = '<div class="list">';
            for (let i = 0, iLen = datasetList.length; i < iLen; i++) {
              const dataset = datasetList[i];
              const isDeleted = dataset.IsDeleted;
              const deletedStr = isDeleted ? `(${this.$t('modelManage.deleted')})` : '';
              const canJump = !isDeleted && dataset.OwnerName;
              result += (canJump ? `<div class="item"><a target="_blank" href="/datasets/detail/${dataset.OwnerName}/${dataset.DatasetName}">${dataset.OwnerName}/${dataset.DatasetAlias || dataset.DatasetName}</a></div>` :
                `<div class="item">${dataset.OwnerName ? (dataset.OwnerName + '/') : ''}${dataset.DatasetAlias || dataset.DatasetName}${deletedStr}</div>`);
            }
            result += '</div>';
          }
          break;
        case 'model':
          const modelList = (data.ModelLists || []);
          if (modelList.length) {
            result = '<div class="list">';
            for (let i = 0, iLen = modelList.length; i < iLen; i++) {
              const model = modelList[i];
              const isDeleted = model.IsDeleted;
              const deletedStr = isDeleted ? `(${this.$t('modelManage.deleted')})` : '';
              const canJump = !isDeleted && model.OwnerName;
              result += (canJump ? `<div class="item"><a target="_blank" href="/models/detail/${model.OwnerName}/${model.ModelName}">${model.OwnerName}/${model.ModelAlias || model.ModelName}</a></div>` :
                `<div class="item">${model.OwnerName ? (model.OwnerName + '/') : ''}${model.ModelAlias || model.ModelName}${deletedStr}</div>`);
            }
            result += '</div>';
          }
          break;
        default:
          result = '--';
      }
      return result;
    }
  },
  beforeMount() {
    const pathname = window.location.pathname;
    const id = pathname.split('/')[3];
    if (!id) {
      this.emptyPage = true;
      return;
    }
    this.loading = true;
    getAiTaskTmpl({ id }).then(res => {
      res = res.data;
      if (res.code == 0) {
        this.data = {
          ...res.data,
          JobTypeStr: getListValueWithKey(TmplTaskTypes, res.data.JobType),
          ComputeSourceStr: getListValueWithKey(TmplComputerResouces, res.data.ComputeSource),
          DescriptionStr: escapeHTML(res.data.Description),
          CreatedUnixStr: formatDate(new Date(res.data.CreatedUnix * 1000), 'yyyy-MM-dd HH:mm:ss'),
          UpdatedUnixStr: formatDate(new Date(res.data.UpdatedUnix * 1000), 'yyyy-MM-dd HH:mm:ss'),
        }
        const configs = configCreateManager.getResourceConfig(this.data.JobType, this.data.ComputeSource)
        if (!configs) {
          console.error('[tmplEdit] get configs error');
          this.emptyPage = true;
          return;
        }
        this.formCfg = configs.form || {};
        // console.log('formCfg', this.formCfg);
        this.$nextTick(() => {
          this.clipboardHandler && this.clipboardHandler.destroy();
          this.clipboardHandler = initClipboard();
        });
      } else {
        this.emptyPage = true;
      }
    }).catch(err => {
      console.log(err);
      this.emptyPage = true;
    }).finally(() => {
      this.loading = false;
    });
  },
  mounted() { },
  beforeDestroy() { },
};
</script>

<style>
.full.height {
  background-color: rgba(249, 249, 249, 1);
}
</style>
<style scoped lang="less">
.content {
  background-color: rgba(249, 249, 249, 1);

  .container {
    display: flex;

    .left {
      width: 0;
      flex: 1;
      padding-top: 20px;

      .title {
        color: rgba(16, 16, 16, 0.5);
        font-size: 14px;
      }

      .box {
        margin-top: 10px;
        border-radius: 10px;
        background-color: rgba(255, 255, 255, 1);
        border: 1px solid rgba(224, 227, 234, 1);
        min-height: 250px;
        padding: 50px 60px;

        .field-tit {
          color: rgb(118, 162, 211);
          font-size: 18px;
          font-weight: 700;
          width: 100%;
          border-bottom: 1px solid rgba(16, 16, 16, 0.2);
          height: 40px;
          display: flex;
          align-items: center;
          margin-bottom: 20px;
        }

        .fields {
          margin-bottom: 25px;

          .item {
            display: flex;
            font-size: 14px;

            .item-l {
              width: 180px;
              height: 32px;
              display: flex;
              justify-content: flex-end;
              color: rgba(16, 16, 16, 0.5);
            }

            .item-r {
              overflow: hidden;
              width: 0;
              flex: 1;
              white-space: nowrap;
              text-overflow: ellipsis;

              /deep/ .list {
                margin-bottom: 8px;

                .item {
                  display: block;
                  margin-bottom: 5px;
                }
              }
            }
          }
        }
      }
    }

    .right {
      width: 340px;
      margin-left: 20px;
      padding-top: 38px;

      .row {
        display: flex;

        .row-l {
          height: 36px;
          width: 148px;
          color: rgba(16, 16, 16, 0.5);
          display: flex;
          justify-content: flex-start;
        }

        .row-r {
          min-height: 36px;
          width: 0;
          flex: 1;
          line-height: 20px;

          &.descr {
            white-space: pre-wrap;
            padding-bottom: 10px;
          }
        }

        .avatar-c {
          display: flex;
          align-items: center;

          .avatar {
            display: inline-block;
            width: 24px;
            height: 24px;
            border-radius: 100%;
            margin-right: 5px;
          }
        }

        .private {
          display: inline-flex;
          align-items: center;
          background-color: rgb(250, 140, 22);
          color: rgba(255, 255, 255, 1);
          font-size: 12px;
          border-radius: 4px;
          padding: 0px 3px;
          gap: 3px;

          svg:not([stroke]) {
            fill: rgb(255, 255, 255);
          }
        }
      }
    }
  }

  &.zh {
    .item-l {
      width: 170px !important;
    }

    .right {
      .row-l {
        width: 80px !important;
      }
    }
  }
}
@media only screen and (max-width: 768px) {
  .left{
    width: 100% !important;
    .box{
      padding: 20px 10px 10px 10px !important;
      .item-l {
        width: 100px !important;
      }
    }
  }
  .right{
    width: 100% !important;
  }
}
</style>
