<template>
  <div class="main-right-c">
    <div class="notice-w main-item">
      <div class="header-b">
        <div class="b-title">{{$t('dashboard.announcements')}}</div>
        <a class="b-desc" href="/home/notice" target="_blank">{{$t('dashboard.moreAnnouncements')}}...</a>
      </div>
      <div class="content-b">
        <el-skeleton style="height:100%" :loading="loadingRight" animated >
          <template slot="template">
            <div v-for="item in skeletonList" :key="item" class="skeleton-item-wrap" >
              <el-skeleton-item variant="text" style="height: 100%; width: 90px; margin-right: 10px;" />
              <el-skeleton-item variant="text" style="flex:1; height: 100%;" />
            </div>
          </template>
          <template>
            <div class="swiper-container">
              <div class="swiper-wrapper">
                <div class="swiper-slide" v-for="(item,index) in overviewConfig.noticeInfo.notices" 
                    :key="`${item.title}-${index}`">
                  <a class="content-item" :href="item.link">
                    <div class="date">{{item.date}}</div>
                    <span class="context nowrap">{{ lang=='en-US' ? item.title_en : item.title }}</span>
                  </a>
                </div>
              </div>
            </div>
          </template>
        </el-skeleton>
      </div>
    </div>
    <div class="action-w main-item">
      <ActivityHeatmap/>
      <ActionList/>
    </div>
    <div class="activity-w main-item">
      <!-- <div class="header-b">
        <div class="b-title">{{$t('dashboard.activities')}}</div>
      </div> -->
      <el-skeleton style="flex:1;height:0;" :loading="loadingRight" animated >
        <template slot="template"><el-skeleton-item variant="image" style="height: 100%;"/>
        </template>
        <template>
          <a class="card-wrap" :href="overviewConfig.activityImageInfo.imageLink">
            <img :src="overviewConfig.activityImageInfo.imageUrl" alt="">
          </a>
        </template>
      </el-skeleton>
    </div>
  </div>
</template>

<script>
import ActivityHeatmap from './ActivityHeatmap.vue';
import ActionList from './ActionList.vue';
import { getOverviewConfig } from '~/apis/modules/dashboard';
import { lang } from '~/langs';
const skeletonList = ['first','second','third','fourth','fifth','sixth']

export default {
  name: "rightCard",
  components: {
    ActivityHeatmap,
    ActionList
  },
  data() {
    return {
      loadingRight: true,
      overviewConfig: {
        noticeInfo: {
          notices: [],
          commitId: '',
        },
        activityImageInfo: {
          imageUrl: '',
          imageLink: '',
        },
      },
      skeletonList: skeletonList,
      lang: lang,
    };
  },
  mounted() {
    this.initData();
  },
  methods: {
    async getOverviewPromote(){
      this.loadingRight = true
      try{
        const response = await getOverviewConfig();
        const res = response.data;
        if(res.Code === 0){
          const data = res.Data;
          console.log('获取数据成功:', data); 
          this.formatAndSetOverviewConfig(data)
        }else{
          this.$message.error(res.Msg || '获取数据失败');
        }
      }catch(error){
        console.error('获取推广数据失败:', error);
        this.$message.error(error.message || '网络错误，请稍后重试');
      }finally{
        this.loadingRight = false
      }
    },
    formatAndSetOverviewConfig(data){
      
      // 转换数据格式（保持响应式）
      Object.assign(this.overviewConfig.noticeInfo, {
        notices: data.notice_info?.notices || [],
        commitId: data.notice_info?.commit_id || ''
      });
      Object.assign(this.overviewConfig.activityImageInfo, {
        imageLink: data.activity_image_info?.image_link || '',
        imageUrl: data.activity_image_info?.image_url || ''
      });
    },
    initSwiper(loop = true) {
      // 检查 Swiper 容器是否存在
      if (!document.querySelector('.swiper-container')) {
        console.warn('Swiper container not found');
        return;
      }
      
      // 销毁旧的 Swiper 实例（防止重复初始化）
      if (this.swiperHandler && this.swiperHandler.destroy) {
        this.swiperHandler.destroy(true, true);
      }
      let length = this.overviewConfig.noticeInfo.notices.length;
      this.swiperHandler = new Swiper(".swiper-container", {
        slidesPerView: 3,
        direction: 'vertical',  // 改为垂直方向
        loop: loop && length > 3, // 只有一张图片时不需要 loop
        spaceBetween: 0,
        autoplay: {
          delay: 2500,
          disableOnInteraction: false,
        },
      });
    },
    async initData() {
      try{
        await Promise.all([
          this.getOverviewPromote(),
        ]);
      } catch(error){
        console.error('获取概览数据失败:', error);
        this.$message.error(error.message || '网络错误，请稍后重试');
      }
      
      this.$nextTick(() => {
        this.initSwiper(true);
      });
    },
  },
}
</script>

