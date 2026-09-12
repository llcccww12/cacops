<template>
      <div>
          <div class="input-search">
            
            
            <el-input v-model="input" clearable  :autofocus="true" @input="changeValue" id="topics_input" @keyup.enter.native="postTopic" :placeholder="i18n['searchOrCreateTopics']">

            </el-input>
            <div class="scrolling-menu">
                <div v-if="showSearchTopic" class="item-text" v-for="(arr,i) in array" @click="addTopics(i,arr)">
                        <div class="icon-wrapper">
                            <i style="line-height: 1.5;color: #303643;font-weight: 900;" v-if="showInitTopic[i]" class="el-icon-check" ></i>
                        </div>
                        <div class="text">{{arr.topic_name.toLowerCase()}} </div>
                </div>
                <div v-if="showInputValue" class="addition item-text"  @click="postTopic">
                    {{i18n['clickOrEnterToAdd']}}<b class="user-add-label-text">{{input.toLowerCase()}}</b>{{i18n['topic']}}
                </div>
                <div v-if="showAddTopic" class="item-text" @click="addPostTopic">
                        <div class="icon-wrapper">
                            <i  style="line-height: 1.5;color: #303643;font-weight: 900;" v-if="showAddFlage" class="el-icon-check" ></i>
                        </div>
                        <div class="text">{{input.toLowerCase()}}</div>
                </div>

            </div>

          </div>
            
      </div>
</template>

<script>

const {AppSubUrl, StaticUrlPrefix, csrf} = window.config;

import $ from 'jquery'


