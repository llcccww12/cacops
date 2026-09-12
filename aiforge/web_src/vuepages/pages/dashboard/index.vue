<template>
  <div class="body-container">
    <div class="body-content">
      <div class="header-wrap">
        <div class="header-left-c">
          <img class="header-img" :src="avatarUrl"></img>
          <div class="header-context">
            <div class="header-title">
              Hi, <span>{{ userName }}</span> {{ $t('dashboard.welcomeTips') }}
            </div>
            <div class="header-bind">
              <div class="bind-item active" @click="toBindWx">{{ isBindWechatText }}</div>
            </div>
          </div>
        </div>
      </div>
      <div class="main-wrap">
        <div class="notices-section">
          <div class="header-b">
            <div class="b-title">{{ $t('dashboard.announcements') }}</div>
            <a class="b-desc" href="/home/notice" target="_blank">{{ $t('dashboard.moreAnnouncements') }}...</a>
          </div>
          <div class="content-b">
            <div class="swiper-container">
              <div class="swiper-wrapper">
                <a class="swiper-slide content-item" v-for="(item, index) in overviewConfig.noticeInfo.notices"
                  :key="`${item.title}-${index}`" :href="item.link">
                  <div class="date">{{ item.date }}</div>
                  <span class="context nowrap">{{ lang == 'en-US' ? item.title_en : item.title }}</span>
                </a>
              </div>
            </div>
          </div>
        </div>
        <div class="quick-section main-item">
          <div class="quick-t">
            <div class="icon-item">
              <img src="/img/home/dashboard-1.webp" />
            </div>
            <span>{{ $t('dashboard.quickAccess') }}</span>
          </div>
          <div class="quick-c">
            <a class="quick-c-item" href="/cloudbrains/create">
              <img src="/img/svg-icon/computingpower.svg" />
              <div class="title">{{ $t('dashboard.createTask') }}</div>
            </a>
            <a class="quick-c-item" href="/repo/create">
              <img src="/img/svg-icon/code.svg" />
              <div class="title">{{ $t('dashboard.createProject') }}</div>
            </a>
            <a class="quick-c-item" href="/guide/create_dataset">
              <img src="/img/svg-icon/dataset.svg" />
              <div class="title  now">{{ $t('datasetObj.createDataset') }}</div>
            </a>
            <a class="quick-c-item" href="/guide/create_model">
              <img src="/img/svg-icon/model.svg" />
              <div class="title">{{ $t('modelManage.createModel') }}</div>
            </a>
          </div>
        </div>
        <div class="task-section main-item">
          <div class="item-header-w">
            <a class="quick-t" href="/cloudbrains">
              <div class="icon-item">
                <img src="/img/home/dashboard-2.webp" />
              </div>
              <span>{{ $t('notebook.sameTaskTips6') }}</span>
            </a>
          </div>
          <div class="right-b-item">
            <div class="b-item">
              <div class="item">
                <span class="label">{{ $t('dashboard.myTaskCount') }}</span>
                <div class="value"><span>{{ overviewData.aiTaskInfo.totalAiTasks }}</span> {{ $t('dashboard.unitTask') }}
                </div>
              </div>
            </div>
            <div class="b-item">
              <div class="item">
                <span class="label">{{ $t('dashboard.runningTasks') }}</span>
                <div class="value"><span>{{ overviewData.aiTaskInfo.runningAiTasks }}</span> {{ $t('dashboard.unitTask') }}
                </div>
              </div>

            </div>
            <div class="b-item">
              <div class="item">
                <span class="label">{{ $t('dashboard.totalGPUHourd') }}</span>
                <div class="value"><span
                    style="font-size: 18px;">{{ overviewData.aiTaskInfo.allCardDurationCount }}</span>
                  {{ $t('dashboard.runtimeCards') }}</div>
              </div>

            </div>
          </div>
        </div>
        <div class="task-section main-item">
          <div class="item-header-w">
            <a class="quick-t" href="/ai_task_tmpl/list_my">
              <div class="icon-item">
                <img src="/img/home/dashboard-2.webp" />
              </div>
              <span>{{ $t('taskTmplObj.taskTmpl') }}</span>
            </a>
          </div>
          <div class="right-b-item">
            <div class="b-item" style="justify-content: center;">
              <div class="item">
                <span class="label">{{ $t('dashboard.templateNum') }}</span>
                <div class="value" style="text-align: center;"><span
                    style="font-size: 18px;">{{ overviewData.aiTaskInfo.templateCount }}</span>
                  {{ $t('dashboard.unitTask') }}</div>
              </div>
            </div>
            <div class="b-item">
              <div class="item">
                <span class="label">{{ $t('dashboard.accumulatedRuns') }}</span>
                <div class="value"><span style="font-size: 18px;">{{ overviewData.aiTaskInfo.templateUseCount }}</span>
                  {{ $t('dashboard.unitTemplate') }}</div>
              </div>
            </div>
          </div>
        </div>
        <div class="task-section main-item">
          <div class="item-header-w">
            <a class="quick-t" href="/reward/point">
              <div class="icon-item">
                <img src="/img/home/dashboard-4.webp" />
              </div>
              <span>{{ $t('dashboard.credits') }}</span>
            </a>
          </div>
          <div class="right-b-item">
            <div class="b-item">
              <div class="item">
                <span class="label">{{ $t('dashboard.currentAvailable1') }}</span>
                <div class="value"><span>{{ overviewData.pointInfo.balance }}</span> {{ $t('dashboard.unitPoints') }}</div>
              </div>
            </div>
            <div class="b-item">
              <div class="item">
                <span class="label">{{ $t('dashboard.totalGained1') }}</span>
                <div class="value"><span style="font-size: 18px;">{{ overviewData.pointInfo.totalEarned }}</span>
                  {{ $t('dashboard.unitPoints') }}</div>
              </div>
            </div>
            <div class="b-item">
              <div class="item">
                <span class="label">{{ $t('dashboard.totalConsumed1') }}</span>
                <div class="value"><span style="font-size: 18px;">{{ overviewData.pointInfo.totalConsumed }}</span>
                  {{ $t('dashboard.unitPoints') }}</div>
              </div>
            </div>
          </div>
        </div>
        <div class="repo-data-model-section">
          <div class="section main-item">
            <div class="item-header-w">
              <a class="quick-t" href="/repositories">
                <div class="icon-item">
                  <img src="/img/home/dashboard-2.webp" />
                </div>
                <span>{{ $t('repos.repos') }}</span>
              </a>
            </div>
            <div class="chart-container">
              <div ref="chartRef1" class="chart"></div>
            </div>
          </div>
          <div class="section main-item">
            <div class="item-header-w">
              <a class="quick-t" href="/explore/datasets_my">
                <div class="icon-item">
                  <img src="/img/home/dashboard-2.webp" />
                </div>
                <span>{{ $t('dataset') }}</span>
              </a>
            </div>
            <div class="chart-container">
              <div ref="chartRef2" class="chart"></div>
            </div>
          </div>
          <div class="section main-item">
            <div class="item-header-w">
              <a class="quick-t" href="/explore/models_my">
                <div class="icon-item">
                  <img src="/img/home/dashboard-2.webp" />
                </div>
                <span>{{ $t('repos.model') }}</span>
              </a>
            </div>
            <div class="chart-container">
              <div ref="chartRef3" class="chart"></div>
            </div>
          </div>
        </div>
        <div class="storage-section main-item">
          <div class="item-header-w">
            <a class="quick-t" href="/storages">
              <div class="icon-item">
                <img src="/img/home/dashboard-3.webp" />
              </div>
              <span>{{ $t('storage.quota') }}</span>
            </a>
          </div>
          <div class="item-box-t">
            <div class="progress-container">
              <svg>
                <!-- 蓝色背景，只显示上半圆（180度） -->
                <circle cx="70" cy="60" r="50" stroke='rgba(19,194,194,1)' stroke-width="10" fill="none"
                  stroke-dasharray="157 314" stroke-dashoffset="0" stroke-linecap="round"
                  transform="rotate(-180 70 60)"> <!-- 旋转让起点在顶部中间 -->
                </circle>
                <text x="70" y="60" text-anchor="middle" font-size="20" fill="rgba(16,16,16,0.7);"
                  class="progress-percent">
                  {{ remainingPercentage }}
                </text>
                <!-- 灰色进度，0%时长度为0，100%时长度为157，覆盖蓝色上半圆 -->
                <circle cx="70" cy="60" r="50" stroke='rgba(221,216,216,1)' stroke-width="11" fill="none"
                  :stroke-dasharray="progressDasharray" stroke-dashoffset="0" stroke-linecap="round"
                  transform="rotate(-180 70 60)" :style="{ opacity: storageUsagePercentage > 0 ? 1 : 0 }"
                  class="progress-circle">
                </circle>
              </svg>
            </div>
            <div class="remain-container">
              <span class="label">{{ $t('dashboard.remainingQuota') }}</span>
              <div class="value"><span>{{ formattedRemaining.value }}</span> {{ formattedRemaining.label }}</div>
            </div>
          </div>
          <div class="right-b-item" style="margin-top: 20px;">
            <div class="b-item" style="justify-content: center;">
              <div class="item">
                <span class="label">{{ $t('dashboard.datasetStorage') }}</span>
                <div class="value"><span>{{ formattedData.value }} </span>{{ formattedData.label }}</div>
              </div>
            </div>
            <div class="b-item">
              <div class="item">
                <span class="label">{{ $t('dashboard.modelStorage') }}</span>
                <div class="value"><span>{{ formattedModel.value }} </span>{{ formattedModel.label }}</div>
              </div>
            </div>
          </div>
        </div>

      </div>
    </div>
  </div>
