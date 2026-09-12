<template>
  <div class="myTrees">
    <el-tree
      ref="tree"
      :data="treeData"
      node-key="filePath"
      empty-text="暂无数据"
      highlight-current
      @node-click="handleLeftclick"
    >
      <div slot-scope="{ node , data }" class="custom-tree-node">
        <span>
          <i :class="getIcon(node.data)" />
          {{ node.label }}
        </span>
        <span>
          <el-dropdown trigger="click" class="el-dropdown"
           v-show="data.type == 'tree'"
          >
            <span class="el-dropdown-link">
              <svg class="icon el-dropdown-svg" aria-hidden="true">
                <use xlink:href="#icon-a-bianzu31"></use>
              </svg>
            </span>
            <el-dropdown-menu slot="dropdown">
              <el-dropdown-item
                @click.native="addChildNode('leaf',data)"
                >新建文件</el-dropdown-item
              >
              <el-dropdown-item
                @click.native="addChildNode('',data)"
                >新建文件夹</el-dropdown-item
              >
              <!-- <el-dropdown-item @click.native="deleteNode">删除</el-dropdown-item> -->
            </el-dropdown-menu>
          </el-dropdown>
        </span>
      </div>
    </el-tree>
  </div>
</template>
<!-- file-icon ide-icon word-icon dark-blue -->
<link rel="stylesheet" href="/web_src/js/components/treeIcon.css" />
<script>
import {icons} from "./icons";
export const sufix = {
  "py":        "python",
  "c":         "cpp",
  "h":         "cpp",
  "g4":        "cpp",
  "sy":        "cpp",
  "cc":        "cpp",
  "cxx":       "cpp",
  "c++":       "cpp",
  "cu":        "cpp",
  "cpp":       "cpp",
  "dynamips":  "cpp",
  "java":      "java",
  "php":       "php",
  "html":      "html",
  "css":       "css",
  "scss":      "scss",
  "go":        "go",
  "r":         "r",
  "graphql":   "graphql",
  "swift":     "swift",
  "xml":       "xml",
  "yaml":      "yaml",
  "json":      "json",
  "lua":       "lua",
  "scheme":    "scheme",
  "less":      "less",
  "ini":       "ini",
  "jpg":       "jpg",
  "jpeg":      "jpeg",
  "png":       "png",
  "gif":       "gif",
  "webp":      "webp",
  "bmp":       "bmp",
  "avi":       "avi",
  "mp4":       "mp4",
  "mov":       "mov",
  "mp3":       "mp3",
  "wav":       "wav",
  "ogg":       "ogg",
  "coffee":    "coffeescript",
  "litcoffee": "coffeescript",
  "js":        "javascript",
  "vue":       "javascript",
  "ejs":       "html",
  "cs":        "csharp",
  "kt":        "kotlin",
  "md":        "markdown",
  "sql":       "mysql",
  "ctrl":      "mysql",
  "m":         "objective-c",
  "mm":        "objective-c",
  "pas":       "pascal",
  "perl":      "perl",
  "pl":        "perl",
  "rb":        "ruby",
  "rs":        "rust", "rust": "rust",
  "tsx":   "typescript",
  "ipynb": "json",
  "sh":    "shell",
  "bash":  "shell",
 };
