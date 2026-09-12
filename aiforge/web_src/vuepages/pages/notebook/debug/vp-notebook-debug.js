import Vue from 'vue';
import {Dialog,Loading} from 'element-ui';
import 'element-ui/lib/theme-chalk/index.css';
import { i18n, lang } from '~/langs';
import App from './index.vue';

Vue.use(Dialog)
Vue.use(Loading.directive)
new Vue({
  i18n,
  render: (h) => h(App),
}).$mount('#__vue-root');
