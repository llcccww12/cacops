<template>
    <div class="ui container">
      <el-steps :active="formactive" class="steps" finish-status="finish" v-if="showSteps">
          <el-step :title="title[0]"></el-step>
          <el-step :title="title[1]"></el-step>
      </el-steps>
        <div class="form-container">
            <div class="form-head">
                <h4>{{ formTtitle }}</h4>
            </div>
            <div class="form-content">
                <slot></slot>
            </div>
        </div>
    </div>
</template>
  
<script>

export default {
    name: "CreateForm",
    props: {
        title: { type: Array, default: () => [] },
        active: {  type: Number, default: 0 },
        showSteps: {  type: Boolean, default: true }
    },
    components: { },
    data() {
        return {
          formactive: 0,
          formTtitle: this.title[0]
        }
    },
    watch:{
        active(val){
          this.formactive = val
          this.formTtitle = this.title[val]
        }
    },
    methods: {

    },
    mounted() { 
      
    },
}
</script>

<style lang='less' scoped>
.form-container {
  border-radius: 10px;
  background-color: rgba(255,255,255,1);
  box-shadow: 0px 2px 6px 0px rgba(178,192,255,0.5);
  .form-head {
    height: 45px;
    display: flex;
    align-items: center;
    h4{
      margin-left: 20px;
      border-bottom: 2px solid rgba(16,16,16,0.8);;
      height: 100%;
      line-height: 48px;
      color: rgba(16,16,16,0.8);;
    }
  }
  .form-content {
    padding: 3rem 12rem 3rem 2rem;
    
  }
}
.steps {
  width: 100%;
  justify-content: center;
  align-items: center;
  height: 35px;
  margin-bottom: 50px;
  ::v-deep .el-step {
    height: 100%;
    // 设置图标和步骤条的行高
    .el-step__head {
      line-height: 35px;
    }
    .el-step__icon{
      width: 36px;
      height: 36px;
      z-index: 99;
      font-size: 18px;
    }
    // 步骤条
    .el-step__line {
      background-color: rgba(0, 0, 0, 0.15);
      top: 50%;
      left: 150px;
      height: 1px;
      right: 20px;
    }
    .el-step__head.is-process {
      color: #0066ff;
      border-color: #0066ff;
      .el-step__icon{
        border: 1px solid #6195f7;
        color: #6195f7;
      }
    }
      // title样式
    .el-step__title {
      z-index: 66;
      position: absolute;
      top: 0px;
      left: calc(10%);
      font-size: 14px;
      z-index: 66;
      color: #404040;
    }
    .el-step__title.is-process{
      color:#3291f8;
    }
    // 已完成步骤条的边框色和字体颜色
    .el-step__head.is-finish {
      color: #0066ff;
      border-color: #0066ff;
      .el-step__icon{
        background-color: rgba(50,145,248,0.1);
      }
    }
    .el-step__title.is-finish {
      // 已完成步骤的title
      font-weight: 700;
      color: #3291f8;
    }

    
  }
  // 第一个步骤
  ::v-deep .el-step:first-child {
    flex-basis: 20% !important;
    .el-step__head.is-process {
      padding-left: 10px !important;
    }
    .el-step__head.is-success {
      padding-left: 10px !important;
    }
    .el-step__title {
      padding-left: 26px !important;
    }

  }
  // 第二个步骤
  ::v-deep .el-step:nth-child(2) {
    flex-basis: 20% !important;
    .el-step__title {
      padding-left: 16px !important;
    }
  }
}
@media only screen and (max-width: 1200px) {
  .form-container .form-content  {
    padding: 1rem 1rem 1rem 0;
  }
}
@media only screen and (max-width: 768px) {
  .form-container .form-content  {
    padding: 0;
  }
}
</style>
  