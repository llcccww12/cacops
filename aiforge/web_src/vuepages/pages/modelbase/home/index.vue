<template>
  <div>
    <TopHeader>
      <div class="title">
        <div class="title-1">大模型基地</div>
        <div class="tab-c">
          <div class="tab" :class="tabIndex == '1' ? 'focused' : ''" @click="changeTab(1)">自然语言处理模型</div>
          <div class="tab" :class="tabIndex == '2' ? 'focused' : ''" @click="changeTab(2)">计算机视觉模型</div>
          <div class="tab" :class="tabIndex == '3' ? 'focused' : ''" @click="changeTab(3)">相关工具系统</div>
        </div>
      </div>
    </TopHeader>
    <div v-if="tabIndex == '1' || tabIndex == '2'">
      <div class="model-category-c model-category-c-bg" >
        <div class="ui container model-category-more">
          <div v-if="loading" class="more-model">数据加载中...</div>
        </div>
      </div>
      <div v-if="!loading">
        <div class="model-category-c model-category-c-bg">
          <div class="catalog-c ui container ">
            <a class="catalog-item" v-for="(item, index) in models" :key="index" :href="`#model-category-${index}`">
              {{ item.category }}
            </a>
          </div>
        </div>
        <div class="model-category-c" :class="index % 2 == 1 ? 'model-category-c-bg' : ''"
          v-for="(item, index) in  models   " :key="index">
          <div class="model-category ui container">
            <a :name="`model-category-${index}`"></a>
            <div class="model-sum">
              <div class="model-sum-head">
                <div class="model-sum-title">{{ item.category }}</div>
              </div>
              <div class="model-sum-descr">{{ item.descr }}</div>
              <div class="model-sum-op">
                <a v-if="item.repoUrl" target="_blank" :href="item.repoUrl">项目主页</a>
                <a v-if="item.homeUrl" target="_blank" :href="item.homeUrl">项目官网</a>
              </div>
            </div>
            <div class="model-list">
              <div class="model-item" :class="_item.recommend ? 'highlight' : ''" v-for="(_item, _index) in item.models"
                :key="index + '-' + _index">
                <div class="model-title">
                  <div class="tit" :title="_item.name">{{ _item.name }}</div>
                  <div style="flex-shrink: 0;">
                    <div v-if="_item.recommend" class="recommend-tip">推荐体验</div>
                  </div>
                </div>
                <div class="model-descr" :title="_item.descr">{{ _item.descr }}</div>
                <div class="model-op" :class="_item.ops.length > 1 ? 'justify-content-space-between' : ''">
                  <div class="model-op-btn-c" v-for="(op, opIndex) in _item.ops"
                    :key="index + '-' + _index + '-' + opIndex">
                    <a v-if="op.url" class="model-op-btn" :href="op.url">
                      <i v-if="op['left-icon']" :class="op['left-icon']"></i>
                      <span>{{ op.label }}</span>
                      <i v-if="op['right-icon']" :class="op['right-icon']"></i>
                    </a>
                    <div v-else class="model-op-btn-c" style="cursor: pointer;color:#3291f8;" @click="showDialog">
                      <i v-if="op['left-icon']" :class="op['left-icon']"></i>
                      <span>{{ op.label }}</span>
                      <i v-if="op['right-icon']" :class="op['right-icon']"></i>
                    </div>
                  </div>
                </div>
              </div>
            </div>
          </div>
        </div>
        <div class="model-category-c" :class="models.length % 2 == 0 ? 'model-category-c-bg' : ''">
          <div class="ui container model-category-more">
            <div class="more-model">更多大模型接入中...</div>
          </div>
        </div>
      </div>
    </div>
    <div v-if="tabIndex == '3'">
      <div class="tool-c">
        <div class="tool-list ui container">
          <div class="model-tool-list">
            <a class="model-tool" :class="`tool-bg-${index + 1}`" :href="item.href" target="_blank"
              v-for="(item, index) in tools" :key="index">
              <div class="tool-title">{{ item.name }}</div>
            </a>
          </div>
        </div>
      </div>
      <div class="model-category-c">
        <div class="ui container model-category-more">
          <div class="more-model">更多大模型工具系统接入中...</div>
        </div>
      </div>
    </div>
    <el-dialog class="task-already-dlg" :visible.sync="taskAlreadyDialogShow" :lock-scroll="false" :show-close="false">
      <div class="err-msg-box-already">
        <div class="msg-content">
          <i class="ri-information-line"></i>
          <div class="msg-content-tip">
            <div class="line-1" v-html="$t('cloudbrainObj.sameTaskTips1', { count: 1 })"></div>
            <div class="line-2" v-html="$t('cloudbrainObj.sameTaskTips2')"></div>
          </div>
        </div>
      </div>
    </el-dialog>
  </div>
