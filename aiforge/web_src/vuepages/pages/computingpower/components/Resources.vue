<template>
  <div class="resources-c">
    <el-checkbox class="filter-mobile-show" v-model="filterShow">
      {{ $t('computingPowerObj.displayFilteringConditions') }}
    </el-checkbox>
    <div v-if="filterShow" class="conds-c">
      <div class="conds-item">
        <div class="conds-item-tit">{{ $t('resourcesManagement.computeResource') }}：</div>
        <div class="conds-item-content">
          <div class="sel-item" :class="item.k == conds.compute_resource ? 'active' : ''"
            v-for="(item, index) in computeResourceList" :key="index" @click="changeConds(item.k, 'compute_resource')">
            {{ item.v }}
          </div>
        </div>
      </div>
      <div class="conds-item">
        <div class="conds-item-tit">{{ $t('resourcesManagement.accCardType') }}：</div>
        <div class="conds-item-content">
          <div class="sel-item" :class="item.k == conds.card_type ? 'active' : ''" v-for="(item, index) in cardTypeList"
            :key="index" @click="changeConds(item.k, 'card_type')">
            {{ item.v }}
          </div>
        </div>
      </div>
      <div class="conds-item">
        <div class="conds-item-tit">{{ $t('computingPowerObj.accCardCount') }}：</div>
        <div class="conds-item-content">
          <div class="sel-item" :class="item.k == conds.acc_cards_num ? 'active' : ''"
            v-for="(item, index) in cardNumList" :key="index" @click="changeConds(item.k, 'acc_cards_num')">
            {{ item.v }}
          </div>
        </div>
      </div>
      <div class="conds-item">
        <div class="conds-item-tit">{{ $t('resourcesManagement.aiCenter') }}：</div>
        <div class="conds-item-content">
          <div class="sel-item" :class="item.k == conds.ai_center ? 'active' : ''" v-for="(item, index) in aiCenterList"
            :key="index" @click="changeConds(item.k, 'ai_center')">
            {{ item.v }}
          </div>
        </div>
      </div>
      <div class="conds-item">
        <div class="conds-item-tit">{{ $t('computingPowerObj.priceRange') }}：</div>
        <div class="conds-item-content">
          <el-input class="price-s" v-model="conds.price_start" @input="inputPrice('price_start')"
            @blur="checkPrice('price_start')"></el-input><span class="to"> - </span><el-input v-model="conds.price_end"
            @input="inputPrice('price_end')" @blur="checkPrice('price_end')" class="price-e"></el-input><span
            class="unit">{{ $t('computingPowerObj.point_hr') }}</span>
          <el-button type="primary" size="mini" v-if="conds.price_start != '' || conds.price_end != ''"
            @click="clearPriceConds">{{ $t('clear') }}</el-button>
        </div>
      </div>
    </div>
    <div class="content-c">
      <div class="list-c" v-loading="loading">
        <div class="list-item" v-for="(item, index) in list" :key="index">
          <div class="top">
            <div class="left">
              <div class="title">
                <div class="name">{{ item.AccCardTypeShow || item.AccCardType }}</div>
              </div>
            </div>
            <div class="right">
              <span class="price">{{ item.UnitPrice.toFixed(2) }}</span>{{ $t('computingPowerObj.point_hr') }}
              <div class="right use-position-mobile-show">
                <JumpButton :title="$t('computingPowerObj.use')" size="small"
                  :tooltip="$t('computingPowerObj.createComputingTask')" @click="goAiTask(item)" />
              </div>
            </div>
          </div>
          <div class="mid">
            <div class="left">
              <div class="attributes">
                <div class="attribute">
                  <span class="tit">{{ item.ComputeResource }}：</span><span class="val">{{ item.AccCardsNum }} * {{
                    item.AccCardTypeShow || item.AccCardType }}<span v-if="item.GPUMemGiB"><span class="val">({{
                      $t('resourcesManagement.gpuMem') }}:</span><span class="val">{{ item.GPUMemGiB
                        }}GB)</span>
                    </span>
                  </span>
                </div>
                <div class="attribute" v-if="item.CpuCores">
                  <span class="tit">CPU：</span><span class="val">{{ item.CpuCores }}</span>
                </div>
                <div class="attribute" v-if="item.MemGiB">
                  <span class="tit">{{ $t('memory') }}：</span><span class="val">{{ item.MemGiB }}GB</span>
                </div>
              </div>
            </div>
            <div class="right"></div>
          </div>
          <div class="bottom">
            <div class="left">
              <div class="ai-center" :title="center.AiCenterName" v-for="(center) in item.AICenterList"
                :key="center.AiCenterCode">
                {{ center.AiCenterName }}
              </div>
            </div>
            <div class="right use-position-mobile-hide">
              <JumpButton :title="$t('computingPowerObj.use')" size="small"
                :tooltip="$t('computingPowerObj.createComputingTask')" @click="goAiTask(item)" />
            </div>
          </div>
        </div>
        <div class="demand-item no-data" v-if="(!list.length && !loading)">
          <div class="item-empty">
            <div class="item-empty-icon"></div>
            <div class="item-empty-tips">{{ $t('modelObj.model_square_empty') }}</div>
          </div>
        </div>
      </div>
    </div>
    <div class="pagination-c">
      <el-pagination background layout="total, sizes, prev, pager, next, jumper" :current-page.sync="conds.page"
        :page-size.sync="conds.page_size" :page-sizes="paginationInfo.pageSizes" :total="paginationInfo.total"
        @current-change="currentChange" @size-change="sizeChange">
      </el-pagination>
    </div>
  </div>
