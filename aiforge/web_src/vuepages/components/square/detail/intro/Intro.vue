<template>
  <div class="ui container content">
    <div class="content-l" v-loading="loading">
      <div class="intro-content" v-if="!(introEmpty && !editing) && !loading">
        <div class="read-mode" v-if="!editing">
          <div class="head">
            <div>
              <i class="el-icon-warning-outline"></i>
              <span>{{ introFileName }}</span>
            </div>
            <div>
              <i v-if="canEdit" class="icon pencil alternate" :title="$t('modelManage.edit')"
                @click="toggleEdit(true)"></i>
            </div>
          </div>
          <div class="content-box">
            <div class="file-view markdown markdown" v-html="htmlContent"></div>
          </div>
        </div>
        <div class="edit-mode" v-if="canEdit && editing">
          <div class="head">
            <div>
              <i class="el-icon-warning-outline"></i>
              <span>{{ introFileName }}</span>
            </div>
            <div>
              <i v-if="canEdit" class="icon reply" :title="$t('cancel')" @click="toggleEdit(false)"></i>
            </div>
          </div>
          <div class="item-tab-c">
            <div class="item-tab" :class="editTab == 'edit' ? 'focus' : ''" tab="edit" @click="changeEditTab('edit')">
              <svg class="svg octicon-code" width="16" height="16" aria-hidden="true">
                <use xlink:href="#octicon-code"></use>
              </svg>
              <span>{{ $t('modelManage.editFiles') }}</span>
            </div>
            <div class="item-tab" :class="editTab == 'preview' ? 'focus' : ''" tab="preview"
              @click="changeEditTab('preview')">
              <svg class="svg octicon-eye" width="16" height="16" aria-hidden="true">
                <use xlink:href="#octicon-eye"></use>
              </svg>
              <span>{{ $t('modelManage.preview') }}</span>
            </div>
          </div>
          <div class="tab-content edit-content" v-show="editTab == 'edit'">
            <div ref="editContainerRef" class="monaco-editor-container" v-loading="editLoading"></div>
          </div>
          <div class="tab-content preview-content" v-show="editTab == 'preview'" v-loading="previewLoading">
            <div ref="previewContainerRef" class="preview-container markdown" v-html="previewContent">
            </div>
          </div>
          <div class="conmit-btn-c">
            <el-button size="default" class="btn confirm-btn" :disabled="this.content == this.editContent"
              v-loading="submitLoading" :class="this.content == this.editContent ? '_disabled_' : ''" @click="submit">
              {{ $t('submit') }} </el-button>
            <el-button size="default" class="btn" @click="toggleEdit(false)">{{ $t('cancel') }}</el-button>
          </div>
        </div>
      </div>
      <div class="empty" v-if="introEmpty && !editing && !loading">
        <div class="icon-c">
          <div class="empty-icon"></div>
        </div>
        <div class="tips">{{ type == 'dataset' ? $t('datasetObj.hasNoIntroForModel') :
          $t('modelManage.hasNoIntroForModel')}}</div>
        <div class="ops" v-if="canEdit">
          <el-button class="create-btn" size="default" @click="createIntro">
            {{ type == 'dataset' ? $t('datasetObj.createModelIntro') : $t('modelManage.createModelIntro') }}
          </el-button>
        </div>
      </div>
    </div>
    <div class="content-r" :style="{ marginTop: !introEmpty ? '28px' : '' }" v-show="!editing">
      <div class="summary">
        <div class="row">
          <div class="label">{{ $t('datasetObj.dataset_owner') }}：</div>
          <div class="value">
            <a :href="`/${dataObj.owner_name}`">{{ dataObj.owner_name }}</a>
          </div>
        </div>
        <!-- Dataset specific fields -->
        <template v-if="type === 'dataset'">
          <div class="row">
            <div class="label">{{ $t('datasetObj.category') }}：</div>
            <div class="value">
              <span v-if="dataObj.tags.length" v-for="item in dataObj.tags" :key="item">
                {{ $t(`datasets.${item}`) }}
              </span>
              <span v-else>--</span>
            </div>
          </div>
          <div class="row">
            <div class="label">{{ $t('datasetObj.application') }}：</div>
            <div class="value">
              <span v-if="dataObj.tasks.length" v-for="item in dataObj.tasks" :key="item">
                {{ $t(`datasets.${item}`) }}
              </span>
              <span v-else>--</span>
            </div>
          </div>
          <div class="row">
            <div class="label">{{ $t('datasetObj.license') }}：</div>
            <div class="value">
              <span>{{ dataObj.licenses || '--' }}</span>
            </div>
          </div>
        </template>
        <template v-else>
          <div class="row">
            <div class="label">{{ $t('modelManage.license') }}：</div>
            <div class="value">
              <a v-if="dataObj && dataObj.licenseInfo && dataObj.licenseInfo.name" target="_blank"
                :href="dataObj && dataObj.licenseInfo && dataObj.licenseInfo.linkUrl">
                {{ dataObj && dataObj.licenseInfo && dataObj.licenseInfo.name }}
              </a>
              <span v-else>--</span>
            </div>
          </div>
          <div class="row">
            <div class="label">{{ $t('modelManage.modelEngine') }}：</div>
            <div class="value">
              <span>{{ dataObj.engineName || '--' }}</span>
            </div>
          </div>
          <div class="row">
            <div class="label">{{ $t('modelManage.modelSource') }}：</div>
            <div class="value">
              <span v-if="dataObj.aimodel_type == 0">
                <span class="model-type online">{{ $t('modelManage.online') }}</span><span
                  v-html="dataObj.displayJobNameHtml"></span>
              </span>
              <span class="model-type local" v-if="dataObj.aimodel_type == 1">{{ $t('modelManage.local') }}</span>
              <div style="display:flex;">
                <span class="model-type external" v-if="dataObj.aimodel_type == 2">{{ $t('modelManage.external')
                }}</span>
                <MigrateModelSync v-if="dataObj.aimodel_type == 2" :data="dataObj"></MigrateModelSync>
              </div>
            </div>
          </div>
        </template>
        <div class="row">
          <div class="label">{{ type == 'dataset' ? $t('datasetObj.dataset_acess') : $t('modelManage.modelAccess') }}：
          </div>
          <div v-if="dataObj.is_private" class="value" style="height: 21px;">
            <span class="model-private">
              <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 32 32" width="16" height="16">
                <defs></defs>
                <g>
                  <path
                    d="M25.333 13.333h1.333c0.736 0 1.333 0.597 1.333 1.333v0 13.333c0 0.736-0.597 1.333-1.333 1.333v0h-21.333c-0.736 0-1.333-0.597-1.333-1.333v0-13.333c0-0.736 0.597-1.333 1.333-1.333v0h1.333v-1.333c0-5.155 4.179-9.333 9.333-9.333s9.333 4.179 9.333 9.333v0 1.333zM6.667 16v10.667h18.667v-10.667h-18.667zM14.667 18.667h2.667v5.333h-2.667v-5.333zM22.667 13.333v-1.333c0-3.682-2.985-6.667-6.667-6.667s-6.667 2.985-6.667 6.667v0 1.333h13.333z">
                  </path>
                </g>
              </svg>
              {{ formatAccess(dataObj.is_private) }}
            </span>

          </div>
          <div v-else class="value">{{ formatAccess(dataObj.is_private) }}
          </div>
        </div>
        <div class="row">
          <div class="label">{{ type == 'dataset' ? $t('datasetObj.dataset_size') : $t('modelManage.modelSize') }}：</div>
          <div class="value">{{ formatSize(dataObj.size) }}</div>
        </div>
        <div class="row">
          <div class="label">{{ $t('modelManage.createTime') }}：</div>
          <div class="value">{{ formatTime(dataObj.created_unix) }}</div>
        </div>
        <div class="row">
          <div class="label">{{ $t('modelManage.updateTime') }}：</div>
          <div class="value">{{ formatTime(dataObj.updated_unix) }}</div>
        </div>
      </div>

      <!-- <template v-if="modelUseTaskList.length">
        <div class="divider-column-vertical"></div>
        <div class="summary">
          <div class="title">{{ $t('modelManage.modelUseTaskList') }}：</div>
          <div class="detail-address" style="margin-left: 1rem;" v-for="(item, index) in modelUseTaskList" :key="index">
            <li class="nowrap">
              <a :href="item.url" :title="item.display_job_name">{{ item.display_job_name }}</a>
            </li>
          </div>
        </div>
      </template> -->
      <template v-if="repoUseTaskList.length">
        <div class="divider-column-vertical"></div>
        <div class="summary">
          <div class="title">{{ $t('modelManage.trainUsedRepo') }}：</div>
          <div class="detail-address" style="margin-left: 1rem;" v-for="(item, index) in repoUseTaskList" :key="index">
            <li class="nowrap">
              <a :href="item.url" :title="item.showName">{{ item.showName }}</a>
            </li>
          </div>
        </div>
      </template>
      <template v-if="trainUsedDataList.length">
        <div class="divider-column-vertical"></div>
        <div class="summary">
          <div class="title">{{ $t('modelManage.trainUsedDataList') }}：</div>
          <div class="detail-address" style="margin-left: 1rem;" v-for="item in trainUsedDataList" :key="item.name">
            <span style="display: flex;" v-if="!item.is_delete">
              <i style="margin-right: 0.5rem;" class="ri-stack-line"></i>
              <a class="nowrap" :href="item.url" :title="item.showName">{{ item.showName }}</a>
            </span>
            <span v-else style="display: flex;color:#888888" class="nowrap">
              <i style="margin-right: 0.5rem;" class="ri-stack-line"></i>
              <span style="width: 70%;" class="nowrap">{{ item.showName }}</span>
              <span>{{ $t('modelManage.deleted') }}</span>
            </span>
          </div>
        </div>
      </template>
    </div>
  </div>
