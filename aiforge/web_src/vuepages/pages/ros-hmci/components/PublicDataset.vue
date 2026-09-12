<template>
  <div v-loading="loading">
    <div class="dataset_head_wrap">
      <el-checkbox
        v-model="checked"
        style="padding: 0.5rem 1rem"
        @change="handleCheckedChange"
        >{{ $t("datasets.platform_recommendations") }}</el-checkbox
      >
      <el-dropdown trigger="click" style="cursor: pointer">
        <span class="el-dropdown-link">
          {{ $t("datasets.sort") }}<i class="el-icon-caret-bottom el-icon--right"></i>
        </span>
        <el-dropdown-menu slot="dropdown">
          <el-dropdown-item
            :class="{ active: item.active }"
            v-for="item in sortList"
            :key="item.name"
            @click.native="handleSort(item)"
            >{{ $t("datasets." + item.name) }}</el-dropdown-item
          >
        </el-dropdown-menu>
      </el-dropdown>
    </div>

    <div v-if="!showEmpty" class="ui two cards">
      <div
        class="ui card dataset_card_wrap"
        v-for="(item, index) in publicDataList"
        @click="gotoDataset(item)"
      >
        <div class="content" style="border-bottom: none">
          <div class="dataset_content_wrap">
            <span :title="item.Title" class="nowrap" style="display: inline-block">{{
              item.Title
            }}</span>
            <img v-if="item.Recommend" src="/img/jian.svg" style="margin-left: 0.5rem" />
            <span class="dataset_icon_wrap" @click.stop="postSquareStar(item, index)">
              <div class="dataset_icon_content">
                <svg
                  width="1.4em"
                  height="1.4em"
                  viewBox="0 0 32 32"
                  class="heart-stroke"
                  :class="{ stars_active: item.IsStaring }"
                >
                  <path
                    d="M4.4 6.54c-1.761 1.643-2.6 3.793-2.36 6.056.24 2.263 1.507 4.521 3.663 6.534a29110.9 29110.9 0 0010.296 9.633l10.297-9.633c2.157-2.013 3.424-4.273 3.664-6.536.24-2.264-.599-4.412-2.36-6.056-1.73-1.613-3.84-2.29-6.097-1.955-1.689.25-3.454 1.078-5.105 2.394l-.4.319-.398-.319c-1.649-1.316-3.414-2.143-5.105-2.394a7.612 7.612 0 00-1.113-.081c-1.838 0-3.541.694-4.983 2.038z"
                  ></path>
                </svg>
              </div>
              <span style="line-height: 1; color: #101010">{{ item.NumStars }}</span>
            </span>
          </div>
          <div class="dataset_label_content">
            <span
              v-if="item.Category"
              class="ui repo-topic label topic"
              @click.stop="chooseLabel(item.Category, 'category')"
              >{{ $t("datasets." + item.Category) }}</span
            >
            <span
              v-if="item.Task"
              class="ui repo-topic label topic"
              @click.stop="chooseLabel(item.Task, 'task')"
              >{{ $t("datasets." + item.Task) }}</span
            >
            <span
              v-if="item.License"
              class="ui repo-topic label topic"
              @click.stop="chooseLabel(item.License, 'license')"
              >{{ item.License }}</span
            >
          </div>
          <div class="description dataset_desc_wrap">
            <p>{{ item.Description }}</p>
          </div>
        </div>
        <div class="extra content" style="border-top: none !important">
          <div style="display: flex; align-items: center">
            <a
              v-if="item.UserID === 0"
              :href="`/${item.Repo.OwnerName}`"
              :title="item.Repo.OwnerName"
            >
              <img
                class="ui avatar image"
                style="width: 22px; height: 22px"
                :src="`/user/avatar/${item.Repo.OwnerName}/-1`"
              />
            </a>

            <a v-else :href="`/${item.User.Name}`" :title="item.User.Name">
              <img
                class="ui avatar image"
                style="width: 22px; height: 22px"
                :src="`/user/avatar/${item.User.Name}/-1`"
              />
            </a>
            <span class="dataset_extra_time">{{ item.CreatedUnix | DateTransfer }}</span>
            <span class="dataset_extra_link" :title="$t('datasets.downloadtimes')">
              <i class="ri-link"></i>
              <span class="dataset_extra_content">{{ item.UseCount }}</span>
            </span>
            <span class="dataset_extra_download" :title="$t('datasets.citations')">
              <i class="ri-download-line"></i>
              <span class="dataset_extra_content">{{ item.DownloadTimes }}</span>
            </span>
          </div>
        </div>
      </div>
    </div>
    <el-empty
      v-else
      :description="$t('noDataset')"
      :image-size="100"
      style="
        background-color: rgba(245, 245, 246, 0.5);
        min-height: 400px;
        padding: 40px 0 80px 0;
      "
    ></el-empty>
    <div class="center" style="margin-top: 2rem" v-if="!showEmpty">
      <el-pagination
        background
        @size-change="handleSizeChange"
        @current-change="handleCurrentChange"
        :current-page="params.page"
        :page-sizes="[30]"
        :page-size="params.pageSize"
        layout="total, sizes, prev, pager, next, jumper"
        :total="total"
      >
      </el-pagination>
    </div>
  </div>
