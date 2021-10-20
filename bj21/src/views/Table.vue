<template>
  <div>
    <Header msg="牌桌"></Header>
    <el-row class="login-card">
      <el-col style="padding: 10px">
        <el-card shadow="always">
          <el-row>
            <el-col :span="3" v-for="o in 8" :key="o" style="padding: 5px">
              <el-card shadow="hover" :body-style="{height: '150px', display: 'flex', 'justify-content': 'center', 'align-items': 'center'}">{{ o }}</el-card>
            </el-col>
          </el-row>
          <el-divider></el-divider>
          <el-button-group>
            <el-button @click="exitTable(this.table.seq)" class="button" icon="el-icon-back"></el-button>
            <el-button class="button">{{ table.player_me.name }}</el-button>
            <el-button class="button">{{ table.name }}:{{ table.seq }}</el-button>
            <el-button class="button">{{ table.player_you.name }}</el-button>
            <el-button @click="startGame(this.table.seq)" class="button" icon="el-icon-caret-right"></el-button>
          </el-button-group>
          <el-divider></el-divider>
          <el-row>
            <el-col :span="3" v-for="o in 8" :key="o" style="padding: 5px">
              <el-card shadow="hover" :body-style="{height: '150px', display: 'flex', 'justify-content': 'center', 'align-items': 'center'}">{{ o }}</el-card>
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
import Header from "../components/Header.vue";
export default {
  name: "Table",
  components: {
    Header,
  },
  created() {
    this.table.seq = this.$route.params.seq;
    this.getTableInfo(this.table.seq);
  },
  mounted() {
    ipcRenderer.on("reply-tableinfo", (event, arg) => {
      let pyou_name = "",
        pme_name = "";
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
/* h3 {
  margin: 40px 0 0;
}
ul {
  list-style-type: none;
  padding: 0;
}
li {
  display: inline-block;
  margin: 0 10px;
}
a {
  color: #42b983;
} */
.login-card {
  display: flex;
  justify-content: center;
  align-items: center;
}
</style>
