<template>
  <a class="item" :class="`item__${type}`"  target="_blank" :href="getItemLink(data)">
    <div class="top">
      <div class="top-head">
        <div class="name nowrap"><span :title="lang=='en-US' ? data.name : data.alias">{{ lang=='en-US' ? data.name : data.alias }}</span></div>
        <div class="reconmend-icon" v-if="data.recommend">
          <svg xmlns="http://www.w3.org/2000/svg" class="dZJqQS svg-icon-path-icon fill" viewBox="0 0 24 24" width="20" height="20"><defs></defs><g><path d="M11.99 2C6.47 2 2 6.48 2 12s4.47 10 9.99 10C17.52 22 22 17.52 22 12S17.52 2 11.99 2zm4.24 16L12 15.45 7.77 18l1.12-4.81-3.73-3.23 4.92-.42L12 5l1.92 4.53 4.92.42-3.73 3.23L16.23 18z"></path></g></svg>
        </div>
        <div v-if="data.is_private">
          <svg xmlns="http://www.w3.org/2000/svg" xmlns:xlink="http://www.w3.org/1999/xlink" viewBox="0 0 20 20" width="20.000000" height="20.000000" fill="none">
            <rect id="私有" width="20.000000" height="20.000000" x="0.000000" y="0.000000"/>
            <circle id="椭圆 6" cx="10" cy="10" r="9" fill="rgb(83,121,194)"/>
            <path id="矢量 26" d="M6.99987 8.25012L6.99987 7.75024C6.99987 6.09349 8.34312 4.75024 9.99987 4.75024C11.6566 4.75024 12.9999 6.09349 12.9999 7.75024L12.9999 8.25012L14 8.25012C14.276 8.25012 14.4999 8.47399 14.4999 8.74999L14.4999 14.75C14.4999 15.026 14.276 15.2499 14 15.2499L6.00012 15.2499C5.72412 15.2499 5.50024 15.026 5.50024 14.75L5.50024 8.74999C5.50024 8.47399 5.72412 8.25012 6.00012 8.25012L7.00024 8.25012L6.99987 8.25012ZM13.4997 9.24987L6.49962 9.24987L6.49962 14.2497L13.4997 14.2497L13.4997 9.24987ZM9.49999 12.116C9.19887 11.9397 9.00012 11.618 9.00012 11.2501C9.00012 10.6977 9.44787 10.25 10.0002 10.25C10.5526 10.25 11.0004 10.6977 11.0004 11.2501C11.0004 11.6184 10.8012 11.9401 10.505 12.1137L10.5001 12.1164L10.5001 13.2504L9.49999 13.2504L9.49999 12.1164L9.49999 12.116ZM7.99999 8.25012L12.0001 8.25012L12.0001 7.75024C12.0001 6.64549 11.1046 5.75037 10.0002 5.75037C8.89587 5.75037 8.00037 6.64587 8.00037 7.75024L8.00037 8.25012L7.99999 8.25012Z" fill="rgb(255,255,255)" fill-rule="nonzero"/>
          </svg>
        </div>
      </div>
      <div class="top-bottom">
        <div class="labels">
          <template v-if="data.aimodel_type!==2">
            <span v-for="(item, index) in data.tags" :key="item" class="label">{{ $t(`datasets.${item}`)  }}</span>
            <span v-for="(item, index) in data.tasks" :key="item" class="label">{{ $t(`datasets.${item}`) }}</span>
            <span v-if="data.licenses" class="label"> {{ data.licenses }} </span>
            <span v-if="(data.engineName || data.engine)" class="label"> {{ data.engineName || data.engine }}</span>
            <span v-for="(item, index) in data.labels" :title="item" :key="index" class="label nowrap">{{ item }}</span>
            <span v-if="data.external_source === 'modelscope'" class="label">魔塔社区</span>
          </template>
          <span v-else class="label nowrap">{{ $t('modelObj.model_source')  }}{{data.external_name}}</span>
        </div>
      </div>
    </div>
    <div class="footer">
      <div class="footer-l">
        <a v-if="data.Owner && data.Owner.Name" :href="`/${data.Owner.Name}`" class="avatar-c">
          <img class="avatar" :src="data.Owner.RelAvatarLink">
        </a>
        <span style="margin-left:3px;margin-right:8px;"> {{ $t('repos.updated') }} {{ data.updateTimeStr }} </span>
        <span style="white-space: nowrap;" v-if="hasOnlineUrl">
          <span class="greenPoint"></span>
          {{ $t('modelObj.can_online_infer') }}
        </span>
      </div>
      <div class="footer-r">
        <div class="fav-c r-item" :class="canChangeFav || isSetting ? '' : 'fav-disabled'" @click.prevent.stop="changeFav(data)">
          <i v-if="!isCollected" class="heart outline icon"
            :title="canChangeFav ? $t('star') : $t('datasets.moststars')"></i>
          <i v-if="isCollected" class="heart icon" :title="canChangeFav ?$t('unStar') : $t('datasets.moststars')"></i>
          <span>{{ collected_count }}</span>
        </div>
        <div class="line"></div>
        <span class="r-item" :title="$t('datasets.citations')">
          <i class="el-icon-link"></i>
          <span>{{ data.use_count }}</span>
        </span>
        <div class="line"></div>
        <span class="r-item" :title="$t('datasets.downloadtimes')">
          <i class="el-icon-download"></i>
          <span>{{ data.download_count }}</span>
        </span>
        <template v-if="type==='aimodel'">
          <div class="line"></div>
          <span class="r-item" :title="$t('modelManage.derivativeTimes')">
            <i class="ri-git-merge-line"></i>
            <span>{{ data.derivative_count }}</span>
          </span>
        </template>
        
      </div>
    </div>
  </a>
