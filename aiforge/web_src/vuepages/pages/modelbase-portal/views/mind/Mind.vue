<template>
  <div>
    <div class="header">
      <div class="title-0">“中国算力网(China Computing NET, C²NET)”专项重大开源应用示范</div>
      <div class="title">鹏城 · 脑海</div>
      <div class="descr">两千亿参数规模大语言模型，期待与您共同书写脑海大模型的新篇章！</div>
    </div>
    <div class="top">
      <div class="ui container content">
        <div class="content-title">鹏城 · 脑海大模型开源计划</div>
        <div class="content-c">
          <div class="content-l">
            <div class="">
              <p>信息爆炸时代，开源大模型的力量无以估量。</p>
              <p>我们很荣幸面向全社会开源两千亿参数规模的大语言模型基座—“鹏城·脑海”，希望它开启一个充满无限可能的创新之门。</p>
              <p>
                即使现在，它不尽完美，但我们坚信开源的力量会让脑海大模型得到真正的成长与繁荣。为此，我们将向开源社区开放“鹏城·脑海”多参数版本及其成长过程中的各阶段版本，我们诚邀各方贤达基于脑海大模型开展进一步研究，探索脑海无限可能，共促人工智能发展。
              </p>
              <p>期待与您共同书写脑海大模型的新篇章！</p>
              <p style="text-align:right;padding-right:10px;">——鹏城实验室脑海大模型研究组</p>
            </div>
            <div class="apply-c">
              <a class="apply-btn" download
                href="https://openi.pcl.ac.cn/OpenIOSSG/promote/raw/branch/master/%e9%b9%8f%e5%9f%8e%e8%84%91%e6%b5%b7%e6%a8%a1%e5%9e%8b%e7%94%b3%e8%af%b7%e8%a1%a8.docx">模型申请</a>
              <a class="apply-btn research-plan-btn" href="#research-plan">相关研究问题开源计划</a>
            </div>
            <div class="apply-descr">下载表格，填写后发送到：secretariat@openi.org.cn</div>
          </div>
          <div class="content-r"></div>
        </div>
      </div>
    </div>
    <div class="ui container">
      <div class="models">
        <div class="model-card-c">
          <div class="model-card" v-for="(item, index) in models" :key="index">
            <div class="model-name-1">{{ item.name1 }}</div>
            <div class="model-name-2">{{ item.name2 }}</div>
            <div class="model-status">
              <div class="status" v-if="item.status.indexOf(1) > -1">
                <svg xmlns="http://www.w3.org/2000/svg"
                  class="styles__StyledSVGIconPathComponent-sc-gbp7ch-0 iVwQlg svg-icon-path-icon fill"
                  viewBox="0 0 32 32" width="16" height="16" fill="rgb(22, 132, 252)">
                  <defs data-reactroot=""></defs>
                  <g>
                    <path
                      d="M16 29.333c-7.364 0-13.333-5.969-13.333-13.333s5.969-13.333 13.333-13.333 13.333 5.969 13.333 13.333-5.969 13.333-13.333 13.333zM16 26.667c5.891 0 10.667-4.776 10.667-10.667s-4.776-10.667-10.667-10.667v0c-5.891 0-10.667 4.776-10.667 10.667s4.776 10.667 10.667 10.667v0zM14.671 21.333l-5.657-5.657 1.885-1.885 3.772 3.772 7.541-7.543 1.887 1.885-9.428 9.428z">
                    </path>
                  </g>
                </svg>
                <span>可申请</span>
              </div>
              <div class="status" v-if="item.status.indexOf(2) > -1">
                <svg xmlns="http://www.w3.org/2000/svg"
                  class="styles__StyledSVGIconPathComponent-sc-gbp7ch-0 iVwQlg svg-icon-path-icon fill"
                  viewBox="0 0 32 32" width="16" height="16" fill="rgb(22, 132, 252)">
                  <defs data-reactroot=""></defs>
                  <g>
                    <path
                      d="M16 2.667c0.004 0 0.009 0 0.014 0 4.147 0 7.852 1.897 10.292 4.871l0.019 0.023-8.439 8.439 8.439 8.439c-2.459 2.997-6.163 4.895-10.31 4.895-0.005 0-0.009 0-0.014 0h0.001c-7.364 0-13.333-5.969-13.333-13.333s5.969-13.333 13.333-13.333zM16 5.333c-5.89 0.001-10.665 4.776-10.665 10.667s4.776 10.667 10.667 10.667c2.353 0 4.527-0.762 6.291-2.052l-0.030 0.021 0.277-0.209-8.427-8.427 8.427-8.428-0.277-0.208c-1.627-1.194-3.648-1.94-5.839-2.022l-0.019-0.001-0.404-0.008zM16 6.667c1.105 0 2 0.895 2 2s-0.895 2-2 2v0c-1.105 0-2-0.895-2-2s0.895-2 2-2v0z">
                    </path>
                  </g>
                </svg>
                <span>持续训练中<span v-if="item.updateTime">（{{ item.updateTime }}）</span></span>
              </div>
            </div>
            <div class="model-descr-c">
              <template v-for="(descr, descrIndex) in item.descr">
                <div class="descr" :key="descrIndex">
                  <div v-if="descr[0]" class="descr-l1">{{ descr[0] }}</div>
                  <div v-if="descr[1]" class="descr-l2">{{ descr[1] }}</div>
                  <div v-if="descr[2]" class="descr-l3">{{ descr[2] }}</div>
                </div>
                <div class="line" v-if="descrIndex != item.descr.length - 1" :key="descrIndex + 100"></div>
              </template>
            </div>
            <div class="link">
              <a :href="item.link">
                <span>查看项目</span>
                <svg xmlns="http://www.w3.org/2000/svg"
                  class="styles__StyledSVGIconPathComponent-sc-gbp7ch-0 bRiIWI svg-icon-path-icon fill"
                  viewBox="0 0 32 32" width="16" height="16" fill="rgb(16, 16, 16)">
                  <defs data-reactroot=""></defs>
                  <g>
                    <path
                      d="M21.563 14.667l-7.152-7.152 1.885-1.885 10.371 10.371-10.371 10.371-1.885-1.885 7.152-7.152h-16.229v-2.667z">
                    </path>
                  </g>
                </svg>
              </a>
            </div>
            <div class="ops">
              <a :href="op.link" v-for="(op, opIndex) in item.ops" :key="opIndex">{{ op.text }}</a>
            </div>
          </div>
        </div>
      </div>
    </div>
    <div class="ui container">
      <div class="model-tab">
        <div class="tab-item" v-for="(item, index) in tabList" :key="index" :class="activeTab == index ? 'active' : ''"
          @click="changeTab(index)">
          {{ item }}
        </div>
      </div>
      <div class="model-tab-content">
        <div class="descr">
          <p v-for="(descr, descrIndex) in tabContent[activeTab].descr" :key="descrIndex">{{ descr }}</p>
        </div>
        <div class="life-stage" v-if="tabContent[activeTab].stage">
          <div class="title">成长生命期阶段版本<span v-if="activeTab == 0">（当前最新版本——进阶版1.5T Tokens）</span></div>
          <div class="descr">您提交的申请体验通过审核后，我们将提供中间 checkpoints 供您研究使用；随着训练的推进，我们将提供对应阶段的 checkpoints 。</div>
          <div class="stage-c">
            <div class="stage" v-for="(item, index) in tabContent[activeTab].stage" :key="index">
              <div class="stage-item" v-if="item.name">
                <div class="name">{{ item.name }}</div>
                <div class="tokens">{{ item.tokens }}</div>
              </div>
              <div v-else class="stage-item">······</div>
              <div class="stage-arr">
                <svg xmlns="http://www.w3.org/2000/svg" style="overflow: visible; filter: none;" view-box="0 0 32 32"
                  width="32" height="9">
                  <path d="M 0 0.5 L 32 0.5" stroke="rgba(22,132,252,1)" stroke-width="1px"></path>
                </svg>
                <svg xmlns="http://www.w3.org/2000/svg" style="margin-left:-13px;" fill="rgba(22,132,252,1)"
                  viewBox="0 0 32 32" width="17" height="17">
                  <defs data-reactroot=""></defs>
                  <g>
                    <path d="M17.563 16l-6.6-6.6 1.885-1.885 8.485 8.485-8.485 8.485-1.885-1.885z"></path>
                  </g>
                </svg>
              </div>
            </div>
            <div class="stage">
              <div class="stage-item">······</div>
            </div>
          </div>
        </div>
        <div class="data-distribution" v-if="tabContent[activeTab].chart && tabContent[activeTab].chart.length">
          <div v-if="activeTab == 0">
            <div class="title">训练数据分布</div>
            <div class="descr">
              <p>我们统计了40个学科领域的数据，此处只罗列了占比前5的学科领域的数据分布情况。</p>
              <p>
                40个学科领域包括：文学、经济、法律、教育学、医学、政治、历史、体育、游戏娱乐、影视影评、心理学、休闲旅游、计算机科学、艺术、食品科学、农业科学、电子商务、电子信息科学、管理学、交通运输、能源科学、军事、轻工业、环境科学、安全科学与消防工程、哲学、生物科学、城市规划与建设、地理、材料科学、广告营销、宗教、天文学、气象学、化学科学、物理科学、数学科学、重工业、通信技术、其他。
              </p>
            </div>
            <div class="chart-el-c">
              <div class="chart-el" :class="`chart-el-${index}`" v-for="(item, index) in tabContent[activeTab].chart"
                :key="index"></div>
            </div>
          </div>
          <div v-if="activeTab == 1">
            <div class="title">训练数据分布</div>
            <div class="descr">
              <p>我们构建了人机协作模式的预训练数据加工线，共计生产了1TB Tokens的高质量预训练数据。预训练数据的中文、英文、代码、多语言数据比例约为 30%:60%:5%:5%，截止目前，新版PengCheng
                Mind-7B模型已完成440B Tokens数据的持续学习训练。</p>
              <p>
                针对中文训练数据，我们完成了40个学科领域的数据划分和分布统计，具体学科领域包括：文学、经济、法律、教育学、医学、政治、历史、体育、游戏娱乐、影视影评、心理学、休闲旅游、计算机科学、艺术、食品科学、农业科学、电子商务、电子信息科学、管理学、交通运输、能源科学、军事、轻工业、环境科学、安全科学与消防工程、哲学、生物科学、城市规划与建设、地理、材料科学、广告营销、宗教、天文学、气象学、化学科学、物理科学、数学科学、重工业、通信技术、其他。
              </p>
            </div>
            <div class="chart-el-c">
              <div class="chart-el chart-el-full" :class="`chart-el-0`"></div>
            </div>
            <div class="descr">
              <p>
                针对英文训练数据，我们完成了38个学科领域的数据划分和分布统计，具体学科领域包括：Entertainment and tourism、Medicine and health、Business、Language
                and literature、History、Arts、Law、Computer Science、Energy science、Mathematics、Political Science、Management
                Science、Sociology、Physics、Religious Studies、Education、Military science、Human physical performance and
                recreation、Environment studies and forestry、Agriculture、Archaeology、Information and communication
                sciences、Engineering and
                technology、Philosophy、Astronomy、Geography、Biology、Transportation、Economics、Materials
                Science、Chemistry、Psychology、Food science、Communication technology、Earth Science、Aerography、Architecture
                and design、Others。
              </p>
            </div>
            <div class="chart-el-c">
              <div class="chart-el chart-el-full" :class="`chart-el-1`"></div>
            </div>
          </div>
        </div>
        <div class="performance-evaluation" v-if="tabContent[activeTab].evaluation">
          <div class="title">性能评测</div>
          <div class="descr">
            <p>{{ tabContent[activeTab].evaluation.descr }}</p>
          </div>
          <div class="chart-el-c">
            <div class="chart-el  chart-el-full" :class="`chart-el-eval`" style="height:450px;"></div>
          </div>
        </div>
        <div class="features" v-if="tabContent[activeTab].features">
          <div class="sum">{{ tabContent[activeTab].features.sum }}</div>
          <div class="feature" v-for="(feature, index) in tabContent[activeTab].features.list" :key="index">
            <div class="title">{{ feature.title }}</div>
            <div class="descr">{{ feature.descr }}</div>
          </div>
        </div>
      </div>
    </div>
    <div class="research-plan-field" v-if="researchPlan">
      <a name="research-plan"></a>
      <div class="ui container">
        <div class="research-plan">
          <div class="title">{{ researchPlan.title }}</div>
          <div class="descr-c">
            <div class="descr">{{ researchPlan.descr }}</div>
          </div>
          <div class="main-pic">
            <img :src="researchPlan.imgSrc" alt="" />
          </div>
          <div class="cards-c">
            <div class="cards">
              <div class="card" v-for="(technology, technologyIndex) in researchPlan.technologys"
                :key="technologyIndex">
                <div class="card-title">
                  <div class="card-num">{{ technologyIndex + 1 }}</div>
                  <div class="card-name-c">
                    <span class="card-name" :title="technology.name">{{ technology.name }}</span>
                  </div>
                </div>
                <div v-for="(reserch, index) in technology.researchs" :key="index">
                  <ul class="card-researchs">
                    <li class="card-reserch" v-for="(content, index) in reserch.list" :key="index">
                      <div class="card-reserch-content">{{ content }}</div>
                    </li>
                  </ul>
                </div>
                <div class="card-link">
                  <a class="card-link-btn" :href="technology.link">
                    <span>进入开源社区</span>
                    <svg xmlns="http://www.w3.org/2000/svg"
                      class="styles__StyledSVGIconPathComponent-sc-gbp7ch-0 iVwQlg svg-icon-path-icon fill"
                      viewBox="0 0 32 32" width="16" height="16" fill="rgb(22, 132, 252)">
                      <defs data-reactroot=""></defs>
                      <g>
                        <path
                          d="M21.563 14.667l-7.152-7.152 1.885-1.885 10.371 10.371-10.371 10.371-1.885-1.885 7.152-7.152h-16.229v-2.667z">
                        </path>
                      </g>
                    </svg>
                  </a>
                </div>
              </div>
              <div class="card card-more">
                <div class="card-more-icon">…</div>
                <div class="card-more-content">更多问题期待您的参与</div>
              </div>
              <template v-if="(researchPlan.technologys.length + 1) % 3 != 0"></template>
              <div class="card card-placeholder" v-for="item in (3 - (researchPlan.technologys.length + 1) % 3)"
                :key="'place-' + item">
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>