<style lang="less" scoped>
.main-item{
  background-color: rgba(255,255,255,1);
  border-radius: 10px;
  border: 1px solid rgba(255,255,255,1);
  padding: 16px 16px 20px 24px;
  display: flex;
  flex-direction: column;
}
.main-right-c{
  flex: 1;
  height: 100%;
  display: flex;
  flex-direction: column;
  width: 0;
  gap: 35px;
  .notice-w{
    height: 180px;
    width: 100%;
    .header-b{
      display: flex;
      justify-content: space-between;
      align-items: flex-end;
      margin-bottom: 18px;
      .b-title{
        height: 30px;
        text-align: center;
        line-height: 30px;
        border-radius: 20px;
        background-color: rgba(255,150,0,1);
        padding: 0 20px;
        color: rgba(255,255,255,1);
      }
      .b-desc{
        line-height: 28px;
        color: rgb(0, 54, 255);
        border-bottom: 1px solid rgba(0, 54, 255, 1);
      }
    }
    .content-b{
      flex: 1;
      height: 0;
      overflow: hidden;
      .skeleton-item-wrap{ 
        display: flex;
        height: 23px;
        margin-bottom: 15px;
      }
      /deep/ .swiper-container{
        height: 100%; // 确保容器有高度
        .swiper-wrapper{
          // flex-direction: column;
          // display: flex;
          // gap: 14px;
          .content-item{
            display: flex;
            align-items: center;
            height: auto !important; // 关键：让高度自适应内容
            .date{
              flex-shrink: 0;
              height: 23px;
              line-height: 23px;
              border-radius: 5px;
              background-color: rgba(228,241,255,1);
              color: rgba(0,98,255,1);
              padding: 0 10px;
            }
            .context{
              margin: 0 8px;
              color: rgba(16,16,16,0.7);
            }
            .new{
              height: 24px;
              line-height: 24px;
              border-radius: 4px;
              background-color: rgba(255,150,0,1);
              color: rgba(255,255,255,1);
              padding: 0 10px 0 6px;
            }
          }
        }
      }
      
    }
  }
  .action-w{
    height: 350px;
    width: 100%;
  }
  .activity-w{
    width: 100%;
    height: 200px;
    .header-b{
      display: flex;
      justify-content: space-between;
      align-items: flex-end;
      margin-bottom: 18px;
      .b-title{
        height: 30px;
        text-align: center;
        line-height: 30px;
        border-radius: 20px;
        background-color: rgba(105,192,255,1);
        padding: 0 20px;
        color: rgba(255,255,255,1);
      }
    }
    /deep/ .el-skeleton{
      height: 100%;
    }
    .card-wrap{
      display: flex;
      height: 100%;
      align-items: center;
      justify-content: center;
      img{
        // width: 318px;
        width: 100%;
        height: 100%;
      }
    }
    
  }
}

@media screen and (min-width: 768px) and (max-width: 1500px) {
  /* 在这里编写只在768px到1500px之间生效的样式 */
  .main-right-c{
    .notice-w{
      max-width: 100% !important;
    }
  }
}
</style>
