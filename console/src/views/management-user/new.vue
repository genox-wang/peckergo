<template>
  <edit-base ref="base" :form="form" :rules="rules" :save-promise="save">
    <Row>
        <Col span="16">
          <FormItem label="用户名" prop="username" >
            <Input placeholder="用户名" v-model="form.username" @on-enter.stop="$refs.base.save()"></Input>
          </FormItem>
        </Col>
        <Col span="16">
          <FormItem label="昵称" prop="display_name">
            <Input placeholder="昵称" v-model="form.display_name" @on-enter.stop="$refs.base.save()"></Input>
          </FormItem>
        </Col>
        <Col span="16">
          <FormItem label="新密码" prop="password">
            <Input placeholder="新密码" v-model="form.password" @on-enter.stop="$refs.base.save()"></Input>
          </FormItem>
        </Col>
        <Col span="16">
          <FormItem label="再次输入" prop="confirm">
            <Input placeholder="再次输入" v-model="form.confirm" @on-enter.stop="$refs.base.save()"></Input>
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
    const validatePassord = (rule, value, callback) => {
      if (value === '') {
        callback(new Error('请输入密码'));
      } else {
        if (this.form.confirm !== '') {
          // 对第二个密码框单独验证
          this.$refs.base.$refs.form.validateField('confirm');
        }
        callback();
      }
    };
    const validateConfirm = (rule, value, callback) => {
      if (value === '') {
        callback(new Error('请再次输入密码'));
      } else if (value !== this.form.password) {
        callback(new Error('两次输入的密码不一致!'));
      } else {
        callback();
      }
    };
    return {
      form: {
        username: '',
        password: '',
        display_name: '',
        role: 1
      },
      rules: {
        username: [
          { required: true, message: '账号不能为空', trigger: 'blur' }
        ],
        display_name: [
          { required: true, message: '昵称不能为空', trigger: 'blur' }
        ],
        password: [
          { validator: validatePassord, trigger: 'blur' }
        ],
        confirm: [
          { validator: validateConfirm, trigger: 'blur' }
        ]
      }
    };
  },
  methods: {
    save () {
      return this.$store.dispatch('create_user', this.form);
    }
  }
};
</script>
