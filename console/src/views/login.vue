<style lang="less">
  @import './login.less';
</style>

<template>
  <div class="login" @keydown.enter="handleSubmit">
    <div class="login-con">
      <Card :bordered="false">
        <p slot="title">
          <Icon type="md-log-in"></Icon>
          欢迎登录
        </p>
        <div class="form-con">
          <Form ref="loginForm" :model="form" :rules="rules">
            <FormItem prop="userName">
              <Input v-model="form.userName" placeholder="请输入用户名">
              <span slot="prepend">
                <Icon :size="16" type="person"></Icon>
              </span>
              </Input>
            </FormItem>
            <FormItem prop="password">
              <Input type="password" v-model="form.password" placeholder="请输入密码">
              <span slot="prepend">
                <Icon :size="14" type="locked"></Icon>
              </span>
              </Input>
            </FormItem>
            <Row>
              <Col span="16">
                <img :src="captcha" @click="resetCaptcha"/>
              </Col>
              <Col span="8">
                <Input v-model="captchaValue" placeholder="验证码"/>
              </Col>
            </Row>

            <input type="hidden" :value="captchaKey"/>

            <FormItem>
              <Button @click="handleSubmit" type="primary" long>登录</Button>
            </FormItem>
          </Form>

        </div>
      </Card>
    </div>
  </div>
</template>

<script>
// import Cookies from 'js-cookie';
export default {
  data () {
    return {
      captchaKey: '',
      captcha: '',
      captchaValue: '',
      form: {
        userName: '',
        password: ''
      },
      rules: {
        userName: [{
          required: true,
          message: '账号不能为空',
          trigger: 'blur'
        }],
        password: [{
          required: true,
          message: '密码不能为空',
          trigger: 'blur'
        }]
      }
    };
  },
  methods: {
    handleSubmit () {
      this.$refs.loginForm.validate((valid) => {
        if (valid) {
          this.$store.dispatch('login', {
            username: this.form.userName,
            password: this.form.password,
            Captcha_key: this.captchaKey,
            Captcha: this.captchaValue
          })
            .then(() => {
              this.$store.commit('setAvator',
                'https://ss1.bdstatic.com/70cFvXSh_Q1YnxGkpoWK1HF6hhy/it/u=3448484253,3685836170&fm=27&gp=0.jpg'
              );
              this.$router.push({
                name: 'home_index'
              });
            })
            .catch(() => {
              this.resetCaptcha();
            });
        }
      });
    },
    resetCaptcha () {
      let self = this;
      this.$store.dispatch('get_captcha').then((data) => {
        self.captchaKey = data.key;
        self.captcha = data.captcha;
      });
    }
  },
  mounted () {
    this.resetCaptcha();
  }
};
</script>

<style>
</style>
