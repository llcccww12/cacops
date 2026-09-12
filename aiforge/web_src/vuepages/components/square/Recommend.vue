<template>
  <div v-if="showFlag" class="recommend-container">
    <div class="swiper-container">
      <div class="swiper-wrapper">
        <a :href="item.link" class="swiper-slide" v-for="(item,index) in list" :key="item.Id">
          
          <div class="item-c" :style="{'background-color': item.bgColor}">
            <svg v-if="type=='dataset'" xmlns="http://www.w3.org/2000/svg" class="bg-icon" viewBox="0 0 48 48" width="131" height="131">
              <defs></defs>
              <g>
                <rect width="48" height="48" fill="white" fill-opacity="0.01"></rect>
                <path
                  d="M34 12V20V21C31.0449 21 28.3892 22.2818 26.5585 24.3198C24.9678 26.0906 24 28.4323 24 31C24 31.5789 24.0492 32.1463 24.1436 32.6983C24.6579 35.7046 26.5143 38.2529 29.0741 39.7046C26.4116 40.5096 22.8776 41 19 41C10.7157 41 4 38.7614 4 36V28V20V12"
                  fill="none" stroke="#ffffff0a" stroke-width="4" stroke-linecap="round" stroke-linejoin="round"></path>
                <path
                  d="M34 12C34 14.7614 27.2843 17 19 17C10.7157 17 4 14.7614 4 12C4 9.23858 10.7157 7 19 7C27.2843 7 34 9.23858 34 12Z"
                  fill="none" stroke="#ffffff0a" stroke-width="4" stroke-linecap="round" stroke-linejoin="round"></path>
                <path d="M4 28C4 30.7614 10.7157 33 19 33C20.807 33 22.5393 32.8935 24.1436 32.6983" fill="none"
                  stroke="#ffffff0a" stroke-width="4" stroke-linecap="round" stroke-linejoin="round"></path>
                <path d="M4 20C4 22.7614 10.7157 25 19 25C21.7563 25 24.339 24.7522 26.5585 24.3198" fill="none"
                  stroke="#ffffff0a" stroke-width="4" stroke-linecap="round" stroke-linejoin="round"></path>
                <path
                  d="M44 31C44 36.5228 39.5228 41 34 41C32.2091 41 30.5281 40.5292 29.0741 39.7046C26.5143 38.2529 24.6579 35.7046 24.1436 32.6983C24.0492 32.1463 24 31.5789 24 31C24 28.4323 24.9678 26.0906 26.5585 24.3198C28.3892 22.2818 31.0449 21 34 21C39.5228 21 44 25.4772 44 31Z"
                  fill="none" stroke="#ffffff0a" stroke-width="4" stroke-linecap="round" stroke-linejoin="round"></path>
                <path d="M37 29L39 31L37 33" fill="none" stroke="#ffffff0a" stroke-width="4" stroke-linecap="round"
                  stroke-linejoin="round">
                </path>
                <path d="M31 33L29 31L31 29" fill="none" stroke="#ffffff0a" stroke-width="4" stroke-linecap="round"
                  stroke-linejoin="round">
                </path>
              </g>
            </svg>
            <svg v-if="type=='aimodel'" xmlns="http://www.w3.org/2000/svg" class="bg-icon model-svg" viewBox="0 0 512 512" width="131" height="131"><defs></defs><g><path d="M488.6 250.2L392 214V105.5c0-15-9.3-28.4-23.4-33.7l-100-37.5c-8.1-3.1-17.1-3.1-25.3 0l-100 37.5c-14.1 5.3-23.4 18.7-23.4 33.7V214l-96.6 36.2C9.3 255.5 0 268.9 0 283.9V394c0 13.6 7.7 26.1 19.9 32.2l100 50c10.1 5.1 22.1 5.1 32.2 0l103.9-52 103.9 52c10.1 5.1 22.1 5.1 32.2 0l100-50c12.2-6.1 19.9-18.6 19.9-32.2V283.9c0-15-9.3-28.4-23.4-33.7zM358 214.8l-85 31.9v-68.2l85-37v73.3zM154 104.1l102-38.2 102 38.2v.6l-102 41.4-102-41.4v-.6zm84 291.1l-85 42.5v-79.1l85-38.8v75.4zm0-112l-102 41.4-102-41.4v-.6l102-38.2 102 38.2v.6zm240 112l-85 42.5v-79.1l85-38.8v75.4zm0-112l-102 41.4-102-41.4v-.6l102-38.2 102 38.2v.6z"></path></g></svg>            
            <svg v-if="type=='ai_task_template'" xmlns="http://www.w3.org/2000/svg" class="bg-icon" viewBox="0 0 48 48" width="131" height="131"><defs></defs><g><path d="M48 0H0V48H48V0Z" fill="white" fill-opacity="0.01"></path><path d="M44 14L24 4L4 14V34L24 44L44 34V14Z" fill="none" stroke="#ffffff0a" stroke-width="4" stroke-linejoin="round"></path><path d="M4 14L24 24" fill="none" stroke="#ffffff0a" stroke-width="4" stroke-linecap="round" stroke-linejoin="round"></path><path d="M24 44V24" fill="none" stroke="#ffffff0a" stroke-width="4" stroke-linecap="round" stroke-linejoin="round"></path><path d="M44 14L24 24" fill="none" stroke="#ffffff0a" stroke-width="4" stroke-linecap="round" stroke-linejoin="round"></path><path d="M34 9L14 19" fill="none" stroke="#ffffff0a" stroke-width="4" stroke-linecap="round" stroke-linejoin="round"></path></g></svg>
            <div class="item">
              <div class="name">
                <span class="nowrap">{{ lang=='en-US' ? item.Name : item.Alias }}</span>
                <svg v-if="item.Recommend" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24" width="24" height="24"><defs></defs><g><path fill="#fcca00" d="M11.99 2C6.47 2 2 6.48 2 12s4.47 10 9.99 10C17.52 22 22 17.52 22 12S17.52 2 11.99 2zm4.24 16L12 15.45 7.77 18l1.12-4.81-3.73-3.23 4.92-.42L12 5l1.92 4.53 4.92.42-3.73 3.23L16.23 18z"></path></g></svg>
            </div>
              <div class="label-c">
                <span v-for="(tag, index) in item.Tags" :key="tag" class="label"># {{ $t(`datasets.${tag}`)  }}</span>
                <span v-for="(task, index) in item.Tasks" :key="task" class="label"># {{ $t(`datasets.${task}`) }}</span>
                <span v-if="item.Licenses" class="label"># {{ item.licenses }} </span>
                <span v-if="item.licenses" class="label"> {{ item.licenses }} </span>
                <span v-if="(item.engineName || item.Engine)" class="label"># {{ item.engineName || item.Engine }}</span>
                <span v-for="(label, index) in item.labels" :title="label" :key="index" class="label nowrap"># {{ label }}</span>
              </div>
            </div>
            <!-- <div class="reconmend-icon"><img src="/img/recommend.png" alt=""></div> -->
          </div>
      </a>
      </div>
      <!-- Navigation buttons -->
    </div>
  </div>
