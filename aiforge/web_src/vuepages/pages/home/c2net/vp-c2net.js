import Vue from 'vue';
import App from './App.vue';
import Index from './index.vue';
import { i18n } from '~/langs'
import * as echarts from 'echarts/core'
import { MapChart } from 'echarts/charts'
import { CanvasRenderer } from 'echarts/renderers';
import { VisualMapComponent } from 'echarts/components';
import china from './china.json'
echarts.use([MapChart, CanvasRenderer, VisualMapComponent])
echarts.registerMap('china', china)
Vue.prototype.$echarts = echarts
let UA = navigator.userAgent
let View = App
if (/(Android|webOS|iPhone|iPod|tablet|BlackBerry|Mobile)/i.test(UA))
      View = Index

new Vue({
  i18n,
  render: (h) => h(View),
}).$mount('#__vue-root');
