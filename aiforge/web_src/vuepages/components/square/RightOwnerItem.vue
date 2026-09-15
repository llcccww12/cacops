<template>
  <a class="item" :class="`item__${type}`"  target="_blank" :href="getItemLink(data)">
    <div class="top">
      <div class="top-head">
        <Icons style="flex-shrink: 0;width:24px;height:24px;" :type="type" :isPrivate="data.is_private" :modelType="data.aimodel_type"></Icons>
        <div class="name-c">
          <div class="name nowrap"><span :title="lang=='en-US' ? data.name : data.alias">{{ lang=='en-US' ? data.name : data.alias }}</span></div>
          <div class="reconmend-icon" v-if="data.recommend">
            <svg xmlns="http://www.w3.org/2000/svg" class="dZJqQS svg-icon-path-icon fill" viewBox="0 0 24 24" width="20" height="20"><defs></defs><g><path d="M11.99 2C6.47 2 2 6.48 2 12s4.47 10 9.99 10C17.52 22 22 17.52 22 12S17.52 2 11.99 2zm4.24 16L12 15.45 7.77 18l1.12-4.81-3.73-3.23 4.92-.42L12 5l1.92 4.53 4.92.42-3.73 3.23L16.23 18z"></path></g></svg>
          </div>
        </div>
        
        <div class="top-r">
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
      <div class="top-bottom">
        <div class="labels">
          <template v-if="data.aimodel_type!==2">
            <span v-for="(item, index) in data.tags" :key="item" class="label">{{ $t(`datasets.${item}`)  }}</span>
            <span v-for="(item, index) in data.tasks" :key="item" class="label">{{ $t(`datasets.${item}`) }}</span>
            <span v-if="data.licenses" class="label"> {{ data.licenses }} </span>
            <span v-if="(data.engineName || data.engine)" class="label"> {{ data.engineName || data.engine }}</span>
            <span v-for="(item, index) in data.labels" :title="item" :key="index" class="label nowrap">{{ item }}</span>          
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
      <div class="footer-r" v-if="operaFlag">
        <div class="upload-bth" @click.prevent.stop="uploadClick(data,type)">
          <svg width="14" height="14" viewBox="0 0 48 48" fill="none" xmlns="http://www.w3.org/2000/svg"><mask id="icon-04f8237f10b972" maskUnits="userSpaceOnUse" x="0" y="0" width="48" height="48" style="mask-type: alpha"><path d="M48 0H0V48H48V0Z" fill="#0066ff"/></mask><g mask="url(#icon-04f8237f10b972)"><path d="M6 24.0083V42H42V24" stroke="#0066ff" stroke-width="4" stroke-linecap="round" stroke-linejoin="round"/><path d="M33 15L24 6L15 15" stroke="#0066ff" stroke-width="4" stroke-linecap="round" stroke-linejoin="round"/><path d="M23.9917 32V6" stroke="#0066ff" stroke-width="4" stroke-linecap="round" stroke-linejoin="round"/></g></svg>
          <span style="line-height: 1px;">{{$t('modelManage.uploadFile')}}</span>
        </div>
        <div class="delete-bth" @click.prevent="deleteClick(data)">
          <svg width="14" height="14" viewBox="0 0 48 48" fill="none" xmlns="http://www.w3.org/2000/svg"><path d="M9 10V44H39V10H9Z" fill="none" stroke="#0066ff" stroke-width="4" stroke-linejoin="round"/><path d="M20 20V33" stroke="#0066ff" stroke-width="4" stroke-linecap="round" stroke-linejoin="round"/><path d="M28 20V33" stroke="#0066ff" stroke-width="4" stroke-linecap="round" stroke-linejoin="round"/><path d="M4 10H44" stroke="#0066ff" stroke-width="4" stroke-linecap="round" stroke-linejoin="round"/><path d="M16 10L19.289 4H28.7771L32 10H16Z" fill="none" stroke="#0066ff" stroke-width="4" stroke-linejoin="round"/></svg>
          <span style="line-height: 1px;">{{$t('modelManage.delete')}}</span>
        </div>
      </div>
    </div>
  </a>
