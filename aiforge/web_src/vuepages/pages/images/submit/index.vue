<template>
  <div>
    <div class="ui container">
      <div class="form-container" v-loading="submitLoading">
        <div class="form-head">
          <h4 v-if="pageType == 'edit'">{{ $t('imagesObj.editImage') }}</h4>
          <h4 v-if="pageType == 'submit' || pageType == 'submitAdmin'">{{ $t('imagesObj.submitImage') }}</h4>
          <h4 v-if="pageType == 'apply'">{{ $t('imagesObj.appyImage') }}</h4>
        </div>
        <div class="form-body">
          <div class="form-body-content">
            <div class="form-row form-row-cluster" v-show="false">
              <div class="title align-items-center"><span class="required">{{ $t('modelManage.useCluster') }}</span>
              </div>
              <div class="content">
                <div class="list">
                  <a class="item" href="javascript:;" v-for="item in typeList" :key="item.k" @click="changeType(item)"
                    :class="form.type == item.k ? 'focus' : ''">
                    <i class="icon ri-archive-drawer-line"></i>
                    <span>{{ item.v }}</span>
                  </a>
                </div>
              </div>
            </div>
            <div class="form-row">
              <div class="title"><span class="required">{{ $t('imagesObj.imageTag') }}</span></div>
              <div class="content" :class="errors.tag ? 'error' : ''">
                <el-input class="field-input" :disabled="pageType != 'submit' && pageType != 'submitAdmin'"
                  v-model="form.tag" @blur="checkTag" @input="checkTag" :maxlength="50"
                  :placeholder="$t('imagesObj.imageTagPlaceholder')" :autofocus="true"></el-input>
                <div class="tips">{{ $t('imagesObj.imageTagInputTips') }}</div>
              </div>
            </div>
            <div class="form-row" style="visibility: hidden;height: 0;margin: 0;">
              <div class="title"><span class="required">{{ $t('imagesObj.imageTaskType') }}</span></div>
              <div class="content">
                <el-select class="field-input" multiple v-model="form.trainType" :disabled="pageType != 'submitAdmin'"
                  @change="checkNull('trainType')">
                  <el-option v-for="item in trainCopyTypeList" :key="item.k" :value="item.k"
                    :label="item.v"></el-option>
                </el-select>
              </div>
            </div>
            <div class="form-row">
              <div class="title"><span class="required">{{ $t('cloudbrainObj.computeResource') }}</span></div>
              <div class="content">
                <el-select class="field-input" v-model="form.compute_resource" :disabled="pageType != 'submitAdmin'"
                  @change="changeComputeResource">
                  <el-option v-for="item in computeResourceList" :key="item.k" :value="item.k"
                    :label="item.v"></el-option>
                </el-select>
              </div>
            </div>
            <div class="form-row" v-if="pageType == 'submitAdmin'">
              <div class="title"><span :class="form.compute_resource == 'GPU' ? 'required' : ''">{{
                $t('imagesObj.imageAdress') }}</span></div>
              <div class="content" :class="errors.place ? 'error' : ''">
                <el-input class="field-input" v-model="form.place" @blur="checkNull('place')"
                  @input="checkNull('place')" :maxlength="300"
                  :placeholder="$t('imagesObj.imageAdressPlaceholder')"></el-input>
              </div>
            </div>
            <div class="form-row" v-if="pageType == 'submitAdmin' || (pageType == 'edit' && pageFrom == 'imageAdmin')">
              <div class="title"><span>{{ $t('resourcesManagement.aiCenter') }}</span></div>
              <div class="content">
                <el-input class="field-input" type="textarea" v-model="form.aICenterImage" :maxlength="1500" :rows="3"
                  :placeholder="$t('imagesObj.imageAiCenterPlaceholder')"></el-input>
              </div>
            </div>
            <div class="form-row">
              <div class="left">
                <div class="title"><span class="required">{{ $t('imagesObj.frameworkName') }}</span></div>
                <div class="content" :class="errors.framework ? 'error' : ''">
                  <el-select class="field-input" v-model="form.framework" @change="changeFramework">
                    <el-option v-for="item in frameworkList" :key="item.k" :value="item.k" :label="item.v"></el-option>
                  </el-select>
                </div>
              </div>
              <div class="right">
                <div v-show="frameworkVersionList.length" class="title"><span class="required">{{
                  $t('imagesObj.version')
                    }}</span></div>
                <div v-show="frameworkVersionList.length" class="content"
                  :class="errors.framework_version ? 'error' : ''">
                  <el-select class="field-input" v-model="form.framework_version"
                    @change="checkNull('framework_version')" filterable>
                    <el-option v-for="item in frameworkVersionList" :key="item.k" :value="item.k"
                      :label="item.v"></el-option>
                  </el-select>
                </div>
              </div>
            </div>
            <div class="form-row">
              <div class="title"><span class="required">{{ $t('imagesObj.pyVersion') }}</span></div>
              <div class="content" :class="errors.python ? 'error' : ''">
                <el-select class="field-input" v-model="form.python" @change="checkNull('python')" filterable>
                  <el-option v-for="item in pythonList" :key="item.k" :value="item.k" :label="item.v"></el-option>
                </el-select>
              </div>
            </div>
            <div class="form-row" v-if="form.compute_resource == 'GPU'">
              <div class="title"><span class="required">{{ $t('imagesObj.cudaVersion') }}</span></div>
              <div class="content" :class="errors.cuda ? 'error' : ''">
                <el-select class="field-input" v-model="form.cuda" @change="checkNull('cuda')" filterable>
                  <el-option v-for="item in cudaList" :key="item.k" :value="item.k" :label="item.v"></el-option>
                </el-select>
              </div>
            </div>
            <div class="form-row" v-if="form.compute_resource == 'NPU'">
              <div class="title"><span class="required">{{ $t('imagesObj.cannVersion') }}</span></div>
              <div class="content" :class="errors.cann ? 'error' : ''">
                <el-select class="field-input" v-model="form.cann" @change="checkNull('cann')" filterable>
                  <el-option v-for="item in cannList" :key="item.k" :value="item.k" :label="item.v"></el-option>
                </el-select>
              </div>
            </div>
            <div class="form-row" v-if="form.compute_resource == 'DCU'">
              <div class="title"><span class="required">{{ $t('imagesObj.dtkVersion') }}</span></div>
              <div class="content" :class="errors.dtk ? 'error' : ''">
                <el-select class="field-input" v-model="form.dtk" @change="checkNull('dtk')" filterable>
                  <el-option v-for="item in dtkList" :key="item.k" :value="item.k" :label="item.v"></el-option>
                </el-select>
              </div>
            </div>
            <div class="form-row">
              <div class="left">
                <div class="title"><span>{{ $t('imagesObj.operationSystemName') }}</span></div>
                <div class="content">
                  <el-input class="field-input" v-model="form.operationSystem" :maxlength="100"
                    :placeholder="$t('imagesObj.operationSystemNamePlaceholder')"></el-input>
                </div>
              </div>
              <div class="right">
                <div class="title"><span>{{ $t('imagesObj.version') }}</span></div>
                <div class="content">
                  <el-input class="field-input" v-model="form.operationSystemVersion" :maxlength="50"
                    :placeholder="$t('imagesObj.operationSystemVersionPlaceholder')"></el-input>
                </div>
              </div>
            </div>
            <div class="form-row">
              <div class="title"><span>{{ $t('imagesObj.thirdPackages') }}</span></div>
              <div class="content">
                <el-input class="field-input" type="textarea" v-model="form.thirdPackages" :rows="3"
                  @blur="checkThirdPackages" :placeholder="$t('imagesObj.thirdPackagesPlaceholder')"
                  :maxlength="1000"></el-input>
              </div>
            </div>
            <div class="form-row">
              <div class="title"><span>{{ $t('imagesObj.topic') }}</span></div>
              <div class="content">
                <el-select class="field-input topic-input" popper-class="popper-topic-input" v-model="form.topics"
                  remote :remote-method="topicRemoteMethod" :loading="topicLoading" multiple filterable allow-create
                  default-first-option clearable :placeholder="$t('imagesObj.topicPlaceholder')">
                  <el-option v-for="item in topicList" :key="item.k" :value="item.k" :label="item.v"></el-option>
                </el-select>
                <div class="tips tips-ext">
                  <i class="ri-error-warning-line"></i>
                  <span v-html="$t('imagesObj.topicTips')"></span>
                </div>
              </div>
            </div>
            <div class="form-row">
              <div class="title"><span class="required">{{ $t('imagesObj.descr') }}</span></div>
              <div class="content" :class="errors.description ? 'error' : ''">
                <el-input class="field-input" type="textarea" v-model="form.description" :rows="3"
                  @blur="checkNull('description')" @input="checkNull('description')"
                  :placeholder="$t('imagesObj.descrPlaceholder')" :maxlength="1000"></el-input>
              </div>
            </div>
            <div class="form-row" style="margin-top:-10px" v-if="pageType == 'submitAdmin'">
              <div class="title"><span></span></div>
              <div class="content" style="padding-top:8px">
                <el-radio-group v-model="form.isRecommend">
                  <el-radio label="true">{{ $t('imagesObj.recommend') }}</el-radio>
                  <el-radio label="false">{{ $t('imagesObj.notRecommend') }}</el-radio>
                </el-radio-group>
              </div>
            </div>
            <template v-if="pageType == 'submit'">
              <div class="form-row" style="margin-top:-10px" v-if="form.compute_resource == 'GPU' || isXJUCenter">
                <div class="title"><span></span></div>
                <div class="content">
                  <div class="tips tips-ext" style="font-size:14px;">
                    <i class="ri-error-warning-line"></i>
                    <span>
                      <span class="light" v-html="$t('imagesObj.submitTips')"></span>
                    </span>
                  </div>
                </div>
              </div>
              <div class="form-row" style="margin-top:-10px" v-if="form.compute_resource == 'NPU' && !isXJUCenter">
                <div class="title"><span></span></div>
                <div class="content">
                  <div class="tips tips-ext" style="font-size:14px;">
                    <i class="ri-error-warning-line"></i>
                    <span>
                      <span class="light" v-html="$t('imagesObj.submitNpuTip1')"></span>
                    </span>
                  </div>
                  <div class="tips tips-ext" style="font-size:14px;">
                    <i class="ri-error-warning-line"></i>
                    <span>
                      <span class="light" v-html="$t('imagesObj.submitNpuTip2')"></span>
                    </span>
                  </div>
                </div>
              </div>
              <div class="form-row" style="margin-top:-10px" v-if="form.compute_resource == 'DCU' && !isXJUCenter">
                <div class="title"><span></span></div>
                <div class="content">
                  <!-- <div class="tips tips-ext" style="font-size:14px;">
                    <i class="ri-error-warning-line"></i>
                    <span>
                      <span class="light" v-html="$t('imagesObj.submitDcuTip1')"></span>
                    </span>
                  </div> -->
                  <div class="tips tips-ext" style="font-size:14px;">
                    <i class="ri-error-warning-line"></i>
                    <span>
                      <span class="light" v-html="$t('imagesObj.submitDcuTip2')"></span>
                    </span>
                  </div>
                </div>
              </div>
            </template>
            <div class="form-row">
              <div class="title"></div>
              <div class="content">
                <el-button type="primary" size="default" class="submit-btn" @click="submit">{{
                  pageType == 'apply' ? $t('imagesObj.submitApply') : $t('submit') }}</el-button>
                <el-button class="cancel-btn" size="default" @click="cancel">{{ $t('cancel') }}</el-button>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import { getStaticFile } from '~/apis/modules/common';
