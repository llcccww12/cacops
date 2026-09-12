<template>
  <div>
    <div class="item">
      <div class="item-top">
        <img v-if="data.RelAvatarLink" class="avatar" :src="data.RelAvatarLink" />
        <img v-else class="avatar" :avatar="data.Name" />
        <div class="content">
          <div class="title">
            <div class="title-l">
              <a :href="`/${data.OwnerName}/${data.Name}`" :title="`${data.OwnerName}/${data.Name}`">
                <span class="title-1">{{ data.OwnerName }}</span>
                <span class="title-1"> / </span>
                <span class="title-2" v-html="data.NameShow"></span>
              </a>
              <i v-if="data.IsArchived" class="archive icon archived-icon"></i>
              <svg v-if="data.IsFork" class="svg octicon-repo-forked" width="15" height="15" aria-hidden="true">
                <use xlink:href="#octicon-repo-forked"></use>
              </svg>
              <svg v-if="data.IsMirror" class="svg octicon-repo-clone" width="15" height="15" aria-hidden="true">
                <use xlink:href="#octicon-repo-clone"></use>
              </svg>
              <svg v-if="(data.IsPrivate || data.IsOwnerPrivate)" style="color:#a1882b!important"
                class="svg octicon-lock" width="15" height="15" aria-hidden="true">
                <use xlink:href="#octicon-lock"></use>
              </svg>
            </div>
            <span class="title-r only-mobile-hidden">
              <span class="t-item" :title="$t('repos.watch')">
                <i class="ri-eye-line"></i>
                <span>{{ data.NumWatches }}</span>
              </span>
              <span class="t-item" :title="$t('repos.star')">
                <i class="ri-star-line"></i>
                <span>{{ data.NumStars }}</span>
              </span>
              <span class="t-item" :title="$t('repos.fork')">
                <svg class="svg octicon-repo-forked" width="13" height="13" aria-hidden="true">
                  <use xlink:href="#octicon-repo-forked"></use>
                </svg>
                <span>{{ data.NumForks }}</span></span>
            </span>
          </div>
          <div class="descr" v-show="data.DescriptionShow" v-html="data.DescriptionShow"></div>
          <div class="tags" :class="topicLink ? '' : 'hide-link'" v-show="data.Topics && data.Topics.length">
            <a v-for="(item, index) in data.TopicsShow" :key="index" class="tag"
              :class="(item.topic.toLocaleLowerCase() == topic.toLocaleLowerCase() ? 'tag-focus' : '')"
              :href="topicLink ? `/explore/repos?q=&topic=${item.topic}&sort=hot` : 'javascript:;'"
              v-html="item.topicShow"></a>
          </div>
          <!-- <div class="repo-datas" v-show="(data.DatasetCnt > 0) || (data.ModelCnt > 0) || (data.AiTaskCnt > 0)">
            <span class="repo-datas-item" v-show="(data.DatasetCnt > 0)">
              <i class="ri-stack-line"></i>
              <span class="label only-mobile-hidden">{{ $t('repos.dataset') }}：</span>
              <span class="value">{{ data.DatasetCnt }}</span>
            </span>
            <span class="repo-datas-item" v-show="(data.ModelCnt > 0)">
              <i class="ri-send-plane-2-line"></i>
              <span class="label only-mobile-hidden">{{ $t('repos.model') }}：</span>
              <span class="value">{{ data.ModelCnt }}</span>
            </span>
            <span class="repo-datas-item" v-show="(data.AiTaskCnt > 0)">
              <i class="ri-order-play-line"></i>
              <span class="label only-mobile-hidden">{{ $t('repos.aiTask') }}：</span>
              <span class="value">{{ data.AiTaskCnt }}</span>
            </span>
          </div> -->
        </div>
      </div>
      <div class="item-bottom">
        <div>
          <span>{{ $t('repos.updated') }}</span>
          <el-tooltip effect="dark" :content="dateFormat(data.UpdatedUnix)" placement="top-start">
            <span>{{ calcFromNow(data.UpdatedUnix) }}</span>
          </el-tooltip>
          <span style="margin-left:8px;" v-if="data.PrimaryLanguage"><i class="color-icon"
              :style="{ backgroundColor: data.PrimaryLanguage.Color }"></i>{{ data.PrimaryLanguage.Language }}</span>
        </div>
        <div class="contributors">
          <span class="contributors-count" v-show="data.Contributors && data.Contributors.length">
            {{ $t('repos.contributors') }}&nbsp;
          </span>
          <span class="contributors-avatar">
            <a :href="item.UserName ? `/${item.UserName}` : (item.Email ? `mailto:${item.Email}` : 'javascript:;')"
              class="avatar-c" v-for="(item, index) in data.Contributors" :key="index">
              <img class="avatar" v-show="item.UserName" :src="item.RelAvatarLink">
              <span class="avatar" v-show="!item.UserName" :style="{ backgroundColor: item.bgColor }">
                {{ (item.Email[0] || '').toLocaleUpperCase() }}</span>
            </a>
          </span>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import relativeTime from 'dayjs/plugin/relativeTime';
