<template>
  <div class="content-box">
    <!-- 任务运行中悬浮提示（只在广场视图显示） -->
    <div class="task-running-tip" v-if="taskRunning && activeTab === 'all'">
      <div class="circle-wrap">
        <img src="/img/agent/dialog-1.png" alt="">
      </div>
      <div class="task-main">
        <div class="task-info">{{ $t('agentPortal.taskRunningTip', { count, numPoint }) }}</div>
        <div class="task-detail" @click="changeTab({ key: 'mine' })">
          <span>{{ $t('agentPortal.viewRunningAgents') }}</span>
        </div>
      </div>
      <div class="close-wrap" @click="closeTaskTip">
        <svg width="12" height="12" viewBox="0 0 48 48" fill="none" xmlns="http://www.w3.org/2000/svg">
          <path d="M8 8L40 40" stroke="#505559" stroke-width="4" stroke-linecap="round" stroke-linejoin="round" />
          <path d="M8 40L40 8" stroke="#505559" stroke-width="4" stroke-linecap="round" stroke-linejoin="round" />
        </svg>
      </div>
      <img style="position: absolute; right: 0;" src="/img/agent/dialog-2.png" alt="">
    </div>

    <div class="head-banner-wrap">
      <div class="banner-content">
        <div class="banner-title">{{ $t('agentPortal.bannerTitle') }}</div>
        <div class="banner-desc">{{ $t('agentPortal.bannerDesc') }}</div>
      </div>
      <div class="banner-img">
        <img src="/img/agent/banner-1.png" alt="">
      </div>
    </div>

    <div class="main-body">

      <div class="filter-c">
        <div class="tab-c">
          <div class="tab-item nowrap" :class="{ active: activeTab === item.key }" v-for="(item) in tabs"
            :key="item.key" @click="changeTab(item)">{{
              item.label }}</div>
        </div>
        <div class="right">
          <div class="search-input-wrap">
            <input :placeholder="$t('agentPortal.searchPlaceholder')" :value="searchText" @input="onSearchInput"
              @keyup.enter="handleSearch">
            <i class="search icon" @click.stop="handleSearch"></i>
          </div>
        </div>

      </div>
      <div class="main-content" v-loading="loading">
        <template v-if="displayList.length">
          <div class="card-grid">
            <div v-for="agent in displayList" :key="agent.id" class="agent-card">
              <div class="card-content">
                <div class="card-icon">
                  <img :src="agent.avatar_url || agent.avatar || '/img/agent/agent-placeholder.png'" />
                </div>
                <div class="card-info">
                  <div class="card-name">{{ agent.name }}</div>
                  <div class="card-desc" :title="agent.description">{{ agent.description }}</div>
                  <div class="card-labels" :ref="'labels-' + agent.id">
                    <span class="label-item" v-for="(label, idx) in (agent.tag || [])" :key="idx"
                      v-show="idx < labelVisibleMap[agent.id]">{{ label }}</span>
                    <el-popover
                      v-if="(agent.tag || []).length - (labelVisibleMap[agent.id] || (agent.tag || []).length) > 0"
                      placement="bottom" trigger="click" popper-class="labels-popover">
                      <div class="labels-popover-content">
                        <span class="pop-label-item"
                          v-for="(label, idx) in (agent.tag || []).slice(labelVisibleMap[agent.id] || (agent.tag || []).length)"
                          :key="idx">{{ label
                          }}</span>
                      </div>
                      <span class="label-overflow" slot="reference">+{{ (agent.tag || []).length -
                        (labelVisibleMap[agent.id] || (agent.tag || []).length) }}</span>
                    </el-popover>
                  </div>
                </div>
              </div>
              <div class="card-middle" v-if="agent.spec">
                <span>{{ getSpecStr(agent.spec).specStr }}</span>
                <div class="point-wrap">
                  <svg xmlns="http://www.w3.org/2000/svg" class="" viewBox="0 0 1024 1024" width="16" height="16">
                    <defs></defs>
                    <g>
                      <path
                        d="M44.8 512c0 256 211.2 467.2 467.2 467.2S979.2 768 979.2 512 768 44.8 512 44.8 44.8 256 44.8 512z"
                        fill="#FBE945" p-id="8950"></path>
                      <path
                        d="M153.6 512c0 198.4 160 358.4 358.4 358.4s358.4-160 358.4-358.4S710.4 153.6 512 153.6 153.6 313.6 153.6 512z"
                        fill="#FBB11B" p-id="8951"></path>
                      <path
                        d="M870.4 512c0-76.8-25.6-153.6-70.4-211.2-57.6-44.8-134.4-70.4-211.2-70.4-198.4 0-358.4 160-358.4 358.4 0 76.8 25.6 153.6 70.4 211.2 57.6 44.8 134.4 70.4 211.2 70.4 198.4 0 358.4-160 358.4-358.4z"
                        fill="#FDC72F" p-id="8952"></path>
                      <path
                        d="M550.4 307.2l76.8 153.6 166.4 32-115.2 128 19.2 172.8-147.2-76.8-153.6 76.8 25.6-172.8-115.2-128 160-32 83.2-153.6"
                        fill="#FBB11B" p-id="8953"></path>
                      <path
                        d="M512 268.8l83.2 153.6 160 32-115.2 128 19.2 172.8L512 678.4l-147.2 76.8 19.2-172.8-115.2-128 160-32L512 268.8"
                        fill="#FBE945" p-id="8954"></path>
                    </g>
                  </svg>
                  <span>{{ getSpecStr(agent.spec).pointStr }}</span>
                </div>

              </div>
              <div class="card-footer">
                <div class="footer-l" :ref="'footerL-' + agent.id">
                  <div class="card-meta">
                    <a :href="`/${agent.user_name}`" class="avatar-c">
                      <img class="avatar" :src="agent.user_avatar">
                    </a>
                    <span class="update-time" style="margin-left:8px;"
                      :title="`${$t('agentPortal.updated')} ${formatTime(agent.updated_unix)}`"><span class="time-1">{{
                        $t('agentPortal.updated') }}</span> {{
                          formatTime(agent.updated_unix) }}</span>
                  </div>
                  <div class="card-count">
                    <svg width="16" height="16" viewBox="0 0 48 48" fill="none" xmlns="http://www.w3.org/2000/svg">
                      <path
                        d="M24 44C35.0457 44 44 35.0457 44 24C44 12.9543 35.0457 4 24 4C12.9543 4 4 12.9543 4 24C4 35.0457 12.9543 44 24 44Z"
                        fill="none" stroke="#888" stroke-width="4" stroke-linejoin="round" />
                      <path d="M20 24V17.0718L26 20.5359L32 24L26 27.4641L20 30.9282V24Z" fill="none" stroke="#888"
                        stroke-width="4" stroke-linejoin="round" />
                    </svg>
                    <span style="margin-left: 5px;">{{ agent.use_count }}</span>
                  </div>
                </div>

                <div class="footer-r" :ref="'footerR-' + agent.id">
                  <div class="card-detail" v-if="agent.detail_url && isDetailVisible(agent.id)">
                    <a :href="agent.detail_url" target="_blank">{{ $t('agentPortal.viewDetails') }}</a>
                    <svg width="12" height="12" viewBox="0 0 48 48" fill="none" xmlns="http://www.w3.org/2000/svg">
                      <path d="M19 11H37V29" stroke="#101010" stroke-width="4" stroke-linecap="round"
                        stroke-linejoin="round" />
                      <path d="M11.5439 36.4559L36.9997 11" stroke="#101010" stroke-width="4" stroke-linecap="round"
                        stroke-linejoin="round" />
                    </svg>
                  </div>

                  <!-- 一键部署按钮 -->
                  <div class="deploy-btn oper-btn"
                    v-if="agent.computedStatus === 'idle' || agent.computedStatus === 'stopped' || agent.computedStatus === 'failed'"
                    @click="handleRun(agent)">
                    <svg width="12" height="12" viewBox="0 0 48 48" fill="none" xmlns="http://www.w3.org/2000/svg">
                      <path d="M24 44C12.9543 44 4 35.0457 4 24C4 12.9543 12.9543 4 24 4C35.0457 4 44 12.9543 44 24"
                        stroke="#0056ff" stroke-width="4" stroke-linecap="round" stroke-linejoin="round" />
                      <path d="M20 24V17.0718L26 20.5359L32 24L26 27.4641L20 30.9282V24Z" fill="none" stroke="#0056ff"
                        stroke-width="4" stroke-linejoin="round" />
                      <path d="M37.0508 32L37.0508 42" stroke="#0056ff" stroke-width="4" stroke-linecap="round"
                        stroke-linejoin="round" />
                      <path d="M42 36.9497L32 36.9497" stroke="#0056ff" stroke-width="4" stroke-linecap="round"
                        stroke-linejoin="round" />
                    </svg>
                    <span>{{ $t('agentPortal.oneClickDeploy') }}</span>
                  </div>

                  <!-- 失败提示 + 一键部署 -->
                  <div class="failed-tip" v-if="agent.computedStatus === 'failed'">
                    <span class="failed-text">{{ $t('agentPortal.deployFailed') }}</span>
                  </div>

                  <!-- 部署中状态 -->
                  <div class="run-wrap" v-if="agent.computedStatus === 'deploying'">
                    <div class="oper-btn wait-btn">
                      <svg width="12" height="12" viewBox="0 0 48 48" fill="none" xmlns="http://www.w3.org/2000/svg">
                        <path d="M7 4H41" stroke="#0056ff" stroke-width="4" stroke-linecap="round"
                          stroke-linejoin="round" />
                        <path d="M7 44H41" stroke="#0056ff" stroke-width="4" stroke-linecap="round"
                          stroke-linejoin="round" />
                        <path d="M11 44C13.6667 30.6611 18 23.9944 24 24C30 24.0056 34.3333 30.6722 37 44H11Z"
                          fill="none" stroke="#0056ff" stroke-width="4" stroke-linejoin="round" />
                        <path d="M37 4C34.3333 17.3389 30 24.0056 24 24C18 23.9944 13.6667 17.3278 11 4H37Z" fill="none"
                          stroke="#0056ff" stroke-width="4" stroke-linejoin="round" />
                      </svg>
                      <span>{{ $t('agentPortal.deploying') }}</span>
                    </div>
                    <div class="oper-btn stop-btn" @click="handleStop(agent)">
                      <svg xmlns="http://www.w3.org/2000/svg" fill="#ff6200" viewBox="64 64 896 896" width="12"
                        height="12">
                        <defs></defs>
                        <g>
                          <path
                            d="M512 64C264.6 64 64 264.6 64 512s200.6 448 448 448 448-200.6 448-448S759.4 64 512 64zm0 820c-205.4 0-372-166.6-372-372 0-89 31.3-170.8 83.5-234.8l523.3 523.3C682.8 852.7 601 884 512 884zm288.5-137.2L277.2 223.5C341.2 171.3 423 140 512 140c205.4 0 372 166.6 372 372 0 89-31.3 170.8-83.5 234.8z">
                          </path>
                        </g>
                      </svg>
                      <span>{{ $t('agentPortal.stop') }}</span>
                    </div>
                  </div>

                  <!-- 运行成功状态 -->
                  <div class="run-wrap" v-if="agent.computedStatus === 'running'">
                    <div class="oper-btn start-btn" @click="handleUse(agent)">
                      <svg xmlns="http://www.w3.org/2000/svg" fill="#fff" viewBox="0 0 32 32" width="16" height="16">
                        <defs></defs>
                        <g>
                          <path
                            d="M16 2.667c7.36 0 13.333 5.973 13.333 13.333s-5.973 13.333-13.333 13.333-13.333-5.973-13.333-13.333 5.973-13.333 13.333-13.333zM16 26.667c5.893 0 10.667-4.773 10.667-10.667s-4.773-10.667-10.667-10.667-10.667 4.773-10.667 10.667 4.773 10.667 10.667 10.667zM20.715 9.4l1.885 1.885-6.6 6.6-1.885-1.885 6.6-6.6z">
                          </path>
                        </g>
                      </svg>
                      <span>{{ $t('agentPortal.startUsing') }}</span>
                    </div>
                    <div class="oper-btn stop-btn" @click="handleStop(agent)">
                      <svg xmlns="http://www.w3.org/2000/svg" fill="#ff6200" viewBox="64 64 896 896" width="12"
                        height="12">
                        <defs></defs>
                        <g>
                          <path
                            d="M512 64C264.6 64 64 264.6 64 512s200.6 448 448 448 448-200.6 448-448S759.4 64 512 64zm0 820c-205.4 0-372-166.6-372-372 0-89 31.3-170.8 83.5-234.8l523.3 523.3C682.8 852.7 601 884 512 884zm288.5-137.2L277.2 223.5C341.2 171.3 423 140 512 140c205.4 0 372 166.6 372 372 0 89-31.3 170.8-83.5 234.8z">
                          </path>
                        </g>
                      </svg>
                      <span>{{ $t('agentPortal.stop') }}</span>
                    </div>
                  </div>
                </div>
              </div>

            </div>

          </div>
          <div class="pagination-wrapper">
            <el-pagination background :small="isMiniScreen" @current-change="currentChange" @size-change="sizeChange"
              :current-page.sync="page" :page-sizes="pageSizes" :page-size.sync="pageSize"
              layout="total, sizes, prev, pager, next, jumper" :total="paginationTotal">
            </el-pagination>
          </div>
        </template>
        <div class="no-data" v-else>
          <span>{{ $t('noData') }}</span>
        </div>
      </div>



    </div>
  </div>
