<template>
  <div>
    <div class="__mobile-tip" >
      <img style="width: 132px;height: 97px;" src="/img/model/pc-view.png" alt="">
      <div style="margin-top: 2rem;">{{$t('useInPcWeb')}}</div>
    </div>
    <div class="ui container __content-box">
      <div ref="graphContainerRef" class="graph-container"></div>
    </div>
  </div>
</template>

<script>
import { getModelEvolutionMap } from '~/apis/modules/modelmanage';
import { lang } from '~/langs';
import { ModelGraph } from './model-graph';
import './model-graph.css';

let modelGraph = null;

export default {
  props: {
    dataObj: { type: Object, default: () => ({}) },
  },
  data() {
    return {
      repoOwnerName: location.pathname.split('/')[1],
      repoName: location.pathname.split('/')[2],
    };
  },
  components: {},
  methods: {
    getGraphData() { 
      getModelEvolutionMap({
        aimodel_id: this.dataObj.id,
      }).then(res => {
        const data = res.data;
        function runData(node) {
          if (node.type == 1) { //模型
            const model = node.aimodel_info || {};
            node.Model = node.aimodel_info; //适配之前Node结构数据
            node.type = 'model';
            node.name = lang==='zh-CN' ? model.alias: model.name;
            node.isCurrent = node.is_current;
            node.isParent = node.is_parent;
            node.isDerive = !node.is_current && !node.is_parent;
            node.creator = model.owner_name;
            const link = `/models/detail/${model.owner_name}/${model.name}`;
            model.link = link;
            node.isPrivate = node.is_private;
            node.isCanOper = node.visible;
            if (!node.isCurrent && !node.isParent) {
              node.link = link;
            }
            if (node.isParent) {
              node.Models4Parent = node.parents || []; //适配之前Node结构数据
              for (let i = 0, iLen = node.Models4Parent.length; i < iLen; i++) {
                const parent = node.Models4Parent[i];
                parent._child = node;
                runData(parent);
              }
              const models = node.parents || [];
              models.map(item => {
                item.link = `/models/detail/${item.aimodel_info.owner_name}/${item.aimodel_info.name}`;
              })
              node.name = models.map((itm) => {
                return lang==='zh-CN' ? itm.aimodel_info.alias : itm.aimodel_info.name
              }).join(',');
              node.link = '';
            }
          }
          if (node.type == 0) {  //训练任务
            node.type = 'task';
            const task = node.task_info || {};
            node.Task = node.task_info; //适配之前Node结构数据
            node.name = task.display_job_name;
            node.isCanOper = node.visible;
            node.isPrivate = node.is_private;
          }
          node.children = node.next || [];
          for (let i = 0, iLen = node.children.length; i < iLen; i++) {
            const child = node.children[i];
            child._parent = node;
            runData(child);
          }
        }
        if(data.code === 0){
          runData(data.data);
          this.$nextTick(() => {
            modelGraph = new ModelGraph();
            modelGraph.init(this.$refs.graphContainerRef, data.data, {});
          });
        }else{
          this.$message.error(data.message);
        }
        
      }).catch(err => {
        console.log(err);
      });
    },
  },
  beforeMount() {
    
  },
  mounted() { 
    this.getGraphData()
  },
  beforeDestroy() { },
};
</script>

<style scoped lang="less">
.graph-container {
  position: relative;
  border: 1px solid rgb(225, 227, 230);
  border-radius: 5px;
  margin-top: 32px;
  height: 75vh;
  overflow: hidden;
}
</style>
