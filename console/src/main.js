import Vue from 'vue';
import iView from 'iview';
import dayjs from 'dayjs';
import lodash from 'lodash';
import {router} from './router/index';
import {appRouter} from './router/router';
import store from './store';
import App from './app.vue';
import 'iview/dist/styles/iview.css';
import util from './libs/util';
import Cookies from 'js-cookie';

Vue.prototype.$d = dayjs;
Vue.prototype._ = lodash;

Vue.use(iView);

let vm = new Vue({
  el: '#app',
  router: router,
  store: store,
  render: h => h(App),
  data: {
    currentPageName: ''
  },
  mounted () {
    this.currentPageName = this.$route.name;
    // 显示打开的页面的列表
    this.$store.commit('setOpenedList');
    this.$store.commit('initCachepage');
    // 权限菜单过滤相关
    this.$store.commit('updateMenulist');
  },
  created () {
    let tagsList = [];
    appRouter.map((item) => {
      if (item.children.length <= 1) {
        tagsList.push(item.children[0]);
      } else {
        tagsList.push(...item.children);
      }
    });
    this.$store.commit('setTagsList', tagsList);
  }
});

// axios 拦截请求
util.ajax.interceptors.request.use((config) => {
  const c = config;
  c.headers = { Authorization: Cookies.get('authToken') };
  return c;
},
error => Promise.reject(error));

// util.ajax 拦截响应
util.ajax.interceptors.response.use((response) => {
  if (response.data.msg) {
    vm.$Message.success(response.data.msg);
  }
  return response;
},
(error) => {
  if (error.response) {
    vm.$Message.error(`${error.response.status}:${error.response.data.msg}`);
  } else {
    vm.$Message.error(error.message);
  }

  if (error.response) {
  // 401 表示认证失败
    if (error.response.status === 401) {
    // 登陆操作
      store.commit('logout');
      store.commit('clearOpenedSubmenu');

      // 返回登录页面
      router.push({ name: 'login' });
    }
  }

  return Promise.reject(error);
});