import localizedFormat from 'dayjs/plugin/localizedFormat';
import 'dayjs/locale/zh-cn';
import 'dayjs/locale/en';
import dayjs from 'dayjs';
import { lang } from '~/langs';
import { timeSinceUnix } from '~/utils';

dayjs.locale(lang == 'zh-CN' ? 'zh-cn' : 'en');
dayjs.extend(relativeTime);
dayjs.extend(localizedFormat);

export default {
  name: "ReposItem",
  props: {
    data: { type: Object, default: () => ({}) },
    topic: { type: String, default: '' },
    topicLink: { type: Boolean, default: true },
  },
  components: {},
  data() {
    return {
      contributors: [],
    };
  },
  methods: {
    calcFromNow(unix) {
      // return dayjs(unix * 1000).fromNow();
      return timeSinceUnix(unix, Date.now() / 1000);
    },
    dateFormat(unix) {
      return lang == 'zh-CN' ? dayjs(unix * 1000).format('YYYY年MM月DD日 HH时mm分ss秒') :
        dayjs(unix * 1000).format('ddd, D MMM YYYY HH:mm:ss [CST]');
    }
  },
  mounted() { },
};
</script>

<style scoped lang="less">
@media only screen and (max-width: 767px) {
  .only-mobile-hidden {
    display: none !important;
  }
}

.item {
  width: 100%;
  border-color: rgba(157, 197, 226, 0.4);
  border-width: 1px;
  border-style: solid;
  box-shadow: rgba(157, 197, 226, 20%) 0px 5px 10px 0px;
  border-radius: 15px;
  font-size: 14px;
  padding: 20px 26px 10px 26px;
  margin-bottom: 40px;
}

.item-top {
  display: flex;
}

.item-top .avatar {
  width: 38px;
  height: 38px;
  margin-right: 10px;
  border-radius: 100%;
}

.content {
  flex: 1;
  overflow: hidden;
}

.content .title {
  display: flex;
  align-items: center;
  height: 30px;
  margin: 4px 0 8px;
}

.content .title-l {
  flex: 1;
  overflow: hidden;
  width: 100%;
  text-overflow: ellipsis;
  white-space: nowrap;
}

.content .title-1 {
  font-size: 18px;
  color: rgba(16, 16, 16, 0.6);
}

.content .title-2 {
  font-size: 18px;
  color: rgba(16, 16, 16, 1);
  font-weight: bold;
  margin-right: 3px;
}

.content .title-r {
  display: flex;
  align-items: center;
  font-weight: 400;
  font-size: 12px;
  color: rgba(26, 40, 51, 1);
  justify-content: flex-end;
}

.content .t-item {
  margin-left: 12px;
  display: flex;
  align-items: center;
}

.content .t-item i {
  margin-right: 4px;
}

.content .descr {
  font-weight: 300;
  font-size: 14px;
  color: rgba(16, 16, 16, 0.8);
  margin-bottom: 16px;
  overflow: hidden;
  text-overflow: ellipsis;
  word-break: break-all;
  display: -webkit-box;
  -webkit-box-orient: vertical;
  -webkit-line-clamp: 6;
  max-height: 120px;
  white-space: break-spaces;
}

.content .tags {
  margin-bottom: 16px;
  overflow: hidden;
  text-overflow: ellipsis;
  white-space: nowrap;

  &.hide-link {
    .tag {
      cursor: default;
    }
  }
}

.content .tag {
  color: rgba(16, 16, 16, 0.8);
  border-radius: 4px;
  font-size: 14px;
  background: rgba(232, 232, 232, 0.6);
  padding: 2px 6px;
  margin-right: 8px;

  &.tag-focus {
    color: red;
  }
}

.content .repo-datas {
  display: flex;
  align-items: center;
  margin-top: 20px;
  margin-bottom: 10px;
}

.content .repo-datas-item {
  display: flex;
  align-items: center;
  margin-right: 24px;
}

.content .repo-datas-item i {
  color: rgba(2, 107, 251, 0.54);
  margin-right: 4px;
  font-size: 16px;
}

.content .repo-datas-item .label {
  color: rgba(2, 107, 251, 0.54);
  margin-right: 4px;
}

.content .repo-datas-item .value {
  font-weight: bold;
}

.item-bottom {
  display: flex;
  align-items: center;
  justify-content: space-between;
  border-top: 1px solid rgba(157, 197, 226, 0.2);
  margin-top: 10px;
  padding-top: 10px;
  font-size: 12px;
  color: rgba(16, 16, 16, 0.6);
}

.item-bottom .contributors {
  display: flex;
  align-items: center;
}

.item-bottom .contributors-avatar {
  display: flex;
  align-items: center;
  margin-left: 16px;

  .avatar-c {

    img[src=""],
    img:not([src]) {
      // opacity: 0;
    }
  }
}

.item-bottom .avatar {
  display: block;
  width: 25px;
  height: 25px;
  margin-left: -6px;
  border-radius: 100%;
  border: 1px solid white;
  font-size: 16px;
  line-height: 24px;
  text-align: center;
  color: white;
  background-color: white;
  font-weight: bold;
}
</style>
