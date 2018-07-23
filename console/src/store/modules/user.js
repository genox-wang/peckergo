import Cookies from 'js-cookie';
import { Base64 } from 'js-base64';
import util from '../../libs/util';

let ajax = util.ajax;

const user = {
  state: {
    token: Cookies.get('authToken')
  },

  getters: {
    userInfo (state) {
      if (state.token) {
        const jwtTokens = state.token.split('.');
        if (jwtTokens.length > 1) {
          return JSON.parse(Base64.decode(jwtTokens[1]));
        }
      }
      return '';
    }
  },

  actions: {
    // 登陆操作
    login ({
      commit
    }, payload) {
      return ajax.post('/console/login', payload)
        .then((resp) => {
        // 保存 token 到  Cookie
          Cookies.set('authToken', resp.data.token);
          const jwtTokens = resp.data.token.split('.');
          let user = {};
          if (jwtTokens.length > 1) {
            user = JSON.parse(Base64.decode(jwtTokens[1]));
          }
          Cookies.set('access', user.role);
          // 设置 state.token
          commit('setToken', resp.data.token);
          return resp;
        });
    },

    get_captcha () {
      return ajax.get('/console/captcha').then(resp => {
        return resp.data;
      });
    },

    get_users ({}, {limit, page, orders, filters}) {
      return ajax.get(`/console/users/?${util.getAllQuery(limit, page, orders, filters)}`)
        .then(resp => {
          return resp.data;
        });
    },

    get_user_by_id ({}, payload) {
      return ajax.get(`/console/users/${payload}`)
        .then(resp => {
          return resp.data;
        });
    },

    create_user: ({}, payload) => ajax.post('/console/users/', payload)
      .then(resp => {
        return resp.data;
      }),

    update_user: ({}, payload) => ajax.put(`/console/users/${payload.id}`, payload)
      .then(resp => {
        return resp.data;
      }),

    delete_user: ({}, payload) => ajax.delete(`/console/users/${payload}`)
      .then(resp => {
        return resp.data;
      })

  },

  mutations: {
    setToken (state, token) {
      state.token = token;
    },

    logout (state, vm) {
      Cookies.remove('authToken');
      Cookies.remove('access');
      // 恢复默认样式
      // let themeLink = document.querySelector('link[name="theme"]');
      // themeLink.setAttribute('href', '');
      // 清空打开的页面等数据，但是保存主题数据
      let theme = '';
      if (localStorage.theme) {
        theme = localStorage.theme;
      }
      localStorage.clear();
      if (theme) {
        localStorage.theme = theme;
      }
    }
  }
};

export default user;
