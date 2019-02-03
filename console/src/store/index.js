import Vue from 'vue';
import Vuex from 'vuex';

import app from './modules/app';
import user from './modules/user';
import logManagement from './modules/log-management';
import dbtables from './modules/dbtables';
// ph-store-index-import don't remove this line

Vue.use(Vuex);

const store = new Vuex.Store({
  state: {
    //
  },
  mutations: {
    //
  },
  actions: {

  },
  modules: {
    app,
    user,
    logManagement,
    dbtables,
    // ph-store-index-modules don't remove this line
  },
});

export default store;
