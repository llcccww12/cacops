<template>
  <div class="content-box">
    <Header>
      <template slot="left">
        <div class="special-label">{{ $t('modelSquare.largeModelsList') }}</div>
      </template>
      <template slot="right">
        <div class="filter-c">
          <el-checkbox v-model="filterExperience" @change="filterHandler">{{ $t('modelObj.can_online_infer')
            }}</el-checkbox>
          <el-checkbox v-model="filterFinetuning" @change="filterHandler">{{ $t('modelObj.can_fine_tune')
            }}</el-checkbox>
          <el-input class="search-inp" v-model="searchVal" :placeholder="$t('modelObj.model_search_placeholder')"
            @input="filterHandler"></el-input>
          <el-button class="search-btn" @click="filterHandler">{{ $t('repos.search') }}</el-button>
        </div>
      </template>
    </Header>
    <div class="main-body">
      <div class="menu">
        <template v-for="item in menuData">
          <div class="menu-item menu-group" :class="menuActive == item.key ? 'active' : ''" v-if="item.title"
            @click="changeCategory(item)"><i class="ri-arrow-down-s-line"></i>{{ item.title }}</div>
          <div class="menu-item" :class="menuActive == _item.key ? 'active' : ''" v-for="_item in item.children"
            @click="changeCategory(_item)">
            <div class="title" :title="_item.title">{{ _item.title }}</div>
          </div>
        </template>
      </div>
      <div class="content" ref="contentRef">
        <div class="content-body">
          <div v-for="(item) in modelData">
            <div v-for="_item in item.list">
              <div class="model-category">
                <a class="anchor-nav" :id="encodeURIComponent(_item.category)"
                  :category="encodeURIComponent(_item.category)" href="javascript:;"></a>
                <div class="category-head">
                  <div class="category-icon">
                    <img :src="_item.icon.url || '/img/model/model-placeholder.png'" :style="_item.icon.style || ''" />
                  </div>
                  <div class="category-title">
                    <div :title="_item.category">{{ _item.category }}</div>
                  </div>
                </div>
                <div class="category-body">
                  <div class="model-item" v-for="model in _item.models">
                    <a class="anchor-nav" :id="encodeURIComponent(model.name)"
                      :category="encodeURIComponent(model.category)" href="javascript:;"></a>
                    <div class="model-item-head">
                      <div class="model-item-icon">
                        <img src="/img/model/model-placeholder.png" alt="">
                      </div>
                      <div class="model-item-title">
                        <div :title="model.name" v-html="renderMatchText(model.name)"></div>
                      </div>
                    </div>
                    <div class="model-item-descr">
                      <div :title="model.descr">{{ model.descr }}</div>
                    </div>
                    <div class="model-item-ops">
                      <a class="model-item-btn" v-if="model.detail" target="_blank" :href="model.detail">
                        {{ `查看详情` }}
                      </a>
                      <div class="model-item-btn" v-if="model.experience"
                        :class="model.experience.recommend ? 'recommend' : ''" @click="goExperience(model)">
                        {{ model.experience.recommend ? '推荐体验' : '在线体验' }}
                      </div>
                      <div class="model-item-btn" v-if="model.experienceGpu"
                        :class="{ recommend: model.experienceGpu.recommend, free: model.experienceGpu.free }"
                        @click="goExperience({ ...model, experience: model.experienceGpu })">
                        {{ model.experienceGpu.title || 'GPU部署体验' }}
                      </div>
                      <div class="model-item-btn" v-if="model.experienceNpu"
                        :class="{ recommend: model.experienceNpu.recommend, free: model.experienceNpu.free }"
                        @click="goExperience({ ...model, experience: model.experienceNpu })">
                        {{ model.experienceNpu.title || 'NPU在线体验' }}
                      </div>
                      <div class="model-item-btn" v-if="model.finetuning" @click="goFinetuning(model)">{{ `模型微调` }}
                      </div>
                    </div>
                  </div>
                </div>
              </div>
            </div>
          </div>
          <div class="extra-info-c"
            v-if="!loading && menuActive == 'all' && (searchVal.trim() == '' && !filterExperience && !filterFinetuning) && modelData.length != 0">
            <div class="extra-info">更多大模型接入中...</div>
          </div>
        </div>
        <div class="extra-info-c no-data" v-if="!loading && modelData.length == 0">
          <div class="extra-info">暂无数据</div>
        </div>
        <div class="extra-info-c loading" v-if="loading">
          <div class="extra-info">数据加载中...</div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import Header from '../../components/Header.vue';
