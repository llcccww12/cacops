<template>
    <div>
        <AppBanner></AppBanner>
        <div class="flex-col">
            <div class="flex-col items-start section_6 space-y-5">
                <div class="flex-row items-center space-x-14">
                    <img class="shrink-0 image_10" src="/img/ros-hmci/aabc8fe78ef76d15c2805169fa04c4c9.png" />
                    <span class="text_15">资源列表/人机协同侦察系统</span>
                </div>
                <div class="flex-row justify-center group_3 space-x-10">
                    <div class="flex-col justify-start text-wrapper_2">
                        <span class="font_3 text_16">人机协同智能操作系统...</span>
                    </div>
                    <div class="flex-col justify-start text-wrapper_2">
                        <span class="font_3 text_17">虚拟学习引擎</span>
                    </div>
                    <div class="flex-col justify-start text-wrapper_2">
                        <span class="font_3 text_18">作业环境仿真器与数字...</span>
                    </div>
                </div>
            </div>
            <div class="flex-col section_7 space-y-17">
                <div class="flex-col section_9">
                    <div class="flex-col items-start section_8">
                        <div class="flex-row justify-center group_4 space-x-36">
                            <div class="flex-row items-start justify-center">
                                <img class="shrink-0 image_11" src="/img/ros-hmci/mbz628.png" />
                                <span class="text_19">概述</span>
                            </div>
                        </div>
                        <div class="section_10"></div>
                    </div>
                    <div class="flex-col group_5">
                        <div class="flex-row items-center self-start space-x-16">
                            <span class="font_4 text_bq">标签</span>
                            <span class="font_2">
                                {{ resourceDetails.name }}
                            </span>
                        </div>
                        <div class="flex-row items-center self-start group_7">
                            <span class="font_4">版本号</span>
                            <span class="font_2 text_331">{{ resourceDetails.version }}</span>
                        </div>
                        <div class="flex-row items-center group_7">
                            <span class="font_4">资源地址</span>
                            <img class="image_13" src="/img/ros-hmci/e6222f48f76ae9ba015d4748d05ca0f6.png" />
                            <a :href="resourceDetails.address" target="_blank" class="font_2">{{
                                resourceDetails.address
                            }}</a>
                        </div>
                        <div class="flex-row items-center self-start group_7 space-x-16">
                            <span class="font_4">开源协议</span>
                            <span class="font_2">{{ resourceDetails.License }}</span>
                        </div>
                        <div class="flex-row items-center self-start group_7 space-x-16">
                            <span class="font_4">最后更新</span>
                            <span class="font_2">{{ resourceDetails.update_time }}</span>
                        </div>
                        <div class="flex-row items-center self-start group_7 space-x-16">
                            <span class="font_4">资源简介</span>
                            <span class="font_2">{{ resourceDetails.synopsis }}</span>
                        </div>
                    </div>
                </div>
                <div class="flex-col markdown">
                    <div class="flex-col items-start justify-center section_88">
                        <span class="text_detail">详细介绍</span>
                    </div>

                    <div class="md_content" v-html="htmlContent"></div>
                </div>
            </div>
        </div>
    </div>
</template>

<script>
import AppBanner from "../components/AppBanner.vue";
import axios from "axios";
import MarkdownIt from "markdown-it";
export default {
    name: "ResourceDetail",
    props: {
        resourceName: {
            type: String,
            required: true,
        },
    },
    data() {
        return {
            resourceDetails: {},
            htmlContent: "",
        };
    },
    async created() {
        const resourceName = this.$route.params.name;
        const response = await fetch(
            `/fanshuai/ROS-hmci-resource/raw/branch/master/resource/${resourceName}.json`
        );
        const data = await response.json();
        this.resourceDetails = data[0]; // 假设响应中的第一个对象包含所需的详细信息
    },

    components: {
        AppBanner,
    },

    methods: {},
    async mounted() {
        try {
            const resourceName = this.$route.params.name;
            const response = await axios.get(
                `/fanshuai/ROS-hmci-resource/raw/branch/master/markdown/${resourceName}.md`
            );
            this.markdownContent = response.data;

            const md = new MarkdownIt();
            this.htmlContent = md.render(this.markdownContent);
        } catch (error) {
            console.error(error);
        }
    },
};
</script>

