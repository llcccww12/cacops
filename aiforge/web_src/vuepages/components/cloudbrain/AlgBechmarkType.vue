<template>
  <div class="form-row">
    <div class="left-area">
      <div class="title align-items-center" v-if="showTitle">
        <span :class="required ? 'required' : ''">{{ '评测类型' }}</span>
      </div>
      <div class="content" :class="errStatus ? 'error' : ''">
        <div class="content-l">
          <el-select class="field-input" v-model="mainType" @change="handleInputChange">
            <el-option v-for="item in mainTypeList" :key="item.key" :value="item.key" :label="item.value"></el-option>
          </el-select>
        </div>
        <div class="content-r">
          <div class="title align-items-center"><span class="required">{{ '子类型' }}</span></div>
          <el-select class="field-input" v-model="childType" @change="handleInputChange">
            <el-option v-for="item in childTypeList" :key="item.key" :value="item.key" :label="item.value"></el-option>
          </el-select>
        </div>
      </div>
    </div>
    <div class="right-area"></div>
  </div>
</template>

<script>

export default {
  name: "AlgBechmarkType",
  props: {
    value: { type: Array, required: true },
    required: { type: Boolean, default: true },
    title: { type: String, default: '' },
    showTitle: { type: Boolean, default: true },
  },
  data() {
    return {
      userName: location.pathname.split('/')[1],
      repoName: location.pathname.split('/')[2],

      mainType: '',
      mainTypeList: [{ key: '1', value: 'Object Detection' }, { key: '2', value: 'Object ReID' }, { key: '3', value: 'Multi-Object-Tracking' }],

      childType: '',
      childTypeList: [],

      sampleUrl: 'https://openi.pcl.ac.cn/OpenIOSSG/MNIST_PytorchExample_GPU',
      errStatus: false,
    };
  },
  watch: {
    value: {
      immediate: true,
      deep: true,
      handler(newVal) {
        const [mainType, childType] = (newVal || []);
        this.mainType = mainType || '';
        this.childType = childType || '';
      }
    }
  },
  methods: {
    check() {

      return !this.errStatus;
    },
    handleInput(value) {
      this.currentValue = value;
      this.$emit('input', value);
    },
    handleInputChange(value) {
      this.$emit('change', value);
      this.check();
    },
  },
  beforeMount() { },
  mounted() { },
};
</script>

<style scoped lang="less">
@import 'cloudbrain.less';

.form-row {
  .content {
    display: flex;

    .content-l {
      display: flex;
      flex: 4;
      margin-right: 15px;
    }

    .content-r {
      display: flex;
      flex: 5;

      .title {
        width: auto;
      }
    }
  }
}
</style>
