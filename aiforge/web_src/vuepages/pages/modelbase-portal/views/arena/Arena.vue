<template>
  <div class="content-box">
    <Header>
      <template slot="right">
        <el-button :disabled="viewType != 'conversation'" class="op-btn" type="primary" @click="startNewRound">
          <div class="btn-content">{{ $t('modelSquare.newRound') }}</div>
        </el-button>
      </template>
    </Header>
    <div class="main-body">
      <div class="content">
        <div class="arena-container">
          <div class="top">
            <div class="left" v-loading="initLoading">
              <ModelSelect :model="model1" :models="modelInfo" :disabled-model="model2.name" v-if="viewType == 'select'"
                @changeModel="changeModel1">
              </ModelSelect>
              <SessionWin ref="sessionWinRef1" key="sessionWinRef1" v-if="viewType == 'conversation'" :model="model1"
                :config="experienceConfig" @finish="onFinish1" headBg="rgba(191, 221, 255, 0.3)"></SessionWin>
            </div>
            <div class="right" v-loading="initLoading">
              <ModelSelect :model="model2" :models="modelInfo" :disabled-model="model1.name" v-if="viewType == 'select'"
                @changeModel="changeModel2">
              </ModelSelect>
              <SessionWin ref="sessionWinRef2" key="sessionWinRef2" v-if="viewType == 'conversation'" :model="model2"
                :config="experienceConfig" @finish="onFinish2">
              </SessionWin>
            </div>
          </div>
          <div class="bottom">
            <div class="send-box" :class="canSend ? '' : 'disabled-send'">
              <textarea class="send-txt" :disabled="!canSend" :placeholder="$t('modelSquare.promptPlaceholder')"
                v-model="prompt" @keydown.enter="carriageReturn($event)"></textarea>
              <div class="btn btn-1" v-if="canSend" :class="responsing ? 'disabled-send' : ''" @click="submit">
                <i class="ri-send-plane-line"></i>
                <span>{{ $t('modelSquare.send') }}</span>
              </div>
              <div class="btn btn-2" v-else>{{ $t('modelSquare.pleaseSelectModelAndConversation') }}</div>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import Header from '../../components/Header.vue';
import ModelSelect from './comps/ModelSelect.vue';
import SessionWin from './comps/SessionWin.vue';
import { getPromoteData } from '~/apis/modules/common';

export default {
  name: 'Arena',
  data() {
    return {
      viewType: 'select', // select|conversation
      experienceConfig: {},
      modelInfo: [],
      model1: {
        name: '',
        category: '',
        icon: ''
      },
      model2: {
        name: '',
        category: '',
        icon: ''
      },
      model1Finish: false,
      model2Finish: false,
      initLoading: true,
      canSend: false,
      prompt: '',
      responsing: false,
    };
  },
  components: { Header, ModelSelect, SessionWin },
  methods: {
    changeModel1(value) {
      this.model1 = { ...value };
      this.checkSend()
    },
    changeModel2(value) {
      this.model2 = { ...value };
      this.checkSend()
    },
    checkSend() {
      this.canSend = this.model1.name && this.model2.name
    },
    onFinish1() {
      this.model1Finish = true
      if (this.model2Finish) {
        this.responsing = false;
      }
    },
    onFinish2() {
      this.model2Finish = true
      if (this.model1Finish) {
        this.responsing = false;
      }
    },
    carriageReturn(event) {
      event.preventDefault()
      if (event.ctrlKey && event.keyCode == 13) {
        this.prompt = this.prompt + '\n'
      } else {
        this.submit()
      }
    },
    startNewRound() {
      this.viewType = 'select';
      this.model1Finish = false;
      this.model2Finish = false;
      this.prompt = '';
      this.responsing = false;
      this.checkSend()
    },
    async submit() {
      if (this.responsing) return;
      const re = new RegExp("^[ ]+$")
      if (!this.prompt || re.test(this.prompt)) {
        this.$message.error(this.$t('modelSquare.inputNotEmpty'))
        return
      }
      this.viewType = 'conversation';
      this.responsing = true;
      this.model1Finish = false;
      this.model2Finish = false;
      const prompt = this.prompt;
      this.prompt = '';
      // console.log('send', prompt)
      this.$nextTick(() => {
        this.$refs.sessionWinRef1.submit(prompt)
        this.$refs.sessionWinRef2.submit(prompt)
      })
    },
  },
  async beforeCreate() {
    this.initLoading = true;
    try {
      const modelDataRes = await getPromoteData('model/modelbasenew.json')
      const modelexperienceRes = await getPromoteData(`model/modelexperiencedeepseeknpu${document.documentElement.attributes["lang"].nodeValue == "zh-CN" ? '' : '_en'}.json`)
      const modelData = JSON.parse(modelDataRes.data);
      const modelexperienceData = JSON.parse(modelexperienceRes.data);
      this.experienceConfig = modelexperienceData;
      const modelMap = {}
      const modelIconImgMap = {
        'DeepSeek': 'DeepSeek',
        'Qwen': 'Qw',
        'ChatGLM': 'GLM',
      }
      for (let i = 0, iLen = modelData.length; i < iLen; i++) {
        const dataI = modelData[i];
        for (let j = 0, jLen = dataI.list.length; j < jLen; j++) {
          const dataJ = dataI.list[j];
          const category = dataJ.category;
          for (let k = 0, kLen = dataJ.models.length; k < kLen; k++) {
            const dataK = dataJ.models[k];
            if (dataK.experienceNpu && dataK.name != 'DeepSeek-Dynamic') {
              if (modelMap[category]) {
                modelMap[category].models.push(dataK)
              } else {
                modelMap[category] = {
                  category: category,
                  icon: modelexperienceData.imgList[modelIconImgMap[category]],
                  iconSize: category == 'DeepSeek' ? 30 : 27,
                  models: [dataK],
                }
              }
            }
          }
        }
      }
      const modelInfo = []
      for (let key in modelMap) {
        modelInfo.push(modelMap[key])
      }
      this.modelInfo = modelInfo;
    } catch (err) {
      console.log(err)
    } finally {
      this.initLoading = false;
    }
  },
  mounted() { },
  beforeDestroy() { },
};
</script>

