<template>
  <div>
    <Form ref="form" :label-width="120" :model="form" :rules="rules" @submit.native.prevent>
      <slot></slot>
    </Form>
    <br>
    <Button type="primary" shape="circle" @click="save" ghost>保存</Button>
  </div>
</template>

<script>
export default {
  props: {
    // form 用于验证的表单数据
    form: {},
    // rules 验证规则 参考https://github.com/yiminghe/async-validator#async-validator
    rules: {},
    // successMessage 表单提交成功反馈文本，默认'保存成功'
    successMessage: '',
    // nextRoute 表单提交成功跳转路由， 默认返回前一个页面
    nextRoute: '',
    // savePromise 表单提交Promise
    savePromise: null
  },
  data () {
    return {
      routeName: ''
    };
  },
  methods: {
    save () {
      this.$refs.form.validate((valid) => {
        if (valid) {
          if (this.savePromise) {
            this.savePromise().then(() => {
              let message = this.successMessage ? this.successMessage : '保存成功';
              this.$Message.success(message);
              this.close();
            });
          } else {
            this.close();
          }
        }
      });
    },
    close () {
      this.$store.commit('removeTag', this.routeName);
      this.$store.commit('closePage', this.routeName);
      if (this.nextRoute) {
        this.$router.push({name: this.nextRoute});
      } else {
        this.$router.go(-1);
      }
    }
  },
  mounted () {
    this.routeName = this.$route.name;
  }
};
</script>
