<template>
  <div class="ui container">
    <div class="title">
      <div class="title-1"><span>{{ $t('user.inviteFriends') }}</span></div>
      <div class="title-2"><span>{{ $t('user.inviteFriendsTips') }}</span></div>
    </div>
    <div class="content-1">
      <div class="img-c">
        <img class="img" :src="bannerImg" />
        <div class="txt">{{ bannerTitle }}</div>
      </div>
      <div class="descr">
        <span>{{ pageLinkDesc }}</span>
        <a :href="pageLink" target="_blank">{{ $t('user.clickToViewTheEventDetails') }}</a>
      </div>
    </div>
    <div class="content-2">
      <div class="txt-c">
        <div class="txt-1">
          <span>{{ pageOpeniDesc }}</span>
        </div>
        <div class="txt-2"><span>{{ $t('user.registrationAdress') }}</span><span>{{ invitationLink + invitationCode
            }}</span></div>
        <div class="txt-3"><span>{{ $t('user.recommender') }}</span><span>{{ invitationCode }}</span></div>
        <el-button class="__copy_link_btn__" type="primary">{{ $t('user.copyRegistrationInvitationLink') }}</el-button>
      </div>
      <div class="qr-code">
        <div id="__qr-code__" style="width:120px;height:120px;"></div>
      </div>
    </div>
    <div class="table-container">
      <div>
        <el-table border :data="tableData" style="width:100%" v-loading="loading" stripe>
          <el-table-column prop="ID" :label="$t('user.invitedFriends')" align="left" header-align="center">
            <template slot-scope="scope">
              <div style="display:flex;align-items:center;padding-left:20px;">
                <img :src="scope.row.avatarSrc" alt="" style="height:45px;width:45px;margin-right:10px;" />
                <a v-if="scope.row.userLink" :href="scope.row.userLink" style="font-weight:500;font-size:15px;">{{
                  scope.row.userName }}</a>
                <span v-else tyle="font-weight:500;font-size:15px;">{{ scope.row.userName }}</span>
              </div>
            </template>
          </el-table-column>
          <el-table-column prop="statusStr" :label="$t('status')" align="center" header-align="center">
            <template slot-scope="scope">
              <span :style="{ color: scope.row.statusColor }">{{ scope.row.statusStr }}</span>
            </template>
          </el-table-column>
          <el-table-column prop="regTime" :label="$t('user.registrationTime')" align="center" header-align="center">
          </el-table-column>
          <template slot="empty">
            <span>{{
              loading ? $t('loading') : $t('noData')
            }}</span>
          </template>
        </el-table>
      </div>
      <div class="__r_p_pagination">
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
</template>

<script>
import Clipboard from 'clipboard';
import QRCode from 'qrcodejs2';
import { formatDate } from 'element-ui/lib/utils/date-util';
import { getUserInvitationCode } from '~/apis/modules/userinvite';

