import Vue from 'vue';
import App from './index.vue';
import { i18n} from '~/langs'
new Vue({
  i18n,
  render: (h) => h(App),
}).$mount('#__vue-root');