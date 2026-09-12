<template>
  <div style="height: 100%;">
    <div class="__mobile-tip">
      <img style="width: 132px;height: 97px;" src="/img/model/pc-view.png" alt="">
      <div style="margin-top: 2rem;">{{ $t('useInPcWeb') }}</div>
    </div>
    <div class="__reward-pointer-c __content-box">
      <div class="ui container">
        <div class="__r_p_header">
          <div>
            <p class="__title">{{ $t('storage.capacity_details') }}</p>
          </div>
        </div>
        <div class="__r_p_summary">
          <div class="__r_p_summary-item">
            <div class="__r_p_summary-progress __r_p-total-usage">
              <div class="__r_p-data-usage" :style="{ width: dataDisplayWidth }"></div>
              <div class="__r_p-model-usage" :style="{ left: dataDisplayWidth, width: modelDisplayWidth }"></div>
            </div>
            <div class="__r_p_summary-total">{{ $t('storage.quota') }}：{{ formattedTotal }}</div>
          </div>
          <!-- 存储信息展示 -->
          <div class="__r_p_-storage-info">
            <div class="__r_p_-storage-item">
              <span class="color-block __r_p-data-usage"></span>
              <span>{{ $t('storage.data_usaged') }}：<span class="_ehance-color">{{ formattedData }}</span> ({{
                dataPercentage }})</span>
            </div>
            <div class="__r_p_-storage-item">
              <span class="color-block __r_p-model-usage"></span>
              <span>{{ $t('storage.model_usaged') }}：<span class="_ehance-color">{{ formattedModel }}</span> ({{
                modelPercentage }})</span>
            </div>
            <div class="__r_p_-storage-item">
              <span class="color-block __r_p-total-usage"></span>
              <span>{{ $t('storage.remaining_available') }}：<span class="_ehance-color">{{ formattedRemaining }}</span>
                ({{ remainingPercentage }})</span>
            </div>
          </div>
        </div>
        <div class="__r_p_tab">
          <div class="__r_p_tab-item" :class="tabIndex === 0 ? '__focus' : ''" style="border-radius: 5px 0px 0px 5px"
            @click="tabChange(0)">
            {{ $t('dataset') }}
          </div>
          <div class="__r_p_tab-item" :class="tabIndex === 1 ? '__focus' : ''" style="border-radius: 0px 5px 5px 0px"
            @click="tabChange(1)">
            {{ $t('repos.model') }}
          </div>
        </div>
        <div class="__r_p_table">
          <div v-show="tabIndex === 0">
            <el-table ref="tableRef" :data="tableData" row-key="id" style="width: 100%;font-size: 14px;"
              v-loading="loading" v-if="tableData.length" @sort-change='sortTableFun'
              @selection-change="selectDataChange">
              <el-table-column type="selection" width="45"></el-table-column>
              <el-table-column column-key="alias" prop="alias" :label="$t('datasetObj.dataset_name1')" min-width="300">
                <template slot-scope="scope">
                  <div style="display: flex;align-items: center;">
                    <a style="margin-right:8px" class="nowrap"
                      :href="`/datasets/detail/${scope.row.owner_name}/${scope.row.name}`" target="_blank">{{
                        scope.row.alias || scope.row.name }}</a>
                    <svg v-if="scope.row.recommend" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24" width="16"
                      height="16">
                      <defs></defs>
                      <g>
                        <path fill="#ff6200"
                          d="M11.99 2C6.47 2 2 6.48 2 12s4.47 10 9.99 10C17.52 22 22 17.52 22 12S17.52 2 11.99 2zm4.24 16L12 15.45 7.77 18l1.12-4.81-3.73-3.23 4.92-.42L12 5l1.92 4.53 4.92.42-3.73 3.23L16.23 18z">
                        </path>
                      </g>
                    </svg>
                  </div>
                </template>
              </el-table-column>
              <el-table-column column-key="size" prop="size" sortable='custom' :label="$t('modelManage.fileSize')"
                width="180">
                <template slot-scope="scope">{{ formatBytes(scope.row.size) }}</template>
              </el-table-column>
              <el-table-column column-key="is_private" prop="is_private" :label="$t('status')" width="160"
                align="center" header-align="center">
                <template slot-scope="scope">
                  <div style="display:flex;justify-content: center;align-items: center;width: 100%;"
                    v-if="scope.row.is_private">
                    <i class="ri-lock-2-line" style="color: #fa8c16;margin-left:6px;"></i>
                    <span style="color: #ff9959;margin-left:8px">{{ $t('modelManage.modelAccessPrivate') }}</span>
                  </div>
                  <div style="display:flex;justify-content: center;align-items: center;width: 100%;" v-else>
                    <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 32 32" width="16" height="16">
                      <defs></defs>
                      <g>
                        <path fill="#27b148"
                          d="M29.696 14.464v15.040c0 0.288-0.224 0.48-0.48 0.48h-17.28c-0.288 0-0.512-0.192-0.512-0.48v-15.040c0-0.256 0.224-0.48 0.512-0.48h1.344c0.288 0 0.512-0.096 0.512-0.256v-0.224c-0.032-1.6 0-3.2-0.032-4.8-0.096-2.112-2.016-3.936-4.128-4.064-2.208-0.096-4.16 1.408-4.608 3.552-0.061 0.331-0.096 0.711-0.096 1.1 0 0.007 0 0.014 0 0.021v-0.001 4.416s-0.224 0.16-0.48 0.16h-1.664c-0.256 0-0.48-0.224-0.48-0.512 0-1.632-0.032-3.264 0.032-4.864 0.16-2.976 2.336-5.504 5.216-6.272 4.032-1.088 8.16 1.6 8.768 5.76 0.096 0.672 0.096 1.408 0.096 2.080v3.68s0.224 0.224 0.48 0.224h12.32c0.256 0 0.48 0.224 0.48 0.48z">
                        </path>
                      </g>
                    </svg>
                    <span style="color: #27b128;margin-left:8px">{{ $t('modelManage.modelAccessPublic') }}</span>
                  </div>
                </template>
              </el-table-column>
              <el-table-column column-key="created_unix" prop="created_unix" :label="$t('modelManage.updateTime')"
                align="center" header-align="center" width="200">
                <template slot-scope="scope">
                  <span>{{ dateFormat(scope.row.updated_unix) }}</span>
                </template>
              </el-table-column>
              <el-table-column prop="operation" :label="$t('operation')" align="center" width="200" fixed="right"
                header-align="center">
                <template slot-scope="scope">
                  <div class="op-wrap">
                    <!-- <a href="javascript:;" @click="goToPreview(scope.row)">{{ $t('modelManage.download') }}</a> -->
                    <a href="javascript:;" style="color: red;" @click="deleteDataSet(scope.row)">{{
                      $t('cloudbrainObj.delete') }}</a>
                  </div>
                </template>
              </el-table-column>
              <template slot="empty">
                <span>{{ loading ? $t('loading') : $t('noData') }}</span>
              </template>
            </el-table>
            <el-empty v-else :image-size="140" :description="$t('modelObj.model_square_empty')"></el-empty>
          </div>
          <div v-show="tabIndex === 1">
            <el-table ref="tableModelRef" :data="tableModel" row-key="id" style="width: 100%;font-size: 14px;"
              v-loading="loading" stripe v-if="tableModel.length" @sort-change='sortTableFun'
              @selection-change="selectModelChange">
              <el-table-column type="selection" width="45"></el-table-column>
              <el-table-column column-key="alias" prop="alias" :label="$t('modelManage.modelName')" min-width="300">
                <template slot-scope="scope">
                  <div style="display: flex;align-items: center;">
                    <a style="margin-right:8px" class="nowrap"
                      :href="`/models/detail/${scope.row.owner_name}/${scope.row.name}`" target="_blank">{{
                        scope.row.alias || scope.row.name }}</a>
                    <svg v-if="scope.row.recommend" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24" width="16"
                      height="16">
                      <defs></defs>
                      <g>
                        <path fill="#ff6200"
                          d="M11.99 2C6.47 2 2 6.48 2 12s4.47 10 9.99 10C17.52 22 22 17.52 22 12S17.52 2 11.99 2zm4.24 16L12 15.45 7.77 18l1.12-4.81-3.73-3.23 4.92-.42L12 5l1.92 4.53 4.92.42-3.73 3.23L16.23 18z">
                        </path>
                      </g>
                    </svg>
                  </div>
                </template>
              </el-table-column>
              <el-table-column column-key="size" prop="size" sortable='custom' :label="$t('modelManage.fileSize')"
                width="180">
                <template slot-scope="scope">{{ formatBytes(scope.row.size) }}</template>
              </el-table-column>
              <el-table-column column-key="is_private" prop="is_private" :label="$t('status')" width="160"
                align="center" header-align="center">
                <template slot-scope="scope">
                  <div style="display:flex;justify-content: center;align-items: center;width: 100%;"
                    v-if="scope.row.is_private">
                    <i class="ri-lock-2-line" style="color: #fa8c16;margin-left:6px;"></i>
                    <span style="color: #ff9959;margin-left:8px">{{ $t('modelManage.modelAccessPrivate') }}</span>
                  </div>
                  <div style="display:flex;justify-content: center;align-items: center;width: 100%;" v-else>
                    <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 32 32" width="16" height="16">
                      <defs></defs>
                      <g>
                        <path fill="#27b148"
                          d="M29.696 14.464v15.040c0 0.288-0.224 0.48-0.48 0.48h-17.28c-0.288 0-0.512-0.192-0.512-0.48v-15.040c0-0.256 0.224-0.48 0.512-0.48h1.344c0.288 0 0.512-0.096 0.512-0.256v-0.224c-0.032-1.6 0-3.2-0.032-4.8-0.096-2.112-2.016-3.936-4.128-4.064-2.208-0.096-4.16 1.408-4.608 3.552-0.061 0.331-0.096 0.711-0.096 1.1 0 0.007 0 0.014 0 0.021v-0.001 4.416s-0.224 0.16-0.48 0.16h-1.664c-0.256 0-0.48-0.224-0.48-0.512 0-1.632-0.032-3.264 0.032-4.864 0.16-2.976 2.336-5.504 5.216-6.272 4.032-1.088 8.16 1.6 8.768 5.76 0.096 0.672 0.096 1.408 0.096 2.080v3.68s0.224 0.224 0.48 0.224h12.32c0.256 0 0.48 0.224 0.48 0.48z">
                        </path>
                      </g>
                    </svg>
                    <span style="color: #27b128;margin-left:8px">{{ $t('modelManage.modelAccessPublic') }}</span>
                  </div>
                </template>
              </el-table-column>
              <el-table-column column-key="updated_unix" prop="updated_unix" :label="$t('modelManage.updateTime')"
                align="center" header-align="center" width="200">
                <template slot-scope="scope">
                  <span>{{ dateFormat(scope.row.updated_unix) }}</span>
                </template>
              </el-table-column>
              <el-table-column prop="operation" :label="$t('operation')" align="center" min-width="80" fixed="right"
                header-align="center">
                <template slot-scope="scope">
                  <div class="op-wrap">
                    <a href="javascript:;" style="color: red;" @click="deleteModel(scope.row)">{{
                      $t('cloudbrainObj.delete') }}</a>
                  </div>
                </template>
              </el-table-column>
              <template slot="empty">
                <span>{{ loading ? $t('loading') : $t('noData') }}</span>
              </template>
            </el-table>
            <el-empty v-else :image-size="140" :description="$t('modelObj.model_square_empty')"></el-empty>
          </div>
          <div class="__r_p_pagination" v-if="showFooter">
            <div style="margin-top: 2rem">
              <div class="list-foot">
                <el-button class="left-btn" plain icon="el-icon-delete" @click="batchDelete" size="small">{{
                  $t('cloudbrainObj.batchDelete') }}</el-button>
                <el-pagination class="center-btn" background @current-change="currentChange"
                  :current-page="pageInfo.curpage" :page-sizes="pageInfo.pageSizes" :page-size="pageInfo.pageSize"
                  layout="total, sizes, prev, pager, next, jumper" :total="pageInfo.total">
                </el-pagination>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
    <!-- 删除确认组件 -->
    <delete-confirm-modal ref="deleteModal" :type="dataType" :data-obj="dataObj" @confirm-delete="handleDeleteModel"
      @delete-success="handleDeleteSuccess" />
    <LoadingMask :loading="maskLoading" :content="maskLoadingContent"></LoadingMask>
  </div>
