<template>
    <div class="ui container">
        <div class="row git-user-content">
                <h3 class="ui header">
                    <div class="ui breadcrumb">
                        <a class="section" :href="url_code">{{i18n['code']}}</a>
                        <div class="divider"> / </div>
                        <div class="active section" >{{i18n['contributors']}}&nbsp;({{totalNum}})</div>
                    </div>
                </h3>
                <div class="ui horizontal relaxed list">
                        <div class="item user-list-item" v-for="(contributor,i) in contributors_list_page" >
                            <a v-if="contributor.user_name" :href="AppSubUrl +'/'+ contributor.user_name"><img class="ui avatar s16 image js-popover-card" :src="contributor.rel_avatar_link"></a>
                            <a v-else :href="'mailto:' + contributor.email "><img class="ui avatar s16 image js-popover-card" :avatar="contributor.email"></a>
                            <div class="content">
                                <div class="header" >
                                    <a class="username"  v-if="contributor.user_name" :href="AppSubUrl +'/'+ contributor.user_name">
                                        <img src="/images/gold.png" width="24" v-show="contributor.commit_cnt > 1000"/>
                                        <img src="/images/silver.png" width="24" v-show="contributor.commit_cnt > 500 && contributor.commit_cnt  < 1000"/>
                                        <img src="/images/bronze.png" width="24" v-show="contributor.commit_cnt > 50 && contributor.commit_cnt  < 500"/>
                                    {{contributor.user_name}}</a>
                                    <a v-else :href="'mailto:' + contributor.email ">{{contributor.email}}</a>
                                </div>
                                <span class="commit-btn">Commits: {{contributor.commit_cnt}}</span>
                            </div>
                        </div>
                </div>
        </div>

        <div class="ui container" style="margin-top:50px;text-align:center">
                <el-pagination
                    background
                    @current-change="handleCurrentChange"
                    :current-page="currentPage"
                    :page-size="pageSize"
                    layout="total, prev, pager, next"
                    :total="totalNum">
                </el-pagination>
        </div>
    </div>
</template>

<style>
.username{
    display: flex;
    align-items: center;
    img{
        margin-right: 10px;
    }
}</style>
<script>

const {AppSubUrl, StaticUrlPrefix, csrf} = window.config;

export default {

  data() {
    return {
        url:'',
        url_infor:'',
        href_:'',
        contributors_list:[],
        contributors_list_page:[],
        currentPage:1,
        pageSize:50,
        totalNum:0,
        AppSubUrl:AppSubUrl
    };
  },
methods: {

    getContributorsList(){
        this.$axios.get(this.url+'/list?'+this.url_infor).then((res)=>{
            this.contributors_list = res.data.contributor_info
            this.totalNum = this.contributors_list.length
            this.contributors_list_page = this.contributors_list.slice(0,this.pageSize)
        })
    },
    handleCurrentChange(val){
        this.contributors_list_page = this.contributors_list.slice((val-1)*this.pageSize,val*this.pageSize)
    },

},
computed:{

},
watch: {

},
created(){
    this.i18n = window.i18n;
    const url = window.location.pathname;
    this.url = url;
    let strIndex = this.url.indexOf("contributors")
    this.url_code = this.url.substr(0,strIndex)

    this.href_ = window.location.href;
    let index = this.href_.indexOf("?")
    this.url_infor = this.href_.substring(index+1,this.href_.length)
    this.getContributorsList()

},

updated(){
    if(document.querySelectorAll('img[avatar]').length!==0){
        window.LetterAvatar.transform()
    }
}
};
</script>

