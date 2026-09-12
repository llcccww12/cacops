<template>
    <div>
        <div class="ui container" v-if="errorMsgBoxShow" style="width: 1050px !important;">
            <div class="err-msg-box">
                <p>{{ errorMsg }}</p>
            </div>
        </div>
        <div class="ui container area">
            <div class="area-title">{{$t('modelSquare.modelChatTask')}}</div>
            <div class="area-content">
                <div class="ui container">
                   <div style="padding: 5rem 5rem 2rem 0;">
                        <div class="form-row">
                            <div class="left-area">
                                <div class="title">
                                    <span class="required">{{ $t('modelManage.modelName') }}</span>
                                </div>
                                <div class="content">
                                    <el-input class="field-input" v-model="modelName" readonly></el-input>
                                </div>
                            </div>
                            <div class="right-area"></div>
                        </div>
                        <!-- <FormTop ref="formTopRef" :repoOwnerName="repoOwnerName" :repoName="repoName" :configs="pageCfg" :queueNum="queueNum"></FormTop> -->
                        <div class="form-row">
                            <div class="left-area">
                                <div class="title">
                                    <span class="required">{{ $t('resourcesManagement.computeResource') }}</span>
                                </div>
                                <div class="content">
                                    <div class="list">
                                        <div class="item focus">
                                            <i class="icon ri-archive-drawer-line"></i>
                                            <span>{{$t('computeResourceTitle.GPU')}}</span>
                                        </div>
                                    </div>
                                </div>
                            </div>
                        </div>
                        <div class="form-row tips-c">
                            <div class="left-area">
                                <div class="title"></div>
                                <div class="content">
                                    <div class="tips tips-1">
                                        <i class="ri-error-warning-line"></i>
                                        <span>
                                        {{ $t('cloudbrainObj.waitCountStart') }} <span>{{ queueNum }}</span> {{ $t('cloudbrainObj.waitCountEnd') }}
                                        </span>
                                    </div>
                                </div>
                            </div>
                        </div>
                        
                        <SpecSelect ref="specRef" v-model="state.spec" :required="true" :configs="specConfigs" networkType="all" :loading="loading"></SpecSelect>
                        <TaskName ref="taskNameRef" v-model="state.taskName" :required="true" :userName="loginName"></TaskName>
                        <div class="form-row">
                            <div class="left-area">
                                <div class="title"></div>
                                <div class="content">
                                <el-button type="primary"
                                    :disabled="maskLoading ||  !specConfigs.specs['all']"
                                    size="default" class="submit-btn" @click="submit">{{ $t('cloudbrainObj.createTask')
                                    }}</el-button>
                                <el-button class="cancel-btn" size="default" @click="cancel">{{ $t('cancel') }}</el-button>
                                </div>
                            </div>
                        </div>
                   </div>
                </div>
            </div>
            <LoadingMask :loading="maskLoading" :content="maskLoadingContent"></LoadingMask>
        </div>
        
    </div>
</template>
<script>
import LoadingMask from '../../../modelbase/components/cloudbrain/LoadingMask.vue';
import { getAiTaskPrepareInfo, createAiTask } from '~/apis/modules/cloudbrain';
import FormTop from '~/components/cloudbrain/FormTop.vue';
import SpecSelect from '~/components/cloudbrain/SpecSelect.vue';
import TaskName from '~/components/cloudbrain/TaskName.vue';
import { getPromoteData } from '~/apis/modules/common';

