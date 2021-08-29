# 说明

### 一 简介

- 前端技术 [Vue](https://github.com/vuejs/vue) + [Vue-Router](https://router.vuejs.org/zh/) + [Vuex](https://vuex.vuejs.org/zh/guide/) + [iView](https://github.com/iview/iview) + [iView-Admin](https://github.com/iview/iview-admin) + [Webpack](https://github.com/webpack/webpack) 
- 后端技术 [Gin](https://github.com/gin-gonic/gin) + [Gorm](https://github.com/gin-gonic/gin) + [Viper](https://github.com/spf13/viper)


### 二 使用方法

#### 1. 运行后端代码 (需要 golang 10 以上 + glide)

```
// mysql 数据库创建 yourApp 数据库


cd yourApp

// 下载依赖
go mod

// 启动后端 api 服务
go run main.go

```

#### 2. 运行前端代码 (需要 npm 6 + node 10)

```
cd console

// 下载前端依赖
npm install

// 运行
npm run serve

```

#### 3. 项目部署 (需要 golang 10 + glide + npm 6 + node 10 + docker + docekr-compose)

```
// 项目克隆到服务器 src 文件下 

cp yourApp

./build.sh
```

### 三 开发准备工作 

安装 [packer](https://github.com/genox-wang/pecker)

### 四 开发

#### 1. 生成 model

```
// 进入项目根目录

pecker model yourModelName

// 工具根据 /templates 目录下的模板自动生成对应 model 相关的前后端代码

```

#### 2. 编辑生成的代码

- 搜索 `TODO`
- 根据 `TODO` 的指引进行选择编辑。 /templates 文件夹下的文件不需要编辑
- 编辑完运行对生成的 model 进行图形化编辑了。可选择进行分表

### 五 常用功能实现

#### 1. 表单字段筛选

`console/src/views/{model}/index.vue`

```javascript
// ui 相关部分代码略过

...

// 在 formatFilters 方法中添加代码

methods: {
  formatFilters () {
    this.filters = [];
    ...
    // 过滤 channel_id。 注意这里的 channel_id 是对应数据库内表的字段 this.selectedChannelId 是要查找的值
    // 在 console./src/vies/mixins/data_table_helper.js 里有其他过滤的方法
    if (this.selectedChannelId) {
      this.fPushEqual('channel_id', this.selectedChannelId);
    }

```

基于 `model.TableFilterMode` 定义的结构实现，过滤和分页

前端相关代码

`lib/util.js`

``` javascript
util.getAllQuery = function (limit, page, orders, filters) {
  limit = limit || 10;
  page = page || 0;
  orders = Array.isArray(orders) ? orders : [];
  filters = Array.isArray(filters) ? filters : [];
  let query = `limit=${limit}&page=${page}`;
  orders.forEach(e => {
    query += `&order=${e}`;
  });
  filters.forEach(e => {
    query += `&${e}`;
  });
  return query;
};

// we wrap qurey for equal
util.we = (f, v) => `${f}=${v}`;
// wgt wrap qurey for greater than
util.wgt = (f, v) => `${f}=^^^${v}`;
// wgle wrap qurey for greater than or equal
util.wgle = (f, v) => `${f}=^^${v}`;
// wlt wrap qurey for lesser than
util.wlt = (f, v) => `${f}=___${v}`;
// wlte wrap qurey for lesser than or equal
util.wlte = (f, v) => `${f}=__${v}`;
// wr wrap qurey for range
util.wr = (f, v1, v2) => `${f}=[${v1},${v2}]`;
```

#### 2. 表单排序

`console/src/views/{model}/index.vue`

```javascript
 data () {
    return {
      ...
      columns: [
        {
          title: '名称',
          key: 'name',
          width: 150,
          sortable: 'custom', // 添加改字段支持过滤，注意要支持过滤，表中的字段要与 key 一致
          tooltip: true
        },

```

#### 3. 前端路由修改

`console/src/router/router.js`

```
// 这里会自动添加编辑和新增路由，如果不需要这些路由可以删除
// 作为Main组件的子页面展示但是不在左侧菜单显示的路由写在otherRouter里
export const otherRouter = {
  path: '/',
  name: 'otherRouter',
  redirect: '/home',
  component: Main,
  children: [
    ...
    ,{ path: '/management/channels/new', title: '新建渠道', name: 'channel_new', access: [1], component: () => import('@/views/management-channel/new.vue') },
    { path: '/management/channels/:id', title: '编辑渠道', name: 'channel_edit', access: [1], component: () => import('@/views/management-channel/edit.vue') }
  ]
};

// 这里可以编辑1级或者2级导航 
// 作为Main组件的子页面展示并且在左侧菜单显示的路由写在appRouter里
export const appRouter = [
  {
    path: '/management',
    icon: 'md-cog',
    name: 'management',
    title: '管理',
    component: Main,
    children: [
      {
        path: 'channels',
        icon: 'logo-buffer',
        name: 'channels',
        title: '渠道',
        access: [1], // 支持的角色字段，角色id 定义在 api/model/user.go
        component: () => import('@/views/management-channel/index.vue')
      },
      // ph-appRouter don't remove this line
    ]
  }
];
```

#### 4. 前端ID-Map映射展示

```javascript
...
// 这里负责展示 idMap
</Col>
  <Col span="12">
  <FormItem label="渠道" prop="channel_id">
    <Select v-model="form.channel_id" filterable>
      <Option v-for="(item, idx) in channelMap" :value="idx" :key="idx">{{ item }}</Option>
    </Select>
  </FormItem>
</Col>

...

// 这里取到 后端定义的 idMap
computed: {
  channelMap () {
    return this.$store.state.channel.idNameMap;
  },
},
```

#### 5. 快捷开关


```javascript
data () {
  return {
    ...
    columns: [
      {
        title: '开关',
        key: 'enabled',
        width: 100,
        fixed: 'right',
        render: (h, {row}) => {
          return h('i-switch', {
            props: {
              size: 'large',
              value: row.enabled
            },
            on: {
              'on-change': (enabled) => {
                // 修改这里调用要请求后端的接口
                // 这里注意后端对于字段为0/false，不会更新，所以不要用bool类型定义字段，用int, 1 代表 true, 2 代表 false。 方便更新操作
                this.$store.dispatch('update_channel', {
                   id: row.id,
                   enabled: enabled ? 1 : 2
                });
              }
            }
          }, [
            h('span', {
              slot: 'open',
              props: {
                value: 1
              }
            }, '开启'),
            h('span', {
              slot: 'close',
              props: {
                value: 0
              }
            }, '关闭')
          ]);
        }
      }
```

### 六 自定义模板

模板代码在 `/templates` 文件夹下

`/templates/model` 文件夹下定义新增的模板
`/templates/model_append` 文件夹定义累加更新模板

#### 1. 宏定义
{{projectName}} 项目名
{{modelName}} model 小写驼峰
{{model_name}} model 小写蛇形
{{model-name}} model -形小写
{{ModelName}} model 大写驼峰
{{MODEL_NAME}} model 大写蛇形

#### 2. 新增的模板

- 文件名定义里新增文件目标位置 

`api^controller^{{model_name}}_controller.go` 代表新增文件 `api/controller/{{model_name}}_controller.go

- 模板内容为新增文件的内容，可以添加宏

#### 3. 累加更新模板

- 文件名定义累加的目标文件

`api^model^init.go` 代表目标文件时  `api/model/init.go`

- 模板内容定义累加的文件内容，最后一行定义累加文件的替换位置

以下代表要把下面内容替换指定文件内 `//ph-AutoMigrate don't remove this line`
```
new({{ModelName}}),
		//ph-AutoMigrate don't remove this line
```

### 七 其他

#### 1. 状态管理

前端项目状态管理基于 [Vuex](https://vuex.vuejs.org/zh/guide/) 实现，请求结构基本都定义在 `store`文件夹下

在`views/model/index.vue` 一般这样条用内部定义接口 `this.$store.dispatch('update_app', this.form);`

```javascript
import util from '../../libs/util';

let ajax = util.ajax;

const app = {
  state: {
    idNameMap: {}
  },

  getters: {
  },

  actions: {
    update_app: ({}, payload) => ajax.put(`/console/apps/${payload.id}`, payload)
      .then(resp => {
        return resp.data;
      }),

    ...

    get_app_id_name_map: ({
      commit
    }) => ajax.get('/console/map/apps/')
      .then((resp) => {
        commit('SET_APP_ID_NAME_MAP', resp.data);
      })

  },

  mutations: {
    SET_APP_ID_NAME_MAP (state, payload) {
      state.idNameMap = payload;
    }
  }
};

export default app;
```

#### 2. js 第三方库使用

`src/main.js` 为项目入口，这里可以引入一些第三方库，并设置全局使用

```
import dayjs from 'dayjs';
import lodash from 'lodash';

...

// 在 vue 内调用 `this.$d` 或者`this._` 就可以使用第三方库功能

Vue.prototype.$d = dayjs;
Vue.prototype._ = lodash;
```