import * as echarts from "echarts";
import { getPromoteData } from '~/apis/modules/common';

const templateOptions1 = {
  title: {
    text: '',
    left: 'center',
  },
  grid: {
    top: 48,
  },
  tooltip: {
    show: true,
    formatter: '{b0}: {c0}%'
  },
  xAxis: {
    type: 'category',
    axisLabel: {},
    data: [],
  },
  yAxis: {
    type: 'value',
    axisLabel: {
      formatter: '{value}%'
    },
  },
  series: [{
    data: [],
    type: 'bar',
    itemStyle: {
      color: '#188df0',
    },
    barWidth: 20,
    label: {
      show: true,
      position: 'top',
      formatter: '{c}%'
    }
  }]
};

const templateOptions2 = {
  title: {
    text: '',
    top: 0,
    left: 'center',
    textStyle: {
      lineHeight: 22,
    },
  },
  grid: {
    top: 68,
    bottom: 78,
    containLabel: true
  },
  tooltip: {
    trigger: 'axis'
  },
  xAxis: {
    type: 'category',
    data: [],
  },
  yAxis: {
    type: 'value',
    min: 20,
    max: 70,
  },
  series: []
};

const chartHandlers = [];

export default {
  data() {
    return {
      models: [{
        name1: '鹏城·脑海两千亿参数',
        name2: 'PengCheng-Mind-200B-1.0',
        status: [1, 2], // 1-可申请、2-持续训练中
        descr: [
          ['基于', '中国算力网', '深圳枢纽节点(鹏城云脑Ⅱ)'],
          ['模型分片', '3456'],
          ['基于', 'MindSpore'],
        ],
        link: 'https://openi.pcl.ac.cn/PengChengMind/PengCheng.Mind',
        ops: [],
        updateTime: '2023年11月',
      }, {
        name1: '鹏城·脑海七十亿参数',
        name2: 'PengCheng-Mind-7B-2.0',
        status: [1, 2],
        descr: [
          ['基于', '中国算力网', '西安、成都、大连节点'],
          ['支持', 'NPU转换GPU'],
          ['基于', 'MindSpore'],
        ],
        link: 'https://openi.pcl.ac.cn/PengChengMind/PengChengMind-7B',
        ops: [],
        updateTime: '2024年11月',
      }, {
        name1: '鹏城·脑海二十亿参数',
        name2: 'PengCheng-Mind-2B-2.0',
        status: [1, 2],
        descr: [
          ['基于', '中国算力网', '深圳枢纽(鹏城云脑Ⅱ)、武汉节点'],
          ['支持', 'NPU转换GPU'],
          ['基于', 'MindSpore'],
        ],
        link: 'https://openi.pcl.ac.cn/PCL-Platform.Intelligence/PengCheng.Mind-mindformer',
        ops: [],
        updateTime: '2024年11月',
      }, {
        name1: '鹏城·脑海 — 开元',
        name2: 'PCMind-2.1-Kaiyuan-2B',
        status: [1, 2],
        descr: [
          ['基于', '中国算力网', '深圳枢纽节点(鹏城云脑Ⅱ)'],
          ['支持', 'NPU转换GPU'],
          ['基于', 'MindSpore'],
        ],
        link: 'https://openi.pcl.ac.cn/PengChengMind/PCMind-Kaiyuan-2B',
        ops: [],
        updateTime: '',
      }],
      activeTab: 0,
      tabList: ['PengCheng-Mind-200B-1.0', 'PengCheng-Mind-7B-2.0', 'PCMind-2.1-Kaiyuan-2B'],
      tabContent: [{
        descr: [
          `PengCheng-Mind 200B 脑海大模型是基于Transformer 架构的2010亿参数自回归式语言模型。模型全流程基于中国算力网枢纽节点（鹏城云脑Ⅱ）的全自主安全可控国产软硬件平台进行开发和训练，采用MindSpore框架实现在大规模集群上长期稳定的多维分布式并行训练。训练模型保存为3456个切片，容量约7.6TB，推理模型保存为16个切片，容量约375GB。支持4096个token的上下文窗口长度，主要聚焦中文核心能力，兼顾英文和部分多语言能力。当前模型已完成训练1T Tokens数据量，仍在持续训练迭代中。`,
          `大模型的开发充满未知与挑战，随着中国算力网的不断发展壮大，希望通过模型训练全流程开源开放的方式，将鹏城·脑海200B 模型培育为算力网上一个种子应用，吸引产学研界更多人参与到通用大语言模型基座的研究和应用中，一起助力国产智能软硬件生态的发展。`,
        ],
        stage: [
          { name: '第一阶段', tokens: '100B Tokens', },
          { name: '第二阶段', tokens: '200B Tokens', },
          { name: '第三阶段', tokens: '300B Tokens', },
          { name: '第四阶段', tokens: '400B Tokens', },
          { name: '第五阶段', tokens: '500B Tokens', },
        ],
        chart: [{
          title: '0-150B Tokens主要数据分布',
          x: ['文学', '经济', '法律', '教育学', '医学'],
          y: [37.04, 11.98, 7.03, 4.83, 4.8],
        }, {
          title: '150-330B Tokens主要数据分布',
          x: ['文学', '经济', '法律', '医学', '历史'],
          y: [28.42, 9.59, 7.44, 5.53, 4.23],
        }, {
          title: '330-530B Tokens主要数据分布',
          x: ['文学', '经济', '法律', '医学', '教育学'],
          y: [24.94, 10.29, 7.73, 5.17, 3.97],
        }]
      }, {
        descr: [
          `PengCheng-Mind 7B 脑海大模型是基于中国算力网训练，可实现跨算力中心迁移的70亿参数自回归式语言模型，该模型支持NPU版本到GPU版本的转换，具有强大的语言理解和生成能力，可深入理解输入文本，自动生成连贯、合理的语句。模型可用于多种自然语言处理任务，如文本生成、机器翻译、问答系统等。PengCheng-Mind 7B在自然语言处理领域具有广阔的应用前景，有利于加速人工智能的学术研究和智能应用的开发过程。`,
        ],
        chart: [{
          type: 'full',
          title: '中文数据分布',
          x: ['文学', '经济', '其他', '政治', '教育学', '法律', '医学', '历史', '体育', '影视影评', '心理学', '艺术', '计算机科学', '农业科学', '交通运输', '安全科学与消防工程', '游戏娱乐', '环境科学', '电子信息科学', '管理学', '能源科学', '食品科学', '军事', '哲学', '城市规划与建设', '生物科学', '电子商务', '材料科学', '休闲旅游', '天文学', '数学科学', '气象学', '物理科学', '轻工业', '化学科学', '地理', '宗教', '通信技术', '广告营销', '重工业'],
          y: [7.87, 6.3, 6.3, 4.73, 4.73, 4.73, 3.94, 3.94, 3.15, 3.15, 3.15, 3.15, 3.15, 2.95, 2.36, 2.36, 2.36, 2.36, 2.36, 2.36, 2.36, 2.36, 1.58, 1.58, 1.58, 1.58, 1.58, 1.26, 1.18, 1.18, 1.18, 1.18, 1.18, 1.18, 0.79, 0.79, 0.79, 0.47, 0.39, 0.39],
        }, {
          type: 'full',
          title: '英文数据分布',
          x: ['Entertainment and tourism', 'Medicine and health', 'Business', 'Language and literature', 'History', 'Arts', 'Law', 'Computer Science', 'Energy science', 'Mathematics', 'Political Science', 'Management Science', 'Sociology', 'Physics', 'Religious Studies', 'Education', 'Military science', 'Human physical performance and recreation', 'Environment studies and forestry', 'Agriculture', 'Archaeology', 'Information and communication sciences', 'Engineering and technology', 'Philosophy', 'Astronomy', 'Geography', 'Biology', 'Transportation', 'Economics', 'Materials Science', 'Chemistry', 'Psychology', 'Food science', 'Communication technology', 'Earth Science', 'Aerography', 'Architecture and design'],
          y: [9.75, 4.86, 4.73, 4.66, 4.65, 4.16, 4.13, 4.02, 3.96, 3.73, 3.69, 3.38, 3.23, 2.89, 2.82, 1.98, 1.75, 1.59, 1.51, 1.39, 1.32, 1.27, 1.25, 1.16, 1.13, 1.12, 1.11, 0.87, 0.87, 0.85, 0.84, 0.84, 0.73, 0.69, 0.65, 0.48, 0.43],
        }],
        evaluation: {
          descr: '预训练模型PengCheng Mind-7B：学科、语言、知识、理解、推理能力持续提升，模型440B节点Checkpoint的学科综合任务(C-EVAL)得分47.3，知识问答任务(Webqa)得分64.3，蕴含推理任务上(CMNLI)得分49.75，科学文献理解(C3-m)得分65.39，指代消解语言任务(winogrande)得分65.19。',
          chartTitle: 'PengCheng-Mind 7B\nCheckpoints(in billions of tokens)',
          category: ['40B', '100B', '170B', '220B', '320B', '440B'],
          data: [{
            name: 'WebQA-知识',
            label: { show: true, },
            data: [
              { value: 32.2, label: { position: 'left' } },
              { value: 44.53, label: { position: 'top' } },
              { value: 49.51, label: { position: 'top' } },
              { value: 50.91, label: { position: 'top' } },
              { value: 61.45, label: { position: 'bottom' } },
              { value: 64.3, label: { position: 'bottom' } }
            ]
          }, {
            name: 'C-EVAL-学科',
            label: { show: true, },
            data: [
              { value: 27.5, label: { position: 'bottom' } },
              { value: 35.2, label: { position: 'bottom' } },
              { value: 39.5, label: { position: 'bottom' } },
              { value: 42, label: { position: 'bottom' } },
              { value: 44.3, label: { position: 'bottom' } },
              { value: 47.3, label: { position: 'bottom' } }
            ]
          }, {
            name: 'CMNLI-推理',
            label: { show: true, },
            data: [
              { value: 37.91, label: { position: 'top' } },
              { value: 42.56, label: { position: 'bottom' } },
              { value: 47.34, label: { position: 'bottom' } },
              { value: 49.71, label: { position: 'bottom' } },
              { value: 52.34, label: { position: 'bottom' } },
              { value: 49.75, label: { position: 'top' } }
            ]
          }, {
            name: 'c3-m-理解',
            label: { show: true, },
            data: [
              { value: 53.15, label: { position: 'bottom' } },
              { value: 58.41, label: { position: 'bottom' } },
              { value: 61.28, label: { position: 'top' } },
              { value: 61.47, label: { position: 'bottom' } },
              { value: 62.23, label: { position: 'left' } },
              { value: 65.39, label: { position: 'top' } }
            ]
          }, {
            name: 'winogrande-语言',
            label: { show: true, },
            data: [
              { value: 56.82, label: { position: 'top' } },
              { value: 59.82, label: { position: 'top' } },
              { value: 61.24, label: { position: 'bottom' } },
              { value: 62.98, label: { position: 'top' } },
              { value: 63.93, label: { position: 'top' } },
              { value: 65.19, label: { position: 'right' } }
            ]
          }]
        }
      }, {
        descr: [
          `PCMind-2.1-Kaiyuan-2B 是一款在昇腾 910A 集群上训练的、数据与权重全开源的语言模型。模型规模为 14 亿参数（加上词表后为20亿），在约 2.2 万亿个 token 上完成预训练。在同类全开源模型中，该模型表现突出，已能够在多项指标上对标乃至接近同规模的主流开源模型。`,
        ],
        features: {
          sum: '围绕高效训练与高质量模型的目标，我们在数据预处理与预训练流程中进行了系统性的设计与优化，主要包括以下创新点：',
          list: [
            {
              title: '1. 数据集质量基准评测',
              descr: '针对多个主流开源预训练数据集（如 DCLM Baseline、FineWeb-Edu），我们采用分位质量评测方法，并结合小规模对比实验系统刻画其质量分布，从而为后续的数据筛选与配比策略提供量化依据。'
            },
            {
              title: '2. 多阶段预训练策略',
              descr: '将整个训练过程划分为 5 个阶段，在阶段推进过程中逐步提升推理密集型与知识密集型样本的占比；同时，对其中质量较高的数据片段进行有选择的重复采样与利用，以更高效地发挥优质数据的价值。'
            },
            {
              title: '3. 多领域课程学习机制',
              descr: '在不同数据集之间，我们基本保持总体数据配比稳定；在各数据集内部，则按样本质量由低到高进行排序，并结合更适配课程学习的学习率衰减与模型合并策略（详见论文 “How Learning Rate Decay Wastes Your Best Data in Curriculum-Based LLM Pretraining”），从而充分发挥课程学习在大规模预训练中的潜力。'
            },
            {
              title: '4. 高性能数据预处理框架',
              descr: '我们构建了一个基于 Spark 的开源数据处理框架，并通过 Chukonu 进行深度优化，在大规模去重、排序等核心任务上实现了极高的吞吐与可扩展性，显著降低了数据工程环节的时间与资源成本。'
            },
            {
              title: '5. 面向训练稳定性的架构设计',
              descr: 'Kaiyuan-2B 的架构针对昇腾 910A 集群（主要是 FP16 的训练精度）进行了优化，引入并集成了 QK Norm、Sandwich Norm 与 Soft-Capping 等技术，有效提升了大规模预训练过程中的数值稳定性与训练鲁棒性。'
            },
          ]
        }
      }],
      researchPlan: null,
    };
  },
  components: {},
  methods: {
    changeTab(index) {
      this.activeTab = index;
      this.$nextTick(() => {
        this.drawChart();
      })
    },
    disposeChart() {
      for (let i = 0, iLen = chartHandlers.length; i < iLen; i++) {
        const chartHandler = chartHandlers[i];
        chartHandler && chartHandler.dispose();
      }
      chartHandlers.splice(0);
    },
    resize() {
      for (let i = 0, iLen = chartHandlers.length; i < iLen; i++) {
        const chartHandler = chartHandlers[i];
        chartHandler && chartHandler.resize();
      }
    },
    drawChart() {
      this.disposeChart();
      const charts = this.tabContent[this.activeTab].chart || [];
      for (let i = 0, iLen = charts.length; i < iLen; i++) {
        const chart = charts[i];
        const options = JSON.parse(JSON.stringify(templateOptions1));
        options.title.text = chart.title;
        options.xAxis.data = chart.x.slice(0);
        options.series[0].data = chart.y.slice(0);
        const chartEl = document.querySelector(`.chart-el-${i}`);
        if (!chartEl) continue;
        if (chart.type == 'full') {
          const clientWidth = chartEl.clientWidth;
          options.grid.containLabel = true;
          options.grid.bottom = 5;
          options.xAxis.axisLabel = {
            interval: 'auto',
            rotate: 45,
            width: 100,
            overflow: 'truncate',
            ellipsis: '...',
          };
          options.dataZoom = [{
            type: 'inside',
            start: 0,
            end: 100
          }];
          options.series[0].label.rotate = 30;
          options.series[0].barWidth = 'auto';
          options.series[0].barMaxWidth = 20;
          if (clientWidth < 1000) {
            options.series[0].label.show = false;
          }
        }
        if (chartEl) {
          const chartHandler = echarts.init(chartEl);
          chartHandler.setOption(options);
          chartHandlers.push(chartHandler);
        }
      }
      const evaluation = this.tabContent[this.activeTab].evaluation;
      const evaluationEl = document.querySelector(`.chart-el-eval`);
      if (evaluation && evaluationEl) {
        const options = JSON.parse(JSON.stringify(templateOptions2));
        options.title.text = evaluation.chartTitle;
        options.xAxis.data = evaluation.category.slice(0);
        options.legend = {
          top: 380,
          data: evaluation.data.map(item => item.name),
        };
        options.series = evaluation.data.map(item => {
          return {
            type: 'line',
            ...item,
          }
        })
        const chartHandler = echarts.init(evaluationEl);
        chartHandler.setOption(options);
        chartHandlers.push(chartHandler);
      }
    },
  },
  beforeCreate() {
    getPromoteData('/model/mind-open-source-plan.json').then(res => {
      try {
        const data = JSON.parse(res.data);
        this.researchPlan = data || '';
      } catch (err) {
        console.log(err);
      }
    }).catch(err => {
      console.log(err);
    })
  },
  mounted() {
    this.changeTab(0);
    window.addEventListener('resize', this.resize);
  },
  beforeDestroy() {
    this.disposeChart();
    window.removeEventListener('resize', this.resize);
  },
};
</script>

