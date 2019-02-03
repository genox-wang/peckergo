<template>
  <div>
    <Collapse>
      <Panel name="1">
        选项
        <div slot="content">
          <Row>
            <Form :label-width="80">

            </Form>
          </Row>
          <Row>
            <Col span="24">
              <Button type="primary" shape="circle" icon="ios-search" style="float: right"  @click="reset" ghost></Button>
            </Col>
          </Row>
        </div>
      </Panel>
    </Collapse>
    <br>
    <data-table-base
      :data="items"
      :loading="tableLoading"
      :columns="columns"
      :total="pagination.total"
      :current="pagination.page"
      :page-size="pagination.size"
      @on-sort-change="sortChanged"
      @on-change="pageChanged"
      @on-page-size-change="sizeChanged"
    >
    </data-table-base>

  </div>
</template>

<script>
import DataTableBase from '@/views/base/data_table_base';
import vueRouterKeepaliveReset from '@/views/mixins/vue_router_keepalive_reset';
import tableFilterHelper from '@/views/mixins/data_table_helper';

export default {
  mixins: [
    vueRouterKeepaliveReset,
    tableFilterHelper
  ],
  components: {
    DataTableBase
  },
  data () {
    return {
      apiGet: 'get_dbtabless',
      deleteModel: false,
      tableLoading: false,
      selectedID: 0,
      pagination: {
        total: 0,
        page: 1,
        size: 10
      },
      items: [],
      orders: [],
      columns: [
        {
          title: 'Name',
          key: 'name',
          width: 300,
          tooltip: true
        },
        {
          title: 'Rows',
          key: 'rows',
          width: 150
        },
        {
          title: 'DataSize',
          key: 'data_size',
          width: 200,
          render: (h, {row}) => {
            let size = row.data_size;
            let suffixs = ['B', 'KB', 'MB', 'GB', 'TB'];
            let idx = 0;
            while (size > 1024) {
              size = size / 1024;
              idx++;
            }
            return h('div', `${size.toFixed(2)} ${suffixs[idx]}`);
          }
        },
        {
          title: 'IndexSize',
          key: 'index_size',
          width: 200,
          render: (h, {row}) => {
            let size = row.index_size;
            let suffixs = ['B', 'KB', 'MB', 'GB', 'TB'];
            let idx = 0;
            while (size > 1024) {
              size = size / 1024;
              idx++;
            }
            return h('div', `${size.toFixed(2)} ${suffixs[idx]}`);
          }
        },
        {
          title: 'TotalSize',
          key: 'total_size',
          width: 200,
          render: (h, {row}) => {
            let size = row.total_size;
            let suffixs = ['B', 'KB', 'MB', 'GB', 'TB'];
            let idx = 0;
            while (size > 1024) {
              size = size / 1024;
              idx++;
            }
            return h('div', `${size.toFixed(2)} ${suffixs[idx]}`);
          }
        }
      ]
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

    formatFilters () {
      this.filters = [];
    }

  }
};
</script>
