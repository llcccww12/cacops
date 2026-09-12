<template>
  <div class="navbar">
    <div class="nanbar_cont justify-center flex-row items-center">
      <div id="nav1" class="flex-row items-center nb_section" @mouseover="addLineToImgSrc(1)"
        @mouseout="removeLineFromImgSrc(1)" :class="{ 'active-route': $route.name === 'HomePage' }">
        <img class="nb_img" :src="imgSrc(1)" />
        <a @click="gotoHomePage" class="nb_txt">首页</a>
      </div>
      <div class="line"></div>
      <div id="nav2" class="flex-row items-center nb_section" @mouseover="addLineToImgSrc(2)"
        @mouseout="removeLineFromImgSrc(2)" :class="{ 'active-route': $route.name === 'OSSystem' }">
        <img class="nb_img" :src="imgSrc(2)" />
        <a @click="gotoOSSystem" class="nb_txt">操作系统</a>
      </div>
      <div class="line"></div>
      <div id="nav3" class="flex-row items-center nb_section" @mouseover="addLineToImgSrc(3)"
        @mouseout="removeLineFromImgSrc(3)" :class="{ 'active-route': $route.name === 'OpenApp' }">
        <img class="nb_img" :src="imgSrc(3)" />
        <a @click="gotoOpenApp" class="nb_txt">开源应用</a>
      </div>
      <div class="line"></div>
      <div id="nav4" class="flex-row items-center nb_section" @mouseover="addLineToImgSrc(4)"
        @mouseout="removeLineFromImgSrc(4)" :class="{ 'active-route': $route.name === 'OpenDataset' }">
        <img class="nb_img" :src="imgSrc(4)" />
        <a @click="gotoOpenData" class="nb_txt">社区数据集</a>
      </div>
      <div class="line"></div>
      <div id="nav5" class="flex-row items-center nb_section" @mouseover="addLineToImgSrc(5)"
        @mouseout="removeLineFromImgSrc(5)" :class="{ 'active-route': $route.name === 'OpenModel' }">
        <img class="nb_img" :src="imgSrc(5)" />
        <a @click="gotoOpenModel" class="nb_txt">社区智能模型</a>
      </div>
      <div class="line"></div>
      <div id="nav6" class="flex-row items-center nb_section" @mouseover="addLineToImgSrc(6)"
        @mouseout="removeLineFromImgSrc(6)" :class="{ 'active-route': $route.name === 'CommunitySource' }">
        <img class="nb_img" :src="imgSrc(6)" />
        <a @click="gotoComunitySource" class="nb_txt">社区资源</a>
      </div>
      <div class="line"></div>
      <div id="nav7" class="flex-row items-center nb_section" @mouseover="addLineToImgSrc(7)"
        @mouseout="removeLineFromImgSrc(7)" :class="{ 'active-route': $route.name === 'HelpCenter' }">
        <img class="nb_img" :src="imgSrc(7)" />
        <a @click="gotoHelpCenter" class="nb_txt">帮助中心</a>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  data() {
    return {};
  },
  mounted() {
  },
  computed: {
    imgSrc() {
      return (id) => {
        const isActiveRoute = this.$route.name === this.getRouteNameById(id);
        return `/img/ros-hmci/icon${isActiveRoute ? '-' : ''}${id}.png`;
      };
    },
  },
  methods: {
    getRouteNameById(id) {
      const routeNames = {
        1: 'HomePage',
        2: 'OSSystem',
        3: 'OpenApp',
        4: 'OpenDataset',
        5: 'OpenModel',
        6: 'CommunitySource',
        7: 'HelpCenter',
      };
      return routeNames[id] || '';
    },
    addLineToImgSrc(id) {
      if (this.$route.name === this.getRouteNameById(id)) return;

      const imgElement = document.getElementById(`nav${id}`).getElementsByTagName('img')[0];
      imgElement.src = imgElement.src.replace(`icon${id}`, `icon-${id}`);
    },
    removeLineFromImgSrc(id) {
      if (this.$route.name === this.getRouteNameById(id)) return;

      const imgElement = document.getElementById(`nav${id}`).getElementsByTagName('img')[0];
      imgElement.src = imgElement.src.replace(`icon-${id}`, `icon${id}`);
    },
    gotoHomePage() {
      // console.log(this.$router);
      this.$router.push("/");
    },
    gotoOSSystem() {
      // console.log(this.$router);
      this.$router.push("/os-system");
    },
    gotoOpenApp() {
      // console.log(this.$router);
      this.$router.push("/open-app");
    },
    gotoOpenData() {
      // console.log(this.$router);
      this.$router.push("/open-data");
    },
    gotoOpenModel() {
      // console.log(this.$router);
      this.$router.push("/open-model");
    },
    gotoComunitySource() {
      // console.log(this.$router);
      this.$router.push("/community-source");
    },
    gotoHelpCenter() {
      // console.log(this.$router);
      this.$router.push("/help-center");
    },
  },
};
</script>

<style scoped>
.nb_section:hover {
  background-image: linear-gradient(180deg, #1ad1c3 0%, #02e588 100%);
  /* 规定背景绘制区域*/
  -webkit-background-clip: text;
  -webkit-text-fill-color: transparent;
}

.nb_section:hover,
.active-route {
  background-image: linear-gradient(180deg, #1ad1c3 0%, #02e588 100%);
  /* 规定背景绘制区域*/
  -webkit-background-clip: text;
  -webkit-text-fill-color: transparent;
}

.nb_txt {
  font-family: PingFang SC;
  color:#4d4d4d;
  font-size: 18px;
  text-align: center;
}

.nb_img {
  width: 17px;
  height: 17px;
  margin-right: 7px;
}

.line {
  margin-left: 46px;
  margin-right: 46px;
  opacity: 80%;
  width: 0px;
  height: 15.4px;
  border: 1px solid;
  border-color: rgba(77, 77, 77, 0.28);
}

.navbar {
  display: inline-flex;
  align-items: center;
  justify-content: center;
  flex-direction: row;
  background-color:rgba(205, 219, 228, 0.5);
  width: 100%;
  height: 108px;
  position: absolute;
  left: 0;
  bottom: 0px;
  z-index: 100;
}

.info {
  position: absolute;
  top: 104px;
  left: 18%;
}

.info>*:not(:first-child) {
  margin-top: 36px;
}

.info-div {
  width: 486.09px;
  height: 53px;
  background-image: linear-gradient(150.44deg, #00e850 0%, #0100be 100%);
  border-radius: 28px;
}
.nanbar_cont{
  width: 1200px;
  margin: 0 auto;
}
</style>
