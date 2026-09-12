<template>
  <div class="form-row">
    <div class="left-area">
      <div class="title">
        <span class="required">{{ $t('cloudbrainObj.image') }}</span>
      </div>
      <div class="content" :class="errStatus ? 'error' : ''">
        <el-select class="field-input" v-model="currentValue" @change="handleChange">
          <el-option v-for="(item, index) in images" :key="`${index}-${item.image_id}`" :value="item.image_id"
            :label="item.image_name"></el-option>
        </el-select>
        <div class="dynamic-tips" v-if="tips" v-html="tips"></div>
      </div>
    </div>
    <div class="right-area"></div>
  </div>
</template>

<script>

import { getAiTaskImgesBySpec } from '~/apis/modules/cloudbrain';
export default {
  name: 'ImageSelectV2',
  props: {
    configs: { type: Object, required: true, },
    value: { type: Object, required: true },
    images: { type: Array, required: true, },
    spec: { type: String, required: true, },
    networkType: { type: String, required: true },
    visualizeRequired: { type: Boolean, default: false },
    required: { type: Boolean, default: true },
    repoOwnerName: { type: String, required: true },
    repoName: { type: String, required: true },
    tips: { type: String, default: '' },
  },
  data() {
    return {
      currentValue: '',
      errStatus: false,
    };
  },
  watch: {
    value: {
      immediate: true,
      deep: true,
      handler(newVal) {
        newVal = newVal === undefined ? {} : newVal;
        this.currentValue = (newVal.image_id || '').toString();
      }
    },
    spec: {
      immediate: true,
      handler(newVal) {
        const imagev2Cfg = this.configs?.form?.imagev2;
        if (imagev2Cfg && imagev2Cfg.relatedSpec) {
          if (!newVal) {
            this.$emit('changeImages', []);
          } else {
            let networkType = 0; // all
            if (this.networkType == 'no_internet') {
              networkType = 1;
            } else if (this.networkType == 'has_internet') {
              networkType = 2;
            }
            this.delayTimer && clearTimeout(this.delayTimer);
            this.delayTimer = setTimeout(() => {
              getAiTaskImgesBySpec({
                repoOwnerName: this.repoOwnerName,
                repoName: this.repoName,
                jobType: this.configs.taskType,
                computeSource: this.configs.computerResouce,
                clusterType: this.configs.clusterType,
                spec: newVal,
                hasInternet: networkType,
                visualizeRequired: this.visualizeRequired,
              }).then(res => {
                const data = res.data;
                if (data.code == 0) {
                  this.$emit('changeImages', data?.data?.images || []);
                }
              }).catch(err => {
                console.log(err);
              })
            }, 50);
          }
        }
      }
    },
  },
  methods: {
    check() {
      if (this.required && !this.currentValue) {
        this.errStatus = true;
      } else {
        this.errStatus = false;
      }
      return !this.errStatus;
    },
    handleChange(value) {
      this.currentValue = value;
      const selectImage = this.images.filter((item) => item.image_id == value)[0];
      this.$emit('input', {
        image_id: selectImage.image_id,
        image_name: selectImage.image_name,
      });
      this.$emit('change', {
        image_id: selectImage.image_id,
        image_name: selectImage.image_name,
      });
    },
  },
  mounted() { },
};
</script>

<style scoped lang="less">
@import 'cloudbrain.less';
</style>
