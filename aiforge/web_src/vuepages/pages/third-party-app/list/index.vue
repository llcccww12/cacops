<template>
  <div class="third-party-app-c">
    <div class="title"><span>{{ $t('thirdPartyApp.thirdPartyApp') }}</span></div>
    <div class="sub-title-1">{{ $t('thirdPartyApp.oauth2App') }}{{ $t('thirdPartyApp.oauth2AppDescr') }}</div>
    <div class="tools-bar">
      <div class="l">
        <el-select class="select" style="width:160px" size="medium" :placeholder="$t('thirdPartyApp.authStatus')"
          clearable v-model="conds.verify_flag" @change="condsChange">
          <el-option-group :label="$t('thirdPartyApp.authStatus')">
            <el-option v-for="item in verifyFlagList" :key="item.k" :label="item.v" :value="item.k" />
          </el-option-group>
        </el-select>
        <el-select class="select" style="width:240px" size="medium" :placeholder="$t('thirdPartyApp.canGetUserInfo')"
          clearable multiple collapse-tags v-model="conds.scopes" @change="condsChange">
          <el-option-group :label="$t('thirdPartyApp.canGetUserInfo')">
            <el-option v-for="item in scopeList" :key="item.k" :label="item.v" :value="item.k"
              :disabled="item.disabled" />
          </el-option-group>
        </el-select>
        <el-input class="search-keyword" style="width:300px" size="medium"
          :placeholder="$t('thirdPartyApp.searchThirdPartyAppPlaceholder')" v-model="conds.q"
          @keyup.enter.native="condsChange">
          <i slot="suffix" class="el-input__icon el-icon-search" @click="search"></i>
        </el-input>
      </div>
      <div class="r">
      </div>
    </div>
    <div class="table-container">
      <div style="min-height:600px;">
        <el-table border :data="tableData" style="width:100%" v-loading="loading" stripe>
          <el-table-column prop="ClientID" :label="$t('thirdPartyApp.clientID')" align="center" header-align="center"
            min-width="200"></el-table-column>
          <el-table-column prop="Name" :label="$t('thirdPartyApp.appName')" align="center" header-align="center"
            min-width="200"></el-table-column>
          <el-table-column prop="VerifyFlagStr" :label="$t('thirdPartyApp.authStatus')" align="center"
            header-align="center" min-width="120">
          </el-table-column>
          <el-table-column prop="AllowedScopes" :label="$t('thirdPartyApp.appCanGetUserInfo')" align="left"
            header-align="center" min-width="260">
            <template slot-scope="scope">
              <div v-for="_scope in scope.row.AllowedScopesList">{{ _scope.v }}</div>
            </template>
          </el-table-column>
          <el-table-column prop="UserName" :label="$t('thirdPartyApp.appCreator')" align="center" header-align="center"
            min-width="120">
          </el-table-column>
          <el-table-column prop="UpdatedUnixStr" :label="$t('modelManage.updateTime')" align="center"
            header-align="center" min-width="140">
          </el-table-column>
          <el-table-column :label="$t('operation')" align="center" header-align="center" min-width="220">
            <template slot-scope="scope">
              <span class="op-btn" v-if="[2].includes(scope.row.VerifyFlag)" @click="showDialog('edit', scope.row)">{{
                $t('edit') }}</span>
              <span class="op-btn" v-if="[2].includes(scope.row.VerifyFlag)" @click="cancelVerify(scope.row)">{{
                $t('thirdPartyApp.cancelVerify')
                }}</span>
              <span class="op-btn" v-if="[1].includes(scope.row.VerifyFlag)" @click="showDialog('verify', scope.row)">{{
                $t('thirdPartyApp.verify') }}</span>
              <span class="op-btn" v-if="[3].includes(scope.row.VerifyFlag)"
                @click="showDialog('verify_again', scope.row)">{{ $t('thirdPartyApp.verifyAgain')
                }}</span>
            </template>
          </el-table-column>
          <template slot="empty">
            <span style="font-size: 12px">{{
              loading ? $t('loading') : $t('noData')
              }}</span>
          </template>
        </el-table>
      </div>
      <div>
        <div style="margin-top: 2rem">
          <div class="center">
            <el-pagination background @current-change="currentChange" @size-change="pageSizeChange"
              :current-page="pageInfo.curpage" :page-sizes="pageInfo.pageSizes" :page-size="pageInfo.pageSize"
              layout="total, sizes, prev, pager, next, jumper" :total="pageInfo.total">
            </el-pagination>
          </div>
        </div>
      </div>
    </div>
    <EditDlg :visible.sync="dialogShow" :type="dialogType" :data="dialogData" @confirm="dialogConfirm"></EditDlg>
  </div>
</template>

<script>
import EditDlg from './components/EditDlg.vue';
import { getApplicaitonList, setEditApplicaiton } from '~/apis/modules/thirdpartyapp';
import { formatDate } from 'element-ui/lib/utils/date-util';
import { getListValueWithKey } from '~/utils';