export default {
  
  data() {
    return {
        input:'',
        params:{},
        showInputValue:false,
        showFlag:-1,
        array:[],
        showAddTopic:false,
        showAddFlage:false,
        showSearchTopic:true,
        postUrl:'',
        arrayTopics:[],
        showInitTopic:[],

    };
  },
  methods: {

    addTopics(item,array){
        
        
        if(!this.showInitTopic[item]){
            if(this.arrayTopics.includes(array.topic_name)){
                return
            }
            else{
                this.arrayTopics.push(array.topic_name)
            
                let topics = this.arrayTopics
                let strTopics = topics.join(',')
                let data = this.qs.stringify({
                    _csrf:csrf,
                    topics:strTopics
                })
                this.Post(data,topics)
                this.$set(this.showInitTopic,item,true)
                $('#repo-topics1').children('span').remove()
            }
            
            
            

        }else{
            this.arrayTopics=this.arrayTopics.filter(ele=>{
                return ele !== array.topic_name

            })
            

            
            let topics = this.arrayTopics
            let strTopics = topics.join(',')
            let data = this.qs.stringify({
                _csrf:csrf,
                topics:strTopics
            })
            this.Post(data,topics)
            this.$set(this.showInitTopic,item,false)
            if(this.arrayTopics.length===0){
                $('#repo-topics1').append(`<span class="no-description text-italic">${this.i18n['noTopics']}</span>`)
            }else{
                $('#repo-topics1').children('span').remove()
            }
            

        }
        
        
        
        
    },
    changeValue(){
        
        if (this.input === ''){
            this.array = this.arrayTopics
            let data = []
            this.showInitTopic = []
            this.array.forEach((element,index) => {
                
                let item = {}
                item.topic_name = element
                
                data.push(item)
                this.showInitTopic.push(true)
                
                
            });
            
            this.array = data
            
            
            this.showInputValue = false
            this.showSearchTopic = true
        }
            
        else if(this.arrayTopics.indexOf(this.input.toLowerCase())>-1){
            this.showInputValue = false
            this.showSearchTopic = false
            
        }else{
            
            this.showInitTopic = []
            let timestamp=new Date().getTime()
            this.params.q = this.input.toLowerCase()
            this.params._ = timestamp
            this.$axios.get('/api/v1/topics/search',{
                params:this.params
            }).then((res)=>{
                this.array = res.data.topics
                
                this.array.forEach((element,index) => {
                
                
                if (this.arrayTopics.indexOf(element.topic_name)>-1){
                    this.showInitTopic.push(true)
                }
                else{
                    this.showInitTopic.push(false)
                }
                this.showInputValue = true
                
            });
                let findelement = this.array.some((item)=>{
                    
                
                    return item.topic_name===this.input.toLowerCase()
                })
                this.showInputValue = !findelement
                
                
            })
            
            
            
            
            this.showSearchTopic = true
        }
        this.showAddTopic = false
        
        

    },
    
    Post(data,topics){
        this.$axios.post(this.postUrl,data).then(res=>{
            const viewDiv = $('#repo-topics1');
            
            viewDiv.children('.topic').remove();
            if (topics.length) {
                const topicArray = topics;
                
                const last = viewDiv.children('a').last();
                for (let i = 0; i < topicArray.length; i++) {
                const link = $('<a class="ui repo-topic small label topic"></a>');
                link.attr(
                    'href',
                    `${AppSubUrl}/explore/repos/square?q=${encodeURIComponent(
                    topicArray[i]
                    )}&topic=`
                );
                link.text(topicArray[i]);
                // link.insertBefore(last);
                viewDiv.append(link)
                }
            }
            
            viewDiv.show();
        })
    },
    postTopic(){
        if(!this.showInputValue){
            return
        }
        const patter = /^[\u4e00-\u9fa5a-zA-Z0-9][\u4e00-\u9fa5a-zA-Z0-9-]{0,34}$/
        let regexp = patter.test(this.input)
        if(!regexp){
            this.$notify({
                
                message: '标签名必须以中文、字母或数字开头，可以包含连字符 (-)，并且长度不得超过35个字符',
                duration: 3000,
                type:'error'
                });
            return
        }else{
            let topic = this.input
            if(this.arrayTopics.includes(topic.toLowerCase())){
                return
            }
            else{
                this.arrayTopics.push(topic.toLowerCase())
            
                let topics = this.arrayTopics
                let strTopics = topics.join(',')
                let data = this.qs.stringify({
                    _csrf:csrf,
                    topics:strTopics
                })
                this.Post(data,topics)
                $('#repo-topics1').children('span').remove()
                this.showInputValue = false
                this.showAddTopic = true
                this.showAddFlage = true
            }
            
        }
        

    },
    addPostTopic(){
        
        if(this.showAddFlage){
            // this.arrayTopics.pop()

            let cancleIndex = this.arrayTopics.indexOf(this.input)
            this.arrayTopics.splice(cancleIndex,1)
            let topics = this.arrayTopics
            let strTopics = topics.join(',')
            let data = this.qs.stringify({
                _csrf:csrf,
                topics:strTopics
            })
            this.Post(data,topics)
            
            if(this.arrayTopics.length===0){
                
                $('#repo-topics1').append(`<span class="no-description text-italic">${this.i18n['noTopics']}</span>`)
            }else{
                $('#repo-topics1').children('span').remove()
            }
        }
        else if(!this.showAddFlage){
            let topic = this.input
            this.arrayTopics.push(topic.toLowerCase())
            
            let topics = this.arrayTopics
            let strTopics = topics.join(',')
            let data = this.qs.stringify({
                _csrf:csrf,
                topics:strTopics
            })
            this.Post(data,topics)
            $('#repo-topics1').children('span').remove()
        }
        this.showAddFlage = !this.showAddFlage
    },
    initTopics(){
        const mgrBtn = $('#manage_topic');
        const editDiv = $('#topic_edit');
       
  
  
        mgrBtn.on('click', (e) => {
            // viewDiv.hide();
            editDiv.css('display', ''); // show Semantic UI Grid
            
            this.input = ''
            if (this.input === ''){
            
            this.array = this.arrayTopics
            
            let data = []
            this.showInitTopic = []
            this.array.forEach((element,index) => {
                
                let item = {}
                item.topic_name = element
                
                data.push(item)
                this.showInitTopic.push(true)
                
                
            });
            
            this.array = data
            
            
            this.showInputValue = false
            this.showSearchTopic = true
            this.showAddTopic = false
            }
            
            stopPropagation(e);
            
        });
        $(document).bind('click',function(){
            editDiv.css('display','none');

        })
        editDiv.click(function(e){
            stopPropagation(e);
        })

       

        function stopPropagation(e) {
            var ev = e || window.event;
            if (ev.stopPropagation) {
                ev.stopPropagation();
            }
            else if (window.event) {
                window.event.cancelBubble = true;//兼容IE
            }
        }
    }
  },
computed:{
   
},
watch: {

    // input(newValue){
        
    //     if (newValue === ''){
    //         this.array = this.arrayTopics
    //         let data = []
    //         this.showInitTopic = []
    //         this.array.forEach((element,index) => {
                
    //             let item = {}
    //             item.topic_name = element
                
    //             data.push(item)
    //             this.showInitTopic.push(true)
                
                
    //         });
            
    //         this.array = data
            
            
    //         this.showInputValue = false
    //         this.showSearchTopic = true
    // }
    // }
},
mounted() {
      const context = this
      
      
      this.postUrl = `${window.location.pathname.split('/').slice(0,3).join('/')}/topics`;
    
      $('#repo-topics1').children('a').each(function(){
          
          context.arrayTopics.push($(this).text())
      });
      if(this.arrayTopics.length===0){
          
            $('#repo-topics1').append(`<span class="no-description text-italic">${this.i18n['noTopics']}</span>`)
      }
      this.changeValue()
    } ,
created(){
    this.i18n = window.i18n;
    this.initTopics();
    this.input=''
    
}
};
</script>

