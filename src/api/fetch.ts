async function fetchAccessToken() {
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

async function fetchUser() {
  const user = localStorage.getItem('user')
  if (user) {
    return JSON.parse(user)
  }

  const accessToken = await fetchAccessToken()
  const resp = await fetch('https://api.spotify.com/v1/me', {
    method: 'GET',
    headers: { 'Content-Type': 'application/json', Authorization: `Bearer ${accessToken}` }
  })

  if (!resp.ok) {
    throw new Error(`Error fetching user: ${resp.status} ${resp.statusText}`)
  }

  const userResp = await resp.json()
  localStorage.setItem('user', JSON.stringify(userResp))
  return userResp
}

export { fetchUser, fetchAccessToken }
