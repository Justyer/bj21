<template>
  <Background>
    <el-row class="login-menu">
      <el-col :span="6">
        <el-input v-model="username" :spellcheck="false" placeholder="V I C T I M" :maxlength="7" :input-style="{'background-color': 'transparent', 'color': '#ff2400'}" prefix-icon="el-icon-knife-fork">
          <template #append>
            <el-button icon="el-icon-right" @click="loginSubmit"></el-button>
          </template>
        </el-input>
        <el-button @click="exitGame" type="text">Exit</el-button>
      </el-col>
    </el-row>
  </Background>
</template>

<script>
import { ipcRenderer } from "electron";
import Background from "@/components/Background.vue";
import { bizKey } from "@/logic/server/logic_conn";
export default {
  name: "Login",
  components: { Background },
  data() {
    return {
      username: "",
    };
  },
  mounted() {
    ipcRenderer.on(bizKey("login"), () => {
      this.$router.push("/home");
    });
  },
  methods: {
    loginSubmit() {
      ipcRenderer.send("mainlogic", {
        cmd: "login",
        text: {
          name: this.username,
          addr: "0.0.0.0:9009",
        },
      });
    },
    exitGame() {
      ipcRenderer.send("mainlogic", {
        cmd: "exit",
        text: {},
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
