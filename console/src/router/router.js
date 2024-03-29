import Main from '@/views/Main.vue';

// 不作为Main组件的子页面展示的页面单独写，如下
export const loginRouter = {
  path: '/login',
  name: 'login',
  meta: {
    title: 'Login - 登录'
  },
  component: () => import('@/views/login.vue')
};

export const page404 = {
  path: '/*',
  name: 'error-404',
  meta: {
    title: '404-页面不存在'
  },
  component: () => import('@/views/error-page/404.vue')
};

export const page403 = {
  path: '/403',
  meta: {
    title: '403-权限不足'
  },
  name: 'error-403',
  component: () => import('@//views/error-page/403.vue')
};

export const page500 = {
  path: '/500',
  meta: {
    title: '500-服务端错误'
  },
  name: 'error-500',
  component: () => import('@/views/error-page/500.vue')
};

// 作为Main组件的子页面展示但是不在左侧菜单显示的路由写在otherRouter里
export const otherRouter = {
  path: '/',
  name: 'otherRouter',
  redirect: '/stats/home',
  component: Main,
  children: [
    // { path: 'home', title: '首页', name: 'home_index', component: () => import('@/views/home/home.vue') },
    { path: '/management/users/new', title: '新建用户', name: 'user_new', access: [1], component: () => import('@/views/management-user/new.vue') },
    { path: '/management/users/:id', title: '编辑用户', name: 'user_edit', access: [1], component: () => import('@/views/management-user/edit.vue') },
    // ph-otherRouter don't remove this line
  ]
};

// 作为Main组件的子页面展示并且在左侧菜单显示的路由写在appRouter里
export const appRouter = [
  {
    path: '/stats',
    icon: 'ios-speedometer-outline',
    name: 'stats',
    title: '数据分析',
    component: Main,
    children: [
      {
        path: 'home',
        icon: 'ios-podium',
        name: 'home_index',
        title: 'PV-UV',
        access: [1],
        component: () => import('@/views/home/home.vue')
      },
    ]
  },
  {
    path: '/tables',
    icon: 'md-speedometer',
    name: 'tables',
    title: '监控',
    component: Main,
    children: [
      {
        path: 'dbtabless',
        icon: 'logo-buffer',
        name: 'dbtabless',
        title: '数据库监控',
        access: [1],
        component: () => import('@/views/management-dbtables/index.vue')
      },
    ]
  },
  {
    path: '/usr_mgr',
    icon: '',
    name: 'usr_mgr',
    title: '用户管理',
    component: Main,
    children: [
      {
        path: 'users',
        icon: 'md-person',
        name: 'users',
        title: '用户配置',
        access: [1],
        component: () => import('@/views/management-user/index.vue')
      },
      // ph-appRouter don't remove this line
    ]
  },
];

// 所有上面定义的路由都要写在下面的routers里
export const routers = [
  loginRouter,
  otherRouter,
  ...appRouter,
  page500,
  page403,
  page404
];
