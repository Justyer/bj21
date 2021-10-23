<template>
  <Background>
    <el-row class="home-menu">
      <el-button style="color: #ff2400;" :size="medium" type="text">
        <strong>V i c t i m -> {{ name }}</strong>
      </el-button>
      <el-button @click="gotoSigleMatch" type="text">Single Match</el-button>
      <el-button @click="gotoTableList" type="text">Table List</el-button>
      <el-button @click="backToLogin" type="text">Back to Login</el-button>
      <el-button @click="exitGame" type="text">Exit</el-button>
    </el-row>
  </Background>
</template>

<script>
import { ipcRenderer } from "electron";
import { ElMessage } from "element-plus";
import Background from "../components/Background.vue";
export default {
  name: "Home",
  components: { Background },
  created() {
    const userinfo = ipcRenderer.sendSync("backend-sys", {
      cmd: "get-user-info",
      text: {},
    });
    this.name = userinfo.name;
  },
  mounted() {
    ipcRenderer.on("reply-logout", (event, arg) => {
      if (typeof arg.text.err === "undefined") {
        this.$router.push("/login");
      } else {
        ElMessage(arg.text.err);
      }
    });
  },
  data() {
    return { name: "" };
  },
  methods: {
    gotoSigleMatch() {},
    gotoTableList() {
      this.$router.push("/table_list");
    },
    backToLogin() {
      ipcRenderer.send("logic-conn", {
        cmd: "logout",
        text: {},
      });
    },
    exitGame() {
      ipcRenderer.send("backend-sys", {
        cmd: "exit",
        text: {},
      });
    },
  },
};
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
.home-menu {
  height: 100%;
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;
}
</style>
