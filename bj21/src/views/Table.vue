<template>
  <div class="table-style">
    <el-row class="login-card">
      <el-col style="padding: 10px">
        <el-card shadow="always" :body-style="{'background-color': 'transparent'}">
          <el-row>
            <el-col :span="3" v-for="o in 8" :key="o" style="padding: 5px;">
              <el-card v-if="o === 1" shadow="hover" :body-style="cardBackStyle">
                <div>!</div>
                <div class="name-in-card">{{ table.player_you.name }}</div>
              </el-card>
              <el-card v-else shadow="hover" :body-style="cardFrontStyle">{{ o }}</el-card>
            </el-col>
          </el-row>
          <el-divider></el-divider>
          <el-button-group>
            <el-button @click="exitTable(this.table.seq)" class="button" icon="el-icon-back"></el-button>
            <el-tooltip :content="table.seq" placement="top" effect="light" :show-after="1000">
              <el-button class="button">
                <i class="el-icon-arrow-left"></i>
                {{ table.name }}
                <i class="el-icon-arrow-right"></i>
              </el-button>
            </el-tooltip>
            <el-button @click="startGame(this.table.seq)" class="button" icon="el-icon-caret-right"></el-button>
          </el-button-group>
          <el-divider></el-divider>
          <el-row>
            <el-col :span="3" v-for="o in 8" :key="o" style="padding: 5px">
              <el-card v-if="o === 1" shadow="hover" :body-style="cardBackStyle">
                <div>{{ o+3 }}</div>
                <div class="name-in-card">{{ table.player_me.name }}</div>
              </el-card>
              <el-card v-else shadow="hover" :body-style="cardFrontStyle">{{ o+3 }}</el-card>
            </el-col>
          </el-row>
        </el-card>
      </el-col>
    </el-row>
  </div>
</template>

<script>
import { ipcRenderer } from "electron";
import { ElMessage } from "element-plus";
export default {
  name: "Table",
  components: {},
  created() {
    this.table.seq = this.$route.params.seq;
    this.getTableInfo(this.table.seq);
  },
  mounted() {
    ipcRenderer.on("reply-tableinfo", (event, arg) => {
      let pyou_name = "<nil>",
        pme_name = "<nil>";
      if (typeof arg.text.table.p1 !== "undefined") {
        if (arg.text.table.me == 1) {
          pme_name = arg.text.table.p1.name;
        } else {
          pyou_name = arg.text.table.p1.name;
        }
      }
      if (typeof arg.text.table.p2 !== "undefined") {
        if (arg.text.table.me == 2) {
          pme_name = arg.text.table.p2.name;
        } else {
          pyou_name = arg.text.table.p2.name;
        }
      }
      this.table = {
        name: arg.text.table.name,
        seq: arg.text.table.seq,
        player_you: {
          name: pyou_name,
        },
        player_me: {
          name: pme_name,
        },
      };
    });
    ipcRenderer.on("reply-standup", (event, arg) => {
      if (typeof arg.text.err === "undefined") {
        this.$router.push({ name: "TableList" });
      } else {
        ElMessage(arg.text.err);
      }
    });
    ipcRenderer.on("reply-table-action", () => {
      this.getTableInfo(this.table.seq);
    });
  },
  data() {
    return {
      table: {
        name: "[unkown name]",
        seq: "[unkown seq]",
        player_you: {
          name: "<nil>",
        },
        player_me: {
          name: "<nil>",
        },
      },
      cardFrontStyle: {
        height: "150px",
        display: "flex",
        // "flex-direction": "column",
        "justify-content": "center",
        "align-items": "center",
        "background-color": "#d2d9e4",
        color: "#f08080",
        // "text-shadow": "0.1em 0.1em 0.15em #d2d9e4",
        "font-size": "100px",
      },
      cardBackStyle: {
        height: "190px",
        padding: "0",
        display: "flex",
        "flex-direction": "column",
        "justify-content": "center",
        "align-items": "center",
        "background-color": "#4768b3",
        color: "#f08080",
        "text-shadow": "0.1em 0.1em 0.15em #d2d9e4",
        "font-size": "50px",
      },
    };
  },
  methods: {
    getTableInfo(seq) {
      ipcRenderer.send("logic-conn", {
        cmd: "tableinfo",
        text: {
          table_seq: seq,
        },
      });
    },
    startGame() {},
    exitTable(seq) {
      ipcRenderer.send("logic-conn", {
        cmd: "standup",
        text: {
          table_seq: seq,
        },
      });
    },
  },
};
</script>
<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
.login-card {
  display: flex;
  justify-content: center;
  align-items: center;
}
.name-in-card {
  font-size: 30px;
}
.table-style {
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;
  background-color: #008080;
  width: 100%;
  height: 100%;
  /* filter: blur(5px); */
  background: url("../assets/home_background.jpeg");
}
</style>
