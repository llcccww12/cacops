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
        <div class="header-right-c">
          <svg width="16" height="16" viewBox="0 0 48 48" fill="none" xmlns="http://www.w3.org/2000/svg">
            <path d="M6 10V22H38L44 16L38 10H6Z" fill="none" stroke="#005cff" stroke-width="4"
              stroke-linejoin="round" />
            <path d="M23 22V44" stroke="#005cff" stroke-width="4" stroke-linecap="round" stroke-linejoin="round" />
            <path d="M23 4V10" stroke="#005cff" stroke-width="4" stroke-linecap="round" stroke-linejoin="round" />
            <path d="M18 44H28" stroke="#005cff" stroke-width="4" stroke-linecap="round" stroke-linejoin="round" />
          </svg>
          <a target="blank" href="/docs/index.html">{{ $t('dashboard.newUserGuide') }}</a>
        </div>
      </div>
      <div class="main-wrap">
        <div class="main-left-c">
          <div class="quick-section main-item">
            <div class="quick-t">
              <div class="icon-item">
                <img src="/img/home/dashboard-1.webp" />
              </div>
              <span>{{ $t('dashboard.quickAccess') }}</span>
            </div>
            <div class="quick-c">
              <a class="quick-c-item" href="/cloudbrains/create">
                <div class="icon-w">
                  <img src="/img/svg-icon/computingpower.svg" />
                </div>
                <div class="content-w">
                  <div class="title">{{ $t('dashboard.createTask') }}</div>
                  <div class="desc" :title="$t('dashboard.taskDesc')">{{ $t('dashboard.taskDesc') }}</div>
                </div>
              </a>
              <a class="quick-c-item" href="/repo/create">
                <div class="icon-w">
                  <img src="/img/svg-icon/code.svg" />
                </div>
                <div class="content-w">
                  <div class="title">{{ $t('dashboard.createProject') }}</div>
                  <div class="desc" :title="$t('dashboard.createRepoDesc')">{{ $t('dashboard.createRepoDesc') }}</div>
                </div>
              </a>
              <a class="quick-c-item" href="/guide/create_dataset">
                <div class="icon-w">
                  <img src="/img/svg-icon/dataset.svg" />
                </div>
                <div class="content-w">
                  <div class="title  now">{{ $t('datasetObj.createDataset') }}</div>
                  <div class="desc" :title="$t('dashboard.uploadDatasetDesc')">{{ $t('dashboard.uploadDatasetDesc') }}
                  </div>
                </div>
              </a>
              <a class="quick-c-item" href="/guide/create_model">
                <div class="icon-w">
                  <img src="/img/svg-icon/model.svg" />
                </div>
                <div class="content-w">
                  <div class="title">{{ $t('modelManage.createModel') }}</div>
                  <div class="desc" :title="$t('dashboard.uploadModelDesc')">{{ $t('dashboard.uploadModelDesc') }}</div>
                </div>
              </a>
            </div>
            <div class="quick-f">
              <span style="flex-shrink: 0;">{{ $t('dashboard.oneClickCreate') }}</span>
              <div class="quick-f-item-wrap" ref="tagsContainer" v-if="list && list.length">
                <a class="quick-f-item" v-for="item in visibleTags" :key="item.ID" :href="item.link" :class="{
                  'truncated': isLastTagTruncated && index === visibleTags.length - 1
                }">
                  <svg style="flex-shrink: 0;" width="12" height="12" viewBox="0 0 48 48" fill="none"
                    xmlns="http://www.w3.org/2000/svg">
                    <path d="M28 6H42V20" stroke="#0066ff" stroke-width="4" stroke-linecap="round"
                      stroke-linejoin="round" />
                    <path
                      d="M42 29.4737V39C42 40.6569 40.6569 42 39 42H9C7.34315 42 6 40.6569 6 39V9C6 7.34315 7.34315 6 9 6L18 6"
                      stroke="#0066ff" stroke-width="4" stroke-linecap="round" stroke-linejoin="round" />
                    <path d="M25.7998 22.1999L41.0998 6.8999" stroke="#0066ff" stroke-width="4" stroke-linecap="round"
                      stroke-linejoin="round" />
                  </svg>
                  <span class="nowrap" :title="item.Alias">{{ item.Alias }}</span>
                </a>
              </div>

              <a target="_blank" href="/ai_task_tmpl/list" class="quick-f-more">{{ $t('cloudbrainObj.more') }}...</a>
            </div>
          </div>
          <template v-loading="loading">
            <div class="task-template-section">
              <div class="task-section main-item">
                <div class="item-header-w">
                  <a class="quick-t" href="/cloudbrains">
                    <div class="icon-item">
                      <img src="/img/home/dashboard-2.webp" />
                    </div>
                    <span>{{ $t('notebook.sameTaskTips6') }}</span>
                  </a>
                </div>
                <div class="item-box-w">
                  <div class="box-left-b" style="width: 70px;"></div>
                  <div class="box-right-b">
                    <div class="right-b-item">
                      <div class="b-item">
                        <div class="item">
                          <span class="label">{{ $t('dashboard.myTaskCount') }}</span>
                          <div class="value"><span class="text-28">{{ overviewData.aiTaskInfo.totalAiTasks }}</span>
                            {{ $t('dashboard.unitTask') }}</div>
                        </div>
                      </div>
                      <div class="b-item">
                        <div class="item">
                          <span class="label">{{ $t('dashboard.runningTasks') }}</span>
                          <div class="value"><span class="text-28" style="color: rgba(89, 219, 255, 1);">{{
                            overviewData.aiTaskInfo.runningAiTasks }}</span>
                            {{ $t('dashboard.unitTask') }}</div>
                        </div>

                      </div>
                      <div class="b-item">
                        <div class="item">
                          <span class="label">{{ $t('dashboard.totalGPUHourd') }}</span>
                          <div class="value"><span class="text-18">{{ overviewData.aiTaskInfo.allCardDurationCount
                              }}</span>
                            {{ $t('dashboard.runtimeCards') }}</div>
                        </div>

                      </div>
                    </div>
                  </div>
                </div>
              </div>
              <div class="template-section main-item">
                <div class="item-header-w">
                  <a class="quick-t" href="/ai_task_tmpl/list_my">
                    <div class="icon-item">
                      <img src="/img/home/dashboard-5.png" />
                    </div>
                    <span>{{ $t('taskTmplObj.taskTmpl') }}</span>
                  </a>
                </div>
                <div class="item-box-w">
                  <div class="box-left-b" style="width: 40px;"></div>
                  <div class="box-right-b">
                    <div class="right-b-item">
                      <div class="b-item">
                        <div class="item">
                          <span class="label">{{ $t('dashboard.templateNum') }}</span>
                          <div class="value"><span class="text-18">{{ overviewData.aiTaskInfo.templateCount }}</span>
                            {{ $t('dashboard.unitTask') }}</div>
                        </div>
                      </div>
                      <div class="b-item">
                        <div class="item">
                          <span class="label">{{ $t('dashboard.accumulatedRuns') }}</span>
                          <div class="value"><span class="text-18">{{ overviewData.aiTaskInfo.templateUseCount }}</span>
                            {{ $t('dashboard.unitTemplate') }}</div>
                        </div>
                      </div>
                    </div>
                  </div>
                </div>
              </div>
            </div>
            <div class="point-section main-item">
              <div class="item-header-w">
                <a class="quick-t" href="/reward/point">
                  <div class="icon-item">
                    <img src="/img/home/dashboard-4.webp" />
                  </div>
                  <span>{{ $t('dashboard.credits') }}</span>
                </a>
              </div>
              <div class="item-box-w">
                <div class="box-left-b"></div>
                <div class="box-right-b">
                  <div class="right-b-item">
                    <div class="b-item">
                      <div class="item">
                        <span class="label">{{ $t('dashboard.currentAvailable') }}</span>
                        <div class="value"><span class="text-28">{{ overviewData.pointInfo.balance }}</span>
                          {{ $t('dashboard.unitPoints') }}</div>
                      </div>
                    </div>
                    <div class="b-item">
                      <div class="item">
                        <span class="label">{{ $t('dashboard.totalGained') }}</span>
                        <div class="value"><span class="text-18">{{ overviewData.pointInfo.totalEarned }}</span>
                          {{ $t('dashboard.unitPoints') }}</div>
                      </div>

                    </div>
                    <div class="b-item">
                      <div class="item">
                        <span class="label">{{ $t('dashboard.totalConsumed') }}</span>
                        <div class="value"><span class="text-18">{{ overviewData.pointInfo.totalConsumed }}</span>
                          {{ $t('dashboard.unitPoints') }}</div>
                      </div>
                    </div>
                  </div>
                </div>
              </div>
            </div>

            <div class="repo-data-model-section">
              <div class="section main-item">
                <div class="item-header-w">
                  <a class="quick-t" href="/repositories">
                    <div class="icon-item">
                      <img src="/img/home/dashboard-6.png" />
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
                      <img src="/img/home/dashboard-7.png" />
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
                      <img src="/img/home/dashboard-8.png" />
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
              <div class="item-box-w">
                <div class="box-left-b">

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

                </div>
                <div class="box-right-b">
                  <div class="right-b-item">
                    <div class="b-item">
                      <div class="item">
                        <span class="label">{{ $t('dashboard.remainingQuota') }}</span>
                        <div class="value"><span class="text-28">{{ formattedRemaining.value }}</span> {{
                          formattedRemaining.label }}</div>
                      </div>
                    </div>
                    <div class="b-item">
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
          </template>
        </div>
        <rightCard />
      </div>
    </div>
  </div>
