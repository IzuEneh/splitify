export type PlaylistsResponse = {
  href: string
  limit: number
  next: string
  offset: number
  previous: string
  total: number
  items: Playlist[]
}

export type Playlist = {
  collaborative: boolean
  description: string
  external_urls: {
    spotify: string
  }
  href: string
  id: string
  images: Array<{
    url: string
    height: number
    width: number
  }>
  name: string
  owner: {
    external_urls: {
      spotify: string
    }
    followers: {
      href: string
      total: number
    }
    href: string
    id: string
    type: string
    uri: string
    display_name: string
  }
  public: false
  snapshot_id: string
  tracks: {
    href: string
    total: number
  }
  type: string
  uri: string
}