</template>

<script>
import TopHeader from '../components/TopHeader.vue';
import { getPromoteData,getModelExperience } from '~/apis/modules/common';

export default {
  data() {
    return {
      tabIndex: '1',
      allModels: {},
      models: [],
      tools: [{
        name: '启智飞轮标注系统',
        href: 'https://pangu-alpha.pcl.ac.cn/dialog',
      }, {
        name: '盘古适配GCU',
        href: 'https://openi.pcl.ac.cn/Enflame/GCU_Pytorch_pangu',
      }],
      loading: false,
      taskAlreadyDialogShow: false,
      getPromoteDataFlag: false
    };
  },
  components: { TopHeader },
  methods: {
    changeTab(tab) {
      this.tabIndex = tab;
      if (tab == '1') {
        this.models = this.allModels.NLP || [];
      } else if (tab == '2') {
        this.models = this.allModels.CV || [];
      }
    },
    showDialog() {
      this.taskAlreadyDialogShow = true
    },
    getModels() {
      this.loading = true;
      getPromoteData('/model/modelbase.json').then(res => {
        try {
          const data = JSON.parse(res.data);
          this.allModels = data || {};
          this.models = data.NLP || [];
          const queryList = this.allModels.CV.concat(this.allModels.NLP)
          const sumModels = queryList.reduce((previousValue, currentValue, index) => {
              return previousValue.concat(currentValue.models)
          }, [])
          if (this.getPromoteDataFlag) {
            getModelExperience().then((res) => {
              this.loading = false;
              sumModels.forEach((element) => {
                element.ops.forEach((item) => {
                  if (item.label === "在线体验") {
                    item.url = '/extension/modelexperience/create?model=' + element.name
                    if(res.data.length){
                      res.data.forEach((task) => {
                        if (task.AppName === element.name) {
                          if (task.Status === "RUNNING") {
                            if (task.LabelName === "sd") {
                              item.url = `/extension/modelexperience/sd?id=${btoa(task.ID)}&modelName=${task.AppName}&type=${task.ComputeResource}`
                            }
                            if (task.LabelName === "chat") {
                              item.url = `/extension/modelexperience/chat?id=${btoa(task.ID)}&modelName=${task.AppName}&type=${task.ComputeResource}`
                            }
                          }else if(task.Status === "WAITING"){
                            item.url = '/cloudbrains'
                          }
                        }
                      })
                    }
                  }
                })
                
              });
            })
          } else {
            this.loading = false;
          }

        } catch (err) {
          this.loading = false;
          console.log(err);
        }
      }).catch(err => {
        console.log(err);
        this.loading = false;
      })
    },

  },
  beforeMount() {
    const isLogin = !!document.querySelector('meta[name="_uid"]');
    this.getPromoteDataFlag = isLogin? true:false
    
    this.getModels();
  },
  mounted() {
    const hash = window.location.hash;
    if (hash) {
      const ele = document.querySelector(`a[href="${hash}"]`);
      if (ele) {
        ele.click();
      }
    }
  },
  beforeDestroy() { },
};
</script>

<style scoped lang="less">
.title {
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;
  height: 100%;
  // margin-top: -10px;

  .title-1 {
    font-weight: 400;
    font-size: 28px;
    color: rgb(16, 16, 16);
    height: 42px;
  }

  .title-2 {
    font-weight: 400;
    font-size: 14px;

    a {
      color: rgba(16, 16, 16, 1);
    }
  }

  .tab-c {
    display: flex;
    align-items: center;
    justify-content: center;
    margin-top: 16px;

    .tab {
      margin: 0 10px;
      color: rgb(5, 127, 255);
      border: 1px solid rgb(5, 127, 255);
      height: 40px;
      min-width: 132px;
      display: flex;
      align-items: center;
      justify-content: center;
      cursor: pointer;

      &.focused {
        color: rgb(255, 255, 255);
        background: rgb(5, 127, 255);
      }
    }
  }
}

