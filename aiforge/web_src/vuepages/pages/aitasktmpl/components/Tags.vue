<template>
  <div class="tags-c" :class="focus ? 'focus' : ''" @click="HandlerFocus">
    <div class="tag" v-for="(tag, index) in tags" :key="tag">{{ tag }} <i class="el-icon-close"
        @click.prevent.stop="removeTag(tag)"></i></div>
    <input ref="newTagInputRef" type="text" v-model="newTag" autocomplete="off" :maxlength="20"
      style="flex-grow: 1;min-width:60px;max-width:90%;" @focus="changeFocus(true)" @blur="changeFocus(false)"
      @keyup.enter="addTag" :placeholder="!tags.length && !newTag ? $t('taskTmplObj.tagsPlaceholder') : ''">
    <!-- <div class="placeholder" v-if="!tags.length && !newTag">{{ $t('taskTmplObj.tagsPlaceholder') }}</div> -->
  </div>
</template>

<script>
export default {
  name: "Tags",
  props: {
    value: { type: Array, default: () => [] }
  },
  components: {},
  data() {
    return {
      maxCount: 10,
      focus: false,
      tags: [],
      newTag: '',
    };
  },
  watch: {
    value: {
      immediate: true,
      handler(newVal) {
        newVal = !!newVal ? newVal : [];
        this.tags = newVal;
      }
    }
  },
  methods: {
    changeFocus(value) {
      this.focus = value;
    },
    HandlerFocus() {
      this.$refs.newTagInputRef.focus();
    },
    removeTag(tag) {
      const index = this.tags.indexOf(tag);
      this.tags.splice(index, 1);
      this.emitChange();
    },
    addTag() {
      this.newTag = this.newTag.trim();
      if (this.newTag) {
        const index = this.tags.indexOf(this.newTag);
        if (index >= 0) {
          this.tags.splice(index, 1);
        } else {
          if (this.tags.length + 1 <= this.maxCount) {
            this.tags.push(this.newTag);
          }
        }
      }
      this.newTag = '';
      this.emitChange();
    },
    emitChange() {
      this.$emit('input', [...this.tags]);
      this.$emit('change', [...this.tags]);
    }
  },
  beforeMount() { },
  mounted() { },
};
</script>

<style scoped lang="less">
.tags-c {
  display: flex;
  flex-wrap: wrap;
  border-radius: 5px;
  padding: 5px 6px;
  background-color: #f5f5f5;
  min-height: 38px;
  align-content: flex-start;
  position: relative;

  .tag {
    color: rgba(16, 16, 16);
    box-sizing: border-box;
    margin: 2px 0 2px 6px;
    padding: 0px 8px;
    background-color: rgb(222, 222, 222);
    border: 1px solid rgb(233, 233, 235);
    border-radius: 4px;
    white-space: nowrap;
    height: 24px;

    .el-icon-close {
      border-radius: 50%;
      text-align: center;
      position: relative;
      cursor: pointer;
      font-size: 12px;
      height: 16px;
      width: 16px;
      line-height: 16px;
      vertical-align: middle;
      color: #909399;
      background-color: #C0C4CC;
      right: -3px;
      top: 0;
      margin-left: -3px;
      transform: scale(0.7);

      &:hover {
        color: rgb(255, 255, 255);
        background-color: rgb(144, 147, 153);
      }
    }
  }

  input {
    border: none;
    outline: 0;
    padding: 0;
    padding-left: 8px;
    color: #666;
    font-size: 14px;
    height: 28px;
    background-color: transparent;

    &::placeholder {
      color: #C0C4CC;
      font-size: 14px;
    }
  }

  &.focus {
    // border-color: #409EFF;
  }

  .placeholder {
    color: #C0C4CC;
    font-size: 14px;
    padding: 0 10px;
    position: absolute;
    top: 0;
    left: 0;
    transform: translate(5px, 9px);
  }
}
</style>
