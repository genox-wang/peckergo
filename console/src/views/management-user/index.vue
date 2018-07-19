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
              <Button type="success" @click="newLine">新建</Button>
            </Col>
            <Col span="12">
              <Button type="primary" style="float: right"  @click="reset">查询</Button>
            </Col>
          </Row>
        </div>
      </Panel>
    </Collapse>
    <!-- <Card>

    </Card> -->
    <br>
    <Table
      :data="items"
      :loading="tableLoading"
      :columns="columns"
      size="small"
      @on-sort-change="sortChanged"
      stripe
      border
      disabled-hover
    >
    </Table>
    <div style="margin: 10px;overflow: hidden">
      <div style="float: right;">
        <Page
          :total="pagination.total"
          :current="pagination.page"
          :page-size="pagination.size"
          :page-size-opts="[5,10,20]"
           @on-change="pageChanged"
           @on-page-size-change="sizeChanged"
           show-total
           show-elevator
           show-sizer
        ></Page>
      </div>
    </div>
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
import util from '@/libs/util';
import vueRouterKeepaliveReset from '@/views/mixins/vue_router_keepalive_reset';

export default {
  mixins: [
    vueRouterKeepaliveReset
  ],
  data () {
    return {
      routeName: '',
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
      filters: [],
      items: [],
      orders: [],
      columns: [
        {
          title: 'ID',
          key: 'id',
          fixed: 'left',
          width: 100,
          sortable: 'custom'
        },
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
        },
        {
          title: '操作',
          key: 'action',
          width: 150,
          fixed: 'right',
          align: 'center',
          render: (h, params) => {
            return h('div', [
              h('Button', {
                props: {
                  type: 'primary',
                  size: 'small'
                },
                style: {
                  marginRight: '5px'
                },
                on: {
                  click: () => {
                    this.$router.push({
                      name: 'user_edit',
                      params: {
                        id: params.row.id
                      }
                    });
                  }
                }
              }, '编辑'),
              h('Button', {
                props: {
                  type: 'error',
                  size: 'small'
                },
                style: {
                  marginRight: '5px'
                },
                on: {
                  click: () => {
                    this.deleteModel = true;
                    this.selectedID = params.row.id;
                  }
                }
              }, '删除')
            ]);
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

    newLine () {
      this.$router.push({
        name: 'user_new'
      });
    },

    pushFilterTimeRange (fieldName, dateTimeRange) {
      if (this._.isEqual(typeof (dateTimeRange[0]), 'object') || this._.isEqual(typeof (dateTimeRange[1]), 'object')) {
        let sTime = dateTimeRange[0].toISOString();
        let eTime = dateTimeRange[1].toISOString();
        this.filters.push(util.wr(fieldName, sTime, eTime));
      }
    },

    pushFilterEqual (fieldName, value) {
      this.filters.push(util.we(fieldName, value));
    },

    formatFilters () {
      this.filters = [];
      this.pushFilterTimeRange('created_at', this.dateTimeRange);
      if (this.role && this.role !== 0) {
        this.pushFilterEqual('role', this.role);
      }
      if (this.displayName !== '') {
        this.pushFilterEqual('display_name', this.displayName);
      }
    },

    reset () {
      this.formatFilters();
      let self = this;
      this.tableLoading = true;
      this.$store.dispatch('get_users', {
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
