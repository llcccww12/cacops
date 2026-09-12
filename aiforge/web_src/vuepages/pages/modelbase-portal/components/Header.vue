<template>
  <div>
    <div class="main-header">
      <div class="area-l">
        <slot name="left">
          <div class="area-l-title" :style="labelStyle">
            <template v-for="(item, index) in breadcrumbData">
              <div v-if="index != 0" class="separator"> / </div>
              <div v-if="index != breadcrumbData.length - 1" class="link label" @click="navHandler(item)">{{
                $t(item.label) }}
              </div>
              <div v-else class="label">{{ $t(item.label) }}</div>

            </template>
            <a v-if="lastPath && lastPath.path" href="javascript:;" class="ui poping up clipboard" id="clipboard-btn"
              data-position="top center" data-variation="inverted tiny" :data-success="$t('copySuccess')"
              :data-content="$t('copy')" :data-original="$t('copy')" :data-clipboard-text="lastPath.path">
              <i class="ri-file-copy-line"></i>
            </a>
            <span class="task-status">
              <i :class="mainData[0] && mainData[0].task.status"></i>
              <span>{{ mainData[0] && mainData[0].task.status }}</span>
              <i v-if="mainData[0] && mainData[0].task.detailed_status === 'dataMigrating' && mainData[0].task.status === 'WAITING'"
                :class="mainData[0].task.detailed_status" :title="$t('cloudbrainObj.migratingData')"></i>
              <i v-if="mainData[0] && mainData[0].task.detailed_status === 'centerPending' && mainData[0].task.status === 'WAITING'"
                :class="mainData[0].task.detailed_status" :title="$t('cloudbrainObj.centerPending')"></i>
              <i v-if="mainData[0] && mainData[0].task.detailed_status === 'ImagePulling' && mainData[0].task.status === 'WAITING'"
                :class="mainData[0].task.detailed_status" :title="$t('cloudbrainObj.imagePulling')"></i>
            </span>
          </div>

        </slot>
      </div>
      <div class="area-r">
        <slot name="right"></slot>
      </div>
    </div>
    <div class="descr" v-if="descr" v-html="descr"></div>
  </div>
</template>

<script>
export default {
  name: 'Header',
  props: {
    useBreadcrumb: { type: Boolean, default: true },
    lastPath: { type: Object, default: null },
    mainData: { type: Array, default: () => { return [] } }
  },
  data() {
    return {
      breadcrumbData: [],
      descr: '',
    };
  },
  components: {},
  watch: {
    lastPath: {
      handler() {
        this.refresh();
      }
    }
  },
  computed: {
    labelStyle() {
      return {
        fontSize: this.$isMobile ? '22px' : '28px', // 动态设置字体大小
      };
    },
  },
  methods: {
    navHandler(item) {
      this.$router.replace(item.path);
    },
    refresh() {
      this.breadcrumbData = [];
      const routes = this.$router.options.routes || [];
      const pathMap = {};
      const walk = (item, parentPath) => {
        const fullPath = `${parentPath}/${item.path}`;
        const cur = {
          ...item,
          path: fullPath,
        };
        if (cur && item.children && item.children.length) {
          item.children.forEach(_item => {
            walk(_item, cur.path);
          });
        }
        pathMap[cur.path] = cur;
      };
      (routes[0]?.children || []).forEach(item => {
        walk(item, '');
      });
      const pathList = this.$route.path.slice(1).split('/');
      let path = '';
      for (let i = 0, iLen = pathList.length; i < iLen; i++) {
        path += ('/' + pathList[i]);
        const routeI = pathMap[path];
        if (routeI && !routeI.redirect) {
          this.breadcrumbData.push({
            label: routeI.meta?.label,
            path: routeI.path,
          });
        }
        if (i == iLen - 1) {
          this.descr = this.$t(routeI?.meta?.descr);
        }
      }
      if (this.lastPath) {
        this.breadcrumbData.push({
          label: this.lastPath.label,
          path: this.lastPath.path,
        });
      }
    }
  },
  created() {
    this.refresh();
  },
  mounted() { },
};
</script>

<style scoped lang="less">
.main-header {
  display: flex;
  align-items: center;
  justify-content: space-between;
  padding: 16px 20px 0 20px;

  .area-l {

    .area-l-title {
      display: flex;
      align-items: center;
      flex-wrap: wrap;
      font-size: 28px;

      .separator {
        color: rgba(111, 118, 165, 1);
        font-family: SourceHanSansSC;
        margin-right: 10px;
      }

      .label {
        line-height: 39px;
        margin-right: 10px;
        font-family: SourceHanSansSC;
        font-weight: 500;

        &.link {
          color: rgba(3, 102, 214, 1);
          font-weight: 400;
          cursor: pointer;

          &:hover {
            color: rgba(3, 102, 214, .8);
          }
        }
      }

      i {
        color: rgba(145, 145, 145, 1);
        font-size: 16px;
        margin-left: 8px;
      }

      .task-status {
        display: flex;
        align-items: center;
        margin-right: 12px;
        margin-left: 10px;

        i,
        span {
          margin-right: 4px;
          font-size: 14px;
        }
      }
    }


    .special-label {
      line-height: 39px;
      margin-right: 10px;
      font-family: SourceHanSansSC;
      font-size: 28px;
      font-weight: 500;
    }
  }

  .real-r {
    margin-left: 20px;
  }
}

.descr {
  color: rgba(80, 85, 89, 1);
  font-size: 14px;
  font-family: SourceHanSansSC;
  font-weight: 400;
  line-height: 20px;
  margin-top: 8px;
  padding-left: 20px;
}

/* 手机端样式 */
@media (max-width: 768px) {
  .area-l {

    // display: none;
    .area-l-title {
      .lable {
        font-size: 22px !important;
      }
    }

    .special-label {
      display: none;
    }

  }

}
</style>
