<template>
  <div class="ui container">
    <div class="top-header"><span>模型部署体验</span></div>
    <div class="content">
      <div class="content-left">
        <div class="header">
          <div class="tips1">请输入测试内容：</div>
          <div class="tips2" v-show="jobCategory == 1 || jobCategory == 2 || jobCategory == 3">* 请参考示例格式输入</div>
        </div>
        <div class="text-input-c">
          <el-input class="text-input" type="textarea" v-model="inputTxt" :rows="9"></el-input>
        </div>
        <div class="generate-btn-c">
          <el-button type="primary" size="default" @click="generate" :loading="generating">
            {{ generating ? '生成中' : '开始生成' }}
          </el-button>
        </div>
      </div>
      <div class="content-right">
        <div class="header">
          <div class="tips1">生成内容：</div>
        </div>
        <div class="text-output-c">
          <el-input class="text-output" type="textarea" v-model="outputTxt" readonly :rows="9"></el-input>
        </div>
      </div>
    </div>
    <div class="footer-tips-c">
      <div class="tips">* 部署的模型可以体验30分钟<span v-if="endTimeStr">，预计 <span style="color:red">{{ endTimeStr }}</span>
          后将关闭</span>。</div>
      <div class="tips">* 本项目处于前沿探索阶段，体验功能仅供学术测试使用。请勿输入违反法律内容，同时，未经许可，禁止分享，传播输入及生成文本内容。感谢理解！</div>
    </div>
  </div>
</template>

<script>
import dayjs from 'dayjs';
import { getFinetuneServiceStatus, getFinetuneServiceInference } from '~/apis/modules/modelbase';
import { getUrlSearchParams } from '~/utils';

// 文本分类
const sample1 = [{
  q: '对文本进行分类:\n有没有新疆旅游攻略？\n选项：游戏，财经，旅游，农业\n答案：',
  a: '旅游',
}];

// 中英翻译
const sample2 = [{
  q: '请将下面这段话改写成英文：\n本组织的目标:确保所有工作人员,包括特派团人员,以健康的身体履行职责,并促进和维持工作人员的健康\n答案：',
  a: 'Objective of the Organization: To ensure that all staff members, including those on mission, are fit to carry out their duties, and to promote and maintain the health of staff'
}];

// 情感分类
const sample3 = [{
  q: '情感分类:\n这个比精装的只少了一个大盒子和苏打志（写真书），价格只要一半，还是比较划算的，内容都是一样的精彩～\n选项：积极，消极\n答案：',
  a: '积极'
}];

// 开放问答示例
// const sample3 = [{
//   q: '问题：力的单位是什么？\n答案：',
//   a: '牛顿',
// }, {
//   q: '问题：2012年奥运在哪里举行？\n答案：',
//   a: '伦敦',
// }, {
//   q: '问题：郑州是那个省的\n答案：',
//   a: '河南',
// }, {
//   q: '问题：山东,山西,这里的山指\n答案：',
//   a: '太行山',
// }, {
//   q: '问题：北宋是中国历史上以汉族为主体建立的封建王朝，建都在哪里？\n答案：',
//   a: '开封',
// }, {
//   q: '问题：赌城在哪（美国）？\n答案：',
//   a: '拉斯维加斯',
// }];

const randomSample = (sampleList) => {
  return { ...sampleList[Math.floor(Math.random() * sampleList.length)] };
};

