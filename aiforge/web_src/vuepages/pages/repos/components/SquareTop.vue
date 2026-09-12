<template>
  <div class="ui _repo_container_bg">
    <div class="ui container _repo_container _content_container">
      <div class="_repo_top_left">
        <a :href="bannerData[0] ? bannerData[0].url : ''">
          <div class="_repo_top_left_img">
            <img :src="bannerData[0] ? bannerData[0].src : ''">
          </div>
        </a>
      </div>
      <div class="_repo_top_middle">
        <div class="_repo_top_middle_header">
          <div class="_repo_top_mid_item" :class="(tabIndex == index) ? '_foucs' : ''" v-for="(item, index) in tabs"
            :key="index" @click="changeTab(item, index)">
            <i :class="item.icon"></i>
            <span>{{ item.label }}</span>
          </div>
        </div>
        <div class="_repo_top_middle_content">
          <div class="_repo_top_mid_repo_list">
            <div class="swiper-wrapper" id="_repo_top_mid_repo"></div>
            <div class="swiper-pagination _repo_top-swiper-pagination"></div>
          </div>
        </div>
      </div>
      <div class="_repo_top_right">
        <a :href="bannerData[1] ? bannerData[1].url : ''">
          <div class="_repo_top_right_img">
            <img :src="bannerData[1] ? bannerData[1].src : ''">
          </div>
        </a>
      </div>
    </div>
  </div>
</template>

<script>
import { getPromoteData } from '~/apis/modules/common';
import { getReposSquareTabData } from '~/apis/modules/repos';
import LetterAvatar from '~/utils/letteravatar';

export default {
  name: "SquareTop",
  props: {
    static: { type: Boolean, default: false },
    staticSwiperData: { type: Array, default: () => [] },
    staticBannerData: { type: String, default: '[]' },
  },
  components: {},
  data() {
    return {
      swiperHandler: null,
      tabIndex: 0,
      tabs: [{
        key: 'preferred',
        icon: 'ri-fire-line',
        label: this.$t('repos.preferred'),
      }, {
        key: 'incubation',
        icon: 'ri-award-line',
        label: this.$t('repos.openIIncubation'),
      }, {
        key: 'hot-paper',
        icon: 'ri-file-damage-line',
        label: this.$t('repos.hotPapers'),
      }],
      bannerData: [],
    };
  },
  methods: {
    initSwiper() {
      this.swiperHandler = new Swiper("._repo_top_mid_repo_list", {
        slidesPerView: 1,
        spaceBetween: 20,
        pagination: {
          el: "._repo_top-swiper-pagination",
          clickable: true,
        },
        autoplay: {
          delay: 4500,
          disableOnInteraction: false,
        },
        breakpoints: {
          768: {
            slidesPerView: 2,
          },
          1024: {
            slidesPerView: 2,
          },
          1200: {
            slidesPerView: 3,
          },
        },
      });
    },
    renderSwiper(data) {
      const swiperEl = document.getElementById("_repo_top_mid_repo");
      let html = '';
      const width = swiperEl.parentNode.clientWidth;
      for (let i = 0, iLen = data.length; i < iLen; i++) {
        html += `<div class="swiper-slide">`;
        for (let j = i; j < i + 2; j++) {
          let dataJ = data[j];
          if (dataJ === undefined) break;
          html += `<div class="_repo_sw_card"><div style="display:flex;">
            ${dataJ["RelAvatarLink"] ? `<img style="border-radius:100%;width:35px;height:35px;margin-bottom:0.6em;" class="left floated mini ui image" src="${dataJ["RelAvatarLink"]}">`
              : `<img style="border-radius:100%;width:35px;height:35px;margin-bottom:0.6em;" class="left floated mini ui image" avatar="${dataJ["Name"]}">`}
                  <span class="header nowrap" style="color:rgb(50, 145, 248);font-size:14px;height:35px;line-height:35px;" href="javascript:;" title="${dataJ["Alias"]}">${dataJ["Alias"]}</span>
						</div><div class="_repo_sw_card_descr _repo_nowrap_line_3" title="${dataJ.Description}">${dataJ.Description}</div>                   
          <a href="/${dataJ.OwnerName}/${dataJ.Name}" style="position:absolute;width:100%;height:100%;top:0;left:0;"></a>         
					</div>`;
        }
        html += `</div>`;
        i++;
      }
      this.swiperHandler.removeAllSlides();
      swiperEl.innerHTML = html;
      this.swiperHandler.updateSlides();
      this.swiperHandler.updateProgress();
      LetterAvatar.transform();
    },
    getBannerData() {
      getPromoteData('/repos/square_banner').then(res => {
        const data = res.data;
        try {
          const list = JSON.parse(data);
          this.bannerData = list;
        } catch (err) {
          console.log(err);
        }
      }).catch(err => {
        this.bannerData = [];
        console.log(err);
      });
    },
    getTabData() {
      getReposSquareTabData(this.tabs[this.tabIndex].key).then(res => {
        res = res.data;
        if (res.Code == 0) {
          const data = res.Data.Repos || [];
          this.renderSwiper(data);
        } else {
          this.renderSwiper([]);
        }
      }).catch(err => {
        console.log(err);
        this.renderSwiper([]);
      });
    },
    changeTab(item, index) {
      this.tabIndex = index;
      this.getTabData();
    },
  },
  mounted() {
    this.initSwiper();
    if (this.static) {
      try {
        this.bannerData = JSON.parse(this.staticBannerData);
      } catch (err) {
        console.log(err);
      }
      this.renderSwiper(this.staticSwiperData);
    } else {
      this.getBannerData();
      this.getTabData();
    }
  },
};
</script>

<style scoped lang="less">
._repo_container_bg {
  width: 100%;
  height: 450px;
  background: url("data:image/svg+xml;charset=utf8,%3Csvg%20xmlns%3D%22http%3A%2F%2Fwww.w3.org%2F2000%2Fsvg%22%20xmlns%3Axlink%3D%22http%3A%2F%2Fwww.w3.org%2F1999%2Fxlink%22%20version%3D%221.1%22%3E%3Cdefs%3E%3ClinearGradient%20id%3D%221%22%20x1%3D%220%22%20x2%3D%221%22%20y1%3D%220%22%20y2%3D%220%22%20gradientTransform%3D%22matrix(0.4999999999999999%2C%20-0.8660000000000001%2C%200.07722000385802469%2C%200.4999999999999999%2C%20-0.183%2C%200.683)%22%3E%3Cstop%20stop-color%3D%22%239aceec%22%20stop-opacity%3D%221%22%20offset%3D%220%22%3E%3C%2Fstop%3E%3Cstop%20stop-color%3D%22%23eeeeee%22%20stop-opacity%3D%221%22%20offset%3D%220.99%22%3E%3C%2Fstop%3E%3C%2FlinearGradient%3E%3C%2Fdefs%3E%3Crect%20width%3D%22100%25%22%20height%3D%22100%25%22%20fill%3D%22url(%231)%22%3E%3C%2Frect%3E%3C%2Fsvg%3E");
  position: relative;
}

._content_container {
  display: flex !important;
  margin: 0 auto !important;
  height: 100% !important;
}

._repo_top_left {
  width: 243px;
  height: 100%;
  position: relative;
}

._repo_top_left_img {
  width: 100%;
  height: 346px;
  position: absolute;
  bottom: 0;
  border-top-left-radius: 10px;
  overflow: hidden;
}

._repo_top_left_img img {
  height: 100%;
  width: 100%;
}

._content_container img[src=""],
img:not([src]) {
  opacity: 0;
}

._repo_top_right {
  width: 243px;
  height: 100%;
  position: relative;
}

._repo_top_right_img {
  width: 100%;
  height: 346px;
  position: absolute;
  bottom: 0;
  border-top-right-radius: 10px;
  overflow: hidden;
}

._repo_top_right_img img {
  height: 100%;
  width: 100%;
}

._repo_top_middle {
  flex: 1;
  height: 100%;
  position: relative;

}

._repo_top_middle_header {
  width: 100%;
  height: 42px;
  position: absolute;
  bottom: 346px;
  display: flex;
  align-items: center;
  background: raba(0, 0, 255, 0.3);
  padding: 0 20px;
}

._repo_top_mid_item {
  padding: 0px 16px;
  font-size: 18px;
  color: rgba(47, 9, 69, 0.64);
  height: 100%;
  display: flex;
  align-items: center;
  cursor: pointer;
  box-sizing: border-box;
  border-bottom: 3px solid transparent;
}

._repo_top_mid_item i {
  margin-right: 5px;
}

._repo_top_mid_item._foucs {
  background-color: rgba(255, 255, 255, 0.3);
  color: rgb(16, 16, 16);
  cursor: default;
  border-bottom: 3px solid rgba(16, 16, 16, 0.8);
}

._repo_top_middle_content {
  width: 100%;
  height: 346px;
  position: absolute;
  bottom: 0;
  background: rgba(255, 255, 255, 0.3);
  padding: 20px 20px;
  overflow: hidden;
}

._repo_top_mid_repo_list {
  overflow: hidden;
}

/deep/._repo_sw_card {
  height: 128px;
  border-color: rgb(255, 255, 255);
  border-width: 1px;
  border-style: solid;
  font-size: 14px;
  padding: 12px;
  background: rgba(255, 255, 255, 0.6);
  margin-bottom: 20px;
  box-sizing: border-box;
  position: relative;
}

/deep/._repo_nowrap {
  overflow: hidden;
  text-overflow: ellipsis;
  white-space: nowrap;
}

/deep/._repo_nowrap_line_3 {
  overflow: hidden;
  text-overflow: ellipsis;
  word-break: break-all;
  font-size: 12px;
  color: rgb(136, 136, 136);
  display: -webkit-box;
  -webkit-box-orient: vertical;
  -webkit-line-clamp: 3;
  max-height: 65px;
  white-space: break-spaces;
}

/deep/._repo_sw_card a {
  color: inherit;
}

/deep/._repo_sw_card_title {
  font-weight: 700;
  font-size: 14px;
  color: rgba(26, 40, 51, 1);
  margin-bottom: 10px;
}

/deep/._repo_sw_card_descr {
  font-size: 12px;
  color: rgba(80, 85, 89, 1);
  margin-bottom: 10px;
  min-height: 42px;
}

/deep/._repo_sw_card_label {
  color: rgb(26, 40, 51);
  font-size: 14px;
  display: flex;
  width: 100%;
  position: relative;
}

/deep/._repo_sw_card_label span {
  border-radius: 4px;
  background: rgba(232, 232, 232, 0.6);
  padding: 0px 6px 2px;
  margin-right: 10px;
  max-width: 50%;
}

/deep/._repo_top_mid_repo_list .swiper-pagination-bullet {
  width: 8px;
  height: 8px;
  border-radius: 100%;
  background: #76cbed;
}

/deep/._repo_top_mid_repo_list .swiper-pagination-bullet-active {
  width: 40px;
  border-radius: 4px;
}
</style>
