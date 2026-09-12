<template>
  <div
    class="inline field margin_b_0"
    :class="{
      dataset_flex: confirmDatasetList && confirmFlag,
      required: required,
    }"
  >
    <label
      class="label-fix-width"
      style="font-weight: normal"
      >{{ i18n.dataset_label }}</label
    >
    <span
      class="dataset-train-span"
      v-if="confirmDatasetList && confirmFlag"
    >
      <input type="hidden" name="attachment" :value="confirmDatasetList" />
      <input
        type="hidden"
        name="dataset_name"
        :value="confirmDatasetNameList"
      />
      <div class="multi-dataset-box">
        <span
          v-for="(item, index) in confirmChecklist"
          :key="index"
          class="select-dataset-label"
          :title="item"
          style="padding: 0.2rem 0"
          >{{ item }};
        </span>
      </div>
    </span>
    <span v-else>
      <input
        type="text"
        class="disabled"
        style="width: 48.5%"
        :placeholder="i18n.dataset_select_placeholder"
        :required="required"
      />
    </span>

    <el-button
      type="text"
      @click="openDataset"
      icon="el-icon-plus"
      :class="
        confirmFlag && confirmDatasetList
          ? 'select-dataset-button'
          : 'select-dataset-button-color'
      "
      >{{ i18n.dataset_select }}
    </el-button>
    <el-dialog
      :title="i18n.dataset_select"
      :visible.sync="dialogVisible"
      :width="dialogWidth"
    >
      <div class="ui icon input dataset-search-vue">
        <i
          class="search icon"
          style="cursor: pointer; pointer-events: auto"
          @click="searchName"
        ></i>
        <input
          type="text"
          :placeholder="i18n.dataset_search_placeholder"
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
          <el-tabs v-model="activeName" @tab-click="handleClick">
            <!-- 当前项目的数据集 -->
            <el-tab-pane
              :label="i18n.dataset_current_repo"
              name="first"
              v-loading="loadingCurrent"
            >
              <el-row>
                <el-tree
                  :data="currentDatasetList"
                  ref="currentTree"
                  highlight-current
                  show-checkbox
                  node-key="id"
                  :default-expanded-keys="initCurrentTreeNode"
                  :props="defaultProps"
                  :index="20"
                  accordion
                  @check="onCheck"
                >
                  <span slot-scope="{ node, data }" class="slot-wrap">
                    <span v-if="data.parent" class="custom-tree-node">
                      <el-tooltip v-if="data.Description" placement="top-start">
                        <div slot="content" class="multiple-wrap">
                          {{ data.Description }}
                        </div>
                        <span class="dataset-title dataset-nowrap"
                          ><div class="dataset_flex">
                            <span
                              style="flex: inherit"
                              class="dataset-nowrap"
                              >{{ node.label }}</span
                            >
                            <img
                              v-if="data.Recommend"
                              style="margin-left: 0.4rem"
                              src="/img/jian.svg"
                            />
                          </div>
                        </span>
                      </el-tooltip>
                      <span v-else class="dataset-title dataset-nowrap"
                        ><div class="dataset_flex">
                          <span style="flex: inherit" class="dataset-nowrap">{{
                            node.label
                          }}</span>
                          <img
                            v-if="data.Recommend"
                            style="margin-left: 0.4rem"
                            src="/img/jian.svg"
                          />
                        </div>
                      </span>
                      <span
                        class="dataset-repolink dataset-nowrap"
                        @click.stop="return false;"
                      >
                        <i
                          class="ri-links-line"
                          style="color: #21ba45; margin-right: 0.3rem"
                          :title="i18n.dataset_relate"
                          v-if="
                            '/' + data.Repo.OwnerName + '/' + data.Repo.Name !==
                            repoLink
                          "
                        ></i>
                        <a
                          :href="
                            '/' +
                            data.Repo.OwnerName +
                            '/' +
                            data.Repo.Name +
                            '/datasets'
                          "
                          target="_blank"
                          :title="data.Repo.OwnerName + '/' + data.Repo.Alias"
                          >{{ data.Repo.OwnerName }}/{{ data.Repo.Alias }}</a
                        >
                      </span>
                    </span>
                    <span v-else style="display: flex">
                      <span class="dataset-nowrap" :title="node.label">
                        {{ node.label }}
                      </span>
                      <span
                        class="zip-loading"
                        v-if="data.DecompressState === 2"
                      >
                        {{ i18n.dataset_unziping }}
                      </span>
                      <span
                        class="unzip-failed"
                        v-if="data.DecompressState === 3"
                      >
                        {{ i18n.dataset_unzip_failed }}
                      </span>
                      <span
                        class="unzip-failed"
                        v-if="data.Size > exceedSize"
                      >
                        {{ i18n.dataset_exceeds_failed }}{{exceedSize/(1024*1024*1024)}}G
                      </span>
                    </span>
                  </span>
                </el-tree>
                <div></div>
              </el-row>
              <div
                class="ui container"
                style="margin-top: 25px; text-align: center"
              >
                <el-pagination
                  background
                  @current-change="currentChangePage"
                  :current-page="initCurrentPage"
                  :page-size="5"
                  layout="total, prev, pager, next"
                  :total="totalNumCurrent"
                >
                </el-pagination>
              </div>
            </el-tab-pane>
            <!-- 我上传的数据集 -->
            <el-tab-pane
              :label="i18n.dataset_my_upload"
              name="second"
              v-loading="loadingMy"
            >
              <el-row>
                <el-tree
                  :data="myDatasetList"
                  ref="myTree"
                  highlight-current
                  show-checkbox
                  node-key="id"
                  :default-expanded-keys="initMyTreeNode"
                  :props="defaultProps"
                  :index="20"
                  accordion
                  @check="onCheck"
                >
                  <span slot-scope="{ node, data }" class="slot-wrap">
                    <span v-if="data.parent" class="custom-tree-node">
                      <el-tooltip v-if="data.Description" placement="top-start">
                        <div slot="content" class="multiple-wrap">
                          {{ data.Description }}
                        </div>
                        <span class="dataset-title dataset-nowrap"
                          ><div class="dataset_flex">
                            <span
                              style="flex: inherit"
                              class="dataset-nowrap"
                              >{{ node.label }}</span
                            >
                            <img
                              v-if="data.Recommend"
                              style="margin-left: 0.4rem"
                              src="/img/jian.svg"
                            />
                          </div>
                        </span>
                      </el-tooltip>
                      <span v-else class="dataset-title dataset-nowrap"
                        ><div class="dataset_flex">
                          <span style="flex: inherit" class="dataset-nowrap">{{
                            node.label
                          }}</span>
                          <img
                            v-if="data.Recommend"
                            style="margin-left: 0.4rem"
                            src="/img/jian.svg"
                          />
                        </div>
                      </span>
                      <span
                        class="dataset-repolink dataset-nowrap"
                        @click.stop="return false;"
                      >
                        <a
                          :href="
                            '/' +
                            data.Repo.OwnerName +
                            '/' +
                            data.Repo.Name +
                            '/datasets'
                          "
                          target="_blank"
                          :title="data.Repo.OwnerName + '/' + data.Repo.Alias"
                          >{{ data.Repo.OwnerName }}/{{ data.Repo.Alias }}</a
                        >
                      </span>
                    </span>
                    <span v-else style="display: flex">
                      <span class="dataset-nowrap" :title="node.label">
                        {{ node.label }}
                      </span>
                      <span
                        class="zip-loading"
                        v-if="data.DecompressState === 2"
                      >
                        {{ i18n.dataset_unziping }}
                      </span>
                      <span
                        class="unzip-failed"
                        v-if="data.DecompressState === 3"
                      >
                        {{ i18n.dataset_unzip_failed }}
                      </span>
                      <span
                        class="unzip-failed"
                        v-if="data.Size > exceedSize"
                      >
                      {{ i18n.dataset_exceeds_failed }}{{exceedSize/(1024*1024*1024)}}G
                      </span>
                    </span>
                  </span>
                </el-tree>
                <div></div>
              </el-row>
              <div
                class="ui container"
                style="margin-top: 25px; text-align: center"
              >
                <el-pagination
                  background
                  @current-change="myChangePage"
                  :current-page="initMyPage"
                  :page-size="5"
                  layout="total, prev, pager, next"
                  :total="totalNumMy"
                >
                </el-pagination>
              </div>
            </el-tab-pane>
            <!-- 公开的数据集 -->
            <el-tab-pane
              :label="i18n.dataset_public"
              name="third"
              v-loading="loadingPublic"
            >
              <el-row>
                <el-tree
                  :data="publicDatasetList"
                  ref="publicTree"
                  highlight-current
                  show-checkbox
                  node-key="id"
                  :default-expanded-keys="initPublicTreeNode"
                  :props="defaultProps"
                  :index="20"
                  accordion
                  @check="onCheck"
                >
                  <span slot-scope="{ node, data }" class="slot-wrap">
                    <span v-if="data.parent" class="custom-tree-node">
                      <el-tooltip v-if="data.Description" placement="top-start">
                        <div slot="content" class="multiple-wrap">
                          {{ data.Description }}
                        </div>
                        <span class="dataset-title dataset-nowrap"
                          ><div class="dataset_flex">
                            <span
                              style="flex: inherit"
                              class="dataset-nowrap"
                              >{{ node.label }}</span
                            >
                            <img
                              v-if="data.Recommend"
                              style="margin-left: 0.4rem"
                              src="/img/jian.svg"
                            />
                          </div>
                        </span>
                      </el-tooltip>
                      <span v-else class="dataset-title dataset-nowrap"
                        ><div class="dataset_flex">
                          <span style="flex: inherit" class="dataset-nowrap">{{
                            node.label
                          }}</span>
                          <img
                            v-if="data.Recommend"
                            style="margin-left: 0.4rem"
                            src="/img/jian.svg"
                          />
                        </div>
                      </span>
                      <span
                        class="dataset-repolink dataset-nowrap"
                        @click.stop="return false;"
                      >
                        <a
                          :href="
                            '/' +
                            data.Repo.OwnerName +
                            '/' +
                            data.Repo.Name +
                            '/datasets'
                          "
                          target="_blank"
                          :title="data.Repo.OwnerName + '/' + data.Repo.Alias"
                          >{{ data.Repo.OwnerName }}/{{ data.Repo.Alias }}</a
                        >
                      </span>
                    </span>
                    <span v-else style="display: flex">
                      <span class="dataset-nowrap" :title="node.label">
                        {{ node.label }}
                      </span>
                      <span
                        class="zip-loading"
                        v-if="data.DecompressState === 2"
                      >
                        {{ i18n.dataset_unziping }}
                      </span>
                      <span
                        class="unzip-failed"
                        v-if="data.DecompressState === 3"
                      >
                        {{ i18n.dataset_unzip_failed }}
                      </span>
                      <span
                        class="unzip-failed"
                        v-if="data.Size > exceedSize"
                      >
                      {{ i18n.dataset_exceeds_failed }}{{exceedSize/(1024*1024*1024)}}G
                      </span>
                    </span>
                  </span>
                </el-tree>
                <div></div>
              </el-row>
              <div
                class="ui container"
                style="margin-top: 25px; text-align: center"
              >
                <el-pagination
                  background
                  @current-change="publicChangePage"
                  :current-page="initPublicPage"
                  :page-size="5"
                  layout="total, prev, pager, next"
                  :total="totalNumPublic"
                >
                </el-pagination>
              </div>
            </el-tab-pane>
            <!-- 我点赞的数据集 -->
            <el-tab-pane
              :label="i18n.dataset_collected"
              name="four"
              v-loading="loadingFavorite"
            >
              <el-row>
                <el-tree
                  :data="MyFavoriteDatasetList"
                  ref="favoriteTree"
                  highlight-current
                  show-checkbox
                  node-key="id"
                  :default-expanded-keys="initFavoriteTreeNode"
                  :props="defaultProps"
                  :index="20"
                  accordion
                  @check="onCheck"
                >
                  <span slot-scope="{ node, data }" class="slot-wrap">
                    <span v-if="data.parent" class="custom-tree-node">
                      <el-tooltip v-if="data.Description" placement="top-start">
                        <div slot="content" class="multiple-wrap">
                          {{ data.Description }}
                        </div>
                        <span class="dataset-title dataset-nowrap"
                          ><div class="dataset_flex">
                            <span
                              style="flex: inherit"
                              class="dataset-nowrap"
                              >{{ node.label }}</span
                            >
                            <img
                              v-if="data.Recommend"
                              style="margin-left: 0.4rem"
                              src="/img/jian.svg"
                            />
                          </div>
                        </span>
                      </el-tooltip>
                      <span v-else class="dataset-title dataset-nowrap"
                        ><div class="dataset_flex">
                          <span style="flex: inherit" class="dataset-nowrap">{{
                            node.label
                          }}</span>
                          <img
                            v-if="data.Recommend"
                            style="margin-left: 0.4rem"
                            src="/img/jian.svg"
                          />
                        </div>
                      </span>
                      <span
                        class="dataset-repolink dataset-nowrap"
                        @click.stop="return false;"
                      >
                        <a
                          :href="
                            '/' +
                            data.Repo.OwnerName +
                            '/' +
                            data.Repo.Name +
                            '/datasets'
                          "
                          target="_blank"
                          :title="data.Repo.OwnerName + '/' + data.Repo.Alias"
                          >{{ data.Repo.OwnerName }}/{{ data.Repo.Alias }}</a
                        >
                      </span>
                    </span>
                    <span v-else style="display: flex">
                      <span class="dataset-nowrap" :title="node.label">
                        {{ node.label }}
                      </span>
                      <span
                        class="zip-loading"
                        v-if="data.DecompressState === 2"
                      >
                        {{ i18n.dataset_unziping }}
                      </span>
                      <span
                        class="unzip-failed"
                        v-if="data.DecompressState === 3"
                      >
                        {{ i18n.dataset_unzip_failed }}
                      </span>
                      <span
                        class="unzip-failed"
                        v-if="data.Size > exceedSize"
                      >
                      {{ i18n.dataset_exceeds_failed }}{{exceedSize/(1024*1024*1024)}}G
                      </span>
                    </span>
                  </span>
                </el-tree>
                <div></div>
              </el-row>
              <div
                class="ui container"
                style="margin-top: 25px; text-align: center"
              >
                <el-pagination
                  background
                  @current-change="favoriteChangePage"
                  :current-page="initFavoritePage"
                  :page-size="5"
                  layout="total, prev, pager, next"
                  :total="totalNumFavorite"
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
            {{ i18n.dataset_selected }}
          </div>
          <div style="flex: 1; margin: 1.5rem 0; overflow-y: auto">
            <el-checkbox-group v-model="checkList">
              <el-checkbox
                v-for="(item, index) in selectDatasetArray"
                :key="index"
                :label="item.label"
                :title="item.label"
                @change="(checked) => changeCheckbox(checked, item)"
              ></el-checkbox>
            </el-checkbox-group>
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
              >{{ i18n.dataset_ok }}</el-button
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
      defaultProps: {
        children: "Attachments",
        label: "label",
      },
      dialogWidth: "65%",
      dialogVisible: false,
      imageAddress: "",
      activeName: "first",
      search: "",
      required: true,
      i18n: {},
      type: 0,
      repoLink: "",
      selectDatasetArray: [],
      checkList: [],
      confirmChecklist: [],
      confirmDatasetList: "",
      confirmDatasetNameList: "",
      confirmFlag: false,

      saveStatusList: [],
      // 初始化已选择的数据集列表
      hasSelectDatasetList: [],
      //当前项目数据集页面配置的初始化
      initCurrentPage: 1,
      totalNumCurrent: 0,
      paramsCurrent: { page: 1, q: "" },
      currentDatasetList: [],
      loadingCurrent: false,
      initCurrentTreeNode: [],
      //我上传的数据集页面配置的初始化
      initMyPage: 1,
      totalNumMy: 0,
      paramsMy: { page: 1, q: "" },
      myDatasetList: [],
      loadingMy: false,
      initMyTreeNode: [],
      //公开的数据集页面配置的初始化
      initPublicPage: 1,
      totalNumPublic: 0,
      paramsPublics: { page: 1, q: "" },
      publicDatasetList: [],
      loadingPublic: false,
      initPublicTreeNode: [],
      //我点赞的数据集页面配置的初始化
      initFavoritePage: 1,
      totalNumFavorite: 0,
      MyFavoriteDatasetList: [],
      paramsFavorite: { page: 1, q: "" },
      loadingFavorite: false,
      initFavoriteTreeNode: [],
      exceedSize:0,
    };
  },
  methods: {
    openDataset() {
      this.dialogVisible = true;
      if (!this.confirmDatasetList) {
        this.confirmFlag = false;
      }

      this.getCurrentRepoDataset();
    },
    handleClick(tab, event) {
      this.search = "";
      if (tab.name == "first") {
        // this.paramsPublic.q = "";
        this.getCurrentRepoDataset();
      }
      if (tab.name == "second") {
        this.getMyUploadDataset();
      }
      if (tab.name == "third") {
        this.getPublicDataset();
      }
      if (tab.name == "four") {
        this.getMyFavoriteDataset();
      }
    },

    //tree 勾选触发事件
    onCheck(data, checkedInfo) {
      this.hasSelectDatasetList = [];
      if (
        this.selectDatasetArray.length === 0 ||
        this.selectDatasetArray.every((item) => item.id !== data.id)
      ) {
        if (
          this.selectDatasetArray.some((item) => {
            let itemIndex = item.label.lastIndexOf('.')
            let dataIndex = data.label.lastIndexOf('.')
            return item.label.substr(0,itemIndex) === data.label.substr(0,dataIndex);
          })
        ) {
          this.$refs[data.ref].setChecked(data.id, false, false);
          this.$message.warning(this.i18n.dataset_not_equal_file);
        } else if (this.selectDatasetArray.length === 5) {
          this.$refs[data.ref].setChecked(data.id, false, false);
          this.$message.error(this.i18n.dataset_most);
        } else {
          this.selectDatasetArray.push(data);
        }
      } else {
        let index = this.selectDatasetArray.findIndex((item) => {
          return item.id === data.id;
        });
        this.selectDatasetArray.splice(index, 1);
      }
      this.checkList = this.selectDatasetArray.map((item) => {
        return item.label;
      });
      this.saveStatusList = this.selectDatasetArray.map((item) => {
        return item.id;
      });
    },
    //已选择数据集checkbox group 勾选事件
    changeCheckbox(checked, data) {
      this.$refs.currentTree.setChecked(data.id, false, false);
      this.$refs.myTree.setChecked(data.id, false, false);
      this.$refs.publicTree.setChecked(data.id, false, false);
      this.$refs.favoriteTree.setChecked(data.id, false, false);
      let index = this.selectDatasetArray.findIndex((item) => {
        return item.id === data.id;
      });
      this.selectDatasetArray.splice(index, 1);
      this.saveStatusList.splice(index, 1);
    },
    tableHeaderStyle({ row, column, rowIndex, columnIndex }) {
      if (rowIndex === 0) {
        return "background:#f5f5f6;color:#606266";
      }
    },

    currentChangePage(val) {
      this.paramsCurrent.page = val;
      this.getCurrentRepoDataset();
    },
    myChangePage(val) {
      this.paramsMy.page = val;
      this.getMyUploadDataset();
    },
    publicChangePage(val) {
      this.paramsPublics.page = val;
      this.getPublicDataset();
    },
    favoriteChangePage(val) {
      this.paramsFavorite.page = val;
      this.getMyFavoriteDataset();
    },

    getCurrentRepoDataset() {
      this.loadingCurrent = true;
      let url = this.repoLink + "/datasets/current_repo_m";
      this.paramsCurrent.type = this.type;
      this.$axios
        .get(url, {
          params: this.paramsCurrent,
        })
        .then((res) => {
          this.loadingCurrent = false;
          let data = JSON.parse(res.data.data);
          this.currentDatasetList = this.transformeTreeData(
            data,
            "currentTree",
            this.paramsCurrent.page
          );
          this.initCurrentTreeNode = this.currentDatasetList[0]?.id
            ? [this.currentDatasetList[0].id]
            : [];
          this.totalNumCurrent = parseInt(res.data.count);
          let setCheckedKeysList = this.currentDatasetList.reduce(
            (pre, cur) => {
              cur.Attachments.forEach((item) => {
                if (
                  this.saveStatusList.includes(item.id) ||
                  this.hasSelectDatasetList.includes(item.id)
                ) {
                  pre.push(item.id);
                }
              });
              return pre;
            },
            []
          );
          this.$refs.currentTree.setCheckedKeys(setCheckedKeysList);
        })
        .catch((error) => {
          this.loadingCurrent = false;
          console.log(error);
        });
    },
    getMyUploadDataset() {
      this.loadingMy = true;
      let url = this.repoLink + "/datasets/my_datasets_m";
      this.paramsMy.type = this.type;
      this.$axios
        .get(url, {
          params: this.paramsMy,
        })
        .then((res) => {
          this.loadingMy = false;
          let data = JSON.parse(res.data.data);
          this.myDatasetList = this.transformeTreeData(
            data,
            "myTree",
            this.paramsMy.page
          );
          this.initMyTreeNode = this.myDatasetList[0]?.id
            ? [this.myDatasetList[0].id]
            : [];
          this.totalNumMy = parseInt(res.data.count);
          let setCheckedKeysList = this.myDatasetList.reduce((pre, cur) => {
            cur.Attachments.forEach((item) => {
              if (this.saveStatusList.includes(item.id)) {
                pre.push(item.id);
              }
            });
            return pre;
          }, []);
          this.$refs.myTree.setCheckedKeys(setCheckedKeysList);
        })
        .catch((error) => {
          console.log(error);
        });
    },

    getPublicDataset() {
      this.loadingPublic = true;
      let url = this.repoLink + "/datasets/public_datasets_m";
      this.paramsPublics.type = this.type;
      this.$axios
        .get(url, {
          params: this.paramsPublics,
        })
        .then((res) => {
          this.loadingPublic = false;
          let data = JSON.parse(res.data.data);
          this.publicDatasetList = this.transformeTreeData(
            data,
            "publicTree",
            this.paramsPublics.page
          );
          this.initPublicTreeNode = this.publicDatasetList[0]?.id
            ? [this.publicDatasetList[0].id]
            : [];
          this.totalNumPublic = parseInt(res.data.count);
          let setCheckedKeysList = this.publicDatasetList.reduce((pre, cur) => {
            cur.Attachments.forEach((item) => {
              if (this.saveStatusList.includes(item.id)) {
                pre.push(item.id);
              }
            });
            return pre;
          }, []);
          this.$refs.publicTree.setCheckedKeys(setCheckedKeysList);
        })
        .catch((error) => {
          this.loadingPublic = false;
          console.log(error);
        });
    },

    getMyFavoriteDataset() {
      this.loadingFavorite = true;
      let url = this.repoLink + "/datasets/my_favorite_m";
      this.paramsFavorite.type = this.type;
      this.$axios
        .get(url, {
          params: this.paramsFavorite,
        })
        .then((res) => {
          this.loadingFavorite = false;
          let data = JSON.parse(res.data.data);
          this.MyFavoriteDatasetList = this.transformeTreeData(
            data,
            "favoriteTree",
            this.paramsFavorite.page
          );
          this.initFavoriteTreeNode = this.MyFavoriteDatasetList[0]?.id
            ? [this.MyFavoriteDatasetList[0].id]
            : [];
          this.totalNumFavorite = parseInt(res.data.count);
          let setCheckedKeysList = this.MyFavoriteDatasetList.reduce(
            (pre, cur) => {
              cur.Attachments.forEach((item) => {
                if (this.saveStatusList.includes(item.id)) {
                  pre.push(item.id);
                }
              });
              return pre;
            },
            []
          );
          this.$refs.favoriteTree.setCheckedKeys(setCheckedKeysList);
        })
        .catch((error) => {
          this.loadingFavorite = false;
          console.log(error);
        });
    },
    transformeTreeData(data, ref, page) {
      return data.reduce((preParent, curParent) => {
        curParent.id = curParent.ID;
        curParent.disabled = true;
        curParent.parent = true;
        curParent.label = curParent.Title;
        let childrenData = curParent.Attachments.reduce(
          (preChild, curchild) => {
            curchild.id = curchild.UUID;
            if (curchild.DecompressState !== 1) {
              curchild.disabled = true;
            }
            if(curchild.Size>this.exceedSize && this.exceedSize){
              curchild.disabled = true;
            }
            curchild.ref = ref;
            curchild.label = curchild.Name;
            preChild.push(curchild);
            return preChild;
          },
          []
        );

        preParent.push(curParent);
        return preParent;
      }, []);
    },
    searchName() {
      if (this.activeName == "first") {
        this.paramsCurrent.q = this.search;
        this.paramsCurrent.page = 1;
        this.getCurrentRepoDataset();
      }
      if (this.activeName == "second") {
        this.paramsMy.q = this.search;
        this.paramsMy.page = 1;
        this.getMyUploadDataset();
      }
      if (this.activeName == "third") {
        this.paramsPublics.q = this.search;
        this.paramsPublics.page = 1;
        this.getPublicDataset();
      }
      if (this.activeName == "four") {
        this.paramsFavorite.q = this.search;
        this.paramsFavorite.page = 1;
        this.getMyFavoriteDataset();
      }
      return false;
    },
    confirmDataset() {
      this.confirmDatasetList = this.saveStatusList.join(";");
      this.confirmDatasetNameList = this.checkList.join(";");
      this.confirmChecklist = this.checkList;
      this.dialogVisible = false;
      this.confirmFlag = true;
    },
    setDialogWidth() {
      const cWidth = document.body.clientWidth;
      if (cWidth > 1600) {
        this.dialogWidth = "1200px";
      } else if (cWidth < 1600 && 1200 < cWidth) {
        this.dialogWidth = "1127px";
      } else if (992 < cWidth && cWidth < 1200) {
        this.dialogWidth = "993px";
      } else {
        this.dialogWidth = "723px";
      }
    },
  },
  watch: {
    search(val) {
      if (!val) {
        switch (this.activeName) {
          case "first": {
            this.paramsCurrent.q = val;
            this.getCurrentRepoDataset();
            break;
          }
          case "second": {
            this.paramsMy.q = val;
            this.getMyUploadDataset();
            break;
          }
          case "third": {
            this.paramsPublics.q = val;
            this.getPublicDataset();
            break;
          }
          case "four": {
            this.paramsFavorite.q = val;
            this.getMyFavoriteDataset();
            break;
          }
          default:
            console.log("error");
        }
      }
    },
  },
  mounted() {
    this.type = $(".cloudbrain-type").data("cloudbrain-type");
    this.repoLink = $(".cloudbrain-type").data("repo-link");
    this.exceedSize = $(".cloudbrain-type").data("exceed-size");
    if ($(".cloudbrain-type").data("dataset-uuid")) {
      this.hasSelectDatasetList = $(".cloudbrain-type")
        .data("dataset-uuid")
        .split(";");
      let hasSelectDatasetName = $(".cloudbrain-type")
        .data("dataset-name")
        .split(";");
      if (
        this.hasSelectDatasetList.length !== 0 &&
        hasSelectDatasetName[0] !== ""
      ) {
        this.saveStatusList = this.hasSelectDatasetList;
        this.checkList = hasSelectDatasetName;
        this.hasSelectDatasetList.forEach((item, index) => {
          this.selectDatasetArray.push({
            id: item,
            label: hasSelectDatasetName[index],
          });
        });
      }
      this.confirmDataset();
    }
    if (
      location.href.indexOf("modelarts/notebook/create") !== -1 ||
      location.href.indexOf("/cloudbrain/create") !== -1 || location.href.indexOf('grampus/notebook/create') !== -1
    ) {
      this.required = false;
    }
    window.onresize = () => {
      return (() => {
        this.setDialogWidth();
      })();
    };
  },
  created() {
    this.setDialogWidth();
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
/deep/ .el-dialog__body {
  position: relative;
}
.el-tree {
  max-height: 400px;
  overflow-y: auto;
  overflow-x: hidden;
  position: relative;
  cursor: default;
  background: #fff;
  color: #606266;
  font-family: SourceHanSansSC-medium;
}
.custom-tree-node {
  display: flex;
  align-items: center;
  justify-content: space-between;
}
.custom-tree-node .dataset-title {
  font-size: 14px;
  color: #101010;
  font-weight: 600;
  flex: 1;
}
.custom-tree-node .dataset-repolink {
  flex: 1;
  text-align: right;
  font-size: 12px;
}
.el-tree /deep/ .el-tree-node__content {
  height: 40px;
  background-color: #f5f5f6;
}
.el-tree /deep/ .el-tree-node__children .el-tree-node__content {
  height: 30px;
  background-color: #fff;
  line-height: 20px;
  font-size: 12px;
}
/deep/ .el-checkbox-group .el-checkbox {
  max-width: 100%;
  min-width: 80%;
}
/deep/ .el-checkbox-group .el-checkbox .el-checkbox__label {
  max-width: 100%;
  overflow: hidden;
  vertical-align: middle;
  text-overflow: ellipsis;
}
.dataset-nowrap {
  overflow: hidden;
  text-overflow: ellipsis;
}
.slot-wrap {
  flex: 1;
  padding-right: 2rem;
  max-width: 93%;
}
.multiple-wrap {
  -webkit-line-clamp: 3;
  -webkit-box-orient: vertical;
  display: -webkit-box;
  max-width: 400px;
  overflow: hidden;
}
.unzip-failed {
  margin-left: 1rem;
  color: red;
}
.zip-loading {
  margin-left: 1rem;
  color: #fcca00;
}
.dataset-search-vue {
  z-index: 9999;
  position: absolute;
  right: 31%;
  height: 30px;
  top: 6px;
}
.select-dataset-label {
  max-width: 100%;
  overflow: hidden;
  text-overflow: ellipsis;
  margin-right: 1rem;
  white-space: nowrap;
}
.dataset_flex {
  display: flex;
  align-items: center;
}
.dataset-debug-span {
  display: inline-block;
  width: 50%;
  margin-left: 0.3rem;
}
.dataset-train-span {
  display: inline-block;
  width: 48.5%;
  margin-left: 0.3rem;
}
.margin_b_0 {
  margin-bottom: 0 !important;
}
.select-dataset-button {
  margin-left: 0.3rem;
  color: #0366d6;
}
.select-dataset-button-color {
  color: #0366d6;
}
.multi-dataset-box {
  border: 1px solid rgba(34, 36, 38, 0.15);
  padding: 0.678571em 1em;
  border-radius: 0.285714rem;
  background: rgb(255, 255, 255);
  display: flex;
  flex-direction: column;
}

@media screen and (min-width: 1200px) and (max-width: 1600px) {
  .dataset-search-vue {
    top: -36px;
  }
}
@media screen and (min-width: 1200px) and (max-width: 1400px) {
  .multiple-wrap {
    max-width: 200px;
  }
}
</style>
