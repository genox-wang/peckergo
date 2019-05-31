<template>
  <edit-base ref="base" :form="form" :rules="rules" :save-promise="save">
    <Row>
      <!-- TODO  添加表单 html -->
      <!-- <Col span="16">
        <FormItem label="昵称" prop="name">
          <Input placeholder="昵称" v-model="form.name" @on-enter.stop="$refs.base.save()"></Input>
        </FormItem>
      </Col> -->
    </Row>
  </edit-base>
</template>

<script>
import editBase from '../base/edit-base.vue';
import vueRouterKeepaliveReset from '@/views/mixins/vue_router_keepalive_reset';

export default {
  mixins: [
    vueRouterKeepaliveReset
  ],
  components: {
    editBase
  },
  data () {
    return {
      form: {
        // TODO 添加表单字段
        // name: ''
      },
      rules: {
        // TODO 添加表单验证策略  键对应表单的 prop
        // https://github.com/yiminghe/async-validator
        // name: [
        //   { required: true, message: '名称不能为空', trigger: 'blur' }
        // ],
      }
    };
  },
  methods: {
    save () {
      // TODO 保存前给表单数据进行处理
      // this.form.name = this.form.name + '';
      return this.$store.dispatch('update_{{model_name}}', this.form);
    },
    reset () {
      let self = this;
      let id = self.$route.params.id;
      this.$store.dispatch('get_{{model_name}}_by_id', this.$route.params.id).then((data) => {
        if (data) {
          // TODO 初始化表单数据
          self.form.id = parseInt(id);
          // self.form.name = data.name;
        }
      });
    }
  }
};
</script>
