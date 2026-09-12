<template>
<div style="height: 100%;">  
  <div class="__mobile-tip" >
      <img style="width: 132px;height: 97px;" src="/img/model/pc-view.png" alt="">
      <div style="margin-top: 2rem;">{{$t('useInPcWeb')}}</div>
  </div>
  <div class="__reward-pointer-c __content-box">
    <div class="ui container" style="width:80%;min-width:1200px;">
      <div class="__r_p_header">
        <div>
          <p class="__title">{{ $t('calcPointDetails') }}</p>
        </div>
        <div style="padding: 0 5px; font-size: 14px">
          <span>
            <i class="question circle icon link" style="color: rgba(3, 102, 214, 1)" data-position="right center"
              data-variation="mini"></i>
            <a href="/reward/point/rule" target="_blank" style="color: rgba(3, 102, 214, 1)">{{
              $t('calcPointAcquisitionInstructions')
              }}</a>
          </span>
        </div>
      </div>
      <div class="__r_p_summary">
        <div class="__r_p_summary_item-c __flex-1" style="position:relative;">
          <div class="__val">{{ summaryInfo.available.toFixed(2) }}</div>
          <div class="__exp">{{ $t('CurrAvailableCalcPoints') }}</div>
          <div class="__r_p_summary_line"></div>
        </div>
        <div class="__r_p_summary_item-c __flex-1">
          <div class="__val">{{ summaryInfo.gain.toFixed(2) }}</div>
          <div class="__exp">{{ $t('totalGainCalcPoints') }}</div>
        </div>
        <div class="__r_p_summary_item-c __flex-1">
          <div class="__val">{{ summaryInfo.used.toFixed(2) }}</div>
          <div class="__exp">{{ $t('totalConsumeCalcPoints') }}</div>
        </div>
      </div>
      <div class="__r_p_tab">
        <div class="__r_p_tab-item" :class="tabIndex === 0 ? '__focus' : ''" style="border-radius: 5px 0px 0px 5px"
          @click="tabChange(0)">
          {{ $t('gainDetail') }}
        </div>
        <div class="__r_p_tab-item" :class="tabIndex === 1 ? '__focus' : ''" style="border-radius: 0px 5px 5px 0px"
          @click="tabChange(1)">
          {{ $t('consumeDetail') }}
        </div>
      </div>
      <div class="__r_p_table">
        <div v-show="tabIndex === 0">
          <el-table :data="tableData" row-key="sn" style="width: 100%" v-loading="loading" stripe
            v-if="tableData.length">
            <el-table-column column-key="sn" prop="sn" :label="$t('serialNumber')" align="center" header-align="center"
              width="180">
            </el-table-column>
            <el-table-column column-key="date" prop="date" :label="$t('time')" align="center" header-align="center"
              width="180">
            </el-table-column>
            <el-table-column column-key="sourceType" prop="sourceType" :label="$t('scene')" align="center"
              header-align="center" width="180"></el-table-column>
            <el-table-column column-key="action" prop="action" :label="$t('behaviorOfPoint')" align="center"
              header-align="center" width="200"></el-table-column>
            <el-table-column column-key="remark" prop="remark" :label="$t('explanation')" align="left" min-width="200"
              header-align="center">
              <template slot-scope="scope">
                <span v-html="scope.row.remark"></span>
              </template>
            </el-table-column>
            <el-table-column column-key="amount" prop="amount" :label="$t('points')" align="center"
              header-align="center" width="120">
              <template slot-scope="scope">
                <span>{{ scope.row.amount.toFixed(2) }}</span>
              </template>
            </el-table-column>
            <template slot="empty">
              <span>{{ loading ? $t('loading') : $t('noData') }}</span>
            </template>
          </el-table>
          <el-empty v-else :image-size="140" :description="$t('noPointGainRecord')"></el-empty>
        </div>
        <div v-show="tabIndex === 1">
          <div class="__r_p_consume_descr" v-html="$t('consumeDescr')"></div>
          <el-table :data="tableData" row-key="sn" style="width: 100%" v-loading="loading" stripe
            v-if="tableData.length">
            <el-table-column column-key="sn" prop="sn" :label="$t('serialNumber')" align="center" header-align="center"
              width="180">
            </el-table-column>
            <el-table-column column-key="date" prop="date" :label="$t('time')" align="center" header-align="center"
              width="180">
            </el-table-column>
            <el-table-column column-key="status" prop="status" :label="$t('status')" align="center"
              header-align="center" width="120">
              <template slot-scope="scope">
                <span :style="{ color: scope.row.statusColor }">{{ scope.row.status }}</span>
              </template>
            </el-table-column>
            <el-table-column column-key="sourceType" prop="sourceType" :label="$t('scene')" align="center"
              header-align="center" width="180"></el-table-column>
            <el-table-column column-key="duration" prop="duration" :label="$t('runTime')" align="center"
              header-align="center" width="120"></el-table-column>
            <el-table-column column-key="remark" prop="remark" :label="$t('explanation')" align="left" min-width="200"
              header-align="center">
              <template slot-scope="scope">
                <span v-html="scope.row.remark"></span>
              </template>
            </el-table-column>
            <el-table-column column-key="taskName" prop="taskName" :label="$t('taskName')" align="center"
              header-align="center" width="180">
              <template slot-scope="scope">
                <span v-html="scope.row.taskName"></span>
              </template>
            </el-table-column>
            <el-table-column column-key="amount" prop="amount" :label="$t('points')" align="center"
              header-align="center" width="120">
              <template slot-scope="scope">
                <span>{{ scope.row.amount.toFixed(2) }}</span>
              </template>
            </el-table-column>
            <template slot="empty">
              <span>{{ loading ? $t('loading') : $t('noData') }}</span>
            </template>
          </el-table>
          <el-empty v-else :image-size="140" :description="$t('noPointConsumeRecord')"></el-empty>
        </div>
        <div class="__r_p_pagination" v-if="tableData.length">
          <div style="margin-top: 2rem">
            <div class="center">
              <el-pagination background @current-change="currentChange" :current-page="pageInfo.curpage"
                :page-sizes="pageInfo.pageSizes" :page-size="pageInfo.pageSize"
                layout="total, sizes, prev, pager, next, jumper" :total="pageInfo.total">
              </el-pagination>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</div>
