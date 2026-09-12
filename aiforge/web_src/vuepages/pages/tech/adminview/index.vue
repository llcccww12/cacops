<template>
  <div>
    <TopHeader :menu="'admin_view'"></TopHeader>
    <div class="ui container">
      <div class="top-container">
        <el-checkbox class="check-toggle" v-model="allChecked" @change="allSelectChange">全选/全不选</el-checkbox>
        <el-button class="btn-agree" @click="batchOperate(true)">批量同意展示</el-button>
        <el-button class="btn-cancel" @click="batchOperate(false)">批量取消展示</el-button>
      </div>
      <div class="table-container">
        <div class="table-title">科技2030项目管理</div>
        <div class="table-wrap">
          <el-table ref="tableRef" border :data="tableData" style="width:100%;" v-loading="loading" stripe row-key="id" max-height="650">
            <el-table-column label="" align="center" header-align="center" width="40" fixed>
              <template slot-scope="scope">
                <el-checkbox v-model="scope.row.checked" @change="rowSelectChange(scope.row)"></el-checkbox>
              </template>
            </el-table-column>
            <el-table-column label="启智项目名称" align="center" header-align="center" fixed min-width="150">
              <template slot-scope="scope">
                <span v-if="scope.row.status == 5">
                  {{ `${scope.row.url.split('/')[3]}/${scope.row.url.split('/')[4]}` }}
                </span>
                <span v-else-if="scope.row.status == 3 || scope.row.status == 4">
                  {{ `${scope.row.url}` }}
                </span>
                <a v-else target="_blank" :href="`/${scope.row.repo_owner_name}/${scope.row.repo_name}`">
                  {{ `${scope.row.repo_owner_name}/${scope.row.repo_name}` }}
                </a>
              </template>
            </el-table-column>
            <el-table-column prop="name" label="科技项目名称" align="center" header-align="center" fixed
              min-width="200"></el-table-column>
            <el-table-column prop="status" label="展示状态" align="center" header-align="center" fixed>
              <template slot-scope="scope">
                <span style="color:rgb(255, 37, 37)"
                  :style="scope.row.status == 1 ? { color: 'rgb(56, 158, 13)' } : ''">{{
                  statusMap[scope.row.status] }}</span>
              </template>
            </el-table-column>
            <el-table-column prop="no" label="项目立项编号" align="center" header-align="center"
              min-width="120"></el-table-column>
            <el-table-column prop="institution" label="项目承担单位" align="center" header-align="center"
              min-width="160"></el-table-column>
            <el-table-column prop="province" label="所属省(省市)" align="center" header-align="center"
              min-width="100"></el-table-column>
            <el-table-column prop="category" label="单位性质" align="center" header-align="center"
              min-width="120"></el-table-column>
            <el-table-column prop="recommend" label="推荐单位" align="center" header-align="center"
              min-width="120"></el-table-column>
            <el-table-column prop="owner" label="项目负责人" align="center" header-align="center"
              min-width="100"></el-table-column>
            <el-table-column prop="phone" label="负责人电话" align="center" header-align="center"
              min-width="120"></el-table-column>
            <el-table-column prop="email" label="负责人邮箱" align="center" header-align="center"
              min-width="150"></el-table-column>
            <el-table-column prop="contact" label="项目联系人" align="center" header-align="center"
              min-width="100"></el-table-column>
            <el-table-column prop="contact_phone" label="联系人电话" align="center" header-align="center"
              min-width="120"></el-table-column>
            <el-table-column prop="contact_email" label="联系人邮箱" align="center" header-align="center"
              min-width="150"></el-table-column>
            <el-table-column prop="apply_year" label="申报年度" align="center" header-align="center"
              min-width="80"></el-table-column>
            <el-table-column prop="execute_month" label="执行周期(月)" align="center" header-align="center"
              min-width="120"></el-table-column>
            <el-table-column prop="execute_period" label="执行期限" align="center" header-align="center"
              min-width="120"></el-table-column>
            <el-table-column prop="type" label="项目类型" align="center" header-align="center"
              min-width="100"></el-table-column>
            <el-table-column prop="start_up" label="启动会时间" align="center" header-align="center"
              min-width="100"></el-table-column>
            <el-table-column prop="number_topic" label="课题数量" align="center" header-align="center"
              min-width="80"></el-table-column>
            <el-table-column prop="topic1" label="课题一级名称及承担单位" align="center" header-align="center"
              min-width="200"></el-table-column>
            <el-table-column prop="topic2" label="课题二级名称及承担单位" align="center" header-align="center"
              min-width="200"></el-table-column>
            <el-table-column prop="topic3" label="课题三级名称及承担单位" align="center" header-align="center"
              min-width="200"></el-table-column>
            <el-table-column prop="topic4" label="课题四级名称及承担单位" align="center" header-align="center"
              min-width="200"></el-table-column>
            <el-table-column prop="topic5" label="课题五级名称及承担单位" align="center" header-align="center"
              min-width="200"></el-table-column>
            <el-table-column prop="all_institution" label="所有参与单位" align="center" header-align="center"
              min-width="300"></el-table-column>
            <el-table-column prop="contribution_institution" label="成果贡献单位" align="center" header-align="center"
              min-width="300"></el-table-column>
            <el-table-column prop="user_name" label="申请账号" align="center" header-align="center" width="120">
              <template slot-scope="scope">
                <a target="_blank" :href="`/${scope.row.user_name}`">{{
                `${scope.row.user_name}` }}</a>
              </template>
            </el-table-column>
            <el-table-column prop="" label="操作" align="center" header-align="center" fixed="right" width="100">
              <template slot-scope="scope">
                <span v-if="scope.row.status == 2" class="op-btn agree" @click="toggleStatus(scope.row)">
                  <i class="el-icon-check"></i>
                  <span>同意展示</span>
                </span>
                <span v-if="scope.row.status == 1" class="op-btn cancel" @click="toggleStatus(scope.row)">
                  <i class="el-icon-close"></i>
                  <span>取消展示</span>
                </span>
              </template>
            </el-table-column>
          </el-table>
        </div>
        <div class="op-tips">
          <i class="el-icon-info"></i>
          提示：批量操作只会对【展示状态】为 “未展示” 或 “已展示” 的数据生效。
        </div>
        <div class="center">
          <el-pagination ref="paginationRef" background @current-change="currentChange" @size-change="sizeChange"
            :current-page.sync="page" :page-sizes="pageSizes" :page-size.sync="pageSize"
            layout="total, sizes, prev, pager, next, jumper" :total="total">
          </el-pagination>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import TopHeader from '../components/TopHeader.vue';
