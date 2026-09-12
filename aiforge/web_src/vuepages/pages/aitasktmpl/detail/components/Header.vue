<template>
  <div class="bg">
    <div class="ui container">
      <div class="title-c">
        <div class="title-l">
          <div class="icon-c">
            <svg xmlns="http://www.w3.org/2000/svg" fill="rgb(16, 16, 16)" viewBox="0 0 48 48" width="32" height="32">
              <defs></defs>
              <g>
                <path d="M48 0H0V48H48V0Z" fill-opacity="0.01"></path>
                <path d="M44 14L24 4L4 14V34L24 44L44 34V14Z" fill="none" stroke="rgb(16, 16, 16)" stroke-width="4"
                  stroke-linejoin="round">
                </path>
                <path d="M4 14L24 24" fill="none" stroke="rgb(16, 16, 16)" stroke-width="4" stroke-linecap="round"
                  stroke-linejoin="round">
                </path>
                <path d="M24 44V24" fill="none" stroke="rgb(16, 16, 16)" stroke-width="4" stroke-linecap="round"
                  stroke-linejoin="round">
                </path>
                <path d="M44 14L24 24" fill="none" stroke="rgb(16, 16, 16)" stroke-width="4" stroke-linecap="round"
                  stroke-linejoin="round">
                </path>
                <path d="M34 9L14 19" fill="none" stroke="rgb(16, 16, 16)" stroke-width="4" stroke-linecap="round"
                  stroke-linejoin="round">
                </path>
              </g>
            </svg>
            <svg v-if="data.IsPrivate" class="lock" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 32 32" width="12"
              height="12">
              <defs></defs>
              <g>
                <path
                  d="M25.333 13.333h1.333c0.736 0 1.333 0.597 1.333 1.333v0 13.333c0 0.736-0.597 1.333-1.333 1.333v0h-21.333c-0.736 0-1.333-0.597-1.333-1.333v0-13.333c0-0.736 0.597-1.333 1.333-1.333v0h1.333v-1.333c0-5.155 4.179-9.333 9.333-9.333s9.333 4.179 9.333 9.333v0 1.333zM6.667 16v10.667h18.667v-10.667h-18.667zM14.667 18.667h2.667v5.333h-2.667v-5.333zM22.667 13.333v-1.333c0-3.682-2.985-6.667-6.667-6.667s-6.667 2.985-6.667 6.667v0 1.333h13.333z">
                </path>
              </g>
            </svg>
          </div>
          <div class="title">
            <div class="nowrap">
              <span v-if="data.Owner && data.Owner.Name">{{ data.Owner.Name }} /</span><span :title="data.Name"> {{
                data.Name }}</span>
            </div>
            <div class="reconmend-icon" v-if="data.Recommend">
              <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24" width="20" height="20">
                <defs></defs>
                <g>
                  <path fill="#FF6200"
                    d="M11.99 2C6.47 2 2 6.48 2 12s4.47 10 9.99 10C17.52 22 22 17.52 22 12S17.52 2 11.99 2zm4.24 16L12 15.45 7.77 18l1.12-4.81-3.73-3.23 4.92-.42L12 5l1.92 4.53 4.92.42-3.73 3.23L16.23 18z">
                  </path>
                </g>
              </svg>
            </div>
          </div>

        </div>
        <div class="title-r">
          <div class="run-btn" @click="goRun">
            <svg xmlns="http://www.w3.org/2000/svg"
              class="styles__StyledSVGIconPathComponent-sc-i3aj97-0 bxMexi svg-icon-path-icon fill" viewBox="0 0 48 48"
              width="18" height="18">
              <defs></defs>
              <g>
                <path d="M24 44C12.9543 44 4 35.0457 4 24C4 12.9543 12.9543 4 24 4C35.0457 4 44 12.9543 44 24"
                  fill="none" stroke="rgb(255, 255, 255)" stroke-width="4" stroke-linecap="round"
                  stroke-linejoin="round"></path>
                <path d="M20 24V17.0718L26 20.5359L32 24L26 27.4641L20 30.9282V24Z" fill="none"
                  stroke="rgb(255, 255, 255)" stroke-width="4" stroke-linejoin="round"></path>
                <path d="M37.0508 32L37.0508 42" fill="none" stroke="rgb(255, 255, 255)" stroke-width="4"
                  stroke-linecap="round" stroke-linejoin="round"></path>
                <path d="M42 36.9497L32 36.9497" fill="none" stroke="rgb(255, 255, 255)" stroke-width="4"
                  stroke-linecap="round" stroke-linejoin="round"></path>
              </g>
            </svg>
            {{ $t('taskTmplObj.run') }}
          </div>
          <div class="count-items">
            <div class="l" style="cursor:pointer;" :title="isCollected ? $t('unStar') : $t('star')"
              @click="changeCollection">
              <svg v-if="isCollected" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 48 48" width="18" height="18">
                <defs></defs>
                <g>
                  <rect width="48" height="48" fill="white" fill-opacity="0.01"></rect>
                  <path
                    d="M15 8C8.92487 8 4 12.9249 4 19C4 30 17 40 24 42.3262C31 40 44 30 44 19C44 12.9249 39.0751 8 33 8C29.2797 8 25.9907 9.8469 24 12.6738C22.0093 9.8469 18.7203 8 15 8Z"
                    fill="rgb(250, 140, 22)" stroke="rgb(250, 140, 22)" stroke-width="4" stroke-linecap="round"
                    stroke-linejoin="round"></path>
                </g>
              </svg>
              <svg v-else xmlns="http://www.w3.org/2000/svg" viewBox="0 0 48 48" width="18" height="18">
                <defs></defs>
                <g>
                  <rect width="48" height="48" fill="white" fill-opacity="0.01"></rect>
                  <path
                    d="M15 8C8.92487 8 4 12.9249 4 19C4 30 17 40 24 42.3262C31 40 44 30 44 19C44 12.9249 39.0751 8 33 8C29.2797 8 25.9907 9.8469 24 12.6738C22.0093 9.8469 18.7203 8 15 8Z"
                    fill="none" stroke="rgb(136, 136, 136)" stroke-width="4" stroke-linecap="round"
                    stroke-linejoin="round"></path>
                </g>
              </svg>
              <span class="tit">{{ $t('star') }}</span>
            </div>
            <div class="r"><span class="num">{{ collectedCount }}</span></div>
          </div>
          <div class="count-items">
            <div class="l"><svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 48 48" width="18" height="18">
                <defs></defs>
                <g>
                  <rect width="48" height="48" fill="white" fill-opacity="0.01"></rect>
                  <path
                    d="M43.8233 25.2305C43.7019 25.9889 43.5195 26.727 43.2814 27.4395C42.763 28.9914 41.9801 30.4222 40.9863 31.6785C38.4222 34.9201 34.454 37 30 37H16C9.39697 37 4 31.6785 4 25C4 18.3502 9.39624 13 16 13L44 13"
                    fill="none" stroke="rgb(136, 136, 136)" stroke-width="4" stroke-linecap="round"
                    stroke-linejoin="round"></path>
                  <path d="M38 7L44 13L38 19" fill="none" stroke="rgb(136, 136, 136)" stroke-width="4"
                    stroke-linecap="round" stroke-linejoin="round"></path>
                </g>
              </svg>
              <span class="tit">{{ $t('taskTmplObj.runTimes') }}</span>
            </div>
            <div class="r"><span class="num">{{ data.UseCount }}</span></div>
          </div>
        </div>
      </div>
      <div class="labels">
        <div class="label" v-if="data.JobType">{{ data.JobTypeStr }}</div>
        <div class="label" v-if="data.ComputeSource">{{ data.ComputeSourceStr }}</div>
        <div class="label" v-for="tag in (data.Tags || [])" :key="tag"> {{ tag }}</div>
      </div>
      <div class="tabs">
        <div class="tab focus">
          <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 48 48" width="16" height="16">
            <defs></defs>
            <g>
              <path
                d="M39 6H9C7.34315 6 6 7.34315 6 9V39C6 40.6569 7.34315 42 9 42H39C40.6569 42 42 40.6569 42 39V9C42 7.34315 40.6569 6 39 6Z"
                fill="none" stroke="rgb(16, 16, 16)" stroke-width="4" stroke-linejoin="round"></path>
              <path d="M26 24H14" stroke="rgb(16, 16, 16)" stroke-width="4" stroke-linecap="round"
                stroke-linejoin="round"></path>
              <path d="M34 15H14" stroke="rgb(16, 16, 16)" stroke-width="4" stroke-linecap="round"
                stroke-linejoin="round"></path>
              <path d="M32 33H14" stroke="rgb(16, 16, 16)" stroke-width="4" stroke-linecap="round"
                stroke-linejoin="round"></path>
            </g>
          </svg>
          {{ $t('taskTmplObj.tmplCoreElements') }}
        </div>
        <a class="tab edit" v-if="data.CanEdit" :href="`/ai_task_tmpl/edit/${data.ID}`">
          <svg xmlns="http://www.w3.org/2000/svg" fill="rgb(16, 16, 16)" viewBox="0 0 32 32" width="16" height="16">
            <defs></defs>
            <g>
              <path
                d="M22.343 4l-2.667 2.667h-13.009v18.667h18.667v-13.009l2.667-2.667v17.009c0 0.736-0.597 1.333-1.333 1.333v0h-21.333c-0.736 0-1.333-0.597-1.333-1.333v0-21.333c0-0.736 0.597-1.333 1.333-1.333v0h17.009zM27.313 2.8l1.887 1.888-12.256 12.256-1.883 0.004-0.003-1.889 12.255-12.259z">
              </path>
            </g>
          </svg>
          {{ $t('taskTmplObj.editTmpl') }}
        </a>
      </div>
    </div>
  </div>