</template>

<script>
import { getPromoteDataset } from "~/apis/modules/dataset";
import { MODEL_ENGINES } from '~/const';
import { getListValueWithKey } from '~/utils';
import { lang } from '~/langs';
import { TmplTaskTypes, TmplComputerResouces } from '~/pages/aitasktmpl/tools';

export default {
  props: {
    swiperOptions: {
      type: Object,
      default: () => ({
        slidesPerView: 3,
        spaceBetween: 20,
        loop: true,
        autoplay: {
          delay: 3000,
          disableOnInteraction: false
        },
        breakpoints: {
          320: {
            slidesPerView: 1,
            spaceBetween: 10
          },
          850: {
            slidesPerView: 2,
            spaceBetween: 15
          },
          1440: {
            slidesPerView: 3,
            spaceBetween: 20
          }
        }
      })
    },
    type: {
      type: String,
      default: 'dataset' // dataset,aimodel,ai_task_template
    },
  },
  data() {
    return {
      list: [],
      swiperHandler: null,
      showFlag: false,
      lang: lang
    };
  },
  methods: {
    async getDatasetShowList(type='dataset') {
      try {
        let key
        if(type==='dataset'){
          key = 'home_v2/dataset.json'
        } else if(type==='aimodel') {
          key = 'home_v2/model.json'
        } else if (type == 'ai_task_template') {
          key = 'home_v2/ai_task_template.json'
        }
        const response = await getPromoteDataset({key: key})
        const res = response.data
        console.log(res)
        if (res.code === 0) {
          this.list = res.data || []
          if (!this.list.length) {
            await this.loadModelScopeRecommend(type)
          }
          if (this.list.length) {
            this.list.forEach((item)=>{
              item.Alias = item?.Alias || item.Name
              if(item.Label){
                item.labels = item.Label.trim().split(/\s+/)
              }
              if(item.Engine != null){
                item.engineName = getListValueWithKey(MODEL_ENGINES, item.Engine)
              }
              if (type == 'dataset') {
                item.link = `/datasets/detail/${item.OwnerName}/${item.Name}`
                item.bgColor = 'rgba(0, 102, 255, 1)'
              }
              if (type == 'aimodel') {
                item.link = `/models/detail/${item.OwnerName}/${item.Name}`
                item.bgColor = 'rgba(96, 0, 203, 1)'
              }
              if (type == 'ai_task_template') {
                const jobTypeStr = getListValueWithKey(TmplTaskTypes, item.JobType)
                const computeSourceStr= getListValueWithKey(TmplComputerResouces, item.ComputeSource)
                item.Alias = item.Name
                item.labels = [jobTypeStr, computeSourceStr]
                item.link = `/ai_task_tmpl/detail/${item.ID}`
                item.bgColor = 'rgb(150, 132, 255)'
                delete item.Tags
              }
            })
            this.showFlag = true
          }
          this.$nextTick(() => {
            this.initSwiper();
          });
        }else{
          this.$message.error(res.msg)
        }
      }catch(error){
        console.log(error)
        this.$message.error(error)
      }
    },
    async loadModelScopeRecommend(type) {
      try {
        const endpoint = type === 'aimodel' ? '/api/v1/modelscope/models' : '/api/v1/modelscope/datasets'
        const response = await fetch(`${endpoint}?page=1&page_size=6&sort=downloads`)
        const payload = await response.json()
        if (payload.code !== 0) return
        const data = payload.data || {}
        const items = type === 'aimodel' ? (data.models || []) : (data.datasets || [])
        this.list = items.slice(0, 6).map((item) => {
          const mapped = {
            ...item,
            Name: item.name,
            Alias: item.alias || item.name,
            OwnerName: item.owner_name,
            Recommend: true,
            link: item.external_url,
            bgColor: type === 'aimodel' ? 'rgba(96, 0, 203, 1)' : 'rgba(0, 102, 255, 1)',
          }
          if (type === 'aimodel' && item.tags) {
            mapped.labels = String(item.tags).trim().split(/\s+/).filter(Boolean)
          }
          if (type === 'dataset') {
            mapped.Tasks = item.tasks || []
            mapped.licenses = item.licenses || item.license
          }
          return mapped
        })
      } catch (error) {
        console.log(error)
      }
    },
    initSwiper() {
      if (this.swiperHandler) {
        this.swiperHandler.destroy();
      }
      // Merge default options with props
      const options = {
        ...this.swiperOptions,
      };
      
      this.swiperHandler = new Swiper('.swiper-container', options);
    }
  },
  mounted() {
    console.log("label",this.type)
    this.getDatasetShowList(this.type);
  },
  beforeDestroy() {
    if (this.swiper) {
      this.swiper.destroy();
    }
  }
};
</script>

