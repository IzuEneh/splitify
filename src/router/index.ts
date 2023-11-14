import { createRouter, createWebHistory } from 'vue-router'
import HomeView from '../views/HomeView.vue'
import Profile from '@/views/ProfileView.vue'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      name: 'home',
      component: HomeView,
      beforeEnter: () => {
        const accessToken = localStorage.getItem('access_token')
        if (accessToken) {
          return { path: '/callback' }
        }

        return true
      }
    },
    {
      path: '/callback',
      name: 'profile',
      component: Profile,
      beforeEnter: async (to) => {
        const accessToken = localStorage.getItem('access_token')
        if (to.query && !accessToken) {
          const code = to.query.code as string
          const verifier = localStorage.getItem('verifier')

          const params = new URLSearchParams()
          params.append('client_id', import.meta.env.VITE_CLIENT_ID)
          params.append('grant_type', 'authorization_code')
          params.append('code', code)
          params.append('redirect_uri', 'http://localhost:5173/callback')
          params.append('code_verifier', verifier!)

          try {
            const result = await fetch('https://accounts.spotify.com/api/token', {
              method: 'POST',
              headers: { 'Content-Type': 'application/x-www-form-urlencoded' },
              body: params
            })

            const { access_token, refresh_token, expires_in } = await result.json()
            localStorage.setItem('access_token', access_token)
            localStorage.setItem('refresh_token', refresh_token)
            localStorage.setItem('expires_in', expires_in)
            localStorage.setItem('access_code_fetched_date', Date.now().toString())
          } catch (error) {
            console.log(`An error occcured getting access token: ${error}`)
            return false
          }
        }

        return true
      }
    }
  ]
})

export default router
