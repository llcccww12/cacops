
<template>
    <BaseDialog :visible="visible" :title="onlineInferTitle" top="14vh" :show-close="showClose">
        <div class="tipsWrapper">
            <h4><strong>{{$t('cloudbrainObj.dialogTips.title1')}}</strong></h4>
            <ol>
                <li class="mr-1">
                  <p>
                    <strong>{{$t('cloudbrainObj.dialogTips.tips1')}}</strong>
                    <br>
                    <strong>{{$t('cloudbrainObj.dialogTips.tips1_1')}}</strong>
                  </p>
                  <pre class="code-block">
                    <code class="chroma">app = gr.mount_gradio_app(app, block, path='/')</code>
                  </pre>
                  <p>
                    <strong>{{$t('cloudbrainObj.dialogTips.tips1_2')}}</strong>
                  </p>
                  <pre class="code-block">
                    <code class="chroma">@app.post("/text2image")</code>
                  </pre>  
                </li>
                <li class="mr-1">
                    <p><strong>{{$t('cloudbrainObj.dialogTips.tips2')}}</strong></p>
                    <pre class="code-block">
                        <code class="chroma">uvicorn.run(app, host='0.0.0.0', port=int(os.getenv('OPENI_SELF_PORT')))</code>
                    </pre>  
                </li>
                <li class="mr-1">
                  <p><strong>{{$t('cloudbrainObj.dialogTips.tips3')}}</strong></p>
                </li>
                <li><span><strong>{{$t('cloudbrainObj.dialogTips.tips61')}}</strong></span><a href="https://openi.pcl.ac.cn/OpenIOSSG/Online-Inference_Example" target="_blank">{{$t('cloudbrainObj.dialogTips.tips62')}}</a></li>
            </ol>
            <div class="btn-wrapper">
                <el-checkbox v-model="checked" style="margin-right: 1rem;">{{$t('cloudbrainObj.dialogTips.tips7')}}</el-checkbox>
                <el-button type="primary" @click="closeDialog">{{$t('cloudbrainObj.dialogTips.tips8')}}</el-button>
            </div>
        </div>
    </BaseDialog>
</template>

<script>
import BaseDialog from '~/components/BaseDialog.vue';
export default {
    name: "DialogTips",
    components:{
        BaseDialog
    },
    props: {
        visible: { type: Boolean, default: false },
    },
    data() {
        return {
            onlineInferTitle:this.$t('cloudbrainObj.dialogTips.tips9'),
            checked:false,
            showClose:false,
        }
    },
    methods:{
        closeDialog(){
            if(this.checked){
                localStorage.setItem("isCloseOnlineTips",true)
            }else{
                localStorage.setItem("isCloseOnlineTips",false)
            }
            this.$emit("closeDialog");
        }
    },
    mounted() {
    },
}
</script>

<style scoped lang="less">
    .btn-wrapper{
        display: flex;
        justify-content: flex-end;
        align-items: center;
        margin-bottom: 1rem;
    }
    .mr-1{
        margin-bottom: 10px;
    }
    .code-block{
        background-color: #f7f7f7;
        border-radius: 3px;
        font-size: 14px;
        display: flex;
        padding: 1rem;
    }
    .chroma{
        line-height: 22px;
    }
    .base-dlg {

        /deep/ .el-dialog {
            // min-width: 50%;
            max-width: 800px;
            width: 100%;
        }
    }
</style>