<style scoped>
.text_detail {
    font-family: Alibaba PuHuiTi;
    color: #24a19b;
    font-size: 18px;
}

.md_content {
    padding: 10px 20px;
}

.markdown {
    width: 1200px;
    background-color: #ffffff;
    border-radius: 6px;
    box-shadow: 0px 3px 12px #0000000a;
    border: solid 1px #ebf7f6;
}

.page {
    background-color: #f3f5f8;
    width: 100%;
    overflow-y: auto;
    overflow-x: hidden;
    height: 100%;
}

.section {
    padding: 17px 80px 12px 123px;
    filter: drop-shadow(0px 2px 3px #0137931f);
    background-color: #282828;
}

.space-x-40>*:not(:first-child) {
    margin-left: 40px;
}

.image_3 {
    width: 63px;
    height: 16px;
}

.section_3 {
    margin-left: 14px;
    background-color: #ffffff;
    width: 2px;
    height: 20.5px;
}

.text {
    margin-left: 18px;
    color: #ffffff;
    font-size: 22px;
    font-family: AlimamaShuHeiTi-Bold;
    font-weight: 700;
    line-height: 20.5px;
    letter-spacing: 1px;
}

.font_1 {
    font-size: 14px;
    font-family: PingFang SC Regular;
    line-height: 14.5px;
    color: #ffffff;
}

.text_3 {
    opacity: 0.9;
}

.text_4 {
    margin-left: 48px;
    font-size: 15px;
    line-height: 14px;
}

.text_5 {
    margin-left: 40px;
    font-size: 15px;
    line-height: 14px;
}

.text_6 {
    margin-left: 38px;
    font-size: 15px;
    line-height: 14px;
}

.text_7 {
    margin-left: 36px;
    font-size: 15px;
    line-height: 14px;
}

.text_8 {
    margin-left: 38px;
    font-size: 15px;
    line-height: 14px;
}

.image_6 {
    margin-left: 38px;
    width: 61px;
    height: 26px;
}

.text_9 {
    margin-left: 40px;
    margin-top: 2px;
    font-size: 15px;
    line-height: 14px;
}

.text_10 {
    margin-left: 38px;
    margin-top: 2px;
    font-size: 15px;
    line-height: 14px;
}

.text_11 {
    margin-left: 38px;
    margin-top: 2px;
    font-size: 15px;
    line-height: 14px;
}

.section_2 {
    border-radius: 15px;
    opacity: 0.7;
    height: 30px;
    border: solid 1px #ffffff;
}

.text_12 {
    font-size: 15px;
    line-height: 14px;
    opacity: 0.63;
}

.pos_2 {
    position: absolute;
    right: 24px;
    top: 50%;
    transform: translateY(-50%);
}

.image_7 {
    opacity: 0.63;
    width: 14px;
    height: 14px;
}

.pos {
    position: absolute;
    left: 25px;
    top: 50%;
    transform: translateY(-50%);
}

.group_2 {
    margin-left: 20px;
    width: 28.5px;
    height: 25px;
}

.image_4 {
    margin-top: 6px;
    opacity: 0.7;
    width: 21px;
    height: 19px;
}

.text-wrapper {
    margin-top: -25px;
    padding: 4px 0;
    background-color: #d92b2f;
    border-radius: 50%;
    width: 12px;
}

.text_2 {
    color: #ffffff;
    font-size: 8px;
    font-family: AlibabaPuHuiTi;
    line-height: 5.5px;
}

.image_2 {
    margin-left: 12px;
    opacity: 0.7;
    width: 20px;
    height: 20px;
}

.section_4 {
    margin-left: 20px;
    background-color: #ffffff7a;
    width: 1px;
    height: 13px;
}

.image {
    margin-left: 20px;
    width: 30px;
    height: 30px;
}

.text_13 {
    margin-left: 10px;
    line-height: 13px;
}

.image_8 {
    margin-left: 8px;
    border-radius: 1px;
    width: 11px;
    height: 6px;
}

.image_5 {
    margin-left: 20px;
    opacity: 0.7;
    width: 20px;
    height: 19px;
}

.section_5 {
    padding: 10px 360px;
    background-color: #757c8d26;
}

.image_9 {
    width: 33px;
    height: 21px;
}

.font_2 {
    font-size: 14px;
    font-family: AlibabaPuHuiTi;
    line-height: 13px;
    color: #757c8d;
}

.text_14 {
    color: #4d4d4d;
}

.section_6 {
    padding: 30px 18.75%;
    background-color: #f6f9fb;
}

.space-y-5>*:not(:first-child) {
    margin-top: 5px;
}

.image_10 {
    border-radius: 50%;
    width: 42px;
    height: 42px;
}

.text_15 {
    color: #181818;
    font-size: 20px;
    font-family: AlibabaPuHuiTi;
    line-height: 18.5px;
}

.group_3 {
    margin-left: 55px;
    width: 446px;
}

.space-x-10>*:not(:first-child) {
    margin-left: 10px;
}

.text-wrapper_2 {
    padding: 6px 0;
    background-color: #e5b37229;
    border-radius: 4px;
    height: 26px;
}

.font_3 {
    font-size: 14px;
    font-family: AlibabaPuHuiTi;
    line-height: 13px;
    color: #bb7f41;
}

.text_16 {
    margin-left: 6px;
}

.text_17 {
    margin-left: 8px;
    margin-right: 4px;
}

.text_18 {
    margin-left: 8px;
    margin-right: 2px;
}

.section_7 {
    align-items: center;
    padding-top: 20px;
    background-color: #ffffff;
    border-radius: 4px;
}

.space-y-17>*:not(:first-child) {
    margin-top: 17px;
}

.section_9 {
    width: 1200px;
    background-color: #ffffff;
    border-radius: 6px;
    box-shadow: 0px 3px 12px #0000000a;
    border: solid 1px #ebf7f6;
}

.section_8 {
    padding: 10px 20px 10px;
    background-color: #fafcff;
    border-radius: 4px;
    border: solid 1px #caced829;
}

.section_88 {
    padding: 10px 20px 10px;
    background-color: #fafcff;
    border-radius: 4px;
    border: solid 1px #caced829;
}

.group_4 {
    padding: 20px 0 10px;
}

.space-x-36>*:not(:first-child) {
    margin-left: 36px;
}

.space-x-6>*:not(:first-child) {
    margin-left: 6px;
}

.image_11 {
    width: 18px;
    height: 18px;
}

.font_4 {
    /* width: 48px; */
    height: 22px;
    font-family: Alibaba PuHuiTi;
    color: #181818;
    font-size: 16px;
    text-align: center;
    letter-spacing: 0px;
}

.text_19 {
    /* color: #24a19b; */
    /* width: 32px; */
    /* height: 22px; */
    font-family: Alibaba PuHuiTi;
    color: #24a19b;
    font-size: 16px;
    /* text-align: center; */
    margin-left: 5px;
}

.image_12 {
    width: 17px;
    height: 14px;
}

.text_20 {
    color: #323232;
}

.section_10 {
    background-image: linear-gradient(90.8deg, #1ad1c3 8.9%, #02e588 81.8%);
    border-radius: 15px;
    width: 49px;
    height: 3px;
}

.group_5 {
    margin: 20px 26px 11px 23px;
    height: 242px;
}

.text_21 {
    margin-top: -14px;
    line-height: 15px;
}

.group_6 {
    margin-top: 0px;
}

.space-x-16>*:not(:first-child) {
    margin-left: 25px;
}

.group_7 {
    margin-top: 20px;
}

.font_5 {
    font-size: 14px;
    font-family: AlibabaPuHuiTi;
    line-height: 16.5px;
    color: #757c8d;
}

.text_22 {
    line-height: 17px;
}

.divider {
    margin-top: 30px;
    background-color: #caced8;
    height: 0px;
}

.text_23 {
    margin-top: 30px;
    line-height: 15px;
}

.group_8 {
    margin-top: 28px;
}

.text_24 {
    width: 62px;
    /* height        : 22px; */
    font-family: Alibaba PuHuiTi;
    color: #181818;
    font-size: 16px;
    /* letter-spacing: 8px; */
    text-align: center;
}

.image_13 {
    margin-left: 27px;
    width: 11px;
    height: 11px;
}

.text_25 {
    margin-left: 6px;
    color: #24a19b;
    line-height: 13.5px;
}

.text_bq {
    letter-spacing: 15px;
}

.text_331 {
    margin-left: 40px;
}

.group_9 {
    margin-top: 18px;
}

.text_26 {
    line-height: 15px;
}

.text_27 {
    line-height: 10px;
}

.group_10 {
    margin-top: 20px;
}

.space-x-14>*:not(:first-child) {
    margin-left: 14px;
}

.text_28 {
    line-height: 12.5px;
}

.group_11 {
    margin-top: 20px;
}

.text_29 {
    margin-left: 28px;
}

.text_30 {
    margin-left: 18px;
}

.group_12 {
    margin-top: 18px;
}

.text_31 {
    margin-left: 8px;
}

.text_32 {
    margin-left: 8px;
}

.text_33 {
    margin-left: 16px;
    line-height: 17px;
}

.image_14 {
    width: 1219px;
    height: 441px;
}

.section_11 {
    background-color: #1b2440;
}

.group_13 {
    padding: 65px 0 60px;
    width: 1143px;
}

.space-y-16>*:not(:first-child) {
    margin-top: 16px;
}

.space-x-194>*:not(:first-child) {
    margin-left: 194px;
}

.space-y-32>*:not(:first-child) {
    margin-top: 32px;
}

.image-wrapper {
    padding: 22px 0 14px;
    background-color: #ffffff;
    border-radius: 13px;
    width: 100px;
}

.image_15 {
    width: 91px;
    height: 64px;
}

.font_8 {
    font-size: 16px;
    font-family: PingFang SC Regular;
    line-height: 14.5px;
}

.text_40 {
    color: #ffffff;
    line-height: 15px;
}

.space-x-148>*:not(:first-child) {
    margin-left: 148px;
}

.font_6 {
    font-size: 18px;
    font-family: 苹方;
    line-height: 16.5px;
    color: #ffffff;
}

.font_7 {
    font-size: 14px;
    font-family: PingFang SC Regular;
    line-height: 13px;
    color: #bdc2d1;
}

.text_34 {
    margin-top: 28px;
}

.text_36 {
    margin-top: 22px;
}

.text_38 {
    margin-top: 22px;
}

.text_35 {
    margin-top: 28px;
}

.text_37 {
    margin-top: 22px;
}

.text_39 {
    margin-top: 22px;
}

.group_14 {
    margin-top: 4px;
}

.space-x-76>*:not(:first-child) {
    margin-left: 76px;
}

.space-y-24>*:not(:first-child) {
    margin-top: 24px;
}

.image_16 {
    border-radius: 4px;
    width: 99px;
    height: 100px;
}

.group_15 {
    height: 25px;
}

.text_41 {
    margin-top: 10px;
    color: #bdc2d1;
    line-height: 15.5px;
}

.group_16 {
    margin-top: -26px;
    padding: 0 12px;
}

.space-x-118-reverse>*:not(:last-child) {
    margin-right: 118px;
}

.section_12 {
    padding: 27px 0;
    background-color: #222a42;
}

.space-y-6>*:not(:first-child) {
    margin-top: 6px;
}

.group_17 {
    width: 389.5px;
    height: 13px;
}

.font_9 {
    font-size: 12px;
    font-family: PingFang SC Regular;
    line-height: 13px;
    color: #bdc2d1;
}

.text_42 {
    margin-top: -12px;
}

.text_43 {
    line-height: 14px;
    text-align: center;
}
</style>
