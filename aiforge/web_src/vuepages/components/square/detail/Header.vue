<template>
  <div class="bg">
    <div class="ui container">
      <div class="title">
        <div class="title-l"  @mouseover="hover = true" @mouseleave="hover = false"
          :style="{ 'max-width': type === 'dataset' ? '74%' : '69%' }">
          <Icons style="flex-shrink: 0;" :type="type" :isPrivate="dataObj.is_private" :modelType="dataObj.aimodel_type"></Icons>
          <div class="t-name nowrap">
            <span class="t-name-data" :title="lang=='en-US' ? dataObj.name : dataObj.alias">{{lang=='en-US' ? dataObj.name : dataObj.alias}}</span>
          </div>
          <!-- <img src="/img/jian.svg" v-if="dataObj.recommend == 1"> -->
          
          <div class="reconmend-icon" v-if="dataObj.recommend">
            <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24" width="20" height="20"><defs></defs><g><path fill="#FF6200" d="M11.99 2C6.47 2 2 6.48 2 12s4.47 10 9.99 10C17.52 22 22 17.52 22 12S17.52 2 11.99 2zm4.24 16L12 15.45 7.77 18l1.12-4.81-3.73-3.23 4.92-.42L12 5l1.92 4.53 4.92.42-3.73 3.23L16.23 18z"></path></g></svg>
          </div>
        </div>
        <div class="title-r" :style="{ 'min-width': type === 'dataset' ? '340px' : '440px' }">
          <div class="btn-group">
            <div @click="changeFav(dataObj)" class="base-btn like-btn" style="cursor: pointer;">
              <template  v-if="!isCollected" >
                <i class="heart outline icon" :title="$t('star')"></i>
                <span>{{ $t('star') }}</span>
              </template>
              <template  v-else>
                <i class="heart icon" :title="$t('unStar')"></i>
                <span>{{ $t('unStar') }}</span>
              </template>
            </div>
            <div class="base-btn">{{ collected_count }}</div>
          </div>
          <div class="btn-group">
            <div class="base-btn">
              <i class="el-icon-link"></i>
              <span>{{ $t('datasets.citations1') }}</span>
            </div>
            <div class="base-btn">{{ dataObj.use_count }}</div>
          </div>
          <div class="btn-group">
            <div class="base-btn">
              <i class="el-icon-download"></i>
              <span>{{ $t('datasets.downloadtimes1') }}</span>
            </div>
            <div class="base-btn">{{ dataObj.download_count }}</div>
          </div>
          <div v-if="type!=='dataset'" class="btn-group">
            <div class="base-btn">
              <i class="ri-git-merge-line"></i>
              <span>{{ $t('modelManage.derivativeTimes') }}</span>
            </div>
            <div class="base-btn">{{ dataObj.derivative_count }}</div>
          </div>
        </div>
      </div>
      <div class="sub-data-title"> 
        <div class="nowrap">
          <a style="color: rgba(3, 102, 214, 1);" :href="`/${dataObj.owner_name}`">{{ dataObj.owner_name }}</a>
          <span> / </span>
          <span style="color: rgba(16, 16, 16, 0.8);">{{ dataObj.name }}</span>
        </div>
        <div class="copy-btn">
          <a href="javascript:;" class="ui poping up clipboard" id="clipboard-btn" data-position="top center"
            data-variation="inverted tiny" :data-success="$t('copySuccess')" :data-content="$t('copy')"
            :data-original="$t('copy')" :data-clipboard-text="`${dataObj.owner_name}/${dataObj.name}`">
            <i class="ri-file-copy-line"></i>
          </a>
        </div>
      </div>
      <div v-if="dataObj.aimodel_type!==2" class="sub-title">
        <div class="labels">
          <span v-for="(item, index) in dataObj.tags" :key="item" class="label">{{ $t(`datasets.${item}`)  }}</span>
          <span v-for="(item, index) in dataObj.tasks" :key="item" class="label">{{ $t(`datasets.${item}`) }}</span>
          <span v-if="dataObj.licenses" class="label"> {{ dataObj.licenses }} </span>
          <span v-if="(dataObj.engineName || dataObj.engine)" class="label"> {{ dataObj.engineName || dataObj.engine }}</span>
          <span v-for="(item, index) in dataObj.labels" :title="item" :key="index" class="label nowrap">{{ item }}</span>
        </div>
      </div>
      <div v-else class="sub-title">
        <div class="labels">
          <span class="label nowrap">{{ $t('modelObj.model_source')  }}{{ dataObj.external_name }}</span>
        </div>
      </div>
      <div class="tabs-wrap">
        <div class="tabs">
          <div class="tabs-l">
            <div 
              class="tab" 
              :class="tabIndex == item.key ? 'focus' : ''"
              v-for="(item) in tabList" 
              :key="item.key" 
              @click="changeTab(item)"
            >
             <i :class="item.icon"></i>
             <span>{{ $t(item.name) }}</span>
            </div>
          </div>
          <div class="tabs-r">
            <div class="tab" @click="changeTab({key:'settings'})" :class="tabIndex == 'settings' ? 'focus' : ''" v-if="dataObj && dataObj.can_manage"> 
                <i class="ri-settings-2-line"></i><span>{{ $t('modelManage.settings') }}</span>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import Icons from './Icons.vue';