export default {
  name: "List",
  props: {
    treeListData: {
      type: Array,
      default: () => [],
    },
    fileInfoParams: {
      type: Object,
      default: () => {},
    },
  },
  data() {
    return {
      fileParams: {},
      treeData: [],
      isShow: false,
      currentData: "",
      currentNode: "",
      menuVisible: false,
      firstLevel: false,
      lastLevel: false,
      filterText: "",
      isLeaf: false,
    };
  },
  watch: {
    treeListData(val) {
      this.treeData = val;
    },
    fileInfoParams(val) {
      this.fileParams = val;
    },
  },

  methods: {
    // 鼠标左击事件
    handleLeftclick(data, node) {
      this.currentData = data;
      this.currentNode = node;
      this.firstLevel = false;
      this.isLeaf = data.isLeaf;
      this.lastLevel = false;
      if (data.type === 'tree') return;
      // if (data.sha) {
      this.$emit("handleChangFile", data, this.treeData);
      // }
    },
    getIcon(data){
      let icon = '';
      if(data.type === 'tree'){
        return 'fa fa-folder ide-icon ide-icon-folder'
      }
      try {
        let suffix = data.name.split(".").pop();
        if(data.name.indexOf(".") > -1){
          suffix = "." + suffix
        }
        icons.forEach(element => {
            if(element[2].test(suffix)){
              icon = element[0] + ' ' + element[1].join(' ');
              throw('')
            }
        })
      } catch (error) {

      }
      return "file-icon ide-icon " + icon;
    },

    // 增加子级节点事件
    addChildNode(shape,data) {
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
            if(!!data.children.filter(item => (item.name === value && item.type === 'blob'))?.length){
              return this.$message.warning("文件已存在，添加失败！");
            }
          }else{
            if(!!data.children.filter(item => (item.name === value && item.type === 'tree'))?.length){
              return this.$message.warning("文件夹已存在，添加失败！");
            }
          }
          const treeD = {
            id,
            label: value,
            operation: "add",
            isEdit: true,
            type: shape === "leaf" ? "blob" : "tree",
            filePath: `${this.currentData.filePath}/${value}`,
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
          this.$refs.tree.append(treeD, this.currentData.filePath);
          setTimeout(() => {
            this.handleLeftclick(treeD,treeD)
          }, 1000);
          this.$emit("handleAddNode", treeD, this.currentData.filePath); // 触发父组件更改提交界面的数据变化
        })
        .catch(() => {});
    },
    // 删除节点
    deleteNode() {
      this.$confirm(`确定删除当前文件,是否继续?`, "提示", {
        confirmButtonText: "确定",
        cancelButtonText: "取消",
        type: "warning",
      })
        .then(() => {
          if (this.currentData.isLeaf) {
            // isLeaf为true 代表文件类型
            if (this.currentData.operation === "add") {
              // 新增的节点删去
              this.$emit("handleDeleteAddNode", this.currentData);
            } else if (this.currentData.sha) {
              // 原本存在的数据 触发父组件更新已存在数据的状态
              this.$emit("handleDeleteOldNode", this.currentData);
            }
          }
          this.$refs.tree.remove(this.currentNode);
        })
        .catch(() => {});
    },
  },
};
</script>
<style lang="less" scoped>
@import "./treeIcon.css";
.myTrees {
  /*background: transparent;*/
  height: calc(100% - 90px);
  overflow: auto;
  background: #FFFFFF;
}
.el-tree {
  /* padding: 20px; */
  /*background: transparent;*/
  background: #FFFFFF;
  color: black;
}
.el-tree .is-current{
  background: #f5f7fa;
}

.custom-tree-node {
  flex: 1;
  display: flex;
  align-items: center;
  justify-content: space-between;
  padding-right: 8px;
  font-size: 14px;
  height: 34px;
  line-height: 34px;
  .el-dropdown-svg{
    height: 16px;
    width: 16px;
    color: #979797;
  }
  .el-dropdown-svg:hover{
    color: #007aff;
  }
}
/deep/ .el-tree-node__content{
  height: 34px;
  line-height: 34px;
  font-size: 14px;
  font-family: PingFangSC-Regular, PingFang SC;
  font-weight: 400;
}
/deep/ .el-tree-node.is-current > .el-tree-node__content{
  background: #F5F7FA;
  color: #2285D0 !important;
  font-family: PingFangSC-Regular, PingFang SC;
  font-weight: 400;
  font-size: 14px;
  color: #2285D0;
}
.el-dropdown-menu{
  padding: 0px !important;
}
.el-popper[x-placement^=bottom] .popper__arrow {
  display: none !important;
}
.el-dropdown-menu__item{
  color: #333333 !important;
  height: 34px !important;
}
li.el-dropdown-menu__item:hover {
  background: #f2f2f2 !important;
}
.icon {
  width: 1em;
  height: 1em;
  vertical-align: -0.15em;
  fill: currentColor;
  overflow: hidden;
}
</style>
<style>
  .myTrees .el-tree-node__expand-icon{
    visibility: hidden;
  }
  .monaco-editor .line-numbers{
    color: #999;
  }
</style>
