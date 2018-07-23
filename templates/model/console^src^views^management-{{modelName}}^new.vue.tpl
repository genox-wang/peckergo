<template>
  <edit-base :form="form" :rules="rules" :save-promise="save">
    <Row>
      <!-- TODO -->
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

export default {
  components: {
    editBase
  },
  data () {
    return {
      form: {
        // TODO
      },
      rules: {
        // TODO
        // username: [
        //   { required: true, message: '账号不能为空', trigger: 'blur' }
        // ],
      }
    };
  },
  methods: {
    save () {
      // TODO
      return this.$store.dispatch('create_{{modelName}}', this.form);
    }
  }
};
</script>