<style scoped lang="less">
.header {
  height: 580px;
  background: url('/img/model/PC-Mind-home.jpg') center center no-repeat;
  background-size: cover;
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;

  .title-0 {
    font-weight: 700;
    font-size: 36px;
    line-height: 48px;
    color: rgb(16, 16, 16);
    text-align: center;
    margin-top: -128px;
    margin-bottom: 48px;
    text-align: center;
    padding: 0 16px;
    max-width: 828px;
  }

  .title {
    font-weight: 700;
    font-size: 48px;
    line-height: 48px;
    color: rgb(16, 16, 16);
    margin-bottom: 48px;
    text-align: center;
  }

  .descr {
    font-weight: 400;
    font-size: 28px;
    color: rgba(16, 16, 16, 0.5);
    line-height: 36px;
    text-align: center;
    padding: 0 10px;
  }
}

.top {
  background-color: #f7f7f7;
  padding: 40px 0 40px 0;

  .content {
    .content-title {
      text-align: center;
      font-weight: 500;
      font-size: 28px;
      line-height: 32px;
      color: rgb(16, 16, 16);
      margin-bottom: 40px;
    }

    .content-c {
      display: flex;
      flex-wrap: wrap;

      .content-l {
        flex: 1;
        font-weight: 400;
        font-size: 16px;
        color: rgb(16, 16, 16);
        line-height: 30px;
        padding-left: 20px;
        padding-right: 20px;
        min-width: 50%;
        margin-bottom: 20px;


        p {
          margin-bottom: 15px;
        }

        .apply-btn {
          display: inline-block;
          height: 40px;
          border-color: rgb(22, 132, 252);
          border-width: 1px;
          border-style: solid;
          color: rgb(22, 132, 252);
          border-radius: 5px;
          font-size: 16px;
          padding: 0 18px;
          text-align: center;
          line-height: 38px;
          margin: 4px 0 20px 0;

          &.research-plan-btn {
            background: rgb(22, 132, 252);
            color: rgb(255, 255, 255);
            margin-left: 10px;
          }
        }

        .apply-descr {
          font-size: 14px;
          color: rgba(16, 16, 16, 0.5);
          line-height: 20px;
        }
      }

      .content-r {
        min-width: 50%;
        flex: 1;
        background: url('/img/model/mind.png') center top no-repeat;
        background-size: contain;
        margin-bottom: 20px;
      }
    }
  }
}


