<template>
  <div>
    <Header msg="主页"></Header>
    <el-row>
      <el-button @click="gotoSigleMatch" class="button" type="text" disabled>单人匹配</el-button>
      <el-button @click="gotoTableList" class="button" type="text">牌桌列表</el-button>
    </el-row>
  </div>
</template>

<script>
import Header from "../components/Header.vue";
export default {
  name: "Home",
  components: {
    Header,
  },
  data() {
    return {};
  },
  methods: {
    gotoSigleMatch () {
      
    },
    gotoTableList() {
      this.$router.push("/table_list");
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
