<template>
  <edit-base ref="base" :form="form" :rules="rules" :save-promise="save">
    <Row>
      <!-- TODO 添加表单 html-->
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
        // TODO 添加表单验证规则 
        // https://github.com/yiminghe/async-validator
        // username: [
        //   { required: true, message: '账号不能为空', trigger: 'blur' }
        // ],
      }
    };
  },
  methods: {
    save () {
      // TODO
      return this.$store.dispatch('create_{{model_name}}', this.form);
    },
    reset () {
      // TODO
    }
  }
};
</script>