<style scoped lang="less">
.op-btn {
  background-color: white;
  color: rgb(0, 102, 255);
  height: 40px;
  border-color: rgb(1, 145, 255);
  border-style: solid;
  border-width: 1px;
  border-radius: 5px;

  .btn-content {
    color: rgba(0, 102, 255, 1);
    font-size: 16px;
    font-family: PingFangSC;
  }

  &.is-disabled {
    border-color: rgba(16, 16, 16, 0.2);

    .btn-content {
      color: rgba(16, 16, 16, 0.5)
    }

    &:hover {
      background-color: white;
    }
  }
}

.content-box {
  height: 100%;
  width: 100%;
  display: flex;
  flex-direction: column;

  .main-body {
    flex: 1;
    height: 0;
    position: relative;
    margin-right: 16px;

    .content {
      height: calc(100% - 32px);
      margin-top: 16px;
      margin-bottom: 16px;
      overflow-y: auto;
      padding: 0 0px 10px 10px;
      scroll-behavior: smooth;

      .arena-container {
        height: 100%;
        position: relative;
        padding-left: 16px;

        .top {
          background-color: gold;
          height: calc(100% - 98px);
          display: flex;
          border: 1px solid rgba(16, 16, 16, 0.1);
          border-radius: 5px;
          background-color: rgb(255, 255, 255);

          .left {
            width: 50%;
            border-right: 1px solid rgba(16, 16, 16, 0.1);
            overflow-y: auto;
            position: relative;
          }

          .right {
            width: 50%;
            overflow-y: auto;
            position: relative;
          }
        }

        .bottom {
          margin-top: 18px;
          height: 80px;
          display: flex;
          align-items: center;
          justify-content: center;

          .send-box {
            display: flex;
            align-items: center;
            width: 900px;
            position: relative;

            .send-txt {
              height: 60px;
              flex: 1;
              font-size: 16px;
              font-weight: 400;
              line-height: 23px;
              font-style: normal;
              color: rgb(16, 16, 16);
              letter-spacing: 0px;
              padding: 7px 10px;
              background: rgb(255, 255, 255);
              border: 2px solid rgb(1, 145, 255);
              border-radius: 10px;
              padding-right: 100px;
              outline: none;
              resize: none;
              box-sizing: border-box;

              &::-webkit-scrollbar {
                width: 0;
              }
            }

            .btn {
              position: absolute;
              right: 0;
              top: 0;
              display: flex;
              align-items: center;

              &.btn-1 {
                color: rgba(0, 102, 255, 1);
                cursor: pointer;
                right: 12px;
                top: 20px;
                padding: 2px 4px;

                &.disabled-send {
                  cursor: not-allowed;
                }

                i {
                  font-size: 20px;
                  margin-right: 5px;
                }

                span {
                  font-size: 16px;
                }
              }

              &.btn-2 {}
            }

            &.disabled-send {
              .send-txt {
                color: rgb(164, 173, 179);
                background: rgba(16, 16, 16, 0.1);
                border-color: transparent;
              }

              .btn-2 {
                top: 10px;
                right: 10px;
                height: 40px;
                background: #3a3a3a;
                border-radius: 4px;
                color: rgb(255, 255, 255);
                padding: 0 10px;
              }
            }
          }
        }
      }
    }
  }
}

@media (max-width: 768px) {
  .content-box {
    .main-body {
      .content {
        margin-bottom: 0;
        padding-bottom: 0;

        .arena-container {
          padding-left: 6px;

          .top {
            flex-direction: column;
            height: calc(100% - 130px);

            .left {
              width: 100%;
              height: 50%;
              border-right: none;
              border-bottom: 1px solid rgba(16, 16, 16, 0.1);
            }

            .right {
              width: 100%;
              height: 50%;
            }
          }

          .bottom {
            margin-top: 10px;
            height: 120px;

            .send-box {
              .send-txt {
                height: 110px;
                padding-right: 10px;
              }

              .btn {
                &.btn-1 {
                  height: 24px;
                  bottom: 12px;
                  top: auto;
                }
              }


              &.disabled-send {
                .btn-2 {
                  top: auto;
                  bottom: 10px;
                }
              }
            }
          }
        }
      }
    }
  }
}
</style>
