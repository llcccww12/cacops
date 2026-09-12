<template>
    <div style="margin-top:32px;">
        <div class="__mobile-tip" >
          <img style="width: 132px;height: 97px;" src="/img/model/pc-view.png" alt="">
          <div style="margin-top: 2rem;">{{$t('useInPcWeb')}}</div>
        </div>
        <create-form :title="title" v-loading="loading" :active="active" class="__content-box">
          <div v-if="active===0">
            <el-form ref="form" :model="form" :rules="rules" hide-required-asterisk label-width="220px">
              <el-form-item :label="$t('modelSquare.model_name')" prop="name" style="margin-bottom: 0;">
                <el-input v-model="form.name" maxlength="100"></el-input>
                <span style="font-size: 12px;color: #888;line-height: 1;margin-top: 0.5em;display: inline-block;">
                    {{ $t('datasetObj.dataset_name_tooltips') }}
                </span>
              </el-form-item>
              <el-form-item :label="$t('modelSquare.model_zh_name')" prop="alias" style="margin-bottom: 0;">
                  <el-input v-model="form.alias" maxlength="100"></el-input>
                  <span style="font-size: 12px;color: #888;line-height: 1;margin-top: 0.5em;display: inline-block;">
                      {{ $t('datasetObj.dataset_name_cn_tooltips') }}
                  </span>
              </el-form-item>
              <el-form-item :label="$t('datasetObj.dataset_owner')" prop="owner_id">
                  <el-select v-model="form.owner_id" :placeholder="$t('datasetObj.select_category')" style="width: 60%;">
                      <el-option v-for="item in owners" :key="item.ID" :value="item.ID" :label="item.Name"></el-option>
                  </el-select>
              </el-form-item>
              <el-form-item label="" prop="is_private">
                  <el-radio-group v-model="form.is_private">
                      <el-radio :label="false">{{$t('modelManage.modelAccessPublic')}}</el-radio>
                      <el-radio :label="true">{{$t('modelManage.modelAccessPrivate')}}</el-radio>
                  </el-radio-group>
              </el-form-item>
              <el-form-item :label="$t('modelManage.modelEngine')" prop="engine">
                  <el-select v-model="form.engine" style="width: 60%;">
                      <el-option v-for="item in engineList" :key="item.k" :value="item.k" :label="item.v"></el-option>
                  </el-select>
              </el-form-item>
              <el-form-item :label="$t('modelManage.license')" prop="license">
                  <el-select v-model="form.license" clearable filterable style="width: 60%;">
                      <el-option v-for="item in licenseList" :key="item.id" :label="item.name" :value="item.id"></el-option>
                  </el-select>
              </el-form-item>
              <el-form-item :label="$t('modelManage.modelLabel')" prop="label">
                  <el-input v-model="form.label" @input="labelInput" :placeholder="$t('modelManage.modelLabelInputTips')" maxlength="255"></el-input>
              </el-form-item>
              <el-form-item class="btn-wrap"> 
                  <el-button round class="cancel" type="info" @click="cancelModel()">{{$t('cancel')}}</el-button>
                  <el-button round class="submit" type="primary" @click="createModel()">{{$t('cloudbrainObj.nextStep')}}</el-button>
              </el-form-item>
            </el-form> 
          </div>
          <div v-if="active===1" >
            <FileUpload :modelId="modelID" subjectType="2" @cancel="cancelUpload" @uploadFinish="uploadFinish"></FileUpload>
          </div>
        </create-form>
    </div>