import { setFavorite , unsetFavorite} from "~/apis/modules/dataset";
import { initClipboard } from '~/utils';
import { lang } from '~/langs';
export default {
  name: "Header",
  props: {
    dataObj: { type: Object, default: () => ({}) },
    tabList: {  type: Array, default: () => [] },
    tab: { type: String, default: 'intro' },
    type: { type: String, default: 'dataset' },
  },
  components: { Icons },
  data() {
    return {
      isCollected: false,
      collected_count: 0,
      isSetting: false,
      isLogin: false,
      tabIndex: this.tab,
      lang: lang
    };
  },
  methods: {
    async changeFav(item) {
      if (this.isSetting) return;
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
        // 检查返回状态码
        if (res.data.code === 401) {
          const redirect = encodeURIComponent(window.location.href);
          window.location.href = `/user/login?redirect_to=${redirect}`;
          return;
        }
        // 处理成功响应
        if (res.data.code === 0) {
          if (this.isCollected) {
            this.collected_count -= 1;
          } else {
            this.collected_count += 1;
          }
          this.isCollected = !this.isCollected;
          this.$message.success(this.isCollected ? this.$t('datasets.starSuccess') : this.$t('datasets.unstarSuccess'));
          // 通知父组件状态变化
          this.$emit('changeFav');
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
    createOnlineDebug() {
      
    },
    changeTab(item){
        this.tabIndex = item.key
        this.$emit('changeTab', item)
    },
  },
  watch: {

  },
  mounted() { 
    this.$nextTick(() => {
        initClipboard('.copy-btn .clipboard');
    })
    this.isCollected = this.dataObj.is_collected;
    this.collected_count = this.dataObj.num_stars;
    this.isLogin = !!document.querySelector('meta[name="_uid"]');
    console.log(this.dataObj);
  },
  beforeMount() { },
};
</script>

<style scoped lang="less">
.bg {
  height: 155px;
  border-color: rgba(204, 204, 255, 0.6);
  border-width: 0px 0px 1px;
  border-style: solid;
  .container {
    position: relative;
    height: 100%;
  }
}

.title {
  display: flex;
  align-items: center;
  justify-content: space-between;
  padding-top: 16px;
  margin-bottom: 2px;

  .title-l {
    display: flex;
    align-items: center;
    max-width: 74%;
    margin-right: 12px;
    .t-name{
      font-size: 16px;
      color: rgb(16, 16, 16);
      margin-left: 10px;
      margin-right: 10px;
      .t-name-data{
        font-size: 20px;
        font-family: Arial;
        font-weight: 700;
        line-height: 28px;
      }
    }
    
    .reconmend-icon{
      margin-top: 5px;
    }
    img {
      height: 19px;
      width: 20px;
      margin-left: 8px;
    }
  }

  .title-r {
    display: flex;
    align-items: center;
    min-width: 340px;
    gap: 10px;
    .btn-group{
      display: flex;
      .base-btn{
        display: flex;
        align-items: center;
        height: 30px;
        font-size: 14px;
        padding: 4px 8px;
        line-height: 21px;
        border: 1px solid rgba(225,227,230,1);
        border-radius: 4px 0px 0px 4px;
        background-color: rgba(247,247,247,1);
        color: rgba(16,16,16,0.5);
        &:last-child {
          border-radius: 0px 4px 4px 0px;
          color: rgba(16,16,16,1);
          background-color: rgba(255,255,255,1);
        }
        i{
          font-size: 16px;
          margin-right: 4px;
        }
      }
      .like-btn {
        cursor: pointer;
        i {
          height: auto;
          color: rgb(250, 140, 22);
        }
      }
    }
    .create-btn {
      font-size: 14px;
      margin-right: 1rem;
      background-color: #1684FC;
    }
  }
}
.sub-data-title{
  color: rgba(16, 16, 16, 1);
  margin-left: 40px;
  line-height: 20px;
  min-height: 20px;
  font-size: 12px;
  display: flex;
  .copy-btn{
    margin-left: 7px;
    i{
      font-size:16px;
      color: #919191;
    }
  }
}
.sub-title {
  margin-top: 8px;
  .labels{
    display: flex;
    gap: 6px;
    .label{
      height: 22px;
      line-height: 22px;
      border-radius: 3px;
      background-color: rgba(219,229,255,1);
      color: rgba(0,102,255,1);
      font-size: 12px;
      padding: 0 8px;
    }
  }
}

.tabs-wrap {
  position: absolute;
  bottom: 0;
  width: 100%;

  .tabs {
    display: flex;
    justify-content: space-between;
    margin-bottom: -1px;

    .tabs-l {
      display: flex;
      align-items: center;
    }

    .tabs-r {
      display: flex;
      align-items: center;
    }

    .tab {
      font-size: 16px;
      padding: 10px 15px;
      color: rgba(16, 16, 16, 0.5);
      display: flex;
      align-items: center;
      cursor: pointer;
      span{
        font-size: 14px;
      }

      i {
        margin-right: 4px;
      }

      &.focus {
        color: #101010;
        border-color: rgba(204, 204, 255, 0.6);
        border-width: 1px 1px 0px;
        border-style: solid;
        border-radius: 4px 4px 0px 0px;
        background: rgb(255, 255, 255);
      }
    }
  }
}

.use-dlg-btn {
  margin-right: 1rem;
  text-decoration: underline;
}
@media screen and (max-width: 767px) {
/* 当视口宽度 ≤ 767px 时生效 */
    /* #loadContainer { display: none; } */
    
    .tabs-l{
      .tab{
        padding: 10px 6px !important;
      }
    }
    .tabs-r{
      span{
        display: none;
      }
    }
    .title-r{
      display: none !important;
    }
    .sub-title{
      span{
        // &:not(:first-child){
        //   display: none !important;
        // }
      }
      
    }
}
</style>