</template>

<script>
import JumpButton from './JumpButton.vue';
import { getAccCardList, getAvailableAiCenterList, getResourceList } from '~/apis/modules/computingpower';
import { ACC_CARD_TYPE } from '~/const';
import { getListValueWithKey } from '~/utils';

export default {
  name: "Resources",
  props: {
    condtions: { type: Object, default: () => ({}) },
    active: { type: Boolean, defalut: false },
  },
  components: { JumpButton },
  data() {
    return {
      filterShow: true,
      computeResourceList: [{ k: '', v: this.$t('all') }],
      cardTypeList: [{ k: '', v: this.$t('all') }],
      cardNumList: [{ k: '', v: this.$t('all') }, ...[1, 2, 4, 8].map(itm => ({ k: itm, v: itm })), { k: 9999, v: this.$t('others') }],
      aiCenterList: [{ k: '', v: this.$t('all') }],
      computeResourceMap: {},
      cardTypeMap: {},
      conds: {
        compute_resource: '',
        card_type: '',
        acc_cards_num: '',
        ai_center: '',
        price_start: '',
        price_end: '',
        page: 1,
        page_size: 15,
      },
      _price_start: '',
      _price_end: '',
      list: [],
      paginationInfo: {
        pageSizes: [15, 30, 60],
        total: 0,
      },
      loading: false,
    };
  },
  computed: {},
  watch: {
    active: {
      handler(newValue, oValue) {
        if (newValue) {
          this.search();
        }
      }
    }
  },
  methods: {
    changeConds(value, type) {
      if (type == 'compute_resource') {
        this.cardTypeList.splice(1, Infinity);
        for (let key in this.cardTypeMap) {
          if (value && value != key) continue;
          this.cardTypeMap[key].forEach(item => {
            this.cardTypeList.push({
              k: item,
              v: getListValueWithKey(ACC_CARD_TYPE, item),
            });
          })
        }
        this.conds['card_type'] = '';
      }
      this.conds[type] = value;
      this.conds.page = 1;
      this.search();
    },
    inputPrice(type) {
      if (!Number.isInteger(Number(this.conds[type])) || this.conds[type].length > 3) {
        this.conds[type] = this.conds[type].slice(0, this.conds[type].length - 1);
      }
    },
    checkPrice(type) {
      if (this.conds.price_start && this.conds.price_end && Number(this.conds.price_start) > Number(this.conds.price_end)) {
        this.conds[type] = '';
      }
      if (this._price_start != this.conds.price_start || this._price_end != this.conds.price_end) {
        this.conds.page = 1;
        this.search();
      }
      this._price_start = this.conds.price_start;
      this._price_end = this.conds.price_end;
    },
    clearPriceConds() {
      this._price_start = this.conds.price_start = '';
      this._price_end = this.conds.price_end = '';
      this.conds.page = 1;
      this.search();
    },
    currentChange(page) {
      this.conds.page = page;
      this.search();
    },
    sizeChange(pageSize) {
      this.conds.page_size = pageSize;
      this.search();
    },
    search() {
      const params = {
        page: this.conds.page,
        pageSize: this.conds.page_size,
        resource: this.conds.compute_resource === '' ? '' : this.computeResourceMap[this.conds.compute_resource] || [this.conds.compute_resource],
        accCardType: this.conds.card_type,
        accCardNum: this.conds.acc_cards_num === '' ? -1 : Number(this.conds.acc_cards_num),
        excludeAccCardNums: this.conds.acc_cards_num == 9999 ? this.cardNumList.slice(1, this.cardNumList.length - 1).map(itm => itm.k).join('|') : undefined,
        centerCode: this.conds.ai_center,
        minPrice: this.conds.price_start === '' ? -1 : Number(this.conds.price_start),
        maxPrice: this.conds.price_end === '' ? -1 : Number(this.conds.price_end),
      };
      // console.log('search conds', this.conds);
      // console.log('search params', params);
      this.loading = true;
      getResourceList(params).then(res => {
        this.loading = false;
        res = res.data;
        if (res.code == 0) {
          const data = res.data || {};
          this.paginationInfo.total = data.total;
          this.list = (data.list || []).map(item => {
            return {
              AccCardTypeShow: getListValueWithKey(ACC_CARD_TYPE, item.AccCardType),
              ...item,
            }
          });
        }
      }).catch(err => {
        this.loading = false;
        console.log(err);
      });
    },
    goAiTask(item) {
      const spec = {
        ComputeSource: item.ComputeResource,
        AccCardsNum: item.AccCardsNum,
        AccCardType: item.AccCardType,
        CpuCores: item.CpuCores,
        MemGiB: item.MemGiB,
        GPUMemGiB: item.GPUMemGiB,
        ShareMemGiB: item.ShareMemGiB,
      }
      window.location.href = `/cloudbrains/create?spec=${encodeURIComponent(JSON.stringify(spec))}`;
    },
    setConds() {
      for (let key in this.conds) {
        if (this.condtions[key]) {
          this.conds[key] = this.condtions[key];
        }
      }
      this._price_start = this.conds.price_start || '';
      this._price_end = this.conds.price_end || '';
    }
  },
  created() {
    // this.setConds();
    getAccCardList().then(res => {
      res = res.data;
      if (res.code == 0) {
        this.computeResourceList.splice(1, Infinity);
        this.cardTypeList.splice(1, Infinity);
        const data = res.data.list || [];
        for (let i = 0, iLen = data.length; i < iLen; i++) {
          const item = data[i];
          const computeSource = item.ComputeSource;
          const cardList = item.CardList;
          const computeSourceKey = computeSource.indexOf('-GPGPU') > 0 ? 'GPGPU' : computeSource;
          if (this.computeResourceMap[computeSourceKey]) {
            this.computeResourceMap[computeSourceKey].push(computeSource);
          } else {
            this.computeResourceMap[computeSourceKey] = [computeSource];
          }
          if (this.cardTypeMap[computeSourceKey]) {
            this.cardTypeMap[computeSourceKey].push(...cardList);
          } else {
            this.cardTypeMap[computeSourceKey] = cardList;
          }
        }
        for (let key in this.computeResourceMap) {
          this.computeResourceMap[key] = Array.from(new Set(this.computeResourceMap[key]));
          this.computeResourceList.push({
            k: key,
            v: key,
          });
        }
        for (let key in this.cardTypeMap) {
          this.cardTypeMap[key] = Array.from(new Set(this.cardTypeMap[key]));
          this.cardTypeMap[key].forEach(item => {
            this.cardTypeList.push({
              k: item,
              v: getListValueWithKey(ACC_CARD_TYPE, item),
            });
          })
        }
      }
    }).catch(err => {
      console.log(err);
    });
    getAvailableAiCenterList().then(res => {
      res = res.data;
      if (res.code == 0) {
        this.aiCenterList.splice(1, Infinity);
        const data = res?.data?.list || [];
        for (let i = 0, iLen = data.length; i < iLen; i++) {
          const item = data[i];
          this.aiCenterList.push({
            k: item.AiCenterCode,
            v: item.AiCenterName,
          });
        }
      }
    }).catch(err => {
      console.log(err);
    });
  },
  mounted() {
    this.search();
    const element = document.querySelector('.filter-mobile-show');
    if (element) {
      const displayValue = window.getComputedStyle(element).getPropertyValue('display');
      // 判断是否为移动端，当为移动端时，默认不显示筛选条件
      if (displayValue === 'block') {
        this.filterShow = false
      }
    } else {
      console.log('Element not found.');
    }
  },
};
</script>