</template>

<script>
import { getOverviewData, getOverviewConfig } from '~/apis/modules/dashboard';
import { getPromoteDataset } from "~/apis/modules/dataset";
import { TmplTaskTypes, TmplComputerResouces } from '~/pages/aitasktmpl/tools';
import { getListValueWithKey } from '~/utils';
import { lang } from '~/langs';
import * as echarts from "echarts";
const UNITS = ['Bytes', 'KiB', 'MiB', 'GiB', 'TiB'];
export default {
  name: 'AiforgeIndex',
  data() {
    return {
      loading: false,
      loadingRight: true,
      swiperHandler: null,
      swiperList: [],
      overviewData: {
        bindingInfo: {
          isBindWechat: false
        },
        aiTaskInfo: {
          totalAiTasks: 0,
          runningAiTasks: 0,
          allCardDurationCount: 0,
          templateCount: 0,
          templateUseCount: 0,
        },
        storageInfo: {
          storageLimit: 0,
          datasetCount: 0,
          datasetUsedStorage: 0,
          modelCount: 0,
          modelUsedStorage: 0,
          usedStorage: 0,
          remainingStorage: 0
        },
        pointInfo: {
          totalEarned: 0,
          balance: 0,
          totalConsumed: 0
        },
        repositoryInfo: {
          repositoryCount: 0,
          selfBuiltCount: 0,
          forkedCount: 0,
          mirroredCount: 0,
          collaborationCount: 0,
        },
        datasetInfo: {
          datasetCount: 0,
          publicCount: 0,
          privateCount: 0
        },
        modelInfo: {
          modelCount: 0,
          publicCount: 0,
          privateCount: 0
        },
      },
      overviewConfig: {
        noticeInfo: {
          notices: [],
          commitId: '',
        },
        activityImageInfo: {
          imageUrl: '',
          imageLink: '',
        },

      },
      lang: lang,
      avatarUrl: '',
      userName: '',
      chartInstances: {
        chart1: null,
        chart2: null,
        chart3: null
      },
      chartConfigs: {
        repositoryChart: {
          colors: ['rgb(80,135,236)', 'rgb(104,187,196)', 'rgb(88,165,92)', 'rgb(242,189,66)'],
          dataKeys: ['selfBuiltCount', 'forkedCount', 'mirroredCount', 'collaborationCount'],
          legendNames: [this.$t('repos.source'), this.$t('repos.fork'), this.$t('repos.mirrors'), this.$t('repos.collaborative')],
          totalKey: 'repositoryCount',
          unit: this.$t('dashboard.unitTask')
        },
        datasetChart: {
          colors: ['rgb(96,113,253)', 'rgb(89,219,255)'],
          dataKeys: ['publicCount', 'privateCount'],
          legendNames: [this.$t('modelManage.modelAccessPublic'), this.$t('modelManage.modelAccessPrivate')],
          totalKey: 'datasetCount',
          unit: this.$t('dashboard.unitTask')
        },
        modelChart: {

          colors: ['rgb(242,189,66)', 'rgb(88,165,92)'],
          dataKeys: ['publicCount', 'privateCount'],
          legendNames: [this.$t('modelManage.modelAccessPublic'), this.$t('modelManage.modelAccessPrivate')],
          totalKey: 'modelCount',
          unit: this.$t('dashboard.unitTask')
        }
      }
    };
  },
  computed: {
    isBindWechatText() {
      const { isBindWechat } = this.overviewData.bindingInfo;
      if (isBindWechat) return this.$t('dashboard.weChatLinked');
      return this.$t('dashboard.weChatNotLinked');
    },
    // 单位转换方法
    formattedData() {
      const { datasetUsedStorage } = this.overviewData.storageInfo;
      return this.formatBytes(datasetUsedStorage);
    },
    formattedModel() {
      const { modelUsedStorage } = this.overviewData.storageInfo;
      return this.formatBytes(modelUsedStorage);
    },
    // 格式化的剩余空间显示（确保不超过 totalStorage）
    formattedRemaining() {
      const { remainingStorage } = this.overviewData.storageInfo;
      return this.formatBytes(Math.max(remainingStorage, 0));
    },
    // 存储使用率百分比（0-100）
    storageUsagePercentage() {
      if (!this.overviewData.storageInfo.storageLimit) return 0;
      return parseFloat((this.overviewData.storageInfo.usedStorage /
        this.overviewData.storageInfo.storageLimit * 100).toFixed(2));
    },
    storageRemainPercentage() {
      if (!this.overviewData.storageInfo.storageLimit) return 0;
      return parseFloat((this.overviewData.storageInfo.remainingStorage /
        this.overviewData.storageInfo.storageLimit * 100).toFixed(2));
    },
    semiCircleCircumference() {
      const radius = 50; // 圆的半径
      return Math.PI * radius; // πr（半圆周长）
    },
    // 计算进度条的实际长度
    progressDasharray() {
      const semiCircumference = this.semiCircleCircumference;
      const progressLength = (semiCircumference * Math.min(this.storageUsagePercentage, 100)) / 100;
      return `${progressLength} 314`;
    },
    // 剩余空间百分比（带精度处理）
    remainingPercentage() {
      const { remainingStorage, storageLimit, datasetUsedStorage, modelUsedStorage, usedStorage } = this.overviewData.storageInfo;
      if (remainingStorage <= 0) {
        return '0%'; // 超限时强制显示0%
      }
      if (remainingStorage === storageLimit) return '100%'
      if ((usedStorage / storageLimit) * 100 < 0.01) return '99.99%'
      if ((remainingStorage / storageLimit) * 100 < 0.01) return '≤0.01%'
      const dataPct = datasetUsedStorage / storageLimit * 100;
      const modelPct = modelUsedStorage / storageLimit * 100;
      const roundedDataStr = this.toPrecision(dataPct, 2);
      const roundedModelStr = this.toPrecision(modelPct, 2);
      // 处理 "<0.01" 的情况，将其视为 0
      const roundedData = roundedDataStr === '≤0.01' ? 0 : Number(roundedDataStr);
      const roundedModel = roundedModelStr === '≤0.01' ? 0 : Number(roundedModelStr);
      console.log('数据占用百分比:', roundedData, roundedModel);
      const roundedRemaining = 100 - roundedData - roundedModel
      return `${roundedRemaining.toFixed(2)}%`;
    },
    textColor() {
      // 根据进度改变文本颜色
      if (this.progress < 30) return '#FF6B6B';
      if (this.progress < 70) return '#FFA726';
      return '#4CAF50';
    }
  },
  methods: {
    toBindWx() {
      if (!this.overviewData.bindingInfo.isBindWechat) {
        window.location.href = '/authentication/wechat/bind?redirect_to=/dashboard';
      }
    },
    async getOverview() {
      try {
        this.loading = true;
        const response = await getOverviewData();
        const res = response.data;
        if (res.Code === 0) {
          const data = res.Data;
          console.log('获取数据成功:', data);
          this.formatAndSetOverviewData(data)
        } else {
          this.$message.error(res.Msg || '获取数据失败');
        }
      } catch (error) {
        console.error('获取概览数据失败:', error);
        this.$message.error(error.message || '网络错误，请稍后重试');
      } finally {
        this.loading = false;
      }
    },
    // 数据处理方法
    // 数据处理方法
    formatAndSetOverviewData(data) {
      if (!data) return;
      // 转换数据格式（保持响应式）
      Object.assign(this.overviewData.bindingInfo, {
        isBindWechat: data.binding_info?.is_bind_wechat || false
      });

      Object.assign(this.overviewData.aiTaskInfo, {
        totalAiTasks: data.ai_task_info?.total_ai_tasks || 0,
        runningAiTasks: data.ai_task_info?.running_ai_tasks || 0,
        allCardDurationCount: data.ai_task_info?.all_card_duration_count || 0,
        templateCount: data.ai_task_info?.template_count || 0,
        templateUseCount: data.ai_task_info?.template_use_count || 0
      });

      Object.assign(this.overviewData.storageInfo, {
        storageLimit: data.storage_info?.storage_limit || 0,
        datasetCount: data.storage_info?.dataset_count || 0,
        datasetUsedStorage: data.storage_info?.dataset_used_storage || 0,
        modelCount: data.storage_info?.model_count || 0,
        modelUsedStorage: data.storage_info?.model_used_storage || 0,
        usedStorage: data.storage_info?.used_storage || 0,
        remainingStorage: data.storage_info?.remaining_storage || 0
      });

      Object.assign(this.overviewData.pointInfo, {
        totalEarned: this.roundToTwo(data.point_info?.total_earned) || 0,
        balance: this.roundToTwo(data.point_info?.balance) || 0,
        totalConsumed: this.roundToTwo(data.point_info?.total_consumed) || 0
      });
      Object.assign(this.overviewData.repositoryInfo, {
        repositoryCount: data.repository_info?.repository_count || 0,
        selfBuiltCount: data.repository_info?.self_built_count || 0,
        forkedCount: data.repository_info?.forked_count || 0,
        mirroredCount: data.repository_info?.mirrored_count || 0,
        collaborationCount: data.repository_info?.collaboration_count || 0
      });
      Object.assign(this.overviewData.datasetInfo, {
        datasetCount: data.dataset_info?.dataset_count || 0,
        publicCount: data.dataset_info?.public_count || 0,
        privateCount: data.dataset_info?.private_count || 0
      });
      Object.assign(this.overviewData.modelInfo, {
        modelCount: data.model_info?.model_count || 0,
        publicCount: data.model_info?.public_count || 0,
        privateCount: data.model_info?.private_count || 0
      });
      console.log('this.overviewData:', this.overviewData);
      this.$nextTick(() => {
        this.renderAllCharts()
      })

    },

    renderAllCharts() {
      // 渲染仓库图表
      this.renderChart('repositoryChart', this.overviewData.repositoryInfo, 'chartRef1')

      // 渲染数据集图表
      this.renderChart('datasetChart', this.overviewData.datasetInfo, 'chartRef2')

      // 渲染模型图表
      this.renderChart('modelChart', this.overviewData.modelInfo, 'chartRef3')
    },
    renderChart(chartType, data, refName) {
      const chartRef = this.$refs[refName]
      if (!chartRef) return

      // 销毁旧实例
      if (this.chartInstances[chartType]) {
        this.chartInstances[chartType].dispose()
      }

      // 创建新实例
      const chartInstance = echarts.init(chartRef)
      this.chartInstances[chartType] = chartInstance

      // 获取配置
      const config = this.chartConfigs[chartType]

      // 准备图表数据
      const chartData = this.prepareChartData(data, config)

      // 计算总数
      const total = data[config.totalKey] || 0

      // 生成配置项
      const options = this.generateOptions(chartData, total, config)

      // 设置配置
      chartInstance.setOption(options)

      // 立即调整大小
      chartInstance.resize()
    },

    prepareChartData(data, config) {
      return config.dataKeys.map((key, index) => ({
        name: config.legendNames[index],
        value: data[key] || 0,
        itemStyle: {
          color: config.colors[index % config.colors.length]
        }
      }))
    },

    generateOptions(chartData, total, config) {
      const graphFloatLetf = (totalNum) => {
        if (totalNum <= 9) {
          return '46%'
        } else if (totalNum <= 99) {
          return '42%'
        } else if (totalNum <= 999) {
          return '36.5%'
        } else if (totalNum < 9999) {
          return '32%'
        } else if (totalNum < 99999) {
          return '30%'
        } else if (totalNum < 999999) {
          return '24%'
        } else if (totalNum < 9999999) {
          return '18%'
        } else if (totalNum < 99999999) {
          return '14%'
        }
      }
      return {
        tooltip: {
          trigger: 'item',
          formatter: (params) => {
            return `${params.name}: ${params.value}${config.unit}`
          }
        },
        series: [
          {
            name: '数量统计',
            type: 'pie',
            radius: ['65%', '80%'],
            center: ['50%', '50%'],
            avoidLabelOverlap: false,
            label: {
              show: false
            },
            labelLine: {
              show: false
            },
            data: chartData,

          }
        ],
        graphic: [
          {
            type: 'text',
            left: `${graphFloatLetf(total)}`,
            top: '38%',  // 数字位置
            style: {
              text: total.toString(),
              textAlign: 'center',
              fill: '#101010',
              fontSize: 18,
              fontWeight: 'bold',
              lineHeight: 18
            }
          },
          {
            type: 'text',
            left: '44%',
            top: '58%',  // 单位位置
            style: {
              text: config.unit,
              textAlign: 'center',
              fill: 'rgba(16,16,16,0.5)',
              fontSize: 16
            }
          }
        ]
      }
    },
    handleResize() {
      // 防抖处理
      if (this.resizeTimer) {
        clearTimeout(this.resizeTimer)
      }
      this.resizeTimer = setTimeout(() => {
        Object.values(this.chartInstances).forEach(chart => {
          if (chart && chart.resize) {
            try {
              chart.resize()
            } catch (error) {
              console.warn('图表resize失败:', error)
            }
          }
        })
      }, 200)
    },
    formatAndSetOverviewConfig(data) {
      // 转换数据格式（保持响应式）
      Object.assign(this.overviewConfig.noticeInfo, {
        notices: data.notice_info?.notices || [],
        commitId: data.notice_info?.commit_id || ''
      });
      Object.assign(this.overviewConfig.activityImageInfo, {
        imageLink: data.activity_image_info?.image_link || '',
        imageUrl: data.activity_image_info?.image_url || ''
      });
    },
    roundToTwo(num) {
      return parseFloat(num.toFixed(2));
    },
    toPrecision(value, decimals = 2) {
      if (typeof value !== "number" || isNaN(value)) return NaN;
      if (value > 0 && value < 0.01) {
        return '≤0.01';  // 极小值显示 ≤0.01%
      }
      // 四舍五入到指定位数
      const roundedValue = Number(value.toFixed(decimals));
      // 检查四舍五入后是否为100.00且原始值小于100
      if (roundedValue === 100.00 && value < 100) {
        const factor = Math.pow(10, decimals);
        const truncated = Math.floor(value * factor) / factor;
        return truncated.toFixed(decimals);
      } else {
        return roundedValue.toFixed(decimals);
      }
    },
    // 智能单位格式化
    formatBytes(bytes) {
      let unitIndex = 0;
      let value = bytes;
      let totalInSameUnit = this.overviewData.storageInfo.storageLimit;  // 总存储的初始单位（字节）
      while (value >= 1024 && unitIndex < UNITS.length - 1) {
        value /= 1024;
        totalInSameUnit /= 1024;
        unitIndex++;
      }
      // 动态判断是否接近总存储的临界值
      const decimals = 2;
      const roundedValue = Number(value?.toFixed(decimals)) || 0;
      const totalRounded = Number(totalInSameUnit?.toFixed(decimals)) || 0;
      // 条件1：四舍五入后等于总存储的显示值，但实际值小于总存储
      // 条件2：四舍五入后等于1024，但实际值小于1024
      if (
        (roundedValue === totalRounded && value < totalInSameUnit) ||
        (roundedValue === 1024.00 && value < 1024)
      ) {
        const factor = Math.pow(10, decimals);
        const truncated = Math.floor(value * factor) / factor;
        return { value: truncated.toFixed(decimals), label: UNITS[unitIndex] };
      } else {
        return { value: parseFloat(roundedValue.toFixed(decimals)), label: UNITS[unitIndex] };
      }

    },
    async getOverviewPromote() {
      this.loadingRight = true
      try {
        const response = await getOverviewConfig();
        const res = response.data;
        if (res.Code === 0) {
          const data = res.Data;
          console.log('获取数据成功:', data);
          this.formatAndSetOverviewConfig(data)
        } else {
          this.$message.error(res.Msg || '获取数据失败');
        }
      } catch (error) {
        console.error('获取推广数据失败:', error);
        this.$message.error(error.message || '网络错误，请稍后重试');
      } finally {
        this.loadingRight = false
      }
    },
    async initData() {
      try {
        await Promise.all([
          this.getOverview(),
          this.getOverviewPromote(),
        ]);
      } catch (error) {
        console.error('获取概览数据失败:', error);
        this.$message.error(error.message || '网络错误，请稍后重试');
      }

      this.$nextTick(() => {
        this.initSwiper(true);
      });
    },
    initSwiper(loop = true) {
      // 检查 Swiper 容器是否存在
      if (!document.querySelector('.swiper-container')) {
        console.warn('Swiper container not found');
        return;
      }

      // 销毁旧的 Swiper 实例（防止重复初始化）
      if (this.swiperHandler && this.swiperHandler.destroy) {
        this.swiperHandler.destroy(true, true);
      }

      this.swiperHandler = new Swiper(".swiper-container", {
        slidesPerView: 1,
        direction: 'vertical',  // 改为垂直方向
        loop: loop && this.swiperList.length > 1, // 只有一张图片时不需要 loop
        spaceBetween: 0,
        autoplay: {
          delay: 2500,
          disableOnInteraction: false,
        },
      });
    }
  },
  mounted() {
    const avatarElement = document.getElementById('_user_avatar');
    if (avatarElement) {
      const avatarUrl = avatarElement.getAttribute('data-avatar');
      const userName = avatarElement.getAttribute('data-name');
      console.log('Avatar URL:', avatarUrl);
      // 可以存储到 Vue 实例中
      this.avatarUrl = avatarUrl;
      this.userName = userName;
    }
    this.initData();
  },
};
</script>