import { getTechAdminList, setTechAdminOperation } from '~/apis/modules/tech';

export default {
  data() {
    return {
      allChecked: false,
      loading: false,
      tableData: [],

      page: 1,
      pageSizes: [15, 30, 50, 100],
      pageSize: 50,
      total: 0,

      statusMap: {
        '1': '已展示',
        '2': '未展示',
        '3': '项目迁移中',
        '4': '项目迁移失败',
        '5': '项目不存在',
      },
    };
  },
  components: { TopHeader },
  methods: {
    getData() {
      this.loading = true;
      getTechAdminList({
        page: this.page,
        pageSize: this.pageSize,
      }).then(res => {
        this.loading = false;
        const { total, data } = res.data;
        this.tableData = data.map((item, index) => {
          return {
            checked: false,
            ...item,
          }
        });
        this.allChecked = false;
        this.total = total;
      }).catch(err => {
        this.loading = false;
        console.log(err);
      });
    },
    allSelectChange() {
      for (let i = 0, iLen = this.tableData.length; i < iLen; i++) {
        this.tableData[i].checked = this.allChecked;
      }
      this.tableData.splice(1, 0);
    },
    rowSelectChange(row) {
      if (row.checked == false) {
        this.allChecked = false;
      } else {
        if (this.tableData.filter((item) => item.checked).length == this.tableData.length) {
          this.allChecked = true;
        }
      }
      this.tableData.splice(1, 0);
    },
    currentChange(page) {
      this.page = page;
      this.getData();
    },
    sizeChange(pageSize) {
      this.pageSize = pageSize;
      this.getData();
    },
    toggleStatus(row) {
      if (row.status != 1 && row.status != 2) return;
      setTechAdminOperation({
        type: row.status == 1 ? 'hide' : 'show',
        id: [row.id],
      }).then(res => {
        res = res.data;
        if (res.Code == 0) {
          this.$message({
            type: 'success',
            message: this.$t('submittedSuccessfully'),
          });
          row.status = row.status == 1 ? 2 : 1;
        } else {
          this.$message({
            type: 'error',
            message: this.$t('submittedFailed'),
          });
        }
      }).catch(err => {
        console.log(err);
        this.$message({
          type: 'error',
          message: this.$t('submittedFailed'),
        });
      });
    },
    batchOperate(showOr) {
      const selectedData = this.tableData.filter((item) => {
        return item.checked && item.status == (showOr ? 2 : 1);
      });
      if (!selectedData.length) {
        this.$message({
          type: 'info',
          message: '请选择符合条件的数据进行操作！',
        });
        return;
      }
      setTechAdminOperation({
        type: showOr ? 'show' : 'hide',
        id: selectedData.map(item => item.id),
      }).then(res => {
        res = res.data;
        if (res.Code == 0) {
          this.$message({
            type: 'success',
            message: this.$t('submittedSuccessfully'),
          });
          this.getData();
        } else {
          this.$message({
            type: 'error',
            message: this.$t('submittedFailed'),
          });
        }
      }).catch(err => {
        console.log(err);
        this.$message({
          type: 'error',
          message: this.$t('submittedFailed'),
        });
      });
    },
  },
  beforeMount() {
    this.getData();
  },
  mounted() { },
  beforeDestroy() { },
};
</script>

