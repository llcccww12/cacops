<template>
  <div class="item">
    <div class="item-top">
      <div class="title-c">
        <div class="title">
          <span class="nowrap">
            <span :title="data.tag">{{ data.tag }}</span>
          </span>
          <div style="margin-left: 6px;display: flex;" v-if="data.type == 5">
            <svg xmlns="http://www.w3.org/2000/svg" class="dZJqQS svg-icon-path-icon fill" viewBox="0 0 24 24" width="20" height="20"><defs></defs><g><path d="M11.99 2C6.47 2 2 6.48 2 12s4.47 10 9.99 10C17.52 22 22 17.52 22 12S17.52 2 11.99 2zm4.24 16L12 15.45 7.77 18l1.12-4.81-3.73-3.23 4.92-.42L12 5l1.92 4.53 4.92.42-3.73 3.23L16.23 18z"></path></g></svg>
          </div>
        </div>
        <div class="fav-c" :class="condition.tab == 1 ? 'fav-disabled' : ''" @click.prevent.stop="changeFav(data)">
          <i v-if="!isStar" class="heart outline icon" :title="$t('star')"></i>
          <i v-if="isStar" class="heart icon" :title="$t('unStar')"></i>
          <span>{{ numStars }}</span>
        </div>
      </div>
      <div class="labels">
        <a class="label label-compute-resource"
          :style="data.computeResourceColor ? `background-color: ${data.computeResourceColor};` : ''">{{
            data.computeResourceShow }}</a>
        <a v-for="(item, index) in data.topics" :key="index" class="label">{{ item }}</a>
      </div>
      <div class="content-row">
        <div class="content-row-item half">
          <div>{{ $t('imagesObj.framework') }}：</div>
          <div class="nowrap" :title="data.framework">{{ data.framework }} {{ data.frameworkVersion }}</div>
        </div>
        <div class="content-row-item half">
          <div>{{ $t('imagesObj.pyVersion') }}：</div>
          <div class="nowrap" :title="data.pythonVersion" v-show="data.pythonVersion">Python {{ data.pythonVersion }}</div>
        </div>
      </div>
      <div class="content-row">
        <div class="content-row-item half">
          <div>{{ $t('imagesObj.cudaVersion') }}：</div>
          <div v-if="data.compute_resource != 'GPU'">--</div>
          <div class="nowrap" :title="data.cudaVersion" v-else>
            <span v-show="data.cudaVersion">Cuda {{ data.cudaVersion }}</span>
          </div>
        </div>
        <div class="content-row-item half">
          <div>{{ $t('imagesObj.operationSystem') }}：</div>
          <div class="nowrap" :title="data.operationSystem">{{ data.operationSystem }} {{ data.operationSystemVersion }}</div>
        </div>
      </div>
      <div class="content-row">
        <div class="content-row-item half">
          <div>{{ $t('imagesObj.cannVersion') }}：</div>
          <div v-if="data.compute_resource != 'NPU'">--</div>
          <div class="nowrap" :title="data.cannVersion" v-else>
            <span v-show="data.cannVersion">Cann {{ data.cannVersion }}</span>
          </div>
        </div>
        <div class="content-row-item half">
          <div>{{ $t('imagesObj.pyPackge') }}：</div>
          <div class="package-wrap" :title="data.thirdPackagesShow">{{ data.thirdPackagesShow }}</div>
        </div>
      </div>
      <div class="content-row">
        <div class="content-row-item half">
          <div>{{ $t('imagesObj.dtkVersion') }}：</div>
          <div v-if="data.compute_resource != 'DCU'">--</div>
          <div class="nowrap" :title="data.dtkVersion" v-else>
            <span v-show="data.dtkVersion">Dtk {{ data.dtkVersion }}</span>
          </div>
        </div>
        <div class="content-row-item half">
          <div>{{ $t('imagesObj.imageTaskType') }}：</div>
          <div class="package-wrap" :title="data.trainTypeShow" v-show="data.trainType && data.status=='1'">{{ data.trainTypeShow }}</div>
        </div>
      </div>
      <div class="content-row">
        <div class="content-row-item">
          <div>{{ $t('imagesObj.descr') }}：</div>
          <div class="descr" :title="data.description"> {{ data.description }} </div>
        </div>
      </div>
    </div>
    <div class="footer">
      <div class="footer-l">
        <a v-if="data.userName" :href="`/${data.userName}`" class="avatar-c" :title="data.userName">
          <img class="avatar" :src="data.relAvatarLink">
        </a>
        <a v-else href="javascript:;" class="avatar-c">
          <img class="avatar" src="/user/avatar/Ghost/-1">
        </a>
        <span style="margin-right:10px;" :title="$t('modelManage.createTime')"> {{ data.createTimeStr }} </span>
        <span style="margin-right:10px;" :title="$t('datasets.citations')">
          <i class="el-icon-link"></i>
          <span style="color:rgba(16, 16, 16, 0.9);">{{ data.useCount }}</span>
        </span>
        <span class="status" style="margin-right:10px;" v-if="condition.tab == 1">
          <el-tooltip v-if="data.status == 0" effect="dark" :content="$t('imagesObj.imageCommitting')" placement="top">
            <i class="CREATING" style="margin-right:3px"></i>
          </el-tooltip>
          <el-tooltip v-if="data.status == 1" effect="dark" :content="$t('imagesObj.imageCommitSuccess')"
            placement="top">
            <i class="SUCCEEDED" style="margin-right:3px"></i>
          </el-tooltip>
          <el-tooltip v-if="data.status == 2 && compute_resource == 'GPU'" effect="dark" :content="$t('imagesObj.imageCommitErrorTips1')"
            placement="top">
            <i class="FAILED" style="margin-right:3px"></i>
          </el-tooltip>
          <el-tooltip v-if="data.status == 2 && compute_resource != 'GPU'" effect="dark" :content="$t('imagesObj.imageCommitErrorTips2')"
            placement="top">
            <i class="FAILED" style="margin-right:3px"></i>
          </el-tooltip>
          <span v-if="data.status === 0">{{ $t('imagesObj.committing') }}</span>
          <span v-if="data.status === 1">{{ $t('imagesObj.commitSuccess') }}</span>
          <span v-if="data.status === 2">{{ $t('imagesObj.commitFailed') }}</span>
        </span>
        <span class="apply-status" v-if="condition.tab == 1">
          <div v-if="data.status === 1">
            <div v-if="data.apply_status === 2" style="display: flex;align-items:center;justify-content:center;">
              <el-tooltip effect="dark" :content="$t('imagesObj.recommendNeedReview')" placement="top">
                <i class="CLOCK" style="margin-right:3px"></i>
              </el-tooltip>
              <span style="color: rgb(250, 140, 22);">{{ $t('imagesObj.recommendNeedReview') }}</span>
            </div>
            <div v-if="data.apply_status === 3" style="display: flex;align-items:center;justify-content:center;">
              <el-tooltip effect="dark" :content="$t('imagesObj.recommendReviewApproved')" placement="top">
                <i class="SUCCEEDED" style="margin-right:3px"></i>
              </el-tooltip>
              <span style="color: rgb(19, 194, 141);">{{ $t('imagesObj.recommendReviewApproved') }}</span>
            </div>
            <div v-if="data.apply_status === 4" style="display: flex;align-items:center;justify-content:center;">
              <el-tooltip effect="dark" :content="data.message" placement="top">
                <i class="FAILED" style="margin-right:3px"></i>
              </el-tooltip>
              <span style="color: red">{{ $t('imagesObj.recommendReviewFailed') }}</span>
            </div>
          </div>
        </span>
      </div>
      <div class="footer-r" v-if="data.status != 0">
        <a v-if="condition.tab == 1 && data.status === 1 && (data.apply_status == 0 || data.apply_status == 1 || data.apply_status === 4) && data.type != 5"
          class="btn-op" href="javascript:void(0);" @click="apply(data)">
          <i class="ri-rocket-2-line" style="margin-right:0;"></i>
          <span>{{ $t('imagesObj.applyForRecommend') }}</span>
        </a>
        <a v-if="condition.tab == 1 && data.apply_status != 3" class="btn-op" href="javascript:void(0);"
          @click="edit(data)">
          <i class="el-icon-edit-outline"></i>
          <span>{{ $t('modelManage.edit') }}</span>
        </a>
        <a v-if="condition.tab == 1" class="btn-op delete" href="javascript:void(0);" @click="deleteImage(data)">
          <i class="el-icon-delete"></i>
          <span>{{ $t('cloudbrainObj.delete') }}</span>
        </a>
      </div>
    </div>
  </div>
