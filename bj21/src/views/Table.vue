<template>
  <div>
    <Hall msg="牌桌"></Hall>
    <el-row class="login-card">
      <el-col style="padding: 5px">
        <el-card class="box-card" shadow="hover">
          <template #header>
            <div class="card-header">
              <el-button @click="gotoHome" class="button" type="text">{{ table.name }}:{{ table.seq }}</el-button>
            </div>
          </template>
          <el-tag type="success">P1:{{ table.player1.name }}</el-tag>
          <el-tag type="success">P2:{{ table.player2.name }}</el-tag>
        </el-card>
        <router-link to="/table_list">TableList</router-link>
      </el-col>
    </el-row>
  </div>
</template>

<script>
import { ipcRenderer } from "electron";
import Hall from "../components/Hall.vue";
export default {
  name: "Table",
  components: {
    Hall,
  },
  created() {
    this.getTableInfo(this.$route.params.seq);
    console.log(this.$route.params.seq);
  },
  mounted() {
    ipcRenderer.on("reply-tableinfo", (event, arg) => {
      console.log(arg);
      let p1_name, p2_name;
      if (typeof arg.text.table.p1 !== "undefined") {
        p1_name = arg.text.table.p1.name;
      }
      if (typeof arg.text.table.p2 !== "undefined") {
        p2_name = arg.text.table.p2.name;
      }
      this.table = {
        name: arg.text.table.name,
        seq: arg.text.table.seq,
        player1: {
          name: p1_name,
        },
        player2: {
          name: p2_name,
        },
      };
      console.log(this.table);
    });
  },
  data() {
    return {
      seq: "hflshgios",
      table: {
        name: "kamistable",
        player1: {
          name: "dxc",
        },
        player2: {
          name: "zxy",
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
    gotoHome() {
      this.$router.push("/");
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