.models {
  padding: 40px 0;
  background-color: white;

  .model-card-c {
    display: flex;
    align-items: center;
    // justify-content: center;
    flex-wrap: wrap;
    gap: 20px 24px;

    .model-card {
      flex: 1;
      min-width: 40%;
      // max-width: 50%; // 540px;
      align-self: stretch;
      border-color: rgba(201, 188, 234, 0.4);
      border-width: 1px;
      border-style: solid;
      box-shadow: rgba(157, 197, 226, 0.2) 0px 5px 10px 0px;
      color: rgb(16, 16, 16);
      border-radius: 10px;
      font-size: 14px;
      padding: 25px;
      line-height: 20px;
      // margin: 10px 12px;
      background: url("data:image/svg+xml;charset=utf8,%3Csvg%20xmlns%3D%22http%3A%2F%2Fwww.w3.org%2F2000%2Fsvg%22%20xmlns%3Axlink%3D%22http%3A%2F%2Fwww.w3.org%2F1999%2Fxlink%22%20version%3D%221.1%22%3E%3Cdefs%3E%3ClinearGradient%20id%3D%221%22%20x1%3D%220%22%20x2%3D%221%22%20y1%3D%220%22%20y2%3D%220%22%20gradientTransform%3D%22matrix(-0.41100000000000014%2C%201.129%2C%20-0.9908929639889198%2C%20-0.41100000000000014%2C%200.911%2C%20-0.129)%22%3E%3Cstop%20stop-color%3D%22%23f1ebff%22%20stop-opacity%3D%221%22%20offset%3D%220.005%22%3E%3C%2Fstop%3E%3Cstop%20stop-color%3D%22%23eef2ff%22%20stop-opacity%3D%221%22%20offset%3D%220.19%22%3E%3C%2Fstop%3E%3Cstop%20stop-color%3D%22%23ffffff%22%20stop-opacity%3D%221%22%20offset%3D%220.59%22%3E%3C%2Fstop%3E%3C%2FlinearGradient%3E%3C%2Fdefs%3E%3Crect%20width%3D%22100%25%22%20height%3D%22100%25%22%20fill%3D%22url(%231)%22%3E%3C%2Frect%3E%3C%2Fsvg%3E");

      &:first-child {
        margin-left: 0;
      }

      &:last-child {
        margin-right: 0;
      }

      .model-name-1 {
        font-size: 18px;
        color: rgb(16, 16, 16);
        line-height: 26px;
        text-align: left;
        margin-bottom: 10px;
      }

      .model-name-2 {
        font-size: 24px;
        color: rgba(99, 102, 241, 1);
        line-height: 35px;
        text-align: left;
      }

      .model-status {
        display: flex;
        align-items: center;
        margin: 14px 0 8px 0;

        .status {
          font-size: 14px;
          color: rgba(22, 132, 252, 1);
          line-height: 20px;
          margin-right: 20px;
          display: flex;
          align-items: center;

          span {
            margin-left: 4px;
          }

          &.waiting {
            color: rgba(16, 16, 16, 0.5);
          }
        }
      }

      .model-descr-c {
        display: flex;
        margin-top: 20px;

        .descr {
          flex: 1;

          .descr-l1 {
            font-weight: 400;
            font-size: 14px;
            color: rgba(16, 16, 16, 0.5);
            line-height: 20px;
            margin-bottom: 8px;
          }

          .descr-l2 {
            font-weight: 700;
            font-size: 18px;
            color: rgb(16, 16, 16);
            line-height: 29px;
            margin-bottom: 8px;
          }

          .descr-l3 {
            font-weight: 400;
            font-size: 14px;
            color: rgb(16, 16, 16);
            line-height: 20px;
          }
        }

        .line {
          height: 86px;
          width: 1px;
          background-color: rgba(16, 16, 16, 0.1);
          margin: 0 16px;
        }
      }

      .link {
        display: flex;
        margin-top: 16px;

        a {
          display: flex;
          align-items: center;
          height: 40px;
          border-width: 1px;
          border-style: solid;
          color: rgb(16, 16, 16);
          border-radius: 20px;
          padding: 0 16px;
          border-color: rgb(201, 188, 234);
          background: rgb(238, 242, 255);
        }
      }

      .ops {
        a {
          display: inline-block;
          color: rgb(22, 132, 252);
          border-radius: 5px;
          font-size: 16px;
          padding: 8px 16px;
          text-align: center;
          line-height: 20px;
          background: rgba(1, 145, 255, 0.1);
        }
      }
    }
  }
}

