import Vue from 'vue';
import iView from 'iview';
import Util from '../libs/util';
import VueRouter from 'vue-router';
import Cookies from 'js-cookie';
import {routers, otherRouter, appRouter} from './router';

Vue.use(VueRouter);

// 路由配置
const RouterConfig = {
  // mode: 'history',
  routes: routers
};

export const router = new VueRouter(RouterConfig);

router.beforeEach((to, from, next) => {
  iView.LoadingBar.start();
  Util.title(to.meta.title);

  if (!Cookies.get('authToken') && to.name !== 'login') { // 判断是否已经登录且前往的页面不是登录页
    next({
      name: 'login'
    });
  } else if (Cookies.get('authToken') && to.name === 'login') { // 判断是否已经登录且前往的是登录页
    Util.title();
    next({
      name: 'home_index'
    });
  } else {
    const curRouterObj = Util.getRouterObjByName([otherRouter, ...appRouter], to.name);
    if (curRouterObj && curRouterObj.access instanceof Array) { // 需要判断权限的路由
      const access = parseInt(Cookies.get('access'));
      let i = 0;
      for (;i < curRouterObj.access.length; i++) {
        if (curRouterObj.access[i] === access) {
          Util.toDefaultPage([otherRouter, ...appRouter], to.name, router, next); // 如果在地址栏输入的是一级菜单则默认打开其第一个二级菜单的页面
          return;
        }
      }
      next({
        replace: true,
        name: 'error-403'
      });
    } else { // 没有配置权限的路由, 直接通过
      Util.toDefaultPage([...routers], to.name, router, next);
    }
  }
});

router.afterEach((to) => {
  Util.openNewPage(router.app, to.name, to.params, to.query);
  iView.LoadingBar.finish();
  window.scrollTo(0, 0);
});
