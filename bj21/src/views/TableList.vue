<template>
  <div>
    <Hall msg="牌桌列表"></Hall>
    <el-row>
      <el-col :span="4" v-for="table in table_list" :key="table.name" style="padding: 5px">
        <el-card class="box-card" shadow="hover">
          <template #header>
            <div class="card-header">
              <el-button @click="gotoTable(table.seq)" v-model="table.seq" class="button" type="text">{{ table.name }}</el-button>
            </div>
          </template>
          <el-tag type="success">P1:{{ table.player1.name }}</el-tag>
          <el-tag type="success">P2:{{ table.player2.name }}</el-tag>
        </el-card>
      </el-col>
    </el-row>
  </div>
</template>

<script>
import { ipcRenderer } from "electron";
import { ElMessage } from "element-plus";
import Hall from "../components/Hall.vue";
export default {
  name: "TableName",
  created() {
    this.getTableList();
  },
  components: {
    Hall,
  },
  mounted() {
    ipcRenderer.on("reply-tablelist", (event, arg) => {
      let table_list = [];
      arg.text.tables.forEach((table) => {
        let p1_name, p2_name;
        if (typeof table.p1 !== "undefined") {
          p1_name = table.p1.name;
        }
        if (typeof table.p2 !== "undefined") {
          p2_name = table.p2.name;
        }
        table_list.push({
          name: table.name,
          seq: table.seq,
          player1: {
            name: p1_name,
          },
          player2: {
            name: p2_name,
          },
        });
      });
      this.table_list = table_list;
    });
    ipcRenderer.on("reply-sitdown", (event, arg) => {
      if (typeof arg.text.err === "undefined") {
        this.$router.push({
          name: "Table",
          params: { seq: arg.text.table_seq },
        });
      } else {
        ElMessage(arg.text.err);
      }
    });
  },
  data() {
    return {
      table_list: [
        {
          name: "kamistable",
          player1: {
            name: "dxc",
          },
          player2: {
            name: "zxy",
          },
        },
      ],
    };
  },
  methods: {
    getTableList() {
      ipcRenderer.send("logic-conn", {
        cmd: "tablelist",
        text: {},
      });
    },
    gotoTable(seq) {
      ipcRenderer.send("logic-conn", {
        cmd: "sitdown",
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
</style>
