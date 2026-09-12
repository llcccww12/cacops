<template>
  <div>
    <div class="header-wrapper">
      <div class="ui container">
        <el-row class="image_text">
          <h1>{{ i18n['cloudeBrainMirror']['cloud_brain_mirror'] }}</h1>
        </el-row>
      </div>
    </div>
    <div class="ui container" id="header">
      <el-tabs v-model="activeName" @tab-click="handleClick">
        <el-tab-pane :label="i18n['cloudeBrainMirror']['recommendImages']" name="first" v-loading="loadingRecommend">
          <template v-if="tableDataRecommend.length !== 0">
            <el-row style="align-items: center; display: flex">
              <el-col :span="12">
                <div>
                  <!-- <el-checkbox v-model="checked">{{i18n['cloudeBrainMirror']['platform_recommendations']}}</el-checkbox> -->
                </div>
              </el-col>
              <el-col :span="10">
                <div>
                  <el-input :placeholder="i18n['cloudeBrainMirror']['placeholder']" v-model="search"
                    class="input-with-select" @keyup.enter.native="searchName()">
                    <el-button id="success" slot="append" icon="el-icon-search" @click="searchName()">{{
                      i18n['cloudeBrainMirror']['search'] }}</el-button>
                  </el-input>
                </div>
              </el-col>
              <el-col :span="2">
                <el-dropdown @command="handleCommandRecommend" trigger="click"
                  style="border: 1px solid rgba(34,36,38,.15);border-radius: 4px;padding: 0.5rem 1rem;float: right;">
                  <span class="el-dropdown-link">
                    {{ sortRecommend | transformSort(vm) }}<i class="el-icon-caret-bottom el-icon--right"></i>
                  </span>
                  <el-dropdown-menu slot="dropdown">
                    <el-dropdown-item :command="{ label: 'defaultsort', sort: '' }">{{
                      i18n['cloudeBrainMirror']['defaultsort'] }}</el-dropdown-item>
                    <el-dropdown-item :command="{ label: 'moststars', sort: 'moststars' }">{{
                      i18n['cloudeBrainMirror']['moststars'] }}</el-dropdown-item>
                    <el-dropdown-item :command="{ label: 'mostused', sort: 'mostused' }">{{
                      i18n['cloudeBrainMirror']['mostused'] }}</el-dropdown-item>
                    <el-dropdown-item :command="{ label: 'newest', sort: 'newest' }">{{
                      i18n['cloudeBrainMirror']['newest'] }}</el-dropdown-item>
                  </el-dropdown-menu>
                </el-dropdown>
              </el-col>
            </el-row>
            <el-row style="margin-top: 15px">
              <el-table :data="tableDataRecommend" style="width: 100%" :header-cell-style="tableHeaderStyle">
                <el-table-column :label="i18n['cloudeBrainMirror']['mirror_tag']" min-width="19%" align="left"
                  prop="tag">
                  <template slot-scope="scope">
                    <div style="display: flex; align-items: center">
                      <a class="text-over image_title" :title="scope.row.tag">{{
                        scope.row.tag
                      }}</a>
                      <img v-if="scope.row.type == 5" src="/img/jian.svg" style="margin-left: 0.5rem" />
                    </div>
                  </template>
                </el-table-column>
                <el-table-column :label="i18n['cloudeBrainMirror']['mirror_description']" min-width="26%" align="left"
                  prop="description">
                  <template slot-scope="scope">
                    <div class="image_desc" :title="scope.row.description">
                      {{ scope.row.description }}
                    </div>
                    <div v-if="!!scope.row.topics">
                      <span v-for="(topic, index) in scope.row.topics" class="ui repo-topic label topic" :key="index"
                        style="cursor: default">{{ topic }}</span>
                    </div>
                  </template>
                </el-table-column>
                <el-table-column prop="cloudbrainType" :label="i18n['cloudeBrainMirror']['available_clusters']"
                  min-width="15%" align="center">
                  <template slot-scope="scope">
                    {{ scope.row.cloudbrainType | transformType(vm) }}
                  </template>
                </el-table-column>
                <el-table-column prop="creator" :label="i18n['cloudeBrainMirror']['creator']" min-width="8%"
                  align="center">
                  <template slot-scope="scope">
                    <a v-if="scope.row.userName || scope.row.relAvatarLink" :href="'/' + scope.row.userName"
                      :title="scope.row.userName">
                      <img :src="scope.row.relAvatarLink" class="ui avatar image" />
                    </a>
                    <a v-else><img class="ui avatar image" title="Ghost" src="/user/avatar/ghost/-1" /></a>
                  </template>
                </el-table-column>
                <el-table-column prop="createdUnix" :label="i18n['cloudeBrainMirror']['creation_time']" align="center"
                  min-width="14%">
                  <template slot-scope="scope">
                    {{ scope.row.createdUnix | transformTimestamp }}
                  </template>
                </el-table-column>
                <el-table-column align="center" min-width="18%" :label="i18n['cloudeBrainMirror']['operation']">
                  <template slot-scope="scope">
                    <div style="
                        display: flex;
                        justify-content: flex-end;
                        align-items: center;
                      ">
                      <div style="display: flex;align-items: center;padding: 0 1rem;" :title="i18n['citations']">
                        <i class="ri-links-line" style="font-size: 16px;"></i>
                        <span style="line-height: 2;margin-left: 0.3rem;">{{ scope.row.useCount }}</span>
                      </div>
                      <div style="
                          display: flex;
                          align-items: center;
                          cursor: pointer;
                          padding: 0 1rem;
                        " @click="
                          imageStar(
                            scope.$index,
                            scope.row.id,
                            scope.row.isStar
                          )
                          ">
                        <svg width="1.4em" height="1.4em" viewBox="0 0 32 32" class="heart-stroke"
                          :class="{ stars_active: scope.row.isStar }">
                          <path
                            d="M4.4 6.54c-1.761 1.643-2.6 3.793-2.36 6.056.24 2.263 1.507 4.521 3.663 6.534a29110.9 29110.9 0 0010.296 9.633l10.297-9.633c2.157-2.013 3.424-4.273 3.664-6.536.24-2.264-.599-4.412-2.36-6.056-1.73-1.613-3.84-2.29-6.097-1.955-1.689.25-3.454 1.078-5.105 2.394l-.4.319-.398-.319c-1.649-1.316-3.414-2.143-5.105-2.394a7.612 7.612 0 00-1.113-.081c-1.838 0-3.541.694-4.983 2.038z">
                          </path>
                        </svg>
                        <span style="line-height: 2; margin-left: 0.3rem">{{
                          scope.row.numStars
                        }}</span>
                      </div>
                    </div>
                  </template>
                </el-table-column>
              </el-table>
            </el-row>
            <div class="ui container" style="margin-top: 50px; text-align: center">
              <el-pagination background @size-change="handleSizeChangeRecommend"
                @current-change="handleCurrentChangeRecommend" :current-page="currentPageRecommend"
                :page-size="pageSizeRecommend" :page-sizes="[5, 10, 15]" layout="total, sizes, prev, pager, next, jumper"
                :total="totalNumRecommend">
              </el-pagination>
            </div>
          </template>
          <template v-else>
            <el-row style="align-items: center; display: flex">
              <el-col :span="12">
                <div>
                  <!-- <el-checkbox v-model="checked">{{ i18n['cloudeBrainMirror']['platform_recommendations']
                  }}</el-checkbox> -->
                </div>
              </el-col>
              <el-col :span="10">
                <div>
                  <el-input :placeholder="i18n['cloudeBrainMirror']['placeholder']" v-model="search"
                    class="input-with-select" @keyup.enter.native="searchName()">
                    <el-button id="success" slot="append" icon="el-icon-search" @click="searchName()">{{
                      i18n['cloudeBrainMirror']['search'] }}</el-button>
                  </el-input>
                </div>
              </el-col>
              <el-col :span="2">
                <el-dropdown @command="handleCommandRecommend" trigger="click"
                  style="border: 1px solid rgba(34,36,38,.15);border-radius: 4px;padding: 0.5rem 1rem;float: right;">
                  <span class="el-dropdown-link">
                    {{ sortRecommend | transformSort(vm) }}<i class="el-icon-caret-bottom el-icon--right"></i>
                  </span>
                  <el-dropdown-menu slot="dropdown">
                    <el-dropdown-item :command="{ label: 'defaultsort', sort: '' }">{{
                      i18n['cloudeBrainMirror']['defaultsort'] }}</el-dropdown-item>
                    <el-dropdown-item :command="{ label: 'moststars', sort: 'moststars' }">{{
                      i18n['cloudeBrainMirror']['moststars'] }}</el-dropdown-item>
                    <el-dropdown-item :command="{ label: 'mostused', sort: 'mostused' }">{{
                      i18n['cloudeBrainMirror']['mostused'] }}</el-dropdown-item>
                    <el-dropdown-item :command="{ label: 'newest', sort: 'newest' }">{{
                      i18n['cloudeBrainMirror']['newest'] }}</el-dropdown-item>
                  </el-dropdown-menu>
                </el-dropdown>
              </el-col>
            </el-row>
            <el-empty :image-size="200"></el-empty>
          </template>
        </el-tab-pane>
        <el-tab-pane :label="i18n['cloudeBrainMirror']['my_mirror']" name="second" v-loading="loadingCustom">
          <template v-if="tableDataCustom.length !== 0">
            <el-row style="align-items: center; display: flex">
              <el-col :span="12">
                <div style="visibility: hidden">TODO</div>
              </el-col>
              <el-col :span="10">
                <div>
                  <el-input :placeholder="i18n['cloudeBrainMirror']['placeholder']" v-model="search"
                    class="input-with-select" @keyup.enter.native="searchName()">
                    <el-button id="success" slot="append" icon="el-icon-search" @click="searchName()">{{
                      i18n['cloudeBrainMirror']['search'] }}</el-button>
                  </el-input>
                </div>
              </el-col>
              <el-col :span="2">
                <el-dropdown @command="handleCommandCustom" trigger="click"
                  style="border: 1px solid rgba(34,36,38,.15);border-radius: 4px;padding: 0.5rem 1rem;float: right;">
                  <span class="el-dropdown-link">
                    {{ sortCustom | transformSort(vm) }}<i class="el-icon-caret-bottom el-icon--right"></i>
                  </span>
                  <el-dropdown-menu slot="dropdown">
                    <el-dropdown-item :command="{ label: 'moststars', sort: 'moststars' }">{{
                      i18n['cloudeBrainMirror']['moststars'] }}</el-dropdown-item>
                    <el-dropdown-item :command="{ label: 'mostused', sort: 'mostused' }">{{
                      i18n['cloudeBrainMirror']['mostused'] }}</el-dropdown-item>
                    <el-dropdown-item :command="{ label: 'newest', sort: 'newest' }">{{
                      i18n['cloudeBrainMirror']['newest'] }}</el-dropdown-item>
                  </el-dropdown-menu>
                </el-dropdown>
              </el-col>
            </el-row>
            <el-row style="margin-top: 15px">
              <el-table :data="tableDataCustom" style="width: 100%" :header-cell-style="tableHeaderStyle">
                <el-table-column :label="i18n['cloudeBrainMirror']['mirror_tag']" min-width="19%" align="left"
                  prop="tag">
                  <template slot-scope="scope">
                    <div style="display: flex; align-items: center">
                      <a class="text-over image_title" :title="scope.row.tag">{{
                        scope.row.tag
                      }}</a>&nbsp;&nbsp;&nbsp;
                      <!-- <i class="ri-lock-2-line" style="color: #fa8c16" v-if="scope.row.isPrivate"></i> -->
                      <img v-if="scope.row.type == 5" src="/img/jian.svg" style="margin-left: 0.5rem" />
                    </div>
                  </template>
                </el-table-column>
                <el-table-column :label="i18n['cloudeBrainMirror']['mirror_description']" min-width="25%" align="left"
                  prop="description">
                  <template slot-scope="scope">
                    <div class="image_desc" :title="scope.row.description">
                      {{ scope.row.description }}
                    </div>
                    <div v-if="!!scope.row.topics">
                      <span v-for="(topic, index) in scope.row.topics" class="ui repo-topic label topic" :key="index"
                        style="cursor: default">{{ topic }}</span>
                    </div>
                  </template>
                </el-table-column>
                <el-table-column prop="cloudbrainType" :label="i18n['cloudeBrainMirror']['available_clusters']"
                  min-width="12%" align="center">
                  <template slot-scope="scope">
                    {{ scope.row.cloudbrainType | transformType(vm) }}
                  </template>
                </el-table-column>
                <el-table-column prop="sumbimtState" :label="i18n['cloudeBrainMirror']['commit_status']" min-width="9%"
                  align="center">
                  <template slot-scope="scope">
                    <div style="
                        display: flex;
                        align-items: center;
                        justify-content: center;
                      ">
                      <span v-if="scope.row.status === 0">{{ i18n['cloudeBrainMirror']['commiting'] }}</span>
                      <span v-if="scope.row.status === 1">{{ i18n['cloudeBrainMirror']['commit_success'] }}</span>
                      <span v-if="scope.row.status === 2">{{ i18n['cloudeBrainMirror']['commit_failed'] }}</span>
                      <el-tooltip v-if="scope.row.status === 0" class="item" effect="dark"
                        :content="i18n['cloudeBrainMirror']['mirror_committed']" placement="top">
                        <i class="CREATING" style="margin-left: 0.3rem"></i>
                      </el-tooltip>
                      <el-tooltip v-if="scope.row.status === 1" class="item" effect="dark"
                        :content="i18n['cloudeBrainMirror']['mirror_submitted']" placement="top">
                        <i class="SUCCEEDED" style="margin-left: 0.3rem"></i>
                      </el-tooltip>
                      <el-tooltip v-if="scope.row.status === 2" class="item" effect="dark"
                        :content="i18n['cloudeBrainMirror']['check_exceeds_20g']" placement="top">
                        <i class="FAILED" style="margin-left: 0.3rem"></i>
                      </el-tooltip>
                    </div>
                  </template>
                </el-table-column>
                <el-table-column prop="sumbimtState" :label="i18n['cloudeBrainMirror']['recommend_by_plateform']"
                  width="140" align="center">
                  <template slot-scope="scope">
                    <div v-if="scope.row.status === 1" style="">
                      <div v-if="scope.row.apply_status === 2"
                        style="display: flex;align-items:center;justify-content:center;">
                        <span style="color: rgb(250, 140, 22);">{{ i18n['cloudeBrainMirror']['pending_approval']
                        }}</span>
                        <el-tooltip class="item" effect="dark" :content="i18n['cloudeBrainMirror']['pending_approval']"
                          placement="top">
                          <i class="CLOCK" style="margin-left: 0.3rem"></i>
                        </el-tooltip>
                      </div>
                      <div v-if="scope.row.apply_status === 3"
                        style="display: flex;align-items:center;justify-content:center;">
                        <span style="color: rgb(19, 194, 141);">{{ i18n['cloudeBrainMirror']['approved'] }}</span>
                        <el-tooltip class="item" effect="dark" :content="i18n['cloudeBrainMirror']['approved']"
                          placement="top">
                          <i class="SUCCEEDED" style="margin-left: 0.3rem"></i>
                        </el-tooltip>
                      </div>
                      <div v-if="scope.row.apply_status === 4"
                        style="display: flex;align-items:center;justify-content:center;">
                        <span style="color: red">{{ i18n['cloudeBrainMirror']['not_approved'] }}</span>
                        <el-tooltip class="item" effect="dark" :content="scope.row.message" placement="top">
                          <i class="FAILED" style="margin-left: 0.3rem"></i>
                        </el-tooltip>
                      </div>
                      <div
                        v-if="scope.row.apply_status == 0 || scope.row.apply_status == 1 || scope.row.apply_status === 4">
                        <span class="apply-btn" v-if="scope.row.type != 5" @click="applyImage(scope.row.id)">
                          {{ i18n['cloudeBrainMirror']['apply'] }}
                        </span>
                        <div v-else>--</div>
                      </div>
                    </div>
                    <div v-else>--</div>
                  </template>
                </el-table-column>
                <el-table-column prop="createdUnix" :label="i18n['cloudeBrainMirror']['creation_time']" align="center"
                  min-width="13%">
                  <template slot-scope="scope">
                    {{ scope.row.createdUnix | transformTimestamp }}
                  </template>
                </el-table-column>
                <el-table-column align="center" min-width="22%" :label="i18n['cloudeBrainMirror']['operation']">
                  <template slot-scope="scope">
                    <div style="
                        display: flex;
                        justify-content: flex-end;
                        align-items: center;
                      ">
                      <div style="display: flex;align-items: center;padding: 0 1rem;" :title="i18n['citations']">
                        <i class="ri-links-line" style="font-size: 16px;"></i>
                        <span style="line-height: 2;margin-left: 0.3rem;">{{ scope.row.useCount }}</span>
                      </div>
                      <div style="
                          display: flex;
                          align-items: center;
                          cursor: default;
                          padding: 0 1rem;
                        ">
                        <svg width="1.4em" height="1.4em" viewBox="0 0 32 32" class="heart-stroke">
                          <path
                            d="M4.4 6.54c-1.761 1.643-2.6 3.793-2.36 6.056.24 2.263 1.507 4.521 3.663 6.534a29110.9 29110.9 0 0010.296 9.633l10.297-9.633c2.157-2.013 3.424-4.273 3.664-6.536.24-2.264-.599-4.412-2.36-6.056-1.73-1.613-3.84-2.29-6.097-1.955-1.689.25-3.454 1.078-5.105 2.394l-.4.319-.398-.319c-1.649-1.316-3.414-2.143-5.105-2.394a7.612 7.612 0 00-1.113-.081c-1.838 0-3.541.694-4.983 2.038z">
                          </path>
                        </svg>
                        <span style="line-height: 2; margin-left: 0.3rem">{{
                          scope.row.numStars
                        }}</span>
                      </div>
                      <div style="padding-left: 1rem; cursor: pointer">
                        <el-dropdown size="medium">
                          <span class="el-dropdown-link">
                            {{ i18n['cloudeBrainMirror']['more'] }}<i class="el-icon-arrow-down el-icon--right"></i>
                          </span>
                          <el-dropdown-menu slot="dropdown">
                            <el-dropdown-item @click.native="eidtImage(scope.row.id)">{{
                              i18n['cloudeBrainMirror']['edit'] }}</el-dropdown-item>
                            <el-dropdown-item style="color: red" @click.native="deleteImage(scope.row.id)">{{
                              i18n['cloudeBrainMirror']['delete'] }}</el-dropdown-item>
                          </el-dropdown-menu>
                        </el-dropdown>
                      </div>
                    </div>
                  </template>
                </el-table-column>
              </el-table>
            </el-row>
            <div class="ui container" style="margin-top: 50px; text-align: center">
              <el-pagination background @size-change="handleSizeChangeCustom" @current-change="handleCurrentChangeCustom"
                :current-page="currentPageCustom" :page-size="pageSizeCustom" :page-sizes="[5, 10, 15]"
                layout="total, sizes, prev, pager, next, jumper" :total="totalNumCustom">
              </el-pagination>
            </div>
          </template>
          <template v-else>
            <el-row style="align-items: center; display: flex">
              <el-col :span="12">
                <div style="visibility: hidden"></div>
              </el-col>
              <el-col :span="10">
                <div>
                  <el-input :placeholder="i18n['cloudeBrainMirror']['placeholder']" v-model="search"
                    class="input-with-select" @keyup.enter.native="searchName()">
                    <el-button id="success" slot="append" icon="el-icon-search" @click="searchName()">{{
                      i18n['cloudeBrainMirror']['search'] }}</el-button>
                  </el-input>
                </div>
              </el-col>
              <el-col :span="2">
                <el-dropdown @command="handleCommandCustom" trigger="click"
                  style="border: 1px solid rgba(34,36,38,.15);border-radius: 4px;padding: 0.5rem 1rem;float: right;">
                  <span class="el-dropdown-link">
                    {{ sortCustom | transformSort(vm) }}<i class="el-icon-caret-bottom el-icon--right"></i>
                  </span>
                  <el-dropdown-menu slot="dropdown">
                    <el-dropdown-item :command="{ label: 'moststars', sort: 'moststars' }">{{
                      i18n['cloudeBrainMirror']['moststars'] }}</el-dropdown-item>
                    <el-dropdown-item :command="{ label: 'mostused', sort: 'mostused' }">{{
                      i18n['cloudeBrainMirror']['mostused'] }}</el-dropdown-item>
                    <el-dropdown-item :command="{ label: 'newest', sort: 'newest' }">{{
                      i18n['cloudeBrainMirror']['newest'] }}</el-dropdown-item>
                  </el-dropdown-menu>
                </el-dropdown>
              </el-col>
            </el-row>
            <el-empty :image-size="200"></el-empty>
          </template>
        </el-tab-pane>
        <el-tab-pane :label="i18n['cloudeBrainMirror']['my_favorite_mirror']" name="third">
          <template v-if="tableDataStar.length !== 0">
            <el-row style="align-items: center; display: flex">
              <el-col :span="12">
                <div style="visibility: hidden"></div>
              </el-col>
              <el-col :span="10">
                <div>
                  <el-input :placeholder="i18n['cloudeBrainMirror']['placeholder']" v-model="search"
                    class="input-with-select" @keyup.enter.native="searchName()">
                    <el-button id="success" slot="append" icon="el-icon-search" @click="searchName()">{{
                      i18n['cloudeBrainMirror']['search'] }}</el-button>
                  </el-input>
                </div>
              </el-col>
              <el-col :span="2">
                <el-dropdown @command="handleCommandStar" trigger="click"
                  style="border: 1px solid rgba(34,36,38,.15);border-radius: 4px;padding: 0.5rem 1rem;float: right;">
                  <span class="el-dropdown-link">
                    {{ sortStar | transformSort(vm) }}<i class="el-icon-caret-bottom el-icon--right"></i>
                  </span>
                  <el-dropdown-menu slot="dropdown">
                    <el-dropdown-item :command="{ label: 'defaultsort', sort: '' }">{{
                      i18n['cloudeBrainMirror']['defaultsort'] }}</el-dropdown-item>
                    <el-dropdown-item :command="{ label: 'moststars', sort: 'moststars' }">{{
                      i18n['cloudeBrainMirror']['moststars'] }}</el-dropdown-item>
                    <el-dropdown-item :command="{ label: 'mostused', sort: 'mostused' }">{{
                      i18n['cloudeBrainMirror']['mostused'] }}</el-dropdown-item>
                    <el-dropdown-item :command="{ label: 'newest', sort: 'newest' }">{{
                      i18n['cloudeBrainMirror']['newest'] }}</el-dropdown-item>
                  </el-dropdown-menu>
                </el-dropdown>
              </el-col>
            </el-row>
            <el-row style="margin-top: 15px">
              <el-table :data="tableDataStar" style="width: 100%" :header-cell-style="tableHeaderStyle">
                <el-table-column :label="i18n['cloudeBrainMirror']['mirror_tag']" min-width="19%" align="left"
                  prop="tag">
                  <template slot-scope="scope">
                    <div style="display: flex; align-items: center">
                      <a class="text-over image_title" :title="scope.row.tag">{{
                        scope.row.tag
                      }}</a>
                      <img v-if="scope.row.type == 5" src="/img/jian.svg" style="margin-left: 0.5rem" />
                    </div>
                  </template>
                </el-table-column>
                <el-table-column :label="i18n['cloudeBrainMirror']['mirror_description']" min-width="26%" align="left"
                  prop="description">
                  <template slot-scope="scope">
                    <div class="image_desc" :title="scope.row.description">
                      {{ scope.row.description }}
                    </div>
                    <div v-if="!!scope.row.topics">
                      <span v-for="(topic, index) in scope.row.topics" class="ui repo-topic label topic" :key="index"
                        style="cursor: default">{{ topic }}</span>
                    </div>
                  </template>
                </el-table-column>
                <el-table-column prop="cloudbrainType" :label="i18n['cloudeBrainMirror']['available_clusters']"
                  min-width="14%" align="center">
                  <template slot-scope="scope">
                    {{ scope.row.cloudbrainType | transformType(vm) }}
                  </template>
                </el-table-column>
                <el-table-column prop="creator" :label="i18n['cloudeBrainMirror']['creator']" min-width="8%"
                  align="center">
                  <template slot-scope="scope">
                    <a v-if="scope.row.userName || scope.row.relAvatarLink" :href="'/' + scope.row.userName"
                      :title="scope.row.userName">
                      <img :src="scope.row.relAvatarLink" class="ui avatar image" />
                    </a>
                    <a v-else><img class="ui avatar image" title="Ghost" src="/user/avatar/ghost/-1" /></a>
                  </template>
                </el-table-column>
                <el-table-column prop="createdUnix" :label="i18n['cloudeBrainMirror']['creation_time']" align="center"
                  min-width="14%">
                  <template slot-scope="scope">
                    {{ scope.row.createdUnix | transformTimestamp }}
                  </template>
                </el-table-column>
                <el-table-column align="center" min-width="18%" :label="i18n['cloudeBrainMirror']['operation']">
                  <template slot-scope="scope">
                    <div style="
                        display: flex;
                        justify-content: flex-end;
                        align-items: center;
                      ">
                      <div style="display: flex;align-items: center;padding: 0 1rem;" :title="i18n['citations']">
                        <i class="ri-links-line" style="font-size: 16px;"></i>
                        <span style="line-height: 2;margin-left: 0.3rem;">{{ scope.row.useCount }}</span>
                      </div>
                      <div style="
                          display: flex;
                          align-items: center;
                          cursor: pointer;
                          padding: 0 1rem;
                        " @click="imageUnstar(scope.row.id)">
                        <svg width="1.4em" height="1.4em" viewBox="0 0 32 32" class="heart-stroke stars_active">
                          <path
                            d="M4.4 6.54c-1.761 1.643-2.6 3.793-2.36 6.056.24 2.263 1.507 4.521 3.663 6.534a29110.9 29110.9 0 0010.296 9.633l10.297-9.633c2.157-2.013 3.424-4.273 3.664-6.536.24-2.264-.599-4.412-2.36-6.056-1.73-1.613-3.84-2.29-6.097-1.955-1.689.25-3.454 1.078-5.105 2.394l-.4.319-.398-.319c-1.649-1.316-3.414-2.143-5.105-2.394a7.612 7.612 0 00-1.113-.081c-1.838 0-3.541.694-4.983 2.038z">
                          </path>
                        </svg>
                        <span style="line-height: 2; margin-left: 0.3rem">{{
                          scope.row.numStars
                        }}</span>
                      </div>
                    </div>
                  </template>
                </el-table-column>
              </el-table>
            </el-row>
            <div class="ui container" style="margin-top: 50px; text-align: center">
              <el-pagination background @size-change="handleSizeChangeStar" @current-change="handleCurrentChangeStar"
                :current-page="currentPageStar" :page-size="pageSizeStar" :page-sizes="[5, 10, 15]"
                layout="total, sizes, prev, pager, next, jumper" :total="totalNumStar">
              </el-pagination>
            </div>
          </template>
          <template v-else>
            <el-row style="align-items: center; display: flex">
              <el-col :span="12">
                <div style="visibility: hidden"></div>
              </el-col>
              <el-col :span="10">
                <div>
                  <el-input :placeholder="i18n['cloudeBrainMirror']['placeholder']" v-model="search"
                    class="input-with-select" @keyup.enter.native="searchName()">
                    <el-button id="success" slot="append" icon="el-icon-search" @click="searchName()">{{
                      i18n['cloudeBrainMirror']['search'] }}</el-button>
                  </el-input>
                </div>
              </el-col>
              <el-col :span="2">
                <el-dropdown @command="handleCommandStar" trigger="click"
                  style="border: 1px solid rgba(34,36,38,.15);border-radius: 4px;padding: 0.5rem 1rem;float: right;">
                  <span class="el-dropdown-link">
                    {{ sortStar | transformSort(vm) }}<i class="el-icon-caret-bottom el-icon--right"></i>
                  </span>
                  <el-dropdown-menu slot="dropdown">
                    <el-dropdown-item :command="{ label: 'defaultsort', sort: '' }">{{
                      i18n['cloudeBrainMirror']['defaultsort'] }}</el-dropdown-item>
                    <el-dropdown-item :command="{ label: 'moststars', sort: 'moststars' }">{{
                      i18n['cloudeBrainMirror']['moststars'] }}</el-dropdown-item>
                    <el-dropdown-item :command="{ label: 'mostused', sort: 'mostused' }">{{
                      i18n['cloudeBrainMirror']['mostused'] }}</el-dropdown-item>
                    <el-dropdown-item :command="{ label: 'newest', sort: 'newest' }">{{
                      i18n['cloudeBrainMirror']['newest'] }}</el-dropdown-item>
                  </el-dropdown-menu>
                </el-dropdown>
              </el-col>
            </el-row>
            <el-empty :image-size="200"></el-empty>
          </template>
        </el-tab-pane>
      </el-tabs>
    </div>
  </div>
