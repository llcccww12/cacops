<template>
    <div class="flex-col page">
        <AppBanner></AppBanner>





        <div class="flex-col justify-start relative group_3">
            <div class="flex-col relative section_8">
                <span class="self-start font_3">ROS-hmci资源发布</span>
                <div class="flex-col group_5 space-y-18">
                    <div v-for="(resource, index) in paginatedResources" :key="index" class="flex-col resource">
                        <div class="flex-row items-center self-start space-x-20">
                            <img class="shrink-0 image" src="/img/ros-hmci/9b12ac07bff17055334642687e7ea577.png" />
                            <router-link :to="{ name: 'ResourceDetail', params: { name: resource.name } }">
                                <span class="name">{{ resource.name }}</span>
                            </router-link>
                        </div>
                        <span class="self-start font_4 text_19">
                            {{ resource.synopsis }}
                        </span>
                        <div class="flex-col self-start group_6">
                            <img class="shrink-0 self-start image_14" src="/img/ros-hmci/mbz627.png" />
                            <span class="self-center font_2 text_20">提交日期：{{ resource.create_time }}</span>
                        </div>
                        <div class="divider"></div>
                    </div>
                </div>


                <div class="center pagination">
                    <el-pagination @size-change="handleSizeChange" @current-change="handleCurrentChange"
                        :current-page="currentPage" :page-sizes="[15]" :page-size="15"
                        layout="total, sizes, prev, pager, next, jumper" :total="totalItems">
                    </el-pagination>
                </div>


            </div>



            <div class="section_7 pos_3">
                <div class="flex-row justify-between items-center section_7_content">
                    <div></div>
                    <div class="flex-col justify-start items-start relative group_4">
                        <input class="font_1 text_15" outline:none type="text" placeholder="请输入资源名称以搜索" v-model="searchQuery" />
                        <!-- <img class="image_11 pos_4" src="/img/ros-hmci/bcea28a03624dcba382543662dbb0a2c.png" /> -->
                        <button class="ui green button image_11 pos_4">{{ $t('repos.search') }}</button>
                    </div>
                    <div class="flex-row space-x-28">
                        <!-- <div class="flex-row items-center space-x-6">
                            <span class="font_2 text_17">排序</span>
                            <img class="shrink-0 image_13" src="/img/ros-hmci/fa6fa7aef56d86f283de58808b2a8b7d.png" />
                        </div> -->
                        <div id="openForm" class="flex-row space-x-8">
                            <img class="shrink-0 image_12" src="/img/ros-hmci/4b4c4662629e074227c63d23f46271ac.png" />
                            <a :href="links.source" class="font_1 text_16">资源发布</a>
                        </div>
                    </div>
                </div>
            </div>
        </div>
    </div>
</template>
  
<script>
import AppBanner from '../components/AppBanner.vue'
import axios from 'axios'
import linksConfig from './links.js'

export default {
    name: 'CommunitySource',
    data() {
        return {
            searchQuery: '',
            links: linksConfig,
            resources: [], // 存储所有资源
            currentPage: 1, // 当前页码
            pageSize: 15, // 每页显示的资源数量
            totalItems: 0, // 资源总数
        };
    },
    async created() {
        // 获取数据
        const response = await fetch(this.links.source_Api);
        const data = await response.json();

        // 将数据设置为resources
        this.resources = data;
        this.totalItems = this.resources.length;
    },

    components: {
        AppBanner,
    },

    methods: {
        handleSizeChange(newSize) {
            this.pageSize = newSize;
        },

        handleCurrentChange(newPage) {
            this.currentPage = newPage;
        },
    },
    mounted() { },

    computed: {
        filteredResources() {
            if (!this.searchQuery) {
                return this.resources;
            }

            const query = this.searchQuery.toLowerCase();
            return this.resources.filter(
                (resource) =>
                    resource.name.toLowerCase().includes(query) ||
                    resource.synopsis.toLowerCase().includes(query),
            );
        },
        paginatedResources() {
            const start = (this.currentPage - 1) * this.pageSize;
            const end = this.currentPage * this.pageSize;
            return this.filteredResources.slice(start, end);
        },
    },
};

</script>
  
  <!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
.pagination {
    margin-top: 30px;
    background-color: #ffffff;
}