export default {
  data() {
    return {
      verifyFlagList: [
        { k: 2, v: this.$t('thirdPartyApp.authenticated') },
        { k: 3, v: this.$t('thirdPartyApp.cancelAuthentication') },
        { k: 1, v: this.$t('thirdPartyApp.notAuthenticated') },
      ],
      scopeList: [
        { k: 'user.base', v: this.$t('thirdPartyApp.basicInfo'), label: this.$t('thirdPartyApp.basicInfoAll') },
        { k: 'user.email', v: this.$t('thirdPartyApp.email'), label: this.$t('thirdPartyApp.email') },
        { k: 'user.phone', v: this.$t('thirdPartyApp.phoneNumber'), label: this.$t('thirdPartyApp.phoneNumber') },
      ],
      conds: {
        q: '',
        verify_flag: '',
        scopes: [],
      },
      tableData: [],
      pageInfo: {
        curpage: 1,
        pageSizes: [10, 30, 50],
        pageSize: 10,
        total: 0,
      },
      loading: false,

      dialogShow: false,
      dialogType: '', // edit|verify|verify_again
      dialogData: {},
    };
  },
  components: { EditDlg },
  methods: {
    getTableData() {
      const params = {
        ...this.conds,
        scopes: this.conds.scopes.join(','),
        page: this.pageInfo.curpage,
        page_size: this.pageInfo.pageSize,
      };
      // console.log('params', JSON.stringify(params));
      this.loading = true;
      getApplicaitonList(params).then(res => {
        res = res.data;
        this.loading = false;
        if (res.code == 0) {
          res = res.data;
          this.pageInfo.total = res.total;
          this.tableData = (res.list || []).map(item => {
            return {
              ...item,
              VerifyFlagStr: getListValueWithKey(this.verifyFlagList, item.VerifyFlag),
              AllowedScopesList: (item.AllowedScopes || []).map(item => {
                return {
                  k: item,
                  v: getListValueWithKey(this.scopeList, item, 'k', 'label'),
                }
              }),
              CreatedUnixStr: formatDate(new Date(item.CreatedUnix * 1000), 'yyyy-MM-dd HH:mm:ss'),
              UpdatedUnixStr: formatDate(new Date(item.UpdatedUnix * 1000), 'yyyy-MM-dd HH:mm:ss'),
            }
          })
        } else {
          this.tableData = [];
          this.pageInfo.total = 0;
        }
      }).catch(err => {
        console.log(err);
        this.loading = false;
        this.tableData = [];
        this.pageInfo.total = 0;
      });
    },
    search() {
      this.pageInfo.curpage = 1;
      this.getTableData();
    },
    condsChange() {
      this.search();
    },
    currentChange(val) {
      this.pageInfo.curpage = val;
      this.getTableData();
    },
    pageSizeChange(val) {
      this.pageInfo.pageSize = val;
      this.search();
    },
    showDialog(type, data) {
      this.dialogType = type;
      this.dialogData = { ...data };
      this.dialogShow = true;
    },
    cancelVerify(data) {
      this.$confirm(this.$t('thirdPartyApp.cancelVerifyTips'), this.$t('tips'), {
        confirmButtonText: this.$t('confirm'),
        cancelButtonText: this.$t('cancel'),
        type: 'warning'
      }).then(() => {
        setEditApplicaiton({
          ID: data.ID,
          VerifyFlag: 3,
          ScopeList: 'user.base',
        }).then(res => {
          res = res.data;
          if (res.code === 0) {
            this.$message({
              type: 'success',
              message: this.$t('submittedSuccessfully')
            });
            this.getTableData();
          } else {
            this.$message({
              type: 'error',
              message: res.msg || this.$t('submittedFailed')
            });
          }
        }).catch(err => {
          console.log(err);
          this.$message({
            type: 'error',
            message: this.$t('submittedFailed')
          });
        })
      }).catch((err) => { console.log(err) });
    },
    dialogConfirm() {
      this.dialogShow = false;
      this.getTableData();
    },
  },
  beforeMount() { },
  mounted() {
    this.search();
  },
  beforeDestroy() { },
};
</script>

<style scoped lang="less">
.third-party-app-c {
  padding: 12px;

  .title {
    height: 30px;
    display: flex;
    align-items: center;
    margin-bottom: 5px;

    span {
      font-weight: 700;
      font-size: 16px;
      color: rgb(16, 16, 16);
    }
  }

  .sub-title-1 {}

  .sub-title-2 {}

  .tools-bar {
    display: flex;
    align-items: center;
    justify-content: space-between;
    margin-top: 18px;
    margin-bottom: 18px;

    .l {
      display: flex;
      align-items: center;

      .el-input,
      .select {
        margin-right: 15px;
      }

      .select {
        /deep/ .el-input__inner::placeholder {
          color: #606266;
        }
      }
    }

    .r {
      display: flex;
      align-items: center;
    }

    .el-input,
    .select {

      /deep/ .el-input__inner {
        border-radius: 0;
      }
    }

    /deep/ .el-icon-search {
      cursor: pointer;
      color: rgb(16, 16, 16);
    }
  }

  .table-container {
    margin-bottom: 16px;

    /deep/ .el-table__header {
      th {
        background: rgb(245, 245, 246);
        font-size: 12px;
        color: rgb(36, 36, 36);
      }
    }

    /deep/ .el-table__body {
      td {
        font-size: 12px;
      }
    }

    .op-btn {
      cursor: pointer;
      font-size: 12px;
      color: rgb(25, 103, 252);
      margin: 0 5px;
    }
  }

  .center {
    display: flex;
    justify-content: center;
  }
}

.el-select-group__wrap {
  /deep/ .el-select-group__title {
    font-size: 14px;
  }
}
</style>
