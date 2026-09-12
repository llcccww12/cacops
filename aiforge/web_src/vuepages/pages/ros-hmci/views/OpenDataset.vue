<template>
  <div>
    <AppBanner></AppBanner>
    <div class="repos--seach datasetSearch">
      <div class="ui container">
        <div class="ui two column centered grid">
          <div class="fourteen wide mobile ten wide tablet ten wide computer column ui form ignore-dirty">
            <div class="ui fluid action input datasetSearchInput">
              <input name="q" value="" placeholder="搜索数据集" autofocus v-model="searchValue"
                @keyup.enter="searchFlag = !searchFlag" />
              <button class="ui green button" @click="searchFlag = !searchFlag">
                {{ $t("repos.search") }}
              </button>
            </div>
          </div>
        </div>
      </div>
    </div>
    <div class="ui container">
      <div class="ui grid">
        <div class="computer only four wide computer column">
          <div class="ui sixteen wide column">
            <div style="
                            font-size: 24px;
                            color: rgba(16, 16, 16, 1);
                            height: 40px;
                            line-height: 40px;
                            margin-bottom: 2rem;
                          ">
              {{ $t("dataset") }}
            </div>
            <!-- <div class="mg-b-2">
              <div class="flex mg-b-1">
                <h3 class="font-medium">
                  {{ $t("datasets.category") }}
                  <span
                    v-if="categoryFlag"
                    @click="clearSelectLeft('category')"
                    class="mg-l-1 underline text-gray-400 text-sm"
                    style="cursor: pointer"
                    >Clear</span
                  >
                </h3>
              </div>
              <div class="flex flex-wrap">
                <a
                  class="tag"
                  :class="item.active ? 'tag-active' : 'tag-gray'"
                  v-for="item in Category"
                  :key="item.name"
                  @click="selectCategory(item)"
                  ><span>{{ $t("datasets." + item.name) }}</span></a
                >
              </div>
            </div>
            <div class="mg-b-2">
              <div class="flex mg-b-1">
                <h3 class="font-medium">
                  {{ $t("datasets.task") }}
                  <span
                    v-if="taskFlag"
                    @click="clearSelectLeft('task')"
                    class="mg-l-1 underline text-gray-400 text-sm"
                    style="cursor: pointer"
                    >Clear</span
                  >
                </h3>
              </div>
              <div class="flex flex-wrap history-content">
                <a
                  class="tag"
                  :class="item.active ? 'tag-active' : 'tag-gray'"
                  v-for="item in Task"
                  :key="item.name"
                  @click="selectTask(item)"
                  ><span>{{ $t("datasets." + item.name) }}</span></a
                >
              </div>
            </div> -->
            <div class="mg-b-2">
              <div class="flex mg-b-1">
                <h3 class="font-medium">
                  {{ $t("datasets.license") }}
                  <span v-if="licenseFlag" @click="clearSelectLeft('license')"
                    class="mg-l-1 underline text-gray-400 text-sm" style="cursor: pointer">Clear</span>
                </h3>
              </div>
              <div class="flex flex-wrap history-content">
                <a class="tag" :class="item.active ? 'tag-active' : 'tag-gray'" v-for="item in License" :key="item.name"
                  @click="selectLicense(item)"><span>{{ item.name }}</span></a>
              </div>
            </div>
          </div>
        </div>
        <div class="ui sixteen wide mobile sixteen wide tablet twelve wide computer column">
          <div class="ui sixteen wide column">
            <el-tabs v-model="activeName" @tab-click="handleClick">
              <el-tab-pane :label="$t('datasets.publick_dataset')" name="public_datasets">
                <div v-if="activeName === 'public_datasets'">
                  <public-dataset :isSigned="isSigned" :dataGet="activeName" :searchValue="searchValue"
                    :searchFlag="searchFlag" :categoryValue="categoryValue" :taskValue="taskValue"
                    :licenseValue="licenseValue" @getLabel="getChildLabel">
                  </public-dataset>
                </div>
              </el-tab-pane>
              <el-tab-pane :label="$t('datasets.my_dataset')" name="my_datasets" v-if="isSigned === 'true'">
                <div v-if="activeName === 'my_datasets'">
                  <public-dataset :isSigned="isSigned" :dataGet="activeName" :searchValue="searchValue"
                    :searchFlag="searchFlag" :categoryValue="categoryValue" :taskValue="taskValue"
                    :licenseValue="licenseValue">
                  </public-dataset>
                </div>
              </el-tab-pane>
              <el-tab-pane :label="$t('datasets.favorite_dataset')" name="my_favorite_datasets"
                v-if="isSigned === 'true'">
                <div v-if="activeName === 'my_favorite_datasets'">
                  <public-dataset :isSigned="isSigned" :dataGet="activeName" :searchValue="searchValue"
                    :searchFlag="searchFlag" :categoryValue="categoryValue" :taskValue="taskValue"
                    :licenseValue="licenseValue">
                  </public-dataset>
                </div>
              </el-tab-pane>
            </el-tabs>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>
