import { createApp } from 'vue';
import App from './App.vue';
import router from './router';
import { store, key } from './store';
import { directive } from '/@/utils/directive';
import { i18n } from '/@/i18n/index';
import { globalComponentSize } from '/@/utils/componentSize';
import axios from 'axios';
import ElementPlus from 'element-plus';
import 'element-plus/dist/index.css';
import '/@/theme/index.scss';
import mitt from 'mitt';
import screenShort from 'vue-web-screen-shot';
import VueGridLayout from 'vue-grid-layout';
import { getI18nSource } from '/@/utils/i18n'
import { getJsonConfigs } from '/@/utils/jsonConfigs'
import VueClipboard from 'vue3-clipboard'

const app = createApp(App);
app
	.use(router)
	.use(store, key)
	.use(ElementPlus, { i18n: i18n.global.t, size: globalComponentSize })
	.use(i18n)
	.use(screenShort, { enableWebRtc: false })
	.use(VueGridLayout)
	.use(VueClipboard, {autoSetContainer: true,appendToBody: true,})
	.mount('#app');

app.config.globalProperties.mittBus = mitt();
app.config.globalProperties.$axios = axios;
app.config.globalProperties.$i = getI18nSource;
app.config.globalProperties.$cfg = getJsonConfigs;

directive(app);
