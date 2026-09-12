<template>
  <div class="container">
    <div class="title">{{ $t('repos.model') }}</div>
    <div class="block-c">
      <div class="title-c">
        <span class="title">{{ $t('modelManage.modelEngine') }}</span>
        <span class="clear-btn" v-if="engineFlag" @click="clearSelectLeft('engine')">Clear</span>
      </div>
      <div class="list engine-c">
        <div class="item" :class="item.active ? 'active' : ''" v-for="(item, index) in Engine" :key="item.k"
          @click="selectEngine(item)">{{ item.v }}</div>
      </div>
    </div>
    <div class="block-c">
      <div class="title-c">
        <span class="title">{{ $t('modelManage.label') }}</span>
        <span class="clear-btn" v-if="labelFlag" @click="clearSelectLeft('label')">Clear</span>
      </div>
      <div class="list label-c">
        <div class="item" :class="item.active ? 'active' : ''" v-for="(item, index) in Label" :key="item.k"
          @click="selectLabel(item)">{{ item.v }}</div>
      </div>
    </div>
  </div>
</template>

<script>
import { getModelSqaureFilters } from '~/apis/modules/modelsquare';

export default {
  name: "ModelFilters",
  props: {
    condition: { type: Object, default: () => ({}) },
  },
  components: {},
  data() {
    return {
      engineFlag: false,
      labelFlag: false,
      Engine: [],
      Label: [],
      engineValue: '',
      labelValue: ''
    };
  },
  methods: {
    selectEngine(item) {
      this.Engine.forEach(element => {
        if (element.k == item.k) {
          element.active = true
        } else {
          element.active = false
        }
      });
      this.engineFlag = true
      this.engineValue = item.k
      this.search()
    },
    selectLabel(item) {
      this.Label.forEach(element => {
        if (element.k == item.k) {
          element.active = true
        } else {
          element.active = false
        }
      });
      this.labelValue = item.k
      this.labelFlag = true
      this.search()
    },
    clearSelectLeft(type) {
      if (type === 'engine') {
        this.Engine.forEach(element => {
          element.active = false
        });
        this.engineValue = ''
        this.engineFlag = false
      } else if (type === 'label') {
        this.Label.forEach(element => {
          element.active = false
        });
        this.labelValue = ''
        this.labelFlag = false
      }
      this.search()
    },
    search() {
      this.$emit('changeCondition', {
        engine: this.engineValue,
        label: this.labelValue
      });
    }
  },
  beforeMount() {
    getModelSqaureFilters().then(res => {
      const oriData = res.data;
      if (oriData && oriData[0]) {
        const data = {
          engine: oriData[0].frame.split(',').map(item => { return { k: item.split(':')[0], v: item.split(':')[1], active: false } }),
          label: oriData[0].label.split(',').map(item => { return { k: item, v: item, active: false } }),
        };
        this.Engine = data.engine;
        this.Label = data.label;
      }
    }).catch(err => {
      console.log(err);
    });
  },
  mounted() { },
};
</script>

<style scoped lang="less">
.container {

  >.title {
    color: rgb(16, 16, 16);
    font-size: 18px;
    font-weight: bold;
    margin-bottom: 20px;
  }

  .block-c {
    margin-top: 8px;

    .title-c {
      display: flex;
      align-items: flex-end;
      margin-bottom: 8px;

      .title {
        color: rgb(16, 16, 16);
        font-size: 14px;
      }

      .clear-btn {
        text-decoration: underline;
        font-size: .875rem;
        cursor: pointer;
        font-size: 12px;
        margin-left: 8px;
        color: rgb(156, 163, 175);
      }
    }

    .list {
      display: flex;
      flex-wrap: wrap;

      .item {
        border-radius: 4px;
        font-size: 12px;
        color: rgba(14, 37, 69, 1);
        box-shadow: 0px 1px 1px 0px rgba(16, 16, 16, 0.2);
        margin: 0 10px 10px 0;
        padding: 2px 4px;
        display: flex;
        align-items: center;
        justify-content: center;
        cursor: pointer;
      }

      &.engine-c {
        .item {
          border: 1px solid rgba(145, 213, 255, 0.5);
          background: linear-gradient(92.08deg, rgba(226, 245, 255, 1) -0.83%, rgba(247, 252, 255, 0.5) 19.83%, rgba(247, 252, 255, 0.5) 97.53%);

          &.active {
            background: rgba(255, 255, 255, 1);
            border: 1px solid rgba(0, 158, 255, 1);
          }
        }
      }

      &.label-c {
        .item {
          border: 1px solid rgba(255, 198, 145, 0.5);
          background: linear-gradient(92.08deg, rgba(255, 245, 226, 1) -0.83%, rgba(255, 253, 247, 0.5) 19.83%, rgba(255, 253, 247, 0.5) 97.53%);

          &.active {
            background: rgba(255, 255, 255, 1);
            border: 1px solid rgba(255, 122, 0, 1);
          }
        }
      }
    }
  }
}
</style>
