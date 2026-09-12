<template>
  <a class="item" target="_blank" :href="`/ai_task_tmpl/detail/${data.ID}`">
    <div class="top" :style="{
      // background: data.BgColor ? `linear-gradient(249.36deg, ${data.BgColor[1]} -0.02%, ${data.BgColor[0]} 72.64%)` : ''
    }">
      <div class="top-t">
        <div>{{ data.JobTypeStr }}</div>
        <div>{{ data.ComputeSourceStr }}</div>
      </div>
      <div class="top-b">
        <span :title="data.Name">{{ data.Name }}</span>
        <svg v-if="data.Recommend" xmlns="http://www.w3.org/2000/svg" fill="rgb(255, 98, 0)" viewBox="0 0 24 24"
          width="20" height="20">
          <defs></defs>
          <g>
            <path
              d="M11.99 2C6.47 2 2 6.48 2 12s4.47 10 9.99 10C17.52 22 22 17.52 22 12S17.52 2 11.99 2zm4.24 16L12 15.45 7.77 18l1.12-4.81-3.73-3.23 4.92-.42L12 5l1.92 4.53 4.92.42-3.73 3.23L16.23 18z">
            </path>
          </g>
        </svg>
      </div>
    </div>
    <div class="mid">
      <div><span class="tit">{{ $t('cloudbrainObj.dataset') }}：</span>
        <span v-if="(data.DatasetLists || []).length" class="val">
          <span :title="data.DatasetsStr">{{ data.DatasetsStr }}</span>
        </span>
        <span v-else class="val">--</span>
      </div>
      <div><span class="tit">{{ $t('repos.model') }}：</span>
        <span v-if="(data.ModelLists || []).length" class="val">
          <span :title="data.ModelsStr"> {{ data.ModelsStr }}</span>
        </span>
        <span v-else class="val">--</span>
      </div>
      <div><span class="tit">{{ $t('repos.repos') }}：</span>
        <span v-if="data.RepoOwnerName && data.RepoName" class="val"
          :title="`${data.RepoOwnerName}/${data.RepoName}`">{{
            `${data.RepoOwnerName}/${data.RepoName}` }}</span>
        <span v-else class="val">--</span>
      </div>
      <div><span class="tit">{{ $t('cloudbrainObj.image') }}：</span>
        <span v-if="data.ImageName || data.ImageUrl" class="val" :title="data.ImageName">{{ data.ImageName ||
          data.ImageUrl }}</span>
        <span v-else class="val">--</span>
      </div>
    </div>
    <div class="bottom">
      <div class="bottom-l">
        <a v-if="data.Owner" :href="`/${data.Owner.Name}`" class="avatar-c">
          <img class="avatar" :src="data.Owner.RelAvatarLink">
        </a>
        <span class="update-time" style="margin-left:3px;margin-right:8px;"
          :title="`${$t('repos.updated')} ${data.UpdatedUnixStr}`"> {{ $t('repos.updated') }} {{
            data.UpdatedUnixStr }} </span>
      </div>
      <div class="bottom-r">
        <div class="fav-c r-item" :class="canChangeFav ? '' : 'fav-disabled'" @click.prevent.stop="changeFav(data)">
          <i v-if="!isCollected" class="heart outline icon"
            :title="canChangeFav ? $t('star') : $t('datasets.moststars')"></i>
          <i v-if="isCollected" class="heart icon" :title="canChangeFav ? $t('unStar') : $t('datasets.moststars')"></i>
          <span>{{ collectedCount }}</span>
        </div>
        <div class="line"></div>
        <span class="r-item" :title="$t('taskTmplObj.runTimes')">
          <svg xmlns="http://www.w3.org/2000/svg" fill="rgb(136, 136, 136)" viewBox="0 0 48 48" width="12" height="12">
            <defs></defs>
            <g>
              <rect width="48" height="48" fill-opacity="0.01"></rect>
              <path
                d="M43.8233 25.2305C43.7019 25.9889 43.5195 26.727 43.2814 27.4395C42.763 28.9914 41.9801 30.4222 40.9863 31.6785C38.4222 34.9201 34.454 37 30 37H16C9.39697 37 4 31.6785 4 25C4 18.3502 9.39624 13 16 13L44 13"
                stroke="rgb(136, 136, 136)" fill="none" stroke-width="4" stroke-linecap="round" stroke-linejoin="round">
              </path>
              <path d="M38 7L44 13L38 19" fill="none" stroke="rgb(136, 136, 136)" stroke-width="4"
                stroke-linecap="round" stroke-linejoin="round">
              </path>
            </g>
          </svg>
          <span>{{ data.UseCount }}</span>
        </span>
        <span class="run-btn" :title="``" @click.stop.prevent="goRun(data)">
          <svg xmlns="http://www.w3.org/2000/svg" fill="rgb(255, 255, 255)" viewBox="0 0 48 48" width="12" height="12">
            <defs></defs>
            <g>
              <path d="M24 44C12.9543 44 4 35.0457 4 24C4 12.9543 12.9543 4 24 4C35.0457 4 44 12.9543 44 24" fill="none"
                stroke="rgb(255, 255, 255)" stroke-width="4" stroke-linecap="round" stroke-linejoin="round"></path>
              <path d="M20 24V17.0718L26 20.5359L32 24L26 27.4641L20 30.9282V24Z" fill="none"
                stroke="rgb(255, 255, 255)" stroke-width="4" stroke-linejoin="round"></path>
              <path d="M37.0508 32L37.0508 42" fill="none" stroke="rgb(255, 255, 255)" stroke-width="4"
                stroke-linecap="round" stroke-linejoin="round"></path>
              <path d="M42 36.9497L32 36.9497" fill="none" stroke="rgb(255, 255, 255)" stroke-width="4"
                stroke-linecap="round" stroke-linejoin="round"></path>
            </g>
          </svg>
          <span>{{ $t('taskTmplObj.run') }}</span>
        </span>
      </div>
    </div>
  </a>
