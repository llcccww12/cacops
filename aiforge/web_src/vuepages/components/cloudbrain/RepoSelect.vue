<template>
  <div class="form-row">
    <div class="left-area">
      <div class="title">
        <span class="required">{{ $t('modelManage.project') }}</span>
      </div>
      <div class="content">
        <div class="content-wrap">
          <el-select class="type-sel field-input" v-model="typeSel" @change="changeType" :disabled="submitLoading">
            <el-option v-for="(item, index) in types" :key="index" :value="item.k" :label="item.v"></el-option>
          </el-select>
          <div class="content-r" v-if="typeSel == 'fork'" :class="forkErrStatus.source ? 'error' : ''">
            <el-input class="field-input" v-model="fork.source" @input="handleInput" maxlength="200"
              :disabled="submitLoading" @change="handleForkSourceChange"></el-input>
          </div>
          <div class="content-r" v-if="typeSel == 'new'">
            <el-input class="cur-user-input field-input" v-model="curUser.name" readonly disabled></el-input>
            <span class="separator-line">/</span>
            <div class="content-r repo-name-c" :class="errStatus ? 'error' : ''">
              <el-input class="new-repo-input field-input" v-model="newRepo" @input="handleInput" maxlength="100"
                @change="handleInput" :disabled="submitLoading"></el-input>
              <div class="tips">{{ $t('repoPathTips') }}</div>
            </div>
          </div>
          <div class="content-r" v-if="typeSel == 'select'" :class="errStatus ? 'error' : ''" :style="{flexDirection: usage == 'dataset' ?'column':''}">
            <el-select class="repo-sel field-input" v-model="repoSelectSel" :disabled="submitLoading" filterable>
              <el-option v-for="(item, index) in repoSelectList" :key="index" :value="item.k"
                :label="item.v"></el-option>
            </el-select>
            <div v-if="usage == 'dataset'" class="tips">{{ $t('datasetObj.dataset_repo_tips') }}</div>
          </div>
          
        </div>
        <div class="content-wrap content-wrap-fork-dest" v-if="typeSel == 'fork'">
          <div class="fork-dest-tit">{{ $t('modelManage.newRepoPath') }}</div>
          <div class="content-r">
            <el-input class="cur-user-input field-input" v-model="curUser.name" readonly disabled></el-input>
            <span class="separator-line">/</span>
            <div class="content-r repo-name-c" :class="forkErrStatus.dest ? 'error' : ''">
              <el-input class="new-repo-input field-input" v-model="fork.dest" @input="handleInput" maxlength="100"
                @change="handleInput" :disabled="submitLoading"></el-input>
              <div class="tips">{{ $t('repoPathTips') }}</div>
            </div>
          </div>
        </div>
      </div>
    </div>
    <div class="right-area" style="display:none"></div>
  </div>
</template>

<script>
import { getUserRepoList, createRepo, forkRepo } from '~/apis/modules/common';
import { i18n } from '~/langs';

const USE_TYPE = [
  { k: 'fork', v: i18n.t('cloudbrainObj.forkRepo') },
  { k: 'new', v: i18n.t('cloudbrainObj.newRepo') },
  { k: 'select', v: i18n.t('cloudbrainObj.selectRepo') },
];