</template>

<script>
import { putCollectAiTaskTmpl, deleteCollectAiTaskTmpl } from '~/apis/modules/aitasktmpl';

export default {
  name: "Header",
  props: {
    data: { type: Object, default: () => ({}) }
  },
  components: {},
  data() {
    return {
      isCollected: false,
      collectedCount: '',
      isSetting: false,
    };
  },
  watch: {
    data: {
      immediate: true,
      deep: true,
      handler(newVal) {
        this.updateData();
      }
    },
  },
  methods: {
    updateData() {
      this.isCollected = this.data.IsCollected;
      this.collectedCount = this.data.NumCollections === undefined ? '' : Math.max(0, this.data.NumCollections);
    },
    goRun() {
      window.location.href = `/cloudbrains/create?tmpl=${this.data.ID}`;
    },
    goEdit() {
      window.location.href = `/ai_task_tmpl/edit/${this.data.ID}`;
    },
    changeCollection() {
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
        id: this.data.ID,
      }).then(res => {
        this.isSetting = false;
        if (res.data.code == '0') {
          this.isCollected = !this.isCollected;
          this.collectedCount = Math.max((this.collectedCount + (this.isCollected ? 1 : -1)), 0);
          this.$message.success(this.isCollected ? this.$t('datasets.starSuccess') : this.$t('datasets.unstarSuccess'));
        } else {
          this.$message.error(res.data.msg);
        }
      }).catch(err => {
        console.log(err);
        this.$message.error(err);
        this.isSetting = false;
      });
    },
  },
  mounted() {
    this.updateData();
  },
  beforeMount() { },
};
</script>

