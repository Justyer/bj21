<template>
  <div>
    <h1>{{ msg }}</h1>
    <el-divider></el-divider>
  </div>
</template>

<script>
export default {
  name: "Hall",
  props: {
    msg: String,
  },
  setup() {},
  methods: {},
};
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
</style>