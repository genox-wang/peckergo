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
              <Col span="8">
                <FormItem label="角色">
                  <Select v-model="role" style="width:200px" filterable clearable>
                    <Option :value="1">管理员</Option>
                    <Option :value="2">操作员</Option>
                  </Select>
                </FormItem>
              </Col>
              <Col span="8">
                <FormItem label="昵称">
                  <Input v-model="displayName" placeholder="昵称" clearable style="width: 200px"></Input>
                </FormItem>
              </Col>
            </Form>
          </Row>
          <Row>
            <Col span="12">
              <Button type="success" @click="onNew">新建</Button>
            </Col>
            <Col span="12">
              <Button type="primary" style="float: right"  @click="reset">查询</Button>
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
      apiGet: 'get_users',
      deleteModel: false,
      tableLoading: false,
      selectedID: 0,
      pagination: {
        total: 0,
        page: 1,
        size: 10
      },
      dateTimeRange: [],
      role: 0,
      displayName: '',
      items: [],
      orders: [],
      columns: [
        {
          title: '用户名',
          key: 'username',
          minWidth: 100,
          sortable: 'custom'
        },
        {
          title: '昵称',
          key: 'display_name',
          minWidth: 100,
          sortable: 'custom'
        },
        {
          title: '角色',
          align: 'center',
          key: 'role',
          minWidth: 100,
          sortable: 'custom',
          render: (h, params) => {
            const roles = ['未知', '管理员', '操作员'];
            const colors = ['yellow', 'red', 'blue'];
            let id = params.row.role;
            return h('tag', {
              props: {
                color: colors[id]
              }
            }, roles[id]);
          }
        },
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
        name: 'user_new'
      });
    },

    onEdit (id) {
      this.$router.push({
        name: 'user_edit',
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
      if (this.role && this.role !== 0) {
        this.fPushEqual('role', this.role);
      }
      if (this.displayName !== '') {
        this.fPushEqual('display_name', this.displayName);
      }
    },

    deleteOk () {
      this.$store.dispatch('delete_user', this.selectedID).then(() => {
        this.$Message.success('删除成功');
        this.reset();
      });
    }
  }
};
</script>
