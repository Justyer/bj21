import { createRouter, createWebHashHistory } from "vue-router";
import Login from "../views/Login.vue";
import Home from "../views/Home.vue";
import TableList from "../views/TableList.vue";
import Table from "../views/Table.vue";

const routes = [
  {
    path: "/",
    name: "Root",
    redirect: "/login",
  },
  {
    path: "/login",
    name: "Login",
    component: Login,
  },
  {
    path: "/home",
    name: "Home",
    component: Home,
  },
  {
    path: "/table_list",
    name: "TableList",
    component: TableList,
  },
  {
    path: "/table/:seq",
    name: "Table",
    component: Table,
  },
  // {
  //   path: "/about",
  //   name: "About",
  //   component: About,
  //   // route level code-splitting
  //   // this generates a separate chunk (about.[hash].js) for this route
  //   // which is lazy-loaded when the route is visited.
  //   // component: () =>
  //   //   import("../views/About.vue"),
  // },
];

const router = createRouter({
  history: createWebHashHistory(),
  routes,
});

export default router;
