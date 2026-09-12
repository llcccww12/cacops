<template>
  <div class="form-row">
      <div class="content ">
        <div class="task-items-c" :class="disabled ? 'disabled' : ''">
          <div class="task-item" :class="item.k == task ? 'focus' : ''" v-for="(item, index) in tasks" :key="index"
            @click="changeType(item)">
            <div class="task-item-name">{{ item.name }}</div>
            <div class="task-item-desc nowrap" :title="item.desc">{{ item.desc }}</div>
          </div>
        </div>
        <div class="tips task-item-desc-long">
          <span>{{ descLong }}</span>
          <span>{{ $t('cloudbrainObj.runningLimit', { count: limitCount, taskType: taskName}) }} </span>
        </div>
      </div>
  </div>
</template>

<script>
import { i18n } from '~/langs';

const SAMPLE_REPO = {
  'DEBUG': 'OpenIOSSG/OpenI_Cloudbrain_Example',
  'TRAIN': 'OpenIOSSG/OpenI_Cloudbrain_Example',
  'ONLINEINFERENCE': 'OpenIOSSG/Online-Inference_Example',
  'GENERAL': 'OpenIOSSG/OpenI_Cloudbrain_Example',
  'HPC': 'OpenIOSSG/OpenI_Cloudbrain_Example',
};

export default {
  name: 'TaskTypeSelect',
  props: {
    value: { type: String, required: true },
    limitCount: {type: Number, required: true},
    disabled: { type: Boolean, default: false },
  },
  data() {
    return {
      tasks: [{
        k: 'DEBUG',
        name: i18n.t('cloudbrainObj.tabTitDebug'),
        cluster: 'C2Net',
        computerResouce: 'NPU',
        desc: i18n.t('cloudbrainObj.debugTaskDesc'),
        descLong: i18n.t('cloudbrainObj.debugTaskDescLong'),
      }, {
        k: 'TRAIN',
        name: i18n.t('cloudbrainObj.tabTitTrain'),
        cluster: 'C2Net',
        computerResouce: 'NPU',
        desc: i18n.t('cloudbrainObj.trainTaskDesc'),
        descLong: i18n.t('cloudbrainObj.trainTaskDescLong'),
      }, {
        k: 'ONLINEINFERENCE',
        name: i18n.t('cloudbrainObj.tabTitOnlineInference'),
        cluster: 'C2Net',
        computerResouce: 'GPU',
        desc: i18n.t('cloudbrainObj.onlineinferTaskDesc'),
        descLong: i18n.t('cloudbrainObj.onlineinferTaskDescLong'),
      }, {
        k: 'GENERAL',
        name: i18n.t('cloudbrainObj.tabTitGeneral'),
        cluster: 'C2Net',
        computerResouce: 'GPU',
        desc: i18n.t('cloudbrainObj.generalTaskDesc'),
        descLong: i18n.t('cloudbrainObj.generalTaskDescLong'),
      }, {
        k: 'HPC',
        name: i18n.t('superComputeTask'),
        cluster: 'C2Net',
        computerResouce: 'CPU',
        desc: i18n.t('cloudbrainObj.inferenceTaskDesc'),
        descLong: i18n.t('cloudbrainObj.inferenceTaskDescLong'),
      },],
      task: '',
      emitChangeTimeout: null,
      isEmitting: false,
    };
  },
  computed: {
    descLong() {
      const find = this.tasks.filter(item => item.k == this.task)[0];
      return (find ? find.descLong : '');
    },
    taskName() {
      const find = this.tasks.filter(item => item.k == this.task)[0];
      return (find ? find.name : '');
    }
  },
  watch: {
    value: {
      handler(newVal,oldVal) {
        if (this.isEmitting) return; // 防止循环调用
        console.log("value======================",newVal,oldVal)
        newVal = newVal === undefined ? '' : newVal;
        if (this.tasks.filter(item => item.k == newVal).length) {
          this.task = newVal.toString();
          this.debouncedEmitChange();
        }
      }
    }
  },
  methods: {
    changeType(task) {
      console.log("changeType",task)
      this.task = task.k;
      if (this.disabled) return;
      this.debouncedEmitChange()
    },
    emitChange() {
      this.isEmitting = true;
      const findTask = this.tasks.filter(item => item.k == this.task)[0];
      this.$emit('input', this.task, findTask);
      this.$emit('change', this.task, findTask);
    },
    debouncedEmitChange() {
      clearTimeout(this.emitChangeTimeout);
      this.emitChangeTimeout = setTimeout(() => {
        this.emitChange();
      }, 50); // 50ms 延迟，确保多次变更合并为一次
    }
  },
  beforeDestroy() {
    // 清理定时器
    clearTimeout(this.emitChangeTimeout);
  },
  beforeMount() { },
  mounted() { },
};
</script>

<style scoped lang="less">
.form-row {
  .content {
      .task-items-c {
        display: flex;
        align-items: center;
        flex-wrap: wrap;

        .task-item {
          display: flex;
          align-items: center;
          justify-content: center;
          font-size: 12px;
          border-right: 1px solid rgb(194, 199, 204);
          height: 60px;
          width: 160px;
          padding: 0 12px;
          margin-bottom: 10px;
          flex-direction: column;
          cursor: pointer;
          background-color: rgba(0,178,255,0.1);
          .task-item-name {
            margin-bottom: 2px;
            font-weight: 700;
            font-size: 14px;
            color: rgba(64, 64, 64, 1);
          }

          .task-item-desc {
            width: 100%;
            color: rgba(136, 136, 136, 1);
            text-align: center;
          }
          &:last-child {
            border-right: none;
          }
          &.focus {
            background-color: rgba(255,255,255,1);
            border: 1px solid rgb(50, 145, 248);
            border-radius: 5px;
            margin-left: -4px;
            .task-item-name {
              color: rgba(0, 102, 255, 1);
            }

            .task-item-desc {
              color: rgba(50, 145, 248, 1);
            }
          }
          &:hover:not(.focus) {
            background-color: rgba(255,255,255,1);
          }

          


        }

        &.disabled {
          .task-item {
            cursor: not-allowed;
          }
        }
      }

      .task-item-desc-long {
        color: rgb(250, 140, 22);
        font-size: 13px;
        margin-top: 0;
      }
    }
}
</style>