</template>

<script>
import LoadingMask from '~/components/cloudbrain/LoadingMask.vue';
import DeleteConfirmModal from '~/components/square/DeleteConfirmModal.vue'
import dayjs from 'dayjs';
import { getStorageSummary, getStorageDataset, delStorageModel, batchDelStorageDataset, batchDelStorageModel } from "~/apis/modules/storage";
import { getOrgStorageDatasetList, getOrgSelectedModelList, getOrgStorageSummary } from '~/apis/modules/organization';
import { delDataset } from "~/apis/modules/dataset";
import { initClipboard } from '~/utils';
const { csrf } = window.config;
const UNITS = ['Bytes', 'KiB', 'MiB', 'GiB', 'TiB'];

export default {
  props: {
    type: {
      type: String,
      default: 'user'
    },
    orgName: {
      type: String,
      default: ''
    }
  },
  data() {
    return {
      loading: false,
      tabIndex: 0,
      tableData: [],
      tableModel: [],
      pageInfo: {
        curpage: 1,
        pageSize: 12,
        pageSizes: [12],
        total: 0,
      },
      orderByData: 'default',
      orderByModel: 'default',

      totalStorage: 0,   // 总容量
      dataUsage: 0,     // 数据用量
      modelUsage: 0,
      remainingUsage: 0,
      userStorage: 0,
      multipleSelection: [],
      maskLoading: false,
      maskLoadingContent: '',
      dataObj: {},
      dataType: 'datasetObj',
    };
  },
  components: { LoadingMask, DeleteConfirmModal },
  computed: {
    // remainingBytes() {
    //     return Math.max(this.totalStorage - (this.dataUsage + this.modelUsage), 0);
    // },
    // 单位转换方法
    formattedData() {
      return this.formatBytes(this.dataUsage);
    },
    formattedModel() {
      return this.formatBytes(this.modelUsage);
    },
    // 格式化的剩余空间显示（确保不超过 totalStorage）
    formattedRemaining() {
      return this.formatBytes(Math.max(this.remainingUsage, 0));
    },
    formattedTotal() {
      let unitIndex = 0;
      let value = this.totalStorage;

      while (value >= 1024 && unitIndex < UNITS.length - 1) {
        value /= 1024;
        unitIndex++;
      }

      return `${value} ${UNITS[unitIndex]}`;
    },

    // 状态判断
    isOverflow() {
      return this.remainingUsage <= 0;
    },
    // 进度条计算
    dataDisplayWidth() {
      if (this.isOverflow) {
        const ratio = this.dataUsage / this.userStorage * 100;
        return `${this.toPrecision(ratio, 3)}%`;
      }
      const ratio = this.dataUsage / this.totalStorage * 100;
      return `${this.toPrecision(ratio, 3)}%`;
    },
    modelDisplayWidth() {
      if (this.isOverflow) {
        const ratio = this.modelUsage / this.userStorage * 100;
        return `${this.toPrecision(ratio, 3)}%`;
      }
      const ratio = this.modelUsage / this.totalStorage * 100;
      return `${this.toPrecision(ratio, 3)}%`;
    },

    // 百分比显示
    dataPercentage() {
      if (this.dataUsage === this.totalStorage || this.totalStorage === 0) return '100%'
      return `${this.toPrecision((this.dataUsage / this.totalStorage) * 100, 2)}%`
    },
    modelPercentage() {
      if (this.modelUsage === this.totalStorage || this.totalStorage === 0) return '100%'
      return `${this.toPrecision((this.modelUsage / this.totalStorage) * 100, 2)}%`
    },
    // 剩余空间百分比（带精度处理）
    remainingPercentage() {
      if (this.isOverflow || this.remainingUsage === 0) {
        return '0%'; // 超限时强制显示0%
      }
      if (this.remainingUsage === this.totalStorage) return '100%'
      if ((this.userStorage / this.totalStorage) * 100 < 0.01) return '99.99%'
      if ((this.remainingUsage / this.totalStorage) * 100 < 0.01) return '≤0.01%'

      const dataPct = this.dataUsage / this.totalStorage * 100;
      const modelPct = this.modelUsage / this.totalStorage * 100;
      const roundedDataStr = this.toPrecision(dataPct, 2);
      const roundedModelStr = this.toPrecision(modelPct, 2);
      // 处理 "<0.01" 的情况，将其视为 0
      const roundedData = roundedDataStr === '≤0.01' ? 0 : Number(roundedDataStr);
      const roundedModel = roundedModelStr === '≤0.01' ? 0 : Number(roundedModelStr);
      const roundedRemaining = 100 - roundedData - roundedModel
      return `${roundedRemaining.toFixed(2)}%`;
    },
    showFooter() {
      if (this.tabIndex === 0) {
        if (this.tableData.length > 0) return true;
      } else {
        if (this.tableModel.length > 0) return true;
      }
      return false;
    }

  },
  methods: {
    selectDataChange(taskList) {
      this.multipleSelection = []
      taskList.forEach((item) => {
        this.multipleSelection.push(item.id)
      })
    },
    selectModelChange(taskList) {
      this.multipleSelection = []
      taskList.forEach((item) => {
        this.multipleSelection.push(item.id)
      })
    },

    batchDelete() {
      if (this.multipleSelection.length === 0) return;
      const confirmTitle = this.$t('tips');
      const confirmMessage = this.tabIndex === 0 ? this.$t('storage.deleteBatchData') : this.$t('storage.deleteBatchModel')
      this.$confirm(confirmMessage, confirmTitle, {
        confirmButtonText: this.$t('confirm'),
        cancelButtonText: this.$t('cancel'),
        type: 'warning',
        lockScroll: false,
      }).then(() => {
        this.maskLoading = true;
        this.maskLoadingContent = this.$t('storage.deletingTips');
        const isDataSetTab = this.tabIndex === 0;
        const params = isDataSetTab
          ? { dataset_ids: this.multipleSelection.join(',') }
          : { aimodel_ids: this.multipleSelection.join(',') };
        const apiCall = isDataSetTab
          ? batchDelStorageDataset(params, 'dataset')
          : batchDelStorageDataset(params, 'aimodel');
        apiCall
          .then((res) => {
            console.log("res", res.data)
            const deletedCount = res.data.data.success?.length || 0
            if (this.shouldDecrementPage(deletedCount)) {
              this.pageInfo.curpage -= 1;
            }
            this.getSummaryInfo();
            isDataSetTab ? this.getTableData() : this.getTableData('aimodel');
            // Reset loading and selection
            this.maskLoading = false;
            this.multipleSelection = [];
          })
          .catch((err) => {
            this.handleDeleteError(err);
          });
      }).catch(() => {
        this.$message({
          type: 'info',
          message: this.$t('cancelOperate'),
        });
      });
    },

    /**
     * 统一错误处理
     */
    handleDeleteError(error) {
      this.maskLoading = false;
      this.$message.error(error?.response?.data?.message || error || this.$t('operationFailed'));
    },
    shouldDecrementPage(deletedCount) {
      const newTotal = this.pageInfo.total - deletedCount;
      const newTotalPages = Math.ceil(newTotal / this.pageInfo.pageSize);
      return this.pageInfo.curpage > newTotalPages && this.pageInfo.curpage > 1;
    },
    sortTableFun(column) {
      const orderMapping = {
        ascending: 'size_asc',
        descending: 'size'
      }
      if (this.tabIndex === 0) {
        this.orderByData = orderMapping[column.order] || 'default';
        this.getTableData()
      } else {
        this.orderByModel = orderMapping[column.order] || 'default';
        this.getTableData('aimodel')
      }
    },
    currentChange: function (val) {
      this.pageInfo.curpage = val;
      if (this.tabIndex === 0) {
        this.getTableData();
      } else {
        this.getTableData('aimodel')
      }

    },
    tabChange: function (index) {
      if (this.tabIndex === index) return;
      this.multipleSelection = []
      this.tabIndex = index;
      this.pageInfo.curpage = 1;
      this.pageInfo.total = 0;
      if (index === 0) {
        this.getTableData();
        this.$nextTick(() => {
          if (this.$refs.tableRef) {
            this.$refs.tableRef.clearSelection();
          }
        })
      } else {
        this.getTableData('aimodel')
        this.$nextTick(() => {
          if (this.$refs.tableModelRef) {
            this.$refs.tableModelRef.clearSelection();
          }
        })
      }

    },
    dateFormat(unix) {
      return dayjs(unix * 1000).format('YYYY-MM-DD HH:mm:ss');
    },
    // 精确小数处理
    toPrecision1(value, decimals = 2) {
      // 处理极小正值的显示
      if (value > 0 && value < 0.01) {
        return '≤0.01';  // 统一极小值显示为 <0.01%
      }

      // 处理四舍五入到 decimals 位后的边界
      const rounded = Number(value.toFixed(decimals));

      // 避免 100% 溢出（如 99.999 四舍五入为 100）
      if (rounded === 100 && !this.isOverflow) {
        return (100 - Math.pow(10, -decimals)).toFixed(decimals); // 返回 99.99
      }
      // 正常情况返回四舍五入值
      return rounded.toFixed(decimals);
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
      let totalInSameUnit = this.totalStorage;  // 总存储的初始单位（字节）
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
        return `${truncated.toFixed(decimals)} ${UNITS[unitIndex]}`;
      } else {
        return `${roundedValue.toFixed(decimals)} ${UNITS[unitIndex]}`;
      }

    },
    goToPreview(row) {
      //datasets/dirs/{{.UUID}}?type={{$.Type}}
      let url = `/${row.repo_owner}/${row.repo_name}/datasets/dirs/${row.uuid}?type=-1`
      window.open(url, '_blank')
    },
    getFeedUrl(row) {
      let url = 'https://openi.pcl.ac.cn/zeizei/OpenI_Learning/issues/new'
      let issuetitle = `${this.$t('modelManage.datasetfile')}：${row.name} ${this.$t('datasetObj.dataset_unzip_failed')}`
      let body = `${this.$t('modelManage.project')}：${row.repo_owner}/${row.repo_name} ${this.$t('modelManage.datasetfile')}：${row.name}`
      return `${url}?issuetitle=${issuetitle}&body=${body}`
    },
    deleteDataSet(row) {
      // let name = row.alias || row.name
      // this.$confirm(this.$t('storage.deleteDataSetConfirm',{name:name}), this.$t('tips'), {
      //     confirmButtonText: this.$t('confirm'),
      //     cancelButtonText: this.$t('cancel'),
      //     dangerouslyUseHTMLString: true,
      //     type: 'warning',
      //     lockScroll: false,
      // }).then(() => {
      //     delDataset({dataset_id: row.id}).then((res)=>{
      //         this.$message.success(this.$t('storage.deleteDataSetSuccess',{name:row.name}))
      //         this.getSummaryInfo()
      //         if (this.shouldDecrementPage(1)) {
      //           this.pageInfo.curpage -= 1;
      //         }
      //         this.getTableData()
      //     }).catch((err)=>{
      //         this.$message.error(err)
      //     })
      // }).catch((err) => {
      //     this.$message({
      //       type: 'info',
      //       message: this.$t('cancelOperate'),
      //     });
      // });
      this.dataObj = row
      this.dataType = 'datasetObj'
      this.$refs.deleteModal.showModal()
    },
    deleteModel(row) {
      // let name = row.alias || row.name
      // this.$confirm(this.$t('storage.deleteModelConfirm',{name:name}), this.$t('tips'), {
      //     confirmButtonText: this.$t('confirm'),
      //     cancelButtonText: this.$t('cancel'),
      //     dangerouslyUseHTMLString: true,
      //     type: 'warning',
      //     lockScroll: false,
      // }).then(() => {
      //     delDataset({aimodel_id: row.id},'aimodel').then((res)=>{
      //         this.$message.success(this.$t('storage.deleteModelConSuccess',{name:name}))
      //         this.getSummaryInfo();
      //         if (this.shouldDecrementPage(1)) {
      //           this.pageInfo.curpage -= 1;
      //         }
      //         this.getTableData('aimodel')
      //     }).catch((err)=>{
      //         this.$message.error(err)
      //     })
      // }).catch(() => {
      //     this.$message({
      //         type: 'info',
      //         message: this.$t('cancelOperate'),
      //     });
      // });
      this.dataObj = row
      this.dataType = 'modelObj'
      this.$refs.deleteModal.showModal()
    },
    async handleDeleteModel({ type, data, callback }) {
      try {
        let response
        if (type === 'datasetObj') {
          response = await delDataset({ dataset_id: data.id })
        } else {
          response = await delDataset({ aimodel_id: data.id }, 'aimodel')
        }
        if (response.data.code === 0) {
          callback({
            success: true,
            message: this.$t('imagesObj.deleteSuccessTips')
          })
        } else {
          callback({
            success: false,
            error: response.data.msg
          })
        }
      } catch (error) {
        callback({
          success: false,
          error: error.message
        })
      }
    },
    // 删除成功后的回调
    handleDeleteSuccess() {
      // this.getCardList() // 刷新列表
      this.getSummaryInfo()
      if (this.shouldDecrementPage(1)) {
        this.pageInfo.curpage -= 1;
      }
      if (this.dataType === 'datasetObj') {
        this.getTableData();
      } else {
        this.getTableData('aimodel')
      }
    },
    getSummaryInfo() {
      const queryParams = this.type === 'user'
        ? getStorageSummary({})
        : getOrgStorageSummary(this.orgName, {});
      queryParams.then(res => {
        let data = this.type === 'user' ? res.data : res.data.data
        console.log("xxx", data)
        this.dataUsage = data.dataset_used_storage
        this.modelUsage = data.model_used_storage
        this.totalStorage = data.storage_limit
        this.remainingUsage = data.remaining_storage
        this.userStorage = data.used_storage
      }).catch(err => {
        this.$message.error(err)
      })
    },
    getTableData(type = 'dataset') {
      this.loading = true;
      const baseParams = {
        page: this.pageInfo.curpage,
        page_size: this.pageInfo.pageSize,
        order_by: type === 'dataset' ? this.orderByData : this.orderByModel
      }
      const getDatasetQuery = this.type === 'user' ? getStorageDataset : getOrgStorageDatasetList
      const params = this.type === 'user' ? baseParams : { owner_name: this.orgName, ...baseParams }
      getDatasetQuery(params, type).then((res) => {
        this.loading = false;
        const tableData = [];
        let records = []
        const data = res.data;
        if (data.code === 0) {
          if (type === 'dataset') {
            records = data.data.datasets
          } else {
            records = data.data.aimodels
          }
          records && records.forEach(item => {
            tableData.push(item);
          })
          if (type === 'dataset') {
            this.tableData.splice(0, Infinity, ...tableData);
          } else {
            this.tableModel.splice(0, Infinity, ...tableData);
          }
          this.pageInfo.total = data.data.total;
        } else {
          this.$message.error(data.msg);
        }
        this.$nextTick(() => {
          initClipboard('.clipboard-model-name');
        });
      }).catch((err) => {
        this.$message.error(err)
        this.loading = false;
        this.tableData.splice(0, Infinity);
      }).finally(() => {
        this.loading = false
      })
    },
  },
  mounted: function () {
    this.getSummaryInfo();
    this.getTableData()

  },
  beforeDestroy: function () {
    this.multipleSelection = [];
  },
};
</script>

