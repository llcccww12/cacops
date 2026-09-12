<template>
  <div>
    <div class="ui container" id="header">
      <el-row style="margin-top: 15px">
        <el-table ref="table" :data="tableData" style="min-width: 100%" row-key="rowKey" lazy :load="load"
          :tree-props="{ children: 'Children', hasChildren: 'hasChildren' }" :header-cell-style="tableHeaderStyle">
          <el-table-column prop="name" :label="i18n.model_name" align="left" min-width="20%">
            <template slot-scope="scope">
              <div class="model-name-c">
                <div class="expand-icon" v-if="scope.row.hasChildren === false">
                  <i class="el-icon-arrow-right"></i>
                </div>
                <span v-if="!scope.row.Children">
                  <span v-if="scope.row.modelType == 0" class="m-online">{{ i18n.online }}</span>
                  <span v-if="scope.row.modelType == 1" class="m-local">{{ i18n.local }}</span>
                  <span v-if="scope.row.modelType == 2" class="m-external">{{ i18n.external }}</span>
                </span>
                <a class="text-over" :href="showinfoHref + encodeURIComponent(scope.row.name)"
                  :title="scope.row.name">{{ scope.row.name }}</a>
                <i class="ri-file-copy-2-line clipboard poping up" :data-original="i18n.copy" :data-content="i18n.copy"
                  :data-success="i18n.copySuccess" data-variation="inverted tiny" :data-clipboard-text="scope.row.name"
                  data-position="top center" style="color:#919191;font-size:14px;margin-left:5px;cursor:pointer;">
                </i>
              </div>
            </template>
          </el-table-column>
          <el-table-column prop="version" :label="i18n.model_version" align="center" min-width="6%">
            <template slot-scope="scope">
              <span class="text-over" :title="scope.row.version">{{
                scope.row.version
              }}</span>
            </template>
          </el-table-column>
          <el-table-column prop="versionCount" :label="i18n.model_version_num" align="center" min-width="7%">
            <template slot-scope="scope">
              <span class="text-over" :title="scope.row.versionCount">{{
                scope.row.versionCount
              }}</span>
            </template>
          </el-table-column>
          <el-table-column prop="size" :label="i18n.model_size" align="center" min-width="10%">
            <template slot-scope="scope">
              <span class="text-over">{{ renderSize(scope.row.size) }}</span>
            </template>
          </el-table-column>
          <el-table-column prop="engineName" :label="i18n.model_egine" align="center" min-width="8%">
            <template slot-scope="scope">
              <span class="text-over" :title="scope.row.engineName">{{
                scope.row.engineName
              }}</span>
            </template>
          </el-table-column>
          <el-table-column prop="isPrivate" :label="i18n.model_status" align="center" min-width="6.75%">
            <template slot-scope="scope">
              <span class="text-over" :title="scope.row.status_title">
                <i style="vertical-align: middle" :class="scope.row.status_str"></i></span>
              <span style="color: #fa8c16;" v-if="scope.row.isPrivate">{{ i18n.modelaccess_private }}</span>
              <span style="color: #13c28d;" v-else>{{ i18n.modelaccess_public }}</span>
            </template>
          </el-table-column>
          <el-table-column prop="createdUnix" :label="i18n.model_update_time" align="center" min-width="13.75%">
            <template slot-scope="scope">
              {{ transTime(scope.row.updatedUnix) }}
            </template>
          </el-table-column>
          <el-table-column prop="userName" :label="i18n.model_creator" align="center" min-width="6.75%">
            <template slot-scope="scope">
              <a :href="!scope.row.userName ? '#' : '/' + scope.row.userName"
                :title="scope.row.userName || defaultAvatarName">
                <img class="ui avatar image" :src="scope.row.userRelAvatarLink || defaultAvatar" />
              </a>
            </template>
          </el-table-column>
          <el-table-column :label="i18n.model_operation" min-width="15%" align="center">
            <template slot-scope="scope">
              <div class="space-around">
                <a class="op-btn"
                  :href="url + 'model_setting?type=1&name=' + encodeURIComponent(scope.row.name) + '&id=' + scope.row.id + '&back=' + encodeURIComponent(curHref)"
                  :class="{ disabled: !scope.row.isCanOper || scope.row.status == 1 }">{{ i18n.modify }}</a>
                <a class="op-btn" style="color: #13c28d;"
                  v-if="scope.row.modelType != 2 && scope.row.status != 1 && repoIsPrivate == false && scope.row.isPrivate == true && scope.row.isCanOper"
                  @click="
                    modifyModelStatus(scope.row.id, scope.row.cName, scope.row.rowKey, false)
                    ">{{ i18n.modelaccess_setpublic }}</a>
                <a class="op-btn" style="color: #fa8c16;"
                  v-if="scope.row.modelType != 2 && scope.row.status != 1 && repoIsPrivate == false && scope.row.isPrivate == false && scope.row.isCanOper"
                  @click="
                    modifyModelStatus(scope.row.id, scope.row.cName, scope.row.rowKey, true)
                    ">{{ i18n.modelaccess_setprivate }}</a>
                <a class="op-btn" @click="downloadAllModel(scope.row.size, scope.row.id,scope.row.name)"
                  :class="{ disabled: !scope.row.isCanDownload }">{{ i18n.model_download }}</a>
                <a class="op-btn" :class="{ disabled: !scope.row.isCanDelete || scope.row.status == 1 }" @click="
                  deleteModel(scope.row.id, scope.row.cName, scope.row.rowKey)
                  ">{{ i18n.model_delete }}</a>
              </div>
            </template>
          </el-table-column>
        </el-table>
      </el-row>
      <div class="ui container" style="margin-top: 50px; text-align: center">
        <el-pagination background @size-change="handleSizeChange" @current-change="handleCurrentChange"
          :current-page="currentPage" :page-sizes="[5, 10, 15]" :page-size="pageSize"
          layout="total, sizes, prev, pager, next, jumper" :total="totalNum">
        </el-pagination>
      </div>
    </div>
    <CommonTipsDialog ref="childModelTips" type="modelFile" :promotePath="codeUsePromotePath" :data="sdkData"
      :title="i18n.codeDownDlgTitle" :closeText="i18n.close">
    </CommonTipsDialog>
  </div>