export default {
  data() {
    return {
      userName: '',
      repoName: 'openi-notebook',

      jobId: '',
      type: '',
      jobCategory: '',
      inputTxt: '',
      outputTxt: '',
      endTimeStr: '',

      generating: false,
      printTimer: null,
    };
  },
  components: {},
  methods: {
    generateSample() {
      if (this.jobCategory == 1) {
        const sample = randomSample(sample1);
        this.inputTxt = sample.q;
        // this.outputTxt = sample.a;
        this.printText(sample.a);
      } else if (this.jobCategory == 2) {
        const sample = randomSample(sample2);
        this.inputTxt = sample.q;
        // this.outputTxt = sample.a;
        this.printText(sample.a);
      } else if (this.jobCategory == 3) {
        const sample = randomSample(sample3);
        this.inputTxt = sample.q;
        // this.outputTxt = sample.a;
        this.printText(sample.a);
      } else if (this.jobCategory == 0) {

      }
    },
    generate() {
      const inputTxt = this.inputTxt.trim();
      if (!inputTxt) {
        this.$message({
          type: 'info',
          message: '请先输入测试内容',
        });
        return;
      }
      this.generating = true;
      getFinetuneServiceInference({
        userName: this.userName,
        jobId: this.jobId,
        text: inputTxt
      }).then(res => {
        this.generating = false;
        const data = res.data;
        if (data && data.text != undefined) {
          // this.outputTxt = data.text;
          this.printText(data.text);
        }
      }).catch(err => {
        this.generating = false;
        console.log(err);
        this.$message({
          type: 'error',
          message: err.message,
        });
      });
    },
    printText(text) {
      this.printTimer && clearInterval(this.printTimer);
      let len = text.length;
      this.printTimer = setInterval(() => {
        if (len < 0) clearInterval(this.printTimer);
        this.outputTxt = text.slice(0, text.length - (len - 1));
        len--;
      }, 18);
    }
  },
  beforeMount() {
    const metaEl = document.querySelectorAll('meta[name="_uid"]');
    if (!metaEl.length) { // 未登录
      window.location.href = `/user/login?redirect_to=${encodeURIComponent(window.location.href)}`;
      return;
    }
    const uid = metaEl[0].getAttribute('content');
    const uname = metaEl[0].getAttribute('content-ext');
    this.userName = uname;
    const urlParams = getUrlSearchParams();
    this.jobId = urlParams.jobid;
    this.type = urlParams.type;
    this.jobCategory = urlParams.jobcategory;
    // console.log(this.jobId, this.type, this.jobCategory);
    getFinetuneServiceStatus({
      userName: this.userName,
      jobId: this.jobId,
    }).then(res => {
      const data = res.data;
      if (data.code == 1) {
        this.$message({
          type: 'error',
          message: data.message,
        });
        setTimeout(() => {
          window.location.href = `/extension/modelbase/pangufinetune`;
        }, 3000);
        return;
      }
      if (data.fineTuneDeployStatus == 'SUCCEEDED') {
        if (data.fineTuneDeployFinishTime && data.fineTuneDeployFinishTime != 0) {
          this.endTimeStr = dayjs(data.fineTuneDeployFinishTime * 1000).format('YYYY-MM-DD HH:mm');
        }
        this.generateSample();
      } else {
        window.location.href = `/extension/modelbase/pangufinetune`;
      }
    }).catch(err => {
      console.log(err);
    });
  },
  mounted() { },
  beforeDestroy() { },
};
</script>

<style scoped lang="less">
.top-header {
  font-weight: 400;
  font-size: 28px;
  color: rgb(16, 16, 16);
  text-align: center;
  margin: 40px 0;
  height: 30px;
  line-height: 30px;
}

.content {
  display: flex;
  flex-wrap: wrap;

  .content-left {
    flex: 1;
    margin: 0px 18px;
    min-width: 400px;

    .header {
      display: flex;
      justify-content: space-between;
      margin-bottom: 8px;

      .tips1 {
        color: rgba(136, 136, 136, 1);
      }

      .tips2 {
        color: rgb(250, 140, 22);
      }
    }

    .text-input-c {
      .text-input {
        /deep/ .el-textarea__inner {
          font-size: 16px;
          color: rgb(16, 16, 16);
        }
      }
    }

    .generate-btn-c {
      text-align: right;
      margin-top: 20px;
    }
  }

  .content-right {
    flex: 1;
    margin: 0px 18px;
    min-width: 400px;

    .header {
      display: flex;
      margin-bottom: 8px;

      .tips1 {
        color: rgba(136, 136, 136, 1);
      }
    }

    .text-output-c {
      .text-output {
        /deep/ .el-textarea__inner {
          font-size: 16px;
          color: rgb(16, 16, 16);
          background-color: rgba(229, 231, 235, 0.1);

          &:focus {
            border-color: #DCDFE6;
          }
        }
      }
    }
  }
}

.footer-tips-c {
  margin-top: 40px;
  color: rgb(250, 140, 22);
  padding: 18px;

  .tips {
    margin-bottom: 5px;
  }
}
</style>
