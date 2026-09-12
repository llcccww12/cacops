<template>
  <div>
    <div class="title" v-if="type == 'add'">请在下方填写您的算力使用需求，部分大额或独占资源使用涉及付费，平台运营人员会与您联系沟通详情。</div>
    <div class="form-c">
      <el-form ref="formRef" class="ignore-dirty" size="default" :model="form" :rules="formRules" label-width="110px"
        @validate="validateEvent">
        <el-form-item label="计算资源" prop="compute_resource">
          <el-select :disabled="disabledEdit" v-model="form.compute_resource" placeholder="请选择计算资源"
            @change="changeComputeResource">
            <el-option v-for="itm in computeResourceList" :key="itm.k" :label="itm.v" :value="itm.k"></el-option>
          </el-select>
        </el-form-item>
        <el-form-item label="卡类型" prop="card_type">
          <el-select :disabled="disabledEdit" v-model="form.card_type" placeholder="请选择卡类型">
            <el-option v-for="itm in cardTypeList" :key="itm.k" :label="itm.v" :value="itm.k"></el-option>
          </el-select>
        </el-form-item>
        <el-form-item label="卡数(卡)" prop="acc_cards_num">
          <el-input :disabled="disabledEdit" ref="accCardsNumRef" v-model="form.acc_cards_num"
            placeholder="请输入卡数，例如数值(4)或者范围(8-16)"></el-input>
        </el-form-item>
        <el-form-item label="存储容量(GB)" prop="disk_capacity">
          <el-input :disabled="disabledEdit" v-model.number="form.disk_capacity" placeholder="请输入存储容量"></el-input>
        </el-form-item>
        <el-form-item label="资源性质" prop="resource_type">
          <el-select :disabled="disabledEdit" v-model="form.resource_type" placeholder="请选择资源性质">
            <el-option v-for="itm in resourceTypeList" :key="itm.k" :label="itm.v" :value="itm.k"></el-option>
          </el-select>
        </el-form-item>
        <el-form-item label="使用时间" prop="use_time">
          <el-date-picker :disabled="disabledEdit" v-model="form.use_time" type="daterange" range-separator="至"
            start-placeholder="开始日期" end-placeholder="结束日期">
          </el-date-picker>
        </el-form-item>
        <el-form-item label="联系人姓名" prop="contact">
          <el-input :disabled="disabledEdit" v-model="form.contact" placeholder="请输入" maxlength="255"></el-input>
        </el-form-item>
        <el-form-item label="微信" prop="wechat">
          <el-input :disabled="disabledEdit" v-model="form.wechat" placeholder="请输入" maxlength="127"></el-input>
        </el-form-item>
        <el-form-item label="电话" prop="phone_number">
          <el-input :disabled="disabledEdit" v-model="form.phone_number" placeholder="请输入" maxlength="255"></el-input>
        </el-form-item>
        <el-form-item label="邮箱" prop="email_address">
          <el-input :disabled="disabledEdit" v-model="form.email_address" placeholder="请输入" maxlength="255"></el-input>
        </el-form-item>
        <el-form-item label="使用组织" prop="org">
          <el-input v-model="form.org" placeholder="请输入平台组织账号，多个用英文分号;分隔" maxlength="255"></el-input>
        </el-form-item>
        <el-form-item label="是否用于科研项目" prop="is_research_project">
          <el-radio v-model="form.is_research_project" :label="false">用于非科研项目</el-radio>
          <el-radio v-model="form.is_research_project" :label="true">用于科研项目</el-radio>
        </el-form-item>
        <el-form-item v-if="form.is_research_project" label="高校或科研机构名称" prop="institution_name">
          <el-input v-model="form.institution_name" placeholder="请输入" maxlength="127"></el-input>
        </el-form-item>
        <el-form-item v-if="form.is_research_project" label="科研项目名称" prop="project_name">
          <el-input v-model="form.project_name" placeholder="请输入" maxlength="127"></el-input>
        </el-form-item>
        <el-form-item v-if="form.is_research_project" label="科研项目编码" prop="project_code">
          <el-input v-model="form.project_code" placeholder="请输入" maxlength="127"></el-input>
        </el-form-item>
        <el-form-item label="使用背景" prop="description">
          <el-input type="textarea" v-model="form.description" placeholder="请输入算力使用需求，如背景介绍、具体使用场景等"
            maxlength="500"></el-input>
        </el-form-item>
        <el-form-item v-if="type == 'edit'" label="提交时间" prop="description">
          <div>{{ timeFormat(data.created_unix) }}</div>
        </el-form-item>
        <el-form-item v-if="type == 'edit'" label="状态" prop="description">
          <div class="status-c">
            <span class="status pending" v-if="data.status == 1"><i
                class="el-icon-stopwatch"></i><span>待处理</span></span>
            <span class="status accepted" v-if="data.status == 2"><i
                class="el-icon-circle-check"></i><span>已接纳</span></span>
            <span class="status refuse" v-if="data.status == 3"><i
                class="el-icon-circle-close"></i><span>未接纳</span></span>
          </div>
        </el-form-item>
        <el-form-item class="button-mobile">
          <el-button type="success" @click="onSubmit">{{ type == 'edit' ? '修改需求' : '提交新需求' }}</el-button>
        </el-form-item>
      </el-form>
    </div>
  </div>
