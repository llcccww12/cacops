<template>
  <div class="repo-extend-info-c">
    <div v-if="true && templateList.length" class="info-block templates">
      <div class="title-c">
        <div class="title">{{ $t('taskTmplObj.taskTmplReferencedRepo') }}</div>
        <div class="more" v-if="templateCanExpand"
          @click="templateIsExpanded ? fold('template') : more('template', true)">{{ templateIsExpanded ?
            $t('taskTmplObj.collapsed') :
            $t('taskTmplObj.more') }}</div>
      </div>
      <div class="list">
        <a class="item template" target="_blank" :href="tmpl.link" v-for="tmpl in templateListShow" :key="tmpl.id">
          <div class="icon-c">
            <svg xmlns="http://www.w3.org/2000/svg" fill="rgba(31, 34, 38, 0.7)" viewBox="0 0 48 48" width="16"
              height="16">
              <defs></defs>
              <g>
                <path d="M48 0H0V48H48V0Z" fill-opacity="0.01"></path>
                <path d="M44 14L24 4L4 14V34L24 44L44 34V14Z" fill="none" stroke="rgba(31, 34, 38, 0.7)"
                  stroke-width="4" stroke-linejoin="round">
                </path>
                <path d="M4 14L24 24" fill="none" stroke="rgba(31, 34, 38, 0.7)" stroke-width="4" stroke-linecap="round"
                  stroke-linejoin="round">
                </path>
                <path d="M24 44V24" fill="none" stroke="rgba(31, 34, 38, 0.7)" stroke-width="4" stroke-linecap="round"
                  stroke-linejoin="round">
                </path>
                <path d="M44 14L24 24" fill="none" stroke="rgba(31, 34, 38, 0.7)" stroke-width="4"
                  stroke-linecap="round" stroke-linejoin="round">
                </path>
                <path d="M34 9L14 19" fill="none" stroke="rgba(31, 34, 38, 0.7)" stroke-width="4" stroke-linecap="round"
                  stroke-linejoin="round">
                </path>
              </g>
            </svg>
            <svg v-if="tmpl.IsPrivate" class="lock" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 32 32" width="6"
              height="6">
              <defs></defs>
              <g>
                <path
                  d="M25.333 13.333h1.333c0.736 0 1.333 0.597 1.333 1.333v0 13.333c0 0.736-0.597 1.333-1.333 1.333v0h-21.333c-0.736 0-1.333-0.597-1.333-1.333v0-13.333c0-0.736 0.597-1.333 1.333-1.333v0h1.333v-1.333c0-5.155 4.179-9.333 9.333-9.333s9.333 4.179 9.333 9.333v0 1.333zM6.667 16v10.667h18.667v-10.667h-18.667zM14.667 18.667h2.667v5.333h-2.667v-5.333zM22.667 13.333v-1.333c0-3.682-2.985-6.667-6.667-6.667s-6.667 2.985-6.667 6.667v0 1.333h13.333z">
                </path>
              </g>
            </svg>
          </div>
          <div class="name" :title="tmpl.showName">{{ tmpl.showName }}</div>
        </a>
      </div>
    </div>
    <div v-if="false && modelList.length" class="info-block models">
      <div class="title-c">
        <div class="title">{{ `` }}</div>
        <div class="more" v-if="modelCanExpand" @click="modelIsExpanded ? fold('model') : more('model', true)">
          {{ modelIsExpanded ? $t('taskTmplObj.collapsed') :
            $t('taskTmplObj.more') }}</div>
      </div>
      <div class="list">
        <a class="item model" target="_blank" :href="model.link" v-for="model in modelListShow" :key="model.id">
          <div class="icon-c">
          </div>
          <div class="name" :title="model.showName">{{ model.showName }}</div>
        </a>
      </div>
    </div>
    <div v-if="false && datasetList.length" class="info-block datasets">
      <div class="title-c">
        <div class="title">{{ `` }}</div>
        <div class="more" v-if="datasetCanExpand" @click="datasetIsExpanded ? fold('dataset') : more('dataset', true)">
          {{ datasetIsExpanded ?
            $t('taskTmplObj.collapsed') :
            $t('taskTmplObj.more') }}</div>
      </div>
      <div class="list">
        <a class="item model" target="_blank" :href="dataset.link" v-for="dataset in datasetListShow" :key="dataset.id">
          <div class="icon-c">
          </div>
          <div class="name" :title="dataset.showName">{{ dataset.showName }}</div>
        </a>
      </div>
    </div>
  </div>