import { searchImageTopics, submitImage } from '~/apis/modules/images';
import { i18n } from '~/langs';
import { getUrlSearchParams } from '~/utils';

export default {
  data() {
    return {
      pageType: window._PageType, // submit|submitAdmin|edit|apply
      pageFrom: window._PageFrom, // submit|imageSquare|imageAdmin|apply
      pageSubmitLink: window._PageSubmitLink,
      pageRepoLink: window._PageRepoLink,
      pageAiCenter: window._ImageAiCenter || '',
      isXJUCenter: (window._ImageAiCenter || '').indexOf('xju') == 0, // 新大中心
      originData: {},
      form: {
        type: window._ImageType || '2', // 0启智GPU|2智算GPU
        tag: '',
        trainType: [],
        compute_resource: window._ImageComputeResource || 'GPU',
        place: '',
        aICenterImage: '',
        framework: '',
        framework_version: '',
        python: '',
        cuda: '',
        dtk: '',
        cann: '',
        operationSystem: '',
        operationSystemVersion: '',
        thirdPackages: '',
        topics: [],
        description: '',
        isRecommend: 'true',
        image_id: ''
      },
      typeList: [{ k: '0', v: i18n.t('imagesObj.openIGPU') }, { k: '2', v: i18n.t('imagesObj.c2netGPU') }],
      computeResourceList: [],
      trainTypeList: [],
      topicLoading: false,
      topicList: [],
      frameworkList: [],
      frameworkVersionList: [],
      pythonList: [],
      cudaList: [],
      dtkList: [],
      cannList: [],
      condtionsData: {},
      errors: {
        tag: false,
        place: false,
        framework: false,
        framework_version: false,
        trainType: false,
        python: false,
        cuda: false,
        dtk: false,
        cann: false,
        description: false,
      },
      submitLoading: false,
    };
  },
  components: {},
  
  computed: {
    trainCopyTypeList() {
      if (['GCU', 'GPU'].includes(this.form.compute_resource)) {
        return this.trainTypeList
      } else {
        return this.trainTypeList.map((item) => {
          if (item.k == 'Notebook') {
            return { k: item.k, v: this.$t('TaskTypeTitle.Notebook1') }
          } else {
            return {...item}
          }
        })
        
      }
      
    },
    
    task() {
      return this.mainDataItem.task || {};
    },
    
    displayJobName() {
      return this.task.display_job_name || '';
    },
    
    jobType() {
      return this.task.job_type || '';
    }
  },
  methods: {
    changeType(item) {
      this.form.type = item.k;
    },
    checkTag() {
      this.form.tag = this.form.tag.trim();
      const reg = /^[a-zA-Z][a-zA-Z0-9_.\-]{0,49}$/;
      this.errors.tag = !reg.test(this.form.tag);
      return this.errors.tag;
    },
    checkNull(key) {
      this.form[key] = this.form[key];
      this.errors[key] = !this.form[key];
      if (key == 'place' && this.form.compute_resource != 'GPU') {
        this.errors[key] = false;
      }
      return this.errors[key];
    },
    checkThirdPackages() {
      const last = this.form.thirdPackages.slice(-1) == '\n' ? '\n' : '';
      const lines = this.form.thirdPackages.split('\n');
      this.form.thirdPackages = lines.map(item => item.trim()).filter(item => item).slice(0, 10).join('\n') + last;
    },
    changeComputeResource() {
      if (this.form.compute_resource == 'GPU') {
        this.form.cuda = this.cudaList.length ? this.cudaList[0].k : '';
      } else if (this.form.compute_resource == 'NPU') {
        this.form.cann = this.cannList.length ? this.cannList[0].k : '';
      } else if (this.form.compute_resource == 'DCU') {
        this.form.dtk = this.dtkList.length ? this.dtkList[0].k : '';
      } else {
        this.form.cuda = '';
        this.errors.cuda = false;
        this.errors.place = false;
        this.form.cann = '';
        this.errors.cann = false;
        this.form.dtk = '';
        this.errors.dtk = false;
      }
    },
    changeFramework(framework) {
      const versions = (this.condtionsData['framework_version'][framework] || []).map(item => {
        return {
          k: item,
          v: item,
        };
      });
      this.form.framework_version = '';
      this.frameworkVersionList = versions;
      if (versions.length) {
        this.form.framework_version = versions[0].k;
      } else {
        this.form.framework_version = '';
      }
      this.checkNull('framework');
    },
    topicRemoteMethod(query) {
      if (query !== '') {
        this.topicLoading = true;
        searchImageTopics({ q: query }).then(res => {
          this.topicLoading = false;
          res = res.data;
          if (res.topics) {
            this.topicList = [];
            const find = res.topics.find(item => item.topic_name === query);
            find && this.topicList.push({ k: find.topic_name, v: find.topic_name });
            for (let i = 0, iLen = res.topics.length; i < iLen; i++) {
              const topic = res.topics[i];
              if (!find || topic.topic_name !== find.topic_name) {
                this.topicList.push({
                  k: topic.topic_name,
                  v: topic.topic_name
                });
              }
            }
          } else {
            this.topicList = [];
          }
        }).catch(err => {
          this.topicLoading = false;
          this.topicList = [];
          console.log(err);
        });
      } else {
        this.topicList = [];
      }
    },
    submit() {
      let hasError = false;
      for (let key in this.errors) {
        if(key!='trainType'){
          this.form[key] = this.form[key].trim();
        }
        let isError = false;
        if (this.pageType == 'submit' && key == 'place') continue;
        if (key == 'place' && this.form.compute_resource != 'GPU') continue;
        if (key == 'tag') {
          isError = this.checkTag();
        } else if (key == 'framework_version') {
          if (this.frameworkVersionList.length) {
            isError = this.checkNull(key);
          }
        } else if (key == 'cuda') {
          if (this.form.compute_resource == 'GPU') {
            isError = this.checkNull(key);
          } else {
            this.form.cuda = '';
          }
        } else if (key == 'dtk') {
          if (this.form.compute_resource == 'DCU') {
            isError = this.checkNull(key);
          } else {
            this.form.dtk = '';
          }
        }else if (key == 'cann') {
          if (this.form.compute_resource == 'NPU') {
            isError = this.checkNull(key);
          } else {
            this.form.cann = '';
          }
        } else {
          isError = this.checkNull(key);
        }
        if (isError) {
          hasError = true;
        }
      }
      if (hasError) return;
      const subData = {
        link: this.pageSubmitLink,
        type: this.form.type,
        tag: this.form.tag.trim(),
        trainType: this.form.trainType.join('&'),
        computeResource: this.form.compute_resource,
        framework: this.form.framework,
        frameworkVersion: this.form.framework_version,
        cudaVersion: this.form.cuda,
        dtkVersion: this.form.dtk,
        cannVersion: this.form.cann,
        pythonVersion: this.form.python,
        operationSystem: this.form.operationSystem.trim(),
        operationSystemVersion: this.form.operationSystemVersion.trim(),
        thirdPackages: this.form.thirdPackages.trim(),
        topics: this.form.topics.join(','),
        description: this.form.description.trim(),
        imageId: this.form.image_id
      };
      if (this.pageType == 'edit' || this.pageType == 'apply') {
        subData.id = this.originData.id;
      }
      if (this.pageType == 'submitAdmin') {
        subData.isRecommend = this.form.isRecommend;
        subData.place = this.form.place;
      }
      if (this.pageType == 'submitAdmin' || (this.pageType == 'edit' && this.pageFrom == 'imageAdmin')) {
        subData.aICenterImage = this.form.aICenterImage;
      }
      // console.log('submit subData', this.pageSubmitLink, this.pageRepoLink, subData);
      // return;
      this.submitLoading = true;
      submitImage(subData).then(res => {
        res = res.data;
        if (res.Code == 0) {
          this.$message({
            type: 'success',
            message: this.$t('submittedSuccessfully'),
          });
          setTimeout(() => {
            this.cancel('success');
          }, 200);
        } else {
          this.submitLoading = false;
          this.$message({
            type: 'error',
            message: res.Message,
          });
        }
      }).catch(err => {
        this.submitLoading = false;
        console.log(err);
        this.$message({
          type: 'error',
          message: this.$t('submittedFailed'),
        });
      });
    },
    cancel(status) {
      if (this.pageFrom == 'submit') {
        if (status == 'success') {
          location.href = `/explore/images_my`;
        } else {
          window.close();
        }
      } else if (this.pageFrom == 'imageSquare' || this.pageFrom == 'apply') {
        location.href = `/explore/images_my`;
      } else if (this.pageFrom == 'imageAdmin') {
        location.href = `/admin/images`
      }
    }
  },
  beforeMount() {
    if (this.pageType == 'submit') {
      this.typeList = this.typeList.filter(item => item.k == this.form.type);
    }
    if (this.pageType == 'submitAdmin') {
      this.form.type = '2';
      this.form.compute_resource = 'GPU';
      this.typeList = this.typeList.filter(item => item.k == this.form.type);
    }
    if (this.pageType == 'edit' || this.pageType == 'apply') {
      const image = window._Image || {};
      console.log(image)
      this.originData = image;
      this.form.type = image.cloudbrainType.toString();
      this.typeList = this.typeList.filter(item => item.k == this.form.type);
      this.form.tag = image.tag;
      this.form.compute_resource = image.compute_resource || 'GPU';
      this.form.place = image.place;
      this.form.aICenterImage = JSON.stringify(image.AiCenterImages || []);
      this.form.operationSystem = image.operationSystem;
      this.form.operationSystemVersion = image.operationSystemVersion;
      this.form.thirdPackages = image.thirdPackages;
      this.form.topics = image.topics || [];
      this.form.description = image.description;
    }
    getStaticFile('/images_version.json').then(res => {
      const data = res.data;
      if (data) {
        this.condtionsData = data;
        this.frameworkList = (data.framework || []).map(item => {
          return { k: item, v: item };
        });
        this.form.framework = this.frameworkList.length ? this.frameworkList[0].k : '';
        if (this.form.compute_resource == 'NPU') {
          this.form.framework = 'MindSpore';
        }
        this.changeFramework(this.form.framework);
        this.cudaList = (data.cuda || []).map(item => {
          return { k: item, v: `Cuda ${item}` };
        });
        this.form.cuda = this.cudaList.length ? this.cudaList[0].k : '';
        this.dtkList = (data.dtk || []).map(item => {
          return { k: item, v: `Dtk ${item}` };
        });
        this.form.dtk = this.dtkList.length ? this.dtkList[0].k : '';
        this.cannList = (data.cann || []).map(item => {
          return { k: item, v: `Cann ${item}` };
        });
        this.form.cann = this.cannList.length ? this.cannList[0].k : '';
        this.pythonList = (data.python || []).map(item => {
          return { k: item, v: `Python ${item}` };
        });
        this.form.python = this.pythonList.length ? this.pythonList[0].k : '';
        if (this.pageType == 'edit' || this.pageType == 'apply') {
          const image = window._Image || {};
          this.form.framework = image.framework;
          this.changeFramework(this.form.framework);
          this.form.framework_version = image.frameworkVersion;
          this.form.cuda = image.cudaVersion;
          this.form.dtk = image.dtkVersion;
          this.form.cann = image.cannVersion;
          this.form.python = image.pythonVersion;
        }
        this.computeResourceList = (data.compute_resource || []).map(item => {
          return { k: item, v: this.$t('computeResourceTitle.' + item) || item };
        });
        console.log("data",data)
        this.trainTypeList = (data.trainType || []).map(item => {
          return { k: item, v: this.$t('TaskTypeTitle.' + item) || item };
        });
        if (this.trainTypeList.length) {
          if (this.form.trainType.length===0) {
            this.form.trainType.push(this.trainTypeList[0].k)
          }
        }
      }
    }).catch(err => {
      console.log(err);
    });
  },
  mounted() {
    const urlParams = getUrlSearchParams();
    if (urlParams.type) {
      this.form.topics.push(urlParams.type);
    }
    if(urlParams.id){
      this.form.image_id = urlParams.id
    }
    if (urlParams.trainType) {
      console.log(urlParams.trainType, urlParams.trainType.split('&'))
      urlParams.trainType.split('_').forEach((item) => {
        console.log(item, this.form.trainType)
        if (!this.form.trainType.includes(item)) {
          this.form.trainType.push(item)
        }
      })
    }
    console.log('this.form.trainType',this.form.trainType)
  },
  beforeDestroy() { },
};
</script>

