<template>
  <div style="overflow: hidden">
    <div v-if="params.type === 'ide'">
      <div class="ide-header">
        <span>
          <svg class="icon ide-header-svg" aria-hidden="true">
            <use xlink:href="#icon-wenjianbeifen"></use>
          </svg>
          {{ `${params.owner}/${params.project}` }}
        </span>
        <i class="el-icon-switch-button" @click="handleBack" />
      </div>
      <div style="display: flex">
        <div class="menu-left">
          <el-menu :default-active="activeKey" class="menu-left-drag-item-drag" background-color="#F5F5F5"
            text-color="#7A7A7B" active-text-color="#000" @select="handleselect">
            <el-menu-item index="1" style="line-height: 20px;height: 90px;padding: 20px 22px 46px 23px;"
              :class="{ activeBack: isActiveStatus }">
              <!--              <img class="menu-left-drag-img" :src="isActiveStatus ? code1 : code2" alt="" />-->
              <svg class="icon menu-left-drag-svg" aria-hidden="true">
                <use xlink:href="#icon-daima1"></use>
              </svg>
              <div slot="title">代码</div>
            </el-menu-item>
            <el-menu-item index="2"
              style="line-height: 20px;height: 90px;background-color: #F5F5F5;padding: 20px 22px 46px 23px;"
              :class="{ activeBack: activeKey === '2' }">
              <!--              <img class="menu-left-drag-img" :src="isActiveStatus ? submitLog2 : submitLog1" alt="" />-->
              <svg class="icon menu-left-drag-svg" aria-hidden="true">
                <use xlink:href="#icon-tijiao"></use>
              </svg>
              <div slot="title">提交</div>
            </el-menu-item>
          </el-menu>
        </div>
        <div ref="box" class="box">
          <div class="inlineBlock leftPaneBase" :style="{ width: leftPanelWidth }">
            <div class="leftPaneBase-submit" style="padding-left:21px;position:relative;">
              {{ isActiveStatus? "代码仓库": "提交" }}
              <el-dropdown @command="handleCommand" v-if="isActiveStatus" style="position:absolute;right:8px;top:18px;">
                <span class="el-dropdown-link">
                  <svg class="icon el-dropdown-svg" aria-hidden="true" style="color:#979797;font-size: 16px;">
                    <use xlink:href="#icon-a-bianzu31"></use>
                  </svg>
                </span>
                <el-dropdown-menu slot="dropdown">
                  <el-dropdown-item command="a">新建文件</el-dropdown-item>
                  <el-dropdown-item command="b">新建文件夹</el-dropdown-item>
                </el-dropdown-menu>
              </el-dropdown>
            </div>
            <VueTree v-loading="treeLoading" v-show="isActiveStatus" :tree-list-data="treeData" :file-info-params="fileInfoParams"
              @handleChangFile="handleChangFile" @deleteNode="handleNode" @handleAddNode="handleAddNode"
              @handleDeleteAddNode="handleDeleteAddNode" @handleDeleteOldNode="handleDeleteOldNode" ref="vuetree" />
            <div v-show="!isActiveStatus && inCludeSub" class="leftPaneBase-wrap">
              <div class="leftPaneBase-wrap-change">
                <span style="font-weight: 400;">更改</span>
                <el-tooltip class="item" effect="dark" content="放弃所有更改" placement="bottom-end">
                  <!--                  <img class="leftPaneBase-wrap-change-img" :src="abandon" alt="" @click="handleResetAll()" />-->
                  <svg class="icon leftPaneBase-wrap-change-img leftPaneBase-wrap-change-svg " aria-hidden="true"
                    @click="handleResetAll()">
                    <use xlink:href="#icon-zhongzhi3"></use>
                  </svg>
                </el-tooltip>
              </div>
            </div>
            <div class="leftPaneBase-wrap" v-show="!isActiveStatus && inCludeSub">
              <div class="leftPaneBase-wrap-fileName">
                <div v-show="!submitAarry.length" class="no-commit">
                  当前没有可提交的文件
                </div>
                <div v-show="(!!submitAarry.length)" v-for="item in submitAarry" :key="item && item.sha">
                  <div v-if="!!item && item.type === 'blob'" class="leftPaneBase-wrap-fileName-item">
                    <span @click="handleDiff(item)">{{ item.name }}</span>
                      <el-tooltip class="item" effect="dark" content="放弃更改" placement="bottom-end">
                        <!--                    <img :src="abandon" alt="" @click="handleload(item)" />-->
                        <svg class="icon leftPaneBase-wrap-fileName-item-img leftPaneBase-wrap-change-svg"
                          @click="handleload(item)" aria-hidden="true">
                          <use xlink:href="#icon-zhongzhi3"></use>
                        </svg>
                      </el-tooltip>
                  </div>
                </div>
              </div>
            </div>
            <div style="width: 100%;padding: 0 20px;">
              <!-- <div style="border-bottom: 1px solid #b1bacd;"></div> -->
            </div>
            <div v-show="!isActiveStatus && inCludeSub" class="leftPaneBase-wrap">
              <div class="leftPaneBase-wrap-info">提交信息</div>
              <el-input :disabled="!submitAarry.length" v-model="textContent" class="leftPaneBase-wrap-input"
                type="textarea" :rows="10" placeholder="请输入提交信息" />
              <div style="width: 100%; padding: 0 20px;">
                <el-button :disabled="!submitAarry.length" v-loading="subLoading" type="primary"
                  element-loading-spinner="el-icon-loading" element-loading-background="rgba(0, 0, 0, 0.8)"
                  class="leftPaneBase-wrap-infoSubmit" @click="handleSubmit">
                  提交
                </el-button>
              </div>
            </div>
          </div>
          <!-- <img :src="drag" alt="" title="收缩侧边栏" class="resize" /> -->
          <div class="resize" title="收缩侧边栏"></div>
          <div class="inlineBlock mid">
            <div class="mid-header">
              <span>{{ isDiff? "": fileInfoParams.path }}</span>
              <div>
                <span v-if="fileInfo.fileType === 'ipynb' && !hideNbView" style="cursor:pointer;margin-right:20px"
                  @click="handleEdit">
                  修改
                  <i class="el-icon-edit" />
                </span>
                <span v-if="fileInfo.fileType === 'ipynb' && hideNbView" style="cursor:pointer;margin-right:20px"
                  @click="handleEdit">
                  预览
                  <i class="el-icon-view" />
                </span>
                <span style="cursor:pointer" @click="handleSetting">设置<i class="el-icon-s-tools" /></span>
              </div>
            </div>
            <!-- <codemirror /> -->
            <div v-show="!fileInfoParams.filePath" class="empty-tip">
              <img width="200" :src="empty" />
              请在左侧文件树中选择一个文件，编辑后即可提交你的修改
            </div>
            <div v-if="containerReload">
              <div class="showImg" v-if="fileInfo.fileType === 'image'">
                <img :src="fileInfo.download_url" />
              </div>
              <div class="showImg" v-else-if="fileInfo.fileType === 'office'">
                <el-button type="primary" @click="downloadUrl(fileInfo.download_url)"><a style="color:white"
                    :href="fileType && fileType.download_url">下载</a></el-button>
              </div>
              <div class="showPdf" v-else-if="fileInfo.fileType === 'pdf'">
                <embed :src="fileInfo.download_url" />
              </div>
              <div class="showImg"
                v-else-if="fileInfo.download_url && fileInfo.fileType !== 'txt' && fileInfo.fileType !== 'ipynb'">
                <el-button type="primary" @click="downloadUrl(fileInfo.download_url)"><a style="color:white"
                    :href="fileType && fileType.download_url">下载</a></el-button>
              </div>
              <div :class="(fileInfo.fileType === 'txt' || fileInfo.fileType === 'ipynb') ? '' : 'hides'">
                <iframe v-if="(fileInfo.fileType === 'ipynb' && !hideNbView && !isDiff)"
                  src="/html/js/nbview/index.html" id="nbview" @load="loadFrame()"></iframe>
                <monaco ref="monaco" v-loading="monacoLoading" element-loading-text="拼命加载中"
                element-loading-spinner="el-icon-loading" element-loading-background="rgba(255, 255, 255, 0.8)"
                :opts="opts" :is-diff="isDiff" :monaca-value="monacaValue" @handleSave="handleSave" />
              </div>
            </div>

          </div>
        </div>
        <el-drawer style="height: calc(100vh - 94px); margin-top: 94px" class="drawer-wrap"
          :visible.sync="drawerVisible" direction="rtl" destroy-on-close size="270" :show-close="false"
          :with-header="false" :modal="false">
          <div class="drawer-section">
            <div class="drawer-section-style" style="padding-bottom:10px;">
              <div class="code-fast">代码格式</div>
              <div class="drawer-section-stydrawer-sectionle-wrap">
                <span class="drawer-wrap-span">显示模式</span>
                <el-select v-model="opts.theme" size="mini" placeholder="请选择" @change="handleChangeModelOptions">
                  <el-option v-for="item in modeOptions" :key="item.value" :label="item.label" :value="item.value" />
                </el-select>
              </div>
              <div class="drawer-section-style-wrap">
                <span class="drawer-wrap-span">字体大小</span>
                <el-select v-model="opts.fontSize" size="mini" placeholder="请选择" @change="handleChangeFontSizeOptions">
                  <el-option v-for="item in fontSizeOptions" :key="item.value" :label="item.label"
                    :value="item.value" />
                </el-select>
              </div>
              <!-- <div class="drawer-section-style-wrap">
                <span>Tab转换</span>
                <el-switch v-model="switchValue" active-color="#2185D0" inactive-color="#858585" />
              </div> -->
            </div>
            <div class="drawer-fast">
              <div class="code-fast">快捷键</div>
              <!-- <div class="drawer-fast-flex">
                <span  class="drawer-wrap-span">保存代码</span><span class="drawer-wrap-span2">Ctrl + S</span>
              </div>
              <div class="drawer-fast-flex">
                <span  class="drawer-wrap-span">唤出快捷键列表</span><span class="drawer-wrap-span2">F1/Alt + F1</span>
              </div> -->
              <div class="drawer-fast-flex">
                <span class="drawer-wrap-span">左右缩进</span><span class="drawer-wrap-span2">Ctrl + ]/[</span>
              </div>
            </div>
          </div>
        </el-drawer>
      </div>
    </div>
    <div v-else v-loading="loading" class="vscode-loading-wrap" element-loading-text="拼命加载中"
      element-loading-spinner="el-icon-loading">
      <div class="vscode-header">
        <span>{{ `${params.owner}/${params.project}` }}</span>
        <div class="vscode-header-right">
          <div class="color inlineBlock mr40">
            <i class="el-icon-time" />
            <span class="vscode-header-right-text">{{ count }}</span>
          </div>
          <div class="inlineBlock mr40 cspointe">
            <i class="el-icon-s-cooperation" />
            <span class="vscode-header-right-text" @click="handleReset">重置环境</span>
          </div>
          <div class="inlineBlock mr40 cspointe">
            <i class="el-icon-basketball" />
            <span class="vscode-header-right-text" @click="handleServicePreview">服务预览</span>
          </div>
          <i style="height:100%;display:flex;align-items:center;" class="el-icon-switch-button cspointe" @click="handleBack" />
        </div>
      </div>
      <div>
        <iframe v-if="headerParams.url" class="iframe-wrap" :src="headerParams.url" />
      </div>
    </div>
  </div>
</template>
<script>
// import VueTreeList from './VueTreeList.vue';
const API = "https://ide-kyxt.pcl.ac.cn";
import VueTree from "./vueTree.vue";
import {
  matterTree,
  RecurveQuery,
  RecurveQueryDelete,
  RecurveAddNode,
  RecurveUpdateValue,
  uniqueArray,
} from "../utils.js";
import monaco from "./monacoeditor.vue";
import code1 from "../../../public/images/code1.png";
import empty from "../../../public/img/empty.png";
import code2 from "../../../public/images/code2.png";
import submitLog2 from "../../../public/images/submitLog2.png";
import submitLog1 from "../../../public/images/submitLog1.png";
import abandon from "../../../public/images/abandon.png";
import warehouse from "../../../public/images/warehouse.png";
import submit from "../../../public/images/submit.png";
import drag from "../../../public/images/drag.png";
import {isBase64} from '../utils.js';
import {sufix} from './vueTree.vue';
import {
  init,
  add,
  readAll,
  readByMainKey,
  deleteDB,
  update,
  remove,
} from "./db.js";

export default {
  name: "IdeProject",
  components: {
    // codemirror,
    monaco,
    // VueTreeList,
    VueTree,
  },
  data() {
    return {
      activeKey: "1",
      drawerVisible: false,
      fileInfo: {},
      treeLoading:true,
      // 图片
      code1,
      code2,
      empty,
      submitLog1,
      submitLog2,
      abandon,
      warehouse,
      submit,
      drag,
      // 图片

      leftPanelWidth: "300px",
      modeOptions: [
        { label: "light", value: "vs" },
        { label: "dark", value: "vs-dark" },
      ],
      fontSizeOptions: [
        { label: "14px", value: 14 },
        { label: "16px", value: 16 },
        { label: "18px", value: 18 },
        { label: "20px", value: 20 },
        { label: "22px", value: 22 },
      ],
      textContent: "", // 文本域输入框的值
      monacaValue: "", // monaca编辑器的值
      fileInfoParams: {}, //  所点击的文件以及参数
      monacoLoading: false,
      sets: {
        language: {
          apex: "apex",
          azcli: "azcli",
          bat: "bat",
          c: "c",
          clojure: "clojure",
          coffeescript: "coffeescript",
          cpp: "cpp",
          csharp: "csharp",
          csp: "csp",
          css: "css",
          dockerfile: "dockerfile",
          fsharp: "fsharp",
          go: "go",
          graphql: "graphql",
          handlebars: "handlebars",
          html: "html",
          ini: "ini",
          java: "java",
          javascript: "javascript",
          json: "json",
          kotlin: "kotlin",
          less: "less",
          lua: "lua",
          markdown: "markdown",
          msdax: "msdax",
          mysql: "mysql",
          "objective-c": "objective-c",
          pascal: "pascal",
          perl: "perl",
          pgsql: "pgsql",
          php: "php",
          plaintext: "plaintext",
          postiats: "postiats",
          powerquery: "powerquery",
          powershell: "powershell",
          pug: "pug",
          python: "python",
          r: "r",
          razor: "razor",
          redis: "redis",
          redshift: "redshift",
          ruby: "ruby",
          rust: "rust",
          sb: "sb",
          scheme: "scheme",
          scss: "scss",
          shell: "shell",
          sol: "sol",
          sql: "sql",
          st: "st",
          swift: "swift",
          tcl: "tcl",
          typescript: "typescript",
          vb: "vb",
          xml: "xml",
          yaml: "yaml",
        },
        theme: {
          vs: "vs",
          "vs-dark": "vs-dark",
          "hc-black": "hc-black",
        },
        colorDecorators: true,
        comments: {
          ignoreEmptyLines: true, // 插入行注释时忽略空行。默认为真。
          insertSpace: true, // 在行注释标记之后和块注释标记内插入一个空格。默认为真。
        },
      },
      opts: {
        value: "",
        diffData: {},
        readOnly: false, // 是否可编辑
        language: "", // 语言类型
        theme: "vs", // 编辑器主题
        selectOnLineNumbers: true,
        roundedSelection: false,
        // readOnly: this.readOnly, // 只读
        cursorStyle: "line", // 光标样式
        automaticLayout: false, // 自动布局
        glyphMargin: true, // 字形边缘
        useTabStops: false,
        fontSize: 14, // 字体大小
        autoIndent: false, // 自动布局
      },
      hideNbView: false,
      isDiff: false,
      inlineDiff: false,
      containerReload: true,
      switchValue: false,
      loading: true,
      treeData: [], // 子组件的数据
      pathname: "", // 编辑文件的文件名
      submitAarry: [], // 提交时候数组数据
      subLoading: false,
      treeRootData: [], // 文件树的节点
      /** 以下为vscode的初始变量 */
      params: this.$route.query, // 从项目列表进来 type === ide | vscode 前者是编辑 后者是嵌套
      headerParams: {
        remainingTime: 0,
        url: "",
      }, // 包含src的url + 倒计时的remainingTime
      count: "00:00:00",
      timer: null,
      // eslint-disable-next-line semi
    };
  },
  computed: {
    isActiveStatus() {
      // 计算当前的activeKey为提交 还是 代码界面
      return this.activeKey === "1";
    },
    inCludeSub() {
      // 计算提交界面的数据是否有数据
      return true;
    },
  },
  mounted() {
    init({
      // 数据库名称
      dbName: this.params.project + '_' + this.params.owner + '_' + this.params.branch,
      // 数据表 数组 如果要创建多个表，则写多个即可
      tableList: [
        {
          // 表名
          tableName: "git",
          // 主键
          keyPath: "name",
          // 对应表的其他属性
          // attr:{
          //    autoIncrement:true, //是否自增
          // },
          indexName: "index", // 索引 不建议使用 因为使用索引 当前表必须有数据否则直接报错
          // unique:false // 对应索引是否唯一
        },
      ],
    });
    this.dragControllerDiv();
    setTimeout(() => {
      if (this.params.type === "vscode") {
        this.getVscodeData();
      } else {
        this.getCodeTree();
      }
    },300)
  },
  provide() {
    return {
      reload: this.reload,
    };
  },
  methods: {
    handleCommand(e){
      const shape = e === 'a' ? "leaf" : "";
      const data = this.treeRootData;
      const id = Math.ceil(Math.random() * 100);
      this.$prompt("请输入名称", "提示", {
        confirmButtonText: "确定",
        cancelButtonText: "取消",
      })
        .then(({ value }) => {
          if (value === '')
            return this.$message.warning("名称不能为空！");
          if (/\s+/g.test(value))
            return this.$message.warning("名称不能包含空格！");

          if(shape === 'leaf'){
            if(!!data.filter(item => (item.name === value && item.type === 'blob'))?.length){
              return this.$message.warning("文件已存在，添加失败！");
            }
          }else{
            if(!!data.filter(item => (item.name === value && item.type === 'tree'))?.length){
              return this.$message.warning("文件夹已存在，添加失败！");
            }
          }
          const treeD = {
            id,
            label: value,
            operation: "add",
            isEdit: true,
            type: shape === "leaf" ? "blob" : "tree",
            filePath: value,
            isLeaf: shape === "leaf",
            name: value,
            fileType:"txt",
            children: [],
          };
          if(shape === "leaf"){
            treeD.content = "";
            treeD.oldContent = "";
            treeD.newContent = "";
            treeD.fileType = 'txt';
            const fix = value.split(".").pop();
            treeD.language = sufix[fix] ? sufix[fix] : "txt"

          }
          treeD.path = treeD.filePath;
          // this.$refs.tree.append(treeD, this.currentData.filePath);
          // setTimeout(() => {
          //   this.handleLeftclick(treeD,treeD)
          // }, 1000);
          // this.$emit("handleAddNode", treeD, this.currentData.filePath); // 触发父组件更改提交界面的数据变化
          this.handleAddNode(treeD,""); // 触发父组件更改提交界面的数据变化
          setTimeout(() => {
            this.$refs.vuetree.handleLeftclick(treeD,treeD)
          },1000)
        })
        .catch((err) => {
          console.error("error:",err)
        });
    },
    handleDiff(data) {
      this.fileInfoParams = data;
      this.fileInfo = data;
      this.$refs.monaco.setDiff(data.oldContent, data.newContent);
    },
    downloadUrl(string) {
      window.location.href = string;
    },
    // 更新提交界面的数据状态
    treeNodeModify(data) {
      const getData = async () => {
        const data = await readAll("git");
        if (data && Array.isArray(data)) {
          this.submitAarry = data
            .filter(
              (name) =>
                name.name !== "treeData" && name.name !== "undefined" && name.d
            )
            .map((item) => item.d);
        }
      };
      getData();
      // const newSubArray = [];
      // function findModifyObj(data) {
      //   data.forEach((item) => {
      //     if (item.isLeaf && item.operation) {
      //       newSubArray.push(item);
      //     } else {
      //       findModifyObj(item.children);
      //     }
      //   });
      // }
      // findModifyObj(data);
    },
    // 新增文件
    handleAddNode(treeD, filePath) {
      this.treeRootData = RecurveAddNode(this.treeRootData, treeD, filePath);
      this.treeData = this.treeRootData;
      this.updateData(treeD);

      setTimeout(async () => {
        await add({
          tableName: "git",
          data: { name: "treeData", data: [...this.treeRootData] },
        });
        this.treeNodeModify(this.treeRootData);
      })
    },

    updateData(d) {
      // const k = await readByMainKey({ tableName: "git", key: "treeData" });
      // const d = { ...k, ...data };
      // this.treeData = matterTree(d.tree);
      // this.treeRootData = JSON.parse(JSON.stringify(this.treeData));
      add({
        tableName: "git",
        data: { name: encodeURI(d.filePath), d },
      });
    },

    loadFrame() {
      document.getElementById("nbview").contentWindow.postMessage(this.fileInfo, '*');
    },
    // 删除新增的文件
    handleDeleteAddNode(valueNode) {
      this.treeRootData = RecurveQueryDelete(
        this.treeRootData,
        valueNode.filePath
      );
      this.treeNodeModify(this.treeRootData);
    },
    // 删除已有的文件
    handleDeleteOldNode(valueNode) {
      this.treeRootData = RecurveQuery(this.treeRootData, valueNode.filePath);
      this.treeNodeModify(this.treeRootData);
    },
    clearAllData() {
      const clear = async () => {
        const data = await readAll("git");
        if (data && Array.isArray(data)) {
          this.fileInfoParams.filePath = null
          this.submitAarry = data.map((item) => {
            // if (item.d) {
            remove({
              tableName: "git",
              key: encodeURI(decodeURIComponent(item.name)),
            });
            // }
          });
        }
        setTimeout(() => {
          this.treeNodeModify();
        }, 200);
      };
      clear();
    },
    // 放弃所有更改
    handleResetAll() {
      this.$confirm(`确定要放弃所有文件的更改吗?`, "提示", {
        confirmButtonText: "确定",
        cancelButtonText: "取消",
        type: "warning",
      })
        .then(async () => {
          // this.handleFileInfo(params);
          this.clearAllData();
          this.isDiff = false;
          this.$refs.monaco.setCloseDiff(false);
        })
        .catch(() => {
        });
    },
    handleEdit() {
      const value = this.$refs.monaco.getVal();
      this.fileInfo.newContent = value;
      this.hideNbView = !this.hideNbView
    },
    // 放弃单个的更改---逻辑未处理
    handleload(value) {
      this.$confirm(`确定要放弃 ${value.path} 中的更改吗?`, "提示", {
        confirmButtonText: "确定",
        cancelButtonText: "取消",
        type: "warning",
      })
        .then(() => {
          // this.handleFileInfo(params);
          // if (value.sha) {
          remove({
            tableName: "git",
            key: encodeURI(value.path),
          });
          this.treeNodeModify();
          // }
        })
        .catch(() => {
        });
    },

    // 提交事件。 接口地址：https://www.apifox.cn/apidoc/shared-906bc1af-5c25-4b9d-b0e0-4d1176ffe4a9/api-30499442。operation为update/add/delete三种
    async handleSubmit() {
      if (!this.textContent.trim())
        return this.$message.warning("提交信息不能为空!");
      try {
        this.subLoading = true;
        await this.$axios
          .post(
            `/api/v1/repos/${this.params.owner}/${this.params.project}/contents/commit`,
            {
              files: this.submitAarry.filter(item => item.type === "blob").map((item) => {
                // item.content = Buffer.from(item.newContent).toString("base64");
                // item.operation = "update";
                return {
                  filePath: item.path,
                  content: Buffer.from(item.newContent).toString("base64"),
                  operation: item.sha ? "update" : "add",
                };
              }),
              branch: (this.params.branch),
              msg: this.textContent,
              currentBranch: (this.params.branch),
              commitBranch: (this.params.branch)
            }
          )
          .then(({ data, status }) => {
            if (status === "200" || status === 200) {
              this.subLoading = false;
              this.submitAarry = []; // 提交的列表
              this.$message.success("提交成功!");
              this.textContent = "";

              this.clearAllData();
              this.getCodeTree();
            }
          });
      } catch (error) {
        this.subLoading = false;
        // 友好提示
        if (error.response.status === 404 || error.response.status === 400) {
          this.$message.error(error.response.data.message);
        }
        console.error("error", error);
      }
    },

    async handleServicePreview(){
      const fromData = new FormData()
      fromData.append('projectId', this.params.project)
      fromData.append('userId', this.params.owner)
      fromData.append('type', 5)
      fromData.append('port', 5060)
      const config = {
        headers: {
          'Content-Type': 'multipart/form-data'
        }
      }
      this.$axios.post(`${API}/vscode/getPort`,fromData,config)
      .then(({ status, data: { data } }) => {
        if (status === "200" || status === 200) {
          let link = document.createElement('a')
          document.body.appendChild(link)
          link.href = data[0].targetUrl
          link.target = "_blank"
          let evt = document.createEvent("MouseEvents")
          evt.initEvent("click", false, false)
          link.dispatchEvent(evt)
          document.body.removeChild(link)
        } else {
          this.$message({
            message: '服务预览失败',
            type: 'warning'
          });
        }
      });
    },
    // 演示
    async handleReset() {
      this.$confirm(`确定要重置环境吗?`, "提示", {
        confirmButtonText: "确定",
        cancelButtonText: "取消",
        type: "warning",
      })
        .then(async () => {
          this.loading = true
          this.headerParams = {}
          await this.$axios
            .post(`${API}/vscode/delete`, {
              projectId: this.params.project,
              userId: this.params.owner,
            })
            .then(({ status, data: { data } }) => {
              if (status === "200" || status === 200) {
                this.getVscodeData();
              } else {
                this.$message({
                  message: '重置环境失败',
                  type: 'warning'
                });
              }
            });
        })
        .catch((e) => {
          console.error(e);
        });

    },

    // 重置环境
    async handleReset() {
      this.$confirm(`确定要重置环境吗?`, "提示", {
        confirmButtonText: "确定",
        cancelButtonText: "取消",
        type: "warning",
      })
        .then(async () => {
          this.loading = true
          this.headerParams = {}
          await this.$axios
            .post(`${API}/vscode/delete`, {
              projectId: this.params.project,
              userId: this.params.owner,
            })
            .then(({ status, data: { data } }) => {
              if (status === "200" || status === 200) {
                this.getVscodeData();
              } else {
                this.$message({
                  message: '重置环境失败',
                  type: 'warning'
                });
              }
            });
        })
        .catch((e) => {
          console.error(e);
        });

    },
    // 返回按钮
    handleBack() {
      // this.$router.go(-1);
      this.$confirm(`确定要退出编辑吗?`, "提示", {
        confirmButtonText: "确定",
        cancelButtonText: "取消",
        type: "warning",
      })
        .then(async () => {
          window.location.href = `/${this.params.owner}/${this.params.project}`
        })
        .catch((e) => {
          console.error(e);
        });
    },
    // 获取vscode界面的接口
    async getVscodeData() {
      try {
        await this.$axios
          .post(`https://code-kyxt.pcl.ac.cn/vscode/getVscode`, {
            projectId: this.params.project,
            userId: this.params.owner,
            language:this.params.language,
            containers:
              "W3siaW1hZ2UiOiJweXRob24zZDctY29kZXNlcnZlcjp2MS4wIiwiY3B1TGltaXQiOjIuMCwiY3B1UmVxdWVzdCI6MC4zLCJtZW1vcnlMaW1pdCI6IjIwNDhNIiwibWVtb3J5UmVxdWVzdCI6IjEwMjRNIiwicmVzb3VyY2VMaW1pdCI6IjEwMDAwSyIsInN0YXJ0VGltZSI6MjAwLCJ0eXBlIjoibWFpbiJ9XQo=",
            tpiGitURL:
              this.params.git,
          })
          .then(({ status, data: { data } }) => {
            if (status === "200" || status === 200) {
              this.loading = false;
              this.headerParams = data[0];
              this.Time();
            }
          });
      } catch (error) {
        console.error("获取目录树：", error);
      }
    },
    handleNode(arr = []) {
      this.submitAarry = this.submitAarry.concat(arr);
    },

    // 获取组件树的接口
    async getCodeTree() {
      try {
        await this.$axios
          .get(
            `/api/v1/repos/${this.params.owner}/${this.params.project}/git/codes/${(this.params.branch)}`,
            {
              headers: {
                Authorization: "token 3551a6fb122b0278133963474fcf18bfeef8baba",
              },
              params: {
                currentBranch: (this.params.branch)
              }
            }
          )
          .then(async ({ status, data }) => {
            if (status === "200" || status === 200) {
              this.treeLoading = false;
              const k = await readByMainKey({
                tableName: "git",
                key: "treeData",
              });

              let dataAll = [];
              try{
                dataAll = await readAll("git");
              }catch(e){
                console.error("git/codes/datall:",e)
              }
              data.tree.push(... dataAll.filter(item => {
                if(item.name.split("/").length === 1 && item.name !== 'undefined' && !data.tree.some(i => i.path === item.name)){
                  return true
                }
                return false
              })?.filter(item => item.d)?.map(item => item.d))

              function loop(d) {
                let path = d[0].path;
                path = path.replace(d[0].name, "");
                if (path.indexOf("/") > -1) {
                  dataAll.some(_item => {
                    let path2 = decodeURIComponent(_item.name).split("/");
                    path2.pop();
                    path2 = path2.join("/") + "/";
                    if (path === path2 && !d.some(i => i.path === path)) {
                      d.unshift(_item.d)
                    }
                  })
                }
                d.map(item => {
                  let path = item.path;
                  if (path) {
                    if (!!item.children.length && item.children) {
                      loop(item.children);
                    }
                  }
                })
              }
              loop(data.tree);

              this.treeData = matterTree(data.tree);
              this.treeRootData = JSON.parse(JSON.stringify(this.treeData));
              add({
                tableName: "git",
                data: { name: "treeData", data: [...this.treeData] },
              });
              this.treeNodeModify();
            }
          });
      } catch (error) {
        console.error("code：", error);
        if (error.response && error.response.status === 403) {
          this.$message.error("权限不足");
        }
      }
    },

    // 点击组件树触发的父组件方法。目前思路。
    // ①如果是新增的文件，点击文件设置monaco编辑器内容为空
    // ②如果是已经存在的文件，点击文件调取接口 this.handleFileInfo
    //
    handleChangFile(params, treeData) {
      this.hideNbView = true;
      this.fileInfoParams = params;
      const getContent = async () => {
        const k = await readByMainKey({
          tableName: "git",
          key: encodeURI(params.filePath),
        });
        if (k.d) {
          this.fileInfo = k.d;
          this.monacaValue = k.d.newContent;
          this.$refs.monaco.setLanguage(k.d.language)
        } else if (params.sha) {
          this.handleFileInfo(params);
        } else {
          this.monacaValue = "";
        }
      };
      getContent();

      // const editorGetValue = Buffer.from(
      //   `${this.$refs.monaco.getVal()}`
      // ).toString("base64"); // 当前编辑器的值
      // console.log("handleChangFile:",params,treeData)
      // const BufferUtf = this.monacaValue; // 当前文件的值
      // if (params.operation === "add") {
      //   // 是新增并且还没有编辑
      //   this.monacaValue = params.oldContent ? params.oldContent : "";
      //   this.fileInfoParams = params;
      //   this.pathname = params.name;
      // } else if (params.sha && !params.oldContent) {
      //   console.log("treeData", treeData);
      //   this.handleFileInfo(params);
      // } else if (editorGetValue && BufferUtf && editorGetValue !== BufferUtf) {
      //   // 当前编辑器的内容是否原有内容一致
      //   this.$confirm(
      //     `当前${this.fileInfoParams.name}文件未保存, 是否继续?`,
      //     "提示",
      //     {
      //       confirmButtonText: "确定",
      //       cancelButtonText: "取消",
      //       type: "warning",
      //     }
      //   )
      //     .then(() => {
      //       this.handleFileInfo(params);
      //     })
      //     .catch(() => {
      //       console.log("取消成功");
      //     });
      // }
    },

    // 获取文件的内容
    async handleFileInfo(params) {
      try {
        this.monacoLoading = true;
        await this.$axios
          .get(
            `/api/v1/repos/${this.params.owner}/${this.params.project}/contents/${params.path}`,
            {
              params: {
                ref: (this.params.branch),
              }
            }
          )
          .then(({ status, data }) => {
            if (status === "200" || status === 200) {
              this.fileInfo = { ...data };
              this.monacoLoading = false;
              let content = data.content;
              if(isBase64(data.content)){
                content = Buffer.from(data.content, "base64").toString("utf8")
              }
              params.oldContent = content;
              this.fileInfoParams = params;
              this.pathname = data.name;
              this.opts.language = data.language
              this.$refs.monaco.setLanguage(data.language)
              this.monacaValue = content;
            }
          });
      } catch (error) {
        this.monacoLoading = false;
        console.error("获取文件信息失败", error);
      }
    },
    // 提交/代码界面按钮切换事件
    handleselect(value) {
      this.hideNbView = true;
      this.fileInfoParams = {};
      this.diffData = {};
      this.isDiff = value == 1 ? false : true;
      this.$refs.monaco.setCloseDiff(true);
      this.activeKey = value;
    },
    // 中间图标的拖拽
    dragControllerDiv() {
      const resize = document.getElementsByClassName("resize");
      const left = document.getElementsByClassName("leftPaneBase");
      const mid = document.getElementsByClassName("mid");
      const box = document.getElementsByClassName("box");

      for (let i = 0; i < resize.length; i++) {
        // 鼠标按下事件
        resize[i].onmousedown = (e) => {
          //颜色改变提醒
          // resize[i].style.background = '#818181';
          var startX = e.clientX;
          resize[i].left = resize[i].offsetLeft;
          // 鼠标拖动事件
          document.onmousemove = function (e) {
            var endX = e.clientX;
            var moveLen = resize[i].left + (endX - startX) - 60; // （endx-startx）=移动的距离。resize[i].left+移动的距离=左边区域最后的宽度
            var maxT = box[i].clientWidth - resize[i].offsetWidth; // 容器宽度 - 左边区域的宽度 = 右边区域的宽度

            if (moveLen < 132) moveLen = 132; // 左边区域的最小宽度为32px
            // if (moveLen > maxT - 250) moveLen = maxT - 250; //右边区域最小宽度为150px

            // resize[i].style.left = moveLen; // 设置左侧区域的宽度

            for (let j = 0; j < left.length; j++) {
              left[j].style.width = moveLen + "px";
              mid[j].style.width = box[i].clientWidth - moveLen - 1 + "px";
            }
          };

          // 鼠标松开事件
          document.onmouseup = (evt) => {
            this.$refs.monaco.init(this.opts)

            //颜色恢复
            resize[i].style.background = "#d6d6d6";
            document.onmousemove = null;
            document.onmouseup = null;
            resize[i].releaseCapture && resize[i].releaseCapture(); //当你不在需要继续获得鼠标消息就要应该调用ReleaseCapture()释放掉
          };
          resize[i].setCapture && resize[i].setCapture(); //该函数在属于当前线程的指定窗口里设置鼠标捕获
          return false;
        };
      }
    },

    // 更改语言
    changeLanguage(val) {
      this.opts.language = val;
    },
    // 更换主题
    changeTheme(val) {
      this.opts.theme = val;
    },
    // 手动获取值
    getValue() {
      console.log(`输出代码:${this.$refs.monaco.getVal()}`);
    },
    // 保存按钮的回调事件
    handleSave(val) {
      // this.monacaValue = Buffer.from(val).toString("base64");
      const readvalue = async () => {
        const k = await readByMainKey({
          tableName: "git",
          key: encodeURI(this.fileInfoParams.filePath),
        });
        if (k.d) {
          if (!k.d.language)
            k.d.language = this.opts.language;
          k.d.fileType = this.fileInfo.fileType;
          if (val === k.d.oldContent) {
            remove({
              tableName: "git",
              key: encodeURI(this.fileInfoParams.filePath),
            });
          } else {
            k.d.newContent = val;
            add({
              tableName: "git",
              data: { name: encodeURI(this.fileInfoParams.filePath), d: k.d },
            });
          }
        } else {
          this.fileInfoParams.newContent = val;
          this.fileInfoParams.language = this.fileInfo.language;
          this.fileInfoParams.fileType = this.fileInfo.fileType;
          add({
            tableName: "git",
            data: {
              name: encodeURI(this.fileInfoParams.filePath),
              d: this.fileInfoParams,
            },
          });
        }
        this.treeNodeModify();
      };
      readvalue();
      // const base64Content = Buffer.from(val).toString("base64");
      // if (this.fileInfoParams.operation === "add") {
      //   // 说明这条是新增的
      //   this.treeRootData = RecurveUpdateValue(
      //     this.treeRootData,
      //     this.fileInfoParams,
      //     base64Content
      //   );
      //   this.treeNodeModify(this.treeRootData);
      // }
    },
    async reload() {
      this.containerReload = false;
      await this.$nextTick();
      this.containerReload = true;
    },

    handleSetting() {
      this.drawerVisible = true;
    },

    handleChangeModelOptions(value) {
      this.opts.theme = value;
      this.$refs.monaco.setTheme(value)
    },
    handleChangeFontSizeOptions(value) {
      this.opts.fontSize = value;
      this.$refs.monaco.updateOptions(this.opts)
    },

    // 倒计时
    countDown() {
      let h = parseInt((this.headerParams.remainingTime / (60 * 60)) % 24);
      h = h < 10 ? `0${h}` : h;
      let m = parseInt((this.headerParams.remainingTime / 60) % 60);
      m = m < 10 ? `0${m}` : m;
      let s = parseInt(this.headerParams.remainingTime % 60);
      s = s < 10 ? `0${s}` : s;
      this.count = `${h}:${m}:${s}`;
    },
    // 定时器没过1秒参数减1
    Time() {
      this.timer = setInterval(() => {
        this.headerParams.remainingTime -= 1;
        if (this.headerParams.remainingTime === 0) clearInterval(this.timer);
        this.countDown();
      }, 1000);
    },
  },
};
</script>
<style lang="less" scoped>
.hides {
  // visibility: hidden;
  position: relative;
  z-index: -1;
}

.showImg {
  width: 100%;
  height: 100%;
  display: flex;
  justify-content: center;
  align-items: center;
  background: #f1f1f1;
}

.showOffice,
.showPdf {
  width: 100%;
  height: 100%;
}

.showOffice iframe,
.showPdf embed {
  width: 100%;
  height: 100%;
  border: none;
  background: white;
}

.showImg img {
  max-height: 100%;
  max-width: 100%;
}

.ide-header {
  background: #363840;
  width: 100%;
  height: 60px;
  //border: 1px solid #363840;
  display: flex;
  padding-left: 32px;
  justify-content: space-between;
  align-items: center;


  span {
    display: flex;
    align-items: center;
    font-size: 14px;
    font-family: PingFangSC-Medium, PingFang SC;
    font-weight: 500;
    color: #ffffff;
    line-height: 60px;

    .ide-header-svg {
      margin-right: 8px;
      width: 13px;
      height: 19px;
    }
  }
}

.el-icon-switch-button {
  line-height: 60px;
  margin-right: 20px;
  cursor: pointer;
  color: #FFFFFF;
  font-size: 14px;
}

.inlineBlock {
  display: inline-block;
  float: left;
}

.menu-left {
  display: inline-block;
  width: 75px;
  background-color: #F5F5F5;
  height: calc(100vh - 60px);
  //.menu-left-drag-img {
  //  width: 24px;
  //  height: 24px;
  //}

  .menu-left-drag {
    li {
      height: 70px;
    }

    &-svg {
      height: 24px;
      width: 30px;
      color: #7A7A7B;
      margin-bottom: 8px;
    }
  }

  .el-menu {
    border-right: none;

    .el-menu-item {
      padding: 10px 20px;
      height: none;
    }
  }
}

.activeBack {
  background: #ffffff !important;
  padding: 19px 21px 45px 22px !important;
  border: 1px solid #D4D4D5;

  .menu-left-drag-svg {
    color: #000000;
  }
}

.leftPaneBase {
  //background: #e6f1f6;
  //border: 1px solid #dededf;
  height: calc(100vh - 40px);

  &-submit {
    background: #F5F5F6;
    border: 1px solid #D4D4D5;
    height: 52px;
    padding: 0 36px;
    font-size: 14px;
    font-family: PingFangSC-Medium, PingFang SC;
    font-weight: 500;
    color: #000000;
    line-height: 52px;
  }

  &-wrap {

    &-change {
      padding: 0 20px;
      height: 54px;
      line-height: 54px;
      display: flex;
      justify-content: space-between;

      &-svg {
        width: 16px;
        height: 16px;
      }

      &-img {
        cursor: pointer;
        margin-top: 20px;
      }
    }

    &-fileName {
      height: calc(100vh - 508px);
      overflow-y: auto;

      &-item {
        padding: 0 20px;
        height: 34px;
        font-size: 14px;
        font-family: PingFangSC-Regular, PingFang SC;
        font-weight: 400;
        color: #000000;
        line-height: 34px;
        overflow: hidden;
        text-overflow: ellipsis;
        white-space: nowrap;
        display: flex;
        align-items: center;
        justify-content: space-between;

        span {
          cursor: pointer;
        }

        &-img {
          cursor: pointer;
          opacity: 0;
        }
      }

      &-item:hover {
        background: #F6F9FC;

        span {
          color: #409EFF;
        }

        .leftPaneBase-wrap-fileName-item-img {
          opacity: 1;
          color: #2285d0;
        }
      }
    }

    &-info {
      padding: 0 20px;
      font-size: 14px;
      font-family: PingFangSC-Regular, PingFang SC;
      font-weight: 400;
      color: #000000;
      margin: 30px 0 10px;
    }

    &-input {
      padding: 0 20px;
      margin-bottom: 10px;

      .el-textarea__inner {
        background: #dae9f3;
        max-height: 200p "max-content";
        max-height: 200px;
      }
    }

    &-infoSubmit {
      width: 100%;
      background: #363742;
      border: 0px;
    }
  }
}

.resize {
  //position: absolute;
  //top: 145px;
  //display: inline-block;
  //width: 2px;
  //height: calc(100% - 145px);
  //border-radius: 5px;
  //margin-top: -51px;
  //cursor: pointer;
  background: #dce3e8;
  //background-position: center;
  /* z-index: 99999; */
  //font-size: 14px;
  //color: white;
  //margin-left: -1px;
  //float: left;
  //cursor: col-resize;
  //z-index: 999;

  position: absolute;
  top: 160px;
  display: inline-block;
  width: 2px;
  height: calc(100vh - 100px);
  border-radius: 5px;
  margin-top: -51px;
  margin-left: -1px;
  cursor: col-resize;
  z-index: 999;
}

.mid {
  float: left;
  width: calc(100% - 300px);
  /*右侧初始化宽度*/
  height: calc(100vh - 42px);
  background: #fff;
  position: relative;
}

.box {
  width: 100%;
  overflow: hidden;
}

.empty-tip {
  text-align: center;
  width: 100%;
  height: 100%;
  display: flex;
  align-items: center;
  justify-content: center;
  flex-direction: column;
  color: #999;
}

.mid-header {
  height: 54px;
  background: #f0f0f0;
  border: 1px solid #dadadb;
  line-height: 54px;
  display: flex;
  width: 100%;
  justify-content: space-between;
  padding: 0 20px;
  font-size: 14px;
  font-family: PingFangSC-Regular, PingFang SC;
  font-weight: 400;
  color: #000000;

  span:nth-child(2) {
    .el-icon-s-tools {
      margin-left: 5px;
      cursor: pointer;
    }
  }
}

.drawer-wrap {
  font-size: 14px;
  font-family: PingFangSC-Regular, PingFang SC;

  .el-drawer {
    background-color: #e6f1f6 !important;
  }

  .drawer-wrap-span {
    height: 20px;
    font-weight: 400;
    color: #7A7A7B;
    line-height: 20px;
  }

  .drawer-wrap-span2 {
    height: 20px;
    font-weight: 400;
    color: #000000;
    line-height: 20px;
  }
}

.code-fast {
  font-weight: 500;
  margin-bottom: 30px;
  height: 20px;
  font-size: 14px;
  font-family: PingFangSC-Medium, PingFang SC;
  font-weight: 500;
  color: #000000;
  line-height: 20px;
}

.drawer-section {
  height: 100%;
  //background: #e6f1f6;
  background: #f0f0f0;
  padding: 25px 26px 0;

  .el-select {
    width: 140px;
    font-weight: 400;
    color: #000000;
  }

  &-stydrawer-sectionle-wrap {
    display: flex;
    align-items: center;
    justify-content: space-between;
    margin-bottom: 20px;
  }

  &-style {
    border-bottom: 1px solid #b1bacd;
    padding-bottom: 30px;
    margin-bottom: 30px;

    &-wrap {
      display: flex;
      justify-content: space-between;
      align-items: center;
      margin-bottom: 20px;
    }
  }

  .drawer-fast {
    &-flex {
      display: flex;
      justify-content: space-between;
      margin-bottom: 20px;
      font-weight: 400;
      color: #333333;
      line-height: 14px;
    }
  }
}

.vscode-header {
  display: flex;
  justify-content: space-between;
  line-height: 40px;
  font-size: 14px;
  font-family: PingFangSC-Medium, PingFang SC;
  font-weight: 500;
  color: #ffffff;
  height: 40px;
  padding: 0 20px;
  background: #051019;
}

.color {
  color: #f8cd47;
}

.mr5 {
  margin-right: 5px;
}

.mr40 {
  margin-right: 40px;
}

.cspointe {
  cursor: pointer;
}

.iframe-wrap {
  width: 100%;
  height: calc(100vh - 40px);
  border: none;
}

.vscode-loading-wrap {
  width: 100%;
  height: 100vh;
}

.no-commit {
  color: #999;
  text-align: center;
  padding-top: 20px;
}
</style>
<style scoped>
/deep/.el-drawer {
  box-shadow: none;
}

:focus {
  outline: none;
}

.icon {
  width: 1em;
  height: 1em;
  vertical-align: -0.15em;
  fill: currentColor;
  overflow: hidden;
}

#nbview {
  position: absolute;
  width: 100%;
  height: calc(100% - 54px);
  left: 0;
  top: 54px;
  z-index: 1;
  background: #FFF;
  border: none;
}
</style>
