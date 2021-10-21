<template>
  <div class="login-style">
    <el-row class="login-card">
      <div class="login-logo">Black! Jack! 21</div>
    </el-row>
    <el-row class="login-card">
      <el-col :span="6">
        <el-input
          id="username"
          v-model="username"
          placeholder="V I C T I M"
          :maxlength="7"
          prefix-icon="el-icon-knife-fork"
        >
          <template #append>
            <el-button icon="el-icon-right" @click="loginSubmit"></el-button>
          </template>
        </el-input>
      </el-col>
    </el-row>
  </div>
</template>

<script>
import { ipcRenderer } from "electron";
export default {
  name: "Login",
  components: {},
  data() {
    return {
      username: "",
    };
  },
  mounted() {
    ipcRenderer.on("reply-login", (event, arg) => {
      localStorage.setItem("user-token", arg.text.token);
      this.$router.push("/home");
    });
  },
  methods: {
    loginSubmit() {
      ipcRenderer.send("logic-conn", {
        cmd: "login",
        text: {
          name: this.username,
        },
      });
    },
  },
};
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
#username {
  margin-bottom: 10px;
}

.login-style {
  background-color: #008080;
  width: 100%;
  height: 100%;
  background: url("../assets/home_background.jpeg");
}

.login-logo {
  color: #470024;
  font-size: 100px;
}

.login-card {
  display: flex;
  justify-content: center;
  align-items: center;
}
</style>