</template>

<script>
import { getDatasets, putDatasetStar } from "~/apis/modules/dataset";
import { Message } from "element-ui";
export default {
  props: {
    isSigned: {
      type: String,
      default: "false",
    },
    dataGet: {
      type: String,
      default: "",
    },
    searchValue: {
      type: String,
      default: "",
    },
    searchFlag: {
      type: Boolean,
      default: false,
    },
    categoryValue: {
      type: String,
      default: "",
    },
    taskValue: {
      type: String,
      default: "",
    },
    licenseValue: {
      type: String,
      default: "",
    },
  },
  data() {
    return {
      publicDataList: [],
      total: 0,
      params: {
        sort: "",
        q: "",
        page: 1,
        pageSize: 30,
        recommend: false,
        category: "",
        task: "ros_hmci_datasets",
        license: "",
      },
      checked: false,
      sortList: [
        { name: "default", active: true },
        { name: "latest", active: false },
        { name: "oldest", active: false },
        { name: "recentupdate", active: false },
        { name: "leastupdate", active: false },
        { name: "downloadtimes", active: false },
        { name: "moststars", active: false },
        { name: "mostusecount", active: false },
      ],
      showEmpty: false,
      loading: false,
    };
  },
  watch: {
    searchValue(newVal) {
      if (!newVal) {
        this.params.page = 1;
        this.params.q = newVal;
        this.getDataList(this.dataGet);
      }
    },
    searchFlag(newVal) {
      this.params.page = 1;
      this.params.q = this.searchValue;
      this.getDataList(this.dataGet);
    },
    categoryValue(val) {
      this.params.category = val;
      this.params.page = 1;
      this.getDataList(this.dataGet);
    },
    taskValue(val) {
      this.params.task = val;
      this.params.page = 1;
      this.getDataList(this.dataGet);
    },
    licenseValue(val) {
      this.params.license = val;
      this.params.page = 1;
      this.getDataList(this.dataGet);
    },
  },
  methods: {
    getDataList(dataType) {
      let url = `/explore/${dataType}`;
      this.loading = true;
      getDatasets(url, this.params)
        .then((res) => {
          if (res.data.result_code === "0") {
            if (res.data.data === "null" || res.data.data == "[]") {
              this.showEmpty = true;
              this.total = 0;
            } else {
              this.publicDataList = JSON.parse(res.data.data);
              if (dataType === "my_favorite_datasets") {
                this.publicDataList.forEach((ele, index) => {
                  this.publicDataList[index].IsStaring = true;
                });
              }
              this.total = Number(res.data.count);
              this.showEmpty = false;
            }
            this.loading = false;
          } else {
            Message.error(res.data.error_msg);
            this.loading = false;
          }
        })
        .catch((err) => {
          Message.error(err);
          this.loading = false;
        });
    },
    handleCurrentChange(val) {
      this.params.page = val;
      this.getDataList(this.dataGet);
    },
    handleSizeChange(val) {
      this.params.pageSize = val;
      this.getDataList(this.dataGet);
    },
    gotoDataset(item) {
      location.href = `/${item.Repo.OwnerName}/${item.Repo.Name}/datasets`;
    },
    postSquareStar(item, index) {
      if (this.isSigned === "false" || !this.isSigned || this.dataGet == "my_datasets")
        return;
      let baseUrl = `/${item.Repo.OwnerName}/${item.Repo.Name}/datasets/${item.ID}/`;
      let url = item.IsStaring ? baseUrl + "unstar" : baseUrl + "star";
      let changeItem = item;
      putDatasetStar(url)
        .then((res) => {
          if (res.data.Code === 0) {
            if (this.dataGet == "my_favorite_datasets") {
              this.getDataList("my_favorite_datasets");
              Message.success(this.$t("datasets.unstarSuccess"));
              return;
            }
            if (item.IsStaring) {
              changeItem.IsStaring = false;
              changeItem.NumStars = changeItem.NumStars - 1;
              this.$set(this.publicDataList, index, changeItem);
              this.$set(this.publicDataList, index, changeItem);
              Message.success(this.$t("datasets.unstarSuccess"));
            } else {
              changeItem.IsStaring = true;
              changeItem.NumStars = changeItem.NumStars + 1;
              this.$set(this.publicDataList, index, changeItem);
              this.$set(this.publicDataList, index, changeItem);
              Message.success(this.$t("datasets.starSuccess"));
            }
          } else {
            Message.error(res.data.Message);
          }
        })
        .catch((err) => {
          Message.error(err);
        });
    },
    handleCheckedChange(val) {
      this.params.recommend = val;
      this.getDataList(this.dataGet);
    },
    handleSort(item) {
      this.sortList.forEach((element) => {
        element.active = false;
      });
      item.active = true;
      this.params.sort = item.name;
      this.getDataList(this.dataGet);
    },
    chooseLabel(item, type) {
      const data = { name: item, active: false, type: type };
      this.$emit("getLabel", data);
    },
  },

  filters: {
    DateTransfer(unix) {
      let date = new Date(unix * 1000);
      return date.toISOString().slice(0, 10);
    },
  },
  mounted() {
    this.getDataList(this.dataGet);
  },
};
</script>