.model-category-c {
  padding: 25px 0 35px 0;

  .model-category {

    .model-sum {
      align-self: baseline;
      padding-top: 10px;
      padding-left: 12px;

      .model-sum-head {
        display: flex;
        align-items: center;

        .model-sum-logo {
          height: 38px;
          width: 38px;
          margin-right: 10px;

          img {
            width: 100%;
            height: 100%;
          }
        }

        .model-sum-title {
          flex: 1;
          font-size: 20px;
          color: rgb(16, 16, 16);
        }
      }

      .model-sum-descr {
        font-weight: 300;
        font-size: 14px;
        color: rgb(136, 136, 136);
        margin: 10px 0 16px 0;
      }

      .model-sum-op {
        a {
          font-size: 14px;
          color: rgb(50, 145, 248);
          font-weight: 500;
          margin-right: 10px;
        }
      }
    }

    .model-list {
      display: flex;
      flex-wrap: wrap;

      .model-item {
        width: 332px;
        margin: 10px 12px 16px 12px;
        border-color: rgba(157, 197, 226, 0.4);
        border-width: 1px;
        border-style: solid;
        box-shadow: rgba(157, 197, 226, 0.2) 0px 5px 10px 0px;
        color: rgb(16, 16, 16);
        border-radius: 6px;
        padding: 24px 16px;
        background: url("data:image/svg+xml;charset=utf8,%3Csvg%20xmlns%3D%22http%3A%2F%2Fwww.w3.org%2F2000%2Fsvg%22%20xmlns%3Axlink%3D%22http%3A%2F%2Fwww.w3.org%2F1999%2Fxlink%22%20version%3D%221.1%22%3E%3Cdefs%3E%3ClinearGradient%20id%3D%221%22%20x1%3D%220%22%20x2%3D%221%22%20y1%3D%220%22%20y2%3D%220%22%20gradientTransform%3D%22matrix(6.123233995736766e-17%2C%201%2C%20-0.2531545429373838%2C%206.123233995736766e-17%2C%200.5%2C%200)%22%3E%3Cstop%20stop-color%3D%22%23eef2ff%22%20stop-opacity%3D%221%22%20offset%3D%220%22%3E%3C%2Fstop%3E%3Cstop%20stop-color%3D%22%23ffffff%22%20stop-opacity%3D%221%22%20offset%3D%221%22%3E%3C%2Fstop%3E%3C%2FlinearGradient%3E%3C%2Fdefs%3E%3Crect%20width%3D%22100%25%22%20height%3D%22100%25%22%20fill%3D%22url(%231)%22%3E%3C%2Frect%3E%3C%2Fsvg%3E");

        .model-title {
          display: flex;
          justify-content: space-between;

          .tit {
            font-weight: 510;
            font-size: 16px;
            color: rgb(16, 16, 16);
            text-align: left;
            margin-right: 10px;
            text-overflow: ellipsis;
            overflow: hidden;
            white-space: nowrap;
          }

          .recommend-tip {
            flex-shrink: 0;
            border-radius: 4px;
            font-size: 12px;
            color: rgb(255, 255, 255);
            background: rgb(255, 98, 0);
            padding: 0px 4px;
          }
        }

        .model-descr {
          font-weight: 300;
          font-size: 12px;
          color: rgb(136, 136, 136);
          margin: 10px 0 14px 0;
          height: 60px;
          overflow: hidden;
          text-overflow: ellipsis;
          word-break: break-all;
          display: -webkit-box;
          -webkit-box-orient: vertical;
          -webkit-line-clamp: 3;
        }

        .model-op {
          display: flex;
          justify-content: center;
          align-items: center;
          flex-wrap: wrap;

          .model-op-btn-c {
            width: 50%;
            display: flex;
            justify-content: center;
            align-items: center;
            margin: 3px 0 0 0;
          }

          a.model-op-btn {
            color: rgb(50, 145, 248);
          }

          div.model-op-btn {
            font-size: 14px;
            color: rgba(136, 136, 136, 1);
          }

          &.justify-content-space-between {
            justify-content: space-between;

            .model-op-btn-c:nth-child(2n + 1) {
              justify-content: flex-start;
              padding-left: 8px;
            }

            .model-op-btn-c:nth-child(2n) {
              justify-content: flex-end;
              padding-right: 8px;
            }
          }
        }

        &.highlight {
          background: url("data:image/svg+xml;charset=utf8,%3Csvg%20xmlns%3D%22http%3A%2F%2Fwww.w3.org%2F2000%2Fsvg%22%20xmlns%3Axlink%3D%22http%3A%2F%2Fwww.w3.org%2F1999%2Fxlink%22%20version%3D%221.1%22%3E%3Cdefs%3E%3ClinearGradient%20id%3D%221%22%20x1%3D%220%22%20x2%3D%221%22%20y1%3D%220%22%20y2%3D%220%22%20gradientTransform%3D%22matrix(6.123233995736766e-17%2C%201%2C%20-0.210790712392706%2C%206.123233995736766e-17%2C%200.5%2C%200)%22%3E%3Cstop%20stop-color%3D%22%23cffff0%22%20stop-opacity%3D%221%22%20offset%3D%220%22%3E%3C%2Fstop%3E%3Cstop%20stop-color%3D%22%23ffffff%22%20stop-opacity%3D%221%22%20offset%3D%221%22%3E%3C%2Fstop%3E%3C%2FlinearGradient%3E%3C%2Fdefs%3E%3Crect%20width%3D%22100%25%22%20height%3D%22100%25%22%20fill%3D%22url(%231)%22%3E%3C%2Frect%3E%3C%2Fsvg%3E");
        }
      }
    }
  }

  &.model-category-c-bg {
    background: url("data:image/svg+xml;charset=utf8,%3Csvg%20xmlns%3D%22http%3A%2F%2Fwww.w3.org%2F2000%2Fsvg%22%20xmlns%3Axlink%3D%22http%3A%2F%2Fwww.w3.org%2F1999%2Fxlink%22%20version%3D%221.1%22%3E%3Cdefs%3E%3ClinearGradient%20id%3D%221%22%20x1%3D%220%22%20x2%3D%221%22%20y1%3D%220%22%20y2%3D%220%22%20gradientTransform%3D%22matrix(0.11899999999999993%2C%201.217%2C%20-0.0901857098765432%2C%200.11899999999999993%2C%200.269%2C%20-0.22)%22%3E%3Cstop%20stop-color%3D%22%23ffffff%22%20stop-opacity%3D%220.47%22%20offset%3D%220%22%3E%3C%2Fstop%3E%3Cstop%20stop-color%3D%22%23e5e7eb%22%20stop-opacity%3D%220.3%22%20offset%3D%221%22%3E%3C%2Fstop%3E%3C%2FlinearGradient%3E%3C%2Fdefs%3E%3Crect%20width%3D%22100%25%22%20height%3D%22100%25%22%20fill%3D%22url(%231)%22%3E%3C%2Frect%3E%3C%2Fsvg%3E");
  }

  .catalog-c {
    padding-left: 10px;
    display: flex;
    align-items: center;
    flex-wrap: wrap;

    .catalog-item {
      display: flex;
      align-items: center;
      justify-content: center;
      margin-right: 12px;
      padding: 0 20px;
      height: 40px;
      border-color: rgb(229, 231, 235);
      border-width: 1px;
      border-style: solid;
      color: rgb(16, 16, 16);
      border-radius: 5px;
      font-size: 16px;
      text-align: center;
      line-height: 23px;
      font-weight: normal;
      font-style: normal;
      background: rgb(249, 250, 251);
      margin-top: 10px;
    }
  }
}

