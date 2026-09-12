<template>
  <div class="content-box">
    <Header> </Header>
    <div class="main-body">
      <div class="content">
        <div class="tools-list">
          <div class="tool-item bg-1">
            <div class="title">盘古GCU</div>
            <div class="descr">PengCheng-Mind 200B 脑海大模型是基于Transformer 架构的2010亿参数自回归式语言模型。</div>
            <div class="detail">
              <a target="_blank" href="https://openi.pcl.ac.cn/Enflame/GCU_Pytorch_pangu">
                <span>查看详情</span><i class="el-icon-right"></i>
              </a>
            </div>
          </div>
          <div class="tool-item bg-2">
            <div class="title">生态桥接工具 MindTorch</div>
            <div class="descr">支持与PyTorch一致的前端表达，兼容PyTorch生态，联合MindSpore框架优势，释放昇腾AI底层计算性能。</div>
            <div class="detail">
              <a target="_blank" href="https://openi.pcl.ac.cn/OpenI/mindtorch_tutorial">
                <span>查看详情</span><i class="el-icon-right"></i>
              </a>
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
  name: 'RelatedTools',
  data() {
    return {};
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
        justify-content: center;
        margin-top: 60px;
        flex-wrap: wrap;

        .tool-item {
          flex: 1;
          border-radius: 10px;
          max-width: 440px;
          min-width: 340px;
          box-shadow: rgba(157, 197, 226, 0.2) 0px 5px 10px 0px;
          padding: 20px;
          margin-bottom: 32px;
          margin-left: 20px;
          margin-right: 20px;

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
            padding: 0 30px;
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
        }
      }
    }
  }
}

@media screen and (min-width: 1300px) {
  .content-box .main-body .content .tools-list .tool-item {
    margin-left: 30px;
    margin-right: 30px;
  }
}
</style>