.model-tab {
  display: flex;
  align-items: center;
  justify-content: center;

  .tab-item {
    display: flex;
    align-items: center;
    justify-content: center;
    width: 320px;
    height: 60px;
    border-width: 1px;
    border-style: solid;
    color: rgba(16, 16, 16, 0.5);
    border-color: transparent;
    border-radius: 10px 10px 0px 0px;
    text-align: center;
    font-weight: 500;
    font-size: 20px;
    line-height: 35px;
    cursor: pointer;

    &.active {
      color: rgba(99, 102, 241, 1);
      border-color: rgb(201, 188, 234);
      border-bottom-color: white;
      background-color: white;
    }
  }
}

.model-tab-content {
  margin-top: -2px;
  border-color: rgb(201, 188, 234);
  border-width: 1px;
  border-style: solid;
  box-shadow: rgba(157, 197, 226, 0.2) 0px 5px 10px 0px;
  color: rgb(16, 16, 16);
  border-radius: 0 0 10px 10px;
  font-size: 14px;
  padding: 32px;

  .descr p {
    font-size: 16px;
    color: #101010;
    line-height: 24px;
  }

  .data-distribution,
  .performance-evaluation,
  .life-stage {
    .title {
      font-weight: 700;
      font-size: 24px;
      color: rgb(16, 16, 16);
      line-height: 35px;
      margin-top: 32px;
      margin-bottom: 18px;
    }

    .descr {
      font-size: 16px;
      color: rgb(16, 16, 16);
      line-height: 24px;
    }

    .stage-c {
      display: flex;
      align-items: center;
      margin: 20px 0;
      padding: 0 20px;
      flex-wrap: wrap;

      .stage {
        display: flex;
        align-items: center;
        margin-bottom: 10px;
        width: 159px;

        .stage-item {
          width: 125px;
          height: 70px;
          color: rgb(16, 16, 16);
          border-radius: 5px;
          font-size: 16px;
          background: rgba(1, 145, 255, 0.1);
          text-align: center;
          margin-left: -6px;
          display: flex;
          flex-direction: column;
          align-items: center;
          justify-content: center;

          .name {
            margin-bottom: 4px;
          }

          .tokens {
            font-size: 14px;
          }
        }

        .stage-arr {
          width: 40px;
          margin-top: 4px;
        }
      }
    }
  }

  .chart-el-c {
    margin-top: 20px;
    display: flex;
    flex-wrap: wrap;

    .chart-el {
      flex-basis: calc(33%);
      height: 400px;
      margin-bottom: 40px;
    }

    .chart-el-full {
      height: 450px;
      flex-basis: calc(100%);
    }
  }

  .features {
    margin-top: 1.2em;

    .sum {
      font-size: 16px;
      color: #101010;
      line-height: 24px;
    }

    .feature {
      margin-top: 15px;

      .title {
        font-size: 16px;
        font-weight: bold;
        color: #101010;
        line-height: 24px;
        margin-bottom: 10px;
      }

      .descr {
        font-size: 16px;
        color: #101010;
        line-height: 24px;
      }
    }
  }
}

