import Vue from 'vue'
import App from './App.vue'
// 导入element-ui组件库
import ElementUI from 'element-ui';
// 导入element-ui组件库的样式
import 'element-ui/lib/theme-chalk/index.css';
// 注意：element-ui组件库，是一个插件，所有的插件都要由Vue去use
import VueClipboard from 'vue-clipboard2';
Vue.use(ElementUI);
Vue.config.productionTip = false
Vue.use(VueClipboard);
new Vue({
  render: h => h(App),
}).$mount('#app')
