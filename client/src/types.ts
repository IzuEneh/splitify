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

export type PlaylistResponse = {
  collaborative: boolean
  description: string
  external_urls: {
    spotify: string
  }
  followers: {
    href: string
    total: number
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
    limit: number
    next: string
    offset: number
    previous: string
    total: number
    items: Track[]
  }
  type: string
  uri: string
}

export type Track = {
  added_at: string
  added_by: {
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
  }
  is_local: boolean
  track: {
    album: {
      album_type: 'compilation'
      total_tracks: 9
      available_markets: ['CA', 'BR', 'IT']
      external_urls: {
        spotify: 'string'
      }
      href: 'string'
      id: '2up3OPMp9Tb4dAKM2erWXQ'
      images: [
        {
          url: 'https://i.scdn.co/image/ab67616d00001e02ff9ca10b55ce82ae553c8228'
          height: 300
          width: 300
        }
      ]
      name: 'string'
      release_date: '1981-12'
      release_date_precision: 'year'
      restrictions: {
        reason: 'market'
      }
      type: 'album'
      uri: 'spotify:album:2up3OPMp9Tb4dAKM2erWXQ'
      artists: [
        {
          external_urls: {
            spotify: 'string'
          }
          href: 'string'
          id: 'string'
          name: 'string'
          type: 'artist'
          uri: 'string'
        }
      ]
    }
    artists: [
      {
        external_urls: {
          spotify: 'string'
        }
        followers: {
          href: 'string'
          total: 0
        }
        genres: ['Prog rock', 'Grunge']
        href: 'string'
        id: 'string'
        images: [
          {
            url: 'https://i.scdn.co/image/ab67616d00001e02ff9ca10b55ce82ae553c8228'
            height: 300
            width: 300
          }
        ]
        name: 'string'
        popularity: 0
        type: 'artist'
        uri: 'string'
      }
    ]
    available_markets: ['string']
    disc_number: 0
    duration_ms: 0
    explicit: false
    external_ids: {
      isrc: 'string'
      ean: 'string'
      upc: 'string'
    }
    external_urls: {
      spotify: 'string'
    }
    href: 'string'
    id: 'string'
    is_playable: false
    linked_from: {}
    restrictions: {
      reason: 'string'
    }
    name: 'string'
    popularity: 0
    preview_url: 'string'
    track_number: 0
    type: 'track'
    uri: 'string'
    is_local: false
  }
}

export type GeneratedPlaylist = {
  id: number
  images: Array<{
    url: string
    height: number
    width: number
  }>
  name: string
  description: string
  owner: {
    display_name: string
  }
  tracks: {
    total: number
    items: Track[]
  }
}
