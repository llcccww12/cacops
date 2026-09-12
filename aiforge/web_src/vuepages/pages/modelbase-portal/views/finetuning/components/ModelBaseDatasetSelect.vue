<template>
  <div class="dataset-select">
    <div class="title">
      <span :class="required ? 'required' : ''">{{ $t('datasetObj.dataset_label') }}</span>
    </div>
    <div class="content" :class="errStatus ? 'error' : ''">
      <div class="tab-item-c">
        <div class="tab-item" :class="tabIndex == 0 ? 'tab-item-focus' : ''" @click="changeTab(0)">{{ $t('datasetObj.exampleDataset') }}</div>
        <div class="tab-item" :class="tabIndex == 1 ? 'tab-item-focus' : ''" @click="changeTab(1)">{{ $t('datasetObj.selectOpeniDataset') }}</div>
      </div>
      <div class="tab-content-c">
        <div v-show="tabIndex == 0" class="tab-content">
          <div style="display:flex;align-items:center;">
            <el-radio-group v-model="sampleDataset" class="radio-wrap">
              <el-radio v-for="item in datasetList" :label="item.id" :key="item.name">
                {{item.name}}
                <span class="download-btn" @click="downloadAllModel(item)">({{$t('modelManage.download')}})</span>
              </el-radio>
              
            </el-radio-group>
          </div>
        </div>
        <div v-show="tabIndex == 1" class="tab-content">
          <div class="platform-dataset">
            <DatasetSelect ref="datasetRef" v-model="selectDatasetList" :multiple="true" :showTitle="false" :maxCount="datasetCount"></DatasetSelect>
          </div>
        </div>
      </div>
      <div class="tips-how">
        <a target="_blank" href="https://openi.pcl.ac.cn/OpenIOSSG/OpenI_LLM_Finetune_Example/src/branch/master/data/dataset_prepare.md">{{$t('datasetObj.buildDatasetTips')}}</a>
      </div>
    </div>
  </div>
</template>

<script>

import DatasetSelect from '~/components/cloudbrain/DatasetSelectV2.vue';
import { getCurrentRepoDataset } from '~/apis/modules/dataset';
export default {
  name: "ModelBaseDatasetSelect",
  props: {
    title: { type: String, default: "数据集" },
    required: { type: Boolean, default: true },
    oriData: { type: Array, default: () => [] },
    tabindex: { type: Number, default: 0 },
    exampleType: { type: Number, default: 1 },
    datasetId: { type: String, default: '' },
    datasetCount: { type: Number, default: 0 },
  },
  components: { DatasetSelect },
  data() {
    return {
      tabIndex: 0,
      datasetList: [],
      sampleDataset: '',
      errStatus: false,
      currentRepoDatasetList: [],
      selectDatasetList: [],
    };
  },
  watch: {
    tabindex: {
      handler(newVal) {
        this.tabIndex = newVal;
      },
      immediate: true,
    },
    oriData: function (val) {
      this.sampleDataset = val[0].id
      this.datasetList = []
      val.forEach((item)=>{
        this.datasetList.push(item)
      })

    },
  },
  methods: {
    changeTab(tab) {
      this.tabIndex = tab;
    },
    getPlatformDataset() {
      if(this.tabIndex===0){
        const selectDataSet = this.datasetList.filter((item)=>{
          return item.id === this.sampleDataset
        })
        return selectDataSet
      }else{
        return this.selectDatasetList;
      }
    },
    downloadAllModel(item){
      let downloadElement = document.createElement('a')
      let href = `/api/v1/dataset/download?dataset_id=${item.id}`
      downloadElement.href = href
      document.body.appendChild(downloadElement)
      downloadElement.click() //点击下载
      document.body.removeChild(downloadElement) //下载完成移除元素
    },

    check() {
      if (this.tabIndex == 0) {
        return true;
      }
      if(this.tabindex === 1){
        if(!this.$refs.datasetSelectRef.check()){
          const parent = document.querySelector('.main-wrap')
          const child = document.querySelector('.dataset-select')
          parent.scrollTo({
            top:child.offsetTop-100,
            behavior:'smooth'
          })
          return false
        }
        return true
      }
      return true
    },
    localUpload() {
      return false;
    }
  },
  mounted() { 
    
  }
};
</script>

<style scoped lang="less">
.dataset-select {
  display: flex;
  margin-bottom: 28px;

  .title {
    width: 130px;
    text-align: right;
    margin-right: 24px;
    color: #101010;
    font-size: 14px;
    display: flex;
    justify-content: flex-end;
    padding-top: 8px;

    .required {
      position: relative;

      &::after {
        position: absolute;
        content: "*";
        top: -3px;
        right: -10px;
        color: red;
      }
    }
  }

  .content {
    flex: 1;
    margin-right: 128px;
    .tab-item-c {
      display: flex;
      .tab-item {
        border: 1px solid #DCDFE6;
        display: flex;
        justify-content: center;
        align-items: center;
        height: 36px;
        padding: 0 20px;
        cursor: pointer;
        border-right: none;
        color: rgb(96, 98, 102);
        background: rgb(245, 245, 246);

        &:first-child {
          border-radius: 5px 0px 0px 0px;
        }

        &:last-child {
          border-right: 1px solid #DCDFE6;
          border-radius: 0px 5px 0px 0px;
        }

        &.tab-item-focus {
          color: rgb(50, 145, 248);
          background: rgb(255, 255, 255);
          border-bottom: none;
        }
      }
    }

    .tab-content-c {
      margin-top: -1px;
      border: 1px solid #DCDFE6;
      border-radius: 0px 5px 5px 5px;
      padding: 10px 16px;

      .tab-content {
        display: flex;
        min-height: 34px;
        .radio-wrap{
          display: flex;
          flex-wrap: wrap;
        }
        .platform-dataset{
          width: 100%;
          .dataset-select{
            margin-bottom: 0px;
          }
        }
      }
    }

    .tips-how {
      margin-top: 5px;
      padding-left: 4px;

      a {
        text-decoration: underline;
      }
    }

    .download-btn {
      color: rgb(0, 102, 255);
      cursor: pointer;
      position: relative;
      display: inline-block;
      font-size: 12px;
    }
  }
}

@media (max-width: 768px) {
  .dataset-select {
    flex-direction: column;
    margin-bottom: 20px; // 减少间距
    .title {
      width: 100%; // 宽度占满
      text-align: left; // 左对齐
      margin-right: 0; // 去掉右边距
      margin-bottom: 8px; // 增加下边距
      padding-top: 0; // 去掉上边距
      justify-content: flex-start; // 左对齐
      font-size: 13px; // 缩小字体
    }

    .content {
      margin-right: 0;
      .radio-wrap{
        gap: 14px;
        .el-radio{
          margin-right: 0px;
        }
      }
    }
  }

  
}

</style>