.research-plan-field {
  margin-top: 40px;
  padding-top: 20px;
  padding-bottom: 40px;
  background: url("data:image/svg+xml;charset=utf8,%3Csvg%20xmlns%3D%22http%3A%2F%2Fwww.w3.org%2F2000%2Fsvg%22%20xmlns%3Axlink%3D%22http%3A%2F%2Fwww.w3.org%2F1999%2Fxlink%22%20version%3D%221.1%22%3E%3Cdefs%3E%3ClinearGradient%20id%3D%221%22%20x1%3D%220%22%20x2%3D%221%22%20y1%3D%220%22%20y2%3D%220%22%20gradientTransform%3D%22matrix(0.11899999999999993%2C%201.217%2C%20-5.771885432098765%2C%200.11899999999999993%2C%200.269%2C%20-0.22)%22%3E%3Cstop%20stop-color%3D%22%23ffffff%22%20stop-opacity%3D%220.47%22%20offset%3D%220%22%3E%3C%2Fstop%3E%3Cstop%20stop-color%3D%22%23e5e7eb%22%20stop-opacity%3D%220.3%22%20offset%3D%221%22%3E%3C%2Fstop%3E%3C%2FlinearGradient%3E%3C%2Fdefs%3E%3Crect%20width%3D%22100%25%22%20height%3D%22100%25%22%20fill%3D%22url(%231)%22%3E%3C%2Frect%3E%3C%2Fsvg%3E");
}

