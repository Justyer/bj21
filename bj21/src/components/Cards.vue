<template>
  <el-row>
    <el-col :span="3" v-for="(card, idx) in player.hands" :key="card.id" style="padding: 5px;">
      <el-card v-if="idx === 0" shadow="hover" :body-style="cardBackStyle">
        <div v-if="card.num === 0">?</div>
        <div v-else>{{ card.num }}</div>
        <div class="name-in-card">{{ player.name }}</div>
        <div class="name-in-card">{{ player.sum_text }}</div>
      </el-card>
      <el-card v-else shadow="hover" v-model="card.id" :body-style="cardFrontStyle">{{ card.num }}</el-card>
    </el-col>
  </el-row>
</template>

<script>
export default {
  name: "Card",
  props: {
    player: {},
  },
  created() {},
  data() {
    return {
      cardFrontStyle: {
        height: "150px",
        display: "flex",
        "justify-content": "center",
        "align-items": "center",
        "background-color": "#d2d9e4",
        color: "#f08080",
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
};
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
.name-in-card {
  font-size: 30px;
}
</style>