<template>
  <div
    class="inline required field"
    :class="{ unite: benchmarkNew, min_title: benchmarkNew }"
  >
    <label
      v-if="benchmarkNew"
      class="label-fix-width"
      style="font-weight: normal"
      >{{ i18n.image_label }}</label
    >
    <label v-else>{{ i18n.image_label }}</label>
    <input
      v-if="benchmarkNew"
      type="text"
      name="image"
      :value="imageAddress"
      style="width: 48.5%"
      :placeholder="i18n.image_select_placeholder"
      required
    />
    <input
      v-else
      type="text"
      name="image"
      :value="imageAddress"
      :placeholder="i18n.image_select_placeholder"
      required
    />
    <el-button
      type="text"
      @click="dialogVisible = true"
      icon="el-icon-plus"
      style="color: #0366d6"
      >{{ i18n.image_select }}
    </el-button>
    <el-dialog
      :title="i18n.image_select"
      :visible.sync="dialogVisible"
      width="50%"
    >
      <div
        class="ui icon input"
        style="z-index: 9999; position: absolute; right: 50px; height: 30px"
      >
        <i
          class="search icon"
          style="cursor: pointer; pointer-events: auto"
        ></i>
        <input
          type="text"
          :placeholder="i18n.image_search_placeholder"
          v-model="search"
        />
      </div>
      <el-tabs v-model="activeName" @tab-click="handleClick">
        <el-tab-pane
          :label="i18n.cloudeBrainMirror.recommendImages"
          name="first"
          v-loading="loadingPublic"
        >
          <div
            style="
              display: flex;
              align-items: center;
              justify-content: space-between;
              padding: 1rem 0;
              border-bottom: 1px solid #f5f5f5;
            "
            v-for="(publicData, index) in tableDataPublic"
            :key="index"
          >
            <div style="width: 90%">
              <div
                style="
                  display: flex;
                  align-items: center;
                  justify-content: space-between;
                "
              >
                <div style="display: flex; align-items: center">
                  <span
                    class="panel_dataset_name text-over"
                    style="margin-left: 0"
                    >{{ publicData.tag }}
                  </span>
                  <img
                    v-if="publicData.type == 5"
                    src="/img/jian.svg"
                    style="margin-left: 0.5rem"
                  />
                </div>

                <div v-if="!!publicData.topics" class="text-over">
                  <span
                    v-for="(topic, index) in publicData.topics"
                    class="ui repo-topic label topic"
                    >{{ topic }}</span
                  >
                </div>
              </div>
              <div style="margin-top: 8px; display: flex">
                <a
                  v-if="publicData.relAvatarLink || publicData.userName"
                  :title="publicData.userName"
                  style="cursor: default"
                >
                  <img
                    class="ui avatar mini image"
                    style="width: 20px; height: 20px"
                    :src="publicData.relAvatarLink"
                  />
                </a>
                <a v-else
                  ><img
                    class="ui avatar mini image"
                    title="Ghost"
                    src="/user/avatar/ghost/-1"
                    style="width: 20px; height: 20px"
                /></a>
                <span class="panel_datset_desc">{{
                  publicData.description
                }}</span>
              </div>
            </div>
            <div>
              <button
                class="ui primary basic button mini"
                @click.stop.prevent="
                  selectImages(publicData.place, publicData.tag)
                "
              >
                {{ i18n.image_use }}
              </button>
            </div>
          </div>
          <div
            class="ui container"
            style="margin-top: 50px; text-align: center"
          >
            <el-pagination
              background
              @current-change="handleCurrentChangePublic"
              :current-page="currentPagePublic"
              :page-size="pageSizePublic"
              layout="total, prev, pager, next"
              :total="totalNumPublic"
            >
            </el-pagination>
          </div>
        </el-tab-pane>
      </el-tabs>
    </el-dialog>
  </div>
</template>

<script>
const { _AppSubUrl, _StaticUrlPrefix, csrf } = window.config;

export default {
  components: {},
  data() {
    return {
      dialogVisible: false,
      benchmarkNew: false,
      imageAddress: "",
      activeName: "first",
      search: "",
      checked: false,
      currentPagePublic: 1,
      pageSizePublic: 5,
      totalNumPublic: 0,
      paramsPublic: {
        page: 1,
        pageSize: 5,
        q: "",
        recommend: false,
        cloudbrainType: 2,
        _csrf: csrf,
      },
      tableDataPublic: [],
      loadingPublic: false,
      i18n: {},
    };
  },
  methods: {
    handleClick(tab, event) {
      this.search = "";
      if (tab.name == "first") {
        this.paramsPublic.q = "";
        this.getImageListPublic();
      }
    },
    tableHeaderStyle({ row, column, rowIndex, columnIndex }) {
      if (rowIndex === 0) {
        return "background:#f5f5f6;color:#606266";
      }
    },

    handleCurrentChangePublic(val) {
      this.paramsPublic.page = val;
      this.getImageListPublic();
    },

    getImageListPublic() {
      this.loadingPublic = true;
      this.$axios
        .get("/api/v1/images/recommend", {
          params: this.paramsPublic,
        })
        .then((res) => {
          this.totalNumPublic = res.data.count;
          this.tableDataPublic = res.data.images;
          this.loadingPublic = false;
        });
    },

    searchName() {
      if (this.activeName == "first") {
        this.paramsPublic.q = this.search;
        this.paramsPublic.page = 1;
        this.getImageListPublic();
      }
    },

    selectImages(place) {
      this.imageAddress = place;
      this.dialogVisible = false;
    },
  },
  watch: {
    search(val) {
      if (this.activeName == "first") {
        this.paramsPublic.q = val;
        this.getImageListPublic();
      }
    },
  },
  mounted() {
    if (document.getElementById("ai_image_name")) {
      this.imageAddress = document.getElementById("ai_image_name").value;
    }
    this.getImageListPublic();
    if (
      location.href.indexOf("benchmark") !== -1 ||
      location.href.indexOf("train-job") !== -1 ||
      location.href.indexOf("modelsafety") !== -1
    ) {
      this.benchmarkNew = true;
    }
  },
  created() {
    if (document.documentElement.attributes["lang"].nodeValue == "en-US") {
      this.i18n = this.$locale.US;
    } else {
      this.i18n = this.$locale.CN;
    }
  },
};
</script>

<style scoped>
.header-wrapper {
  background-color: #f5f5f6;
  padding-top: 15px;
}

.image_text {
  padding: 25px 0 55px 0;
}

#header {
  position: relative;
  top: -40px;
}

#success {
  background-color: #5bb973;
  color: white;
}

.text-over {
  overflow: hidden;
  text-overflow: ellipsis;
  vertical-align: middle;
  white-space: nowrap;
}

.image_title {
  display: inline-block;
  width: 80%;
  cursor: default;
  color: rgb(66, 98, 144);
}

.image_desc {
  -webkit-line-clamp: 2;
  -webkit-box-orient: vertical;
  display: -webkit-box;
  text-overflow: ellipsis;
  overflow: hidden;
}

.heart-stroke {
  stroke: #666;
  stroke-width: 2;
  fill: #fff;
}

.stars_active {
  fill: #fa8c16 !important;
  stroke: #fa8c16 !important;
}
</style>
