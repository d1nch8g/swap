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
      path: "/amlkyc",
      name: "amlkyc",
      component: () => import("../views/Amlkyc.vue"),
    },
    {
      path: "/chats",
      name: "chats",
      component: () => import("../views/Chats.vue"),
    },
    {
      path: "/chat",
      name: "chat",
      component: () => import("../views/Chat.vue"),
    },
    {
      path: "/profile",
      name: "profile",
      component: () => import("../views/Profile.vue"),
    },
    {
      path: "/operator",
      name: "operator",
      component: () => import("../views/Operator.vue"),
    },
    {
      path: "/currencies",
      name: "currencies",
      component: () => import("../views/Currencies.vue"),
    },
    {
      path: "/exchangers",
      name: "exchangers",
      component: () => import("../views/Exchangers.vue"),
    },
    {
      path: "/card-confirmations",
      name: "card-confirmations",
      component: () => import("../views/CardConfirmations.vue"),
    },
    {
      path: "/transfer",
      name: "transfer",
      component: () => import("../views/Transfer.vue"),
    },
    {
      path: "/validate-card",
      name: "validate-card",
      component: () => import("../views/ValidateCard.vue"),
    },
    {
      path: "/order",
      name: "order",
      component: () => import("../views/Order.vue"),
    },
    {
      path: "/orders",
      name: "orders",
      component: () => import("../views/Orders.vue"),
    },
  ],
});

export default router;