<style scoped lang="less">
.__flex-1 {
  flex: 1;
}

.__reward-pointer-c {
  .__r_p_header {
    height: 30px;
    margin: 10px 0;
    display: flex;
    align-items: center;
    justify-content: space-between;

    .__title {
      font-weight: 400;
      font-size: 18px;
      color: rgb(16, 16, 16);
      line-height: 26px;
    }
  }

  .__r_p_summary {
    display: flex;
    flex-direction: column;
    justify-content: center;
    height: 100px;
    background-color: rgb(245, 245, 246);
    width: 100%;
    padding: 0 26px;

    .__r_p_summary-item {
      display: flex;
      align-items: center;
      margin-bottom: 14px;

      .__r_p_summary-progress {
        position: relative;
        width: 80%;
        height: 15px;
        transition: border-color 0.3s ease;
        overflow: hidden;

        /* 防止子元素溢出 */
        .__r_p-data-usage,
        .__r_p-model-usage {
          position: absolute;
          height: 100%;
          top: 0;
          transition: all 0.3s ease;
        }

      }

      .__r_p_summary-total {
        margin-left: 15px;
      }
    }

    .__r_p_-storage-info {
      display: flex;

      .__r_p_-storage-item {
        margin-right: 30px;
        color: #101010;
        display: flex;
        align-items: center;

        .color-block {
          width: 12px;
          height: 12px;
          display: inline-block;
          margin-right: 6px;
        }

        ._ehance-color {
          font-weight: 600;
        }
      }
    }

    .__r_p-total-usage {
      background-color: rgba(251, 251, 251, 1);
      border: 1px solid rgba(187, 187, 187, 1);
    }

    .__r_p-data-usage {
      background-color: rgba(22, 166, 252, 1);
    }

    .__r_p-model-usage {
      background-color: rgba(0, 215, 149, 1);

    }

  }
}



