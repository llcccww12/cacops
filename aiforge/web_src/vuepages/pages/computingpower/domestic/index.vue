<template>
  <div class="ui container">
    <div class="top-field">
      <div class="title">Century Albert 国产算力英雄榜</div>
      <div class="descr">本页面仅统计国产算力在 CacOps 智算服务平台的使用情况。数据更新时间：<span>{{ updateTime }}</span></div>
    </div>
    <div class="sort-field">
      <div class="sort-conds">
        <div class="sort-tab-c">
          <div class="tab" :class="item.key == tabIndex ? 'active' : ''" v-for="(item) in tabList" :key="item.key"
            @click="changeTab(item)">
            {{ item.label }}
          </div>
        </div>
        <div class="sort-type-c">
          <el-select :size="useSmall ? 'small' : 'default'" v-model="sortType" @change="changeSort">
            <el-option v-for="(item) in sortList" :key="item.key" :value="item.key" :label="item.label"></el-option>
          </el-select>
        </div>
      </div>
      <div class="table-container" style="min-height:360px;">
        <el-table class="table" :data="tableData" style="width:100%" v-loading="loading" row-key="id">
          <el-table-column prop="rank" label="排名" align="center" header-align="center" width="80"></el-table-column>
          <el-table-column prop="card_type" label="卡类型" align="left" header-align="left"></el-table-column>
          <el-table-column prop="resource_type" label="计算资源" align="left" header-align="left">
          </el-table-column>
          <el-table-column prop="company" label="厂家" align="left" header-align="left"></el-table-column>
          <el-table-column prop="access_time" label="接入平台时间" align="center" header-align="center"
            min-width="110"></el-table-column>
          <el-table-column prop="sortValue" align="left" header-align="center" min-width="360">
            <template #header>
              <div v-html="sortObj.theader"></div>
            </template>
            <template slot-scope="scope">
              <div class="table-bar">
                <div class="bar-c">
                  <div class="bar"
                    :style="`width:${scope.row.barWith}%;${scope.row.barWith > 0 ? 'min-width:1px' : ''}`">
                    <template v-if="scope.$index == 0 && useShrink">
                      <div></div>
                      <div></div>
                      <div></div>
                    </template>
                  </div>
                </div>
                <div class="bar-value">{{ scope.row.countShow }}</div>
              </div>
            </template>
          </el-table-column>
          <template slot="empty">
            <span style="font-size: 12px">{{
              loading ? $t('loading') : $t('noData')
              }}</span>
          </template>
        </el-table>
      </div>
    </div>
    <div class="partners-field">
      <div class="partners-title">Century Albert 国产算力合作伙伴</div>
      <div class="partners-c">
        <div class="img-c" v-for="(item, index) in partners" :key="index">
          <img :src="item.icon" :alt="item.name">
        </div>
      </div>
    </div>
    <div class="apply-field">
      <a :href="applyLink" class="apply-btn">加速卡申请加入 CacOps 智算服务平台请参考</a>
    </div>
    <div class="card-field">
      <div class="card" v-for="(card, cardIndex) in cards" :key="cardIndex">
        <div class="title-c">
          <div class="title">{{ card.name }}</div>
          <div class="icon-c">
            <img :src="card.icon" alt="" />
          </div>
        </div>
        <div class="descr-c" v-if="card.descr">
          <p v-for="(descr, descrIndex) in card.descr" :key="descrIndex"> {{ descr }}</p>
        </div>
        <div class="feature-c" v-if="card.features">
          <div class="feature" v-for="(feature, featureIndex) in card.features" :key="featureIndex">
            <p class="title">{{ feature.title }}</p>
            <p class="item" v-for="(item, itemIndex) in feature.list" :key="itemIndex"> {{ item }}</p>
          </div>
        </div>
        <div class="table-c" v-if="card.table">
          <table class="table" :class="card.table.tableClass">
            <tr>
              <td class="title" colspan="2">{{ card.table.title }}</td>
            </tr>
            <tr v-for="(row, rowIndex) in card.table.fields" :key="rowIndex">
              <td class="field-title" v-html="row.title"></td>
              <td v-html="row.value"></td>
            </tr>
          </table>
        </div>
        <div class="use-example-c" v-if="card.useExample">
          <div class="title">{{ card.useExample.title }}</div>
          <div v-for="(link, linkIndex) in card.useExample.list" :key="linkIndex">
            <a :href="link"> {{ link }}</a>
          </div>
        </div>
      </div>
      <div class="card placeholder" v-if="cards.length % 2 == 1"></div>
    </div>
  </div>
</template>