.model-category-more {
  display: flex;
  justify-content: center;
  align-items: center;

  .more-model {
    width: 367px;
    height: 60px;
    border-color: rgba(16, 16, 16, 0.1);
    border-width: 1px;
    border-style: solid;
    border-radius: 10px;
    font-size: 14px;
    padding: 0px;
    text-align: center;
    line-height: 20px;
    font-weight: normal;
    font-style: normal;
    background: rgba(208, 231, 255, 0.3);
    display: flex;
    align-items: center;
    justify-content: center;
  }
}

.tool-c {
  display: flex;
  align-items: center;
  background: url("data:image/svg+xml;charset=utf8,%3Csvg%20xmlns%3D%22http%3A%2F%2Fwww.w3.org%2F2000%2Fsvg%22%20xmlns%3Axlink%3D%22http%3A%2F%2Fwww.w3.org%2F1999%2Fxlink%22%20version%3D%221.1%22%3E%3Cdefs%3E%3ClinearGradient%20id%3D%221%22%20x1%3D%220%22%20x2%3D%221%22%20y1%3D%220%22%20y2%3D%220%22%20gradientTransform%3D%22matrix(0.11899999999999993%2C%201.217%2C%20-0.0901857098765432%2C%200.11899999999999993%2C%200.269%2C%20-0.22)%22%3E%3Cstop%20stop-color%3D%22%23ffffff%22%20stop-opacity%3D%220.47%22%20offset%3D%220%22%3E%3C%2Fstop%3E%3Cstop%20stop-color%3D%22%23e5e7eb%22%20stop-opacity%3D%220.3%22%20offset%3D%221%22%3E%3C%2Fstop%3E%3C%2FlinearGradient%3E%3C%2Fdefs%3E%3Crect%20width%3D%22100%25%22%20height%3D%22100%25%22%20fill%3D%22url(%231)%22%3E%3C%2Frect%3E%3C%2Fsvg%3E");
  padding: 65px 0;

  .tool-list {

    .model-tool-list {
      display: flex;
      align-items: center;
      justify-content: center;

      .model-tool {
        width: 310px;
        height: 130px;
        display: flex;
        align-items: center;
        justify-content: center;
        margin: 10px 20px;

        .tool-title {
          font-size: 20px;
          color: rgb(255, 255, 255);
        }
      }
    }
  }
}

