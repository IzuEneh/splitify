<script setup lang="ts">
import { ref } from 'vue';
import { useRoute } from 'vue-router';
import mockPlaylistData from "@/views/mockPlaylist.json";

const clientId = "a69b0d8ab6b74d1598fd9e08c9741d7d"; // Replace with your client id
const route = useRoute()
const { code } = route.query
const playlists = ref(mockPlaylistData.items)

getAccessToken(clientId, code as string)
    .then(accessToken => {
        return getPlaylists("izueneh21", accessToken)
    })
    .then(playlists => console.log(playlists))


async function getAccessToken(clientId: string, code: string): Promise<string> {
    const verifier = localStorage.getItem("verifier");

    const params = new URLSearchParams();
    params.append("client_id", clientId);
    params.append("grant_type", "authorization_code");
    params.append("code", code);
    params.append("redirect_uri", "http://localhost:5173/callback");
    params.append("code_verifier", verifier!);

    const result = await fetch("https://accounts.spotify.com/api/token", {
        method: "POST",
        headers: { "Content-Type": "application/x-www-form-urlencoded" },
        body: params
    });

    const { access_token } = await result.json();
    return access_token;
}

async function getPlaylists(id: any, code: string): Promise<string> {

    const result = await fetch(`https://api.spotify.com/v1/users/${id}/playlists`, {
        method: "GET",
        headers: { "Content-Type": "application/json", "Authorization": `Bearer ${code}` },

    });

    const { playlists } = await result.json();
    console.log(playlists)
    return playlists;
}

</script>

<template>
    <div class="container">
        <header>
            <h1>Pick a playlist</h1>
        </header>

        <main>
            <ul class="playlist-list">
                <li v-for="playlist in playlists" :key="playlist.id">
                    <div class="playlist-item">
                        <img :src="playlist.images[0].url" />
                        <span>{{ playlist.name }}</span>
                    </div>
                </li>
            </ul>
        </main>
    </div>
</template>

<style scoped>
.container {
    display: grid;
    grid-template-columns: 1fr;
    grid-template-rows: 1fr auto;
    grid-template-areas:
        "header"
        "main";
    row-gap: 1rem;
    height: 100vh;
    padding: 0px 16px 32px;
}

header {
    grid-area: header;
}

main {
    grid-area: main
}

.playlist-list {
    display: flex;
    flex-direction: column;
    gap: 8px;
    list-style-type: none;
    padding: 0px;
}

.playlist-item {
    display: flex;
    flex-direction: row;
    align-items: center;
    gap: 16px;
    border: solid white;
    border-width: 1px 0px;
    padding: 8px 0px;
}

.playlist-item:hover {
    transform: scale(1.05);
    cursor: pointer;
    background-color: rgb(43, 43, 43);
}

.playlist-item img {
    width: 50px;
    height: auto;
}
</style>
