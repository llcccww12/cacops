<template>
  <div class="base-dlg">
    <el-dialog class="base-dlg" :visible.sync="dialogShow" :title="title" :width="width" :fullscreen="fullscreen" :top="top"
      :modal="modal" :modal-append-to-body="modalAppendToBody" :append-to-body="appendToBody" :lock-scroll="lockScroll"
      :custom-class="customClass" :close-on-click-modal="closeOnClickModal" :close-on-press-escape="closeOnPressEscape"
      :show-close="showClose" :center="center" :destroy-on-close="destroyOnClose" :before-close="beforeClose"
      @open="open" @opened="opened" @close="close" @closed="closed">
      <template v-slot:title>
        <slot name="title"></slot>
      </template>
      <template v-slot:default>
        <slot name="default"></slot>
      </template>
      <template v-slot:footer>
        <slot name="footer"></slot>
      </template>
    </el-dialog>
  </div>
</template>
<script>
export default {
  name: "BaseDialog",
  props: {
    visible: { type: Boolean, default: false },
    title: { type: String, default: "" },
    width: { type: String, default: "" },
    fullscreen: { type: Boolean, default: false },
    top: { type: String },
    modal: { type: Boolean, default: true },
    modalAppendToBody: { type: Boolean, default: true },
    appendToBody: { type: Boolean, default: false },
    lockScroll: { type: Boolean, default: false },
    customClass: { type: String, default: "" },
    closeOnClickModal: { type: Boolean, default: false },
    closeOnPressEscape: { type: Boolean, default: true },
    showClose: { type: Boolean, default: true },
    beforeClose: { type: Function },
    center: { type: Boolean, default: false },
    destroyOnClose: { type: Boolean, default: false },
  },
  data() {
    return {
      dialogShow: false,
    };
  },
  watch: {
    visible: function (val) {
      this.dialogShow = val;
    },
  },
  methods: {
    open() {
      this.$emit("open");
    },
    opened() {
      this.$emit("opened");
    },
    close() {
      this.$emit("close");
    },
    closed() {
      this.$emit("closed");
      this.$emit("update:visible", false);
    },
  },
  mounted() {},
};
</script>
<style scoped lang="less">
.base-dlg {
  /deep/ .el-dialog__header {
    text-align: left;
    height: 45px;
    background: rgb(240, 240, 240);
    border-radius: 5px 5px 0px 0px;
    border-bottom: 1px solid rgb(212, 212, 213);
    padding: 0 15px;
    display: flex;
    align-items: center;
    font-weight: 500;
    font-size: 16px;
    color: rgb(16, 16, 16);

    .el-dialog__title {
      font-weight: 500;
      font-size: 16px;
      color: rgb(16, 16, 16);
    }

    .el-dialog__headerbtn {
      top: 15px;
      right: 15px;
    }
  }

  /deep/ .el-dialog__body {
    padding: 15px 15px;
  }
}
</style>