.tool-bg-1 {
  background: url("data:image/svg+xml;charset=utf8,%3Csvg%20xmlns%3D%22http%3A%2F%2Fwww.w3.org%2F2000%2Fsvg%22%20xmlns%3Axlink%3D%22http%3A%2F%2Fwww.w3.org%2F1999%2Fxlink%22%20version%3D%221.1%22%3E%3Cdefs%3E%3ClinearGradient%20id%3D%221%22%20x1%3D%220%22%20x2%3D%221%22%20y1%3D%220%22%20y2%3D%220%22%20gradientTransform%3D%22null%22%3E%3Cstop%20stop-color%3D%22%236bcbb5%22%20stop-opacity%3D%221%22%20offset%3D%220%22%3E%3C%2Fstop%3E%3Cstop%20stop-color%3D%22%238799e4%22%20stop-opacity%3D%221%22%20offset%3D%221%22%3E%3C%2Fstop%3E%3C%2FlinearGradient%3E%3C%2Fdefs%3E%3Crect%20width%3D%22100%25%22%20height%3D%22100%25%22%20fill%3D%22url(%231)%22%3E%3C%2Frect%3E%3C%2Fsvg%3E");
}

.tool-bg-2 {
  background: url("data:image/svg+xml;charset=utf8,%3Csvg%20xmlns%3D%22http%3A%2F%2Fwww.w3.org%2F2000%2Fsvg%22%20xmlns%3Axlink%3D%22http%3A%2F%2Fwww.w3.org%2F1999%2Fxlink%22%20version%3D%221.1%22%3E%3Cdefs%3E%3ClinearGradient%20id%3D%221%22%20x1%3D%220%22%20x2%3D%221%22%20y1%3D%220%22%20y2%3D%220%22%20gradientTransform%3D%22matrix(0.9919999999999999%2C%200.803%2C%20-0.14121436004162333%2C%200.9919999999999999%2C%200%2C%200)%22%3E%3Cstop%20stop-color%3D%22%23e2ad71%22%20stop-opacity%3D%221%22%20offset%3D%220%22%3E%3C%2Fstop%3E%3Cstop%20stop-color%3D%22%23f9b96f%22%20stop-opacity%3D%221%22%20offset%3D%220.5%22%3E%3C%2Fstop%3E%3Cstop%20stop-color%3D%22%23f57d34%22%20stop-opacity%3D%221%22%20offset%3D%220.99%22%3E%3C%2Fstop%3E%3C%2FlinearGradient%3E%3C%2Fdefs%3E%3Crect%20width%3D%22100%25%22%20height%3D%22100%25%22%20fill%3D%22url(%231)%22%3E%3C%2Frect%3E%3C%2Fsvg%3E");
}
.task-already-dlg {
  margin-top: 15vh;

  /deep/.el-dialog__header {
    display: none;
  }

  /deep/.el-dialog__body {
    padding: 0px 0px;
  }
}
.err-msg-box-already {
  padding: 1em 1.5em;
  background-color: rgba(242, 113, 28, 0.05);
  border: 2px solid rgba(242, 113, 28, 1);
  border-radius: 5px;

  .msg-content {
    display: flex;
    align-items: center;

    i {
      font-size: 35px;
      color: rgba(242, 113, 28, 1);
    }

    .msg-content-tip {
      text-align: left;
      margin-left: 1rem;

      .line-1 {
        font-weight: 600;
        line-height: 2;

        span {
          color: rgba(242, 113, 28, 1);
        }
      }

      .line-2 {
        color: #939393
      }
    }
  }
}
</style>
