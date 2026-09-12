<template>
  <div class="model-select">
    <div class="model-top-icon">
      <img v-if="iconSrc" :src="iconSrc" alt="">
      <div class="default-img" v-else>
        <svg xmlns="http://www.w3.org/2000/svg" fill="rgba(16, 16, 16, 0.2)" viewBox="0 0 32 32" width="32" height="32">
          <defs></defs>
          <g>
            <path
              d="M12.992 14.4c0.608-0.096 1.216-0.256 1.856-0.448s1.248-0.448 1.76-0.8 0.992-0.736 1.344-1.248c0.352-0.512 0.512-1.12 0.512-1.888 0-1.088-0.384-1.92-1.184-2.528s-1.856-0.864-3.168-0.864c-0.544 0-1.024 0.032-1.44 0.064-0.448 0.032-0.832 0.096-1.184 0.192s-0.736 0.192-1.12 0.32c-0.352 0.096-0.768 0.256-1.152 0.384-0.349-0.596-0.653-1.285-0.877-2.009l-0.019-0.071c-0.224-0.661-0.373-1.425-0.415-2.218l-0.001-0.022c0.672-0.224 1.28-0.416 1.856-0.608 0.544-0.16 1.12-0.288 1.632-0.384 0.576-0.096 1.088-0.192 1.632-0.224s1.12-0.064 1.76-0.064c3.072 0 5.408 0.704 6.976 2.112s2.368 3.328 2.368 5.792c0 1.184-0.256 2.24-0.704 3.104-0.439 0.873-1.007 1.616-1.69 2.234l-0.006 0.006c-0.586 0.56-1.248 1.049-1.967 1.447l-0.049 0.025c-0.704 0.352-1.248 0.64-1.696 0.832v3.488c-0.896 0.16-1.76 0.256-2.592 0.256s-1.664-0.096-2.432-0.256v-6.624zM12.608 29.76c-0.16-0.96-0.224-1.888-0.224-2.816s0.064-1.888 0.224-2.848c0.96-0.16 1.888-0.256 2.816-0.256s1.888 0.096 2.848 0.256c0.16 0.96 0.256 1.92 0.256 2.816 0 0.96-0.096 1.888-0.256 2.848-0.96 0.16-1.888 0.256-2.816 0.256s-1.888-0.096-2.848-0.256z">
            </path>
          </g>
        </svg>
      </div>
    </div>
    <div class="select-c">
      <el-select class="category-select" v-model="selectCategory" :placeholder="$t('modelObj.model_select')"
        size="default" @change="changeCategory">
        <template slot="prefix">
          <div class="icon">
            <div class="img-c">
              <img v-if="iconSrc" :src="iconSrc" :style="iconSize ? `width:${iconSize}px;height:${iconSize}px;` : ''"
                alt="">
            </div>
          </div>
        </template>
        <el-option v-for="(item, index) in categoryList" :key="item.name" :value="item.name">
          <div class="category-item">
            <div class="img-c">
              <img :src="item.icon" :style="item.iconSize ? `width:${item.iconSize}px;height:${item.iconSize}px;` : ''"
                alt="">
            </div>
            <span>{{ item.name }}</span>
          </div>
        </el-option>
      </el-select>
      <el-select class="model-item-select" v-show="selectCategory" size="default" v-model="selectModel" filterable
        @change="changeModel">
        <el-option v-for="(item, index) in modelList" :key="item.name" :value="item.name"
          :disabled="item.name == disabledModel"></el-option>
      </el-select>
    </div>
  </div>
</template>

<script>

export default {
  props: {
    disabledModel: { type: String, default: '' },
    models: { type: Array, default: () => [] },
    model: { type: Object, default: () => { } },
    category: { type: String, default: '' },
  },
  data() {
    return {
      iconSrc: '',
      iconSize: '',
      categoryList: [],
      selectCategory: '',
      modelList: [],
      selectModel: '',
    };
  },
  components: {},
  watch: {
    models: {
      handler(val, oVal) {
        this.init()
      },
      deep: true,
      immediate: true
    }
  },
  methods: {
    init() {
      const categoryList = []
      for (let i = 0, iLen = this.models.length; i < iLen; i++) {
        const dataI = this.models[i];
        const category = {
          name: dataI.category,
          icon: dataI.icon || '/img/chatbot.png',
          iconSize: dataI.iconSize,
          models: dataI.models,
        }
        categoryList.push(category);
      }
      this.categoryList = categoryList;
      if (this.model.category) {
        this.selectCategory = this.model.category;
        this.changeCategory()
        if (this.model.name) {
          this.selectModel = this.model.name;
          this.changeModel()
        }
      }
    },
    changeCategory() {
      const selected = this.categoryList.find(item => item.name == this.selectCategory)
      this.iconSrc = selected.icon;
      this.iconSize = selected.iconSize;
      this.modelList = selected.models.map(item => {
        return {
          name: item.name,
        }
      })
      this.selectModel = ''
      this.changeModel()
    },
    changeModel() {
      this.$emit('changeModel', {
        category: this.selectCategory,
        name: this.selectModel,
        icon: this.iconSrc
      });
    },
  },
  mounted() { },
  beforeDestroy() { },
};
</script>

<style scoped lang="less">
.model-select {
  height: 100%;
  display: flex;
  align-items: center;
  justify-content: center;
  flex-direction: column;

  .model-top-icon {
    margin-bottom: 28px;
    margin-top: -80px;

    img {
      height: 90px;
      width: 90px;
    }

    .default-img {
      height: 90px;
      width: 90px;
      display: flex;
      align-items: center;
      justify-content: center;
      border-color: rgba(16, 16, 16, 0.2);
      border-style: dashed;
      border-width: 2px;
      border-radius: 100%;
      background: rgba(16, 16, 16, 0.05);
    }
  }

  .category-select {
    margin-right: 10px;
    width: 160px;

    /deep/ .el-input--prefix .el-input__inner {
      padding-left: 40px;
      color: rgb(0, 102, 255);

      &::placeholder {
        color: rgb(0, 102, 255);
      }
    }

    /deep/ .el-input__prefix {
      display: flex;
      align-items: center;
    }

    .icon {
      display: flex;
      align-items: center;

      img {
        height: 30px;
        width: 30px;
      }
    }
  }

  .model-item-select {
    width: 260px;

    /deep/ .el-input__inner {
      color: rgb(0, 102, 255);

      &::placeholder {
        color: rgb(0, 102, 255);
      }
    }
  }
}

.img-c {
  height: 30px;
  width: 30px;
  margin-right: 10px;
  display: flex;
  align-items: center;
  justify-content: center;
}

.category-item {
  display: flex;
  align-items: center;

  img {
    height: 30px;
    width: 30px;
  }

  span {}
}

@media (max-width: 768px) {
  .model-select {
    .select-c {
      display: flex;
      flex-direction: column;
      align-items: center;
      justify-content: center;

      .category-select {
        margin-right: 0;
        margin-bottom: 10px;
      }
    }
  }
}
</style>