</template>

<script>
import { modifyModelStatus } from '~/apis/modules/modelmanage';
import CommonTipsDialog from '~/components/CommonTipsDialog.vue';
import { initClipboard } from '~/utils';
const { _AppSubUrl, _StaticUrlPrefix, csrf } = window.config;
const REPOISPRIVATE = window.REPO_IS_PRIVATE;
const maxModelFilesSize = window.MAX_MODEL_SIZE || 0//536870912
export default {
  components: { CommonTipsDialog },
  data() {
    return {
      curHref: window.location.href,
      i18n: {},
      currentPage: 1,
      pageSize: 10,
      totalNum: 0,
      params: { page: 0, pageSize: 10 },
      tableData: [],
      url: "",
      isLoading: true,
      loadNodeMap: new Map(),
      submitId: {},
      repo: location.pathname.split('/').slice(0, 3).join('/'),
      defaultAvatar: "/user/avatar/Ghost/-1",
      defaultAvatarName: "Ghost",
      data: "",
      timer: null,
      timerFlag: false,
      codeUsePromotePath: `tips/modelDown/sdkcode${document.documentElement.attributes["lang"].nodeValue == "zh-CN" ? '' : '_en'}.md`,
      repoIsPrivate: REPOISPRIVATE,
      repoOwnerName: location.pathname.split('/')[1],
      repoName: location.pathname.split('/')[2],
      sdkData:{}
    };
  },
  methods: {
    load(tree, treeNode, resolve) {
      try {
        this.loadNodeMap.set(tree.cName, { tree, treeNode, resolve });
        this.$axios
          .get(this.url + "show_model_child_api", {
            params: {
              name: tree.cName,
            },
          })
          .then((res) => {
            let trainTaskInfo;
            let tableData;
            tableData = res.data;
            for (let i = 0; i < tableData.length; i++) {
              trainTaskInfo = JSON.parse(tableData[i].trainTaskInfo || '{}');
              tableData[i].engineName = this.getEngineName(tableData[i]);
              tableData[i].cName = tableData[i].name;
              tableData[i].rowKey = tableData[i].id + Math.random();
              tableData[i].name = "";
              tableData[i].versionCount = "";
              tableData[i].Children = true;
            }
            resolve(tableData || []);
          });
      } catch (e) {
        this.loading = false;
      }
    },
    tableHeaderStyle({ row, column, rowIndex, columnIndex }) {
      if (rowIndex === 0) {
        return "background:#f5f5f6;color:#606266";
      }
    },
    handleSizeChange(val) {
      this.params.pageSize = val;
      this.getModelList();
    },
    handleCurrentChange(val) {
      this.params.page = val;
      this.getModelList();
    },
    showcreateVue(name, version, label) {
      let title = this.i18n.model_create_version_title;
      $(".ui.modal.second")
        .modal({
          centered: false,
          onShow: function () {
            $("#model_header").text(title);
            $('input[name="name"]').addClass("model_disabled");
            $('input[name="name"]').attr("readonly", "readonly");
            $('input[name="modelSelectedFile"]').attr("readonly", "readonly");
            $('input[name="version"]').addClass("model_disabled");
            $(".ui.dimmer").css({
              "background-color": "rgb(136, 136, 136,0.7)",
            });
            $("#job-name").empty();
            $("#name").val(name);
            $("#label").val(label);
            let version_string = versionAdd(version);
            $("#version").val(version_string);
            loadTrainList();
          },
          onHide: function () {
            document.getElementById("formId").reset();
            $('input[name="name"]').removeClass("model_disabled");
            $('input[name="name"]').removeAttr("readonly");
            $('input[name="modelSelectedFile"]').removeAttr("readonly");
            var cityObj = $("#modelSelectedFile");
            cityObj.attr("value", "");
            $("#choice_model").dropdown("clear");
            $("#choice_version").dropdown("clear");
            $("#choice_Engine").dropdown("clear");
            $(".ui.dimmer").css({ "background-color": "" });
            $(".ui.error.message").text();
            $(".ui.error.message").css("display", "none");
          },
        })
        .modal("show");
    },
    check() {
      let jobid = document.getElementById("jobId").value;
      let versionname = document.getElementById("versionName").value;
      let name = document.getElementById("name").value;
      let version = document.getElementById("version").value;
      let modelSelectedFile =
        document.getElementById("modelSelectedFile").value;
      if (jobid == "") {
        $(".required.ten.wide.field").addClass("error");
        return false;
      } else {
        $(".required.ten.wide.field").removeClass("error");
      }
      if (modelSelectedFile == "") {
        $("#modelSelectedFile").addClass("error");
        return false;
      } else {
        $("#modelSelectedFile").removeClass("error");
      }
      if (versionname == "") {
        $(".required.six.widde.field").addClass("error");
        return false;
      } else {
        $(".required.six.widde.field").removeClass("error");
      }

      if (name == "") {
        $("#modelname").addClass("error");
        return false;
      } else {
        $("#modelname").removeClass("error");
      }
      if (versionname == "") {
        $("#verionname").addClass("error");
        return false;
      } else {
        $("#verionname").removeClass("error");
      }
      return true;
    },
    submit() {
      let context = this;
      let flag = this.check();
      if (flag) {
        let cName = $("input[name='name']").val();
        let version = $("input[name='version']").val();
        let data = $("#formId").serialize();
        const initModel = $("input[name='initModel']").val();
        let url_href =
          version === "0.0.1"
            ? context.url_create_newModel
            : context.url_create_newVersion;
        $("#mask").css({ display: "block", "z-index": "9999" });

        $.ajax({
          url: url_href,
          type: "POST",
          data: data,
          success: function (res) {
            context.getModelList();
            $("input[name='modelSelectedFile']").val("");
            $(".ui.modal.second").modal("hide");
            if (initModel === "0") {
              location.reload();
            }
          },
          error: function (xhr) {
            // 隐藏 loading
            // 只有请求不正常（状态码不为200）才会执行
            $(".ui.error.message").text(xhr.responseText);
            $(".ui.error.message").css("display", "block");
          },
          complete: function (xhr) {
            $("#mask").css({ display: "none", "z-index": "1" });
          },
        });
      } else {
        return false;
      }
    },
    loadrefresh(row) {
      const store = this.$refs.table.store;
      if (!this.loadNodeMap.get(row.cName)) {
        const parent = store.states.data;
        const index = parent.findIndex((child) => child.rowKey == row.rowKey);
        this.getModelList();
      } else {
        let { tree, treeNode, resolve } = this.loadNodeMap.get(row.cName);
        const keys = Object.keys(store.states.lazyTreeNodeMap);
        if (keys.includes(row.rowKey)) {
          this.getModelList();
        } else {
          let parentRow = store.states.data.find(
            (child) => child.cName == row.cName
          );
          let childrenIndex = store.states.lazyTreeNodeMap[
            parentRow.rowKey
          ].findIndex((child) => child.rowKey == row.rowKey);
          parentRow.versionCount = parentRow.versionCount - 1;
          const parent = store.states.lazyTreeNodeMap[parentRow.rowKey];
          if (parent.length === 1) {
            this.getModelList();
          } else {
            parent.splice(childrenIndex, 1);
          }
        }
      }
    },
    modifyModelStatus(id, name, rowKey, isPrivate) {
      let data = { 'id': id, 'isPrivate': isPrivate, 'repo': this.repo };
      modifyModelStatus(data).then(res => {
        res = res.data;
        if (res && res.code == '0') {
          this.getModelList();
        } else {
          this.$message({
            type: 'error',
            message: this.$t('modelManage.infoModificationFailed'),
          });
        }
      }).catch(err => {
        console.log(err);
        this.$message({
          type: 'error',
          message: this.$t('modelManage.infoModificationFailed'),
        });
      });
    },
    deleteModel(id, name, rowKey) {
      let row = { cName: name, id: id, rowKey: rowKey };
      let _this = this;
      let flag = 1;
      $(".ui.basic.modal.first")
        .modal({
          onDeny: function () {
            flag = false;
          },
          onApprove: function () {
            _this.$axios
              .delete(_this.url + "delete_model", {
                params: {
                  id: id,
                },
              })
              .then((res) => {
                _this.loadrefresh(row);
                // _this.getModelList()
              });
            flag = true;
          },
          onHidden: function () {
            if (flag == false) {
              $(".alert")
                .html(_this.i18n['canceled_operation'])
                .removeClass("alert-success")
                .addClass("alert-danger")
                .show()
                .delay(1500)
                .fadeOut();
            } else {
              $(".alert")
                .html(_this.i18n['successfully_deleted'])
                .removeClass("alert-danger")
                .addClass("alert-success")
                .show()
                .delay(1500)
                .fadeOut();
            }
          },
        })
        .modal("show");
    },
    getEngineName(model) {
      if (model.engine == 0) {
        return "PyTorch";
      } else if (model.engine == 1 || model.engine == 121) {
        return "TensorFlow";
      } else if (
        model.engine == 2 ||
        model.engine == 122 ||
        model.engine == 35
      ) {
        return "MindSpore";
      } else if (model.engine == 4) {
        return "PaddlePaddle";
      } else if (model.engine == 5) {
        return "OneFlow";
      } else if (model.engine == 6) {
        return "MXNet";
      } else {
        return "Other";
      }
    },
    intervalModelist() {
      if (!this.timerFlag) {
        this.timer = setInterval(() => {
          this.getModelList();
        }, 10000);
      }
    },
    getModelList() {
      let countStatus = 0;
      try {
        this.loadNodeMap.clear();
        this.$axios
          .get(this.url + "show_model_api", {
            params: this.params,
          })
          .then((res) => {
            $(".ui.grid").removeAttr("style");
            $("#loadContainer").removeClass("loader");
            let trainTaskInfo;
            this.tableData = res.data.data;
            for (let i = 0; i < this.tableData.length; i++) {
              trainTaskInfo = JSON.parse(this.tableData[i].trainTaskInfo || '{}');
              this.tableData[i].cName = this.tableData[i].name;
              this.tableData[i].rowKey = this.tableData[i].id + Math.random();
              this.tableData[i].engineName = this.getEngineName(
                this.tableData[i]
              );
              this.tableData[i].hasChildren = res.data.data[i].versionCount === 1 ? false : true;
              if (this.tableData[i].status !== 1) {
                countStatus++;
              }

              switch (this.tableData[i].status) {
                case 1:
                  this.tableData[i].status_str = "WAITING";
                  this.tableData[i].status_title = this.i18n.model_wait;
                  if (this.tableData[i].modelType == 2) {
                    this.tableData[i].status_title = this.i18n.model_migrating;
                  }
                  break;
                case 2:
                  this.tableData[i].status_str = "FAILED";
                  this.tableData[i].status_title = this.tableData[i].statusDesc;
                  if (this.tableData[i].modelType == 2) {
                    this.tableData[i].status_title = this.i18n.model_migrate_failed;
                  }
                  break;
                default:
                  this.tableData[i].status_str = "SUCCEEDED";
                  this.tableData[i].status_title = this.i18n.model_success;
                  if (this.tableData[i].modelType == 2) {
                    this.tableData[i].status_title = this.i18n.model_migrate_success;
                  }
                  break;
              }
            }
            this.totalNum = res.data.count;
            if (countStatus === this.tableData.length) {
              clearInterval(this.timer);
              this.timer = null;
              this.timerFlag = false;
            } else {
              this.intervalModelist();
              this.timerFlag = true;
            }
            this.$nextTick(() => {
              initClipboard('#header .clipboard');
            })
          })
          .catch((err) => {
            console.log(err);
          });
      } catch (e) {
        console.log(e);
      }
    },
    downloadAllModel(size, id, name) {
      if (size <= maxModelFilesSize) {
        let downloadElement = document.createElement('a')
        let href = `${this.loadhref}${id}`
        downloadElement.href = href
        document.body.appendChild(downloadElement)
        downloadElement.click() //点击下载
        document.body.removeChild(downloadElement) //下载完成移除元素
      } else {
        this.sdkData.owner = this.repoOwnerName
        this.sdkData.repo = this.repoName
        this.sdkData.name = name
        this.$nextTick(() => {
          this.$refs.childModelTips.handlerOpen('',true,maxModelFilesSize)
        })
      }
    }
  },
  computed: {
    loadhref() {
      return this.url + "downloadall?id=";
    },
    showinfoHref() {
      return this.url + "model_readme_tmpl?name=";
    },
    transStatus() {
      return function (state) {
        if (state) {
          return this.i18n.modelaccess_private;
        }
        return this.i18n.modelaccess_public;
      }
    },
    transTime() {
      return function (time) {
        let date = new Date(time * 1000); //时间戳为10位需*1000，时间戳为13位的话不需乘1000
        let Y = date.getFullYear() + "-";
        let M =
          (date.getMonth() + 1 < 10
            ? "0" + (date.getMonth() + 1)
            : date.getMonth() + 1) + "-";
        let D =
          (date.getDate() < 10 ? "0" + date.getDate() : date.getDate()) + " ";
        let h =
          (date.getHours() < 10 ? "0" + date.getHours() : date.getHours()) +
          ":";
        let m =
          (date.getMinutes() < 10
            ? "0" + date.getMinutes()
            : date.getMinutes()) + ":";
        let s =
          date.getSeconds() < 10 ? "0" + date.getSeconds() : date.getSeconds();
        return Y + M + D + h + m + s;
      };
    },
    renderSize() {
      return function (value) {
        if (null == value || value == "") {
          return "0 Bytes";
        }
        var unitArr = new Array(
          "Bytes",
          "KB",
          "MB",
          "GB",
          "TB",
          "PB",
          "EB",
          "ZB",
          "YB"
        );
        var index = 0;
        var srcsize = parseFloat(value);
        index = Math.floor(Math.log(srcsize) / Math.log(1024));
        var size = srcsize / Math.pow(1024, index);
        size = size.toFixed(2); //保留的小数位数
        return size + unitArr[index];
      };
    },
  },
  mounted() {
    this.submitId = document.getElementById("submitId");
    this.intervalModelist();
    this.url = location.href.split("show_model")[0];
    this.submitId.addEventListener("click", this.submit);
    this.url_create_newVersion = this.url + "create_model";
    this.url_create_newModel = this.url + "create_new_model";
  },
  created() {
    if (document.documentElement.attributes["lang"].nodeValue == "en-US") {
      this.i18n = this.$locale.US;
    } else {
      this.i18n = this.$locale.CN;
    }
    this.getModelList();
  },
  beforeDestroy() {
    // 实例销毁之前对点击事件进行解绑
    this.submitId.removeEventListener("click", this.submit);
    clearInterval(this.timer);
  },
};
</script>

