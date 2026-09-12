<template>
  <div class="form-container">
    <div class="ui message" :class="successFlag ? 'success' : 'negative'" v-show="showMessage" style="margin-top: 10px;">
      <p>{{ showMessageText }}</p>
    </div>
    <div>
      <BaseTitle :title="title" class="content-collar" v-loading="loading">
          <div class="main-box">
            <div class="top-op-area" v-if="canChangeDatasetTeams">
              <div id="search-user-box" class="ui search focus">
                <form ref="myForm" @submit.prevent>
                  <div class="ui input">
                    <input ref="myInput" class="prompt" v-model="collaborator" @keydown.enter.prevent="sumbitCollaborate(true)" @input="handlSearch" :placeholder="placeholder" required >
                  </div>
                </form>
                <div v-if="searching" class="results transition visible" >
                  <a class="result">
                    <div class="content">
                        <div class="title">{{$t('datasetObj.dataSearch')}}</div>
                    </div>
                  </a>
                </div>
                <div v-else-if="showResults" class="results transition visible">
                  <a 
                    v-for="user in searchResults.length > 7 ? searchResults.slice(0, 7) : searchResults" 
                    :key="user.id"
                    @click="selectUser(user)"
                    class="result"
                  >
                    <div class="image" v-if="type === 'Collaborators'" >
                      <img :src="user.avatar_url" alt="">
                    </div>
                    <div class="content">
                      <div class="title">
                        <template v-if="type === 'Teams'">
                          {{ user.name }} ({{ filesType==='dataset' ? user.dataset_permission :  user.aimodel_permission }} access)
                        </template>
                        <template v-else>
                          {{ user.full_name ? `${user.username} (${user.full_name})` : user.username }}
                        </template>
                      </div>
                    </div>
                  </a>
                </div>
                <div v-else-if="showEmpty" class="results transition visible" >
                    <a class="result">
                        <div class="content">
                            <div class="title">{{$t('datasetObj.notFoundUser')}}</div>
                        </div>
                    </a>
                </div>
              </div>
              <button class="ui button add-user" @click="sumbitCollaborate">{{btnText}}</button>
            </div>
            <div class="middle-list-area">
              <div class="list-wrap" v-for="item in usersList[type]">
                <template v-if="type === 'Teams'">
                  <div class="name-img">
                    <a :href="`/org/${item.OwnerName}/teams/${item.Name}`">{{item.Name}}</a>
                  </div>
                  <div class="access-i">
                    <i class="ri-shield-keyhole-fill"></i>
                    <span>{{ filesType==='dataset' ? permissionFilter(item.DatasetAuthorize) : permissionFilter(item.AimodelAuthorize)}}</span>
                  </div>
                  <div class="btn-group" v-if="canChangeDatasetTeams">
                    <el-button type="danger" :disabled="!item.canDelete" round 
                      @click="deleteImage(item.ID)" :style="{ backgroundColor: !item.canDelete ? '' : '#ff2525' }">
                      {{$t('delete')}}
                    </el-button>
                  </div>
                </template>
                <template v-if="type === 'Collaborators'">
                  <div class="name-img">
                    <img :src="item.User.RelAvatarLink"/>
                    <a :href="`/${item.User.Name}`">{{item.User.FullName || item.User.Name}}</a>
                  </div>
                  <div class="access-i">
                    <i class="ri-shield-keyhole-fill"></i>
                    <el-dropdown size="medium" trigger="click" @command="handleCommand">
                        <span class="el-dropdown-link">
                            {{ permissionFilter(item.Collaboration.Mode)}}<i class="el-icon-arrow-down el-icon--right"></i>
                        </span>
                        <el-dropdown-menu slot="dropdown">
                            <el-dropdown-item :disabled="item.Collaboration.Mode ===3 " :command="{mode:3, id: item.User.ID}">{{$t('datasetObj.administrators')}}</el-dropdown-item>
                            <el-dropdown-item :disabled="item.Collaboration.Mode ===2 " :command="{mode:2, id: item.User.ID}">{{$t('datasetObj.writePermission')}}</el-dropdown-item>
                            <el-dropdown-item :disabled="item.Collaboration.Mode ===1 " :command="{mode:1, id: item.User.ID}">{{$t('datasetObj.readPermission')}}</el-dropdown-item>
                        </el-dropdown-menu>
                    </el-dropdown>
                  </div>
                  <div class="btn-group">
                    <el-button type="danger" round style="background-color: #ff2525;" @click="deleteImage(item.User.ID)">{{$t('delete')}}</el-button>
                  </div>
                </template>
              </div>
            </div>
          </div>
      </BaseTitle>
    </div>
    <!-- 确认删除弹框 -->
    <div :id="`${type}-${id}`" class="ui small basic delete modal">
        <div class="ui icon header">
          <i class="trash icon"></i> {{$t('datasetObj.deletCollaboratedTitle')}}
        </div>
        <div class="content">
          <p>{{filesType==='dataset' ? $t('datasetObj.deletCollaboratedTips') : $t('modelManage.deletCollaboratedTips')}}</p>
        </div>
        <div class="actions">
          <div class="ui red basic inverted cancel button">
            <i class="remove icon"></i>
            {{$t('cancelOp')}}
          </div>
          <div class="ui green basic inverted ok button">
            <i class="checkmark icon"></i>
            {{$t('confirmOp')}}
          </div>
        </div>
    </div>
  </div>
