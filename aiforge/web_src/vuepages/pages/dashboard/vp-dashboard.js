import Vue from 'vue';
import ElementUI from 'element-ui';
import 'element-ui/lib/theme-chalk/index.css';
import { i18n } from '~/langs';
import App from './App.vue';
import Index from './index.vue';
Vue.use(ElementUI);
let UA = navigator.userAgent
let View = App
console.log("==========",/(Android|webOS|iPhone|iPod|tablet|BlackBerry|Mobile|IEMobile)/i.test(UA))
if (/(Android|webOS|iPhone|iPod|tablet|BlackBerry|Mobile|IEMobile)/i.test(UA)) {
  console.log("mobile")
  View = Index
}
new Vue({
  i18n,
  render: (h) => h(View),
}).$mount('#__vue-root');