.research-plan {
  color: rgba(16, 16, 16, 1);

  .title {
    font-size: 28px;
    line-height: 40px;
    margin-bottom: 32px;
    text-align: center;
  }

  .descr-c {
    margin-bottom: 30px;
    display: flex;
    justify-content: center;

    .descr {
      font-size: 16px;
      line-height: 24px;
      width: 80%;
      max-width: 1000px;
    }
  }


  .main-pic {
    text-align: center;
    margin-bottom: 16px;

    img {
      width: 75%;
      max-width: 800px;
    }
  }

  .cards-c {
    margin-top: 10px;

    .cards {
      display: flex;
      flex-wrap: wrap;
      justify-content: center;

      .card {
        flex-basis: calc(28%);
        max-width: 540px;
        align-self: stretch;
        border-radius: 10px;
        background: linear-gradient(180deg, rgba(255, 255, 255, 1) 0%, rgba(255, 255, 255, 1) 100%);
        text-align: center;
        box-shadow: 0px 5px 10px 0px rgba(157, 197, 226, 0.2);
        border: 1px solid rgba(201, 188, 234, 1);
        padding: 26px 26px;
        margin: 20px;

        .card-title {
          display: flex;

          .card-num {
            width: 50px;
            height: 50px;
            line-height: 20px;
            background: linear-gradient(-47.63deg, rgba(22, 132, 252, 1) 17.44%, rgba(75, 228, 228, 0) 88.44%);
            text-align: center;
            color: rgba(255, 255, 255, 1);
            font-size: 40px;
            border-radius: 50%;
            line-height: 1.2;
            font-weight: bold;
          }

          .card-name-c {
            width: 0;
            flex: 1;
            display: flex;
            align-items: center;

            .card-name {
              text-align: left;
              width: 100%;
              overflow: hidden;
              text-overflow: ellipsis;
              white-space: nowrap;
              color: #101010;
              font-size: 20px;
              margin-left: 8px;
              line-height: 1.2;
            }
          }
        }

        .card-descr {
          margin-top: 20px;
          text-align: left;
          font-size: 16px;
          line-height: 24px;
          margin-bottom: 16px;
        }

        .card-researchs-title {
          text-align: left;
          font-size: 16px;
          line-height: 26px;
          margin-bottom: -10px;
        }

        .card-researchs {
          text-align: left;
          padding-left: 20px;

          .card-reserch {
            .card-reserch-content {
              margin-left: 8px;
              font-size: 16px;
              line-height: 26px;
            }
          }
        }

        .card-link {
          text-align: left;
          display: flex;
          margin-top: 4px;

          .card-link-btn {
            display: flex;
            align-items: center;
            height: 40px;
            // border-color: rgb(22, 132, 252);
            // border-width: 1px;
            // border-style: solid;
            // color: rgb(22, 132, 252);
            // border-radius: 6px;
            font-size: 16px;
            padding: 10px 10px;

            span {
              margin-right: 2px;
            }
          }
        }

        &.card-more {
          background: none;
          display: flex;
          flex-direction: column;
          align-items: center;
          justify-content: center;
          min-height: 250px;

          .card-more-icon {
            border-radius: 50%;
            width: 59px;
            height: 59px;
            line-height: 32px;
            background: linear-gradient(-47.63deg, rgba(201, 188, 234, 1) 17.44%, rgba(187, 187, 187, 0.1) 88.44%);
            margin-top: -30px;
            color: rgba(255, 255, 255, 1);
            font-size: 52px;
            text-align: center;
          }

          .card-more-content {
            margin-top: 26px;
            height: 20px;
            color: rgba(16, 16, 16, 1);
            font-size: 14px;
          }
        }

        &.card-placeholder {
          background: none;
          border-color: transparent;
          box-shadow: none;
        }
      }
    }
  }
}

