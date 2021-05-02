<template>
  <edit-base ref="base" :form="form" :rules="rules" :save-promise="save">
    <Row>
      <!-- TODO  添加表单 html -->
      <!-- <Col span="16">
        <FormItem label="昵称" prop="display_name">
          <Input placeholder="昵称" v-model="form.display_name" @on-enter.stop="$refs.base.save()"></Input>
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
      },
      rules: {
        // TODO 添加表单验证策略
        // https://github.com/yiminghe/async-validator
        // username: [
        //   { required: true, message: '账号不能为空', trigger: 'blur' }
        // ],
      }
    };
  },
  methods: {
    save () {
      // TODO 保存前给表单数据进行处理
      return this.$store.dispatch('update_{{model_name}}', this.form);
    },
    reset () {
      let self = this;
      let id = self.$route.params.id;
      this.$store.dispatch('get_{{model_name}}_by_id', this.$route.params.id).then((data) => {
        if (data) {
          // TODO 初始化表单数据 如 self.form.name = data.name;
          self.form.id = parseInt(id);
        }
      });
    }
  }
};
</script>
