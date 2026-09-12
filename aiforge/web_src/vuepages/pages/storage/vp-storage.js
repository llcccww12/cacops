import Vue from 'vue';
import ElementUI from 'element-ui';
import 'element-ui/lib/theme-chalk/index.css';
import localeEn from 'element-ui/lib/locale/lang/en';
import localeZh from 'element-ui/lib/locale/lang/zh-CN';
import { i18n, lang } from '~/langs';
import App from './index.vue';

Vue.use(ElementUI, {
  locale: lang === 'zh-CN' ? localeZh : localeEn
});

new Vue({
  i18n,
  render: (h) => h(App),
}).$mount('#__vue-root');