</template>

<script>
import { getPoint, getPointAccount, getPointList } from "~/apis/modules/point";
import { getRewardPointRecordInfo } from './utils';
export default {
  data() {
    return {
      loading: false,
      summaryInfo: {
        available: 0,
        gain: 0,
        used: 0,
      },
      tabIndex: 0,
      tableData: [],
      pageInfo: {
        curpage: 1,
        pageSize: 10,
        pageSizes: [10],
        total: 0,
      },
      worker: null,
    };
  },
  components: {},
  methods: {
    currentChange: function (val) {
      this.pageInfo.curpage = val;
      this.getTableData();
    },
    tabChange: function (index) {
      if (this.tabIndex === index) return;
      this.tabIndex = index;
      this.pageInfo.curpage = 1;
      this.pageInfo.total = 0;
      this.getTableData();
      this.getSummaryInfo();
    },
    getSummaryInfo: function () {
      getPointAccount().then(res => {
        if (res.data && res.data.Code === 0) {
          const data = res.data.Data;
          this.summaryInfo.available = data.Balance;
          this.summaryInfo.gain = data.TotalEarned;
          this.summaryInfo.used = data.TotalConsumed;
        }
      }).catch(err => {
        console.log(err);
      })
    },
    getTableData: function () {
      this.loading = true;
      getPointList({
        Operate: this.tabIndex === 0 ? 'INCREASE' : 'DECREASE',
        Page: this.pageInfo.curpage,
        // pageSize: this.pageInfo.pageSize,
      }).then((res) => {
        this.loading = false;
        const tableData = [];
        if (res.data && res.data.Code === 0) {
          const data = res.data.Data;
          const records = data.Records;
          for (let i = 0, iLen = records.length; i < iLen; i++) {
            const record = records[i];
            tableData.push(getRewardPointRecordInfo(record));
          }
          this.tableData.splice(0, Infinity, ...tableData);
          this.pageInfo.total = data.Total;
        }
      })
        .catch((err) => {
          console.log(err);
          this.loading = false;
          this.tableData.splice(0, Infinity);
        });
    },
  },
  mounted: function () {
    this.getSummaryInfo();
    this.getTableData();
    const { AppSubUrl, AppVer, NotificationSettings } = window.config;
    if (NotificationSettings.EventSourceUpdateTime > 0 && !!window.EventSource && !!window.SharedWorker) {
      const worker = new SharedWorker(`${AppSubUrl}/static/eventsource.sharedworker.js?v=${AppVer}`, 'notification-worker');
      worker.port.postMessage({
        type: 'start',
        url: `${window.location.origin}${AppSubUrl}/user/events`,
      });
      worker.port.addEventListener('message', async (event) => {
        if (!event.data || !event.data.type) {
          console.error('unknown worker message event', event);
          return;
        }
        if (event.data.type == 'reward-operation') {
          try {
            this.getSummaryInfo();
            this.getTableData();
          } catch (err) {
            console.error(err);
          }
        }
      });
      worker.port.start();
      this.worker = worker;
    }
  },
  beforeDestroy: function () {
    if (this.worker) {
      worker.port.postMessage({
        type: 'close',
      });
      worker.port.close();
    }
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
    align-items: center;
    height: 100px;
    background-color: rgb(245, 245, 246);

    .__r_p_summary_item-c {
      .__val {
        text-align: center;
        margin: 12px 0;
        font-weight: 400;
        font-size: 28px;
        color: rgb(16, 16, 16);
      }

      .__exp {
        text-align: center;
        font-weight: 400;
        font-size: 14px;
        color: rgba(54, 56, 64, 1);
      }
    }
  }
}

.__r_p_summary_line {
  position: absolute;
  top: 0;
  right: 1px;
  width: 1px;
  height: 100%;
  border-right: 1px solid rgb(212, 212, 213);
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

  .__r_p_consume_descr {
    margin: -5px 0 10px 0;
  }
}
</style>
