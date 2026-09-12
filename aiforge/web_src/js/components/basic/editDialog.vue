<template>

    <el-dialog :close-on-click-modal="!deleteLoading" :title="dialogTitle" :visible.sync="deleteDialog">
        
            <div class="message-box__content">
                
                <div class="message-box-title" >
                    <slot name="title"></slot>
                </div>
                <div class="message-box-info">
                    <slot name="info"></slot>
                </div>
            </div>
            <slot name="content"></slot>
        
        <div slot="footer" class="dialog-footer">

            <button class="ui  button" @click="deleteDialog = false">{{i18n['cancel']}}</button>
            <button class="ui green button" @click="deleteCallback.call(vmContext,deleteParam)">{{i18n['confirm']}}</button>
            <!-- <el-button size="small" style="font-size: 1rem;padding: .78571429em 1.5em .78571429em;border-radius: .28571429rem;" @click="deleteDialog = false">{{"取消"}}</el-button>
            <el-button size="small" style="background-color: #21ba45;color: #fff;font-size: 1rem;padding: .78571429em 1.5em .78571429em;border-radius: .28571429rem;" @click="deleteCallback.call(vmContext,deleteParam)">{{"确定"}}</el-button> -->
        </div>
    </el-dialog>
</template>
<script type="text/javascript">
export default {
  data() {
    return {
      deleteDialog: false,
    };
  },
  props: {

    vmContext: {
      type: Object,
      default() {
        return {};
      }
    },
    dialogTitle: {
      type: String,
      default: '',
    },
    deleteLoading: {
      type: Boolean,
      default: false,
    },
    deleteCallback: {
      type: Function,
      default() {
        return () => {};
      }
    },
    deleteParam: {
      type: String,
      default: ''
    },
    value: {
      type: Boolean,
      default: false,
    }
  },
  computed: {
  },
  watch: {
    deleteDialog() {
      this.$emit('input', this.deleteDialog);
    },
    value() {
      this.deleteDialog = this.value;
    },
  },
  created() {
    this.i18n = window.i18n;
    this.deleteDialog = this.value;
  }
};
</script>

<style scoped>
.el-message-box__content .icon{float:left;}
.message-box__content{}
.message-box__content .icon{float:left;margin-right:20px;}
.message-box__content .message-box-title{font-size:16px;padding:2px 0;color:#333;}
.message-box__content .message-box-p{font-size:16px;padding:20px 0;color:#333;}
.message-box__content .message-box-info{color:#999;padding:2px 0;}


/deep/ .el-dialog__body{
  padding: 1rem;
}
/deep/ .el-dialog__header {
    background: #f0f0f0;
    padding: 1rem;
}
/deep/ .el-dialog__title {
    font-family: PingFangSC-Regular;
    font-size: 1.28571429rem;
    color: rgba(0,0,0,.87);
    font-weight: 200;
    line-height: 25px;
    height: 25px;
}
/deep/ .el-dialog__footer {
    background: #eff3f9;
    padding: 1rem;
}
/deep/ .el-dialog{
    width: 30%;
}
/deep/ .el-form-item__label{
    padding: 0;
}
</style>
