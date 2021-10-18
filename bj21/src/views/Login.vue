<template>
  <div class="hello">
    <h1>{{ msg }}</h1>
    <el-divider></el-divider>
    <el-row>
      <el-col :span="4" v-for="o in 18" :key="o" style="padding: 5px">
        <el-card class="box-card" shadow="hover">
          <template #header>
            <div class="card-header">
              <el-button @click="gotohome(this)" class="button" type="text">Kami's Table</el-button>
            </div>
          </template>
          <div v-for="o in 2" :key="o" class="text item">{{ 'P' + o + ': <nil>' }}</div>
        </el-card>
      </el-col>
    </el-row>
  </div>
</template>

<script>
import { ipcRenderer } from 'electron'
export default {
  name: "Login",
  props: {
    msg: String,
  },
  methods: {
    gotohome: (self) => {
      self.$router.push('/')
      ipcRenderer.sendSync('')
    }
  }
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
