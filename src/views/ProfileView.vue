<script setup lang="ts">
import { ref } from 'vue';
import { useRoute, useRouter } from 'vue-router';

const clientId = "a69b0d8ab6b74d1598fd9e08c9741d7d";
const route = useRoute()
const router = useRouter()
const { code } = route.query
const playlists = ref<Array<any>>([])
const cachedCode = localStorage.getItem("access_code")
const isLoading = ref(true)
const errorMessage = ref("")

if (code != cachedCode) {
    getAccessToken(clientId, code as string)
} else {
    isLoading.value = false
}

if (!isLoading.value && playlists.value.length == 0) {
    getPlaylists("izueneh21")
}

async function getAccessToken(clientId: string, code: string) {
    const verifier = localStorage.getItem("verifier");

    const params = new URLSearchParams();
    params.append("client_id", clientId);
    params.append("grant_type", "authorization_code");
    params.append("code", code);
    params.append("redirect_uri", "http://localhost:5173/callback");
    params.append("code_verifier", verifier!);

    try {
        const result = await fetch("https://accounts.spotify.com/api/token", {
            method: "POST",
            headers: { "Content-Type": "application/x-www-form-urlencoded" },
            body: params
        });

        const { access_token, refresh_token, expires_in } = await result.json();
        localStorage.setItem("access_token", access_token)
        localStorage.setItem("access_code", code)
        localStorage.setItem("refresh_token", refresh_token)
        localStorage.setItem("expires_in", expires_in)
    } catch (error) {
        errorMessage.value = `An error occcured getting access token: ${error}`
    } finally {
        isLoading.value = false
    }
}

async function getPlaylists(id: any) {
    const accessToken = localStorage.getItem("access_token")
    isLoading.value = true

    try {
        const result = await fetch(`https://api.spotify.com/v1/users/${id}/playlists`, {
            method: "GET",
            headers: { "Content-Type": "application/json", "Authorization": `Bearer ${accessToken}` },
        });

        const response = await result.json();
        playlists.value = response.items
    } catch (error) {
        errorMessage.value = `An error occcured getting access token: ${error}`
    } finally {
        isLoading.value = false
    }
}

</script>

<template>
    <div class="container">
        <header>
            <h1>Pick a playlist</h1>
        </header>

        <main>
            <div v-if="isLoading">Loading...</div>
            <p v-else-if="errorMessage.length > 0">{{ errorMessage }}</p>
            <ul v-else class="playlist-list">
                <li v-for="playlist in playlists" :key="playlist.id" @click="router.push(`/playlist/${playlist.id}`)">
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
