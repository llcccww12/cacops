<template>
  <div>
    <template v-if="showFlag">
      <div v-loading="loadingLinkPage">
        <div class="ui container">
          <div class="ui mobile reversed stackable grid">
            <div class="row" style="justify-content: space-between">
              <div class="ui blue small menu compact selectcloudbrain">
                <a class="item" :href="`${repoLink}/datasets`">{{
                  i18n.current_dataset
                }}</a>
                <a
                  class="active item"
                  :href="`${repoLink}/datasets/reference_datasets`"
                  >{{ i18n.linked_datasets }}</a
                >
              </div>
              <!-- <button
                style="margin-right: 2rem"
                class="ui green button"
                :class="{ disabled: !canWrite }"
                @click="openDataset()"
              >
                {{ i18n.linked_datasets }}
              </button> -->
            </div>
            <div class="row">
              <div class="ui two cards" style="width: 100%">
                <div
                  class="ui card refer-dataset-card"
                  v-for="(item, index) in datasetList"
                  :key="index"
                  @click="gotoDataset(item)"
                >
                  <div class="content" style="border-bottom: none">
                    <div class="refer-dataset-card-content">
                      <div class="refer-dataset-card-title">
                        <span
                          :title="item.Title"
                          class="nowrap"
                          style="display: inline-block; max-width: 90%"
                          >{{ item.Title }}</span
                        ><img
                          v-if="item.Recommend"
                          src="/img/jian.svg"
                          style="margin-left: 0.5rem"
                        />
                      </div>
                      <template v-if="item.IsStaring">
                        <div style="display: flex">
                          <button
                            class="ui mini basic button dataset-card-flavor"
                            @click.stop="postStar(item, isSigned)"
                          >
                            <i class="ri-heart-fill" style="color: #fa8c16"></i>
                            <span style="margin-left: 0.3rem">{{
                              i18n.unfavorite
                            }}</span>
                          </button>
                          <a class="ui mini basic button card-flavor-num">
                            {{ item.NumStars }}
                          </a>
                        </div>
                      </template>
                      <template v-else>
                        <div style="display: flex">
                          <button
                            class="ui mini basic button dataset-card-flavor"
                            @click.stop="postStar(item, isSigned)"
                          >
                            <i class="ri-heart-line"></i>
                            <span style="margin-left: 0.3rem">{{
                              i18n.favorite
                            }}</span>
                          </button>
                          <a class="ui mini basic button card-flavor-num">
                            {{ item.NumStars }}
                          </a>
                        </div>
                      </template>
                    </div>
                    <div style="font-size: 12px; margin-top: 5px">
                      <a
                        v-if="item.Category"
                        :href="'/explore/datasets?category=' + item.Category"
                        class="ui repo-topic label topic"
                        @click.stop
                        >{{ i18n[item.Category] || item.Category }}</a
                      >
                      <a
                        v-if="item.Task"
                        :href="'/explore/datasets?task=' + item.Task"
                        class="ui repo-topic label topic"
                        @click.stop
                        >{{ i18n[item.Task] || item.Task }}</a
                      >
                      <a
                        v-if="item.License"
                        :href="'/explore/datasets?license=' + item.License"
                        class="ui repo-topic label topic"
                        @click.stop
                        >{{ item.License }}</a
                      >
                    </div>
                    <div class="description card-flavor-desc">
                      <p>{{ item.Description }}</p>
                    </div>
                  </div>
                  <div
                    class="extra content"
                    style="border-top: none !important"
                  >
                    <div style="display: flex; align-items: center">
                      <a
                        :href="'/' + item.Repo.OwnerName"
                        :title="item.Repo.OwnerName"
                        @click.stop
                      >
                        <img
                          class="ui avatar image"
                          style="width: 22px; height: 22px"
                          :src="'/user/avatar/' + item.Repo.OwnerName + '/-1'"
                        />
                      </a>
                      <span
                        style="
                          color: #999999;
                          font-size: 12px;
                          margin-left: 0.5rem;
                        "
                        >{{ item.CreatedUnix | transformTimestamp }}</span
                      >
                      <span
                        style="
                          display: flex;
                          align-items: center;
                          justify-content: center;
                          margin: 0 1rem;
                        "
                        :title="i18n.citations"
                      >
                        <i class="ri-link"></i>
                        <span
                          style="
                            color: #101010;
                            font-size: 12px;
                            margin-left: 0.2rem;
                          "
                          >{{ item.UseCount }}</span
                        >
                      </span>
                      <span
                        style="display: flex; align-items: center; flex: 1"
                        :title="i18n.downloads"
                      >
                        <i class="ri-download-line"></i>
                        <span
                          style="
                            color: #101010;
                            font-size: 12px;
                            margin-left: 0.2rem;
                          "
                          >{{ item.DownloadTimes }}</span
                        >
                      </span>
                      <button
                        class="ui mini button"
                        :class="{ disabled1: !canWrite }"
                        @click.stop="
                          cancelReferData(item.ID, item.Title, canWrite)
                        "
                      >
                        {{ i18n.disassociate }}
                      </button>
                    </div>
                  </div>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>
    </template>
    <template v-else>
      <div class="ui container">
        <div class="ui mobile reversed stackable grid">
          <div class="row" style="justify-content: space-between">
            <div class="ui blue small menu compact selectcloudbrain">
              <a class="item" :href="`${repoLink}/datasets`">{{
                i18n.current_dataset
              }}</a>
              <a
                class="active item"
                :href="`${repoLink}/datasets/reference_datasets`"
                >{{ i18n.linked_datasets }}</a
              >
            </div>
            <!-- <button class="ui green button" @click="openDataset()">
              {{ i18n.linked_datasets }}
            </button> -->
          </div>
        </div>
        <div class="ui placeholder segment bgtask-none">
          <div class="ui icon header bgtask-header-pic"></div>
          <div class="bgtask-content-header">{{ i18n.not_link_dataset }}</div>
          <div class="bgtask-content">
            <div class="bgtask-content-txt">
              {{ i18n.no_link_dataset_tips1 }}
            </div>
            <div class="bgtask-content-txt">
              {{ i18n.dataset_instructions_for_use
              }}<a href="https://openi.pcl.ac.cn/zeizei/OpenI_Learning">{{
                i18n.dataset_camp_course
              }}</a>
            </div>
          </div>
        </div>
      </div>
    </template>

    <el-dialog
      :visible.sync="dialogVisible"
      :width="dialogWidth"
      @closed="refreshData"
    >
      <template #title>
        <span>{{ i18n.linked_datasets }}</span>
        <span class="refer-data-tips"><span>*</span>{{ i18n.linked_datasets_tips }}</span>
      </template>
      <div class="ui icon input dataset-search-vue">
        <i
          class="search icon"
          style="cursor: pointer; pointer-events: auto"
          @click="searchName"
        ></i>
        <input
          type="text"
          :placeholder="i18n.search_dataset"
          v-model="search"
          @keydown.enter.stop.prevent="searchName"
        />
      </div>
      <el-row>
        <el-col
          :span="17"
          style="
            padding-right: 1rem;
            border-right: 1px solid #f5f5f6;
            position: relative;
          "
        >
          <el-tabs v-model="activeName">
            <el-tab-pane :label="i18n.public_dataset" name="first">
              <el-row v-loading="loadingPublicPage" style="min-height: 50px">
                <el-checkbox-group
                  v-model="checkList"
                  style="font-size: 14px; line-height: 1"
                >
                  <div
                    v-for="(item, index) in publicDatasetList"
                    :key="index"
                    class="select-data-wrap"
                  >
                    <div class="dataset-header-vue">
                      <el-checkbox
                        :label="item.ID"
                        @change="(checked) => changeCheckbox(checked, item)"
                        :title="item.Title"
                        class="select-data-title"
                        style="display: flex; align-items: end"
                        ><span class="ref-data-title">
                          <span
                            style="overflow: hidden; text-overflow: ellipsis"
                          >
                            {{ item.Title }}
                          </span>
                          <img
                            v-if="item.Recommend"
                            src="/img/jian.svg"
                            style="margin-left: 0.5rem"
                          /> </span
                      ></el-checkbox>
                      <a
                        class="select-data-title select-data-href"
                        :href="`/${item.Repo.OwnerName}/${item.Repo.Name}/datasets`"
                        :title="`${item.Repo.OwnerName}/${item.Repo.Alias}`"
                        style="font-size: 12px"
                        target="_blank"
                        >{{ item.Repo.OwnerName }}/{{ item.Repo.Alias }}</a
                      >
                    </div>
                    <div class="data-multiple-wrap" :title="item.Description">
                      {{ item.Description }}
                    </div>
                  </div>
                </el-checkbox-group>
              </el-row>
              <div
                class="ui container"
                style="margin-top: 25px; text-align: center"
              >
                <el-pagination
                  background
                  @current-change="currentChange"
                  :current-page="currentPage"
                  :page-size="5"
                  layout="total, prev, pager, next"
                  :total="totalNum"
                >
                </el-pagination>
              </div>
            </el-tab-pane>
          </el-tabs>
        </el-col>
        <el-col
          :span="7"
          style="
            display: flex;
            flex-direction: column;
            height: 100%;
            right: 0;
            position: absolute;
            padding: 0 1.5rem;
          "
        >
          <div
            style="
              font-size: 14px;
              height: 40px;
              text-align: left;
              color: #0066ff;
              line-height: 40px;
            "
          >
            {{ i18n.selected_data_file }}
          </div>
          <div
            style="
              flex: 1;
              margin-top: 1.5rem;
              margin-bottom: 1rem;
              overflow-y: auto;
            "
          >
            <el-checkbox
              v-for="(item, index) in selectDatasetArray"
              :key="index"
              :label="item.ID"
              :title="item.Title"
              :value="item.isChecked"
              @change="(checked) => changeCheckSelected(checked, item)"
              style="display: flex; margin: 0.5rem 0"
              ><span class="select-data-right">{{
                item.Title
              }}</span></el-checkbox
            >
          </div>
          <div style="text-align: end">
            <el-button
              @click.native="confirmDataset"
              size="small"
              style="
                background: #389e0d;
                color: #fff;
                border: 1px solid #389e0d;
              "
              >{{ i18n.sure }}</el-button
            >
          </div>
        </el-col>
      </el-row>
    </el-dialog>
  </div>
