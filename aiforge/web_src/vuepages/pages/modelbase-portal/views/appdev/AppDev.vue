<template>
  <div class="content-box">
    <Header></Header>
    <div class="main-body">
      <div class="content">
        <div class="tools-list">
          <div v-for="item in list" class="tool-item-c">
            <div class="tool-item" :class="item.bg || ''">
              <div class="title">{{ item.title }}</div>
              <div class="descr">{{ item.descr }}</div>
              <div class="detail">
                <a target="_blank" :href="item.link">
                  <span>{{ item.linkText }}</span><i class="el-icon-right"></i>
                </a>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import Header from '../../components/Header.vue';

export default {
  name: 'AppDev',
  data() {
    return {
      list: [{
        title: '智能服务编排',
        descr: '基于服务编排的交互式智能软件应用构造工具，在大模型辅助下，帮助开发人员根据需求快速编排服务并构建应用。',
        link: 'https://openi.pcl.ac.cn/login/oauth/authorize?client_id=234d0ff8-b82e-4265-8474-3f9e66f1f0ed&response_type=code&redirect_uri=http%3A%2F%2F8.136.53.186%3A4000%2Fflowentry',
        linkText: '查看详情',
        bg: 'bg-1',
      }, {
        title: '智能软件资源库',
        descr: '智能软件资源库：多形态智能软件资源平台，可持续爬取与解析各类开源资源库，大模型辅助智能软件资源语义检索与推荐，实现智能软件资源复用。',
        link: 'https://openi.pcl.ac.cn/login/oauth/authorize?client_id=f174218c-a415-40f2-ad4d-59be53844223&response_type=code&redirect_uri=http%3A%2F%2F123.57.177.171%3A9990%2Frecommend',
        linkText: '查看详情',
        bg: 'bg-2',
      }, /*{
        title: '引擎托管ASE平台',
        descr: '一个专为A.I.算法工程师打造的无服务、全托管式A.I.引擎平台，算法工程师通过集成ASE中提供的插件，即可快速实现 A.I.引擎云服务化。',
        link: 'https://openi.pcl.ac.cn/login/oauth/authorize?client_id=225ff096-e295-4f39-b63d-9f48206dbe85&redirect_uri=https://ase.openi.org.cn:18443/console/ase/userlogin&response_type=code&state=qizhi',
        linkText: '查看详情',
        bg: 'bg-3',
      }, {
        title: '通用AI AGENT平台',
        descr: '大模型AI应用开发组件，定位于帮助开发人员快速定制开发基于GPT大模型的智能应用（AI Agent/AI助手/智能体等）。',
        link: 'https://openi.pcl.ac.cn/login/oauth/authorize?client_id=1cac840c-0e80-4a7c-bdf9-0ec034b2101e&redirect_uri=http://203.83.235.100:10081/ai-agent/%23/gfLogin&response_type=code',
        linkText: '查看详情',
        bg: 'bg-4',
      }*/],
    };
  },
  components: { Header },
  methods: {
  },
  mounted() { },
  beforeDestroy() { },
};
</script>