</template>

<script>
import { getAiTaskTmplList } from '~/apis/modules/aitasktmpl';

const dataEle = document.getElementById('__repos_extend_information-c');

export default {
  data() {
    return {
      repoID: dataEle.getAttribute('repo-id'),
      repoOwnerName: dataEle.getAttribute('repo-owner-name'),
      repoName: dataEle.getAttribute('repo-name'),
      templateList: [],
      templateListShow: [],
      templateIsExpanded: false,
      templateCanExpand: false,
      modelList: [],
      modelListShow: [],
      modelIsExpanded: false,
      modelCanExpand: false,
      datasetList: [],
      datasetListShow: [],
      datasetIsExpanded: false,
      datasetCanExpand: false,
      initLength: 5,
    };
  },
  components: {},
  methods: {
    getData(type, isMore) {
      let getApi = getAiTaskTmplList;
      let params = {};
      if (type == 'template') {
        getApi = getAiTaskTmplList;
        params = {
          type: 'repo',
          repo: this.repoID,
          page: 1,
          page_size: isMore ? 100 : this.initLength
        }
      }
      if (type == 'model') { }
      if (type == 'dataset') { }
      getApi(params).then(res => {
        res = res.data;
        if (type == 'template') {
          if (res.code == 0) {
            const total = res.data?.Total;
            const list = (res.data?.Templates || []).map(item => {
              const owner = item.Owner?.Name || ''
              return {
                showName: `${owner ? (owner + '/') : ''}${item.Name}`,
                id: item.ID,
                link: `/ai_task_tmpl/detail/${item.ID}`
              }
            });
            this.templateList = list;
            this.templateListShow = list;
            this.templateCanExpand = total > this.initLength;
            if (isMore) {
              this.templateIsExpanded = true;
            }
          }
        }
        if (type == 'model') { }
        if (type == 'dataset') { }
      }).catch(err => {
        console.log(err);
      })
    },
    more(type) {
      if (this[`${type}List`].length > this.initLength) {
        this[`${type}ListShow`] = this[`${type}List`];
        this[`${type}IsExpanded`] = true;
      } else {
        this.getData(type, true);
      }
    },
    fold(type) {
      this[`${type}ListShow`] = this[`${type}List`].slice(0, this.initLength);
      this[`${type}IsExpanded`] = false;
    }
  },
  beforeMount() {
    this.getData('template', false);
    // this.getData('model', false);
    // this.getData('dataset', false);
  },
  mounted() { },
  beforeDestroy() { },
};
</script>

<style scoped lang="less">
.repo-extend-info-c {
  padding-top: 4px;

  .info-block {
    border-top: 1px solid rgb(225, 227, 230);
    margin-top: 14px;

    .title-c {
      display: flex;
      justify-content: space-between;
      margin: 8px 0 4px 0;

      .title {
        color: rgb(16, 16, 16);
        font-size: 16px;
        font-weight: 700;
        line-height: 24px;
      }

      .more {
        color: rgba(116, 129, 141, 1);
        font-size: 14px;
        cursor: pointer;
        line-height: 24px;
        margin-left: 10px;
      }
    }

    .list {
      .item {
        display: flex;
        align-items: center;
        height: 28px;
        cursor: pointer;

        .icon-c {
          height: 20px;
          width: 20px;
          display: flex;
          align-items: center;
          justify-content: center;
          position: relative;
          margin-right: 5px;

          .lock {
            position: absolute;
            right: 2px;
            bottom: 5px;
            background: white;
            border-radius: 2px;
          }
        }

        .name {
          color: rgba(98, 120, 199, 1);
          font-size: 14px;
          overflow: hidden;
          text-overflow: ellipsis;
          white-space: nowrap;
        }
      }
    }
  }
}
</style>
