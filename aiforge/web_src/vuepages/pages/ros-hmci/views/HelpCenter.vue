<template>
  <div class="flex-col relative page">
    <div class="banner flex-col relative">
      <NavigationBar></NavigationBar>
      <div class="flex-col pos_6">
        <span class="self-center text_6">在这里,<br/>我们为你提供帮助与支持！</span>
      </div>
    </div>

    <img class="image_8 pos_5" src="/img/ros-hmci/group5727@2x.png" />
    <div class="flex-col relative space-y-64">
      <div class="self-center">
        <div class="flex-row group_3 space-x-14">
          <div class="flex-col shrink-0 self-start relative section_7">
            <div class="flex-col justify-start items-start text-wrapper_2">
              <span class="font_2 text_8">帮助中心</span>
            </div>
            <div class="self-center section_9"></div>
            <div class="flex-row group_6 space-x-8">
              <img class="image_10" src="/img/ros-hmci/mbz640.png" />
              <span class="font_4">常见问题</span>
            </div>
            <div class="flex-row group_6 space-x-8">
              <img class="image_10" src="/img/ros-hmci/mbz637.png" />
              <span class="font_4">入门指南</span>
            </div>
            <div class="flex-row group_6 space-x-8">
              <img class="image_10" src="/img/ros-hmci/mbz636.png" />
              <span class="font_4">安装手册</span>
            </div>
            <div class="flex-row group_6 space-x-8">
              <img class="image_10" src="/img/ros-hmci/mbz635.png" />
              <span class="font_4">API接口</span>
            </div>
          </div>
          <helpCollaps></helpCollaps>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import NavigationBar from "../components/NavigationBar.vue";
import helpCollaps from "../components/helpCollaps.vue";

export default {
  name: "HelpCenter",
  data() {
    return {};
  },
  components: {
    NavigationBar,
    helpCollaps,
  },
  methods: {},
};
</script>

<style scoped>



.page {
  margin-bottom: -80px;
  background-color: #f3f5f8;
  width: 100%;
  overflow-y: hidden;
  /* overflow-x: hidden; */
  height: auto;
}