</template>

<script>
import { getOverviewData } from '~/apis/modules/dashboard';
import { getPromoteDataset } from "~/apis/modules/dataset";
import { TmplTaskTypes, TmplComputerResouces } from '~/pages/aitasktmpl/tools';
import { getListValueWithKey } from '~/utils';
import { lang } from '~/langs';
import rightCard from "./components/rightCard.vue";

import * as echarts from "echarts";
const UNITS = ['Bytes', 'KiB', 'MiB', 'GiB', 'TiB'];
export default {
  name: 'AiforgeIndex',
  components: { rightCard },
  data() {
    return {
      loading: false,
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

      list: [],
      lang: lang,
      avatarUrl: '',

      userName: '',
      visibleTags: [],
      isLastTagTruncated: false,
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
  watch: {
    list: {
      immediate: true,
      handler() {
        this.$nextTick(() => {
          this.calculateVisibleTags();
        });
      }
    }
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
    calculateVisibleTags() {
      if (!this.list || !this.list.length || !this.$refs.tagsContainer) {
        this.visibleTags = [...(this.list || [])];
        this.isLastTagTruncated = false;
        return;
      }

      const container = this.$refs.tagsContainer;
      const moreBtn = container.nextElementSibling;

      // 获取可用宽度
      const containerRect = container.getBoundingClientRect();
      const moreBtnRect = moreBtn.getBoundingClientRect();
      const availableWidth = containerRect.width - (moreBtnRect.width + 6); // 6px是左边距

      // 重置所有标签显示
      this.visibleTags = [...this.list];
      this.isLastTagTruncated = false;

      // 计算可以显示多少标签
      let usedWidth = 0;
      const gap = 10; // CSS中定义的gap
      const visibleIndexes = [];

      for (let i = 0; i < this.list.length; i++) {
        // 模拟标签宽度（最小宽度40px，最大180px）
        const tagText = this.list[i].Alias || '';
        const estimatedWidth = Math.min(Math.max(tagText.length * 8 + 40, 40), 180); // 简单估算宽度

        const currentGap = i > 0 ? gap : 0;
        const tagWidth = estimatedWidth;

        // 如果可以完全显示
        if (usedWidth + tagWidth + currentGap <= availableWidth) {
          usedWidth += tagWidth + currentGap;
          visibleIndexes.push(i);
        }
        // 如果是最后一个可见标签且还有空间，但不够完全显示
        else if (visibleIndexes.length > 0 && i === visibleIndexes[visibleIndexes.length - 1] + 1) {
          const remainingSpace = availableWidth - usedWidth - currentGap;
          if (remainingSpace > 20) { // 至少有20px空间显示部分内容
            visibleIndexes.push(i);
            this.isLastTagTruncated = true;
          }
          break;
        } else {
          break;
        }
      }

      // 设置可见标签
      this.visibleTags = visibleIndexes.map(index => this.list[index]);
    },
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
          return '23%'
        } else if (totalNum <= 99) {
          return '21.5%'
        } else if (totalNum <= 999) {
          return '20%'
        } else if (totalNum < 9999) {
          return '18%'
        } else if (totalNum < 99999) {
          return '16%'
        } else if (totalNum < 999999) {
          return '14%'
        } else if (totalNum < 9999999) {
          return '12%'
        } else if (totalNum < 99999999) {
          return '9%'
        }
      }
      return {
        tooltip: {
          trigger: 'item',
          formatter: (params) => {
            return `${params.name}: ${params.value}${config.unit}`
          }
        },
        legend: {
          type: 'scroll',
          orient: 'vertical',

          top: 'center',
          left: '45%',  // 从60%+5%的边距
          itemGap: 10,
          itemWidth: 8,  // 颜色块宽度
          textStyle: {
            fontSize: 12,
            rich: {
              name: {
                color: 'rgba(16,16,16,0.5)',
              },
              value: {
                color: '#101010',
                fontWeight: 'bold',
                align: 'left'
              },
              unit: {
                color: 'rgba(16,16,16,0.5)',
              }
            }
          },
          formatter: (name) => {
            const item = chartData.find(d => d.name === name)
            if (!item) return name
            return `{name|${name}}：{value|${item.value}} {unit|${config.unit}}`
          }
        },
        series: [
          {
            name: '数量统计',
            type: 'pie',
            radius: ['65%', '80%'],
            center: ['25%', '55%'],
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
            top: '40%',  // 数字位置
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
            left: '23%',
            top: '62%',  // 单位位置
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
    roundToTwo(num) {
      return num.toFixed(2);
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
        return { value: roundedValue.toFixed(decimals), label: UNITS[unitIndex] };
      }

    },

    async getPromoteData() {
      try {
        let key = 'home_v2/ai_task_template.json'
        let type = 'ai_task_template'
        const response = await getPromoteDataset({ key: key })
        const res = response.data
        if (res.code === 0) {
          this.list = res.data || []
          if (this.list.length) {
            this.list.forEach((item) => {
              item.Alias = item?.Alias || item.Name
              if (type == 'ai_task_template') {
                const jobTypeStr = getListValueWithKey(TmplTaskTypes, item.JobType)
                const computeSourceStr = getListValueWithKey(TmplComputerResouces, item.ComputeSource)
                item.Alias = item.Name
                item.labels = [jobTypeStr, computeSourceStr]
                item.link = `/cloudbrains/create?tmpl=${item.ID}`
                item.bgColor = 'rgb(150, 132, 255)'
                delete item.Tags
              }
            })
          }
        } else {
          this.$message.error(res.msg)
        }
      } catch (error) {

      }

    },
    async initData() {
      try {
        await Promise.all([
          this.getOverview(),
          this.getPromoteData()
        ]);
      } catch (error) {
        console.error('获取概览数据失败:', error);
        this.$message.error(error.message || '网络错误，请稍后重试');
      }

    },

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
    window.addEventListener('resize', this.calculateVisibleTags);
    window.addEventListener('resize', this.handleResize)
    this.initData();
  },
  beforeDestroy() {
    window.removeEventListener('resize', this.calculateVisibleTags);
    Object.values(this.chartInstances).forEach(chart => {
      if (chart) {
        chart.dispose()
      }
    })
    // 移除事件监听
    window.removeEventListener('resize', this.handleResize)
  },
};
</script>

<style lang="less" scoped>
.body-container {
  padding: 28px 34px 40px 40px;
  height: 100%;
  background: linear-gradient(180deg, rgba(0, 0, 255, 0.05) 0%, rgba(249, 249, 249, 1.00) 100%);

  .body-content {
    display: flex;
    flex-direction: column;
    height: 100%;

    .header-wrap {
      margin-bottom: 35px;
      display: flex;
      align-items: center;
      justify-content: space-between;
      height: 64px;

      .header-left-c {
        display: flex;

        .header-img {
          width: 64px;
          height: 64px;
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
      flex: 1;
      height: 0;
      display: flex;
      gap: 30px;

      .main-item {
        background-color: rgba(255, 255, 255, 1);
        border-radius: 10px;
        border: 1px solid rgba(255, 255, 255, 1);
        padding: 16px 16px 20px 24px;
        display: flex;
        flex-direction: column;
      }

      .main-left-c {
        flex: 2.2;
        height: 100%;
        display: flex;
        flex-direction: column;
        width: 0;
        gap: 35px;

        .item-header-w {
          margin-bottom: 20px;
        }

        .item-box-w {
          display: flex;

          .box-left-b {
            width: 150px;

            .progress-container {
              width: 100%;
              height: 100%;

              .progress-circle {
                transition: opacity 0.3s ease;
              }
            }

          }

          .box-right-b {
            flex: 1;
            width: 0;
            display: flex;

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
                  }

                  .value {
                    color: rgba(16, 16, 16, 0.5);
                    font-size: 18px;

                    span {
                      color: rgba(16, 16, 16, 1);
                      font-weight: 700;
                    }

                    .text-28 {
                      font-size: 28px;
                    }

                    .text-18 {
                      font-size: 18px;
                    }
                  }

                }

                svg {
                  margin-right: 10px;
                }

                img {
                  margin-right: 10px;
                  width: 32px;
                  height: 32px;
                }
              }
            }
          }
        }

        .quick-t {
          display: flex;
          align-items: center;
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
            font-size: 18px;
          }
        }

        .quick-section {
          height: 216px;

          .quick-c {
            display: flex;
            width: 100%;
            height: 87px;
            margin-top: 26px;
            gap: 15px;

            .quick-c-item {
              width: 25%;
              height: 100%;
              padding: 10px 15px;
              display: flex;

              .icon-w {
                width: 32px;
                height: 32px;

                img {
                  width: 100%;
                  height: 100%;
                }
              }

              .content-w {
                flex: 1;
                width: 0;
                display: flex;
                flex-direction: column;
                justify-content: space-between;
                color: rgba(16, 16, 16, 1);
                font-size: 16px;
                line-height: 24px;
                margin-left: 10px;

                .title {
                  &:hover {
                    color: rgba(0, 102, 255, 1);
                  }
                }

                .desc {
                  line-height: 18px;
                  color: rgba(16, 16, 16, 0.5);
                  font-size: 12px;
                  // max-height: 56px;
                  display: -webkit-box;
                  -webkit-box-orient: vertical;
                  -webkit-line-clamp: 2;
                  /* 限制行数 */

                  /* 可选的其他必要样式 */
                  overflow: hidden;
                  /* 隐藏超出部分 */
                  text-overflow: ellipsis;
                  /* 显示省略号 */
                  word-break: break-word;
                  /* 处理单词换行 */
                }
              }
            }
          }

          .quick-f {
            display: flex;
            align-items: center;
            color: rgba(16, 16, 16, 0.7);
            font-size: 12px;
            padding: 0 15px;
            margin-top: auto;
            overflow: hidden;
            /* 防止内容溢出 */
            gap: 10px;

            .quick-f-item-wrap {
              display: flex;
              flex: 1;
              overflow: hidden;
              gap: 10px;
              flex-wrap: nowrap;
              min-width: 0;

              .quick-f-item {
                display: flex;
                align-items: center;
                justify-content: center;
                border-radius: 6px;
                background-color: rgba(228, 241, 255, 1);
                height: 30px;
                padding: 0 10px;
                font-size: 12px;
                color: rgba(0, 102, 255, 1);
                flex-shrink: 1;
                min-width: 0;
                max-width: 180px;
                overflow: hidden;
                text-overflow: ellipsis;
                white-space: nowrap;
                transition: background-color 0.2s;

                svg {
                  margin-right: 6px;
                }
              }
            }

            .quick-f-more {
              margin-left: 6px;
              color: rgba(0, 102, 255, 1);
              border-bottom: 1px solid rgba(0, 102, 255, 1);
              flex-shrink: 0;
              margin-right: 20px;
            }
          }
        }

        .task-template-section {
          display: flex;
          gap: 24px;
          height: 157px;

          .task-section {
            height: 100%;
            flex: 2;
          }

          .template-section {
            height: 100%;
            flex: 1;
          }
        }

        .repo-data-model-section {
          display: flex;
          gap: 24px;
          height: 180px;

          .section {
            width: 100%;
            flex: 1;

            .item-header-w {
              margin: 0;
            }
          }
        }

        .storage-section,
        .point-section {
          height: 172px;
        }

        .chart-container {
          flex: 1;

          .chart {
            height: 100%;
          }
        }
      }
    }
  }

}

@media screen and (min-width: 768px) and (max-width: 1500px) {

  /* 在这里编写只在768px到1500px之间生效的样式 */
  .main-wrap {
    flex-direction: column;

    .main-left-c {
      width: 100% !important;
    }

    .main-right-c {
      .notice-w {
        max-width: 100% !important;
      }
    }
  }
}

@media screen and (max-width: 1200px) {

  /* 在这里编写只在768px到1500px之间生效的样式 */
  .b-item {
    svg {
      display: none;
    }

    img {
      display: none;
    }
  }
}

:lang(en-US) .b-item {
  .item {
    .line {
      .label {
        min-width: 126px !important;
      }
    }
  }
}
</style>