<style scoped>
.input-search {
    width: 100%;
    display: -webkit-box;
    display: -ms-flexbox;
    display: flex;
    
    min-width: 10rem;
    white-space: nowrap;
    font-size: 1rem;
    position: relative;
    display: inline-block;
    color: rgba(0,0,0,0.8);
    padding: 8px;
}
/deep/ .el-input__inner{
    border-color: #409eff;
}
.scrolling-menu{
    border-top: none !important;
    padding-top: 0 !important;
    padding-bottom: 0 !important;
    display: block;
    position: static;
    overflow-y: auto;
    border: none;
    -webkit-box-shadow: none !important;
    box-shadow: none !important;
    border-radius: 0 !important;
    margin: 0 !important;
    min-width: 100% !important;
    width: auto !important;
    border-top: 1px solid rgba(34,36,38,0.15);
}
.item-text{
    border-top: none;
    padding-right: calc(1.14285714rem + 17px ) !important;
    line-height: 1.333;
    padding-top: 0.7142857rem !important;
    padding-bottom: 0.7142857rem !important;
    position: relative;
    cursor: pointer;
    display: block;
    border: none;
    height: auto;
    text-align: left;
    border-top: none;
    line-height: 1em;
    color: rgba(0,0,0,0.87);
    padding: 0.78571429rem 1.14285714rem !important;
    font-size: 1rem;
    text-transform: none;
    font-weight: normal;
    -webkit-box-shadow: none;
    box-shadow: none;
    -webkit-touch-callout: none;
    display: -webkit-box;
    display: -ms-flexbox;
    display: flex;
    -webkit-box-align: center;
    -ms-flex-align: center;
    align-items: center;
    display: -webkit-box !important;
    display: -ms-flexbox !important;
    display: flex !important;
}
.icon-wrapper{
    text-align: left;
    width: 24px;
    height: 20px;
    -ms-flex-negative: 0;
    flex-shrink: 0;
}
.text{
    max-width: 80%;
    overflow: hidden;
    text-overflow: ellipsis;
    white-space: nowrap;
    font-size: 12px;
    font-weight: 400;
    color: #40485b;
}
.addition{
    background: #f6f6f6;
}
.user-add-label-text{
    max-width: 80%;
    overflow: hidden;
    text-overflow: ellipsis;
    white-space: nowrap;
    margin: 0 4px;
}
</style>