</template>

<script>
import MigrateModelSync from './MigrateModelSync.vue';
import { getMarkdownHtml } from '~/apis/modules/common';
import { getDatasetReadMe, postDatasetReadMe } from "~/apis/modules/dataset";
import { getModelRelatedInfo } from '~/apis/modules/modelsquare';
import { formatDate } from 'element-ui/lib/utils/date-util';
import { transFileSize } from '~/utils';

export default {
  props: {
    dataObj: { type: Object, default: () => ({}) },
    type: { type: String, default: 'dataset' },
  },
  data() {
    return {
      loading: false,
      introFileName: '',
      introEmpty: false,
      canEdit: false,

      editor: null,
      editLoading: false,
      editing: false,
      editTab: 'edit',  // edit|preview
      emptyDefaultContent: '',

      content: '',
      htmlContent: '',
      editContent: '',

      previewLoading: false,
      previewContent: '',

      submitLoading: false,

      trainUsedDataList: [],
      modelUseTaskList: [],
      repoUseTaskList: []

    };
  },
  components: { MigrateModelSync },
  methods: {
    formatAccess(item) {
      return item ? this.$t('modelManage.modelAccessPrivate') : this.$t('modelManage.modelAccessPublic')
    },
    formatSize(size) {
      return transFileSize(size)
    },
    formatTime(time) {
      return formatDate(new Date(time * 1000), 'yyyy-MM-dd HH:mm:ss')
    },
    createIntro() {
      this.content = this.emptyDefaultContent;
      this.toggleEdit(true);
    },
    changeEditTab(tab) {
      console.log(tab == this.editTab)
      if (tab == this.editTab) return;
      this.editTab = tab;
      if (tab == 'preview') {
        getMarkdownHtml(this.editContent).then(res => {
          this.previewContent = res.data;
        }).catch(err => {
          console.log(err);
        });
      } else {
        this.$nextTick(() => {
          this.editor && this.editor.layout();
        });
      }
    },
    toggleEdit(state) {
      if (state) {
        this.editing = true;
        this.editTab = 'edit';
        this.editContent = this.content;
        this.$nextTick(() => {
          this.initEdit();
        });
      } else {
        if (this.editContent != this.content) {
          this.$confirm(this.$t('modelManage.discardFileChanges'), this.$t('tips'), {
            confirmButtonText: this.$t('confirm1'),
            cancelButtonText: this.$t('cancel'),
            type: 'warning',
            lockScroll: false,
          }).then(() => {
            this.editing = false;
            this.$nextTick(() => {
              window.initMarkdownCatalog && window.initMarkdownCatalog();
            });
          }).catch(() => { });
        } else {
          this.editing = false;
          this.$nextTick(() => {
            window.initMarkdownCatalog && window.initMarkdownCatalog();
          });
        }
      }
    },
    resize() {
      this.editor && this.editor.layout();
    },
    async initEdit() {
      this.editLoading = true;
      const monaco = await import('monaco-editor');
      this.editor && this.editor.dispose();
      this.$refs.editContainerRefinnerHTML = '';
      this.editor = monaco.editor.create(this.$refs.editContainerRef, {
        value: this.editContent,
        language: 'markdown',
        theme: document.documentElement.classList.contains('theme-arc-green') ? 'vs-dark' : 'vs',
        wordWrap: 'on'
      });
      const model = this.editor.getModel();
      model.onDidChangeContent(() => {
        this.editContent = this.editor.getValue();
      });
      window.removeEventListener('resize', this.resize);
      window.addEventListener('resize', this.resize);
      this.editLoading = false;
    },
    async submit() {
      if (!this.dataObj.id || this.submitLoading) return;
      if (this.editContent == '') {
        this.$message({
          type: 'info',
          message: this.$t('modelManage.editFileContentFirst'),
        });
        return;
      }
      this.submitLoading = true;
      try {
        let params = {
          [this.type === 'dataset' ? 'dataset_id' : 'aimodel_id']: this.dataObj.id,
        };
        let response = await postDatasetReadMe(params, { content: this.editContent }, this.type);
        const res = response.data;
        if (res.code == 0) {
          let urlText = this.type == 'dataset' ? 'datasets' : 'models'
          window.location.href = `/${urlText}/detail/${this.dataObj.owner_name}/${this.dataObj.name}`;
        } else {
          this.$message({
            type: 'error',
            message: res.msg,
          });
        }
      } catch (err) {
        console.log(err);
        this.$message({
          type: 'error',
          message: this.$t('submittedFailed'),
        });
      } finally {
        this.submitLoading = false;
      }
    },

    async getIntroInfo() {
      if (!this.dataObj.id) return;
      this.loading = true;
      try {
        let params = {
          [this.type === 'dataset' ? 'dataset_id' : 'aimodel_id']: this.dataObj.id
        };
        let reponse = await getDatasetReadMe(params, this.type);
        const res = reponse.data;
        const data = res.data;
        if (res && res.code == 0) {
          if (!data.isExistMDFile) {
            this.introEmpty = true;
            this.emptyDefaultContent = data.content;
          } else {
            this.introEmpty = false;
            this.content = data.content;
            this.htmlContent = data.htmlcontent;
          }
          this.introFileName = data.fileName;
          window.setTimeout(() => {
            this.$nextTick(() => {
              window.initMarkdownCatalog && window.initMarkdownCatalog();
            });
          }, 200);
        }
      } catch (error) {
        console.log(error);
      } finally {
        this.loading = false;
      }
    },
    async getModelRelated() {
      if (!this.dataObj.id || this.type === 'dataset') return;
      try {
        let params = { aimodel_id: this.dataObj.id };
        let reponse = await getModelRelatedInfo(params);
        const res = reponse.data;
        console.log(res)
        if (res.data && res.code == 0) {
          const data = res.data;
          this.trainUsedDataList = data.datasets.map((item) => {
            item.url = `/datasets/detail/{item.owner_name}/${item.name}`
            item.showName = `${item.owner_name}/${item.name}`
            return item
          })
          this.repoUseTaskList = data.repos.map((item) => {
            item.url = `/${item.owner_name}/${item.name}`
            item.showName = `${item.owner_name}/${item.name}`
            return item
          })
        }
      } catch (error) {
        console.log(error);
      } finally {

      }
    }
  },
  beforeMount() {
  },
  mounted() {
    if (this.type === 'dataset') {
      this.canEdit = this.dataObj?.can_edit_file
    } else {
      this.canEdit = this.dataObj?.permission?.can_edit_file
    }
    this.getIntroInfo()
    this.getModelRelated()
    console.log(this.dataObj)
  },
  beforeDestroy() {
    window.removeEventListener('resize', this.resize);
    this.editor && this.editor.dispose();
  },
};
</script>

