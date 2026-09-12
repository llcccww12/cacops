<template>
  <div>
      <h4 id="about-desc" class="ui header desc-home">{{i18n['introduction']}}
        <a class="edit-icon" v-if="IsPermit" id ="editBtn" href="javascript:void(0)" @click="editClick" >
        <i class="gray edit outline icon" style="margin-right: 0;"></i>
        </a>
    </h4>
    <edit-dialog-cmpt 
		:vmContext="vmContext" 
		:dialogTitle="i18n['edit_repository_information']"
		v-model="editDataDialog"
		:deleteCallback="editDataFunc"
		:deleteLoading ="editDataListLoading"
    deleteParam = "ruleForm"
    @input="initForm"
		>
			<div slot="title">
				
			</div>
			<div slot="content">
				<el-form label-position="top" :model="info" :rules="rule" ref="ruleForm">
					<el-form-item :label="i18n['introduction']" prop="desc">
						<el-input v-model="info.desc" type="textarea" :placeholder="i18n['please_enter_the_content']" :autosize="{minRows:4,maxRows:6}" maxlength="255" show-word-limit></el-input>
					</el-form-item>
					<el-form-item :label="i18n['homePage']" prop="index_web" >
						<el-input v-model="info.index_web" :placeholder="`${i18n['homePage']}(eg: https://openi.pcl.ac.cn)`"></el-input>
					</el-form-item>
				</el-form>
			</div>
    </edit-dialog-cmpt>
  </div>
</template>

<script>

const {_AppSubUrl, _StaticUrlPrefix, csrf} = window.config;

import editDialogCmpt from './basic/editDialog.vue';


export default {
  components: {
    editDialogCmpt,
  },
  data() {
    return {
      vmContext: this,
      editDataDialog: false,
      editDataListLoading: false,
      url: '',
      info: {
        desc: '',
        index_web: '',
        repo_name_name: '',
        alias:'',
      },
      IsPermit: false,
    //   rule1:[{min:3,max:5,message:'1',trigger:"blur"}],
      rule: {
        index_web: [
          {required: false, pattern: /(^$)|(^(http|https):\/\/(([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5]).*)|(^(http|https):\/\/[a-zA-Z0-9]+([_\-\.][a-zA-Z0-9]+)*\.[a-zA-Z]{2,10}(:[0-9]{1,10})?(\?.*)?(\/.*)?$)/,message:'请输入有效的URL',tigger:['change','blur']}
        ]
      }
    };
  },
  methods: {
    getIsSigned(){
      var isadmin= document.getElementById("repo-desc").dataset.isadmin;
      var isarchived=document.getElementById("repo-desc").dataset.isarchived;
      // console.log("IsSigned:",this.IsSigned)
      if ((isadmin==true || isadmin=="true")&& (isarchived==false || isarchived=="false")){
        this.IsPermit=true;
      }
    },
    editClick() {
      if (this.IsPermit){
        this.editDataDialog = true;
      }
    },
    getDesc() {
      const el = $('span.description').text();
      this.info.desc = el;
    },
    getWeb() {
      const el = $('a.link.edit-link').text();
      this.info.index_web = el;
    },
    getRepoName() {
      const el = this.url.split('/')[2];
      this.info.repo_name = el;
      this.info.alias = $('input#edit-alias').val()
    },
    initForm(diaolog) {
      if (diaolog === false) {
        this.getRepoName();
        this.getDesc();
        this.getWeb();
      }

    },
    editDataFunc(formName) {
      this.$refs[formName].validate((valid)=>{
        if (valid) {
          this.$axios({
            method: 'post',
            url: this.url,
            header: {'content-type': 'application/x-www-form-urlencoded'},
            data: this.qs.stringify({
              _csrf: csrf,
              action: 'update',
              alias:this.info.alias,
              repo_name: this.info.repo_name,
              description: this.info.desc,
              website: this.info.index_web
            })
          }).then((res) => {
            location.reload();
            this.editDataDialog = false;
          }).catch((error) => {
            this.editDataDialog = false;
          })
        }
        else {
          return false;
        }
      })

    },
    getUrl() {
      const url =document.getElementById("repo-desc").dataset.repolink;
      this.url = `${url}/settings`;
    }
  },
  mounted() {
    this.getIsSigned();
    this.getUrl();
    this.getRepoName();
    this.getDesc();
    this.getWeb();
  },
  watch:{
    'info.desc'(val){
      if(val.length===256){
        this.info.desc = val.substr(0,255)
      }

    }
  },
  created() {
    this.i18n = window.i18n;
    this.getIsSigned();
  }

};
</script>

<style scoped>
.edit-icon{
    color: #8c92a4;
}
.desc-home{
  display: flex;
  justify-content: space-between;
}

</style>
