import Vue from 'vue';
import VueI18n from 'vue-i18n';
import zh from './config/zh-CN';
import en from './config/en-US';

Vue.use(VueI18n);

export const lang = window.config.lang;
export const i18n = new VueI18n({
  locale: lang,
  messages: {
    'zh-CN': zh,
    'en-US': en
  },
  silentTranslationWarn: true
});