.section {
  padding: 17px 80px 17px 123px;
  filter: drop-shadow(0px 2px 3px #0137931f);
  background-color: #282828;
}

.banner {
  background-image: url("/img/ros-hmci/group7348.png");
  background-position: center;
  background-size: cover;
  background-repeat: no-repeat;
  height: 600px;
  margin-top: -1px;
}

.space-x-40>*:not(:first-child) {
  margin-left: 40px;
}

.space-x-48>*:not(:first-child) {
  margin-left: 48px;
}

.space-x-16>*:not(:first-child) {
  margin-left: 16px;
}

.image_3 {
  width: 63px;
  height: 16px;
}

.section_3 {
  background-color: #ffffff;
  width: 2px;
  height: 20.5px;
}

.text_72 {
  font-family: PingFang SC;
  color: #ffffff;
  font-size: 18px;
  text-align: center;
}

.image_41 {
  width: 20px;
  height: 20px;
  margin-right: 5px;
}

.line {
  margin-left: 46px;
  margin-right: 46px;
  opacity: 80%;
  width: 0px;
  height: 15.4px;
  border: 1px solid;
  border-color: #ffffff;
}

.pos_87 {
  align-items: center;
  justify-content: center;
  flex-direction: row;
  display: inline-flex;
  position: absolute;
  top: 355px;
  background-color: rgba(255, 255, 255, 0.15);
  height: 108px;
  width: 100%;
  backdrop-filter: blur(20px);
  z-index: 1;
}

.text {
  color: #ffffff;
  font-size: 22px;
  font-family: AlimamaShuHeiTi-Bold;
  font-weight: 700;
  line-height: 20.5px;
  letter-spacing: 1px;
}

.font_1 {
  font-size: 15px;
  font-family: PingFang SC Regular;
  /* line-height: 14px; */
  color: #ffffff;
}

.text_3 {
  opacity: 0.9;
}

.group {
  width: 368px;
}

.group_2 {
  width: 343px;
}

.view {
  margin-left: 24px;
}

.image_6 {
  opacity: 0.63;
  width: 14px;
  height: 14px;
}

.space-x-124>*:not(:first-child) {
  margin-left: 124px;
}

.text_4 {
  opacity: 0.63;
}

.section_4 {
  background-color: #ffffff7a;
  width: 1px;
  height: 13px;
}

.image {
  margin-left: 20px;
  width: 30px;
  height: 30px;
}

.text_5 {
  margin-left: 10px;
  font-size: 14px;
  line-height: 13px;
}

.image_7 {
  margin-left: 8px;
  border-radius: 1px;
  width: 11px;
  height: 6px;
}

.image_5 {
  margin-left: 20px;
  width: 20px;
  height: 19px;
}

.text-wrapper {
  padding: 4px 0;
  background-color: #d92b2f;
  border-radius: 50%;
  width: 12px;
}

.pos_6 {
  position: relative;
  top: 184px;
}

.pos_2 {
  position: absolute;
  left: 116px;
  top: 0;
}

.text_2 {
  color: #ffffff;
  font-size: 8px;
  font-family: AlibabaPuHuiTi;
  line-height: 5.5px;
}

.image_2 {
  opacity: 0.7;
  width: 20px;
  height: 20px;
}

.pos_3 {
  position: absolute;
  left: 140.5px;
  top: 50%;
  transform: translateY(-50%);
}

.image_4 {
  opacity: 0.7;
  width: 21px;
  height: 19px;
}

.pos_4 {
  position: absolute;
  left: 99.5px;
  top: 50%;
  transform: translateY(-50%);
}

.section_2 {
  border-radius: 15px;
  opacity: 0.7;
  height: 30px;
  border: solid 1px #ffffff;
}

.pos {
  position: absolute;
  left: 0;
  right: 264px;
  top: 0;
}

.image_8 {
  width: 100%;
  /* height: 1522px; */
}

.pos_5 {
  position: absolute;
  left: 0;
  right: 0;
  top: 469px;
}

.space-y-64>*:not(:first-child) {
  margin-top: 64px;
}

.text_6 {
  font-family: Alibaba PuHuiTi;
  color: #4d4d4d;
  font-size: 40px;
  letter-spacing: 8px;
  line-height: normal;
  text-align: center;
  position: relative;
  top: -3vh;
}

.section_6 {
  margin-top: 48px;
  padding-left: 36px;
  background-color: #ffffff;
  border-radius: 12px;
  box-shadow: 0px 3px 12px #0000000a;
  width: 975px;
}

.text_7 {
  color: #757c8d;
}

.image_9 {
  border-radius: 4px 10px 10px 4px;
  width: 79px;
  height: 56px;
}

.group_3 {
  margin: 60px 0;
}

.space-x-14>*:not(:first-child) {
  margin-left: 14px;
}

.section_7 {
  padding-bottom: 40px;
  background-color: #ffffff;
  border-radius: 8px;
  box-shadow: 0px 3px 12px #0000000a;
  height: auto;
}

.text-wrapper_2 {
  padding: 24px 0 14px;
  background-image: url("/img/ros-hmci/bb066358d55b3baa6c30aeb967d32e50.png");
  background-size: 100% 100%;
  background-repeat: no-repeat;
  width: 236px;
}

.font_2 {
  font-size: 15px;
  font-family: AlibabaPuHuiTi;
  line-height: 14px;
  color: #4d4d4d;
}

.text_8 {
  margin-left: 20px;
  color: #181818;
  font-size: 16px;
  line-height: 14.5px;
}

.section_9 {
  margin-top: 7px;
  background-image: repeating-linear-gradient(90deg,
      #caced8,
      #caced8 1.06%,
      transparent 1.06%,
      transparent 2.124%);
  background-position: -1px 0px;
  width: 188px;
  height: 1px;
}

.group_4 {
  margin-top: 22px;
}

.image_10 {
  width: 18px;
  height: 18px;
}

.text_10 {
  color: #24a19b;
}

.group_6 {
  margin-top: 26px;
  padding: 0 18px;
}

.space-x-10>*:not(:first-child) {
  margin-left: 10px;
}

.image_12 {
  width: 17px;
  height: 14px;
}

.font_4 {
  font-size: 15px;
  font-family: PingFang SC Regular;
  color: #4d4d4d;
}

.font_4.now {
  color: #24a19b;
}

.image_13 {
  width: 14px;
  height: 16px;
}

.group_7 {
  margin-top: 24px;
  padding: 0 20px;
}

.space-x-8>*:not(:first-child) {
  margin-left: 8px;
}

.image_14 {
  width: 14px;
  height: 15px;
}

.section_8 {
  padding: 12px 24px 114px;
  background-color: #ffffff;
  border-radius: 6px;
  box-shadow: 0px 3px 12px #0000000a;
  width: 949px;
  height: auto;
  border: solid 1px #ebf7f6;
}

.font_3 {
  font-size: 20px;
  font-family: AlibabaPuHuiTi;
  line-height: 18px;
  color: #181818;
}

.text_9 {
  margin-left: 2px;
  line-height: 18.5px;
}

.group_5 {
  margin-top: 32px;
}

.image_11 {
  width: 17px;
  height: 17px;
}

.text_11 {
  font-size: 16px;
  line-height: 14.5px;
}

.section_10 {
  margin-right: 16px;
  margin-top: 18px;
  padding: 24px 24px 26px;
  background-color: #f3f4f6;
  border-radius: 6px;
}

.space-y-24>*:not(:first-child) {
  margin-top: 24px;
}

.space-y-2>*:not(:first-child) {
  margin-top: 2px;
}

.font_5 {
  font-size: 15px;
  font-family: AlibabaPuHuiTi;
  line-height: 18px;
  color: #757c8d;
}

.font_6 {
  font-size: 15px;
  font-family: AlibabaPuHuiTi;
  line-height: 14px;
  color: #757c8d;
}

.divider {
  background-image: repeating-linear-gradient(90deg,
      #caced8,
      #caced8 0.23%,
      transparent 0.23%,
      transparent 0.456%);
  background-position: -1px 0px;
  height: 1px;
}

.view_2 {
  margin-right: 16px;
  margin-top: 14px;
}

.group_8 {
  padding: 18px 0;
}

.text_12 {
  font-size: 16px;
  line-height: 14.5px;
}

.view_3 {
  margin-right: 16px;
}

.text_13 {
  font-size: 16px;
  line-height: 14.5px;
}

.view_4 {
  margin-right: 16px;
}

.text_14 {
  font-size: 16px;
  line-height: 14.5px;
}

.view_5 {
  margin-right: 16px;
}

.text_15 {
  margin-top: 36px;
}

.group_9 {
  margin-top: 32px;
}

.text_16 {
  font-size: 16px;
  line-height: 14.5px;
}

.view_6 {
  margin-right: 16px;
  margin-top: 18px;
}

.text_17 {
  font-size: 16px;
  line-height: 14.5px;
}

.view_7 {
  margin-right: 16px;
}

.text_18 {
  margin-top: 36px;
  line-height: 18.5px;
}

.group_10 {
  margin-top: 32px;
}

.text_19 {
  font-size: 16px;
  line-height: 15px;
}

.view_8 {
  margin-right: 16px;
  margin-top: 18px;
}

.text_20 {
  font-size: 16px;
  line-height: 15px;
}

.view_9 {
  margin-right: 16px;
}

.text_21 {
  margin-top: 36px;
}

.group_11 {
  margin-top: 32px;
}

.space-x-6>*:not(:first-child) {
  margin-left: 6px;
}

.text_22 {
  font-size: 16px;
  line-height: 14.5px;
}

.view_10 {
  margin-right: 16px;
  margin-top: 18px;
}

.section_11 {
  background-color: #1b2440;
}

.group_12 {
  margin-right: 57px;
  padding: 65px 0 60px;
}

.image-wrapper_2 {
  padding: 22px 0 14px;
  background-color: #ffffff;
  border-radius: 13px;
  width: 100px;
}

.image_15 {
  width: 91px;
  height: 64px;
}

.text_29 {
  margin-top: 32px;
  font-size: 16px;
  line-height: 15px;
}

.font_8 {
  font-size: 15px;
  font-family: PingFang SC Regular;
  line-height: 14px;
  color: #bdc2d1;
}

.text_32 {
  margin-top: 26px;
  font-size: 16px;
  line-height: 15.5px;
}

.group_15 {
  margin-left: 194px;
  margin-top: 2px;
  padding-bottom: 58px;
}

.space-x-148>*:not(:first-child) {
  margin-left: 148px;
}

.space-y-30>*:not(:first-child) {
  margin-top: 30px;
}

.font_7 {
  font-size: 18px;
  font-family: 苹方;
  line-height: 16.5px;
  color: #ffffff;
}

.space-y-22>*:not(:first-child) {
  margin-top: 22px;
}

.text_23 {
  font-size: 14px;
  line-height: 13px;
}

.text_25 {
  font-size: 14px;
  line-height: 13px;
}

.text_27 {
  font-size: 14px;
  line-height: 13px;
}

.text_24 {
  font-size: 14px;
  line-height: 13px;
}

.text_26 {
  font-size: 14px;
  line-height: 13px;
}

.text_28 {
  font-size: 14px;
  line-height: 13px;
}

.group_17 {
  margin-left: 218px;
}

.text-wrapper_3 {
  padding-bottom: 26px;
  width: 99px;
}

.space-x-76>*:not(:first-child) {
  margin-left: 76px;
}

.space-y-18>*:not(:first-child) {
  margin-top: 18px;
}

.image_16 {
  border-radius: 4px;
  width: 99px;
  height: 100px;
}

.text_30 {
  font-size: 14px;
  line-height: 13px;
}

.text_31 {
  font-size: 14px;
  line-height: 13px;
}

.section_12 {
  padding: 26px 0;
  background-color: #222a42;
}

.space-y-6>*:not(:first-child) {
  margin-top: 6px;
}

.group_16 {
  width: 389.5px;
  height: 13px;
}

.font_9 {
  font-size: 12px;
  font-family: PingFang SC Regular;
  line-height: 14px;
  color: #bdc2d1;
}

.text_33 {
  line-height: 13px;
}

.text_34 {
  margin-top: -12px;
  line-height: 13px;
}

.text_35 {
  text-align: center;
  width: 380px;
}
</style>