</template>

<script>
import { setModelFav } from '~/apis/modules/modelsquare';
import { setFavorite, unsetFavorite } from "~/apis/modules/dataset";
import { lang } from '~/langs';
export default {
  name: "RightItem",
  props: {
    data: { type: Object, default: () => ({}) },
    canChangeFav: {  type: Boolean, default: true },
    reloadFlag: {  type: Boolean, default: false },
    type: {  type: String, default: 'dataset' }
  },
  components: {},
  data() {
    return {
      isCollected: false,
      collected_count: 0,
      isSetting: false,
      hasOnlineUrl: 0,
      lang: lang
    };
  },
  watch: {
    data(val, oVal) {
      this.updateData()
    }
  },
  methods: {
    getItemLink(data) {
      if (data.external_url) {
        return data.external_url;
      }
      const safeId = `${data.owner_name}/${data.name}`;
      const url = new URL(location.origin);
      if(this.type==='dataset'){
        url.pathname = `/datasets/detail/${safeId}`
      }else{
        url.pathname = `/models/detail/${safeId}`
      }
      return url.toString();
    },
    async changeFav(item) {
      if (this.isSetting || !this.canChangeFav || item.external_url) return;
      this.isSetting = true;
      let res;
      try {
        let params = {
          [this.type === 'dataset' ? 'dataset_id' : 'aimodel_id']: item.id,
        };
        if (this.isCollected) {
          // 取消收藏
          res = await unsetFavorite(params,this.type);
        } else {
          // 添加收藏
          res = await setFavorite(params,this.type);
        }
        console.log(res);
        // 检查返回状态码
        if (res.data.code === 401) {
          const redirect = encodeURIComponent(window.location.href);
          window.location.href = `/user/login?redirect_to=${redirect}`;
          return;
        }
        // 处理成功响应
        if (res.data.code === 0) {
          this.$emit('changeFav',item.id);
          if (this.isCollected) {
            // 通知父组件状态变化
            if(this.reloadFlag){
              this.$emit('reloadPage');
            }
            this.collected_count -= 1;
          } else {
            this.collected_count += 1;
          }
          this.isCollected = !this.isCollected;
          this.$message.success(this.isCollected ? this.$t('datasets.starSuccess') : this.$t('datasets.unstarSuccess'));
        }else {
          throw new Error(res.data.msg || '操作失败');
        }
      } catch (error) {
        console.error('操作收藏失败:', error);
        this.$message.error(error);
      } finally {
        this.isSetting = false;
      }
    },
    updateData() {
      this.isCollected = this.data.is_collected;
      this.collected_count = this.data.num_stars;
      this.hasOnlineUrl = this.data.hasOnlineUrl;
    }
  },
  beforeMount() {
    this.updateData()
  },
};
</script>

