<template>
  <div>
    <p>{{ msg }}</p>
  </div>

</template>

<script>
  export default {
    name: 'HelloWorld',
    computed: {
      msg() {
        return this.$store.state.msg
      }
    },
    methods: {
      getMessage() {
        this.$store.dispatch("getMessage")
      }
    },
    mounted() {
      this.getMessage()
    }
  }
</script>