</template>

<script>
import { putImageAction, deleteImage, getImageById } from '~/apis/modules/images';

export default {
  name: "Item",
  props: {
    condition: { type: Object, default: () => ({}) },
    data: { type: Object, default: () => ({}) },
  },
  components: {},
  data() {
    return {
      isStar: false,
      numStars: 0,
      isSetting: false,
      hasOnlineUrl: 0,
      refreshTimer: null,
    };
  },
  watch: {
    data: {
      handler(newVal) {
        this.isStar = newVal.isStar;
        this.numStars = newVal.numStars;
        this.refreshStatus();
      },
      immediate: true,
      deep: true,
    },
  },
  methods: {
    changeFav(item) {
      if (this.condition.tab == 1) return;
      if (this.isSetting) return;
      this.isSetting = true;
      putImageAction({
        id: item.id,
        action: this.isStar ? 'unstar' : 'star',
      }).then(res => {
        this.isSetting = false;
        if (res.data.Code == '0') {
          this.isStar = !this.isStar;
          this.numStars = this.numStars + (this.isStar ? 1 : -1);
          this.$message.success(this.isStar ? this.$t('datasets.starSuccess') : this.$t('datasets.unstarSuccess'));
          this.$emit('changeImage');
        } else if (res.data.code == '401') {
          window.location.href = `/user/login?redirect_to=${encodeURIComponent(window.location.href)}`;
        } else {
          this.$message.error(res.data.msg);
        }
      }).catch(err => {
        console.log(err);
        this.isSetting = false;
        if (err.request.responseURL.indexOf('/user/login') >= 0) {
          window.location.href = `/user/login?redirect_to=${encodeURIComponent(window.location.href)}`;
        } else {
          this.$message.error(err);
        }
      });
    },
    copy(item) {
      const tInput = document.createElement("input");
      tInput.value = item.place;
      document.body.appendChild(tInput);
      tInput.select();
      document.execCommand("Copy");
      tInput.remove();
      this.$message({
        type: 'success',
        message: this.$t('copySuccess'),
      });
    },
    refreshStatus() {
      this.refreshTimer && clearInterval(this.refreshTimer);
      this.refreshTimer = setInterval(() => {
        if (this.data.status != 0) {
          this.refreshTimer && clearInterval(this.refreshTimer);
          return;
        }
        getImageById({ id: this.data.id }).then(res => {
          res = res.data;
          if (res.id == this.data.id) {
            const updateData = { ...res };
            delete updateData['userName'];
            delete updateData['relAvatarLink'];
            const trainTypeList = [];
            const trainTypes = res.trainType.split('&');
            trainTypes.forEach(type => {
              if (type) {
                if(type === 'Notebook'){
                  if(['GCU', 'GPU'].includes(res.compute_resource)){
                    trainTypeList.push(this.$t('TaskTypeTitle.Notebook'));
                  }else{
                    trainTypeList.push(this.$t('TaskTypeTitle.Notebook1'));
                  }
                }else{
                  trainTypeList.push(this.$t('TaskTypeTitle.' + type));
                }
              }
            });
            updateData.trainTypeShow = trainTypeList.join('、')
            this.$emit('refreshImage', updateData);
          }
        }).catch(err => {
          console.log(err);
        });
      }, 5 * 1000);
    },
    deleteImage(item) {
      if (this.condition.tab != 1) return;
      if (this.isSetting) return;
      this.$confirm(this.$t('imagesObj.deleteTips'), this.$t('tips'), {
        confirmButtonText: this.$t('confirm1'),
        cancelButtonText: this.$t('cancel'),
        type: 'warning',
        lockScroll: false,
      }).then(() => {
        this.isSetting = true;
        deleteImage({ id: item.id }).then(res => {
          this.isSetting = false;
          res = res.data;
          if (res.Code == '0') {
            this.$message({
              type: 'success',
              message: this.$t('imagesObj.deleteSuccessTips'),
            });
            this.$emit('changeImage');
          } else {
            this.$message({
              type: 'error',
              message: res.Message,
            });
          }
        }).catch(err => {
          this.isSetting = false;
          console.log(err);
          this.$message({
            type: 'error',
            message: this.$t('operationFailed'),
          });
        });
      }).catch(() => { });
    },
    apply(item) {
      location.href = `/image/${item.id}/apply`;
    },
    edit(item) {
      location.href = `/image/${item.id}/imageSquare?trainType=${item.trainType.replace(/&/g, '_')}`;
    }
  },
  beforeMount() {
    this.isStar = this.data.isStar;
    this.numStars = this.data.numStars;
  },
  mounted() {
    this.data.status == 0 && this.refreshStatus();
  },
  beforeDestroy() {
    this.refreshTimer && clearInterval(this.refreshTimer);
  }
};
</script>

