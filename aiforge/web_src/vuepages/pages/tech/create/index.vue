<template>
  <div>
    <TopHeader :menu="''"></TopHeader>
    <div class="ui container">
      <div class="title">您申请的项目将在科技2030项目页面展示</div>
      <div class="form-c" v-show="!selectTechPrj">
        <div class="form-wrap">
          <div class="form-header">申请展示项目</div>
          <div class="form-content">
            <div class="form-row">
              <div class="row-label required">申请项目存放于</div>
              <div class="row-content">
                <el-radio v-model="form.type" label="openi" @input="changeType()">启智社区</el-radio>
                <el-radio v-model="form.type" label="no-openi" @input="changeType()">非启智社区</el-radio>
              </div>
            </div>
            <div class="form-row" :class="form.url_err ? 'form-row-err' : ''">
              <div class="row-label required">现项目地址</div>
              <div class="row-content">
                <el-input size="medium" v-model="form.url" @input="changeUrl" class="can-err"
                  placeholder='请输入现有项目的 HTTP(s) 或 Git " clone" URL，如：https://openi.pcl.ac.cn/OpenI/aiforge'></el-input>
              </div>
            </div>
            <div v-show="form.type == 'no-openi'">
              <div class="form-row" :class="form.alias_err ? 'form-row-err' : ''">
                <div class="row-label required baseline">启智项目名称</div>
                <div class="row-content">
                  <div>
                    <el-input size="medium" v-model="form.repo_alias" @input="changeAlias" class="can-err"></el-input>
                  </div>
                  <div class="tips">
                    请输入中文、字母、数字和-_ .，最多100个字符。
                  </div>
                </div>
              </div>
              <div class="form-row" :class="form.name_err ? 'form-row-err' : ''">
                <div class="row-label required baseline">启智项目路径</div>
                <div class="row-content">
                  <div class="reop-url-c">
                    <el-select size="medium" class="owner-sel" v-model="form.uid" @change="changeOwner">
                      <div slot="prefix" class="owner-item" style="height:100%;padding-left:10px">
                        <img class="owner-img" :src="ownerSelect.RelAvatarLink">
                      </div>
                      <el-option v-for="item in ownerList" :key="item.value" :value="item.value" :label="item.label">
                        <div class="owner-item">
                          <img class="owner-img" :src="item.RelAvatarLink">
                          <span class="owner-name">{{(item.FullName || item.Name)}}</span>
                        </div>
                      </el-option>
                    </el-select>
                    <span style="margin: 0 8px;font-size:22px;"> / </span>
                    <el-input size="medium" v-model="form.repo_name" @input="changeRepoName" class="can-err"></el-input>
                  </div>
                  <div class="tips">
                    启智项目地址：<span class="openi-repo-url">{{ form.repo_url }}</span>
                  </div>
                </div>
              </div>
              <div class="form-row" :class="form.topic_err ? 'form-row-err' : ''">
                <div class="row-label required">关键词</div>
                <div class="row-content">
                  <el-select style="width:100%" size="medium" v-model="form.topics" multiple filterable remote
                    allow-create default-first-option placeholder="请选择标签" :remote-method="searchTopics"
                    :loading="form.topicLoading" class="can-err">
                    <el-option v-for="item in topicsList" :key="item.value" :label="item.label" :value="item.value">
                    </el-option>
                  </el-select>
                </div>
              </div>
              <div class="form-row" :class="form.descr_err ? 'form-row-err' : ''">
                <div class="row-label required baseline">项目简介</div>
                <div class="row-content">
                  <el-input size="medium" type="textarea" :rows="4" placeholder="请输入项目简介(长度不超过255)" :maxlength="255"
                    show-word-limit v-model="form.description" class="can-err"></el-input>
                </div>
              </div>
            </div>
            <div class="form-row" :class="form.tech_err ? 'form-row-err' : ''">
              <div class="row-label required">科技项目</div>
              <div class="row-content">
                <el-input placeholder="请选择科技项目" size="medium" v-model="form.tech_show" :readonly="true" class="can-err">
                  <el-button size="medium" slot="append" class="btn-select" @click="goSelectTechPrj">选择科技项目</el-button>
                  <el-button size="medium" slot="append" class="btn-clear" @click="clearTechPrj">清除</el-button>
                </el-input>
              </div>
            </div>
            <div class="form-row" :class="form.institution_err ? 'form-row-err' : ''">
              <div class="row-label required">成果贡献单位</div>
              <div class="row-content">
                <el-select style="width:100%" size="medium" v-model="form.institution" multiple class="can-err">
                  <el-option v-for="item in institutionList" :key="item.key" :label="item.value"
                    :value="item.value"></el-option>
                </el-select>
              </div>
            </div>
            <div class="form-btn-group">
              <el-button size="medium" type="primary" :loading="submitLoading" class="btn confirm-btn"
                @click="submit">提交申请</el-button>
              <el-button size="medium" class="btn" @click="cancel">{{ $t('cancel') }}</el-button>
            </div>
          </div>
        </div>
      </div>
      <div class="form-select-tech-prj" v-show="selectTechPrj">
        <div class="form-wrap">
          <div class="form-header">
            <span>选择科技项目</span>
            <span>
              <el-button class="form-btn-go-back" @click="goBack">返回</el-button>
            </span>
          </div>
          <div class="form-content">
            <div class="form-row">
              <div class="row-label">请输入科技项目</div>
              <div class="row-content">
                <el-input placeholder="请输入搜索内容" size="medium" v-model="form.tech_search_keyword"
                  @keyup.enter.native="searchTechList" class="input-with-select">
                  <el-select v-model="form.tech_search_sel" style="width:142px" size="medium" slot="prepend">
                    <el-option label="项目立项编号" value="0"></el-option>
                    <el-option label="参与单位" value="1"></el-option>
                    <el-option label="项目名称" value="2"></el-option>
                  </el-select>
                  <el-button size="medium" slot="append" icon="el-icon-search" @click="searchTechList"></el-button>
                </el-input>
              </div>
            </div>
            <div class="form-table">
              <div style="margin: 0 20px 30px 20px">
                <el-table ref="tableRef" border :data="tableData" style="width:100%" v-loading="loading" stripe>
                  <el-table-column prop="no" label="项目立项编号" align="center" header-align="center"
                    width="120"></el-table-column>
                  <el-table-column prop="name" label="科技项目名称" align="center" header-align="center"
                    width="200"></el-table-column>
                  <el-table-column prop="institution" label="项目承担单位" align="center" header-align="center"
                    width="200"></el-table-column>
                  <el-table-column prop="all_institution" label="所有参与单位" align="center"
                    header-align="center"></el-table-column>
                  <el-table-column width="100" :label="$t('operation')" align="center" header-align="center">
                    <template slot-scope="scope">
                      <el-button type="primary" @click="selectedTechPrj(scope.row)">选择</el-button>
                    </template>
                  </el-table-column>
                </el-table>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import TopHeader from '../components/TopHeader.vue';
