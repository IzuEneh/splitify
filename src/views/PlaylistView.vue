<script setup lang="ts">
import { useRoute } from 'vue-router';
import { ref } from 'vue';
import mockPlaylistData from "@/views/mockPlaylist.json";

const route = useRoute()
const { id } = route.params
const playlist = ref(mockPlaylistData.items[0])
const tracks = ref<any>({})
const newPlaylists = ref<Array<any>>([])

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
    console.log(upodatedPlaylists)
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
            <section>
                <ul>
                    <li v-for="track in tracks.items" :key="track.track.id">
                        {{ track.track.name }}
                    </li>
                </ul>
            </section>
            <section>
                <button @click="splitTracks">Split</button>
            </section>
            <section class="new-playlist-container">
                <template v-if="newPlaylists.length > 0">
                    <div v-for="(list, index) in newPlaylists" :key="index" class="new-playlist">
                        <p>New playlist {{ index }} </p>
                        <ul>
                            <li v-for="track in list" :key="track.track.id">
                                {{ track.track.name }}
                            </li>
                        </ul>
                    </div>
                </template>
            </section>
        </main>
    </div>
</template>

<style scoped>
.container {
    display: grid;
    grid-template-columns: 1fr;
    grid-template-rows: 1fr 3fr;
    width: 100vw;
    padding: 0px 16px;
}

header {
    display: grid;
    grid-template-columns: 1fr 3fr;
    grid-template-rows: 1fr;
    align-items: center;
    justify-content: center;
}

header img {
    width: 100px;
    height: auto;
}

header div {
    display: flex;
    flex-direction: column;
    align-items: center;
}

main {
    display: grid;
    grid-template-columns: 1fr 1fr 1fr;
    grid-template-rows: 1fr;
}

.new-playlist-container {
    display: flex;
    flex-direction: column;
    gap: 8px;
}
</style>