<script>
import { getDomesticCardData } from '~/apis/modules/computingpower';
import { formatDate } from 'element-ui/lib/utils/date-util';

const manufacturerIconMap = {
  ascend: '/img/domestic/ascend.svg',
  enflame: '/img/domestic/enflame.svg',
  cambricon: '/img/domestic/cambricon.svg',
  iluvatar: '/img/domestic/iluvatar.svg',
  metax: '/img/domestic/metax.svg',
  hygon: '/img/domestic/hygon.png',
  biren: '/img/domestic/biren.png'
};

export default {
  data() {
    return {
      isOperator: false,
      loading: false,
      useSmall: false,
      tabList: [
        { key: 'all', label: '总排名' },
        { key: '7', label: '近7天排名' },
        { key: '30', label: '近30天排名' },
      ],
      tabIndex: 'all',
      sortList: [
        { key: 'card', label: '使用时长', theader: '使用时长（卡时）<span>说明：卡时=云脑任务运行时长*卡数</span>', unit: '卡时' },
        { key: 'user', label: '使用用户数', theader: '使用用户数（人）', unit: '人' },
        { key: 'task', label: '使用次数（创建云脑任务数）', theader: '使用次数（个）', unit: '个' },
      ],
      sortType: 'card',
      tableData: [],
      updateTime: '',
      useShrink: false,
      partners: [
        { name: '', icon: manufacturerIconMap.ascend },
        { name: '', icon: manufacturerIconMap.enflame },
        { name: '', icon: manufacturerIconMap.cambricon },
        { name: '', icon: manufacturerIconMap.iluvatar },
        { name: '', icon: manufacturerIconMap.metax },
        { name: '', icon: manufacturerIconMap.hygon },
        { name: '', icon: manufacturerIconMap.biren },
      ],
      applyLink: 'https://openi.pcl.ac.cn/OpenIOSSG/promote/src/branch/master/XPRZ.md',
      cards: [{
        name: '华为昇腾Ascend 910',
        icon: manufacturerIconMap.ascend,
        descr: [
          '昇腾（HUAWEI Ascend) 910是业界算力最强的AI处理器，基于自研华为达芬奇架构3D Cube技术，实现业界最佳AI性能与能效，架构灵活伸缩，支持云边端全栈全场景应用。算力方面，昇腾910完全达到设计规格，半精度（FP16）算力达到320 TFLOPS，整数精度（INT8）算力达到640 TOPS，功耗310W。'
        ],
        features: [
          { title: '产品特点', list: ['- 自研华为达芬奇架构NPU', '- 640 TOPS@INT8，320TFLOPS@FP16', '- 最大功耗310W'], },
        ],
        table: {
          title: '关键特性',
          fields: [
            { title: 'Architecture', value: 'HUAWEI Da Vinci' },
            { title: 'Computing Engine', value: '3D Cube' },
            { title: 'Performance', value: '320 TFLOPS @FP16 and 640 TOPS @INT8' },
            { title: 'Max Power', value: '310W' },
            { title: 'Process', value: 'N7+' },
          ]
        },
        useExample: {
          title: '在 CacOps 智算服务平台使用案例：',
          list: [
            'https://openi.pcl.ac.cn/OpenIOSSG/MNIST_Example'
          ]
        },
      }, /*{
        name: '华为昇腾Ascend D910B',
        icon: manufacturerIconMap.ascend,
        descr: [
          '昇腾（HUAWEI Ascend) 910是业界算力最强的AI处理器，基于自研华为达芬奇架构3D Cube技术，实现业界最佳AI性能与能效，架构灵活伸缩，支持云边端全栈全场景应用。算力方面，昇腾910完全达到设计规格，半精度（FP16）算力达到320 TFLOPS，整数精度（INT8）算力达到640 TOPS，功耗310W。'
        ],
        table: {
          title: '关键特性',
          fields: [
            { title: 'Architecture', value: 'HUAWEI Da Vinci' },
            { title: 'Computing Engine', value: '3D Cube' },
            { title: 'Performance', value: '320 TFLOPS @FP16 and 640 TOPS @INT8' },
            { title: 'Max Power', value: '310W' },
            { title: 'Process', value: 'N7+' },
          ]
        }
      },*/ {
        name: '燧原科技云燧ENFLAME-T20',
        icon: manufacturerIconMap.enflame,
        descr: [
          '云燧T20是基于邃思2.0芯片打造的面向数据中心的第二代人工智能训练加速卡，具有模型覆盖面广、性能强、软件生态开放等特点，可支持多种人工智能训练场景。同时具备灵活的可扩展性，提供业界领先的人工智能算力集群方案。产品优势特点如下：'
        ],
        features: [
          { title: '澎湃算力 高精训练', list: ['- 领先的TF32等浮点AI算力', '- 基于HBM2E的高吞吐低延时', '- 动态性特征支持'], },
          { title: '专属通道 算力扩展', list: ['- 独立的高带宽通道，加速卡间通信', '- 独家的机内4卡全互联方案', '- 增强的单机8卡互联方案'], },
          { title: '广泛支持 生态友好', list: ['- C++和Python开发接口', '- 主流框架支持，国产框架适配加速', '- 算子与模型广泛支持'], },
          { title: '工具开放 高效开发', list: ['- 多层次API接口开放', '- 完整工具链，支持高效开发与模型调试', '- 编程模型开放，支持第三方深度定制'], },
        ],
        useExample: {
          title: '在 CacOps 智算服务平台使用案例：',
          list: [
            'https://openi.pcl.ac.cn/Enflame/GCU_Pytorch',
            'https://openi.pcl.ac.cn/Enflame/GCU_PaddlePaddle_Example'
          ]
        },
      }, {
        name: '燧原科技云燧ENFLAME-I20',
        icon: manufacturerIconMap.enflame,
        descr: [
          '云燧I20是基于邃思2.5芯片打造的面向数据中心的第二代人工智能推理加速卡，具有高性能高能效、模型覆盖面广、易部署易运维等特点，可广泛应用于计算机视觉、语音识别与合成、自然语言处理、搜索与推荐等推理场景。产品优势特点如下：'
        ],
        features: [
          { title: '计算引擎', list: ['- 支持FP32、FP16、BF16、INT8等多种数据精度，提供全精度支持和模型性能', '- 高可编程性，支持矢量、张量等多种计算类型，支持超越函数计算加速'], },
          { title: '存储引擎', list: ['- 3层存储结构设计，基于深度学习推理计算负载进行了深度效率优化', '- 国内首张支持HBM2E存储方案的推理加速卡，提供超大存储带宽'], },
          { title: '数据引擎', list: ['- 面向张量AI数据流计算加速的数据引擎，支持切分/逆切分、维度变换、拼接、降采样、镜像、常量填充等', '- 支持融合型张量操作', '- 支持数据多地址广播', '- 硬件多重循环指令，降低同步开销'], },
          { title: '工具开放 高效开发', list: ['- 驭算TopsRider是燧原科技自主知识产权的计算及编程平台，通过软硬件协同的架构设计，充分释放云燧I20的性能。驭算2.0软件栈，通过软硬件结合提供极致的性能和优化的编程环境，进一步降低了用户的迁移成本和学习成本。驭算TopsRider适配主流框架，提供工具链支持开发与调优，不断提升用户的使用体验。'], },
        ],
        useExample: {
          title: '在 CacOps 智算服务平台使用案例：',
          list: [
            'https://openi.pcl.ac.cn/Enflame/SD_Inference',
          ]
        },
      }, {
        name: '寒武纪MLU290',
        icon: manufacturerIconMap.cambricon,
        descr: [
          'MLU290-M5智能加速卡搭载寒武纪首颗训练芯片思元290，采用台积电7nm先进制程工艺，采用MLUv02扩展架构。MLU290-M5智能加速卡采用开放加速模块OAM设计，具备64个MLU Core，1.23TB/s内存带宽以及全新MLU-Link芯片间互联技术，全面支持AI训练、推理或混合型人工智能计算加速任务。',
        ],
        table: {
          title: '思元290-M5 产品规格',
          fields: [
            { title: '产品名称', value: 'MLU290-M5' },
            { title: '核心架构', value: 'Cambricon MLUv02 Extended' },
            { title: '制程工艺', value: '7nm' },
            { title: '自适应精度训练算力', value: '512 TOPS(INT8)<br>256 TOPS(INT16)<br>64 TOPS(CINT32)' },
            { title: 'DirectCV<sup>TM</sup>视频解码', value: '128 Streams 全高清视频' },
            { title: 'DirectCV<sup>TM</sup>图片解码', value: '3200 Frames/s 全高清图片' },
            { title: '内存类型', value: 'HBM2高带宽内存' },
            { title: '内存容量', value: '32GB' },
            { title: '内存位宽', value: '4096 bit' },
            { title: '内存带宽', value: '1228 GB/s' },
            { title: '系统接口', value: 'x 16 PCIe 4.0' },
            { title: 'MLU-Link<sup>TM</sup>接口', value: '6 Ports, 48 Lanes, 50 Gbps' },
            { title: 'MLU-Link<sup>TM</sup>带宽', value: '聚合带宽600GB/s Bi-direction' },
            { title: '最大热功耗', value: '350W' },
            { title: '形态', value: 'OAM (54V)' },
            { title: '尺寸', value: '102mm x 165mm' },
            { title: '含散热器重量', value: '1470g' },
          ]
        }
      }, {
        name: '天垓100（BI -V100）',
        icon: manufacturerIconMap.iluvatar,
        descr: [
          '天垓100通用GPU训练加速卡，产品特点如下：'
        ],
        features: [
          { title: '应用覆盖广', list: ['天垓100聚焦高性能、通用性和灵活性，支持200余种人工智能模型（数量持续增加），支持通用计算、科学计算、大模型、支持业界前沿新算法模型。模型适配速度快，从容面对未来的算法变迁，为人工智能及通用计算和相关垂直应用行业提供匹配行业高速发展的计算力。'], },
          { title: '性能可预期', list: ['天垓100基于通用GPU架构设计，拥有丰富的自研指令集全方位支持标量、矢量、张量运算，提供业界领先的高算力和高能效比，在百余个算法模型的测试平均性能可媲美主流产品。'], },
          { title: '开发易迁移', list: ['适配主流 CPU 芯片 / 服务器厂商，能够支持国内外主流软硬件生态和各种深度学习框架、算法模型和加速库，并通过标准化的软硬件生态接口为行业解决产品使用难、开发平台迁移成本大等痛点，应用迁移成本低、耗时短、无需重新开发，大幅缩短适配验证周期，使客户业务系统几乎无感知地使用天垓 100 产品。'], },
          { title: '全栈可定制', list: ['天垓100 的核心IP 、系统架构、指令集、核心算子、软件栈均由天数科学家团队开发完成，本地支持团队能够根据客户需求提供定制化开发服务。'], },
        ],
        useExample: {
          title: '在 CacOps 智算服务平台使用案例：',
          list: [
            'https://openi.pcl.ac.cn/iluvatar/TianGai100',
            'https://openi.pcl.ac.cn/iluvatar/bert_crf_sequence_labeling',
            'https://openi.pcl.ac.cn/iluvatar/resnet50',
            'https://openi.pcl.ac.cn/iluvatar/paddleyolo',
          ]
        },
      }, {
        name: '智铠100（MR-V100）',
        icon: manufacturerIconMap.iluvatar,
        descr: [
          '智铠100系列通用GPU推理加速卡，产品特点如下：'
        ],
        features: [
          { title: '计算性能高', list: ['支持FP32、FP16、INT8等多精度推理混合计算，实现了指令集增强、算力密度提升、计算存储再平衡，相较于市场上现有主流产品，智铠100将提供2-3倍的实际使用性能。'], },
          { title: '应用覆盖广', list: ['智铠100系列加速卡基于通用GPU架构，支持多种视频规格解码、800+通用指令集、国内外主流深度学习开发框架，拥有丰富编程接口拓展和高性能函数库，可以灵活支持各种算法模型，便于客户自定义开发。'], },
          { title: '使用成本低', list: ['兼容CUDA生态，支持市场主流生态，高达128路视频接入，单路视频性价比高灵活的编程能力，超强的性能及富有吸引力的性价比，为高性能计算和人工智能应用的开发和部署提供了便利。平均迁移时间相较市场主流产品下降50%以上，生态应用迁移迅速。'], },
          { title: '落地支持强', list: ['智铠100全自研架构、核心及软件栈，支持算力开发与优化，可根据客户需求提供定制化开发服务。'], },
        ],
        useExample: {
          title: '在 CacOps 智算服务平台使用案例：',
          list: [
            'https://openi.pcl.ac.cn/iluvatar/iluva202401091747306',
            'https://openi.pcl.ac.cn/iluvatar/bert_crf_sequence_labeling',
            'https://openi.pcl.ac.cn/iluvatar/resnet50',
            'https://openi.pcl.ac.cn/iluvatar/paddleyolo',
          ]
        },
      }, {
        name: '沐曦曦思N100',
        icon: manufacturerIconMap.metax,
        descr: [
          '沐曦在2020年9月成立于上海，它致力于为异构计算提供安全可靠的通用GPU芯片及解决方案，可广泛应用于人工智能、智慧城市、数据中心、云计算、自动驾驶、数字孪生、元宇宙等前沿领域，为数字经济发展提供强大的算力支撑。',
          '此次 CacOps 智算服务平台上线的沐曦GPGPU计算资源主要来自于沐曦的首款产品：“曦思N100” 人工智能推理GPU。曦思N100是一款面向云端数据中心应用的GPU人工智能加速卡，可广泛应用于智慧城市、智慧交通、机器视觉、智能视频处理等场景，该产品具备如下特点：'
        ],
        features: [
          { title: '创新架构，多样算力', list: ['- 面向场景优化的异构GPGPU架构，兼顾计算效率和通用性', '- 单卡提供高达160TOPS INT8 算力和80TFLOPS FP16 算力'], },
          { title: '卓越的视频处理能力', list: ['- 多种视频格式兼容且支持8K分辨率', '- 超高密度视频编解码，在基于视频+AI的场景极具性价比'], },
          { title: '高带宽，低延时', list: ['- 集成HBM2E 高带宽内存', '- 率先支持PCIe5.0'], },
          { title: '成熟易用的软件栈', list: ['- 兼容ONNX、Opencv、FFmpeg等主流框架', '- 完善的Model Zoo和应用示例，开箱即用'], },
        ],
        table: {
          title: '曦思N100人工智能推理GPU产品规格',
          fields: [
            { title: 'AI算力', value: '160TOPS INT8; 80TFLOPS FP16/BF16' },
            { title: '显存', value: '16GB HBM2E, 带宽460GB/s' },
            { title: '视频解码', value: '120x1080P30, H.264/H.265/AV1/AVS2, 支持8K分辨率' },
            { title: '视频编码', value: '120x1080P30, H.264/H.265/AV1, 支持8K分辨率' },
            { title: 'PCIe接口', value: 'PCIe5.0/4.0/3.0 x16' },
            { title: '虚拟化', value: '单卡可支持1/2/4个用户实例' },
            { title: '功耗', value: '最大70W' },
            { title: '散热设计', value: '被动散热' },
            { title: '板卡形态', value: '单槽半高半长(HHHL)' },
          ]
        },
        useExample: {
          title: '在 CacOps 智算服务平台使用案例：',
          list: [
            'https://openi.pcl.ac.cn/Metax/Metax202309121140356/src/branch/master/classification',
            'https://openi.pcl.ac.cn/Metax/Metax202309121140356/src/branch/master/detecion'
          ]
        },
      }, {
        name: '海光DCU',
        icon: manufacturerIconMap.hygon,
        descr: [
          'DCU深算处理器是由海光信息技术有限公司开发的一款高性能通用异构处理器。',
          'DCU深度计算处理器（DCU Z100）基于通用图形处理器理念设计，更加适合为人工智能计算提供强大的算力。DCU Z100拥有64组计算单元，共计8192个计算核心，超高速32GB HBM2内存和高达1TB/s的内存带宽，可以完美支持深度学习训练场景，轻松应对复杂神经网络训练。',
          'DCU Z100拥有10.8T的双精度浮点运算能力，是高性能计算业务的不二之选。优异的架构设计和超高计算能力，可为生命科学、能源勘探、工业设计、航空航天、分子动力学计算等行业提供更高性价比的计算方案。',
        ],
        table: {
          title: 'DCU Z100规格参数',
          fields: [
            { title: 'Category', value: 'Z100' },
            { title: 'Architecture', value: 'Zifang' },
            { title: 'Process (nm)', value: '7' },
            { title: 'Streaming Multiprocessors(SMs/CUs)', value: '64' },
            { title: 'FP32 Units / SM (CU)', value: '64' },
            { title: 'FP32 Units (Vector)', value: '4096' },
            { title: 'FP64 Units (Vector)', value: '4096' },
            { title: 'Tensor Core/Matrix Units', value: '-' },
            { title: 'Threads / Warp', value: '64' },
            { title: 'Max Warps / SM (CU)', value: '40' },
            { title: 'Max Threads / SM (CU)', value: '2560' },
            { title: 'Sparse Matrix', value: 'No' },
            { title: 'Peak FP64 Perf. (TFLOPs)', value: '10.8' },
            { title: 'Peak FP32 Perf. (TFLOPs)', value: '14.7' },
            { title: 'Peak FP16 Perf. (TFLOPs)', value: '29.5' },
            { title: 'Boost Core Clock Speed (MHz)', value: '1319' },
            { title: 'HBM2/HBM2E Memory (GB)', value: '16/32' },
            { title: 'DIE Size (mm<sup>2</sup>)', value: '356' },
            { title: 'TDP (W)', value: '350' },
          ],
          tableClass: 'average',
        }
      }, {
        name: '壁砺™ 106M',
        icon: manufacturerIconMap.biren,
        descr: [
          '壁仞科技通用GPU产品壁砺系列，由多款产品组成，基于壁仞科技原创的训推一体芯片架构，具有高算力、高能效比、高通用性的优势，能够为包括人工智能训练与推理在内的诸多通用计算场景提供强大的算力。壁砺™ 106M产品形态为风冷 OAM 模组，单卡峰值功耗 400W,强大的算力和高速互连能力，可为广泛的人工智能训练与推理场景，提供高能效比的算力解决方案。'
        ],
        features: [
          { title: '通用智算，训推平衡', list: ['壁砺™166系列产品可广泛应用于AI模型训练、微调与推理场景，赋能大模型应用高效落地。壁砺™106M 支持PCIe5.0 X8 主机互连接口技术，并通过自研UBB 主板实现单机8 卡互连。'], },
          { title: '原创设计，软硬协同', list: ['BIRENSUPA™是一个具有完整功能架构的软件开发平台，包括硬件抽象层、壁仞原创BIRENSUPA™编程模型和BRCC编译器，深度学习和通用计算加速库、工具链，支持主流深度学习框架和自研推理加速引擎，并配备针对不同场景的应用SDK等，能够为开发者提供高效的应用开发平台，软硬件协同，探索未来的无限可能。'], },
          { title: '异构聚合，技术创新', list: ['业界首次支持4种异构GPU千卡规模混合训练同一个大模型，异构协同训练效率超过98.5%，助力最终客户实现多种异构算力聚合，最大化提升异构GPU集群利用效率，引领国家标准和行业发展。', '目前，壁砺系列成功在华北、华东、华南等多地的智算中心与终端应用客户处落地，支持包括互联网、能源、金融、教育等众多应用场景。在AI大模型训练与推理场景中，壁仞科技也已经与众多合作伙伴开展落地合作，并完成了主流大模型的适配与调优。'], },
          { title: '智算一体机', list: ['基于壁仞科技 芯片产品与服务器整机厂商合作推出 大模型一体机，已在政务、能源、运营商、科技、金融等行业领域实现广泛应用，具体应用场景包括企业知识库助手、公文报告生成、智慧运维、智慧合规审查、代码助手等。软硬件深度融合，应用高效迁移，产品开箱即用，助力企业快速部署与落地。'], },
        ],
      }]
    };
  },
  components: {},
  computed: {
    sortObj() {
      return this.sortList.filter(itm => itm.key == this.sortType)[0] || {};
    }
  },
  methods: {
    changeTab(tab) {
      if (tab.key == this.tabIndex) return;
      this.tabIndex = tab.key;
      this.getData();
    },
    changeSort() {
      this.getData();
    },
    transformNumber(value, unit) {
      const param = {};
      const k = 10000;
      const sizes = ['', '万', '亿', '万亿'];
      let i;
      if (value < k) {
        param.value = value;
        param.unit = '';
      } else {
        i = Math.floor(Math.log(value) / Math.log(k));
        param.value = ((value / Math.pow(k, i))).toFixed(2);
        param.unit = sizes[i];
      }
      return param.value + ' ' + param.unit + unit;
    },
    getData() {
      this.loading = true;
      getDomesticCardData({
        type: this.tabIndex,
        category: this.sortType
      }).then(res => {
        this.loading = false;
        res = res.data;
        if (res.code == 0) {
          const data = (res.data || []).map((itm, index) => ({
            ...itm,
            rank: index + 1,
          }));
          if (data.length) {
            const first = data[0].count;
            const second = data[1].count;
            this.useShrink = first && second && first / second >= 10;
            data.map((item, index) => {
              item.id = Math.random();
              if (index == 0) {
                item.barWith = 100;
              } else {
                if (this.useShrink) {
                  item.barWith = second == 0 ? 0 : parseFloat((80 * (item.count / second)).toFixed(2));
                } else {
                  item.barWith = first == 0 ? 0 : parseFloat((100 * (item.count / first)).toFixed(2));
                }
              }
              if (first == 0) {
                item.barWith = 0;
              }
              item.countShow = this.transformNumber(item.count, this.sortObj.unit);
            })
            this.updateTime = formatDate(new Date(data[0].updated_unix * 1000), 'yyyy-MM-dd HH:mm:ss');
          } else { }
          this.tableData = data;
        } else {
          console.log(res);
        }
      }).catch(err => {
        this.loading = false;
        console.log(err);
      })
    },
  },
  created() {
    this.isOperator = window.IS_OPERATOR;
    if (!this.isOperator) this.sortList.splice(1, 1);
    this.getData();
  },
  mounted() {
    if (document.body.clientWidth <= 767) {
      this.useSmall = true;
    }
  },
  beforeDestroy() { },
};
</script>