import { getTechs, setOpenIApply, setNoOpenIApply, getCreateRepoUser, getCheckRepoName, getTopics } from '~/apis/modules/tech';

export default {
  data() {
    return {
      form: {
        type: 'openi',
        url: '',
        url_err: false,
        repo_alias: '',
        alias_err: false,
        uid: '',
        repo_name: '',
        name_err: false,
        repo_url: '',
        topics: [],
        topic_err: false,
        topicLoading: false,

        description: '',
        descr_err: false,

        tech_search_sel: '0',
        tech_search_keyword: '',
        tech_obj: null,
        tech_show: '',
        tech_err: false,

        institution: [],
        institution_err: false,
      },
      selectTechPrj: false,
      loading: false,
      tableData: [],
      topicsList: [],
      institutionList: [],

      ownerList: [],
      ownerSelect: {},

      submitLoading: false,
    };
  },
  components: {
    TopHeader,
  },
  methods: {
    resetData() {
      this.form.url = '';
      this.form.url_err = false;
      this.form.repo_alias = '';
      this.form.alias_err = false;
      this.form.uid = this.ownerList.length ? this.ownerList[0].value : '';
      this.form.repo_name = '';
      this.form.repo_url = '';
      this.form.name_err = false;
      this.form.topics = [];
      this.form.topic_err = false;
      this.form.topicLoading = false;
      this.form.description = '';
      this.form.descr_err = false;
      this.form.tech_search_sel = '0';
      this.form.tech_search_keyword = '';
      this.form.tech_obj = null;
      this.form.tech_show = '';
      this.form.tech_err = false;
      this.form.institution = [];
      this.form.institution_err = false;
      this.tableData = [];
      this.topicsList = [];
      this.institutionList = [];
      this.submitLoading = false;
    },
    changeType() {
      this.resetData();
    },
    checkUrl() {
      this.form.url_err = !this.form.url;
      return !this.form.url_err;
    },
    checkRepoAlias() {
      const reg = /^[\u4E00-\u9FA5A-Za-z0-9_.-]{1,100}$/;
      const res = reg.test(this.form.repo_alias);
      this.form.alias_err = !res;
      return res;
    },
    checkRepoName() {
      const reg = /^[A-Za-z0-9_.-]{1,100}$/;
      const res = reg.test(this.form.repo_name);
      this.form.name_err = !res;
      return res;
    },
    checkTopic() {
      this.form.topic_err = this.form.topics.length == 0;
      return !this.form.topic_err;
    },
    checkDescr() {
      this.form.descr_err = !this.form.description;
      return !this.form.descr_err;
    },
    checkTech() {
      this.form.tech_err = !this.form.tech_obj;
      return !this.form.tech_err;
    },
    checkinstitution() {
      this.form.institution_err = this.form.institution.length == 0;
      return !this.form.institution_err;
    },
    changeUrl() {
      if (this.form.type == 'openi') return;
      const owner = this.ownerSelect.Name;
      const urlAdd = location.href.split("/")[0] + "//" + location.href.split("/")[2];
      const repoValue = this.form.url.match(/^(.*\/)?((.+?)(\.git)?)$/)[3];
      this.form.repo_alias = repoValue;
      this.checkRepoAlias();
      getCheckRepoName({ owner: owner, q: repoValue }).then(res => {
        const repo_name = res.data.name;
        this.form.repo_name = repo_name;
        this.form.repo_url = `${urlAdd}/${owner}/${repo_name}.git`;
        this.checkRepoName();
      }).catch(err => {
        console.log(err);
        this.form.repo_name = '';
        this.form.repo_url = '';
      });
    },
    changeAlias() {
      const owner = this.ownerSelect.Name;
      const aliasValue = this.form.repo_alias;
      const urlAdd = location.href.split("/")[0] + "//" + location.href.split("/")[2];
      if (aliasValue && this.checkRepoAlias()) {
        getCheckRepoName({ owner: owner, q: aliasValue }).then(res => {
          const repo_name = res.data.name;
          this.form.repo_name = repo_name;
          this.form.repo_url = `${urlAdd}/${owner}/${repo_name}.git`;
        }).catch(err => {
          console.log(err);
          this.form.repo_name = '';
          this.form.repo_url = '';
        });
      } else {
        this.form.repo_name = '';
        this.form.repo_url = '';
      }
    },
    changeOwner(value) {
      this.ownerSelect = this.ownerList.filter(item => item.value == value)[0];
      this.form.repo_name && this.changeRepoName();
    },
    changeRepoName() {
      const owner = this.ownerSelect.Name;
      const repo_name = this.form.repo_name;
      const urlAdd = location.href.split("/")[0] + "//" + location.href.split("/")[2];
      if (this.checkRepoName()) {
        this.form.repo_url = `${urlAdd}/${owner}/${repo_name}.git`;
      } else {
        this.form.repo_url = '';
      }
    },
    searchTopics(query) {
      if (query !== '') {
        this.form.topicLoading = true;
        getTopics({ q: query }).then(res => {
          this.form.topicLoading = false;
          const topics = res.data.topics || [];
          this.topicsList = topics.map(item => {
            return {
              value: item.topic_name,
              label: item.topic_name,
            }
          });
        }).catch(err => {
          this.topicsList = [];
        });
      } else {
        this.topicsList = [];
      }
    },
    goSelectTechPrj() {
      this.form.tech_search_sel = '0';
      this.form.tech_search_keyword = '';
      this.tableData = [];
      this.selectTechPrj = true;
      this.searchTechList();
    },
    clearTechPrj() {
      this.form.institution = [];
      this.institutionList = [];
      this.form.tech_show = '';
      this.form.tech_obj = null;
      this.form.tech_err = false;
    },
    goBack() {
      this.selectTechPrj = false;
    },
    searchTechList() {
      this.loading = true;
      getTechs({
        no: this.form.tech_search_sel == '0' ? this.form.tech_search_keyword : '',
        institution: this.form.tech_search_sel == '1' ? this.form.tech_search_keyword : '',
        name: this.form.tech_search_sel == '2' ? this.form.tech_search_keyword : '',
      }).then(res => {
        this.loading = false;
        res = res.data;
        this.tableData = res.data || [];
      }).catch(err => {
        console.log(err);
        this.loading = false;
        this.tableData = [];
      });
    },
    selectedTechPrj(item) {
      this.form.institution = [];
      this.institutionList = item.all_institution.split(',').map((item, index) => {
        return {
          key: item,
          value: item,
        }
      });
      this.form.tech_show = `【${item.no}】 ${item.name}`;
      this.form.tech_obj = item;
      this.checkTech();
      this.goBack();
    },
    submit() {
      const subData = {};
      let setApi = null;
      if (this.form.type == 'openi') {
        setApi = setOpenIApply;
        const r1 = this.checkUrl();
        const r2 = this.checkTech();
        const r3 = this.checkinstitution();
        if (r1 && r2 && r3) {
          subData.url = this.form.url;
          subData.no = this.form.tech_obj.no;
          subData.institution = this.form.institution.join(',');
        } else {
          return;
        }
      } else {
        setApi = setNoOpenIApply;
        const r1 = this.checkUrl();
        const r2 = this.checkTech();
        const r3 = this.checkinstitution();
        const r4 = this.checkRepoAlias();
        const r5 = this.checkRepoName();
        const r6 = this.checkTopic();
        const r7 = this.checkDescr();
        if (r1 && r2 && r3 && r4 && r5 && r6 && r7) {
          subData.url = this.form.url;
          subData.uid = this.form.uid;
          subData.alias = this.form.repo_alias;
          subData.repo_name = this.form.repo_name;
          subData.topics = [...this.form.topics];
          subData.description = this.form.description.trim();
          subData.no = this.form.tech_obj.no;
          subData.institution = this.form.institution.join(',');
        } else {
          return;
        }
      }
      // console.log(subData);
      // return;
      this.submitLoading = true;
      setApi(subData).then(res => {
        if (res.data && res.data.code == 0) {
          this.$message({
            type: 'success',
            message: this.$t('submittedSuccessfully'),
          });
          setTimeout(() => {
            window.location.href = '/tech/tech_view';
          }, 1000);
        } else {
          this.submitLoading = false;
          this.$message({
            type: 'info',
            message: res.data.msg
          });
        }
      }).catch(err => {
        this.submitLoading = false;
        this.$message({
          type: 'error',
          message: this.$t('submittedFailed'),
        });
      });
    },
    cancel() {
      window.history.back();
    }
  },
  beforeMount() {
    getCreateRepoUser().then(res => {
      const data = res.data.Data || [];
      this.ownerList = data.map(item => {
        return {
          value: item.ID,
          label: item.FullName || item.Name,
          ...item,
        }
      });
      this.ownerSelect = this.ownerList.length ? this.ownerList[0] : {};
      this.form.uid = this.ownerList.length ? this.ownerList[0].value : '';
    }).catch(err => {
      console.log(err);
    });
  },
  mounted() { },
  beforeDestroy() { },
};
</script>

