<template>
  <div class="item-card" @click.stop.prevent="getItemLink()">
    <div class="card-part1">
      <div class="part1-content">
        <div class="part1-content-display">
          <Icons v-if="data.type !== 'aitasktmpl'" style="flex-shrink: 0;width: 24px;height: 24px;" :type="data.type" :isPrivate="data.is_private" />
          <div v-else class="icon-c">
            <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 48 48" width="24" height="24">
              <defs></defs>
              <g>
                <path d="M48 0H0V48H48V0Z" fill-opacity="0.01"></path>
                <path d="M44 14L24 4L4 14V34L24 44L44 34V14Z" fill="none" stroke="#888" stroke-width="4" stroke-linejoin="round"></path>
                <path d="M4 14L24 24" fill="none" stroke="#888" stroke-width="4" stroke-linecap="round" stroke-linejoin="round"></path>
                <path d="M24 44V24" fill="none" stroke="#888" stroke-width="4" stroke-linecap="round" stroke-linejoin="round"></path>
                <path d="M44 14L24 24" fill="none" stroke="#888" stroke-width="4" stroke-linecap="round" stroke-linejoin="round"></path>
                <path d="M34 9L14 19" fill="none" stroke="#888" stroke-width="4" stroke-linecap="round" stroke-linejoin="round"></path>
              </g>
            </svg>

            <svg v-if="data.IsPrivate" class="lock" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 32 32" width="12" height="12">
              <defs></defs>
              <g>
                <path d="M25.333 13.333h1.333c0.736 0 1.333 0.597 1.333 1.333v0 13.333c0 0.736-0.597 1.333-1.333 1.333v0h-21.333c-0.736 0-1.333-0.597-1.333-1.333v0-13.333c0-0.736 0.597-1.333 1.333-1.333v0h1.333v-1.333c0-5.155 4.179-9.333 9.333-9.333s9.333 4.179 9.333 9.333v0 1.333zM6.667 16v10.667h18.667v-10.667h-18.667zM14.667 18.667h2.667v5.333h-2.667v-5.333zM22.667 13.333v-1.333c0-3.682-2.985-6.667-6.667-6.667s-6.667 2.985-6.667 6.667v0 1.333h13.333z" fill="#888"></path>
              </g>
            </svg>
          </div>
          <div class="item-title" :title="data.alias">{{data.alias}}</div>
          <div v-if="data.recommend" class="item-recommend">
            <svg xmlns="http://www.w3.org/2000/svg" class="styles__StyledSVGIconPathComponent-sc-i3aj97-0 dZJqQS svg-icon-path-icon fill" viewBox="0 0 24 24" width="20" height="20"><defs></defs><g><path d="M11.99 2C6.47 2 2 6.48 2 12s4.47 10 9.99 10C17.52 22 22 17.52 22 12S17.52 2 11.99 2zm4.24 16L12 15.45 7.77 18l1.12-4.81-3.73-3.23 4.92-.42L12 5l1.92 4.53 4.92.42-3.73 3.23L16.23 18z"></path></g></svg>
          </div>
        </div>

        <div class="aitasktmpl-content" v-if="data.type == 'aitasktmpl' && !isMobile">
          <el-row :gutter="20">
            <el-col :span="12">
              <div class="aitasktmpl-item">
                <span class="aitasktmpl-subtitle">{{ $t('cloudbrainObj.dataset') }}：</span>
                <span v-if="(data.DatasetLists || []).length" class="val">
                  <span :title="data.DatasetsStr">{{ data.DatasetsStr }}</span>
                </span>
                <span v-else class="val">--</span>
              </div>
              <div class="aitasktmpl-item">
                <span class="aitasktmpl-subtitle">{{ $t('repos.repos') }}：</span>
                <span v-if="data.RepoOwnerName && data.RepoName" class="val"
                  :title="`${data.RepoOwnerName}/${data.RepoName}`">{{
                    `${data.RepoOwnerName}/${data.RepoName}` }}</span>
                <span v-else class="val">--</span>
              </div>
            </el-col>
            <el-col :span="12">
              <div class="aitasktmpl-item">
                <span class="aitasktmpl-subtitle">{{ $t('repos.model') }}：</span>
                <span v-if="(data.ModelLists || []).length" class="val">
                  <span :title="data.ModelsStr"> {{ data.ModelsStr }}</span>
                </span>
                <span v-else class="val">--</span>
              </div>
              <div class="aitasktmpl-item">
                <span class="aitasktmpl-subtitle">{{ $t('cloudbrainObj.image') }}：</span>
                <span v-if="data.ImageName || data.ImageUrl" class="val" :title="data.ImageName">{{ data.ImageName ||
                  data.ImageUrl }}</span>
                <span v-else class="val">--</span>
              </div>
            </el-col>
          </el-row>
        </div>
        <div class="aitasktmpl-content" v-if="data.type == 'aitasktmpl' && isMobile">
          <el-row :gutter="20">
            <div class="aitasktmpl-item">
              <span class="aitasktmpl-subtitle">{{ $t('cloudbrainObj.dataset') }}：</span>
              <span v-if="(data.DatasetLists || []).length" class="val">
                <span :title="data.DatasetsStr">{{ data.DatasetsStr }}</span>
              </span>
              <span v-else class="val">--</span>
            </div>
            <div class="aitasktmpl-item">
              <span class="aitasktmpl-subtitle">{{ $t('repos.repos') }}：</span>
              <span v-if="data.RepoOwnerName && data.RepoName" class="val"
                :title="`${data.RepoOwnerName}/${data.RepoName}`">{{
                  `${data.RepoOwnerName}/${data.RepoName}` }}</span>
              <span v-else class="val">--</span>
            </div>
            <div class="aitasktmpl-item">
              <span class="aitasktmpl-subtitle">{{ $t('repos.model') }}：</span>
              <span v-if="(data.ModelLists || []).length" class="val">
                <span :title="data.ModelsStr"> {{ data.ModelsStr }}</span>
              </span>
              <span v-else class="val">--</span>
            </div>
            <div class="aitasktmpl-item">
              <span class="aitasktmpl-subtitle">{{ $t('cloudbrainObj.image') }}：</span>
              <span v-if="data.ImageName || data.ImageUrl" class="val" :title="data.ImageName">{{ data.ImageName ||
                data.ImageUrl }}</span>
              <span v-else class="val">--</span>
            </div>
          </el-row>
        </div>

        <div class="part1-content-display tag-part">
          <div v-if="data.type == 'aitasktmpl'">
            <div class="aitasktmpl-tag">
              <div>{{ data.JobTypeStr }}</div>
              <div>{{ data.ComputeSourceStr }}</div>
            </div>
          </div>
          <div v-else style="display: flex">
            <div v-for="(item, index) in data.tags" :key="item">
              <span v-if="data.type == 'dataset'" class="item-tag">{{ $t(`datasets.${item}`) }}</span>
              <span v-else-if="data.type == 'model'" class="item-tag">{{ item }}</span>
            </div>
            <div v-if="data.type == 'dataset'">
              <div v-for="(item, index) in data.tasks" :key="item">
                <span class="item-tag">{{ $t(`datasets.${item}`) }}</span>
              </div>
            </div>
            <div>
              <span v-if="data.licenses" class="item-tag"> {{ data.licenses }} </span>
            </div>
            <div>
              <span v-if="data.engineName" class="item-tag">{{ data.engineName }}</span>
            </div>
          </div>
        </div>
      </div>
    </div>

    <div class="card-part2">
      <div class="time-size">
        <span style="margin-right:5px;" :title="`${$t('repos.updated')} ${data.updated_time}`"> 
          {{ $t('repos.updated') }} {{ data.updated_time }}
        </span>
        <span class="size" v-if="data.type != 'aitasktmpl'" :title="`${$t('datasets.size')} ${data.size}`">
          <span>{{ $t('datasets.size') }}：</span>
          <span style="color: rgba(16,16,16,1)">{{data.size}}</span>
        </span>
      </div>
      <div class="item-link">
        <div :title="$t('datasets.moststars')">
          <i :class="data.is_collected ? 'heart icon' : 'heart outline icon'"></i>
          <span>{{data.num_stars}}</span>
        </div>
        <div class="r-item" :title="data.type != 'aitasktmpl' ? $t('datasets.citations') : $t('taskTmplObj.runTimes')">
          <i v-if="data.type != 'aitasktmpl'" class="el-icon-link"></i>
          <svg v-else xmlns="http://www.w3.org/2000/svg" fill="rgb(136, 136, 136)" viewBox="0 0 48 48" width="12" height="12">
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
          <span>{{ data.use_count }}</span>
        </div>
        <div v-if="data.type != 'aitasktmpl'" class="r-item" :title="$t('datasets.downloadtimes')">
          <i class="el-icon-download"></i>
          <span>{{data.download_count}}</span>
        </div>
        <div v-if="data.type == 'model'" class="r-item" :title="$t('modelManage.derivativeTimes')">
          <i class="ri-git-merge-line"></i>
          <span>{{data.derivative_count}}</span>
        </div>
        <span v-if="data.type == 'aitasktmpl'" class="run-btn" :title="``" @click.stop.prevent="goRun(data)">
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
  </div>