</template>

<script>
const { _AppSubUrl, _StaticUrlPrefix, csrf } = window.config;

export default {
  components: {},
  data() {
    return {
      activeName: "first",
      search: "",
      checked: false,
      currentPageRecommend: 1,
      pageSizeRecommend: 10,
      totalNumRecommend: 0,
      paramsRecommend: { page: 1, pageSize: 10, q: "", recommend: false, cloudbrainType: -1, sort: "", _csrf: csrf, },
      tableDataRecommend: [],
      loadingRecommend: false,
      sortRecommend: 'defaultsort',

      currentPageCustom: 1,
      pageSizeCustom: 10,
      totalNumCustom: 0,
      paramsCustom: { page: 1, pageSize: 10, q: "", cloudbrainType: -1, sort: "newest", _csrf: csrf },
      tableDataCustom: [],
      starCustom: [],
      loadingCustom: false,
      refreshCustomTimer: null,
      sortCustom: 'newest',

      currentPageStar: 1,
      pageSizeStar: 10,
      totalNumStar: 0,
      paramsStar: { page: 1, pageSize: 10, q: "", cloudbrainType: -1, sort: "", _csrf: csrf },
      tableDataStar: [],
      loadingStar: false,
      sortStar: 'defaultsort',
      vm: this,
    };
  },
  methods: {
    handleClick(tab, event) {
      this.search = "";
      this.stopImageListCustomRefresh();
      if (tab.name == "first") {
        this.paramsRecommend.q = "";
        this.getImageListRecommend();
      }
      if (tab.name == "second") {
        this.getImageListCustom();
      }
      if (tab.name == "third") {
        this.getImageListStar();
      }
    },
    handleCommandRecommend(command) {
      this.sortRecommend = command.label
      this.paramsRecommend.sort = command.sort
      this.getImageListRecommend();
    },
    handleCommandCustom(command) {
      this.sortCustom = command.label
      this.paramsCustom.sort = command.sort
      this.getImageListCustom()
    },
    handleCommandStar(command) {
      this.sortStar = command.label
      this.paramsStar.sort = command.sort
      this.getImageListStar();
    },
    tableHeaderStyle({ row, column, rowIndex, columnIndex }) {
      if (rowIndex === 0) {
        return "background:#f5f5f6;color:#606266";
      }
    },
    handleSizeChangeRecommend(val) {
      this.paramsRecommend.pageSize = val;
      this.getImageListRecommend();
    },
    handleCurrentChangeRecommend(val) {
      this.currentPageRecommend = val;
      this.paramsRecommend.page = val;
      this.getImageListRecommend();
    },
    handleSizeChangeCustom(val) {
      this.paramsCustom.pageSize = val;
      this.getImageListCustom();
    },
    handleCurrentChangeCustom(val) {
      this.paramsCustom.page = val;
      this.getImageListCustom();
    },
    handleSizeChangeStar(val) {
      this.paramsStar.pageSize = val;
      this.getImageListStar();
    },
    handleCurrentChangeStar(val) {
      this.paramsStar.page = val;
      this.getImageListStar();
    },
    getImageListRecommend() {
      this.loadingRecommend = true;
      this.$axios
        .get("/api/v1/images/recommend", {
          params: this.paramsRecommend,
        })
        .then((res) => {
          this.totalNumRecommend = res.data.count;
          this.tableDataRecommend = res.data.images;
          this.loadingRecommend = false;
        });
    },
    getImageListCustom() {
      this.loadingCustom = true;
      this.$axios
        .get("/api/v1/images/custom", {
          params: this.paramsCustom,
        })
        .then((res) => {
          this.totalNumCustom = res.data.count;
          this.tableDataCustom = res.data.images;
          this.tableDataCustom.forEach((element) => {
            this.starCustom.push({ id: element.id });
          });
          this.loadingCustom = false;
          this.getImageListCustomRefresh();
        }).catch(err => {
          console.log(err);
          if (err.response.status == 401) {
            window.location.href = `/user/login?redirect_to=${encodeURIComponent(window.location.href)}`;
          }
        });
    },
    getImageListCustomRefresh() {
      this.stopImageListCustomRefresh();
      this.refreshCustomTimer = setInterval(() => {
        this.tableDataCustom.forEach((item) => {
          if (item.status === 0) {
            this.$axios.get(`/image/${item.id}`, {}).then((res) => {
              const newData = res.data;
              this.tableDataCustom.forEach((it) => {
                if (it.id === newData.id) {
                  it.status = newData.status;
                  it.place = newData.place;
                }
              });
            });
          }
        });
      }, 5000);
    },
    stopImageListCustomRefresh() {
      this.refreshCustomTimer && clearInterval(this.refreshCustomTimer);
    },
    getImageListStar() {
      this.loadingStar = true;
      this.$axios
        .get("/api/v1/images/star", {
          params: this.paramsStar,
        })
        .then((res) => {
          this.totalNumStar = res.data.count;
          this.tableDataStar = res.data.images;
          this.loadingStar = false;
        }).catch(err => {
          console.log(err);
          if (err.response.status == 401) {
            window.location.href = `/user/login?redirect_to=${encodeURIComponent(window.location.href)}`;
          }
        });
    },
    deleteImage(id) {
      let flag = 1;
      let _this = this;
      $(".ui.basic.modal.images")
        .modal({
          onDeny: function () {
            flag = false;
          },
          onApprove: function () {
            _this.$axios.delete("/image/" + id).then((res) => {
              _this.getImageListCustom();
            });
            flag = true;
          },
          onHidden: function () {
            if (flag == false) {
              $(".alert")
                .html(_this.i18n['canceled_operation'])
                .removeClass("alert-success")
                .addClass("alert-danger")
                .show()
                .delay(1500)
                .fadeOut();
            } else {
              $(".alert")
                .html(_this.i18n['successfully_deleted'])
                .removeClass("alert-danger")
                .addClass("alert-success")
                .show()
                .delay(1500)
                .fadeOut();
            }
          },
        })
        .modal("show");
    },
    eidtImage(id) {
      location.href = `/image/${id}/imageSquare`;
    },
    applyImage(id) {
      location.href = `/image/${id}/apply`;
    },
    imageStar(index, id, isStar) {
      if (isStar) {
        this.$axios.put(`/image/${id}/action/unstar`).then((res) => {
          if (res.data.Code == 0) {
            this.tableDataRecommend[index].numStars =
              this.tableDataRecommend[index].numStars - 1;
            this.tableDataRecommend[index].isStar = false;
          } else {
            console.log(res.data.Message);
          }
        });
      } else {
        this.$axios.put(`/image/${id}/action/star`).then((res) => {
          if (res.data.Code == 0) {
            this.tableDataRecommend[index].numStars =
              this.tableDataRecommend[index].numStars + 1;
            this.tableDataRecommend[index].isStar = true;
          } else {
            console.log(res.data.Message);
          }
        }).catch(err => {
          console.log(err);
          if (err.request.responseURL.indexOf('/user/login') >= 0) {
            window.location.href = `/user/login?redirect_to=${encodeURIComponent(window.location.href)}`;
          }
        })
      }
    },
    imageUnstar(id) {
      this.$axios.put(`/image/${id}/action/unstar`).then((res) => {
        if (res.data.Code == 0) {
          this.getImageListStar();
        } else {
          console.log(res.data.Message);
        }
      });
    },
    copyUrl(url) {
      const cInput = document.createElement("input");
      cInput.value = url;
      document.body.appendChild(cInput);
      cInput.select();
      document.execCommand("Copy");
      cInput.remove();
      $("body").toast({
        message: this.i18n['cloudeBrainMirror']['copy_succeeded'],
        showProgress: "bottom",
        showIcon: "check circle",
        class: "info",
        position: "top right",
      });
    },
    searchName() {
      if (this.activeName == "first") {
        this.paramsRecommend.q = this.search;
        this.paramsRecommend.page = 1;
        this.getImageListRecommend();
      }
      if (this.activeName == "second") {
        this.paramsCustom.q = this.search;
        this.paramsCustom.page = 1;
        this.getImageListCustom();
      }
      if (this.activeName == "third") {
        this.paramsStar.q = this.search;
        this.paramsStar.page = 1;
        this.getImageListStar();
      }
    },
  },
  filters: {
    transformType(val, vm) {
      if (val == 0) {
        return `${vm.i18n['cloudeBrainMirror']['openi']} GPU`;
      } else {
        return `${vm.i18n['cloudeBrainMirror']['c2net']} GPU`;
      }
    },
    transformSort(val, vm) {
      if (val === 'moststars') {
        return vm.i18n['cloudeBrainMirror']['moststars'];
      } else if (val === 'mostused') {
        return vm.i18n['cloudeBrainMirror']['mostused'];
      } else if (val === 'defaultsort') {
        return vm.i18n['cloudeBrainMirror']['defaultsort'];
      } else {
        return vm.i18n['cloudeBrainMirror']['newest'];
      }
    },
    transformTimestamp(timestamp) {
      const date = new Date(parseInt(timestamp) * 1000);
      const Y = date.getFullYear() + "-";
      const M =
        (date.getMonth() + 1 < 10
          ? "0" + (date.getMonth() + 1)
          : date.getMonth() + 1) + "-";
      const D =
        (date.getDate() < 10 ? "0" + date.getDate() : date.getDate()) + "  ";
      const h =
        (date.getHours() < 10 ? "0" + date.getHours() : date.getHours()) + ":";
      const m =
        (date.getMinutes() < 10 ? "0" + date.getMinutes() : date.getMinutes()) +
        ":";
      const s =
        date.getSeconds() < 10 ? "0" + date.getSeconds() : date.getSeconds(); // 秒
      const dateString = Y + M + D + h + m + s;
      return dateString;
    },
  },
  watch: {
    checked(val) {
      this.paramsRecommend.page = 1;
      this.paramsRecommend.recommend = val;
      this.currentPageRecommend = 1;
      this.getImageListRecommend();
    },
  },
  mounted() {
    const lang = document.querySelector('html').getAttribute('lang');
    if (lang != 'zh-CN') {
      document.getElementsByClassName('el-col-2').forEach((item) => { item.style.width = "10%" })
    }
  },
  created() {
    this.i18n = window.i18n;
    const params = new URLSearchParams(location.search);
    if (params.has("type") && params.get("type") == "myimage") {
      this.activeName = "second";
      this.getImageListCustom();
    } else {
      this.getImageListRecommend();
    }
  },
  beforeDestroy() {
    this.stopImageListCustomRefresh();
  },
};
</script>

<style scoped lang="less">
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

.el-dropdown-menu__item--divided {
  border-top: 1px solid blue;
}

.el-table thead {
  background-color: #f5f5f6;
}

/deep/ .el-tabs__item:hover {
  color: #000;
  font-weight: 500;
}

/deep/ .el-tabs__item.is-active {
  color: #000;
  font-weight: 500;
}

/deep/ .el-tabs__active-bar {
  background-color: #000;
}

@media screen and (max-width: 1600px) {
  /deep/ .el-col-10 {
    width: 40%;
  }

  /deep/ .el-col-2 {
    width: 10%;
  }
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
  stroke: #fa8c16;
  stroke-width: 2;
  fill: #fff;
}

.stars_active {
  fill: #fa8c16 !important;
  stroke: #fa8c16 !important;
}

.apply-btn,
.copy-adress {
  padding: 0 1rem;
  color: #0366d6;
  cursor: pointer;
}

.copy-adress-no {
  cursor: pointer;
  pointer-events: none;
  opacity: .45 !important;
  color: rgba(0, 0, 0, .6);
  padding: 0 1rem;
}
</style>
