<template>
  <div class="dataset-select">
    <div class="title">
      <span :class="required ? 'required' : ''">{{ title }}</span>
    </div>
    <div class="content" :class="errStatus ? 'error' : ''">
      <div class="tab-item-c">
        <div class="tab-item" :class="tabIndex == 0 ? 'tab-item-focus' : ''" @click="changeTab(0)">示例数据集</div>
        <div class="tab-item" :class="tabIndex == 1 ? 'tab-item-focus' : ''" @click="changeTab(1)">本地上传</div>
        <div class="tab-item" :class="tabIndex == 2 ? 'tab-item-focus' : ''" @click="changeTab(2)">选择平台数据集</div>
      </div>
      <div class="tab-content-c">
        <div v-show="tabIndex == 0" class="tab-content">
          <div style="display:flex;align-items:center;height:34px;">
            <el-radio-group v-model="sampleDataset">
              <el-radio :label="1">文本分类数据集</el-radio><a class="download-btn" download
                href="https://openi.pcl.ac.cn/attachments/cc5f7d3f-8be0-482a-8982-ae6c2f0d677b?type=1">(下载)</a>
              <el-radio :label="2">中英翻译数据集</el-radio><a class="download-btn" download
                href="https://openi.pcl.ac.cn/attachments/f5eb1608-d3cb-4cbe-8df1-8eb0903e0edf?type=1">(下载)</a>
              <el-radio :label="3">情感分类数据集</el-radio><a class="download-btn" download
                href="https://openi.pcl.ac.cn/attachments/e9eb36c5-1eea-4138-9de2-7d2ddcad19d9?type=1">(下载)</a>
            </el-radio-group>
          </div>
        </div>
        <div v-show="tabIndex == 1" class="tab-content">
          <div>
            <DatasetFileUploader :datasetId="datasetId" ref="datasetFileUploaderRef" :type="1"></DatasetFileUploader>
          </div>
        </div>
        <div v-show="tabIndex == 2" class="tab-content">
          <div class="platform-dataset">
            <DatasetSelect ref="datasetSelectRef" :userName="userName" :repoName="repoName" :showTitle="false" :type="1"
              :maxCount="1">
            </DatasetSelect>
          </div>
        </div>
      </div>
      <div class="tips-how" v-if="tabIndex != 0">
        <a target="_blank" href="https://openi.pcl.ac.cn/PCL-Platform.Intelligence/pcl_pangu/src/branch/master/docs/README_DATASET.md">如何构建数据集</a>
      </div>
    </div>
  </div>
</template>

<script>
import DatasetFileUploader from './DatasetFileUploader.vue';
import DatasetSelect from './cloudbrain/DatasetSelect.vue';

export default {
  name: "ModelBaseDatasetSelect",
  props: {
    title: { type: String, default: "数据集" },
    required: { type: Boolean, default: true },
    oriData: { type: Array, default: () => [] },
    tabindex: { type: Number, default: 0 },
    exampleType: { type: Number, default: 1 },
    userName: { type: String, default: '' },
    repoName: { type: String, default: '' },
    datasetId: { type: String, default: '' },
  },
  components: { DatasetFileUploader, DatasetSelect },
  data() {
    return {
      tabIndex: 0,

      sampleDataset: 1,
      errStatus: false,
    };
  },
  watch: {
    tabindex: {
      handler(newVal) {
        this.tabIndex = newVal;
      },
      immediate: true,
    },
    exampleType: {
      handler(newVal) {
        this.sampleDataset = newVal;
      },
      immediate: true,
    },
    oriData: function (val) {
      this.name = val;
    },
  },
  methods: {
    changeTab(tab) {
      this.tabIndex = tab;
    },
    getPlatformDataset() {
      return [...this.$refs.datasetSelectRef.selectList];
    },
    check() {
      if (this.tabIndex == 0 && this.sampleDataset >= 1 && this.sampleDataset <= 3) {
        return true;
      }
      if (this.tabIndex == 1) {
        return this.$refs.datasetFileUploaderRef.check()
      }
      if (this.tabIndex == 2) {
        return this.$refs.datasetSelectRef.check()
      }
      return false;
    },
    localUpload() {
      return this.$refs.datasetFileUploaderRef.upload();
    }
  },
};
</script>

<style scoped lang="less">
.dataset-select {
  display: flex;
  margin-bottom: 28px;

  .title {
    width: 200px;
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
      padding: 20px 16px;

      .tab-content {
        min-height: 34px;
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
      position: relative;
      display: inline-block;
      font-size: 12px;
      margin-left: -20px;
      margin-right: 30px;
    }
  }
}

.platform-dataset {
  .dataset-select {
    margin-bottom: 0;
  }
}</style>
