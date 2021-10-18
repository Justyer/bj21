<template>
  <div>
    <Hall msg="主页"></Hall>
    <el-row>
      <el-col :span="4" v-for="o in 18" :key="o" style="padding: 5px">
        <el-card class="box-card" shadow="hover">
          <template #header>
            <div class="card-header">
              <el-button @click="gotohome" class="button" type="text">Kami's Table</el-button>
            </div>
          </template>
          <el-tag v-for="o in 2" :key="o" type="success">Tag 2</el-tag>
        </el-card>
      </el-col>
    </el-row>
  </div>
</template>

<script>
import Hall from "../components/Hall.vue";
export default {
  name: "Home",
  components: {
    Hall,
  },
  data() {
    return {};
  },
  methods: {
    gotohome() {
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
</style>