@media only screen and (max-width: 1600px) {
  .research-plan .cards-c {

    .cards {

      .card {
        flex-basis: calc(44%);
      }
    }
  }
}

@media only screen and (min-width: 1200px) {
  .model-tab-content {
    border-radius: 10px;
  }

  .model-tab-content .chart-el-c {
    .chart-el {
      flex-basis: calc(33.3%);
    }

    .chart-el-full {
      flex-basis: calc(100%);
    }
  }
}

@media only screen and (max-width: 1200px) {
  .model-tab-content .chart-el-c {
    .chart-el {
      flex-basis: calc(49%);
    }

    .chart-el-full {
      flex-basis: calc(100%);
    }
  }
}

@media only screen and (max-width: 1000px) {
  .models .model-card-c {
    display: block;
  }

  .models .model-card-c .model-card {
    margin: 10px auto 30px;
  }

  .research-plan .cards-c .cards {
    display: block;

    .card {
      margin: 40px auto;

      &.card-placeholder {
        display: none;
      }
    }
  }
}

@media only screen and (max-width: 767px) {

  .header .title-0 {
    font-size: 20px;
    margin-bottom: 32px;
  }

  .header .title {
    font-size: 28px;
  }

  .header .descr {
    font-size: 18px;
  }

  .top {
    padding-top: 20px;
    padding-bottom: 20px;
  }

  .top .content .content-title {
    font-size: 20px;
    margin-bottom: 20px;
  }

  .top .content .content-c .content-l .apply-c {
    text-align: center;
  }

  .top .content .content-c .content-l .apply-descr {
    text-align: center;
  }

  .top .content .content-c .content-l p {
    font-size: 14px;
    line-height: 23px;
  }

  .top .content .content-c .content-l,
  .top .content .content-c .content-r {
    min-width: 100%;
    min-height: 200px;
    padding-left: 0px;
    padding-right: 0px;
  }

  .models {
    padding: 20px 0;
  }

  .models .model-card-c {
    display: block;
  }

  .models .model-card-c .model-card {
    margin: 10px auto 20px;

    .model-name-1 {
      font-size: 16px;
    }

    .model-name-2 {
      font-size: 20px;
    }

    .model-descr-c .descr .descr-l1 {
      font-size: 12px;
    }

    .model-descr-c .descr .descr-l2 {
      font-size: 16px;
      line-height: 18px;
    }

    .model-descr-c .descr .descr-l3 {
      font-size: 12px;
      line-height: 14px;
    }
  }

  .models .model-card-c .model-card .link a {
    height: 34px;
    border-radius: 17px;
    padding: 0 12px;
    font-size: 14px;
  }

  .model-tab .tab-item {
    font-size: 14px;
    height: 46px;
    line-height: 18px;
  }

  .model-tab-content {
    padding: 12px;
    margin-top: -2px;

    .descr p {
      padding-top: 2px;
      font-size: 14px;
    }
  }

  .model-tab-content .life-stage .descr {
    font-size: 14px;
  }

  .model-tab-content .life-stage .stage-c {
    display: flex;
    overflow: auto;
    flex-wrap: nowrap;
  }

  .model-tab-content .life-stage .title,
  .model-tab-content .data-distribution .title {
    font-size: 18px;
    margin-top: 20px;
    margin-bottom: 10px;
  }

  .model-tab-content .chart-el-c {
    .chart-el {
      flex-basis: calc(99%);
    }
  }

  .model-tab-content .features .sum,
  .model-tab-content .features .feature .descr {
    font-size: 14px;
  }

  .research-plan-field {
    margin-top: 32px;
  }

  .research-plan {

    .title {
      font-size: 20px;
      margin-bottom: 20px;
    }

    .descr-c {
      margin-bottom: 20px;

      .descr {
        font-size: 14px;
        width: 100%;
      }
    }

    .main-pic {
      img {
        width: 94%;
      }
    }
  }

  .research-plan .cards-c .cards {
    display: block;

    .card {
      margin: 20px auto;
      padding: 20px 20px;

      .card-title {
        .card-num {
          height: 40px;
          width: 40px;
          font-size: 32px;
        }

        .card-name-c {
          .card-name {
            font-size: 18px;
            font-weight: bold;
          }
        }

        .card-descr {
          font-size: 14px;
        }

        .card-researchs {
          font-size: 14px;
        }
      }
    }
  }
}
</style>