<style scoped lang="less">
.recommend-container {
  margin-bottom: 40px;
  position: relative;
  
  .swiper-container {
    padding: 10px 0;
    overflow: hidden;
    
    .swiper-wrapper {
      display: flex;
    }
  }

  .item-c {
    background-color: rgba(0, 102, 255, 1);
    height: 139px;
    box-shadow: 0px 4px 4px 0px rgba(16,16,16,0.1);
    border-radius: 15px;
    position: relative;

    .bg-icon {
      position: absolute;
      right: 0;
      top: 0;
      &.model-svg:not([stroke]) {
        fill: rgba(255, 255, 255, 0.04);
      }
    }

    .item {
      position: absolute;
      width: 100%;
      height: 100%;
      display: flex;
      flex-direction: column;
      align-items: center;
      justify-content: center;

      .name {
        display: flex;
        align-items: flex-end;
        justify-content: center;
        color: rgba(255, 255, 255, 1);
        font-size: 20px;
        font-weight: 500;
        width: 100%;
        height: 40px;
        span{
          max-width: 80%;        
          text-shadow: rgba(0, 0, 0, 0.4) 0px 2px 6px;
        }
        svg{
            margin-left: 5px;
            transform: translateY(-50%);
            color: #ff9900;
        }
      }

      .label-c {
        display: flex;
        align-items: center;
        justify-content: center;
        width: 100%;
        margin-top: 14px;
        flex-wrap: wrap;
        height: 20px;
        overflow: hidden;
        .label {
          color: rgba(255, 255, 255, 0.7);
          font-size: 16px;
          font-weight: 700;
          max-width: 30%;
          overflow: hidden;
          text-overflow: ellipsis;
          white-space: nowrap;
          margin: 0 5px;
        }
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
}

@media (max-width: 768px) {
  .recommend-container {
    .swiper-container{
      padding: 0;
    }
    .swiper-button-prev,
    .swiper-button-next {
      display: none;
    }
  }
}
</style>