</template>

<script>
import { Image } from 'element-ui';
import { getAgentList, getUserAgentList, runAgent, stopAgent } from '~/apis/modules/agent';
import { getAiEndpointUrl } from '~/apis/modules/cloudbrain';
import { renderSpecObject } from '~/utils';
import { formatDate } from 'element-ui/lib/utils/date-util';
// 任务状态常量
const JOB_STATUS = {
  IDLE: 'idle',
  DEPLOYING: 'deploying', // INIT, STARTING, CREATING, WAITING, PREPARING, CONNECTING
  RUNNING: 'running',
  FAILED: 'failed', // FAILED, START_FAILED, CREATE_FAILED, CREATED_FAILED, IMAGE_FAILED, SUBMIT_FAILED
  STOPPED: 'stopped', // STOPPED, KILLED
};

// 部署中状态列表
const DEPLOYING_STATUS = ['INIT', 'STARTING', 'CREATING', 'WAITING', 'PREPARING', 'CONNECTING'];
// 运行成功状态列表
const RUNNING_STATUS = ['RUNNING'];
// 失败状态列表
const FAILED_STATUS = ['FAILED', 'START_FAILED', 'CREATE_FAILED', 'CREATED_FAILED', 'IMAGE_FAILED', 'SUBMIT_FAILED', 'DEPLOY_SERVICE_FAILED', 'CHECK_FAILED'];
// 停止状态列表
const STOPPED_STATUS = ['STOPPED', 'KILLED', 'COMPLETED', 'SUCCEEDED', 'CANCELED', 'LOST', 'DELETED'];