<style scoped lang="less">
.top-field {
  text-align: center;
  padding-top: 42px;
  margin-bottom: 20px;

  .title {
    color: rgba(16, 16, 16, 1);
    font-size: 28px;
    margin-bottom: 24px;
  }

  .descr {
    color: rgba(136, 136, 136, 0.87);
    font-size: 14px;
  }
}

.sort-field {
  margin-bottom: 20px;

  .sort-conds {
    display: flex;
    align-items: center;
    justify-content: space-between;
    flex-wrap: wrap;

    .sort-tab-c {
      display: flex;
      margin-bottom: 16px;
      margin-right: 20px;

      .tab {
        width: 110px;
        height: 40px;
        border-color: rgba(0, 0, 0, 0.1);
        border-width: 1px;
        border-style: solid;
        color: rgba(16, 16, 16, 1);
        display: flex;
        align-items: center;
        justify-content: center;
        border-left: none;
        cursor: pointer;

        &:first-child {
          border-radius: 4px 0px 0px 4px;
          border-left: 1px solid rgba(0, 0, 0, 0.1);
        }

        &:last-child {
          border-radius: 0px 4px 4px 0px;
        }

        &.active {
          color: rgba(91, 185, 115, 1);
          border-color: rgba(33, 186, 69, 1);
          border-left: 1px solid rgba(33, 186, 69, 1);
        }
      }

    }

    .sort-type-c {
      margin-bottom: 16px;
    }
  }

  .table-container {
    .table {
      /deep/ .el-table__header {
        th {
          background: #f1f8ff;
          color: rgba(16, 16, 16, 1);
          font-size: 14px;

          span {
            font-weight: normal;
          }
        }
      }

      /deep/ .el-table__body {
        td {
          font-size: 14px;
        }
      }
    }

    .table-bar {
      width: 100%;
      display: flex;
      display: flex;
      align-items: center;
      min-width: 200px;

      .bar-c {
        flex: 1;
        display: flex;
        align-items: center;
        justify-content: flex-end;


        .bar {
          overflow: hidden;
          height: 8px;
          background: url("data:image/svg+xml;charset=utf8,%3Csvg%20xmlns%3D%22http%3A%2F%2Fwww.w3.org%2F2000%2Fsvg%22%20xmlns%3Axlink%3D%22http%3A%2F%2Fwww.w3.org%2F1999%2Fxlink%22%20version%3D%221.1%22%3E%3Cdefs%3E%3ClinearGradient%20id%3D%221%22%20x1%3D%220%22%20x2%3D%221%22%20y1%3D%220%22%20y2%3D%220%22%20gradientTransform%3D%22matrix(-0.9910000000000001%2C%20-0.014999999999999873%2C%200.000003944485025566073%2C%20-0.9910000000000001%2C%200.999%2C%200.125)%22%3E%3Cstop%20stop-color%3D%22%233291f8%22%20stop-opacity%3D%221%22%20offset%3D%220%22%3E%3C%2Fstop%3E%3Cstop%20stop-color%3D%22%23c9bcea%22%20stop-opacity%3D%221%22%20offset%3D%221%22%3E%3C%2Fstop%3E%3C%2FlinearGradient%3E%3C%2Fdefs%3E%3Crect%20width%3D%22100%25%22%20height%3D%22100%25%22%20fill%3D%22url(%231)%22%3E%3C%2Frect%3E%3C%2Fsvg%3E");

          div {
            float: right;
            height: 100%;
            background-color: white;
            width: 8px;
            margin-right: 8px;
            border-top: 1px solid #cce4fd;
            border-bottom: 1px solid #cce4fd;
            box-sizing: border-box;
          }
        }
      }

      .bar-value {
        width: 100px;
        margin-left: 16px;
      }
    }
  }
}

