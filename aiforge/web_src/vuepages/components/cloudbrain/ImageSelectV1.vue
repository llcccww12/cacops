<template>
  <div class="form-row">
    <div class="left-area">
      <div class="title" v-if="showTitle">
        <span :class="required ? 'required' : ''">{{ $t('cloudbrainObj.image') }}</span>
      </div>
      <div class="content" :class="errStatus ? 'error' : ''">
        <el-input v-if="useId" class="field-input" v-model="image.image_name" :readonly="true"
          :placeholder="$t('cloudbrainObj.selectImage')"></el-input>
        <el-input v-else class="field-input" v-model="image.image_name" @input="imageChange"
          :readonly="configs.computerResouce != 'GPU'" :placeholder="configs.computerResouce == 'GPU' ?
            $t('cloudbrainObj.selectImagePlaceholder') : $t('cloudbrainObj.selectImage')"></el-input>
        <div class="tips" v-if="errStatus && showInnerUrlTip" style="color:#e0b4b4">
          {{ $t('cloudbrainObj.imageInnerUrlErrTips') }}
        </div>
        <div class="tips" v-if="errStatus && showUrlErrTip" style="color:#e0b4b4">
          {{ $t('cloudbrainObj.imageUrlErrTips') }}
        </div>
        <div class="dynamic-tips" v-if="tips" v-html="tips"></div>
      </div>
    </div>
    <div class="right-area">
      <div class="btn-select" @click="dlgShow = true">
        <i class="el-icon-plus"></i>
        <span>{{ $t('cloudbrainObj.selectImage') }}</span>
      </div>
    </div>
    <el-dialog class="model-dlg" :visible.sync="dlgShow" width="1000px"
      :modal="true" :close-on-click-modal="false" :show-close="true" :destroy-on-close="false"
      :before-close="beforeClose" @open="open" @closed="closed">
      <div class="dlg-content">
        <div class="header-c">
          <span class="header-text">{{$t('cloudbrainObj.selectImage')}}</span>
          <el-input size="small" class="search-inp" :placeholder="$t('cloudbrainObj.searchImagePlaceholder')"
            v-model="dlgSearchValue" @keydown.enter.stop.native.prevent="search">
            <div slot="suffix" class="search-inp-icon" @click="search">
              <i class="el-icon-search"></i>
            </div>
          </el-input>
        </div>
        <div class="main-area" v-loading="dlgLoading">
          <div class="image-tabs-c">
            <el-tabs class="image-tabs" v-model="dlgActiveName" @tab-click="dlgTabClick">
              <el-tab-pane :label="$t('cloudbrainObj.recommendImage')" name="first"></el-tab-pane>
              <el-tab-pane :label="$t('cloudbrainObj.myImage')" name="second"></el-tab-pane>
              <el-tab-pane :label="$t('cloudbrainObj.myFavImage')" name="third"></el-tab-pane>
            </el-tabs>
          </div>
          <div class="filter-c">
            <div class="cascader-c">
              <span class="cascader-tit">{{ $t('imagesObj.filterImages') }}：</span>
              <div class="cascader-content-c">
                <div class="cascader-tips">{{ filterImagesPlaceholder }}</div>
                <el-cascader class="image-filter" ref="cascaderFilterRef" v-model="dlgCascaderFilter.value"
                  :props="dlgCascaderProps" :options="dlgCascaderFilter.options" clearable
                  :placeholder="$t('imagesObj.filterImages')" :popper-class="dlgCascaderFilter.popperClass"
                  @expand-change="handleDlgCascaderFilterExpandChange"
                  @visible-change="handleDlgCascaderFilterVisibleChange"
                  @change="handleDlgCascaderFilterChange"></el-cascader>
              </div>
            </div>
            <!-- <el-input size="small" class="search-inp" :placeholder="$t('cloudbrainObj.searchImagePlaceholder')"
              v-model="dlgSearchValue" @keydown.enter.stop.native.prevent="search">
              <div slot="suffix" class="search-inp-icon" @click="search">
                <i class="el-icon-search"></i>
              </div>
            </el-input> -->
          </div>
          <div class="list">
            <div class="item" v-for="item in imageList" :key="item.id">
              <div class="item-s">
                <svg width="16" height="16" viewBox="0 0 48 48" fill="none" xmlns="http://www.w3.org/2000/svg"><path d="M24 6V42" stroke="#888" stroke-width="4" stroke-linecap="round" stroke-linejoin="round"/><path d="M4 34L16 12V34H4Z" fill="none" stroke="#888" stroke-width="4" stroke-linejoin="round"/><path d="M44 34H32V12L44 34Z" fill="none" stroke="#888" stroke-width="4" stroke-linejoin="round"/></svg>
              </div>
              <div class="item-l">
                <div class="item-l-t">
                  <div class="item-title">
                    <span style="color: rgba(16, 16, 16, 0.8);"> {{ item.userName }} / </span>
                    <span class="item-tag nowrap" :title="item.tag">{{ item.tag }}</span>
                    <svg v-if="item.type == 5" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24" width="20" height="20"><defs></defs><g><path fill="#FF6200" d="M11.99 2C6.47 2 2 6.48 2 12s4.47 10 9.99 10C17.52 22 22 17.52 22 12S17.52 2 11.99 2zm4.24 16L12 15.45 7.77 18l1.12-4.81-3.73-3.23 4.92-.42L12 5l1.92 4.53 4.92.42-3.73 3.23L16.23 18z"></path></g></svg>
                  </div>
                </div>
                <div class="item-l-m">
                  <div class="item-topics">
                    <span class="type-compute-resource"
                      :style="item.computeResourceColor ? `background-color: ${item.computeResourceColor};` : ''">{{
                        item.computeResourceShow }}</span>
                    
                    <span class="type-sys" v-for="(topic, index) in item.topicsSys" :key="`sys-${index}`">{{ topic }}</span>
                    <span class="type-pkg" v-for="(topic, index) in item.topicsPkg" :key="`pkg-${index}`">{{ topic }}</span>
                    <span v-for="(topic, index) in item.topics" :key="index">{{ topic }}</span>
                  </div>
                </div>
                <div class="item-l-b">
                  <div>
                    {{$t('imagesObj.imageTaskTips')}}<span class="type-task" v-for="(topic, index) in item.topicsTask" :key="`task-${index}`">{{ topic }} <span v-if="item.topicsTask.length-1 !== index">、</span></span>
                  </div>
                  <!-- <a v-if="item.userName" class="item-creator" :href="`/${item.userName}`">
                    <img class="ui avatar mini image" style="width: 20px; height: 20px"
                      :src="item.relAvatarLink ? item.relAvatarLink : `/user/avatar/ghost/-1`" />
                  </a>
                  <a v-else class="item-creator" href="javascript:;">
                    <img class="ui avatar mini image" style="width: 20px; height: 20px"
                      :src="`/user/avatar/ghost/-1`" />
                  </a>
                  <span class="item-descr" :title="item.description">{{ item.description }}</span> -->
                </div>
              </div>
              <div class="item-r">
                <el-button v-if="item.status == 1" @click="chooseImage(item)">{{ $t('cloudbrainObj.useImage')
                  }}</el-button>
                <span class="error-content" v-if="item.status == 0">
                  <i class="CREATING"></i>
                  <span style="color:#5a5a5a">{{ $t('cloudbrainObj.submitting') }}</span>
                </span>
                <span class="error-content" v-if="item.status == 2">
                  <i class="FAILED"></i>
                  <el-tooltip effect="dark" :content="$t('cloudbrainObj.checkImageSizeTips')" placement="left">
                    <span style="color:red">{{ $t('cloudbrainObj.submitFailed') }}</span>
                  </el-tooltip>
                </span>
              </div>
            </div>
          </div>
          <div class="pagination-c">
            <el-pagination background @current-change="dlgPageChange" :current-page="dlgPage" :page-size="dlgPageSize"
              layout="total, prev, pager, next" :total="dlgTotal">
            </el-pagination>
          </div>
        </div>
      </div>
    </el-dialog>
  </div>