<style scoped lang="less">
.content {
  display: flex;
  margin-top: 32px;

  .content-l {
    flex: 12;
    width: 75%;

    .read-mode {
      .head {
        color: rgb(136, 136, 136);
        display: flex;
        justify-content: space-between;
        align-items: center;
        padding: 0 4px;
        margin-bottom: 8px;

        i {
          font-size: 12px;

          &.pencil {
            cursor: pointer;
          }
        }
      }

      .content-box {
        border-style: solid;
        border-color: rgb(225, 227, 230);
        border-radius: 5px;
        border-width: 1px;
      }
    }

    .edit-mode {
      .head {
        color: rgb(136, 136, 136);
        display: flex;
        justify-content: space-between;
        align-items: center;
        padding: 0 4px;
        margin-bottom: 8px;

        i {
          font-size: 12px;

          &.reply {
            cursor: pointer;
          }
        }
      }

      .item-tab-c {
        display: flex;
        align-items: center;

        .item-tab {
          display: flex;
          align-items: center;
          padding: 10px 15px;
          color: rgba(0, 0, 0, 0.5);
          cursor: pointer;
          margin-bottom: -1px;
          border-left: 1px solid transparent;
          border-right: 1px solid transparent;
          border-top: 1px solid transparent;
          border-bottom: 1px solid transparent;

          &:hover {
            color: rgba(0, 0, 0, 0.8);
          }

          &.focus {
            color: rgba(0, 0, 0, 0.9);
            border-radius: 0.28571429rem 0.28571429rem 0 0 !important;
            border-color: #d4d4d5;
            font-weight: 550;
            background: white;
            border-bottom: 1px solid white;
            margin-bottom: -2px;
          }
        }
      }

      .tab-content {
        border: 1px solid #d4d4d5;
      }

      .preview-container {
        padding: 2em 2em 2em !important;
      }
    }

    .conmit-btn-c {
      margin-top: 20px;

      .btn {
        color: rgb(2, 0, 4);
        background-color: #db2828;
        border-color: #db2828;
        color: #fff;

        &._disabled_ {
          opacity: 0.4;
        }

        &.confirm-btn {
          color: #fff;
          background-color: rgb(56, 158, 13);
          border-color: rgb(56, 158, 13);
        }
      }
    }

    .empty {
      text-align: center;
      background: rgba(245, 245, 246, 0.5);
      height: 391px;

      .icon-c {
        display: flex;
        align-items: center;
        justify-content: center;
        padding-top: 80px;

        .empty-icon {
          height: 100px;
          width: 100px;
          background: url(/img/empty-box.svg) center center no-repeat;
        }
      }

      .tips {
        margin-top: 10px;
        font-size: 18px;
        color: rgb(63, 63, 64);
      }

      .ops {
        margin-top: 26px;

        .create-btn {
          color: #fff;
          background-color: rgb(56, 158, 13);
          border-color: rgb(56, 158, 13);
        }
      }
    }
  }

  .content-r {
    flex: 4;
    width: 25%;
    margin-left: 20px;
    background-color: rgba(247, 247, 247, 1);
    border-radius: 10px;
    padding: 16px 28px;
    padding-right: 12px;

    .title {
      font-size: 16px;
      color: rgb(16, 16, 16);
      line-height: 24px;
      margin-bottom: 8px;
    }

    .descr {
      font-weight: 300;
      font-size: 14px;
      color: rgb(136, 136, 136);
      margin-bottom: 8px;
      line-height: 20px;
    }

    .online-container {
      background: linear-gradient(180deg, rgba(203, 212, 251, 0.3) 0%, rgba(255, 255, 255, 0) 100%);
      border-color: rgb(225, 227, 230);
      border-width: 1px 0px 0px;
      border-style: solid;
      margin-top: 1.5rem;

      .online-address {
        width: 100%;
        display: flex;
        justify-content: center;
        margin-top: 1.5rem;

        a {
          width: 50%;
          background: rgb(255, 255, 255);
          color: rgb(22, 132, 252);
          border-color: rgb(22, 132, 252);
          border-width: 1px;
          border-style: solid;
          text-align: center;
          line-height: 20px;
          padding: 0.4rem 0;
        }
      }
    }


    .model-type {
      color: white;
      padding: 0px 3px;
      border-radius: 4px;
      font-size: 12px;
      margin-right: 4px;

      &.local {
        background-color: #1684fc;
      }

      &.online {
        background-color: #5bb973;
      }

      &.external {
        background-color: #5c246a;
      }
    }

    .model-private {
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

    /deep/.append-txt {
      margin-left: 4px;
      background-color: gainsboro;
      padding: 2px;
      border-radius: 2px;
      font-size: 12px;
    }


    .divider-column-vertical {
      background-color: #fff;
      height: 1px;
      margin-bottom: 1.25rem;
      margin-top: 1.25rem;
    }

    .detail-address {
      margin: 0.5rem 0;
      overflow: hidden;
      white-space: nowrap;
      text-overflow: ellipsis;
      color: rgba(5, 127, 255, 1);

      >a {
        color: rgba(5, 127, 255, 1);
      }
    }

    .summary {
      .row {
        display: flex;
        margin: 10px 0;

        .label {
          width: 100px;
          color: rgb(136, 136, 136)
        }

        .value {
          color: rgb(16, 16, 16);
          flex: 1;

          .model-type {
            color: white;
            padding: 0px 3px;
            border-radius: 4px;
            font-size: 12px;
            margin-right: 4px;

            &.local {
              background-color: #1684fc;
            }

            &.online {
              background-color: #5bb973;
            }

            &.external {
              background-color: #5c246a;
            }
          }
        }
      }

      .title {
        height: 20px;
        color: rgba(136, 136, 136, 1);
        font-size: 14px;
        margin-bottom: 1.2rem;
      }
    }
  }
}

@media screen and (max-width: 767px) {

  /* 当视口宽度 ≤ 767px 时生效 */
  /* #loadContainer { display: none; } */
  .content-r {
    display: none;
  }
}
</style>