.page {
    margin-bottom: -80px;
    /* background-color: #0b0e1d; */
    width: 100%;
    overflow-y: auto;
    overflow-x: hidden;
    height: 100%;
}

.overlay {
    z-index: 1000;
    width: 100%;
    height: 100%;
    position: fixed;
    background-color: rgba(0, 0, 0, 0.27);
}

.resourcePublish {
    z-index: 1;
    border-radius: 12px;
    background-color: #ffffff;
    position: fixed;
    top: 9%;
    left: 50%;
    transform: translateX(-50%);
    height: 806px;
    width: 800px;

}

.hidden {
    display: none;
}

.rp1 {
    width: 800px;
    height: 133px;
    background: no-repeat;
    background-image: url('/img/ros-hmci/组\ 7314.png');
    background-size: 100% 100%;
}

.font_41 {
    margin-left: 15px;
    font-family: PingFang SC;
    color: #181818;
    font-size: 18px;
    height: 25px;
}

.image_36 {
    height: 61.5px;
    margin-left: 54px;
}

.section {
    padding: 17px 80px 12px 123px;
    filter: drop-shadow(0px 2px 3px #0137931f);
    background-color: #282828;
}

.group_25 {
    position: relative;
    top: 0;
    background: url(/img/ros-hmci/mbz605.png);
    background-position: center;
    background-size: cover;
    background-repeat: no-repeat;
    height: 600px;
    width: 100%;
}



.space-x-40>*:not(:first-child) {
    margin-left: 40px;
}

.image_4 {
    width: 63px;
    height: 16px;
}

.text_72 {
    font-family: PingFang SC;
    color: #ffffff;
    font-size: 18px;
    text-align: center;
}

.image_411 {
    width: 7px;
    height: 7px;

}

.group_91 {
    padding: 15px 54px 0px;
}

.group_101 {
    margin-top: 10px;
    border: 1px solid;
    border-color: #caced8;
    width: 692px;
    height: 48px;
    border-radius: 12px;
}

.group_102 {
    margin-top: 10px;
    border: 1px solid;
    border-color: #caced8;
    width: 692px;
    height: 99px;
    border-radius: 12px;
}

.text_62 {
    font-family: PingFang SC;
    color: #24a19b;
    font-size: 16px;
    line-height: 33px;
}

.text_63 {
    font-family: PingFang SC;
    color: #ffffff;
    font-size: 16px;
    line-height: 33px;
}

.text-wrapper_41 {
    margin-right: 15px;
    width: 89px;
    height: 36px;
    border: 1px solid;
    border-color: #24a19b;
    border-radius: 6px;
}

.pos_commit {
    position: absolute;
    bottom: 25px;
    right: 56px;
}

.text-wrapper_51 {
    width: 89px;
    height: 36px;
    background-color: #24a19b;
    border-radius: 6px;
}

.divider_2 {
    margin-top: 38px;
    margin-left: 54px;
    width: 695px;
    height: 0px;
    border: 1px dashed #fff;
    border-color: rgba(157, 163, 175, 0.55);
}

.pos_34 {
    position: absolute;
    top: 14px;
    left: 0;
}

.text_33 {
    margin-left: 20px;
    font-family: PingFang SC;
    color: #181818;
    font-size: 15px;
    /* line-height: 33px; */
}

.text_34 {
    margin-left: 20px;
    font-family: PingFang SC;
    color: #757c8d;
    font-size: 15px;
    /* line-height: 33px; */
}

.font_51 {
    font-family: PingFang SC;
    color: #575757;
    font-size: 16px;
    /* line-height: 33px; */
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
    bottom: 0px;
    background-color: rgba(255, 255, 255, 0.15);
    height: 108px;
    width: 100%;
    backdrop-filter: blur(20px);
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
    font-family: PingFang SC;
    line-height: 13px;
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

.image_7 {
    margin-left: 38px;
    width: 61px;
    height: 26px;
}

.text_9 {
    margin-left: 40px;
    font-size: 15px;
    line-height: 14px;
}

.text_10 {
    margin-left: 38px;
    font-size: 15px;
    line-height: 14px;
}

.text_11 {
    margin-left: 38px;
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
    line-height: 13.5px;
    opacity: 0.63;
}

.pos_2 {
    position: absolute;
    right: 23.5px;
    top: 50%;
    transform: translateY(-50%);
}

.image_8 {
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

.image_5 {
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

.image_3 {
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
    width: 30px;
    height: 30px;
}

.image_2 {
    margin-left: 20px;
}

.text_13 {
    margin-left: 10px;
}

.image_9 {
    margin-left: 8px;
    border-radius: 1px;
    width: 11px;
    height: 6px;
}

.image_6 {
    margin-left: 20px;
    opacity: 0.7;
    width: 20px;
    height: 19px;
}

.section_5 {
    padding: 2px 0 243px;
    background-image: url('/img/ros-hmci/1483edffca1c70a08974d6b7de6a3fdb.png');
    background-size: 100% 100%;
    background-repeat: no-repeat;
}

.image_37 {
    height: 17px;
    position: absolute;
    right: 22px;
    top: 43px;
}

.section_6 {
    padding: 10px 360px;
    background-color: #ffffff26;
}

.space-x-14>*:not(:first-child) {
    margin-left: 14px;
}

.image_10 {
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
    color: #ffffff;
}

.group_3 {
    background-color: #ffffff;
    padding-top: 58px;
}

.section_8 {
    width: 1100px;
    left: 50%;
    transform: translateX(-50%);
    padding: 52px 0;
    background-color: #ffffff;
    border-radius: 4px;
}

.font_3 {
    font-size: 16px;
    font-family: AlibabaPuHuiTi;
    line-height: 15px;
    color: #323232;
}

.group_5 {
    margin-top: 24px;
}

.space-y-18>*:not(:first-child) {
    margin-top: 18px;
}

.space-x-20>*:not(:first-child) {
    margin-left: 20px;
}

.text_18 {
    font-size: 16px;
    font-family: AlibabaPuHuiTi;
    line-height: 15px;
    color: #323232;
    color: #24a19b;
}

.font_4 {
    font-size: 14px;
    font-family: AlibabaPuHuiTi;
    line-height: 13px;
    color: #4d4d4d;
}

.text_19 {
    margin-left: 50px;
    margin-top: 7px;
}

.group_6 {
    margin-left: 50px;
    margin-top: 20px;
    width: 159.5px;
    height: 12px;
}

.image_14 {
    width: 15px;
    height: 15px;
}

.text_20 {
    margin-left: 20px;
    margin-top: -12px;
    font-size: 13px;
    line-height: 12px;
}

.divider {
    margin-left: 50px;
    margin-top: 16px;
    background-color: #caced8;
    height: 0px;
}

.text_21 {
    margin-left: 50px;
    margin-top: 6px;
}

.group_7 {
    width: 960.5px;
}

.space-y-4>*:not(:first-child) {
    margin-top: 4px;
}

.text_22 {
    margin-left: 50px;
}

.view {
    margin-top: 0;
}

.view_2 {
    margin-top: 0;
}

.text_23 {
    margin-left: 52px;
    margin-top: 4px;
}

.text_24 {
    margin-left: 50px;
    margin-top: 7px;
}

.group_8 {
    margin-top: 30px;
    width: 511px;
}

.image_15 {
    border-radius: 4px;
    width: 32px;
    height: 32px;
}

.text-wrapper_3 {
    padding: 10px 0;
    background-color: #ffffff;
    border-radius: 4px;
    width: 32px;
    height: 32px;
    border: solid 1px #caced8;
}

.font_5 {
    font-size: 14px;
    font-family: HelveticaNeue;
    line-height: 10px;
    color: #4d4d4d;
}

.text-wrapper_4 {
    padding: 10px 0;
    background-color: #24a19b;
    border-radius: 4px;
    width: 32px;
    height: 32px;
}

.text_25 {
    color: #ffffff;
}

.text-wrapper_5 {
    padding: 6px 0 12px;
    background-color: #ffffff;
    border-radius: 4px;
    width: 32px;
    height: 32px;
    border: solid 1px #caced8;
}

.text-wrapper_6 {
    padding: 8px 0 14px;
    background-color: #ffffff;
    border-radius: 4px;
    width: 32px;
    height: 32px;
    border: solid 1px #caced8;
}

.text-wrapper_7 {
    padding: 8px 0 12px;
    background-color: #ffffff;
    border-radius: 4px;
    width: 32px;
    height: 32px;
    border: solid 1px #caced8;
}

.text-wrapper_8 {
    padding: 8px 0 12px;
    border-radius: 4px;
    width: 32px;
    height: 32px;
    border: solid 1px #caced8;
}

.text-wrapper_9 {
    padding: 8px 0 10px;
    background-color: #ffffff;
    border-radius: 4px;
    width: 82px;
    height: 32px;
    border: solid 1px #caced8;
}

.font_6 {
    font-size: 14px;
    font-family: PingFang SC;
    line-height: 13px;
    color: #4d4d4d;
}

.text-wrapper_10 {
    padding: 10px 0;
    background-color: #ffffff;
    border-radius: 4px;
    width: 48px;
    height: 32px;
    border: solid 1px #caced8;
}

.text_26 {
    color: #000000a6;
}

.text_27 {
    line-height: 12px;
}

.section_7 {
    text-align: center;
    padding: 24px 0;
    background-color: #f6f9fb;
}
.section_7_content{
    width: 1200px;
}

.pos_3 {
    position: absolute;
    left: 0;
    right: 0;
    top: 0;
}

.group_4 {
    width: 676px;
}

.text-wrapper_2 {
    padding: 10px 0 12px;
    border-radius: 4px;
    width: 442px;
    border: solid 1px #caced8;
}

.text_15 {
    border: 1px solid #caced8;
    padding: 7px 15px;
    border-radius: 4px 0 0 4px;
    width: 602px;
    background-color: #f6f9fb;
    color: #757c8d;
}

.image_11 {
    border-radius: 0 4px 4px 0;
    width: 70px;
    height: 36px;
}

.pos_4 {
    position: absolute;
    right: 0;
}

.space-x-28>*:not(:first-child) {
    margin-left: 28px;
}

.space-x-8>*:not(:first-child) {
    margin-left: 8px;
}

.image_12 {
    width: 16px;
    height: 15px;
    align-self: center;
}

.text_16 {
    color: #24a19b;
    font-size: 15px;
    line-height: 21px;
}

.space-x-6>*:not(:first-child) {
    margin-left: 6px;
}

.text_17 {
    font-size: 15px;
    line-height: 21px;
}

.image_13 {
    border-radius: 1px;
    width: 12px;
    height: 6px;
}

.section_9 {
    background-color: #1b2440;
}

.group_9 {
    padding: 64px 0;
    width: 1143px;
}

.space-y-16>*:not(:first-child) {
    margin-top: 16px;
}

.space-x-195>*:not(:first-child) {
    margin-left: 195px;
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

.image_16 {
    width: 91px;
    height: 64px;
}

.font_9 {
    font-size: 16px;
    font-family: PingFang SC;
}


.space-x-148>*:not(:first-child) {
    margin-left: 148px;
}

.font_7 {
    font-size: 18px;
    font-family: 苹方;
    line-height: 16.5px;
    color: #ffffff;
}

.font_8 {
    font-size: 14px;
    font-family: PingFang SC;
    line-height: 13px;
    color: #bdc2d1;
}

.text_28 {
    margin-top: 28px;
}

.text_30 {
    margin-top: 22px;
}

.text_32 {
    margin-top: 22px;
}

.text_29 {
    margin-top: 28px;
}

.text_31 {
    margin-top: 22px;
}



.group_10 {
    margin: 4px 0 2px;
}

.space-x-76>*:not(:first-child) {
    margin-left: 76px;
}

.space-y-24>*:not(:first-child) {
    margin-top: 24px;
}

.image_17 {
    border-radius: 4px;
    width: 99px;
    height: 100px;
}

.group_11 {
    height: 20.5px;
}

.text_36 {
    margin-top: 6px;
    color: #bdc2d1;
    line-height: 14.5px;
}

.group_12 {
    margin-top: -20px;
    padding: 0 12px;
}

.space-x-118-reverse>*:not(:last-child) {
    margin-right: 118px;
}

.text_35 {
    line-height: 12.5px;
}

.section_10 {
    padding: 28px 0;
    background-color: #222a42;
}

.space-y-6>*:not(:first-child) {
    margin-top: 6px;
}

.group_13 {
    width: 395px;
    height: 12.5px;
}

.font_10 {
    font-size: 12px;
    font-family: PingFang SC;
    line-height: 13px;
    color: #bdc2d1;
}

.text_37 {
    line-height: 12.5px;
}

.text_38 {
    margin-top: -12px;
    line-height: 12.5px;
}

.text_39 {
    line-height: 14px;
    text-align: center;
    width: 381px;
}
</style>