import { getPromoteData, getModelExperience } from '~/apis/modules/common';

export default {
  name: 'ModelList',
  data() {
    return {
      menuData: [{ title: '', children: [{ title: '所有大模型', key: 'all' }] }],
      menuActive: 'all',
      filterExperience: false,
      filterFinetuning: false,
      searchVal: '',
      modelData: [],
      loading: true,
      collectData: {
        originData: [],
        typeList: [],
        categoryList: [],
        modelList: [],
      },
    };
  },
  components: { Header },
  methods: {
    changeCategory(item) {
      this.menuActive = item.key;
      this.searchVal = '';
      this.filterExperience = false;
      this.filterFinetuning = false;
      this.filterData();
    },
    filterHandler() {
      this.filterData();
    },
    filterData() {
      const modelData = [];
      const searchVal = this.searchVal.trim().toLocaleLowerCase();
      for (let i = 0, iLen = this.collectData.originData.length; i < iLen; i++) {
        const dataI = this.collectData.originData[i];
        const typeI = {
          ...dataI,
          list: [],
        };
        for (let j = 0, jLen = dataI.list.length; j < jLen; j++) {
          const dataJ = dataI.list[j];
          dataJ.typeName = dataI.name;
          const listI = {
            ...dataJ,
            models: [],
          };
          if (this.menuActive != 'all' && dataJ.category != this.menuActive && dataJ.typeName != this.menuActive) continue;
          for (let k = 0, kLen = dataJ.models.length; k < kLen; k++) {
            const model = dataJ.models[k];
            if (this.filterExperience && !model.experience && !model.experienceGpu && !model.experienceNpu) continue;
            if (this.filterFinetuning && !model.finetuning) continue;
            if (!searchVal || model.name.toLocaleLowerCase().indexOf(searchVal) > -1) {
              listI.models.push(model);
            }
          }
          if (listI.models.length) {
            typeI.list.push(listI);
          }
        }
        if (typeI.list.length) {
          modelData.push(typeI);
        }
      }
      this.modelData = modelData;
      this.$nextTick(() => {
        this.$refs.contentRef.scrollTo({ top: 0, behavior: 'instant' });
      });
    },
    renderMatchText(string) {
      const searchVal = this.searchVal.trim();
      if (!searchVal) return string;
      const reg = new RegExp(searchVal, 'gi');
      return string.replace(reg, (txt) => {
        return `<span class="mark">${txt}</span>`;
      });
    },
    checkLogin() {
      const isLogin = !!document.querySelector('meta[name="_uid"]');
      if (isLogin) return true;
      window.location.href = `/user/login?redirect_to=${encodeURIComponent(window.location.href)}`;
      return false;
    },
    goExperience(model) {
      if (model.experience.url) {
        window.open(model.experience.url, '_blank');
        return;
      }
      if (!this.checkLogin()) return;
      if (this.operating) return;
      this.operating = true;
      getModelExperience().then(res => {
        this.operating = false;
        res = res.data;
        const task = res.filter(item => item.AppName == model.name)[0];
        if (task) {
          if (task.Status == 'RUNNING') {
            window.open(`/extension/modelexperience/${task.LabelName}?id=${btoa(task.ID)}&modelName=${task.AppName}`, '_blank');
          } else if (task.Status == 'WAITING') {
            this.$router.push('/experience');
          }
        } else {
          this.$router.push({
            path: '/experience/create',
            query: {
              modelName: encodeURIComponent(model.name),
              backpath: '/list',
            }
          });
        }
      }).catch(err => {
        this.operating = false;
        console.log(err);
        this.$message.error(this.$t('operationFailed'));
      })
    },
    goFinetuning(model) {
      if (model.finetuning && this.checkLogin()) {
        if (model.type === "NLP") {
          this.$router.push({
            path: 'nlp/sft/create',
            query: {
              model: encodeURIComponent(model.name),
              backpath: '/list',
            }
          });
        }
        if (model.type === "CV") {
          this.$router.push({
            path: 'cv/sft/create',
            query: {
              model: encodeURIComponent(model.name),
              backpath: '/list',
            }
          });
        }
      }
    },
    getModels() {
      this.loading = true;
      getPromoteData('model/modelbasenew.json').then(res => {
        res = res.data;
        try {
          const finetuneModels = this.finetuneModels.map(item => item.name);
          const data = JSON.parse(res);
          const menuData = [];
          for (let i = 0, iLen = data.length; i < iLen; i++) {
            const dataI = data[i];
            const menu = {
              key: dataI.name,
              title: dataI.name,
              children: [],
            };
            this.collectData.typeList.push(dataI.name);
            for (let j = 0, jLen = dataI.list.length; j < jLen; j++) {
              const dataJ = dataI.list[j];
              dataJ.typeName = dataI.name;
              menu.children.push({ title: dataJ.category, key: dataJ.category });
              this.collectData.categoryList.push(dataJ.category);
              for (let k = 0, kLen = dataJ.models.length; k < kLen; k++) {
                const model = dataJ.models[k];
                model.type = dataI.type;
                model.typeName = dataI.name;
                model.category = dataJ.category;
                model.finetuning = finetuneModels.indexOf(model.name) >= 0;
                this.collectData.modelList.push(model);
              }
            }
            menuData.push(menu);
          }
          this.menuData.push(...menuData);
          this.menuActive = 'all';
          this.modelData = data;
          this.collectData.originData = data;
        } catch (err) {
          console.log(err);
        }
        this.loading = false;
      }).catch(err => {
        console.log(err);
        this.loading = false;
      });
    },
  },
  beforeCreate() {
    this.loading = true;
    Promise.all([new Promise((resolve, reject) => {
      getPromoteData('model/modelfinetune.json').then(res => {
        res = res.data;
        try {
          const finetuneModels = (JSON.parse(res) || {})?.llm?.model || [];
          resolve(finetuneModels);
        } catch (err) {
          console.log(err);
          resolve([]);
        }
      }).catch(err => {
        console.log(err);
        resolve([]);
      });
    }), new Promise((resolve, reject) => {
      getPromoteData('model/modelbasenew.json').then(res => {
        res = res.data;
        try {
          const data = JSON.parse(res);
          resolve(data);
        } catch (err) {
          console.log(err);
          resolve([]);
        }
      }).catch(err => {
        console.log(err);
        resolve([]);
      });
    })]).then(res => {
      const finetuneModels = res[0];
      const modelsInfo = res[1];
      const finetuneModelsNameList = finetuneModels.map(item => item.name);
      // console.log("finetuneModels", finetuneModels)
      // console.log("modelsInfo", modelsInfo)
      // console.log("finetuneModelsNameList", finetuneModelsNameList)
      const menuData = [];
      for (let i = 0, iLen = modelsInfo.length; i < iLen; i++) {
        const dataI = modelsInfo[i];
        const menu = {
          key: dataI.name,
          title: dataI.name,
          children: [],
        };
        this.collectData.typeList.push(dataI.name);
        for (let j = 0, jLen = dataI.list.length; j < jLen; j++) {
          const dataJ = dataI.list[j];
          menu.children.push({ title: dataJ.category, key: dataJ.category });
          this.collectData.categoryList.push(dataJ.category);
          for (let k = 0, kLen = dataJ.models.length; k < kLen; k++) {
            const model = dataJ.models[k];
            model.type = dataI.type;
            model.typeName = dataI.name;
            model.category = dataJ.category;
            model.finetuning = finetuneModelsNameList.indexOf(model.name) >= 0;
            this.collectData.modelList.push(model);
          }
        }
        menuData.push(menu);
      }
      this.menuData.push(...menuData);
      this.menuActive = 'all';
      this.modelData = modelsInfo;
      this.collectData.originData = modelsInfo;
      this.loading = false;
      // console.log("menuData", this.modelData)
    }).catch(err => {
      this.loading = false;
      console.log(err);
    });
  },
  mounted() { },
  beforeDestroy() { },
};
</script>

