<template>
  <div class="resources-c" :class="isZh ? 'zh' : ''">
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
      <div class="conds-item" v-show="cardTypeList.length > 1">
        <div class="conds-item-tit">{{ $t('resourcesManagement.accCardType') }}：</div>
        <div class="conds-item-content">
          <div class="sel-item" :class="item.k == conds.card_type ? 'active' : ''" v-for="(item, index) in cardTypeList"
            :key="index" @click="changeConds(item.k, 'card_type')">
            {{ item.v }}
          </div>
        </div>
      </div>
      <div class="conds-item" v-show="cardNumList.length > 2">
        <div class="conds-item-tit">{{ $t('computingPowerObj.accCardCount') }}：</div>
        <div class="conds-item-content">
          <div class="sel-item" :class="item.k == conds.acc_cards_num ? 'active' : ''"
            v-for="(item, index) in cardNumList" :key="index" @click="changeConds(item.k, 'acc_cards_num')">
            {{ item.v }}
          </div>
        </div>
      </div>
      <div class="conds-item" v-show="providerList.length > 2">
        <div class="conds-item-tit">{{ $t('computingPowerObj.provider') }}：</div>
        <div class="conds-item-content">
          <div class="sel-item" :class="item.k == conds.provider ? 'active' : ''" v-for="(item, index) in providerList"
            :key="index" @click="changeConds(item.k, 'provider')">
            {{ item.v }}
          </div>
        </div>
      </div>
    </div>
    <div class="content-c">
      <div class="list-c" v-loading="loading">
        <div class="list-item" v-for="(item, index) in showList" :key="index">
          <div class="top">
            <div class="left">
              <div class="title">
                <div class="name">{{ item.name }}</div>
                <div v-if="item.source_type" class="type" :class="item.source_type == '独享' ? 'exclusive' : ''">{{
                  item.source_type }}</div>
              </div>
            </div>
            <div class="right">
              <span class="price">{{ item.price }}</span> {{ item.priceUnit }}
            </div>
          </div>
          <div class="mid">
            <div class="left">
              <div class="attributes">
                <div class="attribute" v-for="(attr, index) in ['cpu', 'gpu', 'npu', 'memory', 'storage']"
                  v-if="item[attr] !== undefined">
                  <span class="tit">{{ attrTitle[attr] }}：</span><span class="val">{{ item[attr] }}</span>
                </div>
              </div>
            </div>
            <div class="right"></div>
          </div>
          <div class="bottom">
            <div class="provider"> {{ item.provider }}</div>
            <div class="go-buy" @click="goBuy(item)">{{ $t('computingPowerObj.buyNow') }}</div>
          </div>
        </div>
        <div class="demand-item no-data" v-if="(!showList.length && !loading)">
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
import { getPayResourceList } from '~/apis/modules/computingpower';
import { lang } from '~/langs';

export default {
  name: "ExtraResources",
  props: {
    condtions: { type: Object, default: () => ({}) },
    active: { type: Boolean, defalut: false },
  },
  data() {
    return {
      isZh: lang == 'zh-CN',
      isLogin: false,
      filterShow: true,
      computeResourceList: [{ k: '', v: this.$t('all') }],
      cardTypeList: [{ k: '', v: this.$t('all') }],
      cardNumList: [{ k: '', v: this.$t('all') }],
      providerList: [{ k: '', v: this.$t('all') }],
      attrTitle: {
        cpu: 'CPU',
        gpu: this.$t('resourcesManagement.gpuMem'),
        npu: 'NPU',
        memory: this.$t('resourcesManagement.mem'),
        storage: this.$t('computingPowerObj.systemDisk'),
      },
      computeResourceMap: {},
      cardTypeMap: {},
      cardNumMap: {},
      providerMap: {},
      conds: {
        compute_resource: '',
        card_type: '',
        acc_cards_num: '',
        provider: '',
        page: 1,
        page_size: 15,
      },
      list: [],
      showList: [],
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
          this.getData();
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
            if (!this.cardTypeList.find(itm => itm.k == item)) {
              this.cardTypeList.push({
                k: item,
                v: item
              });
            }
          })
        }
        this.conds['card_type'] = '';
        this.cardNumList.splice(1, Infinity);
        for (let key in this.cardNumMap) {
          if (value && value != key) continue;
          this.cardNumMap[key].forEach(item => {
            if (!this.cardNumList.find(itm => itm.k == item)) {
              this.cardNumList.push({
                k: item,
                v: item
              });
            }
          })
        }
        this.cardNumList.push({
          k: 'other',
          v: this.$t('others')
        });
        this.conds['acc_cards_num'] = '';
      }
      this.conds[type] = value;
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
    goBuy(item) {
      if (this.isLogin) {
        window.open(item.purchaseLink, '_blank');
      } else {
        window.location.href = `/user/login?redirect_to=${encodeURIComponent(window.location.href + '?active=1')}`;
      }
    },
    getData() {
      this.loading = true
      this.computeResourceMap = {};
      this.cardTypeMap = {};
      this.cardNumMap = {};
      this.conds = {
        compute_resource: '',
        card_type: '',
        acc_cards_num: '',
        provider: '',
        page: 1,
        page_size: 15,
      };
      this.paginationInfo.total = 0;
      getPayResourceList({}).then(res => {
        res = res.data;
        if (res.Code == 0) {
          const resourceTypes = res.Data.resourceTypes || []
          this.computeResourceList.splice(1, Infinity);
          this.cardTypeList.splice(1, Infinity);
          this.cardNumList.splice(1, Infinity);
          for (let i = 0, iLen = resourceTypes.length; i < iLen; i++) {
            const resourceTypesI = resourceTypes[i]
            const computeSource = {
              k: resourceTypesI.category,
              v: resourceTypesI.label
            }
            const computeSourceKey = computeSource.k;
            this.computeResourceMap[computeSourceKey] = computeSource;
            const cardList = resourceTypesI.cardTypes || [];
            const cardCounts = resourceTypesI.cardCounts || [];
            const providers = resourceTypesI.providers || [];
            if (this.cardTypeMap[computeSourceKey]) {
              this.cardTypeMap[computeSourceKey].push(...cardList);
            } else {
              this.cardTypeMap[computeSourceKey] = cardList;
            }
            if (this.cardNumMap[computeSourceKey]) {
              this.cardNumMap[computeSourceKey].push(...cardCounts);
            } else {
              this.cardNumMap[computeSourceKey] = cardCounts;
            }
            if (this.providerMap[computeSourceKey]) {
              this.providerMap[computeSourceKey].push(...providers);
            } else {
              this.providerMap[computeSourceKey] = providers;
            }
          }
          for (let key in this.computeResourceMap) {
            const computeSource = this.computeResourceMap[key]
            this.computeResourceList.push(computeSource);
            this.computeResourceMap[key] = [computeSource.k]
          }
          for (let key in this.cardTypeMap) {
            this.cardTypeMap[key] = Array.from(new Set(this.cardTypeMap[key]));
            this.cardTypeMap[key].forEach(item => {
              if (!this.cardTypeList.find(itm => itm.k == item)) {
                this.cardTypeList.push({
                  k: item,
                  v: item,
                });
              }
            })
          }
          for (let key in this.cardNumMap) {
            this.cardNumMap[key] = Array.from(new Set(this.cardNumMap[key]));
            this.cardNumMap[key].forEach(item => {
              if (!this.cardNumList.find(itm => itm.k == item)) {
                this.cardNumList.push({
                  k: item,
                  v: item,
                });
              }
            })
          }
          this.cardNumList.push({
            k: 'other',
            v: this.$t('others')
          });
          for (let key in this.providerMap) {
            this.providerMap[key] = Array.from(new Set(this.providerMap[key]));
            this.providerMap[key].forEach(item => {
              if (!this.providerList.find(itm => itm.k == item)) {
                this.providerList.push({
                  k: item,
                  v: item,
                });
              }
            })
          }
          const dataList = res.Data.products || []
          this.list = dataList.map(item => {
            return {
              ...item,
              computeResource: item.category,
              cardType: item.cardType,
              cardNum: item.cardCount,
              provider: item.provider,
            }
          })
        }
        this.search()
      }).catch(err => {
        console.log(err);
      }).finally(() => {
        this.loading = false;
      })
    },
    search() {
      const params = {
        page: this.conds.page,
        pageSize: this.conds.page_size,
        resource: this.conds.compute_resource === '' ? '' : this.computeResourceMap[this.conds.compute_resource] || [this.conds.compute_resource],
        accCardType: this.conds.card_type,
        accCardNum: this.conds.acc_cards_num === '' ? -1 : this.conds.acc_cards_num,
        provider: this.conds.provider === '' ? '' : this.conds.provider,
      };
      // console.log('search conds', this.conds);
      // console.log('search params', params);
      this.loading = true;
      const useList = this.list.filter((item) => {
        let check1 = true, check2 = true, check3 = true, check4 = true;
        if (params.resource) {
          check1 = params.resource.includes(item.computeResource)
        }
        if (params.accCardType) {
          check2 = item.cardType == params.accCardType
        }
        if (params.accCardNum != -1) {
          if (params.accCardNum == 'other') {
            check3 = !this.cardNumList.slice(1, this.cardNumList.length - 1).map(itm => itm.k).includes(item.cardNum)
          } else {
            check3 = item.cardNum == params.accCardNum
          }
        }
        if (params.provider) {
          check4 = item.provider == params.provider
        }
        return check1 && check2 && check3 && check4
      })
      this.paginationInfo.total = useList.length;
      this.showList = useList.slice((params.page - 1) * params.pageSize, params.page * params.pageSize)
      this.loading = false;
    },
  },
  created() {
    this.isLogin = !!document.querySelector('meta[name="_uid"]');
    this.getData()
  },
  mounted() {
    // this.search();
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
      position: relative;
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
              font-size: 20px;
              margin-right: 10px;
              font-weight: bold;
              overflow: hidden;
              text-overflow: ellipsis;
              white-space: nowrap;
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
              min-width: 34px;

              &.exclusive {
                background-color: rgba(50, 145, 248, 1);
              }
            }
          }
        }

        .right {
          width: 130px;
          padding-top: 4px;
          text-align: right;
          font-size: 14px;

          .price {
            color: rgba(255, 98, 0, 1);
            font-size: 20px;
            font-weight: 700;
            margin-right: 2px;
          }
        }
      }

      .mid {
        padding-top: 6px;
        display: flex;
        align-items: flex-end;
        min-height: 100px;

        .left {
          width: 0;
          flex: 1;
          display: flex;
          flex-wrap: wrap;

          .attributes {
            display: flex;
            flex-direction: column;
            min-height: 104px;

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
          width: 75px;
          display: flex;
          align-items: center;
          justify-content: flex-end;
        }
      }

      .bottom {
        border-top: 1px solid rgba(157, 197, 226, 0.2);
        margin-top: 8px;
        padding-top: 14px;
        display: flex;
        align-items: center;
        justify-content: space-between;

        .provider {
          display: flex;
          align-items: center;
          justify-content: center;
          padding: 0 12px;
          height: 30px;
          border-radius: 4px;
          background-color: rgba(50, 145, 248, 0.1);
          color: #3291f8;
          font-size: 12px;
          text-align: center;
          border: 1px solid rgba(50, 145, 248, 0.6);
        }

        .go-buy {
          display: flex;
          align-items: center;
          justify-content: center;
          height: 30px;
          padding: 0 12px;
          border-radius: 4px;
          color: rgb(255, 255, 255);
          font-size: 12px;
          background: rgb(1, 72, 255);
          cursor: pointer;
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

@media only screen and (max-width: 767px) {
  .filter-mobile-show {
    display: block;
    margin-bottom: 20px;
  }

  .conds-item-tit {
    justify-content: space-between !important;
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

  .content-c {
    .list-c {
      gap: 20px 0;

      .list-item {
        width: 100%;
      }
    }
  }
}

@media only screen and (min-width: 768px) and (max-width: 1350px) {
  .content-c {
    .list-c {
      .list-item {
        width: 48%;
      }
    }
  }
}
</style>
