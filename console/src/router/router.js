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
  redirect: '/home',
  component: Main,
  children: [
    { path: 'home', title: '首页', name: 'home_index', component: () => import('@/views/home/home.vue') },
    { path: '/management/user/new', title: '新建用户', name: 'user_new', access: [0], component: () => import('@/views/management-user/new.vue') },
    { path: '/management/user/edit', title: '编辑用户', name: 'user_edit', access: [0], component: () => import('@/views/management-user/edit.vue') }
  ]
};

// 作为Main组件的子页面展示并且在左侧菜单显示的路由写在appRouter里
export const appRouter = [
  {
    path: '/management',
    icon: 'gear-a',
    name: 'management',
    title: '管理',
    component: Main,
    children: [
      {
        path: 'user',
        icon: 'person-stalker',
        name: 'user',
        title: '用户',
        access: [0],
        component: () => import('@/views/management-user/index.vue')
      }
    ]
  }
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
