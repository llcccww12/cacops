<template>
  <!-- 通用删除确认弹框 -->
  <div :id="modalId" class="ui small modal">
    <div class="header">
      {{ $t(titleKey) }}
    </div>
    <div class="content" v-loading="delLoading">
      <div class="ui warning message text left">
        <p v-html="$t('datasetObj.delete_notices_1')"></p>
        <p v-html="$t(notice2Key, notice2Params)"></p>
      </div>
      <form class="ui form" @submit.prevent>
        <div class="field">
          <label>
            {{ $t(confirmNameKey) }}
            <span class="text red">{{ dataObj.alias }}</span>
          </label>
        </div>
        <div class="required field">
          <label :for="inputId">{{ $t(nameLabelKey) }}</label>
          <input 
            :id="inputId" 
            v-model="name" 
            :name="inputName" 
            required
          >
        </div>

        <div class="text right actions">
          <div class="ui cancel button" @click="handleCancel">{{ $t('cancel') }}</div>
          <button 
            class="ui red button" 
            @click="handleDelete"
            :disabled="!name || delLoading"
          >
            {{ $t('confirm1') }}
          </button>
        </div>
      </form>
    </div>
  </div>
</template>

<script>
export default {
  name: 'DeleteConfirmModal',
  props: {
    // 组件配置
    type: {
      type: String,
      required: true,
    },
    modalId: {
      type: String,
      default: 'delete-confirm-modal'
    },
    
    // 数据对象
    dataObj: {
      type: Object,
      required: true
    },
    
    // 自定义国际化key（可选）
    customKeys: {
      type: Object,
      default: () => ({})
    }
  },
  
  data() {
    return {
      name: '',
      delLoading: false
    }
  },
  
  computed: {
    // i18n key prefix映射：type值到i18n key前缀的转换
    iPrefix() {
      const map = { model: 'modelObj', dataset: 'datasetObj', repos: 'repos' }
      return map[this.type] || this.type
    },

    // 获取用于生成key后缀的基础类型（首字母大写）
    baseTypeCap() {
      return this.capitalizeFirstLetter(this.getBaseType(this.type))
    },

    // 根据类型计算国际化key
    titleKey() {
      return this.customKeys.titleKey || `${this.iPrefix}.deleteThis${this.baseTypeCap}`
    },
    notice2Key() {
      return this.customKeys.notice2Key || `${this.iPrefix}.delete_notices_2`
    },

    notice2Params() {
      return { [this.getBaseType(this.type)]: this.dataObj.alias }
    },

    confirmNameKey() {
      return this.customKeys.confirmNameKey || `${this.iPrefix}.sure${this.baseTypeCap}Name`
    },

    nameLabelKey() {
      return this.customKeys.nameLabelKey || `${this.iPrefix}.${this.getBaseType(this.type)}_name1`
    },

    nameErrorKey() {
      return this.customKeys.nameErrorKey || `${this.iPrefix}.delete${this.baseTypeCap}NameError`
    },
    
    inputId() {
      return `${this.type}_name`
    },
    
    inputName() {
      return `${this.type}_name`
    }
  },
  
  methods: {
    // 获取基础类型：modelObj->model, datasetObj->dataset, repos->repos
    getBaseType(type) {
      if (type.endsWith('Obj')) {
        return type.slice(0, -3)
      }
      return type
    },

    capitalizeFirstLetter(string) {
      return string.charAt(0).toUpperCase() + string.slice(1)
    },
    
    showModal() {
      this.name = ''
      $(`#${this.modalId}`)
        .modal({
          onDeny: () => {
            // this.$emit('cancel')
          },
          onApprove: () => {
            // 不在这里处理，由按钮事件处理
          },
          onHidden: () => {
            // this.$emit('modal-hidden')
          }
        })
        .modal('show')
    },
    
    hideModal() {
      this.resetForm()
      $(`#${this.modalId}`).modal('hide')
    },
    
    handleCancel() {
      this.resetForm()
      this.hideModal()
      this.$emit('cancel')
    },
    resetForm() {
      // 重置表单状态
      this.name = ''
      // 如果使用jQuery的form，可以这样重置
      $(`#${this.modalId} form`)[0]?.reset()
    },
    async handleDelete() {
      // 验证输入的名称
      if (this.name !== this.dataObj.alias) {
        this.$message.error(this.$t(this.nameErrorKey))
        return
      }
      
      this.delLoading = true
      
      try {
        // 发射删除事件，由父组件处理实际的删除逻辑
        this.$emit('confirm-delete', {
          type: this.type,
          data: this.dataObj,
          callback: this.handleDeleteCallback
        })
      } catch (error) {
        this.delLoading = false
        this.$emit('delete-error', error)
      }
    },
    
    handleDeleteCallback(result) {
      this.delLoading = false
      
      if (result.success) {
        this.$message.success(result.message || this.$t('imagesObj.deleteSuccessTips'))
        this.hideModal()
        this.name = ''
        this.$emit('delete-success', this.dataObj)
      } else {
        this.$message.error(result.error || this.$t('common.deleteFailed'))
      }
    }
  }
}
</script>