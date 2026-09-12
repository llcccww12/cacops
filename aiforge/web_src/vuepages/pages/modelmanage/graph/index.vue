<template>
  <div>
    <div v-if="emptyPage" style="padding-top:50px">
      <NotFound></NotFound>
    </div>
    <div v-else>
      <div class="ui container">
        <div ref="graphContainerRef" class="graph-container"></div>
      </div>
    </div>
  </div>
</template>

<script>
import NotFound from '~/components/NotFound.vue';
import { getModelInfoByName, getModelEvolutionMap } from '~/apis/modules/modelmanage';
import { getUrlSearchParams } from '~/utils';
import { ModelGraph } from './model-graph';
import './model-graph.css';

let modelGraph = null;

export default {
  data() {
    return {
      emptyPage: false,

      modelName: '',
      repoOwnerName: location.pathname.split('/')[1],
      repoName: location.pathname.split('/')[2],
      modelData: {},
    };
  },
  components: { NotFound },
  methods: {},
  beforeMount() {
    const urlParams = getUrlSearchParams();
    if (urlParams.name) {
      this.modelName = urlParams.name;
      this.loading = true;
      getModelInfoByName({ repo: `/${this.repoOwnerName}/${this.repoName}`, name: this.modelName }).then(res => {
        const data = res.data;
        this.loading = false;
        if (data && data.length) {
          const model = data[0];
          this.modelData = model;
          getModelEvolutionMap({
            repo: `/${this.repoOwnerName}/${this.repoName}`,
            id: model.id,
          }).then(res => {
            const data = res.data;
            function runData(node) {
              if (node.Type == 1) {
                const model = node.Model || {};
                node.type = 'model';
                node.name = model.name;
                node.isCurrent = node.IsCurrent;
                node.isParent = node.IsParent;
                node.isDerive = !node.IsCurrent && !node.IsParent;
                node.creator = model.userName;
                const link = `/${model.repoOwnerName}/${model.repoName}/modelmanage/model_readme_tmpl?name=${model.name}`;
                model.link = link;
                node.isPrivate = node.IsPrivate;
                node.isCanOper = node.IsCanOper;
                if (!node.isCurrent && !node.isParent) {
                  node.link = link;
                }
                if (node.isParent) {
                  const models = node.Models4Parent || [];
                  models.map(item => {
                    item.link = `/${item.repoOwnerName}/${item.repoName}/modelmanage/model_readme_tmpl?name=${item.name}`;
                  })
                  node.name = models.map(itm => itm.name).join(',');
                  node.link = '';
                }
              }
              if (node.Type == 0) {
                node.type = 'repo';
                node.name = node.RepoOwnerName + ' / ' + `<span style="font-weight:bold">${node.RepoDisplayName}</span>`;
                node.link = `/${node.RepoOwnerName}/${node.RepoName}`;
                node.isPrivate = node.IsPrivate;
                node.isCanOper = node.IsCanOper;
              }
              node.children = node.Next || [];
              for (let i = 0, iLen = node.children.length; i < iLen; i++) {
                const child = node.children[i];
                child._parent = node;
                runData(child);
              }
            }
            runData(data);
            this.$nextTick(() => {
              modelGraph = new ModelGraph();
              modelGraph.init(this.$refs.graphContainerRef, data, {});
            });
          }).catch(err => {
            console.log(err);
          });
        } else {
          this.emptyPage = true;
        }
      }).catch(err => {
        console.log(err);
      });
    } else {
      this.emptyPage = true;
    }
  },
  mounted() { },
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
