<template>
  <div class="form-container">
    <div class="form-head">
        <h4>{{ title }}</h4>
    </div>
    <div class="form-content">
        <slot></slot>
    </div>
  </div>
</template>

<script>
export default {
  name: 'BaseTitle',
  props: {
    title: { type: String, default: '' },
  },
  data() {
    return {
      
    };
  },

  mounted() {
    
  },

  methods: {
    
  },
};
</script>
<style lang='less' scoped>
.form-container {
  border-radius: 10px;
  background-color: rgba(255,255,255,1);
  box-shadow: 0px 2px 6px 0px rgba(178,192,255,0.5);
  margin-top: 20px;
  border: 1px solid rgba(224,227,234,1);
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
  
}
@media screen and (max-width: 767px) {
  .form-container{
    border-radius: 0 !important;
    border: none !important;
    margin-top: 0 !important;
    margin-bottom: 24px !important;
  }
  .form-content {
    margin: 0 10px;
    padding: 10px 0;
  }
}
</style>