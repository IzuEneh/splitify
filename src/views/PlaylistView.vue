<script setup lang="ts">
import { useRoute } from 'vue-router';
import { ref } from 'vue';
import mockPlaylistData from "@/views/mockPlaylist.json";
import format from 'date-fns/format'

const route = useRoute()
const { id } = route.params
const playlist = ref(mockPlaylistData.items[0])
const tracks = ref<any>({})
const newPlaylists = ref<Array<any>>([]);
const table_headers = [
    { title: "#", value: "id" },
    { title: "title", value: "track.name" },
    { title: "album", value: "track.album.name" },
    {
        title: "date added",
        key: "date",
        value: item => format(new Date(item.added_at), "MMM d, yyy")
    },
    {
        title: "length",
        key: "length",
        value: item => format(new Date(item.track.duration_ms), "m:ss")
    }
]

getPlaylist(id as string).then(
    playlist => getTracks(playlist.tracks.href)
)

async function getPlaylist(id: string) {
    const accessToken = localStorage.getItem("access_token")
    const result = await fetch(`https://api.spotify.com/v1/playlists/${id}`, {
        method: "GET",
        headers: { "Content-Type": "application/json", "Authorization": `Bearer ${accessToken}` },
    });

    const res = await result.json();
    playlist.value = res
    return res;
}

async function getTracks(trackURL: string) {
    const accessToken = localStorage.getItem("access_token")
    const result = await fetch(trackURL, {
        method: "GET",
        headers: { "Content-Type": "application/json", "Authorization": `Bearer ${accessToken}` },
    });

    const res = await result.json();
    tracks.value = res
    tracks.value.items = tracks.value.items.map((item, index) => ({ id: index + 1, ...item }))
    return res;
}

function splitTracks() {
    const hour_in_ms = 1000 * 60 * 60
    let time = 0
    let upodatedPlaylists: any[] = []
    let buffer: any[] = []
    tracks.value.items.forEach(track => {
        if (time < hour_in_ms) {
            time += track.track["duration_ms"]
            buffer.push(track)
        } else {
            upodatedPlaylists.push([...buffer])
            time = track.track["duration_ms"]
            buffer = [track]
        }
    });

    if (buffer.length != 0) {
        upodatedPlaylists.push([...buffer])
    }
    newPlaylists.value = upodatedPlaylists
}
</script>

<template>
    <div class="container">
        <header>
            <img :src="playlist.images[0].url" />
            <div>
                <h1>{{ playlist.name }}</h1>
                <span>{{ playlist.description }}</span>
            </div>
        </header>

        <main>
            <!-- section containing action buttons [split, save all?, share?] -->
            <section>
                <button @click="splitTracks">Split</button>
            </section>
            <section>
                <!-- initially contain original list after split display new playlists -->
                <v-data-table v-if="newPlaylists.length === 0" :headers="table_headers" :items="tracks.items">
                </v-data-table>

                <div v-else class="new-playlist-container">
                    <div v-for="(list, index) in newPlaylists" :key="index" class="new-playlist">
                        <p>New playlist {{ index + 1 }} </p>
                        <ul>
                            <li v-for="track in list" :key="track.track.id">
                                {{ track.track.name }}
                            </li>
                        </ul>
                    </div>
                </div>
            </section>
        </main>
    </div>
</template>

<style scoped>
.container {
    display: flex;
    flex-direction: column;
    width: 100vw;
    padding: 0px 16px;
}

header {
    display: flex;
    flex-direction: row;
    gap: 16px;
    align-items: center;
    justify-content: flex-start;
    padding: 16px 0px;
}

header img {
    width: 200px;
    height: auto;
}

header div {
    display: flex;
    flex-direction: column;
    /* align-items: flex-start; */
    /* justify-content: ; */
}

main {
    display: flex;
    flex-direction: column;
    gap: 16px;
}

ul {
    padding: 0px;
}

button {
    background-color: #1db954;
    color: black;
    border-radius: 500px;
    padding: 14px 32px;
    display: flex;
    -moz-box-align: center;
    align-items: center;
    -moz-box-pack: center;
    justify-content: center;
    font-family: spotify-circular, Helvetica, Arial, sans-serif;
    font-size: 18px;
    line-height: 24px;
    font-weight: 700;
    text-align: center;
    text-transform: none;
    border: none;
}

.new-playlist-container {
    display: flex;
    flex-direction: column;
    gap: 8px;
}
</style>