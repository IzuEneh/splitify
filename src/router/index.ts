import { createRouter, createWebHistory } from 'vue-router'
import HomeView from '../views/HomeView.vue'
import Profile from '@/views/ProfileView.vue'
import PlaylistView from '@/views/PlaylistView.vue'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      name: 'home',
      component: HomeView
    },
    {
      path: '/callback',
      name: 'profile',
      component: Profile
    },
    {
      path: '/playlist/:id',
      name: 'playlist',
      component: PlaylistView
    }
  ]
})

export default router