<style lang="less" scoped>
.body-container {
  padding: 28px 14px 40px 14px;
  height: 100%;
  background: linear-gradient(180deg, rgba(0, 0, 255, 0.05) 0%, rgba(249, 249, 249, 1.00) 100%);

  .body-content {
    display: flex;
    flex-direction: column;
    height: 100%;

    .header-wrap {
      margin-bottom: 24px;
      display: flex;
      align-items: center;
      justify-content: space-between;
      height: 64px;

      .header-left-c {
        display: flex;

        .header-img {
          width: 48px;
          height: 48px;
          border-radius: 40px;
        }

        .header-context {
          display: flex;
          flex-direction: column;
          justify-content: space-between;
          margin-left: 17px;

          .header-title {
            line-height: 28px;
            color: rgba(16, 16, 16, 1);
            font-size: 20px;
            font-family: Arial;

            span {
              font-weight: 700;
            }
          }

          .header-bind {
            display: flex;
            margin-top: 8px;
            gap: 5px;

            .bind-item {
              border: 1px solid rgba(187, 187, 187, 1);
              border-radius: 5px;
              line-height: 28px;
              height: 28px;
              color: rgba(16, 16, 16, 1);
              padding: 0 10px;
              cursor: pointer;

              &.active {
                color: #FA8C16;
                border-color: #FA8C16;
              }
            }
          }
        }
      }

      .header-right-c {
        color: rgba(0, 92, 255, 1);
        border: 1px solid rgba(0, 92, 255, 1);
        height: 40px;
        display: flex;
        align-items: center;
        padding: 0 15px;

        svg {
          margin-right: 8px;
        }
      }
    }

    .main-wrap {
      display: flex;
      flex-direction: column;
      gap: 20px;

      .main-item {
        background-color: rgba(255, 255, 255, 1);
        border-radius: 10px;
        border: 1px solid rgba(255, 255, 255, 1);
        padding: 16px 16px 20px 24px;
        display: flex;
        flex-direction: column;
      }

      .notices-section {
        width: 100%;
        margin-bottom: 4px;

        .header-b {
          display: flex;
          justify-content: space-between;
          align-items: flex-end;
          margin-bottom: 12px;

          .b-title {
            height: 30px;
            text-align: center;
            line-height: 30px;
            border-radius: 20px;
            background-color: rgba(255, 150, 0, 1);
            padding: 0 20px;
            color: rgba(255, 255, 255, 1);
          }

          .b-desc {
            line-height: 28px;
            color: rgb(0, 54, 255);
            border-bottom: 1px solid rgba(0, 54, 255, 1);
          }
        }

        .content-b {
          height: 23px;
          overflow: hidden;

          .swiper-container {
            height: 100%; // 确保容器有高度

            .swiper-wrapper {
              .content-item {
                display: flex;
                align-items: center;
                height: auto !important; // 关键：让高度自适应内容

                .date {
                  flex-shrink: 0;
                  height: 23px;
                  line-height: 23px;
                  border-radius: 5px;
                  background-color: rgba(228, 241, 255, 1);
                  color: rgba(0, 98, 255, 1);
                  padding: 0 10px;
                }

                .context {
                  margin: 0 8px;
                  color: rgba(16, 16, 16, 0.7);
                }

              }
            }
          }

        }
      }

      .quick-section {
        height: 172px;
        padding: 20px 24px;

        .quick-c {
          display: flex;
          flex-wrap: wrap;
          flex: 1;
          margin-top: 22px;

          .quick-c-item {
            width: 50%;
            display: flex;
            align-items: center;
            font-size: 16px;

            img {
              width: 32px;
              height: 32px;
            }

            .title {
              color: rgba(16, 16, 16, 1);
              margin-left: 10px;
            }
          }
        }
      }

      .task-section {
        height: 158px;
      }

      .repo-data-model-section {
        display: flex;
        gap: 12px;
        height: 180px;

        .section {
          width: 100%;
          flex: 1;
          padding: 0 !important;

          .item-header-w {
            margin: 0;
            padding: 16px 16px 0px 24px;
          }
        }

        .chart-container {
          flex: 1;

          .chart {
            width: 100%;
            height: 100%;
          }
        }
      }

      .storage-section {
        height: 234px;

        .item-box-t {
          display: flex;

          .progress-container {
            width: 140px;
            height: 70px;
            margin-left: 10px;

            .progress-circle {
              transition: opacity 0.3s ease;
            }
          }

          .remain-container {
            flex: 1;
            height: 100%;
            display: flex;
            flex-direction: column;
            justify-content: space-between;
            margin-left: 10px;

            .label {
              font-size: 14px;
              line-height: 19px;
              color: #101010;
            }

            .value {
              color: rgba(16, 16, 16, 0.5);
              font-size: 16px;
              margin-bottom: 8px;

              span {
                color: rgba(16, 16, 16, 1);
                font-size: 20px;
                font-weight: 700;
              }
            }
          }
        }
      }

      .quick-t {
        display: flex;
        align-items: center;
        font-size: 16px;
        color: rgba(96, 113, 253, 1);

        .icon-item {
          width: 24px;
          height: 24px;
          border-radius: 100%;
          display: flex;
          align-items: center;
          justify-content: center;

          img {
            width: 100%;
            height: 100%;
          }
        }

        span {
          font-weight: 700;
          margin-left: 6px;
          font-size: 16px;
        }
      }

      .item-header-w {
        display: flex;
        justify-content: space-between;
        align-items: flex-end;
        margin-bottom: 20px;

        i {
          font-size: 16px;
          color: rgba(16, 16, 16, 0.5);
        }
      }

      .right-b-item {
        display: flex;
        height: 65px;
        width: 100%;

        .b-item {
          flex: 1;
          display: flex;
          justify-content: center;
          align-items: center;
          height: 100%;
          border-left: 1px solid rgba(16, 16, 16, 0.15);

          &:first-child {
            justify-content: flex-start;
            border-left: none;

            .item {
              .value {
                text-align: left;
              }
            }
          }

          .item {
            display: flex;
            flex-direction: column;
            justify-content: space-between;
            height: 100%;

            .label {
              font-size: 14px;
              line-height: 19px;
              color: #101010;
              text-align: center;
            }

            .value {
              color: rgba(16, 16, 16, 0.5);
              font-size: 16px;
              text-align: center;

              span {
                color: rgba(16, 16, 16, 1);
                font-size: 20px;
                font-weight: 700;
              }
            }

          }

          svg {
            margin-left: 30px;
          }
        }
      }
    }
  }

}
</style>