.partners-field {
  margin-top: 42px;
  text-align: center;

  .partners-title {
    color: rgba(16, 16, 16, 1);
    font-size: 18px;
    margin-bottom: 10px;
  }

  .partners-c {
    display: flex;
    align-items: center;
    justify-content: center;
    flex-wrap: wrap;

    .img-c {
      width: 170px;
      height: 80px;
      margin: 10px 16px;
      display: flex;
      align-items: center;
      justify-content: center;
      padding: 16px;

      img {
        max-height: 100%;
        max-width: 100%;
      }
    }
  }
}

.apply-field {
  margin-top: 12px;
  text-align: center;
  display: flex;
  align-items: center;
  justify-content: center;

  .apply-btn {
    height: 40px;
    border-radius: 5px;
    background: linear-gradient(30deg, rgba(59, 182, 254, 1) 13.4%, rgba(142, 76, 183, 1) 85.87%);
    color: rgba(255, 255, 255, 1);
    font-size: 14px;
    text-align: center;
    display: flex;
    align-items: center;
    justify-content: center;
    padding: 0 20px;
  }
}

.card-field {
  display: flex;
  flex-wrap: wrap;
  justify-content: center;
  margin-top: 16px;

  .card {
    width: 45%;
    min-width: 300px;
    max-width: 650px;
    border-radius: 10px;
    border-color: rgba(201, 188, 234, 0.4);
    border-width: 1px;
    border-style: solid;
    box-shadow: rgba(157, 197, 226, 0.2) 0px 5px 10px 0px;
    margin: 20px;
    padding: 28px;
    background: url("data:image/svg+xml;charset=utf8,%3Csvg%20xmlns%3D%22http%3A%2F%2Fwww.w3.org%2F2000%2Fsvg%22%20xmlns%3Axlink%3D%22http%3A%2F%2Fwww.w3.org%2F1999%2Fxlink%22%20version%3D%221.1%22%3E%3Cdefs%3E%3ClinearGradient%20id%3D%221%22%20x1%3D%220%22%20x2%3D%221%22%20y1%3D%220%22%20y2%3D%220%22%20gradientTransform%3D%22matrix(-0.41100000000000014%2C%201.129%2C%20-0.9330578512396693%2C%20-0.41100000000000014%2C%200.911%2C%20-0.129)%22%3E%3Cstop%20stop-color%3D%22%23f1ebff%22%20stop-opacity%3D%221%22%20offset%3D%220.005%22%3E%3C%2Fstop%3E%3Cstop%20stop-color%3D%22%23eef2ff%22%20stop-opacity%3D%221%22%20offset%3D%220.19%22%3E%3C%2Fstop%3E%3Cstop%20stop-color%3D%22%23ffffff%22%20stop-opacity%3D%221%22%20offset%3D%220.59%22%3E%3C%2Fstop%3E%3C%2FlinearGradient%3E%3C%2Fdefs%3E%3Crect%20width%3D%22100%25%22%20height%3D%22100%25%22%20fill%3D%22url(%231)%22%3E%3C%2Frect%3E%3C%2Fsvg%3E");

    .title-c {
      display: flex;
      align-items: center;
      justify-content: space-between;
      margin-bottom: 10px;

      .title {
        font-weight: 700;
        font-size: 16px;
        color: #101010;
      }

      .icon-c {
        min-width: 120px;
        height: 32px;
        display: flex;
        align-items: center;
        justify-content: flex-end;

        img {
          max-height: 100%;
        }
      }
    }

    .descr-c {
      line-height: 24px;
      color: rgba(16, 16, 16, 1);
      font-size: 14px;

      p {
        margin-bottom: 8px;
        line-height: 24px;
      }
    }

    .feature-c {
      font-size: 14px;

      .feature {
        margin-bottom: 8px;

        p {
          margin-bottom: 8px;
        }

        .title {
          font-weight: bold;
          margin-bottom: 6px;
        }
      }
    }

    .table-c {
      display: flex;
      align-items: center;
      justify-content: center;

      .table {
        width: 85%;

        .title {
          font-weight: bold;
        }

        .field-title {
          text-align: right;
        }

        th,
        td {
          border: 1px solid #EBEEF5;
          padding: 6px 8px;
          font-size: 14px;

          sup {
            font-size: 12px;
          }
        }

        &.average {

          th,
          td {
            width: 50%;
          }
        }
      }
    }

    .use-example-c {
      margin-top: 12px;

      .title {
        font-weight: bold;
        margin-bottom: 8px;
      }

      a {
        margin-bottom: 6px;
        line-height: 1.4285em;
        color: rgb(22, 132, 252);
        word-wrap: break-word;
      }
    }

    &.placeholder {
      box-shadow: none;
      background: none;
      border: none;
    }
  }
}

@media only screen and (max-width: 767px) {
  .sort-field {
    .sort-conds {
      margin-bottom: 8px;

      .sort-tab-c {
        margin-bottom: 8px;

        .tab {
          height: 32px;
          font-size: 12px;
        }
      }

      .sort-type-c {
        margin-bottom: 8px;
      }
    }
  }

  .partners-field {

    .partners-c {

      .img-c {
        width: 150px;
        height: 60px;
        margin: 10px 10px;
        padding: 12px;
      }
    }
  }

  .card-field {
    margin-top: 16px;

    .card {
      width: 100%;
      margin: 20px 0;
      padding: 20px;

      .title-c {

        .icon-c {
          min-width: 80px;
          height: 20px;
          display: flex;
          align-items: center;
          justify-content: flex-end;
        }
      }

      .table-c {
        display: flex;
        align-items: center;
        justify-content: center;

        .table {
          width: 100%;
        }
      }

      &.placeholder {
        display: none;
      }
    }
  }
}
</style>
