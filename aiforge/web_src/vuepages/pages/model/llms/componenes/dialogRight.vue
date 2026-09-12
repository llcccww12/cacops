<template>
    <div class="model-dialog-right">
        <div>
          <p class="use-pattern">{{$t('modelSquare.dialogModeSelect')}}</p>
          <div class="pattern-wrap">
            <el-radio-group v-model="pattern" @change="patternChange">
              <el-radio label="1" border>{{$t('modelSquare.dialogLLM')}}</el-radio>
              <el-radio label="2" border>{{$t('modelSquare.dialogKb')}}</el-radio>
            </el-radio-group>
          </div>
          <div class="component-2" v-if="pattern==='2'" style="min-height: 560px;">
            <div class="label-wrap">
              <span class="titlr">{{$t('modelSquare.configKb')}}</span>
              <span class="icon-rotate" style="transform: rotate(0deg);">▼</span>
            </div>
            <div class="updata-knowledge" @click="recreateVe">
              <span>{{$t('modelSquare.updatekb')}}</span>
            </div>
            <div class="recreate-kb" v-if="recreateFlag">
              <div class="nowrap">
                <i class="el-icon-loading" style="color:rgb(28, 131, 225)"></i>
                <span :title="$t('modelSquare.recreateKb')">{{$t('modelSquare.recreateKb')}}</span>
              </div>
              <div class="nowrap" style="margin-top: 1rem;font-size:12px" :title="recreateVeValue.msg">{{recreateVeValue.msg}}</div>
              <el-progress :percentage="recreateVeValue.percentage"></el-progress>

            </div>
            <div :class="{'disabled': recreateFlag}">
              <span style="display: inline-block;color:#101010">{{$t('modelSquare.selectKb')}}</span>
              <el-select v-model="knowledgeValue" style="width: 100%;margin-top:12px" @change="changeKbValue" v-loading="loading">
                <el-option
                  v-for="item in KnowledgeBaseList"
                  :key="item.value"
                  :label="item.value"
                  :value="item.value">
                </el-option>
              </el-select>
            </div>
            <div class="knowledge-op-btn" :class="{'disabled': recreateFlag}">
              <div @click="dialogVisible = true">
                <span>{{$t('modelSquare.createKb')}}</span>
              </div>
              <div>
                <el-popconfirm
                :title="$t('modelSquare.deleteKbTips',{knowledgeValue})"
                @confirm="confirmDel"
                >
                    <span slot="reference" :class="{'disabled':knowledgeValue===commonKB}">{{$t('modelSquare.deleteKb')}}</span>
                </el-popconfirm>
                
              </div>
            </div>
            <el-tabs v-model="activeName" type="border-card" @tab-click="tabClick" :class="{'disabled': recreateFlag}">
              <el-tab-pane name="upload" :class="{'disabled':knowledgeValue===commonKB}">
                <span slot="label">{{$t('modelSquare.uploadFile')}}</span>
                <div class="upload-knowledge-file">
                  <el-upload
                    class="upload-file"
                    drag
                    action=""
                    multiple
                    :show-file-list="false"
                    :http-request="getUploadFileList"
                    accept=".html, .md, .json, .csv, .txt, .xml, .docx"
                    >
                    <i class="el-icon-upload" style="margin: 1rem 0;"></i>
                    <div class="el-upload__text" v-html="$t('modelSquare.uploadFileTips1')"></div>
                    <div class="el-upload__tip" slot="tip">{{$t('modelSquare.uploadFileTips2')}}</div>
                  </el-upload>
                  <ul class="upload-list">
                    <li class="upload-list-item" v-for="(item,index) in files" :key="item.uid">
                      <a class="upload-item-name" :title="item.name" style="cursor: none;">
                        <i class="el-icon-document"></i>
                        {{item.name}}
                      </a>
                      <label class="upload-item-status">
                        <i v-if="filesStatus[index].status===1"  style="color: #67C23A;" class="el-icon-circle-check"></i>
                        <i v-else-if="filesStatus[index].status===2" :title="filesStatus[index].error" style="color: red;" class="el-icon-warning-outline"></i>
                        <i v-else style="color: #909399;cursor:pointer" class="el-icon-close" @click="delKbFile(index)"></i>
                      </label>
                    </li>
                  </ul>
                </div>
    
                <div class="knowledge-op-btn">
                  <div style="width: 100%;" :class="{disabled:upBtnDisabled}" @click="uploadDocKnowledge">
                    <span >{{$t('modelSquare.addFileToKb')}}</span>
                  </div>
                </div>
              </el-tab-pane>
              <el-tab-pane :label="$t('modelSquare.manageFile')" name="manage" :class="{'disabled':knowledgeValue===commonKB}">
                <div class="knowledege-detail-file">
                  <p>{{$t('modelSquare.deleteKbFileSelect')}}</p>
                  <el-checkbox-group v-model="checkList" size="small" style="max-height: 150px;overflow-y:auto">
                    <el-checkbox :label="kbFile" v-for="kbFile in kbFileList" :key="kbFile"></el-checkbox>
                  </el-checkbox-group>

                  <div class="knowledge-op-btn">
                    <el-popconfirm :title="$t('modelSquare.deleteVbTips')" @confirm="confirmDelFile" style="width: 100%;">
                      <div slot="reference" style="width: 100%;" :class="{'disabled':noCheck}">
                        <span>{{$t('modelSquare.deleteKbFile')}}</span>
                      </div>
                    </el-popconfirm>
                    
                  </div>
                </div>
              </el-tab-pane>

            </el-tabs>
          </div>
          <createKbDialog :modelName="modelName" :dialogVisible="dialogVisible" @close="dialogVisible=false" @refresh="refresh"></createKbDialog>
        </div>
      </div>
