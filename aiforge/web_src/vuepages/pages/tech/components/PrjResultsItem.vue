<template>
  <div class="item">
    <div class="title-c">
      <div class="avatar-c">
        <img v-if="data.rel_avatar_link" class="avatar" :src="data.rel_avatar_link" />
        <img v-else class="avatar" :avatar="data.owner_name" />
      </div>
      <div class="title">
        <a target="_blank" :href="`/${data.owner_name}/${data.name}`">
          <span :title="data.alias">{{ data.alias }}</span>
        </a>
      </div>
    </div>
    <div class="descr" :title="data.description">
      {{ data.description }}
    </div>
    <div class="topics">
      <a v-for="(item, index) in data.topics" :key="index" class="topic" target="_blank"
        :href="`/explore/repos?q=&topic=${item}&sort=hot`">{{ item }}</a>
    </div>
    <div class="footer">
      <div class="contractor" :title="data.institution.split(',').join('、')">{{ data.institution.split(',').join('、') }}
      </div>
      <div class="update-time">
        <span>{{ $t('repos.updated') }}</span>
        <el-tooltip effect="dark" :content="dateFormat(data.updated_unix)" placement="top-start">
          <span>{{ calcFromNow(data.updated_unix) }}</span>
        </el-tooltip>
      </div>
    </div>
  </div>
</template>

<script>
import dayjs from 'dayjs';
import { lang } from '~/langs';
import { timeSinceUnix } from '~/utils';

export default {
  name: "PrjResultsItem",
  props: {
    data: { type: Object, default: () => ({}) },
  },
  components: {},
  data() {
    return {};
  },
  methods: {
    calcFromNow(unix) {
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
.item {
  border-color: rgb(232, 224, 236);
  border-width: 1px;
  border-style: solid;
  box-shadow: rgba(168, 157, 226, 0.2) 0px 5px 10px 0px;
  background: url("data:image/svg+xml;charset=utf8,%3Csvg%20xmlns%3D%22http%3A%2F%2Fwww.w3.org%2F2000%2Fsvg%22%20xmlns%3Axlink%3D%22http%3A%2F%2Fwww.w3.org%2F1999%2Fxlink%22%20version%3D%221.1%22%3E%3Cdefs%3E%3ClinearGradient%20id%3D%221%22%20x1%3D%220%22%20x2%3D%221%22%20y1%3D%220%22%20y2%3D%220%22%20gradientTransform%3D%22matrix(-0.01900000000000005%2C%200.997%2C%20-0.12862587255310284%2C%20-0.01900000000000005%2C%200.995%2C%200.014)%22%3E%3Cstop%20stop-color%3D%22%23f2edf5%22%20stop-opacity%3D%221%22%20offset%3D%220.01%22%3E%3C%2Fstop%3E%3Cstop%20stop-color%3D%22%23ffffff%22%20stop-opacity%3D%221%22%20offset%3D%220.31%22%3E%3C%2Fstop%3E%3C%2FlinearGradient%3E%3C%2Fdefs%3E%3Crect%20width%3D%22100%25%22%20height%3D%22100%25%22%20fill%3D%22url(%231)%22%3E%3C%2Frect%3E%3C%2Fsvg%3E");
  overflow: hidden;
  padding: 15px;
}

.title-c {
  display: flex;
  align-items: center;

  .avatar-c {
    height: 28px;
    width: 28px;
    margin-right: 6px;

    img {
      border-radius: 100%;
      height: 100%;
      width: 100%;
    }
  }

  .title {
    flex: 1;
    overflow: hidden;
    text-overflow: ellipsis;
    white-space: nowrap;

    span {
      color: rgb(16, 16, 16);
      font-size: 16px;
      font-weight: 400;
    }
  }
}

.descr {
  height: 40px;
  margin-top: 8px;
  font-size: 12px;
  font-weight: 300;
  color: rgb(136, 136, 136);
  text-overflow: ellipsis;
  word-break: break-all;
  display: -webkit-box;
  -webkit-box-orient: vertical;
  -webkit-line-clamp: 2;
  max-height: 41px;
  overflow: hidden;
}

.topics {
  overflow: hidden;
  text-overflow: ellipsis;
  white-space: nowrap;
  margin-top: 5px;
  height: 20px;

  .topic {
    color: rgba(16, 16, 16, 0.8);
    border-radius: 4px;
    font-size: 12px;
    background: rgba(232, 232, 232, 0.6);
    padding: 2px 6px;
    margin-right: 8px;
  }
}

.footer {
  margin-top: 12px;
  display: flex;
  align-items: center;
  justify-content: space-between;
  height: 20px;

  .contractor {
    font-size: 14px;
    font-weight: 400;
    color: rgba(0, 40, 192, 0.73);
    overflow: hidden;
    text-overflow: ellipsis;
    white-space: nowrap;
  }

  .update-time {
    text-align: right;
    min-width: 140px;
    font-size: 12px;
    font-weight: 300;
    color: rgb(136, 136, 136);
  }
}
</style>