<style scoped lang="less">
.form-container {
  .form-head {
    background: #f0f0f0;
    border-radius: 0.28571429rem 0.28571429rem 0 0;
    padding: 0.78571429rem 1rem;
    margin: 0 -1px;
    box-shadow: none;
    border: 1px solid #d4d4d5;
  }

  .form-body {
    min-height: 400px;
    border: 1px solid #d4d4d5;
    margin: -1px -1px;
    padding: 3em;
    background-color: #FFF;

    .form-body-content {
      max-width: 1200px;
      margin: 0 auto;
      padding-right: 120px;

      .form-row {
        display: flex;
        margin-bottom: 28px;

        .left {
          display: flex;
          width: 65%;
        }

        .right {
          display: flex;
          width: 35%;

          .title {
            width: 100px;
          }
        }

        .title {
          width: 200px;
          text-align: right;
          margin-right: 24px;
          color: #101010;
          font-size: 14px;
          display: flex;
          justify-content: flex-end;
          padding-top: 8px;
          flex-shrink: 0;

          &.align-items-center {
            padding-top: 0;
            align-items: center;
          }

          .required {
            position: relative;

            &::after {
              position: absolute;
              content: "*";
              top: -3px;
              right: -10px;
              color: red;
            }
          }
        }

        .content {
          flex: 1;

          .field-input {
            width: 100%;
          }

          .tips {
            font-size: 12px;
            color: rgba(136, 136, 136, 1);
            margin-top: 10px;

            &.tips-ext {
              display: flex;
              font-size: 12px;
              align-items: center;

              i {
                color: #f2711c;
                margin-right: 5px;
                font-size: 14px;
              }

              /deep/.light,
              .light {
                color: #f2711c;
              }
            }
          }

          .field-input {
            /deep/.el-input__inner {
              height: 37.6px;
            }

            /deep/.el-input__inner,
            /deep/.el-textarea__inner {
              color: rgba(0, 0, 0, .95);
              font-size: 14px;
              line-height: 22px;

              &:visited {
                border-color: #85b7d9;
              }

              &:focus {
                border-color: #85b7d9;
              }

              &:active {
                border-color: #85b7d9;
              }
            }
          }

          .topic-input {

            /deep/.el-input.is-focus .el-input__inner {
              border-color: #85b7d9;
            }
          }

          &.error {
            .field-input {

              /deep/.el-input__inner,
              /deep/.el-textarea__inner {
                color: #9f3a38;
                background: #fff6f6;
                border-color: #e0b4b4;

                &:visited {
                  border-color: #e0b4b4;
                }

                &:focus {
                  border-color: #e0b4b4;
                }

                &:active {
                  border-color: #e0b4b4;
                }
              }
            }
          }

          .list {
            display: flex;
            align-items: center;
            flex-wrap: wrap;

            .item {
              display: flex;
              align-items: center;
              justify-content: center;
              font-size: 12px;
              color: rgba(0, 0, 0, .87);
              border: 1px solid rgba(34, 36, 38, .15);
              margin-left: -1px;
              height: 38px;
              padding: 0 12px;
              border-left: none;
              margin-bottom: 5px;

              i {
                margin-top: -7px;
                font-size: 14px;
              }

              &.focus {
                color: #0087f5;
                border-color: #0087f5;
                border-left: 1px solid #0087f5;
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

              &:hover:not(.focus) {
                background: rgba(0, 0, 0, .03);
              }
            }
          }
        }
      }

      .submit-btn {
        background-color: #5bb973;
        color: #fff;
        border-color: #5bb973;

        &:hover {
          background-color: #16ab39;
        }

        &:active {
          background-color: #198f35;
        }

        &.is-disabled,
        &.is-disabled:active,
        &.is-disabled:focus,
        &.is-disabled:hover {
          background-color: #5bb973;
          opacity: .45;
        }

        a {
          color: #fff;
          font-weight: 700;
        }
      }

      .cancel-btn {
        background: #e0e1e2;
        color: rgba(0, 0, 0, .6);
        border-color: #e0e1e2;

        &:hover {
          background-color: #cacbcd;
        }

        &:active {
          background-color: #babbbc;
        }
      }
    }
  }
}
@media screen and (max-width: 767px) {
  .form-body{
    width: 100% !important;
    padding: 20px 10px !important;
    .form-body-content{
      padding-right: 0 !important;
      .form-row{
        flex-wrap: wrap;
        gap: 14px;
        .title{
          width: 72px !important;
        }
        .left{
          width: 100% !important;
        }
        .right{
          width: 100% !important;
        }
      }
    }
  }
}
</style>