export default {
    data(){
        return {
            loginName: '',
            modelName:'',
            maskLoading: false,
            maskLoadingContent: '',
            errorMsgBoxShow: false,
            errorMsg: '',
            state: {
                spec: '',
                taskName:'',
            },
            specConfigs: {
                specs: {},
                blance: 0,
                showPoint: false,
            },
            DepModelInfo: {},
            queueNum: 1,
            computeSource: '',
            loading:false
        }
    },
    components: { LoadingMask,SpecSelect,TaskName,FormTop },
    methods: {
        submit() {
            if (this.maskLoading) return;
            let canSubmit = true;
            for (let key in this.state) {
                if (this.$refs[key + 'Ref']) {
                    if (!this.$refs[key + 'Ref'].check()) {
                        canSubmit = false
                    }
                }
            }
            if (!canSubmit) return;

            const subObj = {
                repoOwnerName: this.DepModelInfo[0].ops.repo_owner_name,
                repoName: this.DepModelInfo[0].ops.repo_name,
                job_type: 'MODELEXPERIENCE',
                cluster: 'C2Net',
                compute_source: this.specConfigs.specs.all[0].compute_resource,
            }
            subObj['display_job_name'] = this.state.taskName
            subObj['app_name'] = this.DepModelInfo[0].name
            subObj['description'] = ''
            subObj['branch_name'] = 'master'
            subObj['pretrain_model_id_str'] = this.DepModelInfo[0].ops.model_id //'1c943538-1b83-4fcc-8b72-5272375e7103' //this.DepModelInfo[0].ops.model_id
            subObj['boot_file'] = this.DepModelInfo[0].ops.boot_file //api.py
            subObj['has_internet'] = this.specConfigs.specs.all[0].has_internet
            subObj['spec_id'] = +this.state.spec
            subObj['work_server_number'] = 1
            subObj['image_url'] = this.DepModelInfo[0].ops.image_url
            subObj['label_names'] = this.DepModelInfo[0].type
            this.maskLoadingContent = this.$t('cloudbrainObj.taskPrepareTips');
            this.maskLoading = true;
            this.errorMsg = '';
            this.errorMsgBoxShow = false;
            createAiTask(subObj,{model_experience:true}).then(res => {
                const data = res.data;
                if (data.code == 0) {
                    window.location.href = '/cloudbrains'
                } else {
                    this.maskLoading = false;
                    this.errorMsg = data.msg;
                    this.errorMsgBoxShow = true;
                    document.querySelector('html').scrollTo({ top: 0, behavior: 'smooth' });
                    document.querySelector('body').scrollTo({ top: 0, behavior: 'smooth' });
                }
            }).catch(err => {
                this.maskLoading = false;
                this.$message.error(err)
            });
            
        },
        cancel() {
            if (window.history.length === 1) {
                location.href="/modelbase"
            } else {
                window.history.back()
            }
        },
    },
    mounted(){
        const urlParams = new URLSearchParams(location.search)
        if(urlParams.has('model')){
            this.modelName = urlParams.get('model')
        }
        getPromoteData('model/modelexperience.json').then(res => {
            try {
                const data = JSON.parse(res.data);
                this.DepModelInfo = data.filter((item) => {
                    return item.name === this.modelName
                })
            }   catch (err) {
                this.loading = false
                this.$message.error(err)
            }
            if(this.DepModelInfo.length){
                this.loading = true
                getAiTaskPrepareInfo({
                    repoOwnerName: this.DepModelInfo[0].ops.repo_owner_name,
                    repoName: this.DepModelInfo[0].ops.repo_name,
                    jobType: 'MODELEXPERIENCE',
                    clusterType: 'C2Net',
                    computeSource: 'GPU',
                    model_experience:true
                }).then((res)=>{
                    res = res.data
                    this.loading = false
                    if(res.code===0){
                        const data = res.data;
                        this.queueNum = data.wait_count || 1;
                        this.specConfigs.showPoint = data.pay_switch;
                        this.specConfigs.blance = data.point_account ? data.point_account.balance : 0;
                        this.specConfigs.specs = { 'all': data.specs.all }
                        this.state.spec =  this.specConfigs.specs['all'][0].id.toString() //specsInfo[0].id.toString()
                        this.state.taskName = data.display_job_name;

                    }
                })
            }
        })
    },
    beforeCreate() {
        
    },
    beforeMount() {
        const isLogin = !!document.querySelector('meta[name="_uid"]');
        if (isLogin) {
            this.loginName = document.querySelector('meta[name="_uid"]').getAttribute('content-ext')
        }
    },
}
</script>
<style scoped lang="less">
@import '~/components/cloudbrain/cloudbrain.less';
.err-msg-box {
  display: flex;
  align-items: center;
  justify-content: center;
  background-color: #fff6f6;
  color: #9f3a38;
  box-shadow: 0 0 0 1px #e0b4b4 inset, 0 0 0 0 transparent;
  margin: 1em 0;
  padding: 1em 1.5em;
  border-radius: 0.28571429rem;
  transition: opacity .1s ease, color .1s ease, background .1s ease, box-shadow .1s ease, -webkit-box-shadow .1s ease;

  >p {
    opacity: .85;
  }
}
.area {
    width: 1050px !important;
    margin-top: 40px;
    
    .area-title {
        height: 45px;
        border-color: rgb(212, 212, 213);
        border-width: 1px;
        border-style: solid;
        border-radius: 5px 5px 0px 0px;
        font-size: 14px;
        background: rgb(240, 240, 240);
        line-height: 45px;
        padding-left: 15px;
        font-weight: 550;
        font-size: 16px;
        color: rgb(16, 16, 16);
    }
    .area-content {
        border-color: rgb(212, 212, 213);
        border-width: 1px;
        border-style: solid;
        margin-top: -1px;
        .area-main-wrap {
            display: flex;
            flex-direction: column;
            align-items: center;
            .area-main-btn{
                display: flex;
                justify-content: center;
                margin-bottom: 2rem;
            }
        }
    }
}
.list {
  display: flex;
  align-items: center;
  flex-wrap: wrap;

  .item {
    display: flex;
    align-items: center;
    justify-content: center;
    font-size: 12px;
    color: rgba(0, 0, 0, .87);
    border: 1px solid rgba(34, 36, 38, .15);
    margin-left: -1px;
    height: 38px;
    padding: 0 12px;
    margin-bottom: 5px;

    i {
      margin-top: -7px;
      font-size: 14px;
    }

    &.focus {
      color: #0087f5;
      border-color: #0087f5;
      border-left: 1px solid #0087f5;
    }

    &:first-child {
      border-top-left-radius: 0.28571429rem;
      border-bottom-left-radius: 0.28571429rem;
      border-left: 1px solid rgba(34, 36, 38, .15);

      &.focus {
        border-color: #0087f5;
      }
    }

    &:last-child {
      border-top-right-radius: 0.28571429rem;
      border-bottom-right-radius: 0.28571429rem;
    }

    &:hover:not(.focus) {
      background: rgba(0, 0, 0, .03);
    }
  }
}
.tips-c {
  margin-top: -20px;

  .tips {
    display: flex;
    font-size: 12px;
    align-items: center;

    i {
      color: #f2711c;
      margin-right: 5px;
      font-size: 14px;
    }
  }

  .tips-1 {
    span {
      color: #f2711c;
    }
  }

  .tips-2 {
    span {
      color: #888;
    }
  }
}
</style>