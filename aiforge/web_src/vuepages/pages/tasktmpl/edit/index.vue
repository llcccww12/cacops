<template>
  <div class="tasktmpl-edit-c">
    <div class="tasktmpl-edit-title-c">
      <div class="tasktmpl-edit-title">{{ $t('taskTmplObj.editTmplFile') }}</div>
      <div class="tasktmpl-edit-title-tips">
        <i class="ri-information-line"></i>
        <span>{{ $t('taskTmplObj.dragSortTips') }}</span>
      </div>
    </div>
    <div class="tasktmpl-tab-bar-c" ref="tabBarRef">
      <div class="nav-left" v-show="tabBar.showNav" @click="navScroll('left')">
        <i class="el-icon-d-arrow-left"></i>
      </div>
      <div class="list-group-tabs-c">
        <draggable :list="tabs" class="list-group-tabs" ref="tabBarTabsRef" ghost-class="ghost" animation="200"
          @start="dragging = true" @end="dragging = false">
          <transition-group type="transition" :name1="'flip-list'" :name="!dragging ? 'flip-list' : null">
            <div class="list-group-tabs-item" v-for="item in tabs" :key="item.key" @click="changeNav(item)"
              :class="item.key == tabsValue ? 'is-active' : ''">
              <div class="name" :title="item.name">{{ item.name }}</div>
              <i class="ri-close-fill" @click.stop.prevent="removeTmpl(item)"></i>
            </div>
          </transition-group>
        </draggable>
      </div>
      <div class="nav-right" v-show="tabBar.showNav" @click="navScroll('right')">
        <i class="el-icon-d-arrow-right"></i>
      </div>
      <div class="add-btn" @click="addTmpl">
        <i class="ri-add-line"></i>
        <span v-if="!tabs.length">{{ $t('taskTmplObj.addTmpl') }}</span>
      </div>
    </div>
    <div class="list-group-content">
      <div class="list-group-content-item" v-for="(item, index) in tabs" :key="item.key" v-show="item.key == tabsValue">
        <TmplEdit v-model="tabs[index]" :key="item.key"></TmplEdit>
      </div>
    </div>
  </div>
</template>

<script>
import draggable from 'vuedraggable';
import TmplEdit from '../components/TmplEdit.vue';
import { uuidv4 } from '~/utils';
import { TaskTmplTools } from '../tools';

const taskTmplTools = new TaskTmplTools();

