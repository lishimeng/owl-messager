import {createApp} from 'vue';
import pinia from '/src/stores';
import App from '/src/App.vue';
import router from '/src/router';
import {directive} from '/src/directive';
import {i18n} from '/src/i18n';
import other from '/src/utils/other';
import 'font-awesome/css/font-awesome.min.css'
import ElementPlus from 'element-plus';
import '/src/theme/index.scss';
import VueGridLayout from 'vue-grid-layout';

const app = createApp(App);
directive(app);
other.elSvg(app);
 app.use(pinia)
    .use(router)
    .use(ElementPlus)
    .use(i18n)
    .use(VueGridLayout).mount('#app');
