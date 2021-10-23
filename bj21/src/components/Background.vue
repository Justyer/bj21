<template>
  <div style="height: 100%">
    <div class="bk-style">
      <el-row class="bk-menu">
        <div class="bk-logo">Black! Jack! 21</div>
      </el-row>
      <slot></slot>
    </div>
  </div>
</template>

<script>
export default {
  name: "Background",
};
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
.bk-style {
  display: flex;
  flex-direction: column;
  background-color: #008080;
  width: 100%;
  height: 100%;
  background: url("../assets/home_background.jpeg");
}

.bk-menu {
  display: flex;
  flex-direction: column;
  justify-content: left;
}

.bk-logo {
  color: #00477d;
  font-size: 100px;
}
</style>