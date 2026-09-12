<template>
  <div class="item">
    <div class="header">
      <div class="header-l">
        <div class="title" :title="data.project_name">{{ data.project_name }}</div>
        <div class="prj-type">{{ data.type }}</div>
      </div>
      <div class="header-r">
        <a :href="`/tech/repo_view?project_name=${data.project_name}`">
          <span>更多成果</span>
          <i class="el-icon-arrow-right"></i>
        </a>
      </div>
    </div>
    <div class="content">
      <div class="item-l-c">
        <div class="row">
          <span class="tit">负责单位：</span>
          <span class="val" :title="data.institution">{{ data.institution }}</span>
        </div>
        <div class="row">
          <span class="tit">参与单位：</span>
          <span class="val" :title="data.all_institution.split(',').join('、')">{{
            data.all_institution.split(',').join('、')
          }}</span>
        </div>
        <div class="row">
          <span class="tit">申报年份：</span>
          <span class="val">{{ data.apply_year }}</span>
        </div>
        <div class="row">
          <span class="tit">执行周期：</span>
          <span class="val">{{ data.execute_period }}</span>
        </div>
        <div class="row">
          <span class="tit">项目成果数：</span>
          <span class="val">{{ data.repo_numer }}</span>
        </div>
      </div>
      <div class="item-r-c">
        <a class="repo-item-c" target="_blank" :href="`/${item.owner_name}/${item.name}`"
          v-for="(item, index) in data.Repos">
          <div class="repo-item">
            <div class="repo-hd">
              <div class="repo-avatar">
                <img v-if="item.rel_avatar_link" class="avatar" :src="item.rel_avatar_link" />
                <img v-else class="avatar" :avatar="item.owner_name" />
              </div>
              <div class="repo-tit">{{ item.alias }}</div>
            </div>
            <div class="repo-content">
              <div class="repo-descr">
                {{ item.description }}
              </div>
            </div>
          </div>
        </a>
        <div class="repo-item-c" v-if="!data.Repos.length">
          <div class="bgtask-header-pic" style="margin-top:-20px;"></div>
          <div class="repo-item" style="text-align:center;">项目成果待展示</div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import dayjs from 'dayjs';
import { lang } from '~/langs';
import { timeSinceUnix } from '~/utils';

export default {
  name: "SciAndTechPrjItem",
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

  .header {
    display: flex;
    align-items: center;
    justify-content: space-between;

    .header-l {
      display: flex;
      align-items: center;

      .title {
        color: rgb(16, 16, 16);
        font-size: 18px;
        font-weight: bold;
        overflow: hidden;
        text-overflow: ellipsis;
        white-space: nowrap;
        max-width: 800px;
      }

      .prj-type {
        margin-left: 14px;
        font-size: 12px;
        font-weight: 400;
        color: rgb(100, 59, 159);
        background: rgb(237, 234, 251);
        padding: 2px 5px;
        border-radius: 4px;
      }
    }

    .header-r {
      display: flex;
      align-items: baseline;
      text-align: right;
      color: rgb(50, 145, 248);
      font-size: 12px;
      width: 80px;
      justify-content: flex-end;

      i {
        margin-left: 1px;
      }
    }
  }

  .content {
    display: flex;
    margin-top: 10px;

    .item-l-c {
      flex: 1;

      .row {
        display: flex;
        font-size: 14px;
        color: rgb(16, 16, 16);
        margin: 8px 0;
        font-weight: 400;

        .tit {
          color: rgba(136, 136, 136, 1);
        }

        .val {
          flex: 1;
          width: 0;
          overflow: hidden;
          text-overflow: ellipsis;
          white-space: nowrap;
        }

        &:last-child {
          margin-bottom: 0;
        }
      }
    }

    .item-r-c {
      flex: 2;
      margin-left: 10px;
      display: flex;

      .repo-item-c {
        display: block;
        flex: 1;
        margin: 10px;
        padding: 20px;
        border-color: rgba(157, 197, 226, 0.4);
        border-width: 1px;
        border-style: solid;
        border-radius: 5px;
        box-shadow: rgba(157, 197, 226, 0.2) 0px 5px 10px 0px;
        background: rgb(255, 255, 255);
        width: 0;

        .repo-item {
          .repo-hd {
            display: flex;
            align-items: center;

            .repo-avatar {
              width: 28px;
              height: 28px;

              img {
                width: 100%;
                height: 100%;
                border-radius: 100%;
              }
            }

            .repo-tit {
              flex: 1;
              margin-left: 8px;
              width: 0;
              font-size: 14px;
              color: rgb(50, 145, 248);
              overflow: hidden;
              text-overflow: ellipsis;
              white-space: nowrap;
            }
          }

          .repo-content {
            margin-top: 10px;

            .repo-descr {
              width: 100%;
              font-size: 12px;
              color: rgb(136, 136, 136);
              text-overflow: ellipsis;
              word-break: break-all;
              display: -webkit-box;
              -webkit-box-orient: vertical;
              -webkit-line-clamp: 2;
              max-height: 41px;
              overflow: hidden;
            }
          }
        }
      }
    }
  }
}
</style>
