import { lang, i18n } from '~/langs';
import langMsg from './config';

i18n.mergeLocaleMessage(lang, { $local: langMsg[lang] });
