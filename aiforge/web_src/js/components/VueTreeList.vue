<template>
  <div class="treeNode">
    <span @click="addNode">添加节点</span>
    <vue-tree-list
      :model="data"
      default-tree-node-name="new file"
      default-leaf-node-name="new folder"
      :default-expanded="false"
      @click="onClick"
      @change-name="onChangeName"
      @delete-node="onDel"
      @add-node="onAddNode"
    />
  </div>
</template>

<script>
import {VueTreeList, Tree, TreeNode} from 'vue-tree-list';

export default {
  components: {
    VueTreeList
  },
  props: {
    treeData: {
      type: Array,
      default: () => []
    },
    fileInfoParams: {
      type: Object,
      default: () => {}
    }
  },
  data() {
    return {
      newTree: {},
      fileParams: {}, // 文件相关的信息
      data: new Tree([
        {
          name: 'Node 1',
          id: 1,
          pid: 0,
          children: [
            {
              name: 'Node 1-2',
              id: 2,
              isLeaf: true,
              pid: 0
            }
          ]
        },
        {
          name: 'Node 2',
          id: 3,
          pid: 0,
        },
        {
          name: 'Node 3',
          id: 4,
          pid: 0
        }
      ])
    };
  },
  watch: {
    treeData(val) {
      this.data = new Tree(val);
    },
    fileInfoParams(val) {
      this.fileParams = val;
    }
  },
  mounted() {
  },
  methods: {
    onDel(node) {
      const paramsInfo = [];
      if (!node.children?.length > 0) { // 文件还是文件夹
        paramsInfo[0] = {...node};
      } else { // 文件夹
        node.children.filter((item) => item.sha).map((items) => { // 过滤是新增的元素进行删除
          paramsInfo.push({
            filePath: items.path,
            content: '',
            name: items.name,
            operation: 'delete'
          });
        });
      }
      this.$emit('deleteNode', paramsInfo);
      node.remove();
    },

    onChangeName(params) {
    },

    onAddNode(params) {
    },

    // 点击每个文件/夹事件
    onClick(params) {
      if (!params.isLeaf) return;
      if (this.fileParams?.sha === params.sha) return;
      this.$emit('handleChangFile', params);
    },

    addNode() {
      const node = new TreeNode({name: 'new node', isLeaf: false});

      if (!this.data.children) this.data.children = [];
      this.data.addChildren(node);
    },
  }
};
</script>

<style lang="less" rel="stylesheet/less">
  .vtl {
    .vtl-drag-disabled {
      background-color: #d0cfcf;
      &:hover {
        background-color: #d0cfcf;
      }
    }
    .vtl-disabled {
      background-color: #d0cfcf;
    }
  }
</style>

<style lang="less" rel="stylesheet/less" scoped>
  .icon {
    &:hover {
      cursor: pointer;
    }
  }

  .muted {
    color: gray;
    font-size: 80%;
  }
</style>

<style scoped>
.treeNode{
  padding:20px 20px 0;
  overflow-y: auto;
  overflow-x: hidden;
  text-overflow: ellipsis;
  white-space: nowrap;
}
</style>
