<template>
  <div>
    <Collapse>
      <Panel name="1">
        选项
        <div slot="content">
          <Row>
            <Form :label-width="80">
              <!-- TODO 分表注释以下 -->
              <Col span="8">
                <FormItem label="创建时间">
                  <DatePicker v-model="dateTimeRange" type="datetimerange" format="yyyy-MM-dd HH:mm" placeholder="创建时间" style="width: 260px"></DatePicker>
                </FormItem>
              </Col>
              <Col span="8">
                <FormItem label="用户">
                  <Select v-model="selectedUserID" clearable filterable>
                    <Option v-for="(item,idx) in userMap" :value="idx" :key="idx">{{ item }} </Option>
                  </Select>
                </FormItem>
              </Col>
              <Col span="8">
                <FormItem label="方法">
                  <Select v-model="selectedMethod" clearable filterable>
                    <Option value="POST">POST</Option>
                    <Option value="PUT">PUT</Option>
                    <Option value="DELETE">DELETE</Option>
                  </Select>
                </FormItem>
              </Col>
              <Col span="8">
                <FormItem label="路由">
                  <Input placeholder="路由" v-model="selectedPath"></Input>
                </FormItem>
              </Col>
              <!-- TODO 分表取消注释
              <Col span="8">
                 <FormItem label="日期">
                   <DatePicker  v-model="date" type="date" placeholder="日期"></DatePicker>
                 </FormItem>
              </Col>
              <Col span="8">
                 <FormItem label="时间">
                   <TimePicker v-model="time" type="timerange" placeholder="时间"></TimePicker>
                </FormItem>
              </Col>
              -->
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
      :data="newItems"
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
  computed: {
    userMap () {
      return this.$store.state.user.idNameMap;
    },
    newItems () {
      return this.items.map(e => {
        try {
          e.body = JSON.stringify(JSON.parse(e.body), null, 2);
        } catch (err) {
          e.body = '{}';
        }
        return e;
      });
    }
  },
  data () {
    return {
      apiGet: 'get_log_managements',
      deleteModel: false,
      tableLoading: false,
      // TODO 分表注释下行
      dateTimeRange: [],
      // TODO 分表取消注释下面两行
      // date: '',
      // time: [],
      selectedID: 0,
      selectedMethod: '',
      selectedUserID: '',
      selectedPath: '',
      pagination: {
        total: 0,
        page: 1,
        size: 10
      },
      items: [],
      orders: ['-created_at'],
      columns: [
        {
          title: '日志时间',
          key: 'created_at',
          minWidth: 150,
          sortable: 'custom',
          sortType: 'desc',
          render: (h, {row}) => {
            return h('div', this.$d(row.created_at).format('YYYY-MM-DD HH:mm:ss'));
          }
        },
        {
          title: '用户',
          width: 150,
          render: (h, {row}) => {
            let name = this.userMap[row.user_id] || '未知';
            return h('div', `[${row.user_id}] ${name}`);
          }
        },
        {
          title: '方法',
          width: 100,
          render: (h, {row}) => {
            let colors = {
              post: 'primary',
              put: 'warning',
              delete: 'error',
            };

            return h('div', [h('tag', {
              props: {
                color: colors[row.method.toLowerCase()],
                type: 'border'
              }
            }, row.method)]);
          }
        },
        {
          title: '路由',
          key: 'path',
          width: 150,
          tooltip: true
        },
        {
          title: 'body',
          key: 'body',
          tooltip: true,
          width: 500
        },
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
      // TODO 分表注释下行
      this.fPushTimeRange('created_at', this.dateTimeRange);
      if (this.selectedUserID) {
        this.fPushEqual('user_id', this.selectedUserID);
      }
      if (this.selectedMethod) {
        this.fPushEqual('method', this.selectedMethod);
      }
      if (this.selectedPath) {
        this.fPushEqual('path', this.selectedPath);
      }
      // TODO 分表取消以下注释
      // let date = this.date ? this.$d(this.date) : this.$d();
      // let timeRange = [];
      // if (this.time[0]) {
      //  timeRange[0] = this.$d(date.format('YYYY-MM-DD') + ' ' + this.time[0]);
      //  timeRange[1] = this.$d(date.format('YYYY-MM-DD') + ' ' + this.time[1]);
      // }
      // this.filters.push('suffix=' + date.format('YYMMDD'));
      // this.fPushTimeRange('created_at', timeRange);
    },

    deleteOk () {
      this.$store.dispatch('delete_log_management', this.selectedID).then(() => {
        this.$Message.success('删除成功');
        this.reset();
      });
    },

    _reset () {
      this.$store.dispatch('get_user_id_name_map');
    }
  }
};
</script>
