<template>
  <Background>
    <el-row class="login-menu">
      <el-col :span="6">
        <el-input
          v-model="username"
          :spellcheck ="false"
          placeholder="V I C T I M"
          :maxlength="7"
          :input-style="{'background-color': 'transparent', 'color': '#ff2400'}"
          prefix-icon="el-icon-knife-fork"
        >
          <template #append>
            <el-button icon="el-icon-right" @click="loginSubmit"></el-button>
          </template>
        </el-input>
      </el-col>
    </el-row>
  </Background>
</template>

<script>
import { ipcRenderer } from "electron";
import Background from "../components/Background.vue";
export default {
  name: "Login",
  components: { Background },
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
.login-menu {
  display: flex;
  justify-content: center;
  align-items: center;
}
</style>
