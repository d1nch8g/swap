import { createRouter, createWebHistory } from 'vue-router'
import HomeView from '../views/Home.vue'
import Contacts from '@/views/Contacts.vue'
import Description from '@/views/Description.vue'
import Order from '@/views/Order.vue'
import Transfer from '@/views/Transfer.vue'
import Verify from '@/views/Verify.vue'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      name: 'home',
      component: HomeView
    },
    {
      path: '/contacts',
      name: 'contacts',
      component: Contacts
    },
    {
      path: '/description',
      name: 'description',
      component: Description
    },
    {
      path: '/order',
      name: 'order',
      component: Order
    },
    {
      path: '/transfer',
      name: 'transfer',
      component: Transfer
    },
    {
      path: '/verify',
      name: 'verify',
      component: Verify
    }
  ]
})

export default router
