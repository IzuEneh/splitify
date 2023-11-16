import { createRouter, createWebHistory } from 'vue-router'
import HomeView from '@/views/HomeView.vue'
import Profile from '@/views/ProfileView.vue'
import UnauthorizedView from '@/views/UnauthorizedView.vue'

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
          return { path: '/playlists' }
        }

        return true
      }
    },
    {
      path: '/playlists',
      name: 'profile',
      component: Profile,
      beforeEnter: async (to) => {
        const accessToken = localStorage.getItem('access_token')
        if (to.query && !accessToken) {
          const code = to.query.code as string
          const access_token = await getAccessToken(code)

          return access_token || { path: '/unauthorized' }
        }

        return true
      }
    },
    {
      path: '/unauthorized',
      name: 'unauthorized',
      component: UnauthorizedView
    }
  ]
})

async function getAccessToken(code: string) {
  const verifier = localStorage.getItem('verifier')

  const params = new URLSearchParams()
  params.append('client_id', import.meta.env.VITE_CLIENT_ID)
  params.append('grant_type', 'authorization_code')
  params.append('code', code)
  params.append('redirect_uri', import.meta.env.VITE_REDIRECT_URI)
  params.append('code_verifier', verifier!)

  try {
    const result = await fetch('https://accounts.spotify.com/api/token', {
      method: 'POST',
      headers: { 'Content-Type': 'application/x-www-form-urlencoded' },
      body: params
    })

    if (!result.ok) {
      return false
    }
    const { access_token, refresh_token, expires_in } = await result.json()
    localStorage.setItem('access_token', access_token)
    localStorage.setItem('refresh_token', refresh_token)
    localStorage.setItem('expires_in', expires_in)
    localStorage.setItem('access_code_fetched_date', Date.now().toString())
    return true
  } catch (error) {
    console.log(`An error occcured getting access token: ${error}`)
    return false
  }
}

export default router