</template>

<script>
import BaseTitle from '~/components/BaseTitle.vue'
import { getUsers, postCollaborate, deleteCollaborate, getCollaborate, getTeams, postTeam, deleteTeam, modifyCollaborateAcess } from "~/apis/modules/dataset";
import {debounce} from 'lodash/function';
export default {
  name: 'AiforgeAccess',
  props: {
    id: { type: String, default: '' },
    title: { type: String, default: '' },
    btnText: { type: String, default: '' },
    placeholder: { type: String, default: '' },
    type: { type: String, default: '' },
    org: { type: String, default: '' },
    filesType: { type: String, default: 'dataset' },
  },
  data() {
    return {
      loading: false,
      collaborator: '',
      searching: false,
      showResults: false,
      searchResults: [],
      showEmpty: false,
      usersList: [],
      canChangeDatasetTeams: false,
      showMessage: false,
      showMessageText: '',
      successFlag: false,
      collaboratorUser: ''
    };
  },
  components: {BaseTitle},
  mounted() {
    this.$nextTick(() => {
    if (this.$refs.myInput && document.activeElement === this.$refs.myInput) {
      this.$refs.myInput.blur();
    }
  });
    this.getCollaborateList()
  },
  computed: {
  },
  // filters: {
  //   permissionFilter(value) {
  //     const permissionMap = {
  //       1: "可读权限",
  //       2: "可写权限",
  //       3: "管理员",
  //       4: "所有者",
  //     };
  //     return permissionMap[value] || "未知权限";
  //   },
  // },
  methods: {

    handlSearch: debounce(function() {
      this.searchUsers();
    }, 500),
    async searchUsers() {
      if (!this.collaborator.trim()) {
        this.searchResults = [];
        this.showResults = false;
        this.showEmpty = false
        this.searching = false
        return;
      }
      this.searching = true;
      try {
        // 调用远程搜索API
        const fetchMethod = this.type === 'Teams' ? getTeams : getUsers;
        const queryArgs = this.type === 'Teams'
        ? { q: this.collaborator, org: this.org }
        : { q: this.collaborator };

        const response = await fetchMethod(queryArgs);
        console.log(response)
        const res = response.data
        if (res.ok) {
          if (res.data.length > 2) {
              
            }
            this.searchResults = res.data
            this.showResults = !!res.data.length
            this.showEmpty = !res.data.length
        }else{
            this.$message.error('搜索失败，请稍后重试')
        }
        // this.searchResults = response.data;
      } catch (error) {
        console.error('搜索用户失败:', error);
        this.$message.error('搜索失败，请稍后重试')
        this.searchResults = [];
      } finally {
        this.searching = false;
      }
    },
    // 选择用户
    selectUser(user) {
      if(this.type === 'Teams'){
        this.collaborator = user.name
      }else{
        this.collaborator =  user.full_name ? `${user.username} (${user.full_name})` : user.username
        this.collaboratorUser = user.username
      }
      this.showResults = false;
      // 使用 $nextTick 确保 DOM 更新后聚焦输入框
      this.$nextTick(() => {
        this.$refs.myInput.focus();
      });
    },
    async getCollaborateList(){
      try {
        // 调用远程搜索API
        const data = {
          subject_id: this.id,
          subject_type: this.filesType === 'dataset' ? '1' : '2',
        }
        this.loading = true
        const response = await getCollaborate(data)
        console.log(response)
        const res = response.data
        if(res.code===0){
          this.usersList = res.data
          this.usersList.Teams.forEach((item)=>{
            item.canDelete =  this.filesType === 'dataset' ? item.AllowedToChangeDatasetTeams : item.AllowedToChangeAimodelTeams
          })
          console.log(this.usersList)
          let canOper = this.filesType === 'dataset' ? this.usersList.CanChangeDatasetTeams : this.usersList.CanChangeAimodelTeams
          this.canChangeDatasetTeams = canOper || this.type === 'Collaborators'
        }else{
          this.$message.error(res.msg || '获取失败，请稍后重试')
        }
      } catch (error) {
        console.error('搜索用户失败:', error);
        this.$message.error(error)
      } finally {
       this.loading = false
      }
    },
    async sumbitCollaborate(flag=false){
      if(flag){
        this.collaboratorUser = this.collaborator
      }
      this.$refs.myForm.reportValidity();
      if( !this.collaborator){
        return
      }
      try {
        
        const fetchMethod = this.type === 'Teams' ? postTeam : postCollaborate;
        const queryArgs = {
            subject_id: this.id,
            subject_type: this.filesType === 'dataset' ? '1' : '2',
            ...(this.type === 'Teams' ? {team: this.collaborator} : {collaborator: this.collaboratorUser})
        }
        const response = this.type === 'Teams' 
            ? await fetchMethod(queryArgs, this.filesType)
            : await fetchMethod(queryArgs);
        console.log(response)

        const res = response.data
        if(res.code===0){
          this.collaborator = ''
          await this.getCollaborateList()
          this.showMessageText = this.type === 'Teams' 
            ? this.filesType === 'dataset' 
              ? this.$t('datasetObj.add_team_success') 
              : this.$t('modelManage.add_team_success')
            : this.$t('datasetObj.add_collaborator_success')
          this.successFlag = true
        }else{
          this.showMessageText = res.msg
          this.successFlag = false
        }
      } catch (error) {
        console.error('添加用户失败:', error);
        this.$message.error(error)
      } finally {
        this.collaborator = ''
        this.showResults = false
        this.showEmpty = false
        this.showMessage = true
        setTimeout(() => {
          this.showMessage = false
        }, 2000)
      }
    },
    permissionFilter(value) {
      const permissionMap = {
        1: this.$t('datasetObj.readPermission'),
        2: this.$t('datasetObj.writePermission'),
        3: this.$t('datasetObj.administrators'),
        4: this.$t('datasetObj.dataset_owner'),
      };
      return permissionMap[value] || "未知权限";
    },
    async deleteUser(id){
      try {
        this.loading = true
        const fetchMethod = this.type === 'Teams' ? deleteTeam : deleteCollaborate;
        const queryArgs = {
            subject_id: this.id,
            subject_type: this.filesType === 'dataset' ? '1' : '2',
            ...(this.type === 'Teams' ? {id: id} : {uid: id})
        }
        // 根据方法类型调整参数
        const response = this.type === 'Teams' 
            ? await fetchMethod(queryArgs, this.filesType)
            : await fetchMethod(queryArgs);
        console.log(response)
        const res = response.data
        if(res.code===0){
          this.collaborator = ''
          await this.getCollaborateList()
          this.showMessageText = this.type === 'Teams' 
            ? this.filesType === 'dataset' 
              ? this.$t('datasetObj.remove_team_success') 
              : this.$t('modelManage.remove_team_success')
            : this.$t('datasetObj.remove_collaborator_success')
          this.successFlag = true
          
        }else{
          this.showMessageText = res.msg
          this.successFlag = false
        }
      } catch (error) {
        console.error('添加用户失败:', error);
        this.$message.error(error)
      } finally {
        this.loading = false
        this.showMessage = true
        setTimeout(() => {
          this.showMessage = false
        }, 2000)
      }
    },
    deleteImage(id) {
      let _this = this;
      let ele = `#${_this.type}-${_this.id}`
      $(ele)
        .modal({
          onDeny: function () {
          },
          onApprove: function () {
            _this.deleteUser(id)
          },
        })
        .modal("show");
    },
    async handleCommand(command){
      try {
        const data = {
            uid: command.id,
            mode: command.mode,
            subject_id: this.id,
            subject_type: this.filesType === 'dataset' ? '1' : '2',
            
        }
        const response = await modifyCollaborateAcess(data)
        console.log(response)
        const res = response.data
        if(res.code===0){
            this.$message.success('修改成功！')
            this.getCollaborateList()
        }else{
            this.$message.error(res.msg || '修改失败，请稍后重试')
        }
      } catch (error) {
        console.error('修改失败，请稍后重试:', error);
        this.$message.error(error)
      } finally {
       
      }
    },

  },
};
</script>
<style lang="less" scoped>

