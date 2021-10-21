<template>
  <div class="home-style">
    <el-row class="home-menu">
      <div class="home-logo">Black! Jack! 21</div>
    </el-row>
    <el-row class="home-menu">
      <el-button @click="gotoSigleMatch" style="color: white" type="text">Single Match</el-button>
      <el-button @click="gotoTableList" style="color: white" type="text">Table List</el-button>
      <el-button @click="backToLogin" style="color: white" type="text">Back to Login</el-button>
    </el-row>
  </div>
</template>

<script>
export default {
  name: "Home",
  components: {},
  data() {
    return {};
  },
  methods: {
    gotoSigleMatch() {},
    gotoTableList() {
      this.$router.push("/table_list");
    },
    backToLogin() {
      this.$router.push("/login");
    },
  },
};
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
.home-style {
  display: flex;
  flex-direction: column;
  /* justify-content: center;
  align-items: center; */
  background-color: #008080;
  width: 100%;
  height: 100%;
  /* filter: blur(5px); */
  background: url("../assets/home_background.jpeg");
}

.home-logo {
  color: #470024;
  font-size: 100px;
}

.home-menu {
  display: flex;
  flex-direction: column;
  justify-content: left;
}
</style>
