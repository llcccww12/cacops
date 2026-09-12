<template>
   <div class="ui container main-c">
      <div class="cloudbrain-list-page">
        <div class="page-c">
            <div class="img-show"></div>
            <p class="tip-top">{{$t('cloudbrainObj.cbOffline')}}！</p>
            <p class="tip-bottom" v-html="$t('cloudbrainObj.cbOffflineTips')"></p>
        </div>
      </div>
   </div>
</template>
<script>
export default {
   components: {
     
   },
   props: {
     
   },
   data() {
     return {
       
     }
   },
   computed: {
     
   },
   watch: {
     
   },
   mounted() {
     
   },
   methods: {
     
   }
};
</script>
<style lang='less' scoped>
.cloudbrain-list-page{
    height: 500px;
    width: 100%;
    .page-c{
        display: flex;
        flex-direction: column;
        align-items: center;
        width: 100%;
        height: 100%;
        .img-show{
            margin-top: 60px;
            background-image: url('/img/svg-icon/cb-empty.svg');
            width: 100%;
            height: 200px;
            background-position: center center;     	
            background-repeat: no-repeat;	
            background-size: 100% 100%;
        }
        .tip-top{
            height: 24px;
            line-height: 24px;
            color: rgba(255,98,0,0.88);
            font-size: 18px;
            margin: 20px 0;
        }
        .tip-bottom{
            height: 24px;
            line-height: 24px;
            color: rgba(16,16,16,1);
            font-size: 14px;
        }
    }
}
</style>