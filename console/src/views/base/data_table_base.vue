<template>
  <div>
    <Table
      :data="data"
      :loading="loading"
      :columns="myColumns"
      size="small"
      @on-sort-change="$emit('on-sort-change', arguments[0])"
      stripe
      border
      disabled-hover
    >
    </Table>
    <div style="margin: 10px;overflow: hidden">
      <div style="float: right;">
        <Page
          :total="total"
          :current="current"
          :page-size="pageSize"
          :page-size-opts="[5,10,20]"
          @on-change="$emit('on-change')"
          @on-page-size-change="$emit('on-page-size-change')"
          show-total
          show-elevator
          show-sizer
        ></Page>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  props: {
    current: Number,
    total: Number,
    pageSize: Number,
    data: Array,
    loading: Boolean,
    columns: Array,
    showEdit: Boolean,
    showDelete: Boolean,
    showAction: Boolean,
    actionWith: {
      type: Number,
      default: 150
    },
    actions: {
      type: Function,
      // 对象或数组且一定会从一个工厂函数返回默认值
      default: function (h) {
        return [];
      }
    }
  },
  computed: {
    myColumns () {
      let myColumns = [
        {
          title: 'ID',
          key: 'id',
          fixed: 'left',
          width: 100,
          sortable: 'custom'
        }
      ];

      if (this.showAction) {
        myColumns.push({
          title: '操作',
          key: 'action',
          width: this.actionWith,
          fixed: 'right',
          align: 'center',
          render: (h, params) => {
            let renders = [];
            if (this.showEdit) {
              renders.push(h('Button', {
                props: {
                  type: 'primary',
                  size: 'small'
                },
                style: {
                  marginRight: '5px'
                },
                on: {
                  click: () => {
                    this.$emit('on-edit', params.row.id);
                  }
                }
              }, '编辑'));
            }
            if (this.showDelete) {
              renders.push(h('Button', {
                props: {
                  type: 'error',
                  size: 'small'
                },
                style: {
                  marginRight: '5px'
                },
                on: {
                  click: () => {
                    this.$emit('on-delete', params.row.id);
                  }
                }
              }, '删除'));
            }
            renders.push(...this.actions(h));
            return h('div', renders);
          }
        });
      }
      myColumns.push(...this.columns);
      return myColumns;
    }
  }
};
</script>