<script>
import AppBanner from "../components/AppBanner.vue";
import { Category, Task, License } from "~/pages/dataset/square/constant.js";
//以下组件做过更改
import PublicDataset from "../components/PublicDataset.vue";
export default {
  name: "OpenDataset",
  components: {
    AppBanner,
    PublicDataset,
  },
  data() {
    return {
      Category: Category,
      Task: Task,
      License: License,
      categoryFlag: false,
      taskFlag: false,
      licenseFlag: false,
      categoryValue: "",
      taskValue: "",
      licenseValue: "",
      activeName: "public_datasets",
      isSigned: "false",
      searchValue: "",
      searchFlag: false,
    };
  },
  computed: {},
  methods: {
    handleClick(tab, event) {
      this.searchValue = "";
      this.clearAllSelctLeft();
    },
    selectCategory(item) {
      this.Category.forEach((element) => {
        if (element.name === item.name) {
          element.active = true;
        } else {
          element.active = false;
        }
      });
      this.categoryFlag = true;
      this.categoryValue = item.name;
    },
    selectTask(item) {
      this.Task.forEach((element) => {
        if (element.name === item.name) {
          element.active = true;
        } else {
          element.active = false;
        }
      });
      this.taskValue = item.name;
      this.taskFlag = true;
    },
    selectLicense(item) {
      this.License.forEach((element) => {
        if (element.name === item.name) {
          element.active = true;
        } else {
          element.active = false;
        }
      });
      this.licenseValue = item.name;
      this.licenseFlag = true;
    },
    clearSelectLeft(type) {
      if (type === "category") {
        this.Category.forEach((element) => {
          element.active = false;
        });
        this.categoryValue = "";
        this.categoryFlag = false;
      } else if (type === "task") {
        this.Task.forEach((element) => {
          element.active = false;
        });
        this.taskValue = "";
        this.taskFlag = false;
      } else {
        this.License.forEach((element) => {
          element.active = false;
        });
        this.licenseValue = "";
        this.licenseFlag = false;
      }
    },
    clearAllSelctLeft() {
      if (this.categoryFlag) {
        this.clearSelectLeft("category");
      }
      if (this.taskFlag) {
        this.clearSelectLeft("task");
      }
      if (this.licenseFlag) {
        this.clearSelectLeft("license");
      }
    },
    getChildLabel(data) {
      if (data.type === "category") {
        this.selectCategory(data);
      } else if (data.type === "task") {
        this.selectTask(data);
      } else {
        this.selectLicense(data);
      }
    },
  },
  mounted() {
    const datasets_tmpl = document.getElementById("datasets-square");

    // this.isSigned = datasets_tmpl.getAttribute("data-issigned");
    this.isSigned = true;
  },
};
</script>
<style scoped>
.datasetSearch{
  padding: 20px 0 30px;
  background-color: #f5f5f6;
}

.datasetSearch .datasetSearchInput input{
  background-color: rgba(0, 0, 0, 0);
}
.mg-b-1 {
  margin-bottom: 1rem;
}

.mg-b-2 {
  margin-bottom: 2rem;
}

.mg-l-1 {
  margin-left: 1rem;
}

.text-gray-400 {
  --tw-text-opacity: 1;
  color: rgba(156, 163, 175, var(--tw-text-opacity));
}

.text-sm {
  font-size: 0.875rem;
  line-height: 1.25rem;
}

.underline {
  text-decoration: underline;
}

.flex {
  display: flex;
}

.font-medium {
  font-weight: 500;
}

.flex-wrap {
  flex-wrap: wrap;
}

.tag {
  background-image: linear-gradient(to bottom, var(--tw-gradient-stops));
  border-color: transparent;
  border-radius: 0.5rem;
  border-width: 1px;
  font-size: 0.875rem;
  line-height: 1.25rem;
  overflow: hidden;
  text-overflow: ellipsis;
  white-space: nowrap;
}

.tag-red {
  --tw-gradient-from: #fef2f2;
  --tw-gradient-stops: var(--tw-gradient-from),
    var(--tw-gradient-to, hsla(0, 86%, 97%, 0));
  --tw-gradient-to: #fef2f2;
  --tw-text-opacity: 1;
  color: rgba(153, 27, 27, var(--tw-text-opacity));
}

.tag-purple {
  --tw-gradient-from: #f5f3ff;
  --tw-gradient-stops: var(--tw-gradient-from),
    var(--tw-gradient-to, rgba(245, 243, 255, 0));
  --tw-gradient-to: #f5f3ff;
  --tw-text-opacity: 1;
  color: rgba(91, 33, 182, var(--tw-text-opacity));
}

.tag-blue {
  --tw-gradient-from: #eff6ff;
  --tw-gradient-stops: var(--tw-gradient-from),
    var(--tw-gradient-to, rgba(239, 246, 255, 0));
  --tw-gradient-to: #eff6ff;
  --tw-text-opacity: 1;
  color: rgba(30, 64, 175, var(--tw-text-opacity));
}

.tag.inactive {
  filter: grayscale(100%);
  opacity: 0.5;
}

.tag.tag-active {
  background-color: #0366d6;
  color: #ffffff;
}

.tag-gray {
  background-color: #f8f9fa;
  color: #415058;
}

.tag {
  align-items: center;
  display: inline-flex;
  flex: none;
  height: 2rem;
  margin-bottom: 0.35rem;
  margin-right: 0.35rem;
  max-width: 100%;
}

.tag>span {
  padding: 0.75rem;
  font-size: 14px;
}
</style>
