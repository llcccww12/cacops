<template >
    <div>
        <el-dialog :title="$t('modelSquare.createKb')" :visible.sync="visible" width="30%" @closed="closed" :show-close="false" :close-on-press-escape="false" :close-on-click-modal="false">
            <el-form :model="form" :label-width="formLabelWidth" ref="ruleForm" :rules="rules" style="margin-right: 60px;">
              <el-form-item prop="knowledge_base_name" :label="$t('modelSquare.kbName')">
                <el-input v-model="form.knowledge_base_name" autocomplete="off" :placeholder="$t('modelSquare.createKbPlaceholder')"></el-input>
              </el-form-item>
              <el-form-item :label="$t('modelSquare.vectorType')">
                <el-select v-model="form.vector_store_type" style="width:100%">
                  <el-option label="faiss" value="faiss"></el-option>
                </el-select>
              </el-form-item>
              <el-form-item :label="$t('modelSquare.embedModel')">
                <el-select v-model="form.embed_model" style="width:100%">
                  <el-option label="m3e-base" value="m3e-base"></el-option>
                </el-select>
              </el-form-item>
            </el-form>
            <div slot="footer" class="dialog-footer">
              <el-button @click="cancel" size="small">{{$t('modelSquare.cancel')}}</el-button>
              <el-button type="primary" @click="submitForm('ruleForm')"  size="small">{{$t('modelSquare.create')}}</el-button>
            </div>
          </el-dialog>
    </div>
</template>
<script>
import { llmKbcreate } from '~/apis/modules/llmchat';
export default {
    props:{
        dialogVisible:{type:Boolean,required:true,default:false},
        modelName:{type:String,default:''},
    },
    data() {
        var checkName = (rule, value, callback) => {
            if (value === '') {
              callback(new Error(this.$t('modelSquare.kbNameDetect1')));
            }else if(!(/^[0-9a-zA-Z]+$/g.test(value))){
              callback(new Error(this.$t('modelSquare.kbNameDetect2')));
            }else{
              callback();
            }//[\u4e00-\u9fa5]
        };
        return {
            visible:false,
            knowledgeValue:'',
            form:{
                knowledge_base_name:'',
                vector_store_type:'faiss',
                embed_model:'m3e-base'
            },
            formLabelWidth: '160px',
            rules: {
                knowledge_base_name:[{ validator: checkName, trigger: 'blur' }]
            },
            loading:true
        };
    },
    watch:{
      dialogVisible:{
        handler(newVal,oldVal){
          this.visible = newVal
        },
        deep:true,
        immediate:true
      }
    },
    methods:{
        checkData(rule,value,callback){
            if(value){
              if(/[\u4e00-\u9fa5]/g.test(value)){
                callback(new Error(this.$t('modelSquare.kbNameDetect2')))
              }else{
                callback()
              }
            }
            callback()
        },
        submitForm(formName) {
            this.$refs[formName].validate((valid) => {
              if (valid) {
                const loading = this.$loading({target:'.el-dialog',lock:true})
                llmKbcreate(this.form,{model_name:this.modelName}).then((res)=>{
                  if(res.data.code===200){
                    loading.close()
                    this.visible = false
                    this.$message({
                      type: 'success',
                      message: res.data.msg,
                    });
                    this.$emit('refresh')
                  }else{
                    loading.close()
                    this.$message({
                      type: 'error',
                      message: res.data.msg,
                    });
                  }
                }).catch((err)=>{
                  this.$message({
                    type: 'error',
                    message: err.message,
                  });
                  loading.close()
                })
              } else {
                return false;
              }
            });
        },
        closed(){
          this.form = {knowledge_base_name:'',vector_store_type:'faiss',embed_model:'m3e-base',}
          this.$emit("close");
        },
        cancel(){
          this.visible = false
        }
    }       
}
</script>