<style scoped lang="less">
.title {
  text-align: center;
  margin: 30px 0;
  font-size: 14px;
  color: rgb(16, 16, 16);
}

.form-c,
.form-select-tech-prj {
  display: flex;
  justify-content: center;

  .form-wrap {
    width: 1000px;
    background: rgb(255, 255, 255);
    border-color: rgb(212, 212, 213);
    border-width: 1px;
    border-style: solid;
    border-radius: 5px;
    box-sizing: border-box;

    .form-header {
      height: 45px;
      background: rgb(240, 240, 240);
      border-bottom: 1px solid rgb(212, 212, 213);
      color: rgb(16, 16, 16);
      padding-left: 16px;
      padding-right: 16px;
      font-size: 16px;
      font-weight: 700;
      display: flex;
      align-items: center;
      justify-content: space-between;
    }

    .form-content {
      margin-top: 40px;

      .form-row {
        display: flex;
        align-items: center;
        width: 80%;
        margin: 20px auto;
        padding-right: 50px;

        .row-label {
          flex: 2;
          text-align: right;
          margin-right: 24px;
          position: relative;

          &.required {
            &::after {
              position: absolute;
              top: -2px;
              right: -10px;
              content: '*';
              color: #db2828;
            }
          }

          &.baseline {
            align-self: baseline;
            margin-top: 11px;
          }
        }

        .row-content {
          flex: 9;

          .tips {
            margin-top: 4px;
            font-size: 14px;
            color: rgb(136, 136, 136);


          }
        }

        .reop-url-c {
          display: flex;
          align-items: center;
        }

        .openi-repo-url {
          color: rgba(16, 16, 16, 1);
        }

        &.form-row-err {
          .row-label {
            color: #9f3a38;
          }

          /deep/ .can-err .el-input__inner,
          /deep/ .can-err .el-textarea__inner {
            color: #9f3a38;
            background: #fff6f6;
            border-color: #e0b4b4;
          }
        }
      }

      .form-table {
        margin-top: 30px;
        margin-bottom: 30px;

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

      .form-btn-group {
        display: flex;
        justify-content: center;
        align-items: center;
        margin: 30px 0 40px;
      }
    }
  }
}

.btn {
  color: rgb(2, 0, 4);
  background-color: rgb(194, 199, 204);
  border-color: rgb(194, 199, 204);

  &.confirm-btn {
    color: #fff;
    background-color: rgb(56, 158, 13);
    border-color: rgb(56, 158, 13);
    margin-right: 20px;
  }
}

.btn-select {
  background-color: rgb(50, 145, 248) !important;
  color: white !important;
  border-radius: 0 !important;
  height: 35px;
}

.owner-sel {
  /deep/ .el-input__inner {
    padding-left: 48px;
  }
}

.owner-item {
  display: flex;
  align-items: center;

  .owner-img {
    width: 24px;
    height: 24px;
    margin-right: 10px;
  }
}
</style>
