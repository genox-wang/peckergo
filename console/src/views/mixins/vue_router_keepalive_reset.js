export default {
  data () {
    return {
      routeName: ''
    };
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
