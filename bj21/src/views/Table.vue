<template>
  <div class="table-style">
    <el-row class="table-card">
      <el-col style="padding: 10px">
        <el-card shadow="always">
          <Cards :player="table.player_you"></Cards>
          <el-divider></el-divider>

          <el-button-group>
            <el-button @click="exitTable(table.seq)" class="button" icon="el-icon-back">Escape</el-button>
            <el-tooltip :content="table.seq" placement="top" effect="light" :show-after="1000">
              <el-button class="button">
                <i class="el-icon-arrow-left"></i>
                {{ table.name }}
                <i class="el-icon-arrow-right"></i>
              </el-button>
            </el-tooltip>
            <el-button @click="startGame(table.seq)" class="button" :icon="controller.switchStartEndIcon">{{ controller.switchStartEndText }}</el-button>
            <el-button @click="playerHit(table.seq)" class="button" icon="el-icon-thumb">Hit</el-button>
            <el-button @click="playerStand(table.seq)" class="button" icon="el-icon-hot-water">Stand</el-button>
          </el-button-group>

          <el-divider></el-divider>
          <Cards :player="table.player_me"></Cards>
        </el-card>
      </el-col>
    </el-row>
  </div>
</template>

<script>
import { ElMessage } from "element-plus";
import { onTableInfo } from "@/logic/render/table";
import { logicConnSender, logicConnReply } from "@/logic/render/sender";
import Cards from "@/components/Cards.vue";
import { isUndefined } from "@/utils/judge";
export default {
  name: "Table",
  components: { Cards },
  created() {
    this.table.seq = this.$route.params.seq;
    this.getTableInfo(this.table.seq);
  },
  mounted() {
    onTableInfo((table) => {
      if (table.status === "idle") {
        this.controller.switchStartEndIcon = "el-icon-video-play";
        this.controller.switchStartEndText = "Start Game";
      } else {
        this.controller.switchStartEndIcon = "el-icon-video-pause";
        this.controller.switchStartEndText = "Surrender";
      }
      this.table = table;
    });
    logicConnReply("standup", (event, arg) => {
      if (!isUndefined(arg.text.err)) {
        ElMessage(arg.text.err);
        return;
      }
      this.$router.push({ name: "TableList" });
    });
    logicConnReply("table-action", () => {
      this.getTableInfo(this.table.seq);
    });
    logicConnReply("startgame", (event, arg) => {
      if (!isUndefined(arg.text.err)) {
        ElMessage(arg.text.err);
        return;
      }
      ElMessage("table start!");
    });
  },
  data() {
    return {
      table: {
        name: "[unkown name]",
        seq: "[unkown seq]",
        player_you: {
          name: "<nil>",
          hands: [],
          sum_text: "?/21",
        },
        player_me: {
          name: "<nil>",
          hands: [],
          sum_text: "?/21",
        },
      },
      controller: {
        switchStartEndIcon: "el-icon-video-play",
        switchStartEndText: "Start Game",
      },
    };
  },
  methods: {
    getTableInfo(seq) {
      logicConnSender({
        cmd: "tableinfo",
        text: {
          table_seq: seq,
        },
      });
    },
    startGame(seq) {
      logicConnSender({
        cmd: "startgame",
        text: {
          table_seq: seq,
        },
      });
    },
    playerHit(seq) {
      logicConnSender({
        cmd: "playerhit",
        text: { table_seq: seq },
      });
    },
    playerStand(seq) {
      logicConnSender({
        cmd: "playerstand",
        text: { table_seq: seq },
      });
    },
    exitTable(seq) {
      logicConnSender({
        cmd: "standup",
        text: { table_seq: seq },
      });
    },
  },
};
</script>
<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
.table-card {
  width: 100%;
  display: flex;
  justify-content: center;
  align-items: center;
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
