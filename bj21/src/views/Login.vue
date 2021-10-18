<template>
  <div>
    <Hall msg="登录"></Hall>
    <el-row class="login-card">
      <el-col :span="8">
        <el-card class="box-card" shadow="hover">
          <el-input id="username" v-model="username" placeholder="Username" prefix-icon="el-icon-user-solid" />
          <el-button type="success" icon="el-icon-check" @click="loginSubmit" circle></el-button>
        </el-card>
      </el-col>
    </el-row>
  </div>
</template>

<script>
import { ipcRenderer } from "electron";
import Hall from "../components/Hall.vue";
export default {
  name: "Login",
  components: {
    Hall,
  },
  data() {
    return {
      username: "",
    };
  },
  mounted() {
    ipcRenderer.on("biz-login", (event, arg) => {
      localStorage.setItem("user-token", arg.text.token);
      this.$router.push("/home");
    });
  },
  methods: {
    loginSubmit() {
      let reply = ipcRenderer.send("logic-conn", {
        cmd: "login",
        text: {
          name: this.username,
        },
      });
      console.log("reply", reply);
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
