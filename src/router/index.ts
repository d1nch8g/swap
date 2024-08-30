import { createRouter, createWebHistory } from 'vue-router'
import HomeView from '../views/HomeView.vue'
import AboutView from '../views/AboutView.vue'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      name: 'home',
      component: HomeView
    },
    {
      path: '/ton-rub',
      name: 'ton-rub',
      component: () => AboutView
    },
    {
      path: '/rub-ton',
      name: 'rub-ton',
      component: () => AboutView
    }
  ]
})

export default router