</template>

<script>
const { _AppSubUrl, _StaticUrlPrefix, csrf } = window.config;

export default {
  components: {},
  data() {
    return {
      dialogWidth: "65%",
      dialogVisible: false,
      activeName: "first",
      repoLink: "",
      datasetList: [],
      test: false,
      checkList: [],
      publicDatasetList: [],
      showFlag: true,
      search: "",
      selectDatasetArray: [],
      paramsPublics: { page: 1, q: "" },

      i18n: {},

      totalNum: 0,
      currentPage: 1,
      loadingLinkPage: false,
      loadingPublicPage: false,
      canWrite: true,
      isSigned: false,
      maxReferenceNum: 20,
      isCurrentUrl: true,
    };
  },
  methods: {
    openDataset() {
      this.checkList = this.datasetList.map((item) => {
        this.selectDatasetArray.push({
          ID: item.ID,
          Title: item.Title,
          isChecked: true,
        });
        return item.ID;
      });
      this.dialogVisible = true;
      this.getDatasetList();
    },
    refreshData() {
      this.checkList = [];
      this.selectDatasetArray = [];
    },
    gotoDataset(item) {
      window.open(`/${item.Repo.OwnerName}/${item.Repo.Name}/datasets`);
    },
    currentChange(page) {
      this.paramsPublics.page = page;
      this.getDatasetList();
    },
    searchName() {
      this.paramsPublics.q = this.search;
      this.paramsPublics.page = 1;
      this.getDatasetList();
    },
    cancelReferData(id, name, canWrite) {
      if (!canWrite) {
        return;
      }
      let url = `${this.repoLink}/datasets/reference_datasets/${id}`;
      this.$axios
        .delete(url)
        .then((res) => {
          if (res.data.Code === 0) {
            this.$message.success(this.i18n.cancel_link_dataset.format(name));
            let index = this.datasetList.findIndex((item) => {
              return item.ID === id;
            });
            this.datasetList.splice(index, 1);
            if (this.datasetList.length === 0) {
              this.showFlag = false;
            }
          } else {
            this.$message.error(res.data.Message);
          }
        })
        .catch((err) => {
          this.$message.error(this.i18n.dataset_link_failed);
          console.log(err);
        });
    },
    confirmDataset() {
      this.submitReferDataset();
      this.dialogVisible = false;
    },
    changeCheckbox(checked, item) {
      if (this.checkList.length > this.maxReferenceNum) {
        this.checkList.pop();
        this.$message.error(
          this.i18n.dataset_over_nums.format(this.maxReferenceNum)
        );
        return;
      }
      if (this.checkList.length === this.maxReferenceNum && !checked) {
        this.$message.error(
          this.i18n.dataset_over_nums.format(this.maxReferenceNum)
        );
        return;
      }
      if (checked) {
        this.selectDatasetArray.push({
          ID: item.ID,
          Title: item.Title,
          isChecked: true,
        });
      } else {
        let index = this.selectDatasetArray.findIndex((element) => {
          return element.ID === item.ID;
        });
        this.selectDatasetArray.splice(index, 1);
      }
    },
    changeCheckSelected(checked, item) {
      let index = this.selectDatasetArray.findIndex((element) => {
        return element.ID === item.ID;
      });
      this.selectDatasetArray.splice(index, 1);
      this.checkList.splice(index, 1);
    },
    postStar(item, isSigned) {
      if (!isSigned) {
        return;
      }
      if (item.IsStaring) {
        let url = `${this.repoLink}/datasets/${item.ID}/unstar`;
        this.$axios.put(url).then((res) => {
          if (res.data.Code === 0) {
            this.datasetList.forEach((element, i) => {
              if (element.ID === item.ID) {
                this.datasetList[i].NumStars -= 1;
                this.datasetList[i].IsStaring = !this.datasetList[i].IsStaring;
              }
            });
          }
        });
      } else {
        let url = `${this.repoLink}/datasets/${item.ID}/star`;
        this.$axios.put(url).then((res) => {
          if (res.data.Code === 0) {
            this.datasetList.forEach((element, i) => {
              if (element.ID === item.ID) {
                this.datasetList[i].NumStars += 1;
                this.datasetList[i].IsStaring = !this.datasetList[i].IsStaring;
              }
            });
          }
        });
      }
    },

    getSelectDatasetList() {
      this.loadingLinkPage = true;
      let url = `${this.repoLink}/datasets/reference_datasets_data`;
      this.$axios.get(url).then((res) => {
        this.loadingLinkPage = false;
        if (!res.data) {
          this.showFlag = false;
          this.datasetList = [];
          return;
        } else {
          this.datasetList = res.data;
          this.datasetList.length
            ? (this.showFlag = true)
            : (this.showFlag = false);
        }
      });
    },
    getDatasetList() {
      this.loadingPublicPage = true;
      let url = `${this.repoLink}/datasets/reference_datasets_available`;
      this.$axios
        .get(url, {
          params: this.paramsPublics,
        })
        .then((res) => {
          this.publicDatasetList = JSON.parse(res.data.data);
          this.totalNum = parseInt(res.data.count);
          this.loadingPublicPage = false;
        });
    },
    submitReferDataset() {
      if (this.checkList.length > this.maxReferenceNum) {
        this.$message.error(
          this.i18n.dataset_over_nums.format(this.maxReferenceNum)
        );
        return;
      }
      let url = `${this.repoLink}/datasets/reference_datasets`;
      let data = this.qs.stringify(
        {
          _csrf: csrf,
          dataset_id: this.checkList,
        },
        { arrayFormat: "repeat" }
      );
      this.$axios
        .post(url, data)
        .then((res) => {
          if (res.data.Code === 0) {
            this.$message.success(this.i18n.dataset_link_success);
            this.getSelectDatasetList();
          } else {
            this.$message.error(res.data.Message);
          }
        })
        .catch((err) => {
          this.$message.error(this.i18n.dataset_link_failed);
          console.log(err);
        });
    },
  },

  filters: {
    transformTimestamp(timestamp) {
      const date = new Date(parseInt(timestamp) * 1000);
      const Y = date.getFullYear() + "-";
      const M =
        (date.getMonth() + 1 < 10
          ? "0" + (date.getMonth() + 1)
          : date.getMonth() + 1) + "-";
      const D =
        (date.getDate() < 10 ? "0" + date.getDate() : date.getDate()) + "  ";

      const dateString = Y + M + D;
      return dateString;
    },
  },
  watch: {
    search(val) {
      if (!val) {
        this.searchName();
      }
    },
    showFlag(val) {
      if (!val || this.isCurrentUrl) {
        document
          .getElementById("header-dataset")
          .setAttribute("href", this.repoLink + "/datasets");
      } else {
        document
          .getElementById("header-dataset")
          .setAttribute("href", this.repoLink + "/datasets/reference_datasets");
      }
    },
  },
  mounted() {
    this.getSelectDatasetList();
  },
  created() {
    this.repoLink = $(".reference-dataset").data("repolink") || "";
    this.canWrite = $(".reference-dataset").data("canwrite");
    this.isSigned = $(".reference-dataset").data("is-sign");
    this.maxReferenceNum = $(".reference-dataset").data("max-reference-num");
    this.isCurrentUrl = $(".reference-dataset").data("address");
    if (document.documentElement.attributes["lang"].nodeValue == "en-US") {
      this.i18n = this.$locale.US;
    } else {
      this.i18n = this.$locale.CN;
    }
  },
  beforeDestroy() {},
};
</script>

