<template>
  <edit-base :form="form" :rules="rules" :save-promise="save">
    <Row>
        <Col span="16">
          <FormItem label="用户名" prop="username">
            <Input placeholder="用户名" v-model="form.username"></Input>
          </FormItem>
        </Col>
        <Col span="16">
          <FormItem label="昵称" prop="display_name">
            <Input placeholder="昵称" v-model="form.display_name"></Input>
          </FormItem>
        </Col>
        <Col span="16">
          <FormItem label="新密码" prop="password">
            <Input placeholder="新密码" v-model="form.password"></Input>
          </FormItem>
        </Col>
        <Col span="16">
          <FormItem label="角色">
             <Select v-model="form.role">
              <Option :value="0">管理员</Option>
              <Option :value="1">操作员</Option>
            </Select>
          </FormItem>
        </Col>
      </Row>
  </edit-base>
</template>

<script>
import editBase from '../base/edit-base.vue';

export default {
  components: {
    editBase
  },
  data () {
    return {
      form: {
        id: 0,
        username: '',
        password: '',
        display_name: '',
        role: 0
      },
      rules: {
        username: [
          { required: true, message: '账号不能为空', trigger: 'blur' }
        ],
        display_name: [
          { required: true, message: '昵称不能为空', trigger: 'blur' }
        ]
      }
    };
  },
  methods: {
    save () {
      return this.$store.dispatch('update_user', this.form);
    }
  },
  mounted () {
    let payload = this.$route.params.payload;
    if (payload) {
      this.form.id = payload.id;
      this.form.username = payload.username;
      this.form.display_name = payload.display_name;
      this.form.role = payload.role;
    }
  }
};
</script>