</template>

<script>
import { getDemandCreationRequired, postDemand, updateDemand } from '~/apis/modules/computingpower';
import { formatDate } from 'element-ui/lib/utils/date-util';
import { ACC_CARD_TYPE } from '~/const';
import { getListValueWithKey } from '~/utils';

export default {
  name: "DemandForm",
  props: {
    type: { type: String, default: 'add' }, // add edit
    data: { type: Object, default: () => ({}) },
  },
  data() {
    return {
      form: {
        compute_resource: '',
        card_type: '',
        acc_cards_num: '',
        disk_capacity: '',
        resource_type: '',
        use_time: '',
        contact: '',
        wechat: '',
        phone_number: '',
        email_address: '',
        org: '',
        is_research_project: false,
        institution_name: '',
        project_name: '',
        project_code: '',
        description: '',
      },
      formRules: {
        compute_resource: [{ required: true, message: ' ' }],
        card_type: [{ required: true, message: ' ' }],
        acc_cards_num: [
          { required: true, message: ' ' },
          {
            validator: (rule, value, callback) => {
              const reg1 = /^[1-9]\d*$/;
              const reg2 = /^[1-9]\d*\-[1-9]\d*$/;
              if (reg1.test(value)) {
                callback();
                return;
              }
              if (reg2.test(value)) {
                const nums = value.split('-');
                if (Number(nums[0]) > Number(nums[1])) {
                  callback(new Error('请输入正确的卡数范围'));
                } else {
                  callback();
                }
              }
              callback(new Error('请输入正确的卡数'));
            },
            trigger: 'blur',
          }
        ],
        disk_capacity: [
          { required: true, message: ' ' },
          {
            validator: (rule, value, callback) => {
              if (!Number.isInteger(value)) {
                callback(new Error('请输入数字值'));
              } else {
                if (value <= 0) {
                  callback(new Error('必须大于0GB'));
                } else {
                  callback();
                }
              }
            }
          }
        ],
        resource_type: [{ required: true, message: ' ' }],
        use_time: [{ required: true, message: ' ' }],
        contact: [{ required: true, message: ' ' }],
        wechat: [{ required: true, message: ' ' }],
        phone_number: [
          { required: true, message: ' ' },
          {
            validator: (rule, value, callback) => {
              const reg1 = /^(?:(?:\+|00)86)?1\d{10}$/;
              const reg2 = /^(?:(?:\d{3}-)?\d{8}|^(?:\d{4}-)?\d{7,8})(?:-\d+)?$/;
              if (reg1.test(value) || reg2.test(value)) {
                callback();
              } else {
                callback(new Error('请输入正确的电话'));
              }
            },
            trigger: 'blur',
          }
        ],
        email_address: [
          { required: true, message: ' ' },
          {
            validator: (rule, value, callback) => {
              const reg = /^[A-Za-z0-9_]+@[a-zA-Z0-9_-]+(\.[a-zA-Z0-9_-]+)+$/;
              if (!reg.test(value)) {
                callback(new Error('请输入正确的邮箱'));
              } else {
                callback();
              }
            },
            trigger: 'blur',
          }
        ],
        is_research_project: [{ required: true, message: ' ' }],
        institution_name: [{ required: true, message: ' ' }],
        project_name: [{ required: true, message: ' ' }],
        project_code: [{ required: true, message: ' ' }],
      },
      computeResourceList: [],
      computeResourceCardTypeMap: {
        'GPU': ['A100', 'V100', 'T4', '3090'],
      },
      cardTypeList: [],
      resourceTypeList: [
        { k: 1, v: '共享' },
        { k: 2, v: '独占' },
      ],
      submitLoading: false,
    };
  },
  computed: {
    disabledEdit() {
      return this.type == 'edit' && this.data.status == 2; // 1-待处理，2-已接纳，3-未接纳
    }
  },
  methods: {
    timeFormat(unix) {
      return formatDate(new Date(unix * 1000), 'yyyy-MM-dd HH:mm:ss');
    },
    validateEvent() {
      this.toggleDirtyForm(true);
    },
    resetForm() {
      this.$refs['formRef'].resetFields();
      this.cardTypeList = [];
      this.toggleDirtyForm(false);
    },
    changeComputeResource() {
      this.cardTypeList = this.computeResourceCardTypeMap[this.form.compute_resource].map(itm => ({ k: itm, v: getListValueWithKey(ACC_CARD_TYPE, itm) }));
      this.form.card_type = this.cardTypeList.length ? this.cardTypeList[0].k : '';
    },
    initEdit() {
      this.resetForm();
      for (let key in this.form) {
        this.form[key] = this.data[key];
      }
      this.form.use_time = [new Date(this.data['begin_date']), new Date(this.data['end_date'])];
      this.changeComputeResource();
      this.form.card_type = this.data['card_type'];
    },
    initApply(data) {
      this.form.compute_resource = '';
      this.form.card_type = '';
      this.cardTypeList = [];
      this.form.acc_cards_num = '';
      const computeResource = data.ComputeResource.indexOf('-GPGPU') >= 0 ? 'GPGPU' : data.ComputeResource;
      if (this.computeResourceList.find(item => item.k == computeResource)) {
        this.form.compute_resource = computeResource;
        this.changeComputeResource();
      }
      if (this.cardTypeList.find(item => item.k == data.AccCardType)) {
        this.form.card_type = data.AccCardType;
      }
      this.form.acc_cards_num = data.AccCardsNum.toString();
      this.$refs['accCardsNumRef'].focus();
    },
    onSubmit() {
      if (this.submitLoading) return;
      for (let key in this.form) {
        if (typeof this.form[key] == 'string') {
          this.form[key] = this.form[key].trim();
        }
      }
      this.$refs['formRef'].validate((valid) => {
        if (valid) {
          const submitData = {
            ...this.form,
            begin_date: formatDate(this.form.use_time[0], 'yyyy-MM-dd'),
            end_date: formatDate(this.form.use_time[1], 'yyyy-MM-dd'),
          };
          if (!this.form.is_research_project) {
            submitData.institution_name = '';
            submitData.project_name = '';
            submitData.project_code = '';
          }
          delete submitData['use_time'];
          if (this.type == 'add') {
            this._addDemand(submitData);
          }
          if (this.type == 'edit') {
            submitData.id = this.data.id;
            this._updateDemand(submitData);
          }
        } else {
          console.log('error submit!!');
          return false;
        }
      });
    },
    _addDemand(submitData) {
      this.submitLoading = true;
      postDemand(submitData).then(res => {
        this.submitLoading = false;
        res = res.data;
        if (res.code == 0) {
          this.$message({
            type: 'success',
            message: '需求提交成功！',
          });
          this.resetForm();
          this.$emit('success');
        } else {
          this.$message({
            type: 'error',
            message: '需求提交失败！',
          });
          this.$emit('error');
        }
      }).catch(err => {
        console.log(err);
        this.submitLoading = false;
        this.$message({
          type: 'error',
          message: '需求提交失败！',
        });
        this.$emit('error');
      })
    },
    _updateDemand(submitData) {
      this.submitLoading = true;
      updateDemand(submitData).then(res => {
        this.submitLoading = false;
        res = res.data;
        if (res.code == 0) {
          this.$message({
            type: 'success',
            message: '需求修改成功！',
          });
          this.$emit('success');
        } else {
          this.$message({
            type: 'error',
            message: '需求修改失败！',
          });
          this.$emit('error');
        }
      }).catch(err => {
        console.log(err);
        this.submitLoading = false;
        this.$message({
          type: 'error',
          message: '需求修改失败！',
        });
        this.$emit('error');
      })
    },
    toggleDirtyForm(isDirty) {
      if (isDirty) {
        this.$refs['formRef'].$el.classList.remove('ignore-dirty');
        this.$refs['formRef'].$el.classList.add('dirty');
      } else {
        this.$refs['formRef'].$el.classList.remove('dirty');
        this.$refs['formRef'].$el.classList.add('ignore-dirty');
      }
    },
  },
  mounted() {
    getDemandCreationRequired().then(res => {
      res = res.data;
      if (res.code == 0) {
        const data = res.data;
        Object.assign(this.computeResourceCardTypeMap, { ...data });
        this.computeResourceList = Object.keys(this.computeResourceCardTypeMap).map(itm => ({ k: itm, v: itm }));
        this.cardTypeList = [];
        this.form.card_type = '';
      }
      if (this.type == 'edit') {
        this.$nextTick(() => {
          this.initEdit();
        });
      }
    }).catch(err => {
      console.log(err);
    });
  },
};
</script>
<style scoped lang="less">
@media only screen and (max-width: 767px) {
  /deep/ .el-form-item__label {
    color: #101010 !important;
  }
  .button-mobile {
   /deep/ .el-form-item__content {
      margin: 0 !important;
    }
    /deep/ .el-button {
      width: 100%;
      height: 45px;
    }
  }
}
.title {
  margin-bottom: 16px;
  color: #606266;
}

.form-c {

  .el-select,
  .el-range-editor {
    width: 100%;

    /deep/.el-range-separator {
      min-width: 20px;
    }
  }

  /deep/ .el-form-item__label {
    line-height: 20px;
    height: 40px;
    display: flex;
    align-items: center;
    justify-content: flex-end;
  }
  @media screen and (max-width: 768px) {
    /deep/ .el-input__inner::placeholder, /deep/ .el-range-input::placeholder, /deep/ .el-textarea__inner::placeholder {
      font-size: 12px;
    }
  }
}

.status-c {
  .status {
    margin-left: 8px;

    i {
      margin-right: 3px;
    }

    &.accepted {
      color: rgb(39, 177, 72);
    }

    &.refuse {
      color: rgb(140, 162, 170);
    }

    &.pending {
      color: rgb(50, 145, 248);
    }
  }
}
</style>