<style scoped lang="less">
.bg {
  background-color: rgba(245, 245, 245, 1);
  border-bottom: 1px solid rgba(224, 227, 234, 1);
  box-sizing: border-box;
}

.title-c {
  display: flex;
  align-items: center;
  justify-content: space-between;
  padding-top: 30px;
  flex-wrap: wrap;
  gap: 10px;

  .title-l {
    display: flex;
    align-items: center;
    flex: 1;
    // width: 0;

    .icon-c {
      flex-shrink: 0;
      height: 40px;
      width: 40px;
      display: flex;
      align-items: center;
      justify-content: center;
      position: relative;
      margin-right: 10px;

      .lock {
        position: absolute;
        right: 4px;
        bottom: 9px;
        background: white;
        border-radius: 4px;
      }
    }

    .title {
      color: rgba(16, 16, 16, 1);
      font-size: 18px;
      font-family: Arial;
      font-weight: 700;
      overflow: hidden;
      text-overflow: ellipsis;
      white-space: nowrap;
      flex: 1;
      width: 0;
      line-height: 28px;
      display: flex;
    }

    .reconmend-icon {
      display: flex;
      align-items: center;
      margin-left: 10px;
    }
  }

  .title-r {
    display: flex;
    align-items: center;
    flex-shrink: 0;

    .run-btn {
      height: 30px;
      display: flex;
      align-items: center;
      background: rgb(64, 169, 255);
      border-radius: 4px;
      color: rgb(255, 255, 255);
      padding: 0 8px;
      cursor: pointer;

      svg {
        margin-right: 5px;
      }
    }

    .count-items {
      height: 30px;
      display: flex;
      align-items: center;
      border-radius: 4px;
      border-color: rgb(225, 227, 230);
      border-style: solid;
      border-width: 1px;
      background: rgb(247, 247, 247);
      margin-left: 12px;

      .l {
        height: 100%;
        display: flex;
        align-items: center;
        border-right: 1px solid rgb(225, 227, 230);
        padding: 0 8px;

        svg {
          margin-right: 4px;
        }
      }

      .r {
        height: 100%;
        display: flex;
        align-items: center;
        padding: 0 10px;
      }
    }
  }
}

.labels {
  display: flex;
  align-items: center;
  margin-top: 22px;
  flex-wrap: wrap;

  .label {
    display: flex;
    align-items: center;
    height: 22px;
    line-height: 17px;
    border-radius: 3px;
    background-color: rgba(219, 229, 255, 1);
    color: rgba(0, 102, 255, 1);
    font-size: 12px;
    padding: 0 8px;
    margin-right: 6px;
    margin-bottom: 4px;
  }
}

.tabs {
  display: flex;
  align-items: center;
  justify-content: space-between;
  margin-top: 12px;
  margin-bottom: -1px;

  .tab {
    display: flex;
    align-items: center;
    color: rgba(16, 16, 16, 1);
    font-size: 14px;
    height: 40px;
    padding: 0 10px;


    svg {
      margin-right: 6px;
    }

    &.focus {
      border: 1px solid rgba(16, 16, 16, 0.1);
      border-radius: 5px 5px 0px 0px;
      background-color: rgba(249, 249, 249, 1);
      border-bottom: 1px solid rgba(249, 249, 249, 1);
    }

    &.edit {
      cursor: pointer;
    }
  }
}
</style>