<style scoped lang="less">
.dZJqQS.fill:not([stroke]) {
    fill: rgb(255, 98, 0);
}
.item__dataset{
  background: url("data:image/svg+xml;charset=utf8,%3Csvg%20xmlns%3D%22http%3A%2F%2Fwww.w3.org%2F2000%2Fsvg%22%20xmlns%3Axlink%3D%22http%3A%2F%2Fwww.w3.org%2F1999%2Fxlink%22%20version%3D%221.1%22%3E%3Cdefs%3E%3ClinearGradient%20id%3D%221%22%20x1%3D%220%22%20x2%3D%221%22%20y1%3D%220%22%20y2%3D%220%22%20gradientTransform%3D%22matrix(0.13300000000000006%2C%201.125%2C%20-0.09968631692806922%2C%200.13300000000000006%2C%200.026%2C%20-0.121)%22%3E%3Cstop%20stop-color%3D%22%23e1f4ff%22%20stop-opacity%3D%221%22%20offset%3D%220%22%3E%3C%2Fstop%3E%3Cstop%20stop-color%3D%22%23ffffff%22%20stop-opacity%3D%221%22%20offset%3D%221%22%3E%3C%2Fstop%3E%3C%2FlinearGradient%3E%3C%2Fdefs%3E%3Crect%20width%3D%22100%25%22%20height%3D%22100%25%22%20fill%3D%22url(%231)%22%3E%3C%2Frect%3E%3C%2Fsvg%3E");
}
.item__aimodel{
  background: url("data:image/svg+xml;charset=utf8,%3Csvg%20xmlns%3D%22http%3A%2F%2Fwww.w3.org%2F2000%2Fsvg%22%20xmlns%3Axlink%3D%22http%3A%2F%2Fwww.w3.org%2F1999%2Fxlink%22%20version%3D%221.1%22%3E%3Cdefs%3E%3ClinearGradient%20id%3D%221%22%20x1%3D%220%22%20x2%3D%221%22%20y1%3D%220%22%20y2%3D%220%22%20gradientTransform%3D%22matrix(0.13300000000000006%2C%201.125%2C%20-0.09968631692806922%2C%200.13300000000000006%2C%200.026%2C%20-0.121)%22%3E%3Cstop%20stop-color%3D%22%23e5e5ff%22%20stop-opacity%3D%221%22%20offset%3D%220%22%3E%3C%2Fstop%3E%3Cstop%20stop-color%3D%22%23ffffff%22%20stop-opacity%3D%221%22%20offset%3D%221%22%3E%3C%2Fstop%3E%3C%2FlinearGradient%3E%3C%2Fdefs%3E%3Crect%20width%3D%22100%25%22%20height%3D%22100%25%22%20fill%3D%22url(%231)%22%3E%3C%2Frect%3E%3C%2Fsvg%3E");
}
.item {
  display: flex;
  width: 100%;
  min-height: 128px;
  border-radius: 10px;
  
  border: 1px solid rgba(255,255,255,1);
  display: flex;
  justify-content: space-between;
  flex-direction: column;
  height: 100%;
  &:hover {
    border: 1px solid rgba(0,102,255,1);
    .top .top-head .name {
      color: rgba(0,102,255,1);
    }
  }
  .top {
    .top-head {
      display: flex;
      position: relative;
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
        color: rgba(16,16,16,1);
        font-size: 16px;
        font-weight: 700;
        margin-right: 8px;
        // flex: 1;
        // width: 0;
      }
      .reconmend-icon {
        margin-right: 6px;
      }
    }

    .top-bottom {
      padding: 0 16px;


      .labels {
        display: flex;
        align-items: center;
        flex-wrap: wrap;
        .label{
          margin-right: 12px;
          color: rgba(16,16,16,0.5);
          font-size: 12px;
        }
      }
    }
  }

  .footer {
    display: flex;
    align-items: center;
    justify-content: space-between;
    padding: 16px;
    padding-top: 0;

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