<style scoped lang="less">
.el-tabs__item {
  font-size: 16px;
}

.el-dropdown-link {
  color: rgba(0, 0, 0, 0.87);
  font-weight: 400;
  font-size: 1rem;
}

.el-icon-arrow-down {
  font-size: 1rem;
}

.el-dropdown-menu__item {
  padding: 2px 1rem;
  font-size: 1rem;
}

.el-dropdown-menu__item.active {
  color: #409eff;
  background-color: rgba(179, 216, 255, 0.3);
}

.dataset_card_wrap {
  cursor: pointer;
  box-shadow: 0px 4px 4px 0px rgba(232, 232, 232, 0.6) !important;
  border: 1px solid rgba(232, 232, 232, 1) !important;
}

.dataset_content_wrap {
  font-size: 16px;
  color: #0366d6;
  font-family: SourceHanSansSC-medium;
  height: 34px;
  font-weight: bold;
  display: flex;
  align-items: center;
}

.dataset_icon_wrap {
  display: flex;
  align-items: center;
  justify-content: flex-end;
  cursor: pointer;
  font-size: 12px;
  font-weight: normal;
  flex: 1;
  margin-left: 1.5rem;
}

.dataset_icon_content {
  line-height: 1;
  margin-right: 4px;
  margin-bottom: -2px;
}

.dataset_label_content {
  font-size: 12px;
  margin-top: 5px;
}

.dataset_desc_wrap {
  -webkit-box-orient: vertical;
  -webkit-line-clamp: 2;
  display: -webkit-box;
  overflow: hidden;
  color: #999999 !important;
  font-size: 14px;
  margin-top: 10px;
}

.dataset_extra_time {
  color: #999999;
  font-size: 12px;
  margin-left: 0.5rem;
}

.dataset_extra_link {
  display: flex;
  align-items: center;
  justify-content: center;
  margin: 0 1rem;
}

.dataset_extra_download {
  display: flex;
  align-items: center;
  justify-content: center;
}

.dataset_extra_content {
  color: #101010;
  font-size: 12px;
  margin-left: 0.2rem;
}

.dataset_head_wrap {
  display: flex;
  align-items: center;
  justify-content: space-between;
  margin-bottom: 0.5rem;
}

.heart-stroke {
  stroke: #fa8c16;
  stroke-width: 2;
  fill: #fff;
}

.stars_active {
  fill: #fa8c16 !important;
  stroke: #fa8c16 !important;
}
</style>