.content-collar{
  .main-box{
    .top-op-area{
        display: flex;
        margin: 20px 0 0 40px;
        .add-user{
            background-color: rgba(0,102,255,1);
            color: rgba(255,255,255,1);
            margin-left: 12px;
            border-radius: 100px;
        }
    }
    .middle-list-area{
        margin: 20px;
        .list-wrap{
            border-top: 1px solid rgb(212, 212, 213);
            height: 50px;
            flex: 1;
            min-width: 0;
            display: flex;
            .name-img{
                flex: 1;
                margin-left: 20px;
                display: flex;
                align-items: center;
                gap: 8px;
                img{
                    width: 32px;
                    height: 32px;
                    border-radius: 100%;
                }
            }
            .access-i{
                flex: 1;
                display: flex;
                align-items: center;
                color: #101010;
                font-weight: bold;
                gap: 8px;
                .el-dropdown-link{
                    color: #101010;
                    font-weight: bold;
                }
            }
            .btn-group{
                margin-right: 20px;
                display: flex;
                align-items: center;
            }
        }
    }
  }
}
@media screen and (max-width: 767px) {
  .top-op-area{
    margin: 6px 0 0 0 !important;
  }
  .middle-list-area{
    margin: 20px 0 0 0 !important;
    .name-img{
      margin-left: 0 !important;
    }
    .btn-group{
      margin-right: 0 !important;
    }
  }
  
}
</style>