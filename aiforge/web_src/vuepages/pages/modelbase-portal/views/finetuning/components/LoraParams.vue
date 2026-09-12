<template>
<div class="train-params">
    <div class="train-params-content">
        <div class="title-wrap">
            <div class="title-first">
                <span class="number">1</span>
                <span>{{$t('cloudbrainObj.paramsSetting')}}</span>
            </div>
            <div class="params-dialog" @click="dialogShow=true">
                <span>{{$t('modelFinetune.ProfessionalParameters')}}</span>
                <i class="ri-file-copy-line"></i>
            </div>
        </div>
        <div class="params-item">
            <div style="line-height: 40px;">{{$t('modelFinetune.useBasicModel')}}：</div>
            <el-select v-model="modeOptions[0].value" style="width: 100%;">
                <el-option v-for="item in modeOptions" :key="item.value" :label="item.label" :value="item.value" />
            </el-select>
        </div>
        <div class="params-item">
            <div class="slider-wrap">
                <span>{{$t('modelFinetune.Repeat')}}：</span>
                <el-slider v-model="sections[1].params[0].value" show-input :show-input-controls="false" :min="1" style="width: 60%;padding: 0 6px"></el-slider>
                <span class="label">Repeat</span>
            </div>
            <div class="slider-wrap">
                <span>{{$t('modelFinetune.Epoch')}}：</span>
                <el-slider v-model="sections[0].params[0].value" show-input :show-input-controls="false" :min="1" style="width: 60%;padding: 0 6px;"></el-slider>
                <span class="label">Epoch</span>
            </div>
        </div>
        <div class="params-item total-step">
            <span>{{$t('modelFinetune.totalSteps')}}: </span>
            <span>{{stepTotal ? stepTotal : $t('modelFinetune.uploadImgCalc') }}</span>
        </div>
    </div>
    <div class="train-params-prompt">
        <div style="margin: 16px 0;">{{$t('modelFinetune.Prompt')}}</div>
        <div style="flex-grow: 1;margin-bottom: 24px;">
            <textarea v-model="sections[2].params.at(-1).value" :placeholder="$t('modelFinetune.loraPromptPlace')" class="prompt-input"></textarea>
        </div>
    </div>
    <el-dialog
        :visible.sync="dialogShow"
        custom-class="setting-container"
        :close-on-click-modal="false"
        :close-on-press-escape="false"
        @closed="closed"
        @opened="opened"
        @open="open"
        >
        <div slot="title">
            <div style="display: flex;color: #101010;">
                <i style="font-size: 24px;transform: rotate(90deg);" class="ri-sound-module-line"></i>
                <span style="font-size: 16px;margin-left: 12px;">{{$t('modelFinetune.ProfessionalSetting')}}</span>
            </div>
		</div>
        <div class="setting-wrap">
            <div class="setting-menu">
                <div class="menu-wrap">
                    <div 
                        v-for="(item, index) in sections" 
                        :key="index"
                        class="menu-item"
                        :class="{ 'is-checked': activeSection === item.title }" 
                        @click.prevent="scrollToSection(item.title)"
                    >
                        {{ item.title }}
                    </div>
                </div>
            </div>
            <div class="setting-form" ref="contentForm">
                <div class="content-group" 
                    v-for="(section, index) in sections"
                    :key="index"
                    :id="section.title"
                >
                    <div class="title" ref="sectionTitles" :data-id="section.title">{{section.title}}</div>
                    <div class="content">
                        <div 
                            v-for="(param, i) in section.params"
                            :key="i"
                        >
                            <div class="_item">
                                <div class="label">
                                    <div class="label-en">{{ param.name_en }}</div>
                                    <div class="label-zh">{{ param.name_zh }}</div>
                                </div>
                                <div class="value">
                                    <el-slider v-if="param.type==='slider'" v-model="param.value" show-input :show-tooltip="false" 
                                    :show-input-controls="false" input-size="medium" style="width: 100%;"
                                    :min="param.min" :max="param.max" :step="param.step"></el-slider>
                                    <el-select v-model="param.value" style="width: 100%;" v-else-if="param.type==='select'" :disabled="param.disabled" size="medium">
                                        <el-option
                                        v-for="item in param.options"
                                        :key="item.value"
                                        :label="item.label"
                                        :value="item.value">
                                        </el-option>
                                    </el-select>
                                    <el-input  v-else-if="param.type==='input'" v-model="param.value" size="medium" :disabled="param.disabled"></el-input>
                                    <el-input-number  v-else-if="param.type==='number'" v-model="param.value" size="medium" style="width: 100%;"
                                        controls-position="right" :disabled="param.disabled" :min="param.min" :max="param.max" :step="param.step"></el-input-number>
                                    <el-switch v-else-if="param.type==='radio'" v-model="param.value" :disabled="param.disabled"></el-switch>
                                    <el-input v-else-if="param.type==='texteare'" v-model="param.value" type="textarea" rows="3" :disabled="param.disabled"></el-input>
                                </div>
                            </div>
                        </div>
                    </div>
                </div>
            </div>
        </div>
    </el-dialog>
