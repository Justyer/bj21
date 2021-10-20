<template>
  <div>
    <Header msg="登录"></Header>
    <el-row class="login-card">
      <el-col :span="8">
        <el-card class="box-card" shadow="hover">
          <el-input :input-style="{'margin-bottom': '10px'}" id="username" v-model="username" placeholder="Username" prefix-icon="el-icon-user-solid" />
          <el-button icon="el-icon-check" @click="loginSubmit" circle></el-button>
        </el-card>
      </el-col>
    </el-row>
  </div>
</template>

<script>
import { ipcRenderer } from "electron";
import Header from "../components/Header.vue";
export default {
  name: "Login",
  components: {
    Header,
  },
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
.login-card {
  display: flex;
  justify-content: center;
  align-items: center;
}
</style>