</template>

<script>
import { putCollectAiTaskTmpl, deleteCollectAiTaskTmpl } from '~/apis/modules/aitasktmpl';

export default {
  name: "Item",
  props: {
    condition: { type: Object, default: () => ({}) },
    data: { type: Object, default: () => ({}) },
  },
  components: {},
  data() {
    return {
      isCollected: false,
      collectedCount: 0,
      isSetting: false,
      canChangeFav: true,
    };
  },
  watch: {
    data(val, oVal) {
      this.updateData()
    }
  },
  methods: {
    changeFav(item) {
      if (this.isSetting) return;
      const isLogin = !!document.querySelector('meta[name="_uid"]');
      if (!isLogin) {
        window.location.href = `/user/login?redirect_to=${encodeURIComponent(window.location.href)}`;
        return;
      }
      this.isSetting = true;
      let setApi = null;
      if (this.isCollected) {
        setApi = deleteCollectAiTaskTmpl;
      } else {
        setApi = putCollectAiTaskTmpl;
      }
      setApi({
        id: item.ID,
      }).then(res => {
        this.isSetting = false;
        if (res.data.code == '0') {
          this.isCollected = !this.isCollected;
          this.collectedCount = Math.max((this.collectedCount + (this.isCollected ? 1 : -1)), 0);
          this.$message.success(this.isCollected ? this.$t('datasets.starSuccess') : this.$t('datasets.unstarSuccess'));
          this.$emit('changeFav');
        } else {
          this.$message.error(res.data.msg);
        }
      }).catch(err => {
        console.log(err);
        this.$message.error(err);
        this.isSetting = false;
      });
    },
    updateData() {
      this.isCollected = this.data.IsCollected;
      this.collectedCount = this.data.NumCollections;
      this.canChangeFav = true;
    },
    goRun(item) {
      window.location.href = `/cloudbrains/create?tmpl=${item.ID}`;
    }
  },
  beforeMount() {
    this.updateData()
  },
};
</script>

<style scoped lang="less">
.item {
  display: flex;
  width: 100%;
  min-height: 146px;
  background: rgb(255, 255, 255);
  box-shadow: rgba(16, 16, 16, 0.1) 0px 4px 4px 0px;
  display: flex;
  justify-content: space-between;
  flex-direction: column;
  height: 100%;
  border-radius: 10px;

  .top {
    height: 80px;
    background: rgba(216, 225, 255, 0.85);
    padding: 14px;
    border-radius: 10px 10px 0px 0px;

    .top-t {
      display: flex;

      div {
        display: flex;
        align-items: center;
        justify-content: center;
        height: 22px;
        background: rgba(16, 16, 16, 0.3);
        color: rgb(255, 255, 255);
        border-radius: 3px;
        padding: 0 5px;
        margin-right: 5px;
        font-size: 12px;
      }
    }

    .top-b {
      display: flex;
      align-items: center;
      margin-top: 12px;

      span {
        display: inline-block;
        color: rgba(16, 16, 16, 1);
        font-size: 16px;
        font-weight: 700;
        overflow: hidden;
        text-overflow: ellipsis;
        white-space: nowrap;
        max-width: calc(100% - 20px);
      }

      svg {
        margin-left: 2px;
      }
    }
  }

  .mid {
    padding: 12px 14px;
    padding-bottom: 0;

    div {
      display: flex;
      align-items: center;
      margin-bottom: 4px;

      span {
        font-size: 12px;

        &.tit {
          color: rgba(16, 16, 16, 0.5);
          margin-right: 2px;
        }

        &.val {
          width: 0;
          flex: 1;
          color: rgba(16, 16, 16, 0.8);
          overflow: hidden;
          text-overflow: ellipsis;
          white-space: nowrap;
        }
      }
    }
  }

  .bottom {
    padding: 14px;
    padding-top: 10px;
    display: flex;
    align-items: center;
    justify-content: space-between;
    margin-top: 2px;

    .bottom-l {
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

      .update-time {
        overflow: hidden;
        text-overflow: ellipsis;
        white-space: nowrap;
      }
    }

    .bottom-r {
      display: flex;
      align-items: center;
      justify-content: flex-end;;
      flex-shrink: 0;

      .line {
        height: 16px;
        margin-left: 8px;
        border-right: 1px solid rgba(16, 16, 16, 0.1);
      }

      .r-item {
        display: flex;
        align-items: center;
        margin-left: 8px;
        color: rgba(136, 136, 136, 1);
        font-size: 12px;

        i,
        svg {
          margin-right: 5px;
        }

        span {}
      }

      .fav-c {
        display: flex;
        align-items: flex-start;

        i {
          color: rgb(250, 140, 22);
        }

        &.fav-disabled {
          i {
            cursor: default;
          }
        }
      }

      .run-btn {
        margin-left: 10px;
        display: flex;
        align-items: center;
        background: rgb(0, 102, 255);
        height: 24px;
        font-size: 12px;
        padding: 0 6px;
        border-color: rgba(157, 197, 226, 0.4);
        border-style: solid;
        border-width: 1px;
        border-radius: 4px;
        color: rgb(255, 255, 255);
        cursor: pointer;

        svg {
          margin-right: 4px;
        }
      }
    }
  }

  &:hover {
    .top {
      .top-b {
        span {
          color: rgb(0, 102, 255);
        }
      }
    }
  }
}
</style>