</template>

<script>
import Icons from '~/components/square/Icons.vue';

export default {
  name: "Item",
  props: {
    data: { type: Object, default: () => ({}) },
  },
  components: { Icons },
  data() {
    return {
      isMobile: false,
    };
  },
  mounted() {
    this.checkScreenSize();
    window.addEventListener('resize', this.checkScreenSize);
  },
  beforeDestroy() {
    window.removeEventListener('resize', this.checkScreenSize);
  },
  methods: {
    checkScreenSize() {
      this.isMobile = window.innerWidth <= 767;
    },
    getItemLink() {
      const typeToPathMap = {
        aitasktmpl: (data) => `/ai_task_tmpl/detail/${data.ID || data.id}`,
        dataset: (data) => `/datasets/detail/${data.owner_name}/${data.name}`,
        model: (data) => `/models/detail/${data.owner_name}/${data.name}`
      };

      const getPath = typeToPathMap[this.data.type];
      
      if (!getPath) {
        console.warn(`Unknown data type: ${this.data.type}`);
        return;
      }

      const href = getPath(this.data);
      window.open(href, '_blank');
    },
    goRun(item) {
      window.open(`/cloudbrains/create?tmpl=${item.ID}`, '_blank');
    }
  }

}
</script>

<style scoped lang="less">
.dZJqQS.fill:not([stroke]) {
  fill: rgb(255, 98, 0);
}
.item-card {
  border-radius: 10px;
  background-color: rgba(255,255,255,1);
  color: rgba(16,16,16,1);
  font-size: 14px;
  box-shadow: 0px 0px 20px 0px rgba(221,221,221,0.5);
  font-family: PingFangSC-regular;
  border: 1px solid rgba(255,255,255,1);
  margin-bottom: 25px;
  cursor: pointer;

  &:hover {
    .item-title {
      color: rgb(0, 102, 255) !important;
    }
  }
  
  .card-part1 {
    border-bottom: 1px solid rgba(35,36,38,0.1);

    .part1-content {
      padding: 15px 15px 5px 15px;

      .part1-content-display {
        display: flex;
        align-items: center;

        .item-icon {
          width: 24px;
          height: 24px;
        }

        .icon-c {
          height: 24px;
          width: 24px;
          display: flex;
          align-items: center;
          justify-content: center;
          position: relative;
          // margin-right: 10px;

          .lock {
            position: absolute;
            right: -3px;
            bottom: 0;
            background: white;
            border-radius: 4px;
          }
        }

        .item-title {
          color: rgba(16,16,16,1);
          font-size: 14px;
          font-weight: bold;
          padding:0 10px 0 3px;
        }

        

        .item-recommend {
          display: flex;
        }

        .item-tag {
          border-radius: 3px;
          background-color: rgba(16,16,16,0.05);
          color: rgba(16,16,16,0.7);
          font-size: 12px;
          text-align: center;
          padding: 3px 5px;
          margin-right: 5px;
          margin-top: 10px;
        }

        .aitasktmpl-tag {
          display: flex;
          padding-left: 10px;
          // margin-top: 7px;

          div {
            display: flex;
            align-items: center;
            justify-content: center;
            height: 22px;
            background: rgba(16,16,16,0.05);
            color: rgba(16,16,16,0.7);
            border-radius: 3px;
            padding: 0 5px;
            margin-right: 5px;
            font-size: 12px;
          }
        }
      }

      .aitasktmpl-content {
        margin-top: 5px;
        padding-left: 10px;

        .aitasktmpl-item {
          display: flex;
          padding-top: 2px;
          font-size: 12px;

          .aitasktmpl-subtitle {
            color: rgba(16, 16, 16, 0.5);
          }

          .val {
            width: 0;
            flex: 1;
            color: rgba(16, 16, 16, 0.8);
            overflow: hidden;
            text-overflow: ellipsis;
            white-space: nowrap;
          }
        }

        
      }

      .tag-part {
        margin: 10px 0;
      }
    }
  }

  .card-part2 {
    padding: 15px;
    display: flex;
    color: rgba(136,136,136,1);
    font-size: 12px;
    justify-content: space-between;

      .time-size {
        font-family: Arial-regular;
      }
      .item-link {
        display: flex;
        align-items: center;

        .heart {
          color: rgb(250, 140, 22);
        }

        .r-item {
          display: flex;
          align-items: center;
          margin-left: 10px;
          font-size: 12px;

          i, svg {
            color: rgba(136, 136, 136, 1);
            margin-right: 5px;
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

          &:hover {
            background: linear-gradient(47.69deg, rgba(0,102,255,1) -2.37%,rgba(132,0,235,1) 98.81%);
          }

          svg {
            margin-right: 4px;
          }
        }
      }
  }
}

@media only screen and (max-width: 767px) {
  .size {
    display: none;
  }
  .aitasktmpl-tag {
    padding: 0 !important;
  }
  .run-btn {
    background: linear-gradient(47.69deg, rgba(0,102,255,1) -2.37%,rgba(132,0,235,1) 98.81%) !important;
  }
}

</style>