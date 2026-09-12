<template>
  <div class="operation-group-c" :key="usekey">
    <div ref="btnsRef" class="group-line">
      <slot></slot>
    </div>
    <el-popover v-if="showMore" placement="bottom-end" width="auto" trigger="click" popper-class="group-btn-popup-more">
      <div slot="reference">
        <a class="el-dropdown-link el-dropdown-selfdefine" aria-haspopup="list" role="button" tabindex="0">{{
          $t('cloudbrainObj.more') }}
          <i class="el-icon-arrow-down el-icon--right"></i>
        </a>
      </div>
      <div ref="moreRef" class="operation-group-c dorpdown">
        <slot></slot>
      </div>
    </el-popover>
  </div>
</template>

<script>
export default {
  name: 'OperationBtnGroup',
  props: {
    showLen: { type: Number, default: 3 },
    usekey: { type: String | String, required: true },
  },
  data() {
    return {
      showMore: false,
    };
  },
  components: {},
  watch: {
    showLen() {
      this.$nextTick(() => {
        this.refresh();
      });
    },
  },
  methods: {
    refresh() {
      this.$nextTick(() => {
        this.$refs.btnsRef.children.forEach(ele => ele.classList.remove('group-btn-hide'));
        const showBtns = [].filter.call(this.$refs.btnsRef.children, ele => (ele.computedStyleMap().get('display').value != 'none'));
        if (showBtns.length > this.showLen) {
          this.showMore = true;
          showBtns.forEach((ele, index) => {
            if (index >= this.showLen) {
              ele.classList.add('group-btn-hide');
            } else {
              ele.classList.remove('group-btn-hide');
            }
          });
          this.$nextTick(() => {
            this.$refs.moreRef.children.forEach(ele => ele.classList.remove('group-btn-hide'));
            const showBtns = [].filter.call(this.$refs.moreRef.children, ele => (ele.computedStyleMap().get('display').value != 'none'));
            showBtns.forEach((ele, index) => {
              if (index < this.showLen) {
                ele.classList.add('group-btn-hide');
              } else {
                ele.classList.remove('group-btn-hide');
              }
            });
          });
        } else {
          this.showMore = false;
        }
      });
    }
  },
  beforeMount() { },
  beforeUpdate() {
    this.refresh();
  },
  mounted() {
    this.refresh();
    window.addEventListener('resize', this.refresh);
  },
  beforeDestroy() {
    window.removeEventListener('resize', this.refresh);
  },
};
</script>

<style scoped lang="less">

.operation-group-c {
  display: flex;
  align-items: center;
  justify-content: center;
  flex-wrap: wrap;
  gap: 12px;
  .group-line {
    display: flex;
    align-items: center;
    flex-wrap: wrap;
    a{
      color: #005cff;
    }
  }

  /deep/ a {
    margin: 0 10px;
    font-size: 14px;
    

    &.disabled {
      color: rgba(0, 0, 0, .87);
    }

    &.delete {
      color: red;
    }

    &.group-btn-hide {
      display: none;
    }

    &.el-dropdown-link {
      color: rgba(0, 0, 0, .87);
      display: flex;
      align-items: center;
    }
  }

  .el-dropdown {
    margin: 0 10px;
    color: rgba(0, 0, 0, .87);
    cursor: pointer;
  }

  &.dorpdown {
    flex-direction: column;
    align-items: baseline;

    /deep/ a {
      line-height: 36px;
      width: 100%;
      margin: 0;
      padding: 0 20px;
      overflow: hidden;
      height: 36px;
      text-overflow: ellipsis;
      white-space: nowrap;

      &.group-btn-hide {
        display: none;
      }

      &:hover {
        background-color: #ecf5ff;
      }
    }
  }
}
</style>
<style lang="less">
.group-btn-popup-more {
  padding: 0 !important;
}
</style>
