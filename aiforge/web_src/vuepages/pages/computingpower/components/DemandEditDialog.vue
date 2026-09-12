
<template>
  <div class="common-tips-dlg">
    <div class="trigger-container" @click="handlerOpen">
      <slot></slot>
    </div>
    <BaseDialog :visible="visible" :title="title" :show-close="showClose" @open="openDialog"
      @closed="closeDialog" :appendToBody="appendToBody">
      <div class="body-c">
        <div class="form-c">
          <DemandForm ref="demandFormRef" type="edit" :data="data" @success="successHandle" @error="errorHandle">
          </DemandForm>
        </div>
      </div>
    </BaseDialog>
  </div>
</template>

<script>
import BaseDialog from '~/components/BaseDialog.vue';
import DemandForm from './DemandForm.vue';
export default {
  name: "DemandEditDialog",
  props: {
    title: { type: String, default: 'Title' },
    appendToBody: { type: Boolean, default: false },
    data: { type: Object, default: () => ({}) }
  },
  components: { BaseDialog, DemandForm },
  data() {
    return {
      showClose: true,
      visible: false,
      loading: false,
    }
  },
  methods: {
    handlerOpen() {
      this.visible = true;
    },
    openDialog() {
      this.$refs['demandFormRef']?.initEdit();
    },
    closeDialog() {
      this.visible = false;
    },
    successHandle() {
      this.$emit('success');
      this.closeDialog();
    },
    errorHandle() {
      this.$emit('error');
    }
  },
  mounted() { },
}
</script>

<style scoped lang="less">
/deep/ .el-dialog {
  width: 800px
}
@media only screen and (max-width: 767px) {
  /deep/ .el-dialog {
    width: 100%;
  }
}
.trigger-container {
  display: inline-block;
}

.dialog-footer {
  text-align: center;

  .close-btn {
    color: #fff;
    background-color: #21ba45;
    border-color: #21ba45;
    font-size: 1rem;
    font-weight: 700;
  }
}

.body-c {
  display: flex;
  align-items: center;
  justify-content: center;
  padding-top: 20px;

  .form-c {
    width: 600px;
    text-align: left;
  }
}
</style>
