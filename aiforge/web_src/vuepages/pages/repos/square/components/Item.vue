<template>
  <a class="item" target="_blank" :href="`/${data.OwnerName}/${data.Name}`">
    <div class="top">
      <div class="top-head">
        <div class="icon-c">
          <img v-if="data.RelAvatarLink" class="avatar" :src="data.RelAvatarLink" />
          <img v-else class="avatar" :avatar="data.Name" />
        </div>
        <div class="name" :title="`${data.OwnerName}/${data.Name}`">
          <span class="title-1">{{ data.OwnerName }}</span>
          <span class="title-1"> / </span>
          <span class="title-2" v-html="data.NameShow"></span>
          <i v-if="data.IsArchived" class="archive icon archived-icon"></i>
          <svg v-if="data.IsFork" class="svg octicon-repo-forked" width="15" height="15" aria-hidden="true">
            <use xlink:href="#octicon-repo-forked"></use>
          </svg>
          <svg v-if="data.IsMirror" class="svg octicon-repo-clone" width="15" height="15" aria-hidden="true">
            <use xlink:href="#octicon-repo-clone"></use>
          </svg>
          <svg v-if="(data.IsPrivate || data.IsOwnerPrivate)" style="color:#a1882b!important" class="svg octicon-lock"
            width="15" height="15" aria-hidden="true">
            <use xlink:href="#octicon-lock"></use>
          </svg>
        </div>
      </div>
      <div class="top-descr">
        <div class="descr" :title="data.DescriptionShow" v-show="data.DescriptionShow" v-html="data.DescriptionShow">
        </div>
      </div>
      <div class="top-bottom">
        <div class="labels">
          <a v-for="(item, index) in data.TopicsShow" :key="index" class="label normal">{{ item.topicShow }}</a>
        </div>
      </div>
    </div>
    <div class="footer">
      <div class="footer-l">
        <!-- <a :href="`/${data.userName}`" class="avatar-c">
          <img class="avatar" :src="data.userRelAvatarLink">
        </a> -->
        <div>
          <span>{{ $t('repos.updated') }}</span>
          <el-tooltip effect="dark" :content="dateFormat(data.UpdatedUnix)" placement="top-start">
            <span>{{ calcFromNow(data.UpdatedUnix) }}</span>
          </el-tooltip>
          <span style="margin-left:8px;" v-if="data.PrimaryLanguage"><i class="color-icon"
              :style="{ backgroundColor: data.PrimaryLanguage.Color }"></i>{{ data.PrimaryLanguage.Language }}</span>
        </div>
      </div>
      <div class="footer-r">
        <span class="r-item" :title="$t('repos.watch')">
          <i class="ri-eye-line"></i>
          <span>{{ data.NumWatches }}</span>
        </span>
        <div class="line"></div>
        <span class="r-item" :title="$t('repos.star')">
          <i class="ri-star-line"></i>
          <span>{{ data.NumStars }}</span>
        </span>
        <div class="line"></div>
        <span class="r-item" :title="$t('repos.fork')">
          <svg class="svg octicon-repo-forked" width="13" height="13" aria-hidden="true">
            <use xlink:href="#octicon-repo-forked"></use>
          </svg>
          <span>{{ data.NumForks }}</span></span>
      </div>
    </div>
  </a>
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
  name: "Item",
  props: {
    condition: { type: Object, default: () => ({}) },
    data: { type: Object, default: () => ({}) },
  },
  components: {},
  data() {
    return {};
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
  beforeMount() { },
};
</script>

<style scoped lang="less">
.item {
  display: flex;
  width: 100%;
  min-height: 146px;
  box-shadow: 0px 4px 4px 0px rgba(16, 16, 16, 0.1);
  border-radius: 10px;
  border: 1px solid rgba(255,255,255,1);
  background: linear-gradient(176.14deg, rgba(233,247,255,1) 0.84%,rgba(255,255,255,1) 54.19%);
  justify-content: space-between;
  flex-direction: column;
  height: 100%;
  &:hover {
    border: 1px solid rgba(0,102,255,1);
    .top .top-head .name {
      color: rgba(0,102,255,1);
    }
  }
  .top {

    .top-head {
      display: flex;
      align-items: center;
      position: relative;
      height: 55px;
      border-top-left-radius: 15px;
      border-top-right-radius: 15px;
      padding: 16px;

      .icon-c {
        width: 32px;
        height: 32px;
        margin-right: 6px;
        border-radius: 50%;

        img {
          width: 100%;
          height: 100%;
          border-radius: 50%;
        }
      }

      .name {
        color: rgba(14, 37, 69, 1);
        font-size: 16px;
        flex: 1;
        width: 0;
        display: flex;
        align-items: center;
        display: block;
        overflow: hidden;
        text-overflow: ellipsis;
        white-space: nowrap;

        span {}

        .title-2 {
          margin-right: 3px;
        }
      }

      .reconmend-icon {
        position: absolute;
        right: 15px;
        top: 0px;
        height: 19.28px;
        width: 29.61px;

        img {
          height: 100%;
          width: 100%;
        }
      }
    }

    .top-bottom {
      padding: 16px;
      padding-top: 12px;
      padding-bottom: 0px;

      .labels {
        display: flex;
        align-items: center;
        flex-wrap: wrap;

        .label {
          border-radius: 3px;
          color: rgba(16,16,16,0.7);
          font-size: 12px;
          padding: 0px 6px;
          margin-right: 6px;
          margin-bottom: 6px;
          overflow: hidden;
          text-overflow: ellipsis;
          white-space: nowrap;

          &.normal {
            background-color: rgba(16,16,16,0.05);
          }
        }
      }
    }

    .top-descr {
      padding: 0 16px;

      .descr {
        color: rgba(16, 16, 16, 0.75);
        font-size: 12px;
        overflow: hidden;
        text-overflow: ellipsis;
        word-break: break-all;
        display: -webkit-box;
        -webkit-box-orient: vertical;
        -webkit-line-clamp: 2;
        max-height: 40px;
        white-space: break-spaces;
      }
    }
  }

  .footer {
    display: flex;
    align-items: center;
    justify-content: space-between;
    margin-top: 2px;
    padding: 6px 16px 10px 16px;

    .footer-l {
      font-size: 12px;
      font-weight: 400;
      overflow: hidden;
      text-overflow: ellipsis;
      white-space: nowrap;
      color: rgba(136, 136, 136, 1);
      display: flex;
      align-items: center;

      .avatar-c {
        margin-right: 4px;
        display: inline-block;
        height: 24px;

        .avatar {
          display: inline-block;
          width: 24px;
          height: 24px;
          border-radius: 100%;
        }
      }

      .greenPoint {
        display: inline-block;
        width: 8px;
        height: 8px;
        margin-right: 5px;
        background-color: #5dbf77;
        border-radius: 50%;
      }
    }

    .footer-r {
      display: flex;
      align-items: center;
      justify-content: flex-end;;

      .line {
        height: 16px;
        margin-left: 8px;
        border-right: 1px solid rgba(16, 16, 16, 0.1);
      }

      .r-item {
        display: flex;
        align-items: center;
        margin-left: 8px;
        color: rgba(136, 136, 136, 1);
        font-size: 12px;

        i {
          margin-right: 5px;
        }

        span {}
      }
    }
  }
}
</style>