export default {
  name: 'RepoSelect',
  props: {
    usage: { type: String, default: 'cloudbrain' }, // cloudbrain,model,dataset
    required: { type: Boolean, default: true },
    forkSource: { type: String, default: '' },
    useTypes: { type: Array, default: () => ['fork', 'new', 'select'] },
    delfaultSelect: { type: String, default: 'select' },
    latestTaskRepo: { type: Object, default: () => ({}) },
  },
  data() {
    return {
      types: [],
      typeSel: 'select',
      fork: {
        source: '',
        dest: '',
      },
      curUser: {
        id: '',
        name: '',
      },
      newRepo: '',
      repoSelectList: [],
      repoSelectSel: '',

      submitLoading: false,
      errStatus: false,
      forkErrStatus: {
        source: false,
        dest: false,
      },
    };
  },
  watch: {
    forkSource: {
      immediate: true,
      handler(newVal) {
        newVal = newVal === undefined ? '' : newVal;
        this.fork.source = newVal || '';
        this.fork.dest = this.fork.source.split('/')[1] || '';
      }
    },
    useTypes: {
      deep: true,
      handler(newVal, oldValue) {
        if (newVal.toString() == oldValue.toString()) return;
        this.refreshTypes();
      }
    },
    latestTaskRepo: {
      deep: true,
      handler(newVal, oldValue) {
        if (newVal.repo && this.usage == 'cloudbrain' && this.useTypes.indexOf('select') > -1) {
          this.typeSel = 'select';
          this.changeType(this.typeSel, () => {
            if (this.repoSelectList.map(item => item.k).indexOf(newVal.repo) > -1) {
              this.repoSelectSel = newVal.repo;
            }
          });
        }
      }
    },
  },
  methods: {
    refreshTypes(isFirstInit) {
      if (isFirstInit) {
        this.typeSel = this.delfaultSelect;
      }
      this.types = USE_TYPE.filter(item => this.useTypes.indexOf(item.k) > -1);
      const typesList = this.types.map(item => item.k);
      if (typesList.indexOf(this.typeSel) < 0) {
        this.typeSel = typesList[0] || '';
      }
      this.fork.source = this.forkSource || '';
      this.fork.dest = this.fork.source.split('/')[1] || '';
      if (this.typeSel == 'select' && !this.repoSelectList.length) {
        this.getRepoList();
      }
    },
    check() {
      if (this.typeSel == 'fork') {
        this.errStatus = false;
        this.forkErrStatus.source = false;
        this.forkErrStatus.dest = false;
        const reg1 = /^[.a-zA-Z0-9_-]{1,100}\/[.a-zA-Z0-9_-]{1,100}$/;
        const reg2 = /^[A-Za-z0-9_.-]{1,100}$/;
        if (!reg1.test(this.fork.source)) {
          this.errStatus = true;
          this.forkErrStatus.source = true;
        }
        if (!reg2.test(this.fork.dest)) {
          this.errStatus = true;
          this.forkErrStatus.dest = true;
        }
      }
      if (this.typeSel == 'new') {
        const reg = /^[A-Za-z0-9_.-]{1,100}$/;
        this.errStatus = !reg.test(this.newRepo);
      }
      if (this.typeSel == 'select') {
        this.errStatus = !this.repoSelectSel;
      }
      return !this.errStatus;
    },
    changeType(type, cb) {
      if (this.typeSel == 'fork') {
        this.fork.source = this.forkSource || '';
        this.fork.dest = this.fork.source.split('/')[1] || '';
      } else {
        this.fork.dest = '';
      }
      if (this.typeSel == 'new') {
        this.newRepo = '';
      } else {
        this.newRepo = '';
      }
      if (this.typeSel == 'select') {
        this.getRepoList(cb);
      } else {
        this.repoSelectSel = '';
      }
      this.errStatus = false;
    },
    handleForkSourceChange() {
      const reg = /^[.a-zA-Z0-9_-]{1,100}\/[.a-zA-Z0-9_-]{1,100}$/;
      if (reg.test(this.fork.source)) {
        this.fork.dest = this.fork.source.split('/')[1];
      }
      this.check();
    },
    handleInput() {
      this.check();
    },
    getRepoList(cb) {
      getUserRepoList({
        uid: this.curUser.id,
        sort: 'updated',
        order: 'desc',
        type: this.usage
      }).then(res => {
        res = res.data;
        if (res.ok) {
          const data = (res.data || []).map((item) => {
            return {
              k: item.html_url.split('/').slice(-2).join('/'),
              v: item.full_display_name,
              p: item.private,
              repoId: item.id,
              repo: item,
            }
          });
          this.repoSelectList = data;
          this.repoSelectSel = this.repoSelectList[0] ? this.repoSelectList[0].k : '';
          if (this.typeSel == 'select') {
            this.check();
          }
          cb && cb();
        }
      }).catch(err => {
        console.log(err);
      });
    },
    submit() {
      if (!this.check()) {
        this.$emit('submit', false);
        return;
      }
      if (this.submitLoading) return;
      this.submitLoading = true;
      if (this.typeSel == 'fork') {
        forkRepo({
          owner: this.fork.source.split('/')[0],
          repo: this.fork.source.split('/')[1],
          repo_name: this.fork.dest,
        }).then(res => {
          res = res.data;
          this.submitLoading = false;
          if (res.id) {
            const repoPath = res.html_url.split('/').slice(-2).join('/');
            this.$emit('submit', true, {
              type: this.typeSel,
              repoOwnerName: repoPath.split('/')[0],
              repoName: repoPath.split('/')[1],
              private: false,
              repoId: res.id,
              repo: res,
            });
          } else {
            this.$emit('submit', false);
          }
        }).catch(err => {
          this.submitLoading = false;
          this.$emit('submit', false);
          this.$message({
            type: 'error',
            message: err.response.data.message ?? err,
            duration: 4000,
          });
        });
        return;
      }
      if (this.typeSel == 'new') {
        createRepo({
          name: this.newRepo,
          private: false
        }).then(res => {
          res = res.data;
          this.submitLoading = false;
          if (res.id) {
            const repoPath = res.html_url.split('/').slice(-2).join('/');
            this.$emit('submit', true, {
              type: this.typeSel,
              repoOwnerName: repoPath.split('/')[0],
              repoName: repoPath.split('/')[1],
              private: false,
              repoId: res.id,
              repo: res,
            });
          } else {
            this.$emit('submit', false);
          }
        }).catch(err => {
          console.log(err);
          this.submitLoading = false;
          this.$emit('submit', false);
          this.$message({
            type: 'error',
            message: err.response.data.message ?? err,
            duration: 4000,
          });
        });
        return;
      }
      if (this.typeSel == 'select') {
        const selectRepo = this.repoSelectList.filter((item) => {
          return item.k === this.repoSelectSel
        })
        this.$emit('submit', true, {
          type: this.typeSel,
          repoOwnerName: this.repoSelectSel.split('/')[0],
          repoName: this.repoSelectSel.split('/')[1],
          private: selectRepo[0].p,
          repoId: selectRepo[0].repoId,
          repo: selectRepo[0].repo,
        });
        this.submitLoading = false;
        return;
      }
    },
  },
  beforeMount() {
    const metaEl = document.querySelector('meta[name="_uid"]');
    if (metaEl) {
      this.curUser.id = metaEl.getAttribute('content');
      this.curUser.name = metaEl.getAttribute('content-ext');
    }
    this.refreshTypes(true);
  },
  mounted() { },
};
</script>

<style scoped lang="less">
@import 'cloudbrain.less';

.form-row {
  .left-area {
    .content {
      .content-wrap {
        display: flex;

        &.content-wrap-fork-dest {
          margin-top: 10px;
        }

        .type-sel {
          width: 160px;
          margin-right: 10px;
        }

        .fork-dest-tit {
          width: 160px;
          margin-right: 10px;
          display: flex;
          margin-top: 10px;
          justify-content: flex-end;
        }

        .content-r {
          flex: 1;
          display: flex;
          align-items: baseline;

          .cur-user-input {
            width: auto;
            max-width: 160px;
          }

          .separator-line {
            font-size: 20px;
            margin: 0 4px;
          }

          .new-repo-input {
            flex: 1;
          }

          &.repo-name-c {
            flex-direction: column;
          }
        }

        .error {
          .field-input {

            /deep/.el-input__inner,
            /deep/.el-textarea__inner {
              color: #9f3a38;
              background: #fff6f6;
              border-color: #e0b4b4;

              &:visited {
                border-color: #e0b4b4;
              }

              &:focus {
                border-color: #e0b4b4;
              }

              &:active {
                border-color: #e0b4b4;
              }
            }
          }
        }
      }
    }
  }
}
</style>
