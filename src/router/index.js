import { createRouter, createWebHistory } from "vue-router";
import Home from "../views/Main.vue";

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    // User views
    {
      path: "/",
      name: "home",
      component: Home,
    },
    {
      path: "/contacts",
      name: "contacts",
      component: () => import("../views/Contacts.vue"),
    },
    {
      path: "/login",
      name: "login",
      component: () => import("../views/Login.vue"),
    },
    {
      path: "/register",
      name: "register",
      component: () => import("../views/Register.vue"),
    },
    {
      path: "/rules",
      name: "rules",
      component: () => import("../views/Rules.vue"),
    },
    {
      path: "/orders",
      name: "orders",
      component: () => import("../views/Orders.vue"),
    },
  ],
});

export default router;