export default {
  name: 'AgentList',
  data() {
    return {
      activeTab: 'all',
      q: '', // 已提交的搜索关键词（Enter 时从 searchText 同步）
      agentList: [], // 智能体广场列表
      agentInfoCache: {}, // 跨页智能体基本信息缓存 { agent_id: agent_info }，避免翻页后丢失 spec
      userAgentMap: {}, // 用户智能体状态映射 { agent_id: user_agent_info }
      total: 0,
      pageSize: 12,
      pageSizes: [12],
      page: 1,
      isMiniScreen: false,
      clientWidth: window.innerWidth,
      loading: false,
      pollingTimer: null,
      taskRunningDismissed: false, // 弹窗是否被手动关闭
      wrappedMap: {}, // 记录每个卡片是否换行 { agentId: true/false }
      labelVisibleMap: {}, // 记录每个卡片可见标签数 { agentId: count }
      resizeObserver: null,
    };
  },
  components: {},
  watch: {
    displayList() {
      this.$nextTick(() => {
        this.$nextTick(() => {
          this.checkFooterWrap()
          this.checkLabelOverflow()
        })
      })
    },
  },
  computed: {
    // 合并展示列表
    displayList() {
      if (this.activeTab === 'mine') {
        // 用户视角：只看运行中的（running状态）
        // 搜索过滤通过 agentList（后端已按 q 过滤）与 userAgentMap 取交集，
        // 不依赖 this.q，避免每次按键触发重渲染和 DOM 测量
        const agentIdSet = new Set(this.agentList.map(a => a.id));
        return Object.values(this.userAgentMap).filter(
          agent => agent.computedStatus === JOB_STATUS.RUNNING && agentIdSet.has(agent.id)
        );
      }
      // 广场视角：合并状态
      return this.agentList.map(agent => {
        const userAgent = this.userAgentMap[agent.id];
        if (userAgent) {
          return { ...agent, ...userAgent };
        }
        // 没有用户智能体信息时，视为 idle 状态，显示一键部署按钮
        return { ...agent, computedStatus: JOB_STATUS.IDLE };
      });
    },
    // 分页总数：mine tab 用运行中智能体数量，all tab 用广场总数
    paginationTotal() {
      if (this.activeTab === 'mine') {
        const agentIdSet = new Set(this.agentList.map(a => a.id));
        return Object.values(this.userAgentMap).filter(
          agent => agent.computedStatus === JOB_STATUS.RUNNING && agentIdSet.has(agent.id)
        ).length;
      }
      return this.total;
    },
    // 是否已登录
    isLogin() {
      return !!document.querySelector('meta[name="_uid"]');
    },
    // 标签页列表
    tabs() {
      if (!this.isLogin) {
        return [{ key: 'all', label: '🤖 ' + this.$t('agentPortal.agentSquare') }];
      }
      return [
        { key: 'all', label: '🤖 ' + this.$t('agentPortal.agentSquare') },
        { key: 'mine', label: '🧑‍ ' + this.$t('agentPortal.runningAgents') },
      ];
    },
    // 是否有任务正在运行
    taskRunning() {
      return !this.taskRunningDismissed && Object.values(this.userAgentMap).some(
        agent => agent.computedStatus === JOB_STATUS.RUNNING
      );
    },
    // 运行中的任务信息
    taskInfo() {
      const runningAgents = Object.values(this.userAgentMap).filter(
        agent => agent.computedStatus === JOB_STATUS.RUNNING
      );
      if (runningAgents.length === 0) return '';
      const names = runningAgents.map(a => a.name).join('、');
      return this.$t('agentPortal.agentsRunning', { names });
    },
    // 正在运行的智能体数量
    count() {
      return Object.values(this.userAgentMap).filter(
        agent => agent.computedStatus === JOB_STATUS.RUNNING
      ).length;
    },
    // 正在运行的智能体消耗积分累加
    numPoint() {
      return Object.values(this.userAgentMap)
        .filter(agent => agent.computedStatus === JOB_STATUS.RUNNING)
        .reduce((sum, agent) => {
          // 从 agent.spec.unit_price 获取积分，默认为 0
          const price = agent.spec?.unit_price || 0;
          return sum + price;
        }, 0);
    },
  },
  methods: {
    // 计算智能体状态
    getAgentStatus(jobStatus) {
      if (!jobStatus) return JOB_STATUS.IDLE;
      if (RUNNING_STATUS.includes(jobStatus)) return JOB_STATUS.RUNNING;
      if (FAILED_STATUS.includes(jobStatus)) return JOB_STATUS.FAILED;
      if (STOPPED_STATUS.includes(jobStatus)) return JOB_STATUS.STOPPED;
      if (DEPLOYING_STATUS.includes(jobStatus)) return JOB_STATUS.DEPLOYING;
      return JOB_STATUS.IDLE;
    },
    changeTab(item) {
      this.activeTab = item.key;
      this.page = 1;
      this.wrappedMap = {};
      this.labelVisibleMap = {};
      this.loadData();
    },
    formatTime(time) {
      return formatDate(new Date(time * 1000), 'yyyy-MM-dd')
    },
    getSpecStr(spec) {
      if (!spec) return '';
      const specObj = renderSpecObject({
        id: '',
        compute_resource: spec.compute_resource,
        acc_cards_num: spec.acc_cards_num,
        acc_card_type: spec.acc_card_type,
        gpu_mem_gi_b: spec.gpu_mem_gib || 0,
        mem_gi_b: spec.mem_gib || 0,
        share_mem_gi_b: spec.share_mem_gib || 0,
        cpu_cores: spec.cpu_cores,
        unit_price: spec.unit_price || 0,
      }, true);
      return specObj;
    },
    // DOM 渲染后测量实际标签元素宽度，决定每个卡片显示几个标签
    // 分两步：先全部显示（确保 offsetWidth 不为 0），再测量裁剪
    checkLabelOverflow() {
      // 第一步：将所有标签设为可见，避免被 v-show 隐藏的标签 offsetWidth 为 0
      this.displayList.forEach(agent => {
        if (agent.tag && agent.tag.length) {
          this.$set(this.labelVisibleMap, agent.id, agent.tag.length);
        }
      });
      // 第二步：等 DOM 更新后测量真实宽度
      this.$nextTick(() => {
        this.displayList.forEach(agent => {
          const tags = agent.tag;
          if (!tags || !tags.length) {
            this.$set(this.labelVisibleMap, agent.id, 0);
            return;
          }
          const labelsRef = this.$refs['labels-' + agent.id];
          const containerEl = Array.isArray(labelsRef) ? labelsRef[0] : labelsRef;
          if (!containerEl) return;
          const containerWidth = containerEl.offsetWidth;
          if (containerWidth === 0) return;

          const labelEls = containerEl.querySelectorAll('.label-item');
          if (!labelEls.length) return;

          const overflowBadgeWidth = 30;
          const availableWidth = containerWidth - overflowBadgeWidth;
          let visibleCount = 0;
          let usedWidth = 0;
          for (let i = 0; i < labelEls.length; i++) {
            const elWidth = labelEls[i].offsetWidth + 10;
            if (usedWidth + elWidth > availableWidth) break;
            usedWidth += elWidth;
            visibleCount++;
          }
          this.$set(this.labelVisibleMap, agent.id, Math.max(visibleCount, 1));
        });
      });
    },
    // searchText 是非响应式属性（在 created 中挂载），@input 更新不触发组件重渲染，
    // 从而消除 all tab 下 12 张卡片 vdom diff 造成的输入延迟
    onSearchInput(e) {
      const oldVal = this.searchText;
      this.searchText = e.target.value;
      if (oldVal && !this.searchText) {
        this.q = '';
        this.handleSearch();
      }
    },
    handleSearch() {
      this.q = this.searchText;
      this.page = 1;
      this.loadData();
    },
    closeTaskTip() {
      this.taskRunningDismissed = true;
    },
    currentChange(page) {
      this.page = page;
      this.loadData();
    },
    sizeChange(pageSize) {
      this.page = 1;
      this.pageSize = pageSize;
      this.loadData();
    },
    // 加载广场列表
    async loadAgentList() {
      try {
        const res = await getAgentList({ q: this.q, page: this.page, pageSize: this.pageSize });
        if (res.data.code === 0) {
          this.agentList = res.data.data.agents || [];
          this.total = res.data.data.total || 0;
          // 缓存智能体信息，供 loadUserAgentMap 跨页查找 spec
          this.agentList.forEach(a => { this.$set(this.agentInfoCache, a.id, a); });
        }
      } catch (err) {
        console.error('Failed to load agent list:', err);
      }
    },
    // 加载用户智能体状态
    async loadUserAgentMap() {
      try {
        const res = await getUserAgentList({ page: 1, pageSize: 100 });
        if (res.data.code === 0) {
          const map = {};
          (res.data.data.user_agents || []).forEach(item => {
            // 从跨页缓存中查找智能体基本信息，避免翻页后 spec 丢失
            console.log("xxxxxxxxxxx", item)
            console.log("xxxxxxxxxxx", this.agentInfoCache)
            const agentInfo = this.agentInfoCache[item.agent_id] || {};
            map[item.agent_id] = {
              ...agentInfo,
              ...item,
              id: item.agent_id, // 确保 id 是智能体的 id，而非用户智能体自己的 id
              computedStatus: this.getAgentStatus(item.job_status),
            };
          });
          this.userAgentMap = map;
        }
      } catch (err) {
        console.error('Failed to load user agent map:', err);
      }
    },
    async loadData() {
      this.loading = true;
      try {
        // 先加载广场列表，再加载用户智能体状态（依赖 agentList）
        await this.loadAgentList();
        // 只有登录用户才加载用户智能体状态
        if (this.isLogin) {
          await this.loadUserAgentMap();
        }
      } finally {
        this.loading = false;
      }
      // 初始化 labelVisibleMap，让标签先全部显示，等待 DOM 后再测量裁剪
      this.agentList.forEach(agent => {
        if (agent.tag && agent.tag.length) {
          this.$set(this.labelVisibleMap, agent.id, agent.tag.length);
        }
      });
      // 启动轮询（只有登录用户才启动）
      if (this.isLogin) {
        this.startPolling();
      }
    },
    // 启动轮询（只有存在运行中的任务时才轮询）
    startPolling() {
      this.stopPolling();
      // 如果没有运行中的任务，不启动轮询
      if (Object.keys(this.userAgentMap).length === 0) {
        return;
      }
      // 每5秒轮询一次
      this.pollingTimer = setInterval(() => {
        this.loadUserAgentMap();
      }, 5000);
    },
    // 停止轮询
    stopPolling() {
      if (this.pollingTimer) {
        clearInterval(this.pollingTimer);
        this.pollingTimer = null;
      }
    },
    async handleRun(agent) {
      if (!this.isLogin) {
        window.location.href = '/user/login?redirect_to=' + encodeURIComponent(window.location.href);
        return;
      }
      try {
        const res = await runAgent(agent.id);
        if (res.data.code === 0) {
          this.$message.success(this.$t('agentPortal.startSuccess'));
          // 先 +1，再 loadUserAgentMap，避免两次渲染间的滞后
          const src = this.agentList.find(a => a.id === agent.id);
          if (src) {
            this.$set(src, 'use_count', (src.use_count || 0) + 1);
          }
          // 刷新用户智能体状态
          await this.loadUserAgentMap();
          // 启动轮询（确保部署后能持续更新状态）
          this.startPolling();
          // 如果返回了 cloudbrain_id，保存到 agent
          if (res.data.data && res.data.data.cloudbrain_id) {
            agent.cloudbrain_id = res.data.data.cloudbrain_id;
          }
        } else {
          this.$message.error(res.data.msg || this.$t('agentPortal.startFailed'));
        }
      } catch (err) {
        console.error('Failed to run agent:', err);
        this.$message.error(this.$t('agentPortal.startFailed'));
      }
    },
    async handleStop(agent) {
      try {
        const id = agent.cloudbrain_id || agent.id;
        const res = await stopAgent(id);
        if (res.data.code === 0) {
          this.$message.success(this.$t('agentPortal.stopSuccess'));
          await this.loadUserAgentMap();
          // 根据是否有运行中的任务决定是否继续轮询
          this.startPolling();
        } else {
          this.$message.error(res.data.msg || this.$t('agentPortal.stopFailed'));
        }
      } catch (err) {
        console.error('Failed to stop agent:', err);
        this.$message.error(this.$t('agentPortal.stopFailed'));
      }
    },
    handleUse(agent) {
      // 开始使用 - 跳转到使用页面
      getAiEndpointUrl({
        id: agent.cloudbrain_id,
      }).then(res => {
        res = res.data;
        if (res.code == 0) {
          if (res.data && res.data.url) {
            if (!window.open(res.data.url)) {
              window.location.href = res.data.url;
            }
          }
        } else {
          this.$message({
            type: 'error',
            message: res.msg,
          });
        }
      }).catch(err => {
        this.$message({
          type: 'error',
          message: this.$t('operationFailed'),
        });
      });
    },
    calcScreenInfo() {
      this.isMiniScreen = document.documentElement.clientWidth <= 800;
    },
    isDetailVisible(agentId) {
      return this.wrappedMap[agentId] !== true
    },
    checkFooterWrap() {
      this.displayList.forEach(agent => {
        const footerL = this.$refs['footerL-' + agent.id]
        const footerR = this.$refs['footerR-' + agent.id]
        if (footerL && footerR) {
          const footerLEl = Array.isArray(footerL) ? footerL[0] : footerL
          const footerREl = Array.isArray(footerR) ? footerR[0] : footerR
          const wrapped = footerREl.offsetTop > footerLEl.offsetTop
          // 只在换行时隐藏"了解详情"，不反向恢复，避免振荡
          if (wrapped) {
            this.$set(this.wrappedMap, agent.id, true)
          } else if (!this.wrappedMap[agent.id]) {
            // 从未换行过，保持显示
            this.$set(this.wrappedMap, agent.id, false)
          }
          // 如果曾经换行被隐藏，即使现在不换了也不恢复，
          // 等 resize 或 tab 切换时通过 resetFooterWrap 重新判断
        }
      })
    },
    resetFooterWrap() {
      this.wrappedMap = {}
      this.labelVisibleMap = {}
      this.$nextTick(() => {
        this.checkFooterWrap()
        this.checkLabelOverflow()
      })
    },
    observeCardWidth() {
      this.$nextTick(() => {
        const grid = document.querySelector('.card-grid')
        if (grid && this.resizeObserver) {
          this.resizeObserver.observe(grid)
        }
        this.checkFooterWrap()
      })
    },
  },
  created() {
    this.searchText = ''; // 非响应式属性，按键更新不触发组件重渲染
    window.addEventListener('resize', this.calcScreenInfo);
    this.loadData();
    const params = {
      agent_type: "OPENCLAW",
      template_id: "06d90c64-a586-476e-92c0-78ee4d40d16c",
      status: 1,
      name: "ruuning-2x11",
      avatar_url: "https://example.com/avatar.png",
      back_ground_url: "https://example.com/background.png",
      detail_url: "https://example.com/detail",
      tag: [
        "机器学习-newnewnwew",
        "图像处理-new",
        "第三个",
        "第四个",
        "第五个",
        "第六个"
      ],
      creator_id: 8,
      description: "这是一个智能体的描述xx"
    }
    // createAgent(params).then((res) => {
    //   console.log(res)
    // })
    // deleteAgent(6).then((res) => { })
  },
  mounted() {
    this.resizeObserver = new ResizeObserver(() => {
      this.resetFooterWrap()
    })
    this.$nextTick(() => this.observeCardWidth())
  },
  beforeDestroy() {
    window.removeEventListener('resize', this.calcScreenInfo);
    this.stopPolling();
    if (this.resizeObserver) {
      this.resizeObserver.disconnect()
    }
  },
};
</script>

