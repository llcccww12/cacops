import Vue from 'vue';
import ElementUI from 'element-ui';
import 'element-ui/lib/theme-chalk/index.css';
import localeEn from 'element-ui/lib/locale/lang/en';
import localeZh from 'element-ui/lib/locale/lang/zh-CN';
import { i18n, lang } from '~/langs';
import './langs';
import { setWebpackPublicPath } from '~/utils';
import App from './App.vue';
import router from './router';

setWebpackPublicPath();
// 初始检测设备类型
Vue.prototype.$isMobile = window.innerWidth <= 768;

// 监听窗口大小变化
window.addEventListener('resize', () => {
  Vue.prototype.$isMobile = window.innerWidth <= 768;
});
Vue.use(ElementUI, {
  locale: lang === 'zh-CN' ? localeZh : localeEn,
  size: 'small',
});

new Vue({
  i18n,
  router,
  render: (h) => h(App),
}).$mount('#__vue-root');
