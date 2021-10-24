<template>
  <Background>
    <el-row class="tablelist-menu">
      <el-button @click="backToHome" type="text">Back to Home</el-button>
    </el-row>
    <el-row>
      <el-table
        id="tablelist"
        :data="table_list"
        empty-text="No Tables"
        max-height="200"
        :header-cell-style="{
          'background-color': 'transparent',
          color: 'white',
        }"
        :header-row-style="{
          'background-color': 'transparent',
          color: 'white',
        }"
        :row-style="{
          'background-color': 'transparent',
          color: 'white',
        }"
        :cell-style="{
          'background-color': 'transparent',
          color: 'white',
          border: 'none',
        }"
      >
        <el-table-column align="center" prop="seq" label="Table Id" />
        <el-table-column align="center" label="Table Name">
          <template #default="scope">
            <el-button
              style="color: #af111c"
              @click="gotoTable(scope.row.seq)"
              v-model="scope.row.seq"
              class="button"
              type="text"
            >{{ scope.row.name }}</el-button>
          </template>
        </el-table-column>
        <el-table-column align="center" prop="player1.name" label="P1" />
        <el-table-column align="center" prop="player2.name" label="P2" />
      </el-table>
    </el-row>
  </Background>
</template>

<script>
import { ipcRenderer } from "electron";
import { ElMessage } from "element-plus";
import Background from "@/components/Background.vue";
import { bizKey } from '@/logic/server/logic_conn';
export default {
  name: "TableName",
  created() {
    this.getTableList();
  },
  components: { Background },
  mounted() {
    ipcRenderer.on(bizKey("tablelist"), (event, arg) => {
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
    ipcRenderer.on(bizKey("sitdown"), (event, arg) => {
      if (typeof arg.text.err === "undefined") {
        this.$router.push({
          name: "Table",
          params: { seq: arg.text.table.seq },
        });
      } else {
        ElMessage(arg.text.err);
      }
    });
  },
  data() {
    return {
      table_list: [],
    };
  },
  methods: {
    getTableList() {
      ipcRenderer.send("logicconn", {
        cmd: "tablelist",
        text: {},
      });
    },
    gotoTable(seq) {
      ipcRenderer.send("logicconn", {
        cmd: "sitdown",
        text: {
          table_seq: seq,
        },
      });
    },
    backToHome() {
      this.$router.push("/home");
    },
  },
};
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
.tablelist-menu {
  display: flex;
  flex-direction: column;
  justify-content: left;
}

#tablelist {
  width: 100%;
  margin: 20px;
  color: red;
  border: none;
  background-color: transparent;
}
</style>