<style scoped lang="less">
.content-box {
  width: 100%;
  display: flex;
  flex-direction: column;
  background: #f5f5f5;
  padding: 0 38px 0 42px;
  min-height: 100%;
  box-sizing: border-box;

  .head-banner-wrap {
    width: 100%;
    height: 360px;
    background: radial-gradient(0.5% 1.15% at 36.199999999999996% 63.9%, rgba(0, 199, 255, 1) 0%, rgba(50, 145, 248, 1) 100%);
    margin-top: 37px;
    border-radius: 15px;
    display: flex;
    position: relative;
    overflow: hidden;

    .banner-content {
      margin-top: 117px;
      margin-left: 136px;
      max-width: 60%;
      z-index: 2;

      .banner-title {
        color: rgba(255, 255, 255, 1);
        font-size: 32px;
        font-weight: 700;
        line-height: 42px;
        height: 42px;
      }

      .banner-desc {
        color: rgba(255, 255, 255, 0.7);
        font-size: 16px;
        margin-top: 14px;
        line-height: 21px;
      }
    }

    .banner-img {
      position: absolute;
      right: 78px;
      bottom: -6px;
      z-index: 1;

      img {
        max-width: 100%;
        height: auto;
      }
    }
  }

  .main-body {
    flex: 1;
    display: flex;
    flex-direction: column;

    .filter-c {
      display: flex;
      align-items: center;
      justify-content: space-between;
      margin-top: 38px;
      margin-bottom: 28px;
      flex-wrap: wrap;
      gap: 16px;

      .tab-c {
        display: flex;
        align-items: center;
        flex-wrap: wrap;
        gap: 6px;

        .tab-item {
          height: 32px;
          display: flex;
          align-items: center;
          padding: 0 6px;
          border-bottom: 2px solid rgba(51, 38, 98, 0.3);
          font-size: 16px;
          box-sizing: border-box;
          cursor: pointer;
          color: rgba(16, 16, 16, 0.5);
          white-space: nowrap;

          &:hover {
            color: rgba(16, 16, 16, 1);
            border-color: rgba(51, 38, 98, 1);
          }

          &.active {
            color: rgba(0, 102, 255, 1);
            border-bottom: 2px solid rgba(0, 102, 255, 1);
          }
        }
      }

      .right {
        display: flex;
        align-items: center;
        gap: 14px;

        .search-input-wrap {
          position: relative;
          height: 32px;
          width: 380px;
          max-width: calc(100vw - 300px);
          cursor: pointer;

          input {
            width: 100%;
            height: 100%;
            padding: 0 32px 0 12px;
            border: 1px solid rgba(240, 242, 245, 1);
            border-radius: 4px;
            font-size: 14px;
            box-sizing: border-box;
            cursor: pointer;
            outline: none;
            transition: border-color 0.2s, box-shadow 0.2s;

            &:focus {
              border-color: rgba(0, 102, 255, 0.5);
              box-shadow: 0 0 0 2px rgba(0, 102, 255, 0.1);
            }
          }

          .search.icon {
            position: absolute;
            right: 10px;
            top: 50%;
            transform: translateY(-50%);
            cursor: pointer;
            pointer-events: auto;
          }
        }

        .search-c {
          margin-right: 16px;

          .el-input {
            .el-button {
              color: rgb(16, 16, 16);
            }
          }
        }
      }
    }

    .main-content {
      flex: 1;

      .card-grid {
        display: grid;
        grid-template-columns: repeat(3, minmax(0, 1fr));
        gap: 20px;
        width: 100%;

        .agent-card {
          background: linear-gradient(176.14deg, rgba(228, 252, 255, 1) 0.84%, rgba(255, 255, 255, 1) 94.19%);
          border-radius: 10px;
          padding: 20px;
          box-shadow: 0px 0px 20px 0px rgba(221, 221, 221, 0.5);
          transition: all 0.3s ease;
          display: flex;
          flex-direction: column;

          &:hover {
            background-color: #fff;
          }

          .card-content {
            display: flex;
            gap: 12px;

            .card-icon {
              flex-shrink: 0;
              width: 48px;
              height: 48px;
              border-radius: 6px;
              background-color: rgba(229, 229, 229, 1);

              img {
                width: 100%;
                height: 100%;
                object-fit: cover;
              }
            }

            .card-info {
              flex: 1;
              min-width: 0;

              .card-name {
                font-size: 14px;
                font-weight: 600;
                color: rgba(16, 16, 16, 1);
                margin-bottom: 10px;
                overflow: hidden;
                text-overflow: ellipsis;
                white-space: nowrap;
              }

              .card-desc {
                font-size: 12px;
                line-height: 15px;
                height: 30px;
                color: #666;
                margin-bottom: 12px;
                display: -webkit-box;
                -webkit-line-clamp: 2;
                -webkit-box-orient: vertical;
                overflow: hidden;
              }

              .card-labels {
                display: flex;
                gap: 10px;
                font-size: 12px;
                color: rgba(80, 85, 89, 1);
                overflow: hidden;

                .label-item {
                  border-radius: 3px;
                  background-color: rgba(230, 237, 245, 1);
                  height: 24px;
                  line-height: 26px;
                  padding: 0 8px;
                  flex-shrink: 0;
                  overflow: hidden;
                  text-overflow: ellipsis;
                  white-space: nowrap;

                  // 最后一个能被完整显示的标签，不要设置 flex-shrink，让它自然截断
                  &:last-child {
                    flex-shrink: 1;
                    min-width: 0;
                  }
                }

                .label-overflow {
                  display: flex;
                  border-radius: 3px;
                  background-color: rgba(0, 102, 255, 0.1);
                  color: rgba(0, 102, 255, 1);
                  height: 24px;
                  line-height: 26px;
                  padding: 0 8px;
                  flex-shrink: 0;
                  cursor: pointer;
                }
              }


            }
          }

          .card-middle {
            margin-top: 15px;
            padding: 10px 0;
            border-width: 1px 0px;
            border-color: rgba(157, 197, 226, 0.4);
            border-style: solid;
            display: flex;
            align-items: center;
            font-size: 12px;
            justify-content: space-between;
            color: #101010;
            flex-wrap: wrap;
            gap: 8px;

            .point-wrap {
              display: flex;
              align-items: center;
              gap: 6px;
              color: rgba(242, 113, 28, 1);
            }
          }

          .card-footer {
            margin-top: auto;
            padding-top: 20px;
            display: flex;
            justify-content: space-between;
            align-items: center;
            flex-wrap: wrap;
            gap: 12px;

            .footer-l {
              display: flex;
              gap: 10px;
              color: #999;
              font-size: 12px;

              .card-meta {
                display: flex;
                align-items: center;

                .avatar-c {
                  width: 16px;
                  height: 16px;

                  .avatar {
                    width: 100%;
                    height: 100%;
                    border-radius: 50%;
                  }
                }

                .update-time {
                  white-space: nowrap;
                }
              }

              .card-count {
                display: flex;
                align-items: center;
              }
            }

            .footer-r {
              display: flex;
              gap: 12px;
              align-items: center;
              margin-left: auto;

              .card-detail {
                border-bottom: 1px solid rgba(136, 136, 136, 1);
                font-size: 12px;
                display: flex;
                align-items: center;
                gap: 4px;

                a {
                  color: rgba(16, 16, 16, 1);
                }
              }

              .oper-btn {
                display: flex;
                align-items: center;
                height: 26px;
                padding: 0 10px;
                border-radius: 4px;
                font-size: 12px;
                white-space: nowrap;
                cursor: pointer;

                svg {
                  margin-right: 6px;
                }
              }

              .deploy-btn {
                background-color: rgba(230, 251, 255, 1);
                color: rgba(0, 86, 255, 1);
                border: 1px solid rgba(0, 86, 255, 1);
              }

              .run-wrap {
                display: flex;
                align-items: center;
                gap: 8px;


                // .oper-btn {
                //   border-radius: 4px;
                //   height: 28px;
                //   padding: 0 10px;
                //   display: flex;
                //   align-items: center;
                //   font-size: 12px;
                //   cursor: pointer;
                // }

                .wait-btn {
                  background-color: rgba(0, 86, 255, 0.1);
                  color: #0056ff;

                  svg {
                    animation: rotate 1s linear infinite;
                  }
                }

                .start-btn {
                  background-color: rgba(0, 86, 255, 1, 0.5);
                  color: #fff;

                  &:hover {
                    background-color: rgba(0, 86, 255, 1, 1);
                  }
                }

                .stop-btn {
                  background-color: rgba(255, 98, 0, 0.1);
                  color: #ff6200;

                  &:hover {
                    background-color: rgba(255, 98, 0, 0.2);
                  }
                }
              }
            }
          }
        }
      }

      .pagination-wrapper {
        display: flex;
        justify-content: center;
        margin-top: 24px;
        padding: 16px 0;
        overflow-x: auto;
        padding-bottom: 100px;
      }
    }


    .no-data {
      display: flex;
      justify-content: center;
      align-items: center;
      color: #999;
      font-size: 16px;
      min-height: 300px;
    }
  }
}

