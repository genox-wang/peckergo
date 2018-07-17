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
              <Option :value="1">管理员</Option>
              <Option :value="2">操作员</Option>
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
      routeName: '',
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
    },
    reset () {
      let self = this;
      this.$store.dispatch('get_user_by_id', this.$route.params.id).then((data) => {
        if (data) {
          self.form.id = data.id;
          self.form.username = data.username;
          self.form.display_name = data.display_name;
          self.form.role = data.role;
        }
      });
    }
  },

  watch: {
    '$route' (to, from) {
      if (this._.isEqual(this.routeName, this.$route.name)) {
        this.reset();
      }
    }
  },

  mounted () {
    if (this._.isEqual(this.routeName, '')) {
      this.routeName = this.$route.name;
      this.reset();
    }
  }
};
</script>
