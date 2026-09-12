<template>
  <a class="item" target="_blank"
    :href="`/${data.repoOwnerName}/${data.repoName}/modelmanage/model_readme_tmpl?name=${data.name}`">
    <div class="top">
      <div class="top-head">
        <div class="icon-c"><img src="/img/icons/model-icon.png" alt=""></div>
        <div class="name"><span :title="data.name">{{ data.name }}</span></div>
        <div class="reconmend-icon" v-if="data.recommend == 1"><img src="/img/recommend.png" alt=""></div>
      </div>
      <div class="top-bottom">
        <div class="labels">
          <span v-if="(data.engineName || data.engine)" class="label engine"> {{ data.engineName || data.engine }}
          </span>
          <a v-for="(item, index) in data.labels" :key="index" class="label normal">{{ item }}</a>
        </div>
      </div>
    </div>
    <div class="footer">
      <div class="footer-l">
        <a :href="`/${data.userName}`" class="avatar-c">
          <img class="avatar" :src="data.userRelAvatarLink">
        </a>
        <span style="margin-left:3px;margin-right:8px;"> {{ $t('repos.updated') }} {{ data.updateTimeStr }} </span>
        <span style="white-space: nowrap;" v-if="hasOnlineUrl">
          <span class="greenPoint"></span>
          {{ $t('modelObj.can_online_infer') }}
        </span>
      </div>
      <div class="footer-r">
        <div class="fav-c r-item" :class="canChangeFav ? '' : 'fav-disabled'" @click.prevent.stop="changeFav(data)">
          <i v-if="!isCollected" class="heart outline icon"
            :title="canChangeFav ? $t('star') : $t('datasets.moststars')"></i>
          <i v-if="isCollected" class="heart icon" :title="canChangeFav ?$t('unStar') : $t('datasets.moststars')"></i>
          <span>{{ collectedCount }}</span>
        </div>
        <div class="line"></div>
        <span class="r-item" :title="$t('datasets.citations')">
          <i class="el-icon-link"></i>
          <span>{{ data.referenceCount }}</span>
        </span>
        <div class="line"></div>
        <span class="r-item" :title="$t('datasets.downloadtimes')">
          <i class="el-icon-download"></i>
          <span>{{ data.downloadCount }}</span>
        </span>
        <div class="line"></div>
        <span class="r-item" :title="$t('modelManage.derivativeTimes')">
          <i class="ri-git-merge-line"></i>
          <span>{{ data.derivativeCount }}</span>
        </span>
      </div>
    </div>
  </a>
</template>

<script>
import { setModelFav } from '~/apis/modules/modelsquare';

export default {
  name: "ModelItem",
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
      hasOnlineUrl: 0,
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
      if (this.condition.tab == 2 || this.condition.tab == 5) return;
      if (this.isSetting) return;
      this.isSetting = true;
      setModelFav({
        id: item.id,
        collected: this.isCollected ? false : true,
      }).then(res => {
        this.isSetting = false;
        if (res.data.code == '0') {
          this.isCollected = !this.isCollected;
          this.collectedCount = this.collectedCount + (this.isCollected ? 1 : -1);
          this.$message.success(this.isCollected ? this.$t('datasets.starSuccess') : this.$t('datasets.unstarSuccess'));
          this.$emit('changeFav');
        } else if (res.data.code == '401') {
          window.location.href = `/user/login?redirect_to=${encodeURIComponent(window.location.href)}`;
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
      this.isCollected = this.data.isCollected;
      this.collectedCount = this.data.collectedCount;
      this.hasOnlineUrl = this.data.hasOnlineUrl;
      this.canChangeFav = !(this.condition.tab == 2 || this.condition.tab == 5);
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
  box-shadow: 0px 4px 4px 0px rgba(16, 16, 16, 0.1);
  border-radius: 15px;
  background-color: rgba(250, 250, 255, 1);
  border: 1px solid rgba(16, 16, 16, 0.2);
  display: flex;
  justify-content: space-between;
  flex-direction: column;
  height: 100%;

  .top {

    .top-head {
      display: flex;
      align-items: center;
      position: relative;
      background-color: rgba(255, 255, 255, 1);
      height: 55px;
      border-top-left-radius: 15px;
      border-top-right-radius: 15px;
      padding: 16px;

      .icon-c {
        width: 40px;
        height: 40px;
        margin-right: 6px;

        img {
          width: 100%;
          height: 100%;
        }
      }

      .name {
        color: rgba(14, 37, 69, 1);
        font-size: 16px;
        flex: 1;
        width: 0;

        span {
          display: block;
          overflow: hidden;
          text-overflow: ellipsis;
          white-space: nowrap;
        }
      }

      .reconmend-icon {
        position: absolute;
        right: 15px;
        top: 0px;
        height: 19.28px;
        width: 29.61px;

        img {
          height: 100%;
          width: 100%;
        }
      }
    }

    .top-bottom {
      padding: 16px;
      padding-top: 12px;
      padding-bottom: 0px;

      .labels {
        display: flex;
        align-items: center;
        flex-wrap: wrap;

        .label {
          border-radius: 3px;
          color: rgba(14, 37, 69, 1);
          font-size: 12px;
          padding: 0px 6px;
          margin-right: 6px;
          box-shadow: 0px 1px 1px 0px rgba(16, 16, 16, 0.2);
          margin-bottom: 6px;
          overflow: hidden;
          text-overflow: ellipsis;
          white-space: nowrap;

          &.engine {
            border: 1px solid rgba(145, 213, 255, 0.5);
            background: linear-gradient(92.08deg, rgba(226, 245, 255, 1) -0.83%, rgba(247, 252, 255, 0.5) 19.83%, rgba(247, 252, 255, 0.5) 97.53%);
          }

          &.normal {
            border: 1px solid rgba(255, 198, 145, 0.5);
            background: radial-gradient(0.11261981933593752% 0.6100000000000001% at 31.2% 0.1%, rgba(225, 241, 255, 1) 0%, rgba(229, 255, 246, 0) 100%);
          }
        }
      }
    }
  }

  .footer {
    display: flex;
    align-items: center;
    justify-content: space-between;
    margin-top: 2px;
    padding: 6px 16px 10px 16px;

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

      .greenPoint {
        display: inline-block;
        width: 8px;
        height: 8px;
        margin-right: 5px;
        background-color: #5dbf77;
        border-radius: 50%;
      }
    }

    .footer-r {
      display: flex;
      align-items: center;
      justify-content: flex-end;;

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

        i {
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
    }
  }
}
</style>
