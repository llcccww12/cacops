<template>
  <div>
    <div class="form-row">
      <div class="title required">
        <span>{{ $t('resourcesManagement.resourceSpecification') }}</span>
      </div>
      <div class="content"></div>
      <div class="select-btn">
        <a @click="dlgShow = true" href="javascript:;">{{ $t('resourcesManagement.selectSpec') }}</a>
      </div>
      <div class="base-dlg">
        <BaseDialog :visible.sync="dlgShow" :width="`1200px`" :title="$t('resourcesManagement.selectSpec')" @open="open"
          @opened="opened" @close="close" :appendToBody="true" @closed="closed">
          <div class="dlg-content">
            <div class="left-area">
              <div class="tabs">
                <div class="tab" :class="tabIndex == '1' ? 'focused' : ''" @click="changeTab(1)">{{
                  $t('resourcesManagement.accordingSpec') }}</div>
                <div class="tab" :class="tabIndex == '2' ? 'focused' : ''" @click="changeTab(2)">{{
                  $t('resourcesManagement.accordingQueue') }}</div>
              </div>
              <div class="table-c">
                <div v-if="tabIndex == '1'">
                  <div>
                    <div class="header row">
                      <div class="row-l" style="width:400px;">{{ $t('resourcesManagement.resourceSpecification') }}
                      </div>
                      <div class="row-r" style="flex:1">{{ $t('resourcesManagement.resQueue') }}</div>
                    </div>
                    <div class="table-content">
                      <div class="row" v-for="(item, index) in tableData1" :key="index">
                        <div class="row-l" style="width:400px;">{{ item.SpecStr }}</div>
                        <div class="row-r" style="flex:1">
                          <div class="row-item" v-for="_item in item.queues" :key="_item.ID">
                            <el-checkbox :value="_item.checked" @change="selectChange(_item.ID)">
                              <span
                                v-html="_item.QueueStr + _item.PriceStr + _item.NetworkTypeStr + _item.VisualizationStr + _item.StatusStr + _item.QueueIsExclusiveStr"></span>
                            </el-checkbox>
                          </div>
                          <div class="btn-c">
                            <button @click="selectAll(item.queues)">{{ $t('selectAll') }}</button>
                            <button @click="clearSelectAll(item.queues)">{{ $t('selectNone') }}</button>
                          </div>
                        </div>
                      </div>
                    </div>
                  </div>
                </div>
                <div v-if="tabIndex == '2'">
                  <div class="header row">
                    <div class="row-l" style="width:400px;">{{ $t('resourcesManagement.resQueue') }}</div>
                    <div class="row-r" style="flex:1">{{ $t('resourcesManagement.resourceSpecification') }}</div>
                  </div>
                  <div class="table-content">
                    <div class="row" v-for="(item, index) in tableData2" :key="index">
                      <div class="row-l" style="width:400px;">{{ item.QueueStr }}</div>
                      <div class="row-r" style="flex:1">
                        <div class="row-item" v-for="_item in item.specs" :key="_item.ID">
                          <el-checkbox :value="_item.checked" @change="selectChange(_item.ID)">
                            <span
                              v-html="_item.SpecStr + _item.PriceStr + _item.NetworkTypeStr + _item.VisualizationStr + _item.StatusStr + _item.QueueIsExclusiveStr"></span>
                          </el-checkbox>
                        </div>
                        <div class="btn-c">
                          <button @click="selectAll(item.specs)">{{ $t('selectAll') }}</button>
                          <button @click="clearSelectAll(item.specs)">{{ $t('selectNone') }}</button>
                        </div>
                      </div>
                    </div>
                  </div>
                </div>
              </div>
              <div class="confirm-btn-c">
                <el-button type="primary" class="btn confirm-btn" @click="confirm">{{ $t('confirm') }}</el-button>
              </div>
            </div>
          </div>
        </BaseDialog>
      </div>
    </div>
    <div class="row-detail">
      <div class="table-c">
        <div>
          <div class="header row">
            <div class="row-l" style="width:400px;">{{ $t('resourcesManagement.resourceSpecification') }}</div>
            <div class="row-r" style="flex:1">{{ $t('resourcesManagement.resQueue') }}</div>
          </div>
          <div class="table-content">
            <div class="row" v-for="(item, index) in tableDataShow" :key="index">
              <div class="row-l" style="width:400px;">{{ item.SpecStr }}</div>
              <div class="row-r" style="flex:1">
                <div class="" v-for="(_item, _index) in item.queues" :key="_index">
                  <span
                    v-html="_item.QueueStr + _item.PriceStr + _item.NetworkTypeStr + _item.VisualizationStr + _item.StatusStr + _item.QueueIsExclusiveStr"></span>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import BaseDialog from '~/components/BaseDialog.vue';
import { SPECIFICATION_STATUS } from '~/const';