.__r_p_tab {
  display: flex;
  margin: 18px 0;

  .__r_p_tab-item {
    width: 115px;
    height: 38px;
    display: flex;
    justify-content: center;
    align-items: center;
    border: 1px solid rgb(225, 227, 230);
    color: #101010;
    box-sizing: border-box;
    cursor: pointer;

    &.__focus {
      border-color: rgb(50, 145, 248);
      color: rgb(50, 145, 248);
      cursor: default;
    }
  }
}

.__r_p_table {
  /deep/ .el-table__header {
    th {
      background: rgb(245, 245, 246);
      color: rgb(96, 98, 102);
      font-weight: 400;
      font-size: 14px;
    }
  }

  .diy-popper {
    max-width: 400px;
  }

  .dataset-name-wrap {
    display: flex;
    align-items: center;


  }

  .__item-name {
    display: flex;
    align-items: center;

    span {
      margin: 0 6px;
    }

  }

  .clipboard-model-name {
    margin-left: 5px;
    color: #919191;
    cursor: pointer;
  }

  .op-wrap {
    display: flex;
    align-items: center;
    justify-content: center;

    a {
      margin: 0 10px;
      font-size: 14px;
      color: #1678c2;

      &.disabled {
        color: rgba(0, 0, 0, .87);
      }
    }
  }

}

.__r_p_pagination {
  .list-foot {
    display: flex;
    justify-content: space-between;
    /* 初始设置为两端对齐 */
    align-items: center;
    /* 垂直方向居中 */
    margin-top: 20px;

    .left-btn {
      // margin-right: auto;
      border-color: #ff4d4f;
      color: #ff4d4f;
    }

    .center-btn {
      margin-right: auto;
      margin-left: auto;
    }
  }
}

.datset-status {
  display: flex;
  align-items: center;

  .down-wrap {
    display: flex;
    align-items: center;
    margin-left: 8px;

    i {
      font-size: 10px;
    }
  }

  a {
    flex-shrink: 0;
    color: rgba(254, 221, 89, 1);
    margin-left: 8px;
  }
}
</style>
