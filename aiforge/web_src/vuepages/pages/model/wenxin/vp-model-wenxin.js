import Vue from 'vue';
import { i18n } from '~/langs';
import App from './index.vue';



new Vue({
  i18n,
  render: (h) => h(App),
}).$mount('#__vue-root');