export default {
  name: "SpecSelect",
  props: {
    visible: { type: Boolean, default: false },
    value: { type: Array, required: true },
    specs: { type: Array, required: true },
  },
  components: { BaseDialog },
  data() {
    return {
      dlgShow: false,
      tabIndex: '1',
      selectList: [],
      statusList: [...SPECIFICATION_STATUS],
    };
  },
  watch: {
    visible: function (val) {
      this.dlgShow = val;
    },
    value: {
      immediate: true,
      deep: true,
      handler(newVal) {
        newVal = newVal === undefined ? [] : newVal;
        this.selectList = [...newVal];
      }
    }
  },
  computed: {
    tableData1: function () {
      const map = {};
      for (let i = 0, iLen = this.specs.length; i < iLen; i++) {
        const spec = this.specs[i];
        const key = `${spec.Cluster}|${spec.ComputeResource}|${spec.SourceSpecId || Math.random()}`;
        const queue = {
          ...spec,
          QueueId: spec.QueueId,
          QueueCode: spec.QueueCode,
          AiCenterCode: spec.AiCenterCode,
          AiCenterName: spec.AiCenterName,
          QueueStr: spec.QueueStr,
          QueueIsExclusiveStr: spec.QueueIsExclusiveStr,
          PriceStr: spec.PriceStr,
          NetworkTypeStr: spec.NetworkTypeStr,
          VisualizationStr: spec.VisualizationStr,
          checked: this.selectList.indexOf(spec.ID) > -1,
        };
        if (map[key]) {
          map[key].queues.push(queue);
        } else {
          map[key] = {
            ...spec,
            queues: [queue],
          }
        }
      }
      const data = [];
      for (let key in map) {
        data.push(map[key]);
      }
      return data;
    },
    tableData2: function () {
      const map = {};
      for (let i = 0, iLen = this.specs.length; i < iLen; i++) {
        const spec = this.specs[i];
        const key = `${spec.Cluster}|${spec.ComputeResource}|${spec.QueueId}`;
        const _spec = {
          ...spec,
          QueueId: spec.QueueId,
          QueueCode: spec.QueueCode,
          AiCenterCode: spec.AiCenterCode,
          AiCenterName: spec.AiCenterName,
          QueueStr: spec.QueueStr,
          QueueIsExclusiveStr: spec.QueueIsExclusiveStr,
          PriceStr: spec.PriceStr,
          NetworkTypeStr: spec.NetworkTypeStr,
          VisualizationStr: spec.VisualizationStr,
          checked: this.selectList.indexOf(spec.ID) > -1,
        };
        if (map[key]) {
          map[key].specs.push(_spec);
        } else {
          map[key] = {
            ...spec,
            specs: [_spec],
          }
        }
      }
      const data = [];
      for (let key in map) {
        data.push(map[key]);
      }
      return data;
    },
    tableDataShow: function () {
      const map = {};
      for (let i = 0, iLen = this.specs.length; i < iLen; i++) {
        const spec = this.specs[i];
        if (this.value.indexOf(spec.ID) < 0) continue;
        const key = `${spec.Cluster}|${spec.ComputeResource}|${spec.SourceSpecId || Math.random()}`;
        const queue = {
          ...spec,
          QueueId: spec.QueueId,
          QueueCode: spec.QueueCode,
          AiCenterCode: spec.AiCenterCode,
          AiCenterName: spec.AiCenterName,
          QueueStr: spec.QueueStr,
          QueueIsExclusiveStr: spec.QueueIsExclusiveStr,
          PriceStr: spec.PriceStr,
          NetworkTypeStr: spec.NetworkTypeStr,
          VisualizationStr: spec.VisualizationStr,
        };
        if (map[key]) {
          map[key].queues.push(queue);
        } else {
          map[key] = {
            ...spec,
            queues: [queue],
          }
        }
      }
      const data = [];
      for (let key in map) {
        data.push(map[key]);
      }
      return data;
    },
  },
  methods: {
    resetDataInfo() {
      this.dataInfo = {}
    },
    changeTab(tabIndex) {
      if (this.tabIndex == tabIndex) return;
      this.tabIndex = tabIndex;
    },
    selectChange(id) {
      const index = this.selectList.indexOf(id);
      if (index < 0) {
        this.selectList.push(id);
      } else {
        this.selectList.splice(index, 1);
      }
    },
    selectAll(list) {
      list.forEach(_item => {
        if (this.selectList.indexOf(_item.ID) < 0) {
          this.selectList.push(_item.ID);
        }
      });
    },
    clearSelectAll(list) {
      list.forEach(_item => {
        const index = this.selectList.indexOf(_item.ID);
        if (index >= 0) {
          this.selectList.splice(index, 1);
        }
      });
    },
    open() {
      this.$emit("open");
      this.selectList = [...this.value];
      this.tabIndex = '1';
    },
    opened() {
      this.$emit("opened");
    },
    close() {
      this.$emit("close");
    },
    closed() {
      this.$emit("closed");
      this.$emit("update:visible", false);
    },
    confirm() {
      this.dlgShow = false;
      this.$emit('input', this.selectList);
      this.$emit('change', this.selectList);
    },
    cancel() {
      this.dlgShow = false;
      this.$emit("update:visible", false);
    }
  },
  mounted() {
    this.resetDataInfo();
  },
};
</script>
<style scoped lang="less">
.form-row {
  display: flex;
  min-height: 42px;
  margin-bottom: 4px;

  .title {
    width: 255px;
    display: flex;
    justify-content: flex-end;
    align-items: center;
    margin-right: 20px;
    color: rgb(136, 136, 136);
    font-size: 14px;

    &.required {
      span {
        position: relative;
      }

      span::after {
        position: absolute;
        right: -10px;
        top: -2px;
        vertical-align: top;
        content: '*';
        color: #db2828;
      }
    }
  }

  .content {
    width: 10px;
    display: flex;
    align-items: center;
  }

  .select-btn {
    display: flex;
    align-items: center;
    font-size: 13px;
  }
}

