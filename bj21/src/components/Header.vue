<template>
  <div>
    <h1 style="margin:0">{{ msg }}</h1>
    <el-divider style="margin-top: 10px; margin-bottom: 10px"></el-divider>
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