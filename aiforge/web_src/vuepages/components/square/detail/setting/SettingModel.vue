<template>
  <div class="ui container setting-wrap">
    <BaseTitle :title="$t('modelManage.modelBaseInfo')" v-loading="baseLoading">
        <el-form size="medium" class="form-box" ref="form" :model="form" :rules="rules">
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
                <el-select v-model="owners[0].name" disabled style="width: 60%;">
                    <el-option v-for="item in owners" :key="item.name" :value="item.name" :label="item.name"></el-option>
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
                <el-button round class="submit" style="background: rgb(0, 102, 255);" type="primary" @click="editBaseInfo()">{{$t('confirm')}}</el-button>
            </el-form-item>
        </el-form>
    </BaseTitle>

    <Access :id="dataObj.id" :title="$t('modelManage.modelCollaborator')" :placeholder="$t('datasetObj.searchUsers')" 
        :btnText="$t('datasetObj.addCollaborators')" type="Collaborators" filesType="aimodel"></Access>
    <Access v-if="dataObj.Owner.IsOrganization" :id="dataObj.id" :title="$t('modelManage.managementTeam')" filesType="aimodel"
        :placeholder="$t('datasetObj.searchTeams')" :btnText="$t('datasetObj.addTeam')" type="Teams" :org="dataObj.Owner.Name"></Access>
    <BaseTitle :title="$t('datasetObj.dangerousOperationZone')" class="content-del" v-if="dataObj.can_delete">
        <div class="main-box">
            <div class="text-wrap">
                <p style="font-weight: 700">{{$t('modelObj.deleteThisModel')}}</p>
                <p>{{$t('modelObj.deleteThisModelTips')}}</p>
            </div>
            <el-button size="medium" round plain class="del-btn" @click="showModal">{{$t('modelObj.deleteThisModel')}}</el-button>
            
        </div>
    </BaseTitle>
    <!-- 确认删除弹框 -->
    <div id="dataset-del" class="ui small modal">
        <div class="header">
			{{$t('modelObj.deleteThisModel')}}
		</div>
        <div class="content"  v-loading="delLoading">
			<div class="ui warning message text left">
				<p v-html="$t('datasetObj.delete_notices_1')"></p>
				<P v-html="$t('modelObj.delete_notices_2',{model:dataObj.alias})"></P>
			</div>
			<form class="ui form" @submit.prevent>
				<div class="field">
					<label>
						{{$t('modelObj.sureModelName')}}
						<span class="text red">{{dataObj.alias}}</span>
					</label>
				</div>
				<div class="required field">
					<label for="dataset_name">{{$t('modelObj.model_name1')}}</label>
					<input id="dataset_name" v-model="name" name="dataset_name" required>
				</div>

				<div class="text right actions">
					<div class="ui cancel button">{{$t('cancel')}}</div>
					<button class="ui red button" @click="deleteDataset">{{$t('confirm1')}}</button>
				</div>
			</form>
		</div>
    </div>
  </div>
</template>