</template>
<script >
import CreateForm from '../components/CreateForm.vue'; 
import FileUpload from '../components/FileUpload.vue';
import { getAvailableUsers } from '~/apis/modules/dataset';
import { createModels, getModelLicenseList } from '~/apis/modules/modelmanage';
import { getUrlSearchParams } from '~/utils';
import { MODEL_ENGINES } from '~/const'
const {AppSubUrl} = window.config;
const MAX_LABEL_COUNT = 5;
export default {
    components: {CreateForm,FileUpload },
    data() {
      return {
          owners: [],
          orgId: 0,
          cancelUrl:'',
          form: {
              name: '',
              alias: '',
              aimodel_type: 1,
              is_private: false,
              engine: 0,
              label: '',
              license: '',
              owner_id: '',
          },
          loading:true,
          engineList: MODEL_ENGINES,
          licenseList: [],
          label: "",
          rules: {
              name: [
                  { required: true, message: this.$t('datasetObj.dataset_name_required3'), trigger: "blur" },
                  {
                      validator: (rule, value, callback) => {
                          if (/^[a-zA-Z0-9][a-zA-Z0-9-_.]{0,98}[a-zA-Z0-9]$/.test(value) == false) {
                              callback(new Error(this.$t('datasetObj.dataset_name_required2')));
                          } else {
                              callback();
                          }
                      },
                      trigger: "blur",
                  },
              ],
              alias: [
                  // { required: true, message: this.$t('datasetObj.dataset_name_required3'), trigger: "blur" },
                  {
                      validator: (rule, value, callback) => {
                          if (/^[\u4e00-\u9fa5a-zA-Z0-9-_.]{0,100}$/.test(value) == false) {
                              callback(new Error(this.$t('datasetObj.dataset_name_required2')));
                          } else {
                              callback();
                          }
                      },
                      trigger: "blur",
                  },
              ],
              owner_id: [
                  { required: true},
              ],
              engine: [
                  { required: true},
              ],

          },
          active: 0,
          title: [this.$t('modelManage.createNewModel'), this.$t('modelManage.uploadModelFiles')],
          modelID: '',
          modelObj: {},

      }
    },
    methods: {
      labelInput() {
        const hasEndSpace = this.form.label[this.form.label.length - 1] == ' ';
        const list = this.form.label.trim().split(' ').filter(label => label != '');
        this.form.label = list.slice(0, MAX_LABEL_COUNT).join(' ') + (hasEndSpace && list.length < MAX_LABEL_COUNT ? ' ' : '');
      },
      createModel() {
        this.$refs['form'].validate((valid) => {
            if (valid) {
                this.loading = true
                this.postCreateModel()
            } else {
                return false
            }
        })
      },
      async postCreateModel() {
          try {
              const response = await createModels(this.form)
              console.log(response)
              this.loading = false
              if (response.data.code === 0) {
                  this.active = 1
                  this.modelID = response.data.data.id
                  this.modelObj = response.data.data
              } else {
                  this.$message.error(response.data.msg);
              }
          }catch (error) {
              this.loading = false
              this.$message.error(error);
          }
      },
      cancelModel() {
          if (this.cancelUrl) {
              window.location.href = this.cancelUrl;
          } else {
              const referrer = document.referrer || '';
              if (referrer.indexOf('guide/create_model') > -1) {
                  window.location.href = '/dashboard';
              } else {
                  window.history.back();
              }
          }
      },
      async getUsersList() {
        try {
            const response = await getAvailableUsers({org:this.orgId},'aimodel')
            console.log(response)
            const res = response.data
            if (res.code === 0) {
                this.owners = res.data.users || []
                if (this.owners.length === 0) {
                    this.loading = false
                    return
                }
                if(this.orgId){
                    // 保证 owner_id 一定存在于 owners 中，否则 el-select 会回退显示原始 id
                    const orgOwner = this.owners.find(item => item.ID === +this.orgId)
                    this.form.owner_id = orgOwner ? orgOwner.ID : this.owners[0].ID
                }else{
                    this.form.owner_id = this.owners[0].ID
                }
            } else {
                this.$message.error(response.data.msg);
            }
            this.loading = false
        }catch (error) {
            this.loading = false
            this.$message.error(error);
        }
      },
      cancelUpload(){
        location.href = `/models/detail/${this.modelObj.owner_name}/${this.modelObj.name}`
      },
      uploadFinish(){
        window.setTimeout(() => {
            location.href = `/models/detail/${this.modelObj.owner_name}/${this.modelObj.name}?tab=files`
        }, 1000);
      }
    },
  mounted() { },
  beforeMount() {
      const urlParams = getUrlSearchParams();
      if (urlParams.backurl) {
          this.cancelUrl = urlParams.backurl;
      }
      if(urlParams.org){
          this.orgId = urlParams.org
          console.log(urlParams.org)
      }
      this.getUsersList()
      getModelLicenseList().then(res => {
        res = res.data;
        try {
          const license = JSON.parse(res) || [];
          this.licenseList = license;
          if (license.length) {
            this.form.license = license[0].id;
          }
        } catch (err) {
          console.log(err);
        }
      }).catch(err => {
        console.log(err);
      });
  },
}
</script>

