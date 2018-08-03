import util from '../../libs/util';

let ajax = util.ajax;

const {{modelName}} = {
  state: {
  },

  getters: {
  },

  actions: {
    get_{{model_name}}s ({}, {limit, page, orders, filters}) {
      return ajax.get(`/console/{{model_name}}s/?${util.getAllQuery(limit, page, orders, filters)}`)
        .then(resp => {
          return resp.data;
        });
    },

    get_{{model_name}}_by_id ({}, payload) {
      return ajax.get(`/console/{{model_name}}s/${payload}`)
        .then(resp => {
          return resp.data;
        });
    },

    create_{{model_name}}: ({}, payload) => ajax.post('/console/{{model_name}}s/', payload)
      .then(resp => {
        return resp.data;
      }),

    update_{{model_name}}: ({}, payload) => ajax.put(`/console/{{model_name}}s/${payload.id}`, payload)
      .then(resp => {
        return resp.data;
      }),

    delete_{{model_name}}: ({}, payload) => ajax.delete(`/console/{{model_name}}s/${payload}`)
      .then(resp => {
        return resp.data;
      })

  },

  mutations: {
  }
};

export default {{modelName}};