<style scoped lang="less">
.content-box {
  height: 100%;
  width: 100%;
  display: flex;
  flex-direction: column;

  .filter-c {
    display: flex;
    align-items: center;

    .el-checkbox {
      margin-right: 20px;

      &.is-checked {
        /deep/.el-checkbox__label {
          color: #606266;
        }
      }
    }

    .search-inp {
      color: rgb(136, 136, 136);
      border-radius: 4px 0 0 4px;
      height: 32px;
      display: flex;
      align-items: center;
      justify-content: center;

      /deep/.el-input__inner {
        border: none;
        height: 32px;
        border-radius: 0;
        border-color: rgb(210, 210, 216);
        border-width: 1px;
        border-style: solid;
        border-radius: 4px 0 0 4px;
      }
    }

    .search-btn {
      font-family: PingFangSC;
      font-size: 14px;
      color: rgb(255, 255, 255);
      background: rgb(111, 118, 165);
      border-radius: 0 4px 4px 0;
      border-color: rgb(111, 118, 165);
      border-width: 1px;
      border-style: solid;
      height: 32px;
      display: flex;
      align-items: center;
      justify-content: center;
    }
  }

  .main-body {
    flex: 1;
    height: 0;
    position: relative;
    margin-right: 16px;

    .menu {
      position: absolute;
      top: 16px;
      left: 16px;
      width: 183px;
      height: calc(100% - 32px);
      border-color: rgb(210, 210, 216);
      border-width: 1px;
      border-style: solid;
      border-radius: 10px;
      background-color: rgb(255, 255, 255);
      overflow-y: auto;
      padding: 10px 0 10px 0;

      .menu-item {
        height: 36px;
        display: flex;
        align-items: center;
        margin: 0 10px 0 10px;
        padding-left: 16px;
        cursor: pointer;

        .title {
          width: 100%;
          line-height: 36px;
          overflow: hidden;
          text-overflow: ellipsis;
          white-space: nowrap;
        }

        &.active {
          color: rgb(0, 122, 255);
          background-color: #f5f5f6;
          border-radius: 4px;
        }

        &:hover {
          color: rgba(0, 122, 255, 0.8);
        }
      }

      .menu-group {
        color: rgb(111, 118, 165);
        padding-left: 10px;

        i {
          margin-right: 2px;
        }
      }
    }

    .content {
      height: calc(100% - 32px);
      margin-left: 200px;
      margin-top: 16px;
      margin-bottom: 16px;
      overflow-y: auto;
      padding: 0 10px 10px 10px;
      scroll-behavior: smooth;

      .content-body {
        .model-category {

          .category-head {
            display: flex;
            align-items: center;
            height: 40px;
            margin: 10px 0;
            padding: 0 16px;

            .category-icon {
              height: 36px;
              margin-right: 10px;
              display: flex;
              align-items: center;
              justify-content: center;

              img {
                height: 100%;
              }
            }

            .category-title {
              color: rgb(16, 16, 16);
              font-size: 20px;
              font-family: SourceHanSansSC;
              font-weight: 700;
              line-height: 30px;
              flex: 1;
              width: 0;

              div {
                overflow: hidden;
                text-overflow: ellipsis;
                white-space: nowrap;
              }
            }
          }

          .category-body {
            display: flex;
            flex-wrap: wrap;

            .model-item {
              background: rgb(253, 253, 253);
              border-radius: 6px;
              box-shadow: rgba(157, 197, 226, 0.2) 0px 5px 10px 0px;
              border-color: rgba(157, 197, 226, 0.2);
              border-width: 1px;
              border-style: solid;
              width: 520px;
              height: 151px;
              padding: 14px;
              margin: 10px;
              position: relative;
              // max-width: 520px;

              .anchor-nav {
                position: absolute;
                top: -6px;
              }

              .model-item-head {
                display: flex;
                align-items: center;

                .model-item-icon {
                  height: 40px;
                  width: 40px;
                  margin-right: 10px;

                  img {
                    height: 100%;
                    width: 100%;
                  }
                }

                .model-item-title {
                  color: rgba(16, 16, 16, 1);
                  font-size: 16px;
                  font-family: SourceHanSansSC;
                  font-weight: 500;
                  line-height: 23px;
                  flex: 1;
                  width: 0;

                  div {
                    overflow: hidden;
                    text-overflow: ellipsis;
                    white-space: nowrap;

                    /deep/.mark {
                      color: red;
                    }
                  }
                }
              }

              .model-item-descr {
                margin-left: 50px;
                color: rgb(136, 136, 136);
                font-size: 12px;
                font-family: SourceHanSansSC;
                font-weight: 300;
                line-height: 17px;
                height: 34px;
                overflow: hidden;
                text-overflow: ellipsis;
                word-break: break-all;
                display: -webkit-box;
                -webkit-box-orient: vertical;
                -webkit-line-clamp: 2;
              }

              .model-item-ops {
                margin-left: 50px;
                margin-top: 10px;
                display: flex;
                align-items: center;

                .model-item-btn {
                  border-color: rgb(229, 229, 229);
                  border-width: 1px;
                  border-style: solid;
                  height: 30px;
                  width: 33%;
                  font-family: PingFangSC;
                  font-weight: 400;
                  display: flex;
                  align-items: center;
                  justify-content: center;
                  color: rgba(96, 98, 102, 1);
                  font-size: 12px;
                  cursor: pointer;
                  border-right: none;

                  &.recommend {
                    color: rgb(255, 98, 0);
                    border-color: rgb(255, 98, 0);
                    border-width: 1px;
                    border-style: solid;

                    &:hover {
                      color: rgba(255, 98, 0, 0.8);
                    }

                    &+.recommend {
                      border-left: none;
                    }
                  }

                  &.free {
                    position: relative;

                    &::after {
                      position: absolute;
                      content: "免费";
                      top: -8px;
                      right: -6px;
                      height: 16px;
                      width: 32px;
                      border-radius: 20px;
                      background: red;
                      color: white;
                      display: flex;
                      align-items: center;
                      justify-content: center;
                    }
                  }

                  &:hover {
                    color: rgba(96, 98, 102, 0.8);
                  }

                  &:first-child {
                    border-top-left-radius: 5px;
                    border-bottom-left-radius: 5px;
                  }

                  &:last-child {
                    border-top-right-radius: 5px;
                    border-bottom-right-radius: 5px;
                    border-right: 1px solid rgb(229, 229, 229);

                    &.recommend {
                      border-right: 1px solid rgb(255, 98, 0);
                    }
                  }
                }
              }
            }
          }
        }
      }

      .extra-info-c {
        display: flex;
        align-items: center;
        justify-content: center;
        margin-top: 20px;
        margin-bottom: 10px;

        &.no-data,
        &.loading {
          height: 100%;
          margin-top: 0;
          margin-bottom: 0;
        }

        .extra-info {
          width: 367px;
          height: 60px;
          border: 1px solid rgba(16, 16, 16, .1);
          border-radius: 10px;
          font-size: 14px;
          padding: 0;
          text-align: center;
          line-height: 20px;
          font-weight: 400;
          font-style: normal;
          background: rgba(208, 231, 255, .3);
          display: flex;
          align-items: center;
          justify-content: center;
        }
      }
    }
  }
}

@media only screen and (min-width: 1600px) {
  .content-box .main-body .content .content-body .model-category .category-body .model-item {
    width: calc(33.3% - 20px);
  }
}

@media only screen and (min-width: 1200px) and (max-width: 1599.9px) {
  .content-box .main-body .content .content-body .model-category .category-body .model-item {
    width: calc(50% - 20px);
  }
}

@media only screen and (max-width: 1199.9px) {
  .content-box .main-body .content .content-body .model-category .category-body .model-item {
    width: calc(100% - 20px);
  }
}

/* 手机端样式 */
@media (max-width: 768px) {
  .content-box {
    .main-body {
      .menu {
        display: none;
      }

      .content {
        margin-left: 0;
      }
    }
  }
}
</style>