export default {
  data() {
    return {
      bannerImg: '',
      bannerTitle: '',
      pageLink: '',
      pageLinkDesc: '',
      invitationLink: window.origin + '/user/sign_up?sharedUser=',
      invitationCode: '',
      pageOpeniDesc: '',

      loading: false,
      tableData: [],
      pageInfo: {
        curpage: 1,
        pageSize: 10,
        pageSizes: [10],
        total: 0,
      },
    };
  },
  components: {},
  methods: {
    initCopy() {
      const clipboard = new Clipboard('.__copy_link_btn__', {
        text: () => {
          return `${this.pageOpeniDesc}\n${this.$t('user.registrationAdress')}${this.invitationLink + this.invitationCode}\n${this.$t('user.recommender')}${this.invitationCode}`;
        },
      });
      clipboard.on('success', (e) => {
        this.$message({
          type: 'success',
          message: this.$t('user.theSharedContentHasBeenCopiedToTheClipboard')
        });
      });
      clipboard.on('error', (e) => {
        this.$message({
          type: 'error',
          message: this.$t('user.copyError')
        });
      });
    },
    transRowData(item) {
      return {
        userName: item.Name || 'Ghost',
        avatarSrc: item.Avatar || '/user/avatar/Ghost/-1',
        userLink: item.Name ? (window.origin + '/' + item.Name) : '',
        statusStr: item.Name ? this.$t('user.normal') : this.$t('userWasDel'),
        statusColor: item.Name ? 'rgb(82, 196, 26)' : 'rgb(245, 34, 45)',
        regTime: formatDate(new Date(item.CreatedUnix * 1000), 'yyyy-MM-dd HH:mm:ss'),
      }
    },
    initData() {
      getUserInvitationCode({ page: this.pageInfo.curpage, pageSize: this.pageInfo.pageSize }).then(res => {
        res = res.data;
        if (res) {
          this.bannerImg = res.page_banner_img;
          this.bannerTitle = res.page_banner_title;
          this.pageLink = res.page_link;
          this.pageLinkDesc = res.page_link_desc;
          this.invitationCode = res.invitation_code;
          this.pageOpeniDesc = res.page_openi_desc;
          this.tableData = (res.invitation_users || []).map((item, index) => {
            return this.transRowData(item);
          });
          this.pageInfo.total = res.invitation_users_count;
          const qrCode = new QRCode("__qr-code__", {
            text: this.invitationLink + this.invitationCode,
            width: 120,
            height: 120,
            colorDark: '#000000',
            colorLight: '#ffffff',
            correctLevel: QRCode.CorrectLevel.H
          });
        }
      }).catch(err => {
        console.log(err);
      });
    },
    getTableData() {
      const params = {
        page: this.pageInfo.curpage,
        pageSize: this.pageInfo.pageSize,
      };
      this.loading = true;
      getUserInvitationCode(params).then(res => {
        this.loading = false;
        res = res.data;
        const data = (res.invitation_users || []).map((item, index) => {
          return this.transRowData(item);
        });
        this.tableData = data;
        this.pageInfo.total = res.invitation_users_count;
      }).catch(err => {
        console.log(err);
        this.loading = false;
      });
    },
    currentChange(val) {
      this.pageInfo.curpage = val;
      this.getTableData();
    },
  },
  mounted() {
    this.initData();
    this.initCopy();
  },
  beforeDestroy() {
  },
};
</script>

<style scoped lang="less">
.title {
  margin-top: 15px;
  margin-bottom: 15px;

  .title-1 {
    font-weight: 500;
    font-size: 20px;
    color: rgba(16, 16, 16, 1);
    margin-bottom: 10px;
  }

  .title-2 {
    font-weight: 400;
    font-size: 14px;
    color: rgba(136, 136, 136, 1);
  }
}

.content-1 {
  margin-bottom: 32px;

  .img-c {
    height: 80px;
    position: relative;

    .img {
      width: 100%;
      height: 100%;
    }

    .txt {
      position: absolute;
      width: 100%;
      height: 100%;
      left: 0;
      top: 0;
      line-height: 80px;
      padding-left: 25px;
      font-weight: 500;
      font-size: 24px;
      color: rgb(255, 255, 255);
    }
  }

  .descr {
    font-weight: 300;
    font-size: 16px;
    color: rgb(16, 16, 16);
    padding: 25px;
    border-left: 1px solid rgba(0, 0, 0, 0.1);
    border-right: 1px solid rgba(0, 0, 0, 0.1);
    border-bottom: 1px solid rgba(0, 0, 0, 0.1);
    border-radius: 0px 0px 4px 4px;
  }
}

.content-2 {
  display: flex;
  background-color: rgb(228, 242, 255);
  border-color: rgb(228, 242, 255);
  border-width: 1px;
  border-style: solid;
  border-radius: 5px;
  padding: 25px;
  margin-bottom: 32px;

  .txt-c {
    flex: 1;
    font-weight: 300;
    font-size: 16px;
    color: rgb(16, 16, 16);

    span {
      line-height: 24px;
    }

    div {
      margin-bottom: 6px;
    }

    .txt-3 {
      margin-bottom: 15px;
    }
  }

  .__copy_link_btn__ {
    font-size: 14px;
    padding: 11px 15px;
    background: rgb(21, 114, 255);
    border-radius: 0;
  }

  .qr-code {
    width: 150px;
    display: flex;
    flex-direction: column;
    align-items: flex-end;
  }
}

.table-container {
  margin-bottom: 16px;

  /deep/ .el-table__header {
    th {
      background: rgb(245, 245, 246);
      font-size: 14px;
      color: rgb(36, 36, 36);
      font-weight: 400;
    }
  }

  /deep/ .el-table__body {
    td {
      font-size: 14px;
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
</style>