export default {
  components: { draggable, TmplEdit },
  data() {
    return {
      repoOwnerName: location.pathname.split('/')[1],
      repoName: location.pathname.split('/')[2],
      tabs: [],
      tabsValue: '',
      dragging: false,

      tabBar: {
        showNav: false,
      },
      editorReady: false,
    };
  },
  watch: {
    tabs: {
      deep: true,
      handler(nVal, oVal) {
        if (!this.editorReady) return;
        this.tabsUpdateTimer && clearTimeout(this.tabsUpdateTimer);
        this.tabsUpdateTimer = setTimeout(() => {
          this.tabBarLayoutUpdate();
          this.renderText();
        }, 50);
      }
    }
  },
  methods: {
    changeNav(item) {
      this.tabsValue = item.key;
    },
    getRandomTmplName() {
      return 'New Tmpl ' + uuidv4().slice(-4);
    },
    addTmpl() {
      const key = uuidv4();
      this.tabs.push({
        key,
        ...taskTmplTools.transformData({
          Name: this.getRandomTmplName(),
        })
      });
      this.tabsValue = key;
      setTimeout(() => {
        this.tabBar.showNav && this.navScroll('right', { useMax: true });
      }, 30);
    },
    removeTmpl(item) {
      this.$confirm(this.$t('taskTmplObj.deleteTmplTips'), this.$t('tips'), {
        confirmButtonText: this.$t('confirm'),
        cancelButtonText: this.$t('cancel'),
        type: 'warning',
        lockScroll: false,
      }).then(() => {
        const findIndex = this.tabs.findIndex(_item => _item.key == item.key);
        const findTab = this.tabs.splice(findIndex, 1);
        if (findTab[0].key == this.tabsValue) {
          this.tabsValue = this.tabs[findIndex]?.key || this.tabs[findIndex - 1]?.key || '';
        }
        setTimeout(() => {
          this.tabBarLayoutUpdate();
        }, 100);
      }).catch(err => { });
    },
    renderText() {
      this.renderTimer && clearTimeout(this.renderTimer);
      this.renderTimer = setTimeout(() => {
        const data = this.tabs.map(item => {
          return taskTmplTools.transformDataReverse(item);
        });
        const text = JSON.stringify(data, '', 2);
        if (this.editor) {
          this.editor.setValue(text);
        }
      }, 800);
    },
    // tabbar control
    tabBarLayoutUpdate() {
      const tabBarContainerEl = this.$refs.tabBarRef;
      const tabBarContentEl = this.$refs.tabBarTabsRef?.$el;
      const widthO = tabBarContainerEl?.clientWidth || 0;
      const widthI = tabBarContentEl?.clientWidth || 0;
      if (widthI > 0 && widthI >= widthO - 72) {
        this.tabBar.showNav = true;
        this.navScroll('', { refresh: true });
      } else {
        this.tabBar.showNav = false;
        this.navScroll('', { useMin: true });
      }
    },
    navScroll(type, options = {}) {
      const tabBarContainerEl = this.$refs.tabBarRef.querySelector('.list-group-tabs-c');
      const tabBarContentEl = this.$refs.tabBarTabsRef?.$el;
      const widthO = tabBarContainerEl?.clientWidth || 0;
      const widthI = tabBarContentEl?.clientWidth || 0;
      const max = Math.max(0, Math.abs(widthI - widthO) + 1);
      const transform = tabBarContentEl.style.transform;
      const match = transform.match(/translateX\((.*?)\)/) || [0, 0];
      const translateX = parseInt(match[1]);
      let _next = Math.abs(translateX) + (type == 'left' ? - widthO * 2 / 3 - 1 : widthO * 2 / 3 + 1);
      if (options.refresh) {
        _next = Math.abs(translateX);
      }
      if (options.useMax) {
        _next = max;
      }
      if (options.useMin) {
        _next = 0;
      }
      const next = Math.min(Math.max(0, _next), max);
      tabBarContentEl.style.transform = `translateX(${-next}px)`;
    },
    resize() {
      this.tabBarLayoutUpdate();
    },
    initEditor() {
      return new Promise((resolve) => {
        let count = 0;
        this.editorInitTimer = setInterval(() => {
          count++;
          if (count >= 60) {
            clearInterval(this.editorInitTimer);
            resolve(false); // 初始化失败
          }
          if ($ && $('#edit_area') && $('#edit_area').data('editor')) {
            this.editor = $('#edit_area').data('editor');
            this.editor.updateOptions({
              readOnly: true,
            });
            this.editorReady = true;
            setTimeout(() => {
              window.monaco && window.monaco.editor.setModelLanguage(this.editor.getModel(), 'json');
            }, 500);
            clearInterval(this.editorInitTimer);
            resolve(true); // 初始化成功
          }
        }, 1000);
      });
    },
    check() {
      for (let i = 0, iLen = this.tabs.length; i < iLen; i++) {
        const tab = this.tabs[i];
        if (!tab.name || !tab.taskType || !tab.cluster || !tab.computeResource) {
          this.$message({
            type: 'info',
            message: this.$t('taskTmplObj.completeTmplTips', { n: i + 1 }),
          });
          this.tabsValue = tab.key;
          document.querySelector('html').scrollTo({ top: 0, behavior: 'smooth' });
          document.querySelector('body').scrollTo({ top: 0, behavior: 'smooth' });
          return false;
        }
      }
      return true;
    },
    initSubmitEvent() {
      const self = this;
      document.querySelector('.ui.edit.form').addEventListener('submit', function (e) {
        if (!self.check()) {
          e.preventDefault();
          return false;
        }
        return true;
      });
    },
  },
  beforeMount() {
    const searchParams = new URLSearchParams(window.location.search);
    let tmplIndex = Number(searchParams.get('tmplIndex') || 0);
    let TemplateList = [];
    try {
      TemplateList = JSON.parse(window.TemplateListStr) || [];
      if (!Array.isArray(TemplateList)) {
        TemplateList = [];
      }
    } catch (err) {
      TemplateList = [];
    }
    const isNewFile = location.pathname.indexOf('_new') > 0;
    const isNewTmpl = searchParams.get('new');
    let tmplObj = null;
    try {
      const tmplStr = searchParams.get('tmpl');
      if (tmplStr) {
        tmplObj = JSON.parse(decodeURIComponent(tmplStr));
      }
    } catch (err) {
      console.log(err);
    }
    console.log("tmplObj",tmplObj)
    if ((isNewFile || isNewTmpl) && tmplObj) {
      if (!tmplObj.Name) {
        tmplObj.Name = this.getRandomTmplName();
      }
      TemplateList.push(tmplObj);
      tmplIndex = TemplateList.length - 1;
    }
    if (isNewFile && !tmplObj) {
      setTimeout(() => {
        this.addTmpl();
      }, 1000);
    }
    this.tabs = TemplateList.map(item => {
      const data = taskTmplTools.transformData(item);
      return {
        key: uuidv4(),
        ...data,
      }
    });
    this.tabsValue = this.tabs[tmplIndex]?.key || this.tabs[0]?.key || '';
  },
  async mounted() {
    window.addEventListener('resize', this.resize);
    await this.initEditor();
    this.initSubmitEvent();
    // 确保编辑器就绪后再初始化数据
    this.renderText(); 
  },
  beforeDestroy() {
    window.removeEventListener('resize', this.resize);
  },
};
</script>

