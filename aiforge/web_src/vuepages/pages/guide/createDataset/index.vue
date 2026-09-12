<template>
    <div style="margin-top:32px;">
        <div class="__mobile-tip" >
          <img style="width: 132px;height: 97px;" src="/img/model/pc-view.png" alt="">
          <div style="margin-top: 2rem;">{{$t('useInPcWeb')}}</div>
        </div>
        <create-form :title="title" v-loading="loading" :active="active">
            <div v-if="active===0">
                <el-form ref="form" :model="form" :rules="rules" hide-required-asterisk label-width="220px">
                    <el-form-item :label="$t('datasetObj.dataset_name')" prop="name" style="margin-bottom: 0;">
                        <el-input v-model="form.name" maxlength="100"></el-input>
                        <span style="font-size: 12px;color: #888;line-height: 1;margin-top: 0.5em;display: inline-block;">
                            {{ $t('datasetObj.dataset_name_tooltips') }}
                        </span>
                    </el-form-item>
                    <el-form-item :label="$t('datasetObj.dataset_zh_name')" prop="alias" style="margin-bottom: 0;">
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
                    <el-form-item :label="$t('datasetObj.category')" prop="tags">
                        <el-select v-model="form.tags[0]" :placeholder="$t('datasetObj.select_category')" filterable style="width: 60%;">
                            <el-option v-for="item in Category" :key="item.name" :value="item.name" :label="$t('datasets.'+ item.name)"></el-option>
                        </el-select>
                    </el-form-item>
                    <el-form-item :label="$t('datasetObj.task')" prop="tasks">
                        <el-select v-model="form.tasks[0]" :placeholder="$t('datasetObj.select_task')" filterable style="width: 60%;">
                            <el-option v-for="item in Task" :key="item.name" :value="item.name" :label="$t('datasets.'+ item.name)"></el-option>
                        </el-select>
                    </el-form-item>
                    <el-form-item :label="$t('datasetObj.license')" prop="licenses">
                        <el-select v-model="form.licenses" :placeholder="$t('datasetObj.license_helper')" filterable style="width: 60%;">
                            <el-option v-for="item in License" :key="item.name" :value="item.name" :label="item.name"></el-option>
                        </el-select>
                    </el-form-item>
                    <el-form-item class="btn-wrap"> 
                        <el-button round class="cancel" type="info" @click="cancelDataset()">{{$t('cancel')}}</el-button>
                        <el-button round class="submit" type="primary" @click="createDataset()">{{$t('cloudbrainObj.nextStep')}}</el-button>
                    </el-form-item>
                </el-form>
            </div>
            <div v-if="active===1" >
                <FileUpload :modelId="datasetID" @cancel="cancelUpload" @uploadFinish="uploadFinish"></FileUpload>
            </div>
        </create-form>
    </div>
</template>
<script >
import CreateForm from '../components/CreateForm.vue'; 
import FileUpload from '../components/FileUpload.vue';
import { createDatasets, getAvailableUsers } from '~/apis/modules/dataset';
import { Category, Task, License } from '~/pages/dataset/square/constant.js'
import { getUrlSearchParams } from '~/utils';
export default {
    components: { CreateForm, FileUpload },
    data() {
        return {
            owners: [],
            orgId: 0,
            cancelUrl:'',
            form: {
                name: "",
                alias: "",
                is_private: false,
                tags: [],
                tasks: [],
                licenses: "",
                owner_id: ''
            },
            loading:true,
            Category: Category,
            Task: Task,
            License: License,
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
                tags: [
                    { required: true, message: this.$t('datasetObj.category_required'), trigger: "change" },
                ],
                tasks: [
                    {
                        required: true,
                        message: this.$t('datasetObj.select_task_required'),
                        trigger: "change",
                    },
                ],

            },
            active: 0,
            title: [this.$t('datasetObj.createDataset'), this.$t('uploadDatasetFile')],
            datasetID: '',
            datasetObj: {}
        }
    },
    methods: {
        createDataset() {
            this.$refs['form'].validate((valid) => {
                if (valid) {
                    this.loading = true
                    this.postCreateDataset()
                } else {
                    return false
                }
            })
        },
        async postCreateDataset() {
            try {
                const response = await createDatasets(this.form)
                console.log(response)
                this.loading = false
                if (response.data.code === 0) {
                    this.active = 1
                    this.datasetID = response.data.data.id
                    this.datasetObj = response.data.data
                } else {
                    this.$message.error(response.data.msg);
                }
            }catch (error) {
                this.loading = false
                this.$message.error(error);
            }
        },
        cancelDataset() {
            if (this.cancelUrl) {
                window.location.href = this.cancelUrl;
            } else {
                const referrer = document.referrer || '';
                if (referrer.indexOf('guide/create_dataset') > -1) {
                    window.location.href = '/dashboard';
                } else {
                    window.history.back();
                }
            }
        },
        async getUsersList() {
            try {
                const response = await getAvailableUsers({org:this.orgId})
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
            location.href = `/datasets/detail/${this.datasetObj.owner_name}/${this.datasetObj.name}`
        },
        uploadFinish(){
            window.setTimeout(() => {
                location.href = `/datasets/detail/${this.datasetObj.owner_name}/${this.datasetObj.name}?tab=files`
            }, 1000);
        }
    },
    mounted() { },
    async beforeMount (){
        const urlParams = getUrlSearchParams();
        if (urlParams.backurl) {
            this.cancelUrl = urlParams.backurl;
        }
        if(urlParams.org){
            this.orgId = urlParams.org
            console.log(urlParams.org)
        }
        await this.getUsersList()
        
    },
}
</script>
<style>
.full.height{
    background-color: rgba(245,245,246,1);
}
</style>


<style lang='less' scoped>
::v-deep .is-required .el-form-item__label::after {
  content: "*";
  color: #ff0000;
  margin-left: 4px;
}
.btn-wrap{
    display:flex;
    justify-content: flex-end;
    .cancel{
        background-color: #e0e3ea; 
        color: #020004;     
        border-color: rgb(224, 227, 234);
    }
    .submit{
        background-color: #0066ff;
    }
}

</style>