<style scoped lang="less">
.content-box {
  height: 100%;
  width: 100%;
  display: flex;
  flex-direction: column;

  .main-body {
    flex: 1;
    height: 0;
    position: relative;
    margin-right: 16px;

    .content {
      height: calc(100% - 32px);
      margin-top: 16px;
      margin-bottom: 16px;
      overflow-y: auto;
      padding: 0 10px 10px 10px;
      scroll-behavior: smooth;

      .tools-list {
        display: flex;
        flex-wrap: wrap;
        margin-top: 30px;

        .tool-item-c {
          width: 50%;
          padding: 30px;
          padding-bottom: 20px;

          &:nth-child(odd) {
            display: flex;
            justify-content: flex-end;;
          }

          .tool-item {
            border-radius: 10px;
            box-shadow: rgba(157, 197, 226, 0.2) 0px 5px 10px 0px;
            padding: 20px;
            height: 100%;
            max-width: 440px;

            .title {
              color: rgba(255, 255, 255, 1);
              font-size: 18px;
              font-family: SourceHanSansSC;
              font-weight: 500;
              line-height: 25px;
              text-align: center;
              padding: 0 30px;
            }

            .descr {
              color: rgba(255, 255, 255, 1);
              font-size: 14px;
              font-family: SourceHanSansSC;
              font-weight: 300;
              line-height: 20px;
              text-align: center;
              padding: 0 15px;
              margin-top: 12px;
              opacity: 0.8;
            }

            .detail {
              display: flex;
              align-items: center;
              justify-content: center;
              margin-top: 14px;
              padding: 0 36px;

              a {
                color: rgba(255, 255, 255, 1);
                font-size: 14px;
                font-family: SourceHanSansSC;
                font-weight: 300;
                line-height: 20px;

                i {
                  margin-left: 4px;
                }

                &:hover {
                  opacity: 0.8;
                }
              }
            }

            &.bg-1 {
              background: url("data:image/svg+xml;charset=utf8,%3Csvg%20xmlns%3D%22http%3A%2F%2Fwww.w3.org%2F2000%2Fsvg%22%20xmlns%3Axlink%3D%22http%3A%2F%2Fwww.w3.org%2F1999%2Fxlink%22%20version%3D%221.1%22%3E%3Cdefs%3E%3ClinearGradient%20id%3D%221%22%20x1%3D%220%22%20x2%3D%221%22%20y1%3D%220%22%20y2%3D%220%22%20gradientTransform%3D%22matrix(-0.9550000000000001%2C%200.6080000000000002%2C%20-0.05927810079133697%2C%20-0.9550000000000001%2C%200.949%2C%200)%22%3E%3Cstop%20stop-color%3D%22%23009dff%22%20stop-opacity%3D%221%22%20offset%3D%220.02%22%3E%3C%2Fstop%3E%3Cstop%20stop-color%3D%22%23007fff%22%20stop-opacity%3D%221%22%20offset%3D%221%22%3E%3C%2Fstop%3E%3C%2FlinearGradient%3E%3C%2Fdefs%3E%3Crect%20width%3D%22100%25%22%20height%3D%22100%25%22%20fill%3D%22url(%231)%22%3E%3C%2Frect%3E%3C%2Fsvg%3E");
            }

            &.bg-2 {
              background: url("data:image/svg+xml;charset=utf8,%3Csvg%20xmlns%3D%22http%3A%2F%2Fwww.w3.org%2F2000%2Fsvg%22%20xmlns%3Axlink%3D%22http%3A%2F%2Fwww.w3.org%2F1999%2Fxlink%22%20version%3D%221.1%22%3E%3Cdefs%3E%3ClinearGradient%20id%3D%221%22%20x1%3D%220%22%20x2%3D%221%22%20y1%3D%220%22%20y2%3D%220%22%20gradientTransform%3D%22matrix(-0.9550000000000001%2C%200.6080000000000002%2C%20-0.05927810079133697%2C%20-0.9550000000000001%2C%200.949%2C%200)%22%3E%3Cstop%20stop-color%3D%22%23687dd0%22%20stop-opacity%3D%221%22%20offset%3D%220%22%3E%3C%2Fstop%3E%3Cstop%20stop-color%3D%22%232758c3%22%20stop-opacity%3D%221%22%20offset%3D%221%22%3E%3C%2Fstop%3E%3C%2FlinearGradient%3E%3C%2Fdefs%3E%3Crect%20width%3D%22100%25%22%20height%3D%22100%25%22%20fill%3D%22url(%231)%22%3E%3C%2Frect%3E%3C%2Fsvg%3E");
            }

            &.bg-3 {
              background: url("data:image/svg+xml;charset=utf8,%3Csvg%20xmlns%3D%22http%3A%2F%2Fwww.w3.org%2F2000%2Fsvg%22%20xmlns%3Axlink%3D%22http%3A%2F%2Fwww.w3.org%2F1999%2Fxlink%22%20version%3D%221.1%22%3E%3Cdefs%3E%3ClinearGradient%20id%3D%221%22%20x1%3D%220%22%20x2%3D%221%22%20y1%3D%220%22%20y2%3D%220%22%20gradientTransform%3D%22matrix(-0.9550000000000001%2C%200.6080000000000002%2C%20-0.05927810079133697%2C%20-0.9550000000000001%2C%200.949%2C%200)%22%3E%3Cstop%20stop-color%3D%22%234e96e6%22%20stop-opacity%3D%221%22%20offset%3D%220.02%22%3E%3C%2Fstop%3E%3Cstop%20stop-color%3D%22%236ab88f%22%20stop-opacity%3D%221%22%20offset%3D%221%22%3E%3C%2Fstop%3E%3C%2FlinearGradient%3E%3C%2Fdefs%3E%3Crect%20width%3D%22100%25%22%20height%3D%22100%25%22%20fill%3D%22url(%231)%22%3E%3C%2Frect%3E%3C%2Fsvg%3E");
            }

            &.bg-4 {
              background: url("data:image/svg+xml;charset=utf8,%3Csvg%20xmlns%3D%22http%3A%2F%2Fwww.w3.org%2F2000%2Fsvg%22%20xmlns%3Axlink%3D%22http%3A%2F%2Fwww.w3.org%2F1999%2Fxlink%22%20version%3D%221.1%22%3E%3Cdefs%3E%3ClinearGradient%20id%3D%221%22%20x1%3D%220%22%20x2%3D%221%22%20y1%3D%220%22%20y2%3D%220%22%20gradientTransform%3D%22matrix(-0.9550000000000001%2C%200.6080000000000002%2C%20-0.05927810079133697%2C%20-0.9550000000000001%2C%200.949%2C%200)%22%3E%3Cstop%20stop-color%3D%22%23cbad0c%22%20stop-opacity%3D%221%22%20offset%3D%220%22%3E%3C%2Fstop%3E%3Cstop%20stop-color%3D%22%236161c9%22%20stop-opacity%3D%221%22%20offset%3D%221%22%3E%3C%2Fstop%3E%3C%2FlinearGradient%3E%3C%2Fdefs%3E%3Crect%20width%3D%22100%25%22%20height%3D%22100%25%22%20fill%3D%22url(%231)%22%3E%3C%2Frect%3E%3C%2Fsvg%3E");
            }
          }
        }
      }
    }
  }
}

@media screen and (min-width: 900px) and (max-width: 1600px) {
  .content-box .main-body .content .tools-list .tool-item-c {
    width: 50%;
  }
}

@media screen and (min-width: 700px) and (max-width: 900px) {
  .content-box .main-body .content .tools-list .tool-item-c {
    width: 100%;
    padding-left: 30px;
    padding-right: 30px;
    padding-top: 15px;
    justify-content: center;
    display: flex;

    &:nth-child(odd) {
      justify-content: center;
    }
  }
}

@media screen and (max-width: 700px) {
  .content-box .main-body .content .tools-list .tool-item-c {
    width: 100%;
    padding-top: 15px;
    justify-content: center;
    display: flex;

    &:nth-child(odd) {
      justify-content: center;
    }
  }
}
</style>
