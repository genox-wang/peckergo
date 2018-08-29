import util from '@/libs/util';

export default {
  data () {
    return {
      filters: []
    };
  },

  methods: {
    pageChanged (page) {
      this.pagination.page = page;
      this.reset();
    },

    sizeChanged (size) {
      this.pagination.size = size;
      this.reset();
    },

    sortChanged ({key, order}) {
      this.orders = this._.isEqual(order, 'desc') ? ['-' + key] : [key];
      this.reset();
    },

    fPushTimeRange (fieldName, dateTimeRange) {
      if (this._.isEqual(typeof (dateTimeRange[0]), 'object') || this._.isEqual(typeof (dateTimeRange[1]), 'object')) {
        let sTime = dateTimeRange[0].toISOString();
        let eTime = dateTimeRange[1].toISOString();
        this.filters.push(util.wr(fieldName, sTime, eTime));
      }
    },

    fPushEqual (fieldName, value) {
      this.filters.push(util.we(fieldName, value));
    },

    fPushGreaterThan (fieldName, value) {
      this.filters.push(util.wgle(fieldName, value));
    },

    reset () {
      if (this._reset) {
        this._reset();
      }
      this.formatFilters();
      let self = this;
      this.tableLoading = true;
      this.$store.dispatch(this.apiGet, {
        limit: this.pagination.size,
        page: this.pagination.page - 1,
        orders: this.orders,
        filters: this.filters
      }).then((resp) => {
        this.tableLoading = false;
        self.items = resp.data;
        if (resp.meta) {
          self.pagination.total = resp.meta.pagination.total;
          self.pagination.page = resp.meta.pagination.page + 1;
        }
      }).catch(() => {
        this.tableLoading = false;
      });
    }
  }
}
;
