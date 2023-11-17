import { ref } from 'vue'
import type { Playlist } from '@/types'
import fetchAccessToken from '@/api/fetch'

export default function usePlaylists() {
  const playlists = ref<Playlist[]>([])
  const loading = ref(true)
  const error = ref('')

  async function getPlaylists() {
    let url = `https://api.spotify.com/v1/me/playlists`
    let res: any[] = []
    try {
      const accessToken = await fetchAccessToken()
      while (url != null) {
        const result = await fetch(url, {
          method: 'GET',
          headers: { 'Content-Type': 'application/json', Authorization: `Bearer ${accessToken}` }
        })

        if (!result.ok) {
          error.value = 'error fetching playlists, url: ' + url
          return
        }

        const response = await result.json()
        res = res.concat(response.items)
        url = response.next
      }

      playlists.value = res
    } catch (err) {
      error.value = `An error occcured getting access token: ${err}`
    } finally {
      loading.value = false
    }
  }

  getPlaylists()
  return { playlists, loading, error }
}
