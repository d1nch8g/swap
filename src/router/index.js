import { createRouter, createWebHistory } from "vue-router";
import Home from "../views/Home.vue";

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
    // Admin views
  ],
});

export default router;
