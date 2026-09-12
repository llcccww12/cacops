<template>
  <div class="label-select-c" ref="labelSelectRef" :class="isExpand ? 'expand' : ''" v-show="list.length > 1">
    <div class="label-item" :class="selected == index ? 'active' : ''" v-for="(item, index) in list" :key="index"
      @click="change(item, index)">
      {{ item.v }}
    </div>
    <div class="expand-btn" v-show="showExpand">
      <span v-if="!isExpand" @click="isExpand = true"><i class="ri-arrow-down-s-line"></i>{{ $t('org.unfold') }}</span>
      <span v-else @click="isExpand = false"><i class="ri-arrow-up-s-line"></i>{{ $t('org.fold') }}</span>
    </div>
  </div>
</template>

<script>
export default {
  name: "LabelSelect",
  props: {
    type: { type: String, default: "repo" }, // 'repo|model|dataset'
    pageData: { type: Object, default: () => { } },
    labels: { type: Array, default: () => [] },
  },
  data() {
    return {
      list: [],
      selected: 0,
      showExpand: false,
      isExpand: false,
    };
  },
  watch: {
    labels: {
      handler(val, oVal) {
        if (this.list.length || val.length===0) return;
        this.getLabels(val);
      }
    }
  },
  methods: {
    getLabels(newLabels) {
      if (!newLabels) return;
      const _labels = [{
        k: '',
        v: this.$t('all')
      }];
      const labels = newLabels.map(item => {
        return { ...item };
      });
      this.list = [..._labels, ...labels];
      this.checkExpandStatus();
    },
    change(item, index) {
      this.selected = index;
      this.$emit('change', item.k);
    },
    checkExpandStatus() {
      this.showExpand = false;
      this.isExpand = false;
      const labelSelectRef = this.$refs.labelSelectRef;
      this.$nextTick(() => {
        this.showExpand = labelSelectRef.offsetHeight < labelSelectRef.scrollHeight;
      });
    },
  },
  mounted() {
    // this.getLabels();
    window.addEventListener('resize', this.checkExpandStatus);
  },
  beforeDestroy() {
    window.removeEventListener('resize', this.checkExpandStatus);
  }
};
</script>

<style scoped lang="less">
.label-select-c {
  display: flex;
  flex-wrap: wrap;
  position: relative;
  max-height: 74px;
  overflow: hidden;
  padding-right: 56px;
  margin: 16px 0;

  &.expand {
    max-height: none;
  }

  .label-item {
    border: 1px solid #e8e8e8;
    border-radius: 4px;
    color: #415058;
    font-family: Microsoft Yahei;
    font-size: 14px;
    padding: .3em .5em;
    height: 30px;
    text-align: center;
    margin: .2em;
    max-width: 100%;
    display: inline-flex;
    cursor: pointer;

    &.active {
      background-color: #0366d6 !important;
      color: #fff !important;
    }

    &:hover {
      color: #1e70bf;
    }
  }

  .expand-btn {
    position: absolute;
    right: 0;
    bottom: 5px;
    height: 33px;
    overflow: hidden;

    span {
      margin: 0 5px 0 5px;
      height: 33px;
      display: flex;
      align-items: center;
      justify-content: center;
      color: #0366d6;
      cursor: pointer;
      background-color: white;

      i {
        margin-right: 4px;
      }
    }
  }
}
</style>
