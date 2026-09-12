<template>
  <div class="group">
    <NavigationBar></NavigationBar>
    <div class="info">
      <span class="info1">人机协同智能操作系统</span>
      <div class="info-div">
        <span class="info2">智能操作软件库和工具  构建人工智能应用程序</span>
      </div>
      <span class="info3">ROS - hmci</span>
    </div>
  </div>
</template>

<script>
import NavigationBar from "./NavigationBar.vue";
export default {
  name: "AppBanner",
  props: {},
  components: {
    NavigationBar,
  },
  data() {
    return {
      list: [],
    };
  },
  methods: {},
  mounted() { },
};
</script>

<style scoped>
.group {
  position: relative;
  top: 0px;
  background: url(/img/ros-hmci/mbz605.png);
  background-position: center;
  background-size: cover;
  background-repeat: no-repeat;
  height: 600px;
  width: 100%;
}

.info {
  display: flex;
  flex-direction: column;
  justify-content: center;
  top: 165px;
  width: 1200px;
  margin: 0 auto;
  position: relative;
}

.info>*:not(:first-child) {
  margin-top:36px;
}

.info-div {
  display: inline-flex;
  flex-direction: row;
  align-items: center;
  justify-content: center;
  width: 486.09px;
  height: 53px;
  background-image:linear-gradient(150.44deg,#0feb8d 0%,#50c6a6 100%);
  border-radius: 28px;
}

.info1 {
  font-family: sucaijishikufangti;
  font-weight:700;
  color:#444b66;
  font-size:40px;
  letter-spacing: 7px;
}

.info2 {
  font-family: Alibaba PuHuiTi;
  color: #ffffff;
  font-size: 21px;
  letter-spacing: 1px;
}

.info3 {
  font-family: Alibaba PuHuiTi;
  color:#444b66;
  font-size:36px;
  letter-spacing: 4px;
}
</style>
