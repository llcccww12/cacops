<template>
    <div class="menu-area">
         <div class="menu-select-container">
             <div>
                 <div class="text-select-container">
                     <div class="text-content">
                         <div v-for="key in Object.keys(config)" class="label-wrapper">
                             <template v-if="key==='system_message'">
                                 <div class="item-label">
                                     <span>{{$t('modelSquare.'+key)}}</span>
                                     <el-tooltip class="item" effect="dark" :content="config[key].desc" placement="right">
                                         <svg viewBox="64 64 896 896" focusable="false" data-icon="question-circle" width="1em" height="1em" fill="currentColor" aria-hidden="true"><path d="M512 64C264.6 64 64 264.6 64 512s200.6 448 448 448 448-200.6 448-448S759.4 64 512 64zm0 820c-205.4 0-372-166.6-372-372s166.6-372 372-372 372 166.6 372 372-166.6 372-372 372z"></path><path d="M623.6 316.7C593.6 290.4 554 276 512 276s-81.6 14.5-111.6 40.7C369.2 344 352 380.7 352 420v7.6c0 4.4 3.6 8 8 8h48c4.4 0 8-3.6 8-8V420c0-44.1 43.1-80 96-80s96 35.9 96 80c0 31.1-22 59.6-56.1 72.7-21.2 8.1-39.2 22.3-52.1 40.9-13.1 19-19.9 41.8-19.9 64.9V620c0 4.4 3.6 8 8 8h48c4.4 0 8-3.6 8-8v-22.7a48.3 48.3 0 0130.9-44.8c59-22.7 97.1-74.7 97.1-132.5.1-39.3-17.1-76-48.3-103.3zM472 732a40 40 0 1080 0 40 40 0 10-80 0z"></path></svg>
                                     </el-tooltip>
                                 </div>
                                 <div>
                                     <el-input
                                         type="textarea"
                                         :placeholder="$t('modelSquare.systemPlaceholder')"
                                         v-model="params.messages[0].content"
                                         maxlength="1024"
                                         show-word-limit
                                         rows="3"
                                         resize="none"
                                         style="margin-left: -2px;"
                                     >
                                     </el-input>
                                 </div>
                             </template>
                             <template v-else-if="key==='history'">
                                 <div class="item-label" style="justify-content: space-between;">
                                 <div>
                                     <span>{{$t('modelSquare.'+key)}}</span>
                                     <el-tooltip class="item" effect="dark" :content="config[key].desc" placement="right">
                                         <svg viewBox="64 64 896 896" focusable="false" data-icon="question-circle" width="1em" height="1em" fill="currentColor" aria-hidden="true"><path d="M512 64C264.6 64 64 264.6 64 512s200.6 448 448 448 448-200.6 448-448S759.4 64 512 64zm0 820c-205.4 0-372-166.6-372-372s166.6-372 372-372 372 166.6 372 372-166.6 372-372 372z"></path><path d="M623.6 316.7C593.6 290.4 554 276 512 276s-81.6 14.5-111.6 40.7C369.2 344 352 380.7 352 420v7.6c0 4.4 3.6 8 8 8h48c4.4 0 8-3.6 8-8V420c0-44.1 43.1-80 96-80s96 35.9 96 80c0 31.1-22 59.6-56.1 72.7-21.2 8.1-39.2 22.3-52.1 40.9-13.1 19-19.9 41.8-19.9 64.9V620c0 4.4 3.6 8 8 8h48c4.4 0 8-3.6 8-8v-22.7a48.3 48.3 0 0130.9-44.8c59-22.7 97.1-74.7 97.1-132.5.1-39.3-17.1-76-48.3-103.3zM472 732a40 40 0 1080 0 40 40 0 10-80 0z"></path></svg>
                                     </el-tooltip>
                                 </div>
                                 <el-switch v-model="historyFlag" @change="change" :disabled="config[key].disabled"></el-switch>
                             </div>
                             </template>
                             <template v-else>
                                 <div class="item-label">
                                     <span>{{$t('modelSquare.'+key)}}</span>
                                         <el-tooltip class="item" effect="dark" :content="config[key].desc" placement="right">
                                         <svg viewBox="64 64 896 896" focusable="false" data-icon="question-circle" width="1em" height="1em" fill="currentColor" aria-hidden="true"><path d="M512 64C264.6 64 64 264.6 64 512s200.6 448 448 448 448-200.6 448-448S759.4 64 512 64zm0 820c-205.4 0-372-166.6-372-372s166.6-372 372-372 372 166.6 372 372-166.6 372-372 372z"></path><path d="M623.6 316.7C593.6 290.4 554 276 512 276s-81.6 14.5-111.6 40.7C369.2 344 352 380.7 352 420v7.6c0 4.4 3.6 8 8 8h48c4.4 0 8-3.6 8-8V420c0-44.1 43.1-80 96-80s96 35.9 96 80c0 31.1-22 59.6-56.1 72.7-21.2 8.1-39.2 22.3-52.1 40.9-13.1 19-19.9 41.8-19.9 64.9V620c0 4.4 3.6 8 8 8h48c4.4 0 8-3.6 8-8v-22.7a48.3 48.3 0 0130.9-44.8c59-22.7 97.1-74.7 97.1-132.5.1-39.3-17.1-76-48.3-103.3zM472 732a40 40 0 1080 0 40 40 0 10-80 0z"></path></svg>
                                     </el-tooltip>
                                 </div>
                                 <div>
                                     <el-slider v-model="params[key]" show-input 
                                     :show-input-controls="false" :min="config[key].min" 
                                     :max="config[key].max" :step="config[key].step"></el-slider>
                                 </div>
                             </template>
                         </div>
                     </div>
                 </div>
             </div>
         </div>
     </div>
 </template>
 <script>
 
 export default {
    name: 'ChatMenu',
    props: {
     config:{type:Object,default:()=>({})},
    },
    data() {
      return {
        params: {
            messages: [],
            max_tokens: 1024,
            temperature: 0.5,
            top_p: 0.9,
            repetition_penalty: 0.9,
            model: ''
         },
         historyFlag: false
      }
    },
    computed: {
      
    },
    watch: {
     config:function(value){
         if(JSON.stringify(value) !== '{}'){
             this.initParams()
         }
     }
    },
    methods: {
     change() {
         this.params.messages = [{role: "system",content: "你是一个聊天助手"}]
     },
    //  change() {

    //  },
     sendPamras() {
         this.$emit('changeParams',this.params,this.historyFlag);
     },
     initParams(){
         let params = this.config
         for (let key in params) {
             if(key == 'history'){
                 this.historyFlag = params[key].default
             }else if(params[key] !== null ){
                 this.params[key] = params[key].default
             }
         }
     },
 
    },
    mounted() {
    },
    beforeMount() {
 
    },
   
 };
 </script>
 <style lang='less' scoped>
 .menu-area{
     border-radius: 10px;
     background-color: rgba(249,250,251,1);
     border: 1px solid rgba(16,16,16,0.15);
     color: rgba(16,16,16,1);
     padding: 26px 14px 20px 22px;
     overflow-y: auto;
     min-height: 400px;
     min-width: 368px;
     height: calc(100vh - 250px);
     .menu-select-container{
         display: flex;
         flex-direction: column;
     }
     .text-model{
         font-size: 18px;
         font-weight: 700;
         margin-bottom: 12px;
     }
     .label-wrapper{
         line-height: 32px;
         padding-top: 12px;
         .item-label{
             display: flex;
             align-items: center;
             svg{
                 color: rgb(186, 191, 199);
                 margin-left: 6px;
             }
         }
     }
 }
 </style>