<style scoped lang="less">
.filter-mobile-show {
  display: none;
}

.use-position-mobile-show {
  display: none;
}

@media only screen and (max-width: 767px) {
  .filter-mobile-show {
    display: block;
    margin-bottom: 20px;
  }

  .conds-item-tit {
    justify-content: space-between !important;
  }

  .attributes {
    display: block !important;
  }

  .use-position-mobile-show {
    display: block;
  }

  .use-position-mobile-hide {
    display: none !important;
  }

  .available {
    color: rgba(39, 177, 72, 1);
    font-size: 12px;
  }

  .apply {
    color: rgba(50, 145, 248, 1);
    font-size: 12px;
  }

  .list-item {
    width: 100% !important;
  }
}

.resources-c {
  margin-top: 20px;

  &.zh {
    .conds-c .conds-item .conds-item-tit {
      width: 100px;
    }
  }
}

.conds-c {
  .conds-item {
    display: flex;

    .conds-item-tit {
      width: 175px;
      display: flex;
      padding-top: 7px;
      justify-content: flex-end;
      padding-right: 10px;
    }

    .conds-item-content {
      display: flex;
      flex-wrap: wrap;
      width: 0;
      flex: 1;
      align-items: center;

      .sel-item {
        display: flex;
        align-items: center;
        justify-content: center;
        height: 30px;
        border-radius: 4px;
        background-color: rgba(248, 249, 250, 1);
        color: rgba(65, 80, 88, 1);
        font-size: 12px;
        padding: 0 14px;
        cursor: pointer;
        margin-right: 12px;
        margin-bottom: 16px;

        &.active {
          background-color: rgba(3, 102, 214, 1);
          color: rgba(255, 255, 255, 1);
        }
      }

      .price-s,
      .price-e {
        width: 52px;

        /deep/.el-input__inner {
          text-align: center;
          padding: 0 8px;
          height: 30px;
        }
      }

      .to,
      .unit {
        margin: 0 10px;
      }

      @media screen and (max-width: 403px) {
        .el-button {
          margin-top: 5px;
        }
      }
    }
  }
}

