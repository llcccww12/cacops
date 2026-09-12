import Vue from 'vue'
import App from './Index.vue'
// import router from './router'
import localeEn from 'element-ui/lib/locale/lang/en';
import localeZh from 'element-ui/lib/locale/lang/zh-CN';
import  { i18n, lang } from "~/langs";
import ElementUI from 'element-ui';
import 'element-ui/lib/theme-chalk/index.css';
// import './style/index.css'

Vue.use(ElementUI, { locale: lang === 'zh-CN' ? localeZh : localeEn, 
}),
Vue.config.productionTip = false




new Vue({
  i18n,
//   router,
  render: h => h(App)
}).$mount('#__vue-root');
