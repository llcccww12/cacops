<template>
  <div><router-view></router-view></div>
</template>

<script>
export default {
  name: "App",
};
</script>
<style>
body ::-webkit-scrollbar {
  -webkit-appearance: none;
  width: 7px;
  height: 7px
}

body ::-webkit-scrollbar-track {
  background: transparent;
  border-radius: 7px !important;
}

body ::-webkit-scrollbar-thumb {
  background: rgb(210, 210, 216) !important;
  border-radius: 7px !important;
}

body ::-webkit-scrollbar-thumb:hover {
  background: rgba(210, 210, 216, 0.8) !important;
  border-radius: 7px !important;
}

body ::-webkit-scrollbar-thumb:active {
  background: rgba(210, 210, 216, 0.8) !important;
  border-radius: 7px !important;
}
@media only screen and (max-width: 800px) {
  .el-message-box__wrapper .el-message-box {
    width: calc(100% - 20px);
  }
}
</style>