<style scoped lang="less">
.model-name-c {
  display: flex;
  align-items: center;

  .clipboard {
    display: none;
  }

  &:hover {
    .clipboard {
      display: inherit;
    }
  }
}

.text-over {
  overflow: hidden;
  text-overflow: ellipsis;
  vertical-align: middle;
  white-space: nowrap;
}

.m-local {
  background-color: rgb(22, 132, 252);
  color: white;
  padding: 2px 3px;
  border-radius: 4px;
  font-size: 12px;
  margin-right: 2px;
}

.m-online {
  background-color: rgb(91, 185, 115);
  color: white;
  padding: 2px 3px;
  border-radius: 4px;
  font-size: 12px;
  margin-right: 2px;
}

.m-external {
  background-color: #5c246a;
  color: white;
  padding: 2px 3px;
  border-radius: 4px;
  font-size: 12px;
  margin-right: 2px;
}

.el-icon-arrow-right {
  font-family: element-icons !important;
  speak: none;
  font-style: normal;
  font-weight: 400;
  font-feature-settings: normal;
  font-variant: normal;
  text-transform: none;
  line-height: 1;
  vertical-align: middle;
  display: inline-block;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  border: 1px solid #d4d4d5;
  border-radius: 50%;
  color: #d4d4d5;
  margin-right: 4px;
}

.el-icon-arrow-right::before {
  content: "\e6e0";
}

.expand-icon {
  display: inline-block;
  width: 20px;
  line-height: 20px;
  height: 20px;
  text-align: center;
  margin-right: 3px;
  font-size: 12px;
}

/deep/ .el-table_1_column_1.is-left .cell {
  padding-right: 0px !important;
  white-space: nowrap;
}

/deep/ .el-table__expand-icon .el-icon-arrow-right {
  font-family: element-icons !important;
  speak: none;
  font-style: normal;
  font-weight: 400;
  font-feature-settings: normal;
  font-variant: normal;
  text-transform: none;
  line-height: 1;
  vertical-align: middle;
  display: inline-block;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  border: 1px solid #3291f8;
  border-radius: 50%;
  color: #3291f8;
  margin-right: 4px;
}

.space-around {
  display: flex;
  justify-content: space-around;
}

.op-btn-c {
  text-align: right;
  padding-right: 20px;
}

.op-btn {
  margin: 0 0 0 5px;
}

.disabled {
  cursor: default;
  pointer-events: none;
  color: rgba(0, 0, 0, 0.6) !important;
  opacity: 0.45 !important;
}
</style>