</template>
<script>
import { llmKbDelete, llmKbDeleteDoc,llmKbList,llmKbUploadDocUrl,llmKbUploadDoc,llmKbFileList,llmRecreateVectorStore} from '~/apis/modules/llmchat';
import createKbDialog from './createKbDialog.vue'
let csrf = window.config ? window.config.csrf : ''
export default {
    name: "dialogRight",
    props: {
      commonKB:{type:String,default:''},
      modelName:{type:String,default:''},
    },
    components: { createKbDialog},
    data() {
        return {
            pattern:'1',
            knowledgeValue:'',
            dialogVisible:false,
            files:[],
            form:{
                knowledge_base_name:'',
                vector_store_type:'faiss',
                embed_model:'m3e-base'
            },
            formLabelWidth: '120px',
            KnowledgeBaseList:[],
            
            flag:false,
            kbFileList:[],
            checkList:[],
            uploadFlag:false,
            upBtnDisabled:false,
            filesStatus:[],
            activeName:'upload',
            loading:false,
            recreateVeValue:{},
            recreateFlag:false,
            noCheck:true
        };
    },
    watch:{
      knowledgeValue(val){
        if (val){
          this.$emit('changeKbName',val)
        }
      },
      checkList(val){
        if(val.length===0){
          this.noCheck = true
        }else{
          this.noCheck = false
        }
      }
    },
    methods:{
        patternChange(val){
          this.$emit('radioChange',val)
          if(val==='2'){
            this.getKnowledgeBaseList()
          }
        },
        tabClick(tab, event){
          if(tab.name ==='manage'){
            this.getKbFileDetails()
          } 
        },
        confirmDel(){
            this.deleteKnowledgeBase()
        },
        confirmDelFile(){
          this.delKbFileVe()
        },
        recreateVe(){
          this.recreateVeValue.percentage = 0
          let data = {
            knowledge_base_name:this.knowledgeValue,
            allow_empty_kb:true,
            vs_type:"faiss",
            embed_model:"m3e-base",
          }
          this.recreateFlag = true
          llmRecreateVectorStore({knowledge_base_name:this.knowledgeValue,model_name:this.modelName}).then((response)=>{
            if(response.status===200){
              const reader = response.body.getReader();
              const processBinaryData = async () => {
                while (true) {
                  const { done, value } = await reader.read();
                  if (done) {
                    // The entire response has been processed
                    this.recreateFlag = false
                    this.$message({
                      type: 'success',
                      message: this.$t('modelSquare.recreateKbSuccess',{knowledgeValue:this.knowledgeValue}),
                    });
                    break;
                  }
                  // Handle the binary data in the 'value' variable
                  let chars = new TextDecoder().decode(value)
                  if(chars.indexOf('}{')!==-1){
                    const textArray = chars.split('}{')
                    let str1 = textArray[0] + "}"
                    let str2 = "{" + textArray[1]
                  }else{
                    const paraseChars = JSON.parse(chars)
                    if(paraseChars.code===200){
                      let percentage = Math.ceil(((paraseChars.finished + 1)/paraseChars.total)*100)
                      this.recreateVeValue = {...JSON.parse(chars),percentage:percentage}
                    }else{
                      this.recreateFlag = false
                      this.$message({
                        type: 'error',
                        message: paraseChars.msg,
                      });
                      
                      return
                    }
                  }
                  // You can process the binary data here and update your UI as needed
                }
              };
          
              // Start processing the binary data
              processBinaryData();
            }else if(response.status===403){
              this.recreateFlag = false
              this.$message({
                type: 'error',
                message: this.$t('modelSquare.noPermission'),
              });
            }else{
              this.recreateFlag = false
              this.$message({
                type: 'error',
                message: response.statusText,
              });
            }
          })
          /*llmRecreateVectorStore({knowledge_base_name:this.knowledgeValue}).then((res)=>{
          })*/
        },
        deleteKnowledgeBase(){
          const loading = this.$loading({target:'.component-2',lock:true})
          llmKbDelete({knowledge_base_name:this.knowledgeValue,model_name:this.modelName}).then((res)=>{
              if(res.data.code===200){
                this.$message({
                  type: 'success',
                  message: res.data.msg,
                });
                this.getKnowledgeBaseList()
              }else{
                this.$message({
                  type: 'error',
                  message: res.data.msg,
                });
              }
              loading.close()
            }).catch((err)=>{
              this.$message({
                type: 'error',
                message: err.message,
              });
              loading.close()
            })
        },
        delKbFile(index){
            this.files.splice(index,1)
            this.filesStatus.splice(index,1)
        },
        async delKbFileVe(){
            const loading = this.$loading({target:'.component-2',lock:true})
            if(this.checkList.length!==0){
              const data = {
                knowledge_base_name:this.knowledgeValue,
                file_names: this.checkList
              }
              llmKbDeleteDoc(data,{model_name:this.modelName}).then((res)=>{
                loading.close()
                if(res.data.code===200){
                  const {failed_files} = {...res.data.data}
                  if(Object.keys(failed_files).length===0){
                    this.$message({
                      message: res.data.msg,
                      type: 'success',
                    });
                    this.checkList = []
                    this.getKbFileDetails()
                  }else{
                    this.getKbFileDetails()
                    Object.keys(failed_files).forEach((item)=>{
                      setTimeout(()=>{
                        this.$notify({
                          message: failed_files[item],
                          type: 'error',
                          position: 'bottom-right'
                        })
                      },500);
                    })
                  }
                }else{
                  this.$message({
                    type: 'error',
                    message: res.data.msg,
                  });
                }
              }).catch((error)=>{
                loading.close()
                if(error.response.status===403){
                  this.$message({
                    type: 'error',
                    message: this.$t('modelSquare.noPermission'),
                  })
                  this.checkList=[]
                }else{
                  this.$message({
                    type: 'error',
                    message: error.message,
                  });
                }
              })
            }else{
              loading.close()
            }
        },
        
        getUploadFileList(value){
            if(this.uploadFlag){
                this.files = []
                this.filesStatus = []
                this.uploadFlag = false
                this.upBtnDisabled = false
            }
            if(this.files.length>=10){
              setTimeout(()=>{
                this.$message.error(this.$t('modelSquare.uploadFIleLimit'))
              },0)
              return
            }
            const index = value.file.name.lastIndexOf('.')
            const acceptFileTypes = ['html', 'md', 'json', 'csv', 'txt', 'xml', 'docx'] //
            if(index === -1){
              setTimeout(()=>{
                this.$message({
                  type: 'error',
                  message: `${value.file.name}${this.$t('modelSquare.fileError')}`,
                });
              },0)
              return
            }else if(this.files.findIndex(f=>f.name===value.file.name)!==-1){
              setTimeout(()=>{
                this.$message({
                  type: 'error',
                  message: `${value.file.name}${this.$t('modelSquare.fileExit')}`,
                });
              },0)
              return
            }else if(value.file.size/(1024*1024)>=1){
              setTimeout(()=>{
                this.$message({
                  type: 'error',
                  message: `${value.file.name}${this.$t('modelSquare.fileExceed')}`,
                });
              })
              return
            }else if(!acceptFileTypes.includes(value.file.name.substr(index+1))){
              setTimeout(()=>{
                this.$message({
                  type: 'error',
                  message: `${value.file.name}${this.$t('modelSquare.fileTypeError')}`,
                });
              },0)
              return
            }else{
              value.file.status = 0
              this.files.push(value.file)
              this.filesStatus.push({name:value.file.name,status:0})
            }
        },
        getKnowledgeBaseList(){
            this.loading = true
            this.KnowledgeBaseList=[]
            llmKbList({model_name:this.modelName}).then((res)=>{
              if(res.data.code===200){
                this.knowledgeValue = res.data.data[0]
                this.getKbFileDetails()
                res.data.data.forEach((item)=>{
                  this.KnowledgeBaseList.push({value:item})
                })
              }else{
                this.$message({
                  type: 'error',
                  message: res.data.msg,
                });
              }
              this.loading = false
            }).catch((err)=>{
              this.$message({
                type: 'error',
                message: err.message,
              });
              this.loading = false
            })
            
          },
          uploadDocKnowledge(){
            if(this.files.length!==0){
              const loading = this.$loading({target:'.component-2',lock:true})
              const fd = new FormData()
                fd.append("override",false)
                fd.append("knowledge_base_name",this.knowledgeValue)
                this.files.forEach((file)=>{
                fd.append("files",file)
              })
              llmKbUploadDocUrl({model_name:this.modelName},fd).then((res)=>{
                if(res.data.code===200){
                  const {failed_files} = {...res.data.data}
                  this.filesStatus.forEach((file,index)=>{
                    this.filesStatus[index].status = 1
                  })
                  if(Object.keys(failed_files).length!==this.files.length){
                    this.$message({
                      type: 'success',
                      message: res.data.msg,
                    });
                  }
                  if(Object.keys(failed_files).length!==0){
                    Object.keys(failed_files).forEach((item, index)=>{
                      setTimeout(()=>{
                        this.$notify({
                          message: failed_files[item],
                          type: 'error',
                          position: 'bottom-right'
                        })
                      },500);
                      this.filesStatus.forEach((file,index)=>{
                        if(item===file.name){
                          this.filesStatus[index].status = 2
                          this.filesStatus[index].error = failed_files[item]
                        }
                      })
                      
                    })
                  }
                  this.uploadFlag = true
                  this.upBtnDisabled = true
                }else{
                  this.$message({
                    type: 'error',
                    message: res.data.msg,
                  });
                }
                loading.close()
              }).catch((err)=>{
                if(err.response.status===403){
                  this.$message({
                    type: 'error',
                    message: this.$t('modelSquare.noPermission'),
                  })
                }else{
                  this.$message({
                    type: 'error',
                    message: err.message,
                  });
                }
                loading.close()
              })

            }
            
          },
          changeKbValue(val){
            this.files=[]
            this.upBtnDisabled = false
            this.checkList = []
            this.getKbFileDetails()
          },
          getKbFileDetails(){
            const loading = this.$loading({target:'.knowledege-detail-file',lock:true})
            llmKbFileList({knowledge_base_name:this.knowledgeValue,model_name:this.modelName}).then((res)=>{
              if(res.data.code===200){
                loading.close()
                this.kbFileList = res.data.data
              }else{
                this.$message({
                  type: 'error',
                  message: res.data.msg,
                });
                loading.close()
              }
            }).catch((err)=>{
              this.$message({
                type: 'error',
                message: err.message,
              });
              loading.close()
            })
          },
          refresh(){
            this.dialogVisible = false
            this.getKnowledgeBaseList()
          }
    },
    async mounted() {
        
    
    }
}
</script>
<style lang="less" scoped>
  .model-dialog-right{
    width: 30%;
    .pattern-wrap{
      display: flex;
      flex-wrap: wrap;
      gap:8px;
    }
    .use-pattern{
      font-family: SourceHanSansSC;
      font-weight: 550;
      font-size: 16px;
      color: #101010;
      font-style: normal;
      letter-spacing: 0px;
      line-height: 26px;
      text-decoration: none;
    }
    .component-2{
      border: 1px solid #e5e7eb;
      border-radius: 8px;
      padding: 12px 16px;
      margin-top: 1rem;
      display: flex;
      flex-direction: column;
      .label-wrap{
        font-weight: 550;
        font-size: 16px;
        color: #101010;
        margin-bottom:8px;
        display: flex;
        justify-content: space-between;
        cursor: pointer;
        width:100%;

        .icon-rotate{
          transition: .15s;
        }
        
      }
      .updata-knowledge{
        width: 100%;
        height: 40px;
        font-size: 14px;
        border: 1px solid rgba(1, 145, 255, 0.3);
        background: rgba(1, 145, 255, 0.1);
        color: rgb(16, 16, 16);
        display: flex;
        justify-content: center;
        align-items: center;
        margin: 1.5rem 0;
        border-radius: 5px;
        cursor: pointer;
      }
      .recreate-kb{
        border: 1px solid #e5e7eb;
        border-radius: 4px;
        display: flex;
        flex-direction: column;
        padding: 0.8rem;
        margin-bottom: 1rem;
      }
      .knowledge-op-btn{
        display: flex;
        justify-content: space-between;
        margin: 0.8rem 0;
        div{
          border: 1px solid rgba(1, 145, 255, 0.3);
          background: rgba(1, 145, 255, 0.1);
          color: rgb(16, 16, 16);
          display: flex;
          justify-content: center;
          align-items: center;
          height: 32px;
          width:40%;
          border-radius: 5px;
          cursor:pointer;
        }
      }
      .upload-file{
        width: 100%; 
        margin-top: 1rem;
      }
      .upload-list{
        margin: 0;
        padding: 0;
        list-style: none;
        max-height: 100px;
        overflow-y: auto;
        .upload-list-item{
          transition: all 0.5s cubic-bezier(0.55, 0, 0.1, 1);
          font-size: 14px;
          color: #606266;
          line-height: 1.8;
          margin-top: 5px;
          position: relative;
          box-sizing: border-box;
          border-radius: 4px;
          width: 100%;
          &:first-child{
            margin-top: 10px;
          }
          .upload-item-name{
            color: #606266;
            display: block;
            margin-right: 40px;
            overflow: hidden;
            padding-left: 4px;
            text-overflow: ellipsis;
            transition: color 0.3s;
            white-space: nowrap;
            .el-icon-document{
              color: #909399;
            }
          }
          .upload-item-status{
            position: absolute;
            right: 5px;
            top: 0;
            line-height: inherit;
            display: block;
          }
        }
      }
    }
  }
  /deep/ .el-radio {
    margin-right: 8px;
  }
  /deep/ .el-radio.is-bordered.is-checked{
    border-color: rgb(99, 102, 241);
    background: rgb(99, 102, 241);
  } 
  /deep/ .el-radio__input.is-checked + .el-radio__label{
    color: white;
  }
  /deep/ .el-radio__input.is-checked .el-radio__inner{
    border-color: rgb(79, 70, 229);
    background: rgb(79, 70, 229);
  }
  /deep/ .el-dialog__body{
    padding-bottom: 0;
  } 
  /deep/ .el-upload{
    width: 100%;
  }
  /deep/ .el-upload-dragger{
    width: 100%;
    height: 120px;
  }
  /deep/ .el-tabs--border-card{
    border-radius: 0.5rem;
  }
</style>