.row-detail {
  font-size: 12px;

  .table-c {
    padding-bottom: 2px;
    padding-right: 8px;

    .table-content {
      max-height: 300px;
      overflow: auto;
      margin-top: -1px;
      padding-bottom: 2px;
    }

    .row {
      display: flex;

      &.header {
        background-color: #f5f5f6;

        .row-l {
          justify-content: center;
          border-bottom: 1px solid gainsboro;
        }

        .row-r {
          flex-direction: row;
          justify-content: center;
          border-bottom: 1px solid gainsboro;
        }
      }

      .row-l {
        display: flex;
        align-items: center;
        border: 1px solid gainsboro;
        border-right: none;
        border-bottom: none;
        padding: 4px 8px;
        width: 0;
        overflow: hidden;
      }

      .row-r {
        border: 1px solid gainsboro;
        padding: 4px 8px;
        display: flex;
        flex-direction: column;
        border-bottom: none;
        width: 0;
        overflow: hidden;
        position: relative;
      }

      &:last-child {
        .row-l {
          border-bottom: 1px solid gainsboro;
        }

        .row-r {
          border-bottom: 1px solid gainsboro;
        }
      }
    }
  }

}

.dlg-content {
  margin: 0px 0 0 0;
  display: flex;
  font-size: 12px;

  .left-area {
    flex: 1;
    width: 0;

    .tabs {
      display: flex;
      align-items: center;
      margin-bottom: 10px;

      .tab {
        display: flex;
        align-items: center;
        justify-content: center;
        font-size: 12px;
        color: rgba(0, 0, 0, .87);
        border: 1px solid rgba(34, 36, 38, .15);
        margin-left: -1px;
        height: 30px;
        padding: 0 6px;
        border-left: none;

        &.focused {
          color: #0087f5;
          border-color: #0087f5;
          border-left: 1px solid #0087f5 !important;
        }

        &:first-child {
          border-top-left-radius: 0.28571429rem;
          border-bottom-left-radius: 0.28571429rem;
          border-left: 1px solid rgba(34, 36, 38, .15);

          &.focus {
            border-color: #0087f5;
          }
        }

        &:last-child {
          border-top-right-radius: 0.28571429rem;
          border-bottom-right-radius: 0.28571429rem;
        }

        &:hover:not(.focused) {
          background: rgba(0, 0, 0, .03);
          cursor: pointer;
        }
      }
    }

    .table-c {
      padding-bottom: 2px;
      padding-right: 8px;

      .table-content {
        height: 400px;
        overflow: auto;
        margin-top: -1px;
        padding-bottom: 2px;
      }

      .row {
        display: flex;

        &.header {
          background-color: #f5f5f6;

          .row-l {
            justify-content: center;
            border-bottom: 1px solid gainsboro;
          }

          .row-r {
            flex-direction: row;
            justify-content: center;
            border-bottom: 1px solid gainsboro;
          }
        }

        .row-l {
          display: flex;
          align-items: center;
          border: 1px solid gainsboro;
          border-right: none;
          border-bottom: none;
          padding: 4px 8px;
          width: 0;
          overflow: hidden;
        }

        .row-r {
          border: 1px solid gainsboro;
          padding: 4px 8px;
          display: flex;
          flex-direction: column;
          border-bottom: none;
          width: 0;
          overflow: hidden;
          position: relative;

          /deep/ .el-checkbox__label {
            font-size: 12px;
          }

          .btn-c {
            position: absolute;
            top: 5px;
            right: 5px;
            z-index: 1;

            button {
              background-color: #409EFF;
              color: #FFF;
              margin-left: 3px;
              outline: none;
              cursor: pointer;
              border: none;
              border-radius: 2px;
              padding: 2px 4px;

              &:focus,
              &:hover {
                background: #66b1ff;
              }

              &:active {
                background: #3a8ee6;
              }
            }
          }

          .row-item {
            &:first-child {
              margin-right: 80px;
            }

            /deep/ .el-checkbox {
              white-space: break-spaces;
              display: inline-flex;
              align-items: center;
            }
          }
        }

        &:last-child {
          .row-l {
            border-bottom: 1px solid gainsboro;
          }

          .row-r {
            border-bottom: 1px solid gainsboro;
          }
        }
      }
    }

    .confirm-btn-c {
      margin-top: 15px;
      text-align: right;
    }
  }
}
</style>
