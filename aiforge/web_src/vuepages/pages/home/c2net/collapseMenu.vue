<template>
   <div class="nav-bar-content">
      <div class="header-wrap" @click="openFlag = !openFlag">
            <span class="text">{{title}}</span>
            <i class="ri-arrow-right-s-line arrow" :class="{ open: openFlag }"></i>
      </div>
  
      <div v-show="openFlag" class="content">
            <slot></slot>
        </div>
  </div>
</template>
          
<script>
export default{
    props:{
        title: {
            type: String,
            default: '计算资源'
        },
        open: {
            type: Boolean,
            default: false
        }
    },
    data() {
        return {
            openFlag: this.open
        };
    },
    methods: {
        
    },
}
</script>
