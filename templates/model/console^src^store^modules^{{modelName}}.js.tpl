import util from '../../libs/util';

let ajax = util.ajax;

const {{modelName}} = {
  state: {
  },

  getters: {
  },

  actions: {
    get_{{modelName}}s ({}, {limit, page, orders, filters}) {
      return ajax.get(`/console/{{modelName}}s/?${util.getAllQuery(limit, page, orders, filters)}`)
        .then(resp => {
          return resp.data;
        });
    },

    get_{{modelName}}_by_id ({}, payload) {
      return ajax.get(`/console/{{modelName}}s/${payload}`)
        .then(resp => {
          return resp.data;
        });
    },

    create_{{modelName}}: ({}, payload) => ajax.post('/console/{{modelName}}s/', payload)
      .then(resp => {
        return resp.data;
      }),

    update_{{modelName}}: ({}, payload) => ajax.put(`/console/{{modelName}}s/${payload.id}`, payload)
      .then(resp => {
        return resp.data;
      }),

    delete_{{modelName}}: ({}, payload) => ajax.delete(`/console/{{modelName}}s/${payload}`)
      .then(resp => {
        return resp.data;
      })

  },

  mutations: {
  }
};

export default {{modelName}};