<style scoped lang="less">
.top-container {
  display: flex;
  align-items: center;
  margin: 30px 0 15px 0;

  .check-toggle {
    margin-right: 20px;
  }

  .btn-agree {
    margin-right: 10px;
    background: rgb(50, 145, 248);
    color: rgb(255, 255, 255);

    &:hover {
      opacity: 0.9;
    }

    &:active {
      opacity: 0.8;
    }
  }

  .btn-cancel {
    margin-right: 10px;
    background: rgb(250, 140, 22);
    color: rgb(255, 255, 255);

    &:hover {
      opacity: 0.9;
    }

    &:active {
      opacity: 0.8;
    }
  }
}

.table-container {
  .table-title {
    height: 43px;
    font-size: 16px;
    font-weight: 700;
    padding: 15px;
    color: rgb(16, 16, 16);
    border-color: rgb(212, 212, 213);
    border-width: 1px;
    border-style: solid;
    border-radius: 5px 5px 0px 0px;
    background: rgb(240, 240, 240);
    display: flex;
    align-items: center;
  }

  .table-wrap {
    overflow-x: auto;

    /deep/ .el-table__header {
      th {
        background: rgb(249, 249, 249);
        font-size: 12px;
        color: rgb(136, 136, 136);
        font-weight: normal;
      }
    }

    /deep/ .el-table__body {
      td {
        font-size: 12px;
      }
    }

    /deep/ .el-radio__label {
      display: none;
    }
  }
}

.op-tips {
  margin: 10px 0 10px 0;
  color: #606266;
  font-size: 12px;

  i {
    margin-right: 4px;
  }
}

.op-btn {
  cursor: pointer;

  &.agree {
    color: rgb(56, 158, 13);
  }

  &.cancel {
    color: rgb(50, 145, 248);
  }
}
</style>