<style scoped lang="less">
.flip-list-move {
  transition: transform 0.5s;
}

.no-move {
  transition: transform 0s;
}

.tasktmpl-edit-c {
  background: rgb(240, 240, 240);
  position: relative;
  overflow: hidden;
  border-radius: .28571429rem .28571429rem 0 0;

  .tasktmpl-edit-title-c {
    display: flex;
    align-items: center;
    justify-content: space-between;
    height: 32px;
    padding: 0 14px;
    margin: 8px 0;

    .tasktmpl-edit-title {
      color: rgba(0, 0, 0, 0.9);
      font-size: 18px;
      font-weight: 700;
    }

    .tasktmpl-edit-title-tips {
      display: flex;
      align-items: center;
      color: rgb(130, 130, 130);

      i {
        margin-right: 4px;
        font-size: 16px;
      }

      span {
        color: rgba(0, 0, 0, 0.9);
        font-size: 14px;
      }
    }
  }

  .tasktmpl-tab-bar-c {
    display: flex;
    align-items: center;
    padding: 0 20px;
    height: 38px;

    .nav-left {
      margin-right: 8px;
      cursor: pointer;
    }

    .nav-right {
      margin-left: 8px;
      cursor: pointer;
    }

    .list-group-tabs-c {
      max-width: calc(100% - 72px + 40px);
      overflow: hidden;
    }

    .list-group-tabs {
      position: relative;
      float: left;
      display: block;
      transition: transform 0.5s;

      span {
        display: flex;
        align-items: center;
        justify-content: flex-start;

        .list-group-tabs-item {
          padding: 0 6px 0 8px;
          height: 38px;
          box-sizing: border-box;
          color: rgba(0, 0, 0, 0.9);
          border-radius: 5px 5px 0px 0px;
          border-color: rgb(212, 212, 213);
          border-width: 1px 1px 0px;
          border-style: solid;
          display: flex;
          align-items: center;
          background: rgb(226, 226, 228);
          border-bottom: 1px solid rgba(182, 182, 182, 0.7);

          .name {
            min-width: 80px;
            max-width: 200px;
            overflow: hidden;
            text-overflow: ellipsis;
            white-space: nowrap;
            cursor: pointer;
            text-align: center;
          }

          i {
            margin-left: 8px;
            color: rgb(96, 96, 96);
            cursor: pointer;

            &:hover {
              color: rgb(3, 102, 214);
            }
          }

          &.is-active {
            color: rgb(3, 102, 214);
            border-color: rgba(182, 182, 182, 0.7);
            border-width: 1px 1px 0px;
            border-style: solid;
            background: rgb(249, 249, 249);
          }
        }
      }

      &.has-nav {
        width: calc(100% - 72px);
        overflow: hidden;
      }
    }

    .add-btn {
      display: flex;
      align-items: center;
      justify-content: center;
      cursor: pointer;
      margin-left: 8px;

      i {
        font-size: 20px;
      }

      span {
        margin-top: 2px;
      }
    }
  }

  .list-group-content {
    border-color: rgba(182, 182, 182, 0.7);
    border-width: 1px;
    border-style: solid;
    background: rgb(249, 249, 249);
    margin-top: -1px;
    border-bottom: none;

    .list-group-content-item {
      padding: 20px 30px 12px 30px;
    }
  }
}
</style>
