<template>
  <div>
    <div v-if="emptyPage" style="padding-top:50px">
      <NotFound></NotFound>
    </div>
    <div v-else>
      <div v-if="migrateState == 0 || migrateState == -1" class="ui container content-wrap">
        <div class="header">
          <span class="title">{{ $t('modelObj.migrate_from', { url: migrateUrl }) }}</span>
        </div>
        <div class="content">
          <div class="file-list">
            <div class="file-item" v-for="(item, index) in files" :key="index">
              <div class="file-name-c">
                <i class="el-icon-document"></i>
                <span class="file-name" :title="item.filename">{{ item.filename }}</span>
              </div>
              <div class="file-size">{{ item.sizeShow }}</div>
              <div class="progress-c">
                <div v-if="item.status == -1" class="progress wating"></div>
                <div v-if="item.status == 0" class="progress migrating">
                  <div></div>
                </div>
                <div v-if="item.status == 1" class="progress success"></div>
                <div v-if="item.status == 2" class="progress failed"></div>
              </div>
              <div class="status-c">
                <div class="status wating" v-if="item.status == -1">
                  <svg class="icon" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 32 32" width="16" height="16"
                    fill="rgb(187, 187, 187)">
                    <defs></defs>
                    <g>
                      <path
                        d="M16 29.333c-7.364 0-13.333-5.969-13.333-13.333s5.969-13.333 13.333-13.333 13.333 5.969 13.333 13.333-5.969 13.333-13.333 13.333zM16 26.667c5.891 0 10.667-4.776 10.667-10.667s-4.776-10.667-10.667-10.667v0c-5.891 0-10.667 4.776-10.667 10.667s4.776 10.667 10.667 10.667v0zM9.333 17.333h12v2.667h-5.333v4l-6.667-6.667zM16 12v-4l6.667 6.667h-12v-2.667h5.333z">
                      </path>
                    </g>
                  </svg>
                  <span>{{ $t('modelObj.migrate_watting') }}</span>
                </div>
                <div class="status migrating" v-if="item.status == 0">
                  <i class="el-icon-loading"></i>
                  <span>{{ $t('modelObj.migrating') }}</span>
                </div>
                <div class="status success" v-if="item.status == 1">
                  <i class="el-icon-circle-check"></i>
                  <span>{{ $t('modelObj.migrate_success') }}</span>
                </div>
                <div class="status failed" v-if="item.status == 2">
                  <i class="el-icon-circle-close"></i>
                  <span>{{ $t('modelObj.migrate_failed') }}</span>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>
      <div v-if="migrateState == 2" class="ui container failed-wrap">
        <div class="failed-svg-c"><img class="failed-svg" src="/img/rocket-failed.svg"></div>
        <div class="tips">{{ $t('modelObj.migrate_failed_tips') }}
          <a class="btn" @click="retry" href="javascript:;">{{ $t('modelObj.migrate_retry') }} </a>
          {{ $t('modelObj.migrate_try_later') }}
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import NotFound from '~/components/NotFound.vue';
import { getUrlSearchParams, transFileSize } from '~/utils';
import { getModelInfoByName, getModelMigrateStatus, setModelMigrateRetry } from '~/apis/modules/modelmanage';

