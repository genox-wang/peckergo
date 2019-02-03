import util from '../../libs/util';

let ajax = util.ajax;

const dbtables = {
  state: {
  },

  getters: {
  },

  actions: {
    get_dbtabless ({}, {limit, page, orders, filters}) {
      return ajax.get(`/console/dbtabless/?${util.getAllQuery(limit, page, orders, filters)}`)
        .then(resp => {
          return resp.data;
        });
    }
  },

  mutations: {
  }
};

export default dbtables;
