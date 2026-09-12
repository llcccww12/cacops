<template>
  <div class="form-row">
    <div class="left-area">
      <div class="title align-items-center" v-if="showTitle"><span class="required">{{ 'AI引擎' }}</span></div>
      <div class="content">
        <el-select class="engine-type-sel field-input" v-model="engineType" @change="handleEngineTypeChange">
          <el-option v-for="item in engineTypeList" :key="item.id" :value="item.id" :label="item.name"></el-option>
        </el-select>
        <el-select class="engine-sel field-input" v-model="engine" @change="handleEngineChange">
          <el-option v-for="item in engineList" :key="item.id" :value="item.id" :label="item.name"></el-option>
        </el-select>
      </div>
    </div>
    <div class="right-area"></div>
  </div>
</template>

<script>

export default {
  name: "AIEngineSelect",
  props: {
    value: { type: Array, required: true },
    required: { type: Boolean, default: true },
    title: { type: String, default: '' },
    engines: { type: Array, required: true },
    showTitle: { type: Boolean, default: true },
  },
  data() {
    return {
      engineType: '',
      engine: '',
      errStatus: false,
    };
  },
  computed: {
    engineTypeList() {
      return this.engines;
    },
    engineList() {
      const selectedEngineType = this.engines.find((item) => item.id == this.engineType);
      if (selectedEngineType) {
        return selectedEngineType.engines || [];
      } else {
        return [];
      }
      return [];
    },
  },
  watch: {
    value: {
      immediate: true,
      deep: true,
      handler(newVal) {
        const [engineType, engine] = (newVal || []);
        this.engineType = engineType || '';
        this.engine = engine || '';
      }
    }
  },
  methods: {
    handleEngineTypeChange(item) {
      this.engine = '';
      const selectedEngineType = this.engines.find((item) => item.id == this.engineType);
      let engines = [];
      if (selectedEngineType) {
        engines = selectedEngineType.engines || [];
      }
      if (engines.length) {
        this.engine = engines[0].id || '';
      }
      this.$nextTick(() => {
        this.$emit('input', [this.engineType, this.engine]);
        this.$emit('change', [this.engineType, this.engine]);
      });
    },
    handleEngineChange(item) {
      this.$emit('input', [this.engineType, this.engine]);
      this.$emit('change', [this.engineType, this.engine]);
    },
    check() { },
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

    .engine-type-sel {
      flex: 4;
      margin-right: 10px;
    }

    .engine-sel {
      flex: 6;
    }
  }
}
</style>