<style scoped>
.dataset-search-vue {
  z-index: 9999;
  position: absolute;
  right: 31%;
  height: 30px;
  top: 60px;
}
.refer-dataset-card {
  cursor: pointer;
  box-shadow: 0px 4px 4px 0px rgba(232, 232, 232, 0.6);
  border: 1px solid rgba(232, 232, 232, 1);
}
.refer-dataset-card .refer-dataset-card-content {
  font-size: 16px;
  color: #0366d6;
  font-family: SourceHanSansSC-medium;
  height: 34px;
  font-weight: bold;
  display: flex;
  align-items: center;
  justify-content: space-between;
}
.refer-dataset-card-title {
  display: flex;
  align-items: center;
  max-width: 80%;
  width: 100%;
}
.dataset-card-flavor {
  display: flex;
  align-items: center;
  padding: 0.3rem 0.5rem;
  border: #888888;
  border-top-right-radius: 0 !important;
  border-bottom-right-radius: 0 !important;
  margin-right: -1px;
}
.card-flavor-num {
  padding: 0.5rem;
  border: #888888;
  border-top-left-radius: 0 !important;
  border-bottom-left-radius: 0 !important;
  cursor: default !important;
}
.card-flavor-desc {
  -webkit-box-orient: vertical;
  -webkit-line-clamp: 2;
  display: -webkit-box;
  overflow: hidden;
  color: #999999;
  font-size: 14px;
  margin-top: 10px;
}
.select-data-wrap {
  padding: 1rem 0;
  border-bottom: 1px solid rgba(0, 0, 0, 0.05);
}
.select-data-title {
  flex: 1;
  overflow: hidden;
}
.select-data-title .ref-data-title {
  font-size: 18px;
  color: #454545;
  font-weight: 700;
  display: flex;
  align-items: baseline;
  max-width: 100%;
}
.select-data-href {
  text-align: right;
  text-overflow: ellipsis;
  max-width: 35%;
  word-break: initial;
  margin-left: 1rem;
  white-space: nowrap;
}
/deep/ .el-checkbox-group .el-checkbox .el-checkbox__label {
  display: flex;
  max-width: 90%;
}
.select-data-right {
  overflow: hidden;
  vertical-align: middle;
  text-overflow: ellipsis;
  max-width: 100%;
  display: inline-block;
}
.data-multiple-wrap {
  -webkit-line-clamp: 3;
  -webkit-box-orient: vertical;
  display: -webkit-box;
  max-width: 100%;
  overflow: hidden;
  padding-top: 1rem;
  color: #888888;
  font: 12px;
  line-height: 20px;
  margin-left: 2rem;
}
.disabled1 {
  opacity: 0.45 !important;
}
.refer-data-tips {
  margin-left: 2rem;
    display: inline-block;
    color: rgb(242, 113, 28);
    font-size: 12px;
}
</style>
