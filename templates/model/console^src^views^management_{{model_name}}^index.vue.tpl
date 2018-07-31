<template>
  <div>
    <Collapse>
      <Panel name="1">
        选项
        <div slot="content">
          <Row>
            <Form :label-width="80">
              <Col span="8">
                <FormItem label="创建时间">
                  <DatePicker v-model="dateTimeRange" type="datetimerange" format="yyyy-MM-dd HH:mm" placeholder="创建时间" style="width: 260px"></DatePicker>
                </FormItem>
              </Col>
            </Form>
          </Row>
          <Row>
            <Col span="12">
              <Button type="success"  shape="circle" @click="onNew" ghost>新建</Button>
            </Col>
            <Col span="12">
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
      @on-edit="onEdit"
      @on-delete="onDelete"
      show-action
      show-edit
      show-delete
    >
    </data-table-base>

    <Modal
      v-model="deleteModel"
      title="确认删除"
      @on-ok="deleteOk"
     >
      <p>确认删除该行吗？</p>
    </Modal>
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
      apiGet: 'get_{{model_name}}s',
      deleteModel: false,
      tableLoading: false,
      dateTimeRange: [],
      selectedID: 0,
      pagination: {
        total: 0,
        page: 1,
        size: 10
      },
      items: [],
      orders: [],
      columns: [
        // TODO
        {
          title: '创建时间',
          key: 'created_at',
          minWidth: 150,
          sortable: 'custom',
          render: (h, {row}) => {
            return h('div', this.$d(row.created_at).format('YYYY-MM-DD HH:mm:ss'));
          }
        },
        {
          title: '更新时间',
          key: 'updated_at',
          minWidth: 150,
          sortable: 'custom',
          render: (h, {row}) => {
            return h('div', this.$d(row.updated_at).format('YYYY-MM-DD HH:mm:ss'));
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

    onNew () {
      this.$router.push({
        name: '{{model_name}}_new'
      });
    },

    onEdit (id) {
      this.$router.push({
        name: '{{model_name}}_edit',
        params: {
          id: id
        }
      });
    },

    onDelete (id) {
      this.deleteModel = true;
      this.selectedID = id;
    },

    formatFilters () {
      this.filters = [];
      this.fPushTimeRange('created_at', this.dateTimeRange);
    },

    deleteOk () {
      this.$store.dispatch('delete_{{model_name}}', this.selectedID).then(() => {
        this.$Message.success('删除成功');
        this.reset();
      });
    }
  }
};
</script>