</template>

<script>
import isURL from 'validator/lib/isURL';
import { getImages, getImageAvailabelFilter } from '~/apis/modules/images';
import { COMPUTER_RESOURCES_COLORS, JOB_TYPE } from '~/const';
import { getListValueWithKey } from '~/utils';

export default {
  name: "ImageSelectV1",
  props: {
    value: { type: Object, required: true },
    type: { type: Number, default: 0 }, // -1-全部,0-启智GPU,2-智算GPU
    useId: { type: Boolean, default: false },
    showTitle: { type: Boolean, default: true },
    required: { type: Boolean, default: true },
    spec: { type: String, required: true, },
    networkType: { type: [String, undefined], default: undefined },
    visualize: { type: [Boolean, undefined], default: undefined },
    configs: { type: Object, default: () => ({}) },
    tips: { type: String, default: '' },
  },
  data() {
    return {
      image: {
        image_url: '',
        image_id: '',
        image_name: ''
      },
      imageList: [],
      dlgShow: false,
      dlgLoading: false,
      dlgActiveName: 'first',
      dlgSearchValue: '',
      dlgCascaderFilter: {
        popperClass: `popper-filter-${Math.random().toString().replace('.', '')}`,
        value: [],
        options: []
      },
      dlgCascaderProps: {
        lazy: true,
        lazyLoad: (node, resolve) => {
          const { level } = node;
          const path = node.path || [];
          const framework = path[0];
          const version = path[1];
          const python = path[2];
          let networkType = undefined;
          if (this.networkType == 'no_internet') {
            networkType = 1;
          } else if (this.networkType == 'has_internet') {
            networkType = 2;
          }
          getImageAvailabelFilter({
            index: level,
            framework,
            version,
            python,
            job_type: this.configs.taskType,
            compute_resource: this.configs.computerResouce,
            spec: (this.configs.computerResouce == 'GPU' && this.configs.taskType != 'GENERAL') ? -1 : this.spec,
            recommend: this.dlgActiveName == 'first' ? true : undefined,
            mine: this.dlgActiveName == 'second' ? true : undefined,
            star: this.dlgActiveName == 'third' ? true : undefined,
            has_internet: networkType,
            visualize_required: this.visualize,
          }).then(res => {
            if (res.data.code == 0) {
              if (level == 0) {
                let data = res.data.data || [];
                if (data.indexOf('Other') >= 0) {
                  data = data.filter(item => item !== 'Other');
                  data.push('Other');
                }
                const nodes = data.map(item => ({
                  value: item,
                  label: item,
                  leaf: item == 'Other' ? true : false,
                }));
                resolve(nodes);
              } else {
                const nodes = (res.data.data || []).map(item => ({
                  value: item,
                  label: item ? item : 'None',
                  leaf: level >= (['GPU', 'NPU'].includes(this.configs.computerResouce) ? 3 : 2)
                }));
                if (!nodes.length) {
                  node.config.leaf = true;
                }
                resolve(nodes);
              }
            } else {
              resolve([]);
            }
          }).catch(err => {
            console.log(err);
            resolve([]);
          });
        }
      },
      dlgPage: 1,
      dlgPageSize: 5,
      dlgTotal: 0,

      errStatus: false,
      showInnerUrlTip: false,
      showUrlErrTip: false,
    };
  },
  watch: {
    value: {
      immediate: true,
      handler(newVal) {
        newVal = newVal || {}
        this.image.image_url = newVal.image_url || '';
        this.image.image_id = newVal.image_id || '';
        this.image.image_name = newVal.image_name || '';
        this.$emit('input', this.image);
      }
    },
    spec: {
      immediate: true,
      handler(newVal) {
        this.$emit('changeImage');
        this.resetDlgCascaderFilter();
      }
    },
  },
  computed: {
    filterImagesPlaceholder() {
      if (this.configs.computerResouce == 'GPU') {
        return this.$t('imagesObj.filterImagesPlaceholderCuda')
      }
      if (this.configs.computerResouce == 'NPU') {
        return this.$t('imagesObj.filterImagesPlaceholderCann')
      }
      if (this.configs.computerResouce == 'DCU') {
        return this.$t('imagesObj.filterImagesPlaceholderDtk')
      }
      const list = this.$t('imagesObj.filterImagesPlaceholderCuda').split('/');
      list.pop();
      return list.join('/');
    }
  },
  methods: {
    open() {
      this.dlgTotal = 0;
      this.dlgPage = 1;
      this.dlgSelectedModel = [];
      this.dlgSelectedModelList = [];
      this.searchImageData();
    },
    beforeClose(done) {
      done();
    },
    closed() { },
    dlgTabClick(tab, event) {
      this.dlgTotal = 0;
      this.dlgPage = 1;
      this.resetDlgCascaderFilter();
      this.searchImageData();
    },
    search() {
      this.dlgTotal = 0;
      this.dlgPage = 1;
      this.searchImageData();
    },
    resetDlgCascaderFilter() {
      this.dlgCascaderFilter.value = [];
      this.dlgCascaderFilter.options = [];
      this.handleDlgCascaderFilterExpandChange([]);
    },
    handleDlgCascaderFilterVisibleChange() {
      const popper = document.querySelector(`.${this.dlgCascaderFilter.popperClass}`);
      if (popper && popper.querySelector('.popper-filter-title')) return;
      const title = document.createElement('div');
      title.classList = ['popper-filter-title'];
      title.style = 'display:flex;margin-left:-1px;';
      let innerHtml = '';
      const titles = [this.$t('imagesObj.frameworkName'), this.$t('imagesObj.frameworkVersion'), this.$t('imagesObj.pyVersion')];
      if (this.configs.computerResouce == 'GPU') {
        titles.push(this.$t('imagesObj.cudaVersion'))
      }
      if (this.configs.computerResouce == 'DCU') {
        titles.push(this.$t('imagesObj.dtkVersion'))
      }
      if (this.configs.computerResouce == 'NPU') {
        titles.push(this.$t('imagesObj.cannVersion'))
      }
      titles.forEach((item, index) => {
        innerHtml += `<div class="popper-filter-title-item" 
        style="display:flex;align-items:center;height:30px;width:180px;box-sizing:border-box;color:rgb(136, 136, 136);
            ${index != 0 ? 'border-left:1px solid #E4E7ED;display:none;' : ''}padding-top:10px;padding-left:30px;font-size:12px;">${item}</div>`;
      })
      title.innerHTML = innerHtml;
      popper.prepend(title);
    },
    handleDlgCascaderFilterExpandChange(value) {
      const popper = document.querySelector(`.${this.dlgCascaderFilter.popperClass}`);
      if (!popper) return;
      const title = popper.querySelector('.popper-filter-title');
      if (!title) return;
      const items = title.querySelectorAll('.popper-filter-title-item');
      items.forEach((item, index) => {
        const style = item.style;
        if (index <= value.length) {
          style.display = 'flex';
        } else {
          style.display = 'none';
        }
      })
    },
    handleDlgCascaderFilterChange() {
      this.search();
    },
    searchImageData() {
      const tabName = this.dlgActiveName;
      const typeMap = {
        'first': '0',
        'second': '1',
        'third': '2',
      };
      let networkType = undefined;
      if (this.networkType == 'no_internet') {
        networkType = 1;
      } else if (this.networkType == 'has_internet') {
        networkType = 2;
      }
      const params = {
        type: typeMap[tabName],
        q: this.dlgSearchValue,
        page: this.dlgPage,
        pageSize: this.dlgPageSize,
        cloudbrainType: this.type,
        jobType: this.configs.taskType,
        computeResource: this.configs.computerResouce,
        framework: this.dlgCascaderFilter.value[0],
        frameworkVersion: this.dlgCascaderFilter.value[1],
        python: this.dlgCascaderFilter.value[2],
        cuda: this.configs.computerResouce == 'GPU' ? this.dlgCascaderFilter.value[3] : '',
        dtk: this.configs.computerResouce == 'DCU' ? this.dlgCascaderFilter.value[3] : '',
        cann: this.configs.computerResouce == 'NPU' ? this.dlgCascaderFilter.value[3] : '',
        spec: (this.configs.computerResouce == 'GPU' && this.configs.taskType != 'GENERAL') ? -1 : this.spec,
        // trainType: (this.configs.computerResouce == 'GPU' && this.configs.taskType != 'GENERAL') ?
        //   undefined : getListValueWithKey(JOB_TYPE, this.configs.taskType, 'k', 'train_type'),
        trainType: getListValueWithKey(JOB_TYPE, this.configs.taskType, 'k', 'train_type'),
        onlyOpenIImage: (this.configs.computerResouce == 'GPU' && this.configs.taskType != 'GENERAL') ? true : undefined,
        hasInternet: networkType,
        visualizeRequired: this.visualize,
      }
      this.dlgLoading = true;
      getImages(params).then(res => {
        this.dlgLoading = false;
        const data = res.data?.images || [];
        data.forEach(item => {
          const topicsSys = [];
          if (item.framework) {
            topicsSys.push(`${item.framework} ${item.frameworkVersion}`.trim());
          }
          if (item.pythonVersion) {
            topicsSys.push(`Python ${item.pythonVersion}`);
          }
          if (item.cudaVersion) {
            topicsSys.push(`Cuda ${item.cudaVersion}`);
          }
          if (item.dtkVersion) {
            topicsSys.push(`Dtk ${item.dtkVersion}`);
          }
          if (item.cannVersion) {
            topicsSys.push(`Cann ${item.cannVersion}`);
          }
          if (item.operationSystem) {
            topicsSys.push(`${item.operationSystem} ${item.operationSystemVersion}`.trim());
          }
          const topicsPkg = [];
          const thirdPackages = item.thirdPackages.split('\n');
          thirdPackages.forEach(pkgLine => {
            if (pkgLine) {
              topicsPkg.push(pkgLine.trim().replace('==', ' '));
            }
          });
          item.topicsSys = topicsSys;
          item.topicsPkg = topicsPkg;

          const topicsTask = [];
          const trainTypes = item.trainType.split('&');
          trainTypes.forEach(type => {
            if (type) {
              if(type === 'Notebook'){
                if(['GCU', 'GPU'].includes(item.compute_resource)){
                  topicsTask.push(this.$t('TaskTypeTitle.Notebook'));
                }else{
                  topicsTask.push(this.$t('TaskTypeTitle.Notebook1'));
                }
              }else{
                topicsTask.push(this.$t('TaskTypeTitle.' + type));
              }
            }
          });
          item.topicsTask = topicsTask
          console.log(item.trainType, topicsTask)
          const compute_resource = item.compute_resource || 'GPU';
          item.computeResourceColor = COMPUTER_RESOURCES_COLORS[compute_resource];
          item.computeResourceShow = this.$t('computeResourceTitle.' + compute_resource);
        });
        this.imageList = data;
        this.dlgTotal = parseInt(res.data?.count || 0);
      }).catch(err => {
        this.dlgLoading = false;
        console.log(err);
      });
    },
    imageChange() {
      this.image.image_url = this.image.image_name.trim();
      this.image.image_id = '';
      this.$nextTick(() => {
        this.$emit('input', this.image);
        this.$emit('change', this.image);
        this.check(true);
      });
    },
    chooseImage(data) {
      console.log('chooseImage', data);
      this.image.image_url = (data.place || '').trim();
      this.image.image_id = (data.image_id || '').toString().trim();
      this.image.image_name = (data.tag || '').trim();
      this.dlgShow = false;
      this.$emit('input', this.image);
      this.$emit('change', this.image);
      this.check();
    },
    dlgPageChange(page) {
      this.dlgPage = page;
      this.searchImageData();
    },
    checkInnerUrlErr(url) {
      const regex = /^dockerhub\.pcl\.ac\.cn|^192\./;
      if (url && regex.test(url)) {
        this.showInnerUrlTip = true;
        return true;
      }
      this.showInnerUrlTip = false;
      return false;
    },
    checkUrlErr(url) {
      const exceptExtensionName = [
        'txt', 'doc', 'pdf', 'ppt', 'pptx', 'xls', 'xlsx', 'html', 'htm', 'jpg', '.jpeg', 'png', 'gif', 'mp3', 'mp4',
        'exe', 'dll', 'psd', 'svg', 'json', 'xml', 'csv', 'log', 'bat', 'ini', 'sql', 'java', 'py', 'c', 'cpp',
        'js', 'css', 'git'
      ];
      const extensionName = url.indexOf('.') >= 0 ? url.split('.').pop() : '';
      if (url && exceptExtensionName.indexOf(extensionName) >= 0) {
        this.showUrlErrTip = true;
        return true;
      }
      if (url && !isURL(this.image.image_url, {})) {
        console.log('url', url)
        const reg = /^([0-9a-z\-_.]+\/)?([0-9a-z\-_.]+\/)?([0-9a-z\-_.]+)(\:[0-9a-z\-_.]+)?$/;
        if (!reg.test(url)) {
          console.log("xxxxxxxxxxxxxx")
          this.showUrlErrTip = true;
          return true;
        }
      }
      this.showUrlErrTip = false;
      return false;
    },
    check(flag=false) {
      this.showUrlErrTip = false;
      this.showInnerUrlTip = false;
      this.errStatus = false;
      if (this.required) {
        if (this.image.image_name == '') {
          this.errStatus = true;
        }
        if(this.checkUrlErr(this.image.image_name) && flag){
          this.errStatus = true;
        }
      }
      return !this.errStatus;
    },
  },
  beforeMount() {
    this.resetDlgCascaderFilter();
  },
  mounted() { },
};
</script>

