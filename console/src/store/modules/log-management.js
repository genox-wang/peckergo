import util from '../../libs/util';

let ajax = util.ajax;

const logManagement = {
  state: {
    // 为前端暴露 ID-Name 映射
    // idNameMap: {},
  },

  getters: {
  },

  actions: {
    get_log_managements ({}, {limit, page, orders, filters}) {
      return ajax.get(`/console/log_managements/?${util.getAllQuery(limit, page, orders, filters)}`)
        .then(resp => {
          return resp.data;
        });
    },

    create_log_management: ({}, payload) => ajax.post('/console/log_managements/', payload)
      .then(resp => {
        return resp.data;
      }),

    // 为前端暴露 ID-Name 映射
    // get_log_management_id_name_map: ({
    //   commit
    // }) => ajax.get('/console/map/log_managements/')
    //   .then((resp) => {
    //     commit('SET_LOG_MANAGEMENT_ID_NAME_MAP', resp.data);
    //  }),
  },

  mutations: {
    SET_LOG_MANAGEMENT_ID_NAME_MAP (state, payload) {
      state.idNameMap = payload;
    }
  }
};

export default logManagement;