.labels-popover-content {
  display: flex;
  flex-wrap: wrap;
  gap: 6px;
  max-width: 280px;

  .pop-label-item {
    border-radius: 3px;
    background-color: rgba(230, 237, 245, 1);
    height: 24px;
    line-height: 26px;
    padding: 0 8px;
    font-size: 12px;
    color: rgba(80, 85, 89, 1);
    white-space: nowrap;
  }
}

.task-running-tip {
  position: fixed;
  bottom: 10px;
  left: calc(200px + (100vw - 200px) / 2);
  transform: translateX(-50%);
  z-index: 1000;
  display: flex;
  background: linear-gradient(180deg, rgba(255, 168, 0, 1) 0%, rgba(255, 190, 64, 1) 18%);
  border-radius: 8px;
  max-width: calc(100vw - 40px);
  width: max-content;
  height: auto;
  min-height: 80px;
  padding: 12px 16px;
  box-sizing: border-box;

  .circle-wrap {
    position: relative;
    width: 60px;
    height: 60px;
    background-color: #fff;
    border-radius: 100%;
    flex-shrink: 0;

    img {
      position: absolute;
      top: -34px;
      left: -17px;
      max-width: none;
    }
  }

  .task-main {
    display: flex;
    flex-direction: column;
    justify-content: center;
    line-height: 19px;
    color: rgba(16, 16, 16, 1);
    gap: 5px;
    z-index: 99;
    min-width: 0;

    .task-info {
      span {
        color: rgba(245, 34, 45, 1);
        font-weight: 700;
      }
    }

    .task-detail {
      height: 20px;

      span {
        cursor: pointer;
        border-bottom: 1px solid rgba(16, 16, 16, 0.5);
      }
    }
  }

  .close-wrap {
    width: 20px;
    height: 20px;
    background-color: rgba(255, 255, 255, 0.4);
    border-radius: 100%;
    display: flex;
    align-items: center;
    justify-content: center;
    margin-left: 12px;
    flex-shrink: 0;
    z-index: 99;
    cursor: pointer;
  }
}