</div>
</template>
<script>
export default {
  name: 'LoraParams',
  components: {
     
  },
  mixins: [],
  props: {
    fileLength: { type: Number, default: 0 },
    sections: { type: Array, default: () => [] },
    modelName: { type: String, default: "", },
  },
  data() {
    return {
        dialogShow: false,
        baseModel: 'F1_dev_fp8',
        // modeOptions: [{ label: 'F1_dev_fp8', value: 'F1_dev_fp8' }, { label: 'F1_dev', value: 'F1_dev' },],
        
        state:{
            TrainConfig:{},
            DatasetConfig:{},
            PromptsConfig:{}
        },
        stepsValue: 0,
        
        activeSection: '',
        isScrollingProgrammatically: true,
    }
  },
  computed: {
    stepTotal() {
        if (this.stepsValue) {
            return Math.ceil(this.stepsValue * this.sections[1].params[0].value / this.sections[1].params[1].value )* this.sections[0].params[0].value
        }
    },
    modeOptions() {
        return [{label:this.modelName,value:this.modelName}]
    }
  },
  watch: {
    fileLength(newVal) {
      if (newVal === 0) {
        this.stepsValue = 0;
      } else {
        this.stepsValue = newVal
      }
    }, 
  },
  methods: {
    scrollToSection(id) {
        this.isScrollingProgrammatically = true
        const section = document.getElementById(id);
        if (section) {
            this.activeSection = id;
            section.scrollIntoView({ behavior: 'smooth' });
            setTimeout(() => {
                this.isScrollingProgrammatically = false; // Reset the flag after scrolling completes
            }, 500);
        }
    },
    closed() { },
    open() {
        this.activeSection = this.sections[0].title
    },
    opened() {
        this.setupObserver()
        this.$refs.sectionTitles.forEach((section, index) => { 
            this.observer.observe(section)
        })
    },
    setupObserver() {
        this.observer = new IntersectionObserver((entries) => {
            entries.forEach((entry) => {
                if (entry.isIntersecting && !this.isScrollingProgrammatically) {
                  const sectionId = entry.target.getAttribute('data-id');
                  if (this.activeSection !== sectionId) {
                     this.activeSection = sectionId;
                  }
                }
            });
        }, {
            root: null,
            rootMargin: '0px',
            threshold: 0.25,
        });
    },
    submitParams() {
        const data = {
            TrainConfig:{},
            DatasetConfig:{resolution:[1024,1024]},
            PromptsConfig:{negative:""}
        }
        this.sections.forEach((section) => {
            section.params.forEach((item) => {
                if (item.data_flag) {
                    data.DatasetConfig[item.name_train] = item.value
                } else if (item.prompts_flag) {
                    if (item.name_en === 'resolution') {
                        data.PromptsConfig[item.name_train] = item.value.split(",").map(strNum => +strNum);
                    } else {
                        data.PromptsConfig[item.name_train] = item.value
                    }
                    
                } else {
                    data.TrainConfig[item.name_train] = item.value
                }
            })
        })
        return data 
    },
  },
  mounted() {
    
  },
  beforeDestroy() {
    
  }
};
</script>
<style lang='less' scoped>
.number{
    display: flex;
    align-items: center;
    justify-content: center;
    width: 20px;
    height: 20px;
    margin-right: 8px;
    background: #3162ff;
    border-radius: 10px;
    color: #fff;
    font-size: 14px
}
.train-params{
    border-right: 1px solid rgba(157,197,226,0.4);
    flex-shrink: 0;
    box-sizing: content-box;
    width: 446px;
    height: 100%;
    overflow-y: auto;
    background-color: #fff;
    
    font-family: PingFang SC;
    display: flex;
    flex-direction: column;
    .train-params-content{
        padding: 18px 24px 0 18px;
        display: flex;
        flex-direction: column;
        .title-wrap{
            display: flex;
            align-items: center;
            font-weight: 500;
            justify-content: space-between;
            .title-first{
                display: flex;
                align-items: center;
                
            }
            .params-dialog{
                display: flex;
                align-items: center;
                cursor: pointer;
            }
        }
        .params-item{
            margin-top: 24px;
            /deep/ .el-input__inner {
                color: #191919;
                background: #f2f5f9;
                border: none;
                border-radius: 8px;
            }
            .slider-wrap{
                &:first-child{
                    margin-bottom: 4px;
                }
                display: flex;
                align-items: center;
                .label{
                    color: #888888;
                    margin-left: 4px;
                }
                /deep/ .el-slider__input{
                    width:60px;
                    float: left;
                }
                /deep/ .el-slider__runway.show-input{
                    margin-left: 70px;
                    margin-right: 10px;
                }
                /deep/ .el-slider__button { // 拖动的滑块的样式
                    width: 10px;
                    height: 10px;
                }
            }
        }
        .total-step{
            height: 40px;
            line-height: 40px;
            border-radius: 4px;
            background-color: rgba(245,245,246,1);
            color: rgba(136,136,136,1);
            padding-left: 12px;
        }
    }
    .train-params-prompt{
        flex-grow: 1;
        padding: 16px;
        display: flex;
        flex-direction: column;
        .prompt-input{
            box-sizing: border-box;
            padding: 12px;
            border-radius: 12px;
            height: 100%;
            width: 100%; 
        }
    }
    /deep/ .setting-container{
        width: 784px;
        height: 570px;
        margin-top: 20vh !important;
        border-radius: 12px;
        .setting-wrap{
            display: flex;
            gap: 52px;
            justify-content: space-between;
            height: 528px;
            margin: -34px -20px;
            padding: 4px 0 20px 0;
            .setting-menu{
                position: sticky;
                top: 8px;
                flex-shrink: 0;
                width: 169px;
                overflow: hidden;
                .menu-wrap{
                    position: sticky;
                    top: 8px;
                    flex-shrink: 0;
                    width: 159px;
                    padding-left: 24px;
                    overflow: hidden;
                    .menu-item{
                        position: relative;
                        z-index: 1;
                        display: flex;
                        align-items: center;
                        justify-content: flex-start;
                        width: 135px;
                        height: 44px;
                        margin-top: 12px;
                        padding-left: 20px;
                        color: #999;
                        font-weight: 400;
                        font-size: 14px;
                        cursor: pointer;
                        -webkit-user-select: none;
                        -moz-user-select: none;
                        user-select: none;
                        &.is-checked{
                            position: relative;
                            color: #001ac7;
                            background-color: #f5f5f6;
                            border-radius: 32px;
                        }
                    }
                    
                }
            }
            .setting-form{
                flex-grow: 1;
                height: 100%;
                padding-right: 32px;
                overflow-y: auto;
                .content-group{
                    padding-bottom: 44px;
                    .title{
                        padding-top: 16px;
                        color: #191919;
                        font-weight: 500;
                        font-size: 18px;
                        line-height: 25px;
                    }
                    .content{
                        ._item{
                            display: flex;
                            align-items: flex-start;
                            margin-top: 20px;
                            .label{
                                display: flex;
                                flex-direction: column;
                                flex-shrink: 0;
                                justify-content: center;
                                width: 233px;
                                color: #999;
                                font-weight: 400;
                                .label-en{
                                    font-size: 12px;
                                    line-height: 17px;
                                }
                                .label-en{
                                    color: #191919;
                                    font-size: 14px;
                                    line-height: 20px;
                                    font-weight: 400;
                                }
                            }
                            .value{
                                display: flex;
                                flex-grow: 1;
                                justify-content: flex-end;
                                margin-top: 4px;
                            }
                            .el-slider__input{
                                width:80px
                            }
                            .el-slider__runway.show-input{
                                margin-right: 90px;
                            }
                            .el-slider__button { // 拖动的滑块的样式
                                width: 12px;
                                height: 12px;
                            }
                            .el-input__inner{
                                background: #f2f5f9;
                                border-radius: 8px;
                            }
                        }
                    }
                }
            }
        }
        
    }
}
::-webkit-scrollbar {
    width: 0; /* 尝试隐藏滚动条，但效果可能因浏览器而异 */
    background: transparent; /* 设置滚动条背景为透明 */
}
</style>