export default {
  data() {
    return {
      isLogin: false,
      emptyPage: false,
      migrateState: '', // 0-成功,1-迁移中,2-失败
      migrateUrl: '',
      loading: false,
      files: [],
      state: {
        model_path: '',
        hf_repo_id: '',
      },
      model_id: '',
    };
  },
  components: { NotFound },
  methods: {
    goDetail() {
      window.location.href = `/models/detail/${this.state.model_path}`
    },
    getMigrateStatus() {
      getModelMigrateStatus({ hf_repo_id: this.state.hf_repo_id }).then(async(res) => {
        res = res.data;
        if (res.code == 1) {
          this.files = (res.data || []).map((item => {
            return {
              ...item,
              sizeShow: transFileSize(item.size),
            };
          }));
          await this.getDatasetDetail(true)
        } else {
          this.fileError = res.msg;
        }
      }).catch(err => {
        console.log(err);
        this.$message({
          type: 'error',
          message: err,
        });
      });
    },
    refreshMigrateStatus() {
      this.stopTimer();
      this.refreshTimer = setInterval(() => {
        this.getMigrateStatus();
      }, 8000);
    },
    stopTimer() {
      this.refreshTimer && clearInterval(this.refreshTimer);
    },
    retry() {
      if (this.loading) return;
      this.loading = true;
      setModelMigrateRetry({ hf_repo_id: this.state.hf_repo_id }).then(res => {
        res = res.data;
        if (res.code == 1) {
          window.location.reload();
        } else if (res.code == -1) {
          this.loading = false;
          this.$message({
            type: 'error',
            message: res.msg,
          });
        } else {
          this.loading = false;
          this.$message({
            type: 'error',
            message: this.$t('modelObj.re_migrate_failed'),
          });
        }
      }).catch(err => {
        this.loading = false;
        console.log(err);
        this.$message({
          type: 'error',
          message: this.$t('modelObj.re_migrate_failed'),
        });
      });
    },
    async getDatasetDetail(flag=false){
        try {
         this.loading = true
         const response = await getModelInfoByName({aimodel_id: this.model_id})
         const res = response.data
         let ObjTemp = {}
         console.log("xxxx1",res)
         if(res.code === 0 ){
            ObjTemp = res.data
            if(ObjTemp.aimodel_type == 2){
              this.migrateState = ObjTemp.migration.status;
              if(!flag){
                this.state.hf_repo_id = ObjTemp.external_name
                this.state.model_path = `${ObjTemp.owner_name}/${ObjTemp.name}`
                this.migrateUrl = `https://huggingface.co/${ObjTemp.external_name}`;
              }
              if ( this.migrateState === 0 || this.migrateState === -1){ // 迁移中
                if(!flag){
                  this.getMigrateStatus();
                  this.refreshMigrateStatus();
                }
              }else if (this.migrateState === 2){
                return 
              }else{ // 1 迁移成功
                this.goDetail();
              }
            }else{
              this.goDetail();
            }
         } else {
          if (res.code === 9004) {
            this.emptyPage = true
          } else {
            this.$message.error(res.msg)
          }
         }
        } catch (error) {
            this.$message.error(error)
        }finally{
          this.loading = false
        }
    },
  },
  async beforeMount() {
    this.isLogin = !!document.querySelector('meta[name="_uid"]');
    if (!this.isLogin) {
      window.location.href = `/user/login?redirect_to=${encodeURIComponent(window.location.href)}`;
      return;
    }
    const urlParams = getUrlSearchParams();
    if (urlParams.model_id) {
      this.model_id = urlParams.model_id
      await this.getDatasetDetail()
    } else {
      this.emptyPage = true;
    }
  },
  mounted() { },
  beforeDestroy() {
    this.stopTimer();
  },
};
</script>

<style scoped lang="less">
.content-wrap {
  margin-top: 50px;

  .header {
    height: 45px;
    border-color: rgb(212, 212, 213);
    border-width: 1px;
    border-style: solid;
    border-radius: 5px 5px 0px 0px;
    font-size: 14px;
    background: rgb(240, 240, 240);
    display: flex;
    align-items: center;

    .title {
      font-weight: 600;
      font-size: 16px;
      color: rgb(16, 16, 16);
      margin-left: 10px;
    }
  }

  .content {
    margin-top: -1px;
    border-color: rgb(212, 212, 213);
    border-width: 1px;
    border-style: solid;
    padding: 30px 0;
    border-top: none;

    .file-list {
      margin: 0 120px;

      .file-item {
        display: flex;
        height: 40px;
        align-items: center;
        box-sizing: border-box;
        border-bottom: 1px solid #f1f1f1;

        .file-name-c {
          width: 280px;
          display: flex;
          align-items: center;

          i {
            margin-right: 4px;
          }

          .file-name {
            flex: 1;
            width: 0;
            overflow: hidden;
            text-overflow: ellipsis;
            white-space: nowrap;
          }
        }

        .file-size {
          width: 104px;
          padding-left: 20px;
        }

        .progress-c {
          flex: 1;
          display: flex;
          margin: 0 20px;

          .progress {
            width: 100%;
            height: 10px;
            border-radius: 5px;
            background: rgba(16, 16, 16, 0.05);

            &.success {
              background: rgb(56, 158, 13);
            }

            &.migrating {

              div {
                height: 100%;
                width: 50%;
                border-radius: 5px;
                background: linear-gradient(90deg, rgba(16, 16, 16, 0.05) 0%, rgb(4, 100, 255) 100%);
                animation: animationWidth 2.5s infinite ease-in-out;
              }
            }
          }
        }

        .status-c {
          width: 100px;

          .status {
            color: rgb(187, 187, 187);
            display: flex;
            align-items: center;

            i,
            svg {
              margin-right: 4px;
            }

            span {
              line-height: 16px;
            }

            &.success {
              color: rgb(56, 158, 13);
            }

            &.failed {
              color: red;
            }

            &.migrating {
              color: rgb(4, 100, 255);
            }
          }
        }
      }
    }
  }
}

.failed-wrap {
  display: flex;
  flex-direction: column;
  align-items: center;
  margin: 80px 0 40px 0;

  .failed-svg-c {
    height: 328px;
    width: 328px;
  }

  .tips {
    text-align: center;
    margin-top: 10px;
    font-size: 14px;
    color: rgb(16, 16, 16);

    .btn {
      color: rgb(4, 100, 255);
    }
  }
}

@keyframes animationWidth {
  0% {
    width: 0;
  }

  100% {
    width: 100%;
  }
}
</style>
