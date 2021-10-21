<template>
  <div>
    <h1 style="margin:0">{{ msg }}</h1>
    <el-divider class="header-divider"></el-divider>
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
.header-divider {
  margin-top: 12px;
  margin-bottom: 0;
}
</style>