<style scoped lang="less">
.dZJqQS.fill:not([stroke]) {
    fill: #ff6200;
}
.item {
  border: 1px solid rgba(255,255,255,1);
  background: linear-gradient(176.14deg, rgba(229, 246, 239, 1) 0.84%, rgba(255, 255, 255, 1) 62.45%);
  border-radius: 10px;
  box-shadow: 0px 0px 20px 0px rgba(221,221,221,0.5);
  overflow: hidden;
  padding: 15px;
  display: flex;
  justify-content: space-between;
  flex-direction: column;
  height: 100%;

  &:hover {
    border-color: rgba(204, 204, 255, 0.6);
    border: 1px solid rgba(187,193,246,0.7);
    background: rgba(255,255,255,1);

    .title span {
      color: rgba(0,102,255,1);
    }
  }
}

.title-c {
  display: flex;
  align-items: center;

  .title {
    flex: 1;
    overflow: hidden;
    text-overflow: ellipsis;
    white-space: nowrap;
    display: flex;
    align-items: center;
    color: #101010;
    font-weight: 700;

    a {
      max-width: calc(100% - 40px);
      overflow: hidden;
      text-overflow: ellipsis;

      span {
        font-weight: 500;
        font-size: 16px;
        color: rgb(3, 102, 214);
      }
    }
  }
}