<style scoped lang="less">
@import 'cloudbrain.less';

.form-row {
  .right-area {
    display: flex;
    align-items: center;

    // .btn-select {
    //   align-self: flex-start;
    //   margin-top: 9px;
    //   display: flex;
    //   justify-content: center;
    //   align-items: center;
    //   cursor: pointer;
    //   color: rgb(3, 102, 214);

    //   i {
    //     margin-right: 4px;
    //   }
    // }
  }
}

.model-dlg {
  /deep/.el-dialog__body {
    padding-top: 0;
  }

  .dlg-content {
    display: flex;
    flex-direction: column;
    min-height: 300px;
    .header-c{
      display: flex;
      align-items: center;
      height: 40px;
      .header-text{
        font-size: 18px;
        font-weight: 700;
        line-height: 20px;
        color: #101010;
      }
      .search-inp {
        overflow: hidden;
        width: 246px;
        z-index: 5;
        position: relative;
        margin-left: 16px;
        /deep/ .el-input__inner{
          border: none;
          background: #f5f5f5;
        }
        .search-inp-icon {
          height: 100%;
          display: flex;
          justify-content: center;
          align-items: center;
          width: 22px;
          cursor: pointer;
        }
      }
    }
    .main-area {
      flex: 1;
      position: relative;
      overflow: hidden;

      .image-tabs-c {
        display: flex;
        align-items: center;

        .image-tabs {
          flex: 1;
          overflow: hidden;
          margin-right: 5px;
        }
      }

      .filter-c {
        display: flex;
        align-items: flex-end;
        justify-content: space-between;

        .cascader-c {
          display: flex;
          align-items: flex-end;

          .cascader-tit {
            margin-bottom: 6px;
          }

          .cascader-content-c {
            .cascader-tips {
              font-size: 12px;
              padding-left: 16px;
            }
          }
        }

        .image-filter {
          margin-top: -1px;
          width: 340px;
        }

        .search-inp {
          overflow: hidden;
          width: 330px;
          z-index: 5;
          position: relative;
          margin-top: -1px;

          .search-inp-icon {
            height: 100%;
            display: flex;
            justify-content: center;
            align-items: center;
            width: 22px;
            cursor: pointer;
          }
        }
      }

      .list {
        min-height: 390px;

        .item {
          display: flex;
          // align-items: center;
          padding: 1rem 0px;
          border-bottom: 1px solid rgb(245, 245, 245);
          .item-s {
            margin-right: 10px;
            margin-top: 2px;
          }
          .item-l {
            flex: 1;
            overflow: hidden;

            .item-l-t {
              display: flex;
              align-items: center;
              justify-content: space-between;
              margin-left: 2px;
              .item-title {
                display: flex;
                align-items: center;
                max-width: 80%;
                overflow: hidden;
                text-overflow: ellipsis;
                white-space: nowrap;
                .item-tag{
                  color: rgba(0, 102, 255, 1);
                  font-size: 14px;
                  // font-family: Arial;
                  font-weight: 700;
                }
                svg{
                  min-width: 20px;
                  margin-left: 10px;
                }
                // span {
                //   font-size: 15px;
                //   color: #0366d6;
                //   margin-right: 5px;
                //   vertical-align: middle;
                // }

                // img {
                //   vertical-align: middle;
                // }
              }
            }

            .item-l-m {
              margin-top: 8px;

              .item-topics {
                display: flex;
                flex-wrap: wrap;

                span {
                  display: flex;
                  align-items: center;
                  justify-content: center;
                  font-size: .85714286rem;
                  margin: 0 0.14285714em;
                  padding: 0.3em 0.5em;
                  background-color: #e8e8e8;
                  color: rgba(0, 0, 0, .6);
                  font-weight: 700;
                  border-radius: 0.28571429rem;
                  cursor: pointer;
                  line-height: 1;
                  margin-bottom: 3px;
                  
                  &.type-sys {
                    background: rgba(50, 145, 248, 0.2);
                  }
                  &.type-task {
                    background: rgba(255,245,226,1);
                  }
                  &.type-pkg {
                    background: rgba(91, 185, 115, 0.2);
                  }

                  &.type-compute-resource {
                    color: white;
                  }
                }
              }
            }

            .item-l-b {
              margin-top: 8px;
              display: flex;
              align-items: center;
              color: rgba(136, 136, 136, 1);
              margin-left: 2px;
              // .item-creator {
              //   margin-right: 6px;
              //   flex-shrink: 0;
              // }

              // .item-descr {
              //   overflow: hidden;
              //   text-overflow: ellipsis;
              //   display: inline-block;
              //   white-space: nowrap;
              // }
            }
          }

          .item-r {
            width: 120px;
            display: flex;
            // align-items: center;
            justify-content: flex-end;

            button {
              color: #3291f8;
              border-color: #3291f8;
              height: 30px;
            }

            .error-content {
              display: flex;
              align-items: center;
              justify-content: center;
              font-size: 12px;

              i {
                margin-right: 5px;
              }
            }
          }
        }
      }

      .pagination-c {
        margin-top: 25px;
        text-align: center;
      }
    }
  }
}
</style>