@keyframes rotate {
  from {
    transform: rotate(0deg);
  }

  to {
    transform: rotate(360deg);
  }
}

@media (max-width: 1550px) {
  .content-box {
    .banner-content {
      margin-left: 20px !important;
    }

    .main-body {
      .main-content {
        .card-grid {
          grid-template-columns: repeat(2, minmax(0, 1fr));
        }
      }
    }
  }
}

@media (max-width: 1120px) {
  .content-box {
    .main-body {
      .main-content {
        .card-grid {
          grid-template-columns: repeat(1, minmax(0, 1fr));
        }
      }
    }
  }
}

@media (max-width: 992px) {
  .content-box {
    padding: 0 20px;

    .head-banner-wrap {
      height: 280px;

      .banner-content {
        margin-top: 80px;
        margin-left: 40px;

        .banner-title {
          font-size: 24px;
          line-height: 32px;
          height: 32px;
        }

        .banner-desc {
          font-size: 14px;
        }
      }

      .banner-img {
        right: 20px;
        bottom: 0;

        img {
          width: 180px;
          height: auto;
        }
      }
    }
  }
}

@media (max-width: 768px) {
  .content-box {
    padding: 0 12px;

    .head-banner-wrap {
      height: 200px;
      margin-top: 16px;

      .banner-content {
        margin-top: 40px;
        margin-left: 16px;
        max-width: 70%;

        .banner-title {
          font-size: 18px;
          line-height: 24px;
          height: 24px;
        }

        .banner-desc {
          font-size: 12px;
          margin-top: 8px;
          line-height: 16px;
        }
      }

      .banner-img {
        right: 8px;
        bottom: 0;

        img {
          width: 120px;
          height: auto;
        }
      }
    }

    .main-body {
      .filter-c {
        margin-top: 16px;
        margin-bottom: 16px;
        flex-direction: column;
        align-items: stretch;

        .tab-c {
          justify-content: flex-start;
          overflow-x: auto;
          padding-bottom: 8px;

          .tab-item {
            padding: 6px 14px;
            font-size: 14px;
          }
        }

        .right {
          width: 100%;

          .search-input-wrap {
            width: 100%;
            max-width: none;
          }
        }
      }

      .main-content {
        .pagination-wrapper {
          .el-pagination {

            /deep/ .el-pagination__total,
            /deep/ .el-pagination__sizes,
            /deep/ .el-pagination__jump {
              display: none;
            }
          }
        }

        .card-grid {
          grid-template-columns: minmax(0, 1fr);
          gap: 12px;

          .agent-card {
            padding: 14px;

            .card-footer {
              .time-1 {
                display: none;
              }
            }
          }
        }

      }
    }
  }

  .task-running-tip {
    bottom: 16px;
    left: 16px;
    right: 16px;
    transform: none;
    width: auto;
    max-width: none;
    min-width: 0;
    flex-direction: row;
    align-items: center;
    padding: 10px 12px;
    height: auto;

    .circle-wrap {
      width: 48px;
      height: 48px;
      margin-right: 12px;

      img {
        top: -28px;
        left: -14px;
        width: 60px;
        height: auto;
      }
    }

    .task-main {
      flex: 1;
      min-width: 0;

      .task-info {
        font-size: 13px;
      }

      .task-detail {
        font-size: 12px;
      }
    }

    .close-wrap {
      width: 24px;
      height: 24px;
      margin-left: 8px;
    }
  }
}
</style>