<template>
  <div class="repo-selected-bg">
    <div class="ui container _repo_container _repo-selected-container" style="padding-top:3rem;padding-bottom:3rem;">
      <div class="_repo_title"><span>{{ $t('repos.selectedFields') }}</span></div>
      <div class="_repo-selected-list">
        <div class="swiper-wrapper" id="_repo-selected"></div>
        <div class="swiper-pagination _repo-selected-swiper-pagination"></div>
      </div>
    </div>
  </div>
</template>

<script>
import { getHomePageData } from '~/apis/modules/repos';
import LetterAvatar from '~/utils/letteravatar';

export default {
  name: "RecommendRepos",
  props: {
    static: { type: Boolean, default: false },
    staticSwiperData: { type: Array, default: () => [] },
  },
  components: {},
  data() {
    return {
      swiperHandler: null,
    };
  },
  methods: {
    renderRepos(json) {
      var selectedRepoEl = document.getElementById("_repo-selected");
      var html = "";
      if (json != null && json.length > 0) {
        var repoMap = {};
        for (var i = 0, iLen = json.length; i < iLen; i++) {
          var repo = json[i];
          var label = repo.Label;
          if (repoMap[label]) {
            repoMap[label].push(repo);
          } else {
            repoMap[label] = [repo];
          }
        }
        for (var label in repoMap) {
          var repos = repoMap[label];
          var labelSearch = repos[0].Label;
          html += `<div class="swiper-slide"><div><a style="color:rgb(50, 145, 248);font-size:16px;font-weight:550;" href="/explore/repos?q=&topic=${labelSearch}&sort=hot"># ${label}</a></div>`;
          for (var i = 0, iLen = repos.length; i < iLen; i++) {
            if (i >= 4) break;
            var repo = repos[i];
            html += `<div class="ui fluid card">             
                <div class="content">
                  ${repo["Avatar"] ? `<img style="border-radius: 100%;" class="left floated mini ui image" src="${repo["Avatar"]}">` : `<img style="border-radius: 100%;" class="left floated mini ui image" avatar="${repo["Name"]}">`}
                  <span class="header nowrap" style="color:rgb(50, 145, 248);font-size:14px;" href="javascript:;" title="${repo["Alias"]}">${repo["Alias"]}</span>
                  <div class="description nowrap-2" style="rgba(136,136,136,1);;font-size:12px;" title="${repo["Description"]}">${repo["Description"]}</div>
                </div>
                <a style="position:absolute;height:100%;width:100%;" href="/${repo["OwnerName"]}/${repo["Name"]}"></a>
            </div>`;
          }
          html += '</div>'
        }
        this.swiperHandler = new Swiper("._repo-selected-list", {
          slidesPerView: 1,
          spaceBetween: 25,
          pagination: {
            el: "._repo-selected-swiper-pagination",
            clickable: true,
          },
          autoplay: {
            delay: 4500,
            disableOnInteraction: false,
          },
          breakpoints: {
            768: {
              slidesPerView: 3,
            },
            1024: {
              slidesPerView: 4,
            },
            1200: {
              slidesPerView: 4,
            },
            1600: {
              slidesPerView: 4,
            }
          },
        });
        selectedRepoEl.innerHTML = html;
        this.swiperHandler.updateSlides();
        this.swiperHandler.updateProgress();
        LetterAvatar.transform();
      }
    }
  },
  mounted() {
    if (this.static) {
      this.renderRepos(this.staticSwiperData);
    } else {
      getHomePageData().then(res => {
        this.renderRepos(res.data.repo);
      }).catch(err => {
        console.log(err);
      });
    }
  },
};
</script>

<style scoped lang="less">
.repo-selected-bg {
  background: url("data:image/svg+xml;charset=utf8,%3Csvg%20xmlns%3D%22http%3A%2F%2Fwww.w3.org%2F2000%2Fsvg%22%20xmlns%3Axlink%3D%22http%3A%2F%2Fwww.w3.org%2F1999%2Fxlink%22%20version%3D%221.1%22%3E%3Cdefs%3E%3ClinearGradient%20id%3D%221%22%20x1%3D%220%22%20x2%3D%221%22%20y1%3D%220%22%20y2%3D%220%22%20gradientTransform%3D%22matrix(0.11899999999999993%2C%201.217%2C%20-0.24039506172839506%2C%200.11899999999999993%2C%200.269%2C%20-0.22)%22%3E%3Cstop%20stop-color%3D%22%23ffffff%22%20stop-opacity%3D%220.47%22%20offset%3D%220%22%3E%3C%2Fstop%3E%3Cstop%20stop-color%3D%22%23e5e7eb%22%20stop-opacity%3D%220.3%22%20offset%3D%221%22%3E%3C%2Fstop%3E%3C%2FlinearGradient%3E%3C%2Fdefs%3E%3Crect%20width%3D%22100%25%22%20height%3D%22100%25%22%20fill%3D%22url(%231)%22%3E%3C%2Frect%3E%3C%2Fsvg%3E");
}

._repo_title {
  font-size: 18px;
  color: rgb(16, 16, 16);
  text-align: center;
  margin-bottom: 1em;
  font-weight: bold;
}

._repo-selected-list {
  overflow: hidden;
  padding: 1em 1em 3em 1em;
  text-align: left;
  position: relative;
}

/deep/._repo-selected-swiper-pagination .swiper-pagination-bullet {
  width: 8px;
  height: 8px;
  border-radius: 100%;
  background: #76cbed;
}

/deep/._repo-selected-swiper-pagination .swiper-pagination-bullet-active {
  width: 40px;
  border-radius: 4px;
}

/deep/ ._repo-selected-list .card {
  border-radius: 6px;
  background-color: #FFF;
  box-shadow: 0px 5px 10px 0px rgba(105, 192, 255, .3);
  border: 1px solid rgba(105, 192, 255, .4);
}

/deep/ ._repo-selected-list .header {
  line-height: 40px !important;
}
</style>
