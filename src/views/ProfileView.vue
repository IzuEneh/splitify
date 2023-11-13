<script setup lang="ts">
import PlaylistList from '@/components/PlaylistList.vue';
import { ref } from 'vue';
import { useRoute, useRouter } from 'vue-router';

const route = useRoute()
const router = useRouter()
const { code } = route.query
const playlists = ref<Array<any>>([])
const cachedCode = localStorage.getItem("access_code")
const isLoading = ref(true)
const errorMessage = ref("")

if (code != cachedCode) {
    getAccessToken(code as string)
} else {
    isLoading.value = false
}

if (!isLoading.value && playlists.value.length == 0) {
    getPlaylists("izueneh21")
}

async function getAccessToken(code: string) {
    const verifier = localStorage.getItem("verifier");

    const params = new URLSearchParams();
    params.append("client_id", import.meta.env.VITE_CLIENT_ID);
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
        <section class="content-area">
            <span>Your Playlists</span>
            <div>
                <PlaylistList :playlists="playlists" />
            </div>
        </section>
    </div>
</template>

<style scoped>
.container {
    display: flex;
    gap: 1rem;
    padding: 16px;
    height: 100vh;
    width: 100vw;
}

.content-area {
    background-color: var(--color-background-soft);
    border-radius: 10px;
    width: 25vw;
    display: flex;
    flex-direction: column;
    padding: 16px;
}

.content-area span {
    font-size: large;
    font-weight: 500;
}

.content-area div {
    overflow-y: scroll;
}
</style>