.content-c {
  margin-top: 20px;

  .list-c {
    display: flex;
    flex-wrap: wrap;
    gap: 30px 20px;

    .list-item {
      padding: 18px 22px;
      border-radius: 15px;
      background-color: rgba(255, 255, 255, 1);
      box-shadow: 0px 5px 10px 0px rgba(157, 197, 226, 0.2);
      border: 1px solid rgba(157, 197, 226, 0.4);
      width: 32%;
      min-height: 160px;

      .top {
        display: flex;

        .left {
          width: 0;
          flex: 1;

          .title {
            display: flex;
            align-items: center;

            .name {
              color: rgba(16, 16, 16, 1);
              font-size: 16px;
              margin-right: 10px;
              font-weight: bold;
            }

            .type {
              display: flex;
              align-items: center;
              justify-content: center;
              padding: 0 4px;
              height: 18px;
              border-radius: 2px;
              background-color: rgba(91, 185, 115, 1);
              color: rgba(251, 251, 251, 1);
              font-size: 12px;
              text-align: center;

              &.exclusive {
                background-color: rgba(50, 145, 248, 1);
              }
            }
          }
        }

        .right {
          width: 130px;
          padding-top: 0px;
          text-align: right;
          font-size: 14px;

          .price {
            color: rgb(64, 123, 237);
            font-size: 20px;
            margin-right: 2px;
          }
        }
      }

      .mid {
        padding-top: 6px;
        display: flex;
        align-items: flex-end;
        min-height: 60px;

        .left {
          width: 0;
          flex: 1;
          display: flex;
          flex-wrap: wrap;

          .attributes {
            display: flex;
            flex-direction: column;
            min-height: 76px;

            .attribute {
              margin: 4px 0;

              &:first-child {
                margin-top: 0;
              }

              &:last-child {
                margin-bottom: 0;
              }

              .tit {
                color: rgb(136, 136, 136);
              }
            }
          }
        }

        .right {
          width: 5px;
          display: flex;
          align-items: center;
          justify-content: flex-end;
        }
      }

      .bottom {
        border-top: 1px solid rgba(157, 197, 226, 0.2);
        margin-top: 14px;
        padding-top: 6px;
        display: flex;

        .left {
          width: 0;
          flex: 1;
          display: flex;
          flex-wrap: wrap;

          .ai-center {
            padding: 0 12px;
            margin-right: 10px;
            margin-top: 10px;
            height: 30px;
            line-height: 30px;
            border-radius: 4px;
            background-color: rgba(50, 145, 248, 0.1);
            color: rgba(50, 145, 248, 1);
            font-size: 12px;
            text-align: center;
            border: 1px solid rgba(50, 145, 248, 0.6);
            max-width: 90%;
            text-overflow: ellipsis;
            white-space: nowrap;
            overflow: hidden;
          }
        }

        .right {
          width: 100px;
          display: flex;
          justify-content: flex-end;
          padding-top: 10px;
        }
      }
    }
  }
}

.no-data {
  display: flex;
  justify-content: center;
  padding: 0 0;
  width: 100%;

  .item-empty {
    height: 391px;
    width: 100%;
    overflow: hidden;
    padding: 15px;
    background: transparent;
    display: flex;
    flex-direction: column;
    justify-content: center;
    background-color: rgba(245, 245, 246, 0.5);

    .item-empty-icon {
      height: 80px;
      width: 100%;
      background: url(/img/empty-box.svg) center center no-repeat;
    }

    .item-empty-tips {
      text-align: center;
      margin-top: 20px;
      font-size: 18px;
      color: rgb(63, 63, 64);
    }
  }
}

.pagination-c {
  text-align: center;
  margin: 10px 0;
}
</style>
