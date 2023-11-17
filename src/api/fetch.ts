export default async function fetchAccessToken() {
  const accessToken = localStorage.getItem('access_token')
  const expiresIn = localStorage.getItem('expires_in')
  const lastFetchedAccessToken = localStorage.getItem('access_code_fetched_date')
  const currentTimeSecs = Date.now() / 1000
  const elapsedTime = parseInt(lastFetchedAccessToken!) / 1000 + parseInt(expiresIn!)
  if (elapsedTime > currentTimeSecs) {
    return accessToken
  }

  const refreshToken = localStorage.getItem('refresh_token')
  const params = new URLSearchParams()
  params.append('client_id', import.meta.env.VITE_CLIENT_ID)
  params.append('grant_type', 'refresh_token')
  params.append('refresh_token', refreshToken!)

  const body = await fetch('https://accounts.spotify.com/api/token', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/x-www-form-urlencoded'
    },
    body: params
  })

  if (!body.ok) {
    throw new Error(`Error fetching refresh token: ${body.status} ${body.statusText}`)
  }

  const { access_token, expires_in, refresh_token } = await body.json()
  localStorage.setItem('access_token', access_token)
  localStorage.setItem('refresh_token', refresh_token)
  localStorage.setItem('expires_in', expires_in)
  localStorage.setItem('access_code_fetched_date', Date.now().toString())
  return access_token
}