<script>
import BaseTitle from '~/components/BaseTitle.vue'
import Access from './Access.vue'
import { getModelLicenseList } from '~/apis/modules/modelmanage';
import { editDatasetsDetail, delDataset } from "~/apis/modules/dataset";
import { MODEL_ENGINES } from '~/const'
const MAX_LABEL_COUNT = 5;
export default {
  name: 'AiforgeSetting',
  props: {
    dataObj: { type: Object, default: () => ({}) },
  },
  data() {
    return {
      form: {
        name: this.dataObj.name || "",
        alias: this.dataObj.alias || "",
        aimodel_type: this.dataObj.aimodel_type || 1,
        is_private: this.dataObj.is_private || false,
        engine: this.dataObj.engine || 0,
        label: this.dataObj.labels.join(" ") || "",
        license: this.dataObj?.licenseInfo?.id || "",
      },
      baseLoading: false,

      delLoading: false,

      visible: false,
      name: '',
      owners: [{ name: this.dataObj.owner_name}],
      engineList: MODEL_ENGINES,
      licenseList: [],
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
        engine: [
            { required: true},
        ],

      },
    };
  },
  components: { BaseTitle, Access},
  mounted() {
    console.log(this.dataObj.id)
    getModelLicenseList().then(res => {
      res = res.data;
      try {
        const license = JSON.parse(res) || [];
        this.licenseList = license;
        // if (license.length) {
        // this.form.license = license[0].id;
        // }
      } catch (err) {
        console.log(err);
      }
    }).catch(err => {
    console.log(err);
    });
  },
  computed: {
  },
  methods: {
    labelInput() {
      const hasEndSpace = this.form.label[this.form.label.length - 1] == ' ';
      const list = this.form.label.trim().split(' ').filter(label => label != '');
      this.form.label = list.slice(0, MAX_LABEL_COUNT).join(' ') + (hasEndSpace && list.length < MAX_LABEL_COUNT ? ' ' : '');
    },
    editBaseInfo() {
            this.$refs['form'].validate((valid) => {
                if (valid) {
                    this.baseLoading = true
                    this.editDataset()
                } else {
                    return false
                }
            })
        },
    async editDataset() {
        try {
            const response = await editDatasetsDetail({aimodel_id:this.dataObj.id},this.form,'aimodel')
            console.log(response)
            this.baseLoading = false
            if (response.data.code === 0) {
                this.$message.success("更新成功" || response.data.msg);
                if (this.dataObj.name !== this.form.name) {
                    location.href = `/models/detail/${this.dataObj.owner_name}/${this.form.name}`    
                } else {
                    this.$emit('editSuccess')
                }
            } else {
                this.$message.error(response.data.msg);
            }
        }catch (error) {
            this.baseLoading = false
            this.$message.error(error);
        }
    },

    showModal (){
      let _this = this;
      let ele = `#dataset-del`
      $(ele)
        .modal({
          onDeny: function () {
            console.log('Denied')
          },
          onApprove: function () {
            console.log('Approved')
            // _this.deleteUser(id)
          },
        })
        .modal("show");
    },
    async deleteDataset() {
        if(this.name !== this.dataObj.alias) {
            this.$message.error(this.$t('modelObj.deleteModelNameError'))
            return
        }
        try {
            this.delLoading = true
            const response = await delDataset({aimodel_id:this.dataObj.id},'aimodel')
            console.log(response)
            this.delLoading = false
            if (response.data.code === 0) {
                this.$message.success(this.$t('imagesObj.deleteSuccessTips'));
                window.location.href = '/explore/models'
            } else {
                this.$message.error(response.data.msg);
            }
        }catch (error) {
            this.delLoading = false
            this.$message.error(error);
        }
    },
    
  },
};
</script>
<style lang="less" scoped>
.form-box{
    padding: 3rem 12rem 3rem 2rem;
    /deep/ .el-form-item__label{
        width: 200px;
    }
    /deep/ .el-form-item__content{
        margin-left: 200px;
    }
    .el-select{
        width: 60%;
    }
}
.content-del{
    box-shadow: 0px 2px 6px 0px rgba(255,178,186,0.5);
    border: 1px solid rgba(234,163,170,1);
    /deep/ h4{
        border-bottom-color:rgba(234,163,170,1);
    }
    .main-box{
        display: flex;
        justify-content: space-between;
        padding: 20px 32px 32px 32px;
        .del-btn{
            color: rgba(255,37,37,1);
            border: 1px solid rgba(255,37,37,1);
            height: 34px;
        }
    }
}

.required-name{
    margin-top: 12px;
    margin-bottom: 6px;
    &::after{
        margin: -.2em 0 0 .2em;
        content: '*';
        color: #db2828;
        display: inline-block;
        vertical-align: top;
    }
}
@media screen and (max-width: 767px) {
/* 当视口宽度 ≤ 767px 时生效 */
  .setting-wrap{
    margin: 0 !important;
  }
  
  .form-box{
    padding: 0 0 1px 0 !important;
    /deep/ .el-form-item__label{
      width: auto !important;
    }
    /deep/ .el-form-item__content{
      margin-left: 0 !important;
    }
    .el-select{
      width: 100% !important;
    }
  }
  .content-del{
    .main-box{
        padding: 0 0 12px 10px  !important;
        
        display: block !important;
        .text-wrap{
            margin-bottom: 20px;
        }
    }
  }
} 
</style>