<style lang="less" scoped>
::v-deep .is-required .el-form-item__label::after {
  content: "*";
  color: #ff0000;
  margin-left: 4px;
}

::v-deep .el-form-item__label{
  color: #101010;
}
::v-deep .el-radio-button__orig-radio:checked + .el-radio-button__inner{
  background-color: #f9f9f9;
  border-color: #3291f8;
  color: #3291f8;
}
::v-deep .el-button--primary{
  background-color: #007aff;
}
::v-deep .el-button--info{
  color: #020004;
  background-color: #c2c7cc;
  border-color: #c2c7cc;
}
// 可定义局部公共样式
.steps {
  width: 100%;
  padding: 0 20px 0 120px;
  
  height: 35px;
  ::v-deep .el-step {
    height: 100%;
    // 设置图标和步骤条的行高
    .el-step__head {
      line-height: 35px;
    }
    .el-step__icon{
      width: 48px;
      height: 48px;
      z-index: 99;
      font-size: 18px;
    }
    // 步骤条
    .el-step__line {
      background-color: rgba(0, 0, 0, 0.15);
      top: 50%;
      left:110px;
      height: 1px;
    }
    .el-step__head.is-process {
      color: #0066ff;
      border-color: #0066ff;
      .el-step__icon{
        border: 1px solid #6195f7;
        color: #6195f7;
      }
    }
      // title样式
    .el-step__title {
      z-index: 66;
      position: absolute;
      top: 5px;
      left: calc(10%);
      font-size: 14px;
      background-color: #fff;
      z-index: 66;
      color: #404040;
    }
    .el-step__title.is-process{
      color:#3291f8
    }
    // 已完成步骤条的边框色和字体颜色
    .el-step__head.is-finish {
      color: #0066ff;
      border-color: #0066ff;
    }
    .el-step__title.is-success {
      // 已完成步骤的title
      font-weight: 700;
      color: #3291f8;
    }
    // .el-step__title.is-process {
    //   //未完成步骤的title
    //   color: #3291f8;
    // }
    

    // 已完成图标背景色
    // .el-step__icon {
    //   width: 48px;
    //   height: 48px;
    //   font-size: 18px;
      
    //   z-index: 99;
    //   // 已完成图标字体颜色
    //   .el-step__icon-inner {
    //     font-weight: unset !important;
        
    //   }
    // }
    // 未完成图标背景色
    // .is-process .el-step__icon.is-text {
    //   z-index: 99;
    //   // background: #6195f7;
    //   // 未完成图标字体颜色
    //   .el-step__icon-inner {
    //     color: #0066ff;
    //   }
    // }
    // // title样式
    // .el-step__title {
    //   z-index: 66;
    //   position: absolute;
    //   top: 0;
    //   left: calc(2%);
    //   width: 150px;
    //   font-size: 14px;
    //   background-color: #fff;
    //   z-index: 66;
    //   color: #404040;
    // }
    // .el-step__title.is-process{
    //   // 防止最后一个title会加粗
    //   font-weight: normal !important;
    // }
    
  }
  // 第一个步骤
  ::v-deep .el-step:first-child {
    flex-basis: 40% !important;
    .el-step__head.is-process {
      padding-left: 10px !important;
    }
    .el-step__head.is-success {
      padding-left: 10px !important;
    }
    // .el-step__line {
    //   width: 100%;
    //   margin-left: 25% !important;
    // }
    
    .el-step__title {
      padding-left: 26px !important;
    }

  }
  // 第二个步骤
  ::v-deep .el-step:nth-child(2) {
    flex-basis: 40% !important;
    .el-step__title {
      padding-left: 16px !important;
    }
    // .el-step__line {
    //   width: 100%;
    //   margin-left: 28% !important;
    // }
    // .el-step__title {
    //   width: 90px;
    //   padding-left: 20px !important;
    // }
    // .el-step__description {
    //   margin-left: 28px !important;
    // }
  }
  // 第三个步骤
  ::v-deep .el-step:last-child {
    flex-basis: 20% !important;
    .el-step__title {
      padding-left: 38px !important;
      z-index: 66;
    }
    
  }
}
 </style>
