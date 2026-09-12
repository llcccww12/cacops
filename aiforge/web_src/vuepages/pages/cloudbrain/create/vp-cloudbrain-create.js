import Vue from 'vue';
import ElementUI from 'element-ui';
import 'element-ui/lib/theme-chalk/index.css';
import localeEn from 'element-ui/lib/locale/lang/en';
import localeZh from 'element-ui/lib/locale/lang/zh-CN';
import { i18n, lang } from '~/langs';
import App from './index.vue';
function isMobileDevice() {
  return /Android|webOS|iPhone|iPad|iPod|BlackBerry|IEMobile|Opera Mini/i.test(navigator.userAgent);
}
Vue.prototype.$isMobile = isMobileDevice()
Vue.use(ElementUI, {
  locale: lang === 'zh-CN' ? localeZh : localeEn,
  size: 'small',
});

new Vue({
  i18n,
  render: (h) => h(App),
}).$mount('#__vue-root');
