<template>
    <div style="margin-bottom: 14px;min-height: 125px;">
        <div v-show="isLoading">
          <div class="ui active centered inline indeterminate text loader">{{$t('dashboard.heatmapLoading')}}</div>
        </div>
        <div class="total-contributions" v-if="!isLoading">
          <div class="rect"></div>
          <span>{{ $t('dashboard.contributionCount',{count:totalContributions}) }}</span>
        </div>
        <calendar-heatmap v-show="!isLoading" :locale="locale" :no-data-text="locale.no_contributions" :tooltip-unit="locale.contributions" :end-date="endDate" :values="values" :range-color="colorRange"/>
    </div>
</template>

<script>
import {CalendarHeatmap} from 'vue-calendar-heatmap';
import { getHeatMap } from "~/apis/modules/common";

export default {
    name: "ActivityHeatmap",
    components: {
        CalendarHeatmap
    },
    data() {
        return {
            isLoading: true,
            colorRange: ['#f4f4f4','#d8efbf','#9fdb81','#66c74b','#609926','#025900'],
            endDate: null,
            values: [],
            totalContributions: 0,
            user: '',
            locale: {
                contributions: 'contributions',
                no_contributions: 'No contributions',
            },
        };
    },
    mounted() {
        const metaEl = document.querySelector('meta[name="_uid"]');
        if (metaEl) {
          this.user = metaEl.getAttribute('content-ext');
        }
        this.endDate = new Date();
        this.loadHeatmap(this.user);
    },
    methods: {
        async loadHeatmap(userName) {
          // 将上述请求修改成下面的的代码
          try {
            const response = await getHeatMap(userName);
            const chartRawData = response.data; // 假设getHeatMap返回的数据结构包含data字段
            const chartData = [];
            let totalContributions = 0;
            
            for (let i = 0; i < chartRawData.length; i++) {
              totalContributions += chartRawData[i].contributions;
              chartData.push({
                date: new Date(chartRawData[i].timestamp * 1000),
                count: chartRawData[i].contributions
              });
            }
            
            this.totalContributions = totalContributions;
            this.values = chartData;
            this.isLoading = false;
            
          } catch(error) {
            console.error('获取热力图数据失败:', error);
            this.isLoading = false;
          }
        }
    },
}
</script>

<style lang="less" scoped>
.total-contributions{
  display: flex;
  height: 23px;
  line-height: 23px;
  color: rgba(16,16,16,1);
  font-size: 16px;
  font-weight: bold;
  align-items: center;
  margin-bottom: 16px;
  gap: 14px;
  .rect{
    width: 10px;
    height: 20px;
    border-radius: 20px;
    background-color: rgba(73,115,223,1);
  }
}
</style>
