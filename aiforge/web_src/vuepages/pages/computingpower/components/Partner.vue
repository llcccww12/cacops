<template>
  <div class="partner-c">
    <div class="title-c">
      <div class="l"><img src="/img/computingpower/title-decoration.png" alt=""></div>
      <div class="tit">{{ $t('computingPowerObj.computingPowerPartner') }}</div>
      <div class="r"><img src="/img/computingpower/title-decoration.png" alt=""></div>
    </div>
    <div class="part-title-c">
      <div class="part-title">
        <div class="l"></div>
        <div class="tit">{{ $t('computingPowerObj.computingPowerOperationPlatform') }}</div>
        <div class="r"></div>
      </div>
    </div>
    <div class="computing-platform-c">
      <div class="item" v-for="(item, index) in computingPlatform" :key="index">
        <img :src="item.logo" :style="item.height ? `height:${item.height}px` : ''" alt="">
      </div>
    </div>
    <div class="part-title-c">
      <div class="part-title">
        <div class="l"></div>
        <div class="tit">{{ $t('computingPowerObj.chipManufacturers') }}</div>
        <div class="r"></div>
      </div>
    </div>
    <div class="chip-vendor-c">
      <div class="item" v-for="(item, index) in chipVendor" :key="index">
        <img :src="item.logo" :style="item.height ? `height:${item.height}px` : ''" alt="">
      </div>
    </div>
    <div class="part-title-c">
      <div class="part-title">
        <div class="l"></div>
        <div class="tit">{{ $t('resourcesManagement.aiCenter') }}</div>
        <div class="r"></div>
      </div>
    </div>
    <div class="ai-center-c">
      <div class="item" v-for="(item, index) in aiCenter" :key="index">
        <img :src="item.logo" alt="">
      </div>
    </div>
  </div>
</template>

<script>

import { getPartners } from '~/apis/modules/computingpower';

export default {
  name: "Partner",
  props: {},
  data() {
    return {
      computingPlatform: [],
      chipVendor: [],
      aiCenter: [],
    };
  },
  methods: {},
  mounted() {
    getPartners({}).then(res => {
      res = res.data;
      if (res.Code == 0) {
        const data = res.Data;
        this.computingPlatform = (data.computing_platform || []);
        this.chipVendor = (data.chip_vendor || []);
        this.aiCenter = (data.ai_center || []);
      }
    }).catch(err => {
      console.log(err);
    })
  },
};
</script>
<style scoped lang="less">
.partner-c {
  margin-top: 50px;

  .title-c {
    display: flex;
    align-items: center;
    justify-content: center;
    margin-bottom: 32px;

    .l {
      width: 162px;
      height: 40px;

      img {
        height: 100%;
        width: 100%;
      }
    }

    .tit {
      font-size: 28px;
    }

    .r {
      width: 162px;
      height: 40px;

      img {
        height: 100%;
        width: 100%;
        transform: rotate(180deg);
      }
    }
  }

  .part-title-c {
    display: flex;
    align-items: center;
    justify-content: center;
    margin-bottom: 15px;

    .part-title {
      display: flex;
      align-items: center;
      justify-content: center;
      width: 268px;

      .tit {
        font-size: 16px;
        color: rgba(16, 16, 16, 0.5);
        margin: 0 16px;
      }

      .l,
      .r {
        flex: 1;
        height: 0;
        width: 268px;
        border-bottom: 1px solid rgb(187, 187, 187);
      }
    }
  }

  .computing-platform-c {
    display: flex;
    align-items: center;
    justify-content: center;
    gap: 0 70px;
    margin-bottom: 20px;
    flex-wrap: wrap;

    .item {
      display: flex;
      height: 60px;
      margin-bottom: 23px;
      align-items: center;
      justify-content: center;

      img {
        height: 60px;
      }
    }
  }

  .chip-vendor-c {
    display: flex;
    align-items: center;
    justify-content: center;
    gap: 0 35px;
    margin-bottom: 22px;
    flex-wrap: wrap;

    .item {
      height: 60px;
      margin-bottom: 23px;
      display: flex;
      align-items: center;
      justify-content: center;

      img {
        height: 55px;
      }
    }
  }

  .ai-center-c {
    display: flex;
    align-items: center;
    justify-content: center;
    gap: 15px 4px;
    margin-bottom: 20px;
    flex-wrap: wrap;

    .item {
      height: 65px;
      margin-bottom: 0px;
      align-items: center;
      justify-content: center;

      img {
        height: 100%;
      }
    }
  }
}
</style>