.fav-c {
  display: flex;
  font-size: 12px;

  i {
    color: rgb(250, 140, 22);
    cursor: pointer;
  }

  span {
    color: rgb(16, 16, 16);
  }

  &.fav-disabled {
    i {
      cursor: default;
    }
  }
}

.content-row {
  display: flex;
  align-items: center;
  font-size: 12px;
  color: rgb(136, 136, 136);
  margin-top: 4px;

  .content-row-item {
    display: flex;
    min-width: 100%;

    &.half {
      min-width: 50%;
    }

    div:nth-child(1) {
      width: 90px;
      text-align: right;
    }

    div:nth-child(2) {
      flex: 1;
      width: 0;
      color: rgb(16, 16, 16);
      word-break: break-all;
    }
    
    .package-wrap{
      word-break: break-all;
      display: -webkit-box;
      -webkit-line-clamp: 2; /* 限制显示的行数 */
      -webkit-box-orient: vertical;
      overflow: hidden;
      text-overflow: ellipsis;
    }
  }
}

.descr {
  font-size: 12px;
  font-weight: 300;
  color: rgb(136, 136, 136);
  text-overflow: ellipsis;
  word-break: break-all;
  display: -webkit-box;
  -webkit-box-orient: vertical;
  -webkit-line-clamp: 1;
  max-height: 41px;
  overflow: hidden;
}

.labels {
  display: flex;
  flex-wrap: wrap;
  margin-top: 8px;

  .label {
    color: rgba(16, 16, 16, 0.8);
    border-radius: 4px;
    font-size: 12px;
    background: rgba(232, 232, 232, 0.6);
    padding: 0 6px;
    margin-right: 8px;
    margin-bottom: 6px;
    cursor: default;
  }

  .label-compute-resource {
    color: white;
  }
}

.footer {
  margin-top: 12px;
  display: flex;
  align-items: center;
  flex-wrap: wrap;
  justify-content: space-between;
  min-height: 20px;
  font-size: 12px;

  .footer-l {
    font-size: 12px;
    font-weight: 400;
    overflow: hidden;
    text-overflow: ellipsis;
    white-space: nowrap;
    color: rgba(136, 136, 136, 1);
    display: flex;
    align-items: center;

    .avatar-c {
      margin-right: 4px;
      display: inline-block;
      height: 24px;

      .avatar {
        display: inline-block;
        width: 24px;
        height: 24px;
        border-radius: 100%;
      }
    }

    .status {
      display: flex;
      align-items: center;
    }

    .apply-status {
      display: flex;
      align-items: center;
    }
  }

  .footer-r {
    display: flex;
    justify-content: flex-end;;
    font-size: 12px;
    font-weight: 300;
    color: rgb(136, 136, 136);
    margin-left: auto;
    .btn-op {
      margin-left: 12px;
      display: flex;
      align-items: center;

      i {
        margin-right: 2px;
      }

      &.delete {
        color: rgb(255, 37, 37);
      }
    }
  }
}

:lang(en-US) .content-row {
  .content-row-item {
    div:nth-child(1) {
      width: 112px;
    }
  }
}
@media only screen and (max-width: 800px){
  .center {
    .el-pagination {
      /deep/ .el-pagination__total,
      /deep/ .el-pagination__sizes,
      /deep/ .el-pagination__jump {
        display: none;
      }
    }
  }
  .status{
    span{
      display: none;
    }
  }
  .btn-op{
    span{
      display: none;
    }
  }

}

</style>