</template>

<script>
import Icons from './Icons.vue';
import { setModelFav } from '~/apis/modules/modelsquare';
import { setFavorite, unsetFavorite } from "~/apis/modules/dataset";
import { lang } from '~/langs';
export default {
  name: "RightItem",
  props: {
    data: { type: Object, default: () => ({}) },
    canChangeFav: {  type: Boolean, default: true },
    reloadFlag: {  type: Boolean, default: false },
    operaFlag: {  type: Boolean, default: false },
    type: {  type: String, default: 'dataset' }
  },
  components: { Icons },
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
    isRealExternalUrl(url) {
      return typeof url === 'string' && /^https?:\/\//i.test(url);
    },
    getItemLink(data) {
      if (this.isRealExternalUrl(data.external_url)) {
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
      if (this.isSetting || !this.canChangeFav) return;
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
    },
    uploadClick(data, type){
      const safeId = `${data.owner_name}/${data.name}`;
      if(type==='dataset'){
        window.open(`/datasets/detail/${safeId}?tab=files`)
      }else{
        window.open(`/models/detail/${safeId}?tab=files`)
        
      }
    },
    deleteClick(data){
      this.$emit('deleteEvent', data)
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
  box-shadow: 0px 0px 20px 0px rgba(221,221,221,0.5);
  background: linear-gradient(176.14deg, rgba(224,236,255,1) 0.84%,rgba(255,255,255,1) 54.19%);
}
.item__aimodel{
  box-shadow: 0px 0px 20px 0px rgba(221,221,221,0.5);
  background: linear-gradient(176.14deg, rgba(230,224,255,1) 0.84%,rgba(255,255,255,1) 54.19%);
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
    background: #fff;
    border: 1px solid rgba(255,255,255,1);
    .top .top-head .name {
      color: rgba(0,102,255,1);
    }
  }
  .top {
    .top-head {
      display: flex;
      position: relative;
      align-items: center;
      border-top-left-radius: 15px;
      border-top-right-radius: 15px;
      padding: 22px 19px 11px 25px;

      .icon-c {
        width: 40px;
        height: 40px;
        margin-right: 6px;

        img {
          width: 100%;
          height: 100%;
        }
      }
      .name-c { 
        flex: 1;
        width: 0;
        display: flex;
        margin-left: 16px;
      }
      .name {
        color: rgba(16,16,16,1);
        font-size: 14px;
        font-weight: 700;
        margin-right: 8px;
        // flex: 1;
        // width: 0;
      }
      .reconmend-icon {
        display: flex;
        margin-right: 6px;
      }
      .top-r {
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

    .top-bottom {
      padding: 0 19px 14px 67px;
      .labels {
        display: flex;
        align-items: center;
        flex-wrap: wrap;
        gap: 8px;
        .label{
          color: rgba(16,16,16,0.5);
          font-size: 12px;
          height: 22px;
          line-height: 22px;
          padding: 0 6px;
          border-radius: 3px;
          background-color: rgba(16,16,16,0.05);
        }
      }
    }
  }

  .footer {
    display: flex;
    align-items: center;
    justify-content: space-between;
    padding: 13px 19px 16px 22px;
    border-top: 1px solid rgba(0,0,255,0.05);
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
        height: 20px;

        .avatar {
          display: inline-block;
          width: 20px;
          height: 20px;
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
    .footer-r{
      display: flex;
      align-items: center;
      justify-content: flex-end;;
      color: rgba(0,102,255,1);
      font-size: 12px;
      .upload-bth{
        display: flex;
        align-items: center;
      }
      .delete-bth{
        display: flex;
        align-items: center;
        margin-left: 23px;
      }
      svg{
        margin-right: 6px;
      }
    }
    
  }
}
@media only screen and (max-width: 849.99px) {
  .delete-bth {
    margin-left: 13px !important;
  }
}
</style>
