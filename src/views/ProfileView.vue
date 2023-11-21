<script setup lang="ts">
import { ref, watch } from 'vue';
import { useRouter } from 'vue-router';
import PlaylistList from '@/components/PlaylistList.vue';
import PlaylistView from '@/components/PlaylistView.vue';
import type { PlaylistResponse, GeneratedPlaylist, Track } from '@/types';
import usePlaylists from '@/composables/usePlaylists';
import { fetchAccessToken, fetchUser } from '@/api/fetch';

const router = useRouter()
const { playlists, loading, error } = usePlaylists()
const selectedID = ref("")
const selectedPlaylist = ref<PlaylistResponse | GeneratedPlaylist | null>(null)
const newPlaylists = ref<GeneratedPlaylist[]>([])
const activeWindow = ref(0)
const showModal = ref(false)
const modalMessage = ref("")

watch(playlists, () => {
    if (playlists.value.length != 0) {
        getPlaylist(playlists.value[0].id)
    }
})
watch(error, () => {
    if (!loading && error.value.length > 0) {
        console.log(error)
        router.push({ name: 'unauthorized' })
    }
})

async function getPlaylist(id: string) {
    try {
        const accessToken = await fetchAccessToken()
        const result = await fetch(`https://api.spotify.com/v1/playlists/${id}`, {
            method: "GET",
            headers: { "Content-Type": "application/json", "Authorization": `Bearer ${accessToken}` },
        });

        if (!result.ok) {
            await router.push({ name: 'unauthorized' })
            return;
        }

        const res = await result.json();
        selectedPlaylist.value = res
        selectedPlaylist.value!.tracks.items = selectedPlaylist.value!.tracks.items.map((item, index) => ({ id: index + 1, ...item }))

    } catch (err) {
        console.log("Error fetching playlist: ", err)
    }
}

function createNewPlaylist(id: number, name: string, tracks: Track[]) {
    return ({
        id,
        images: [{ url: "https://source.unsplash.com/random/60×60/?music", height: 60, with: 60 }],
        name: name,
        description: "new playlist created by ai",
        owner: {
            display_name: "current user"
        },
        tracks: {
            total: tracks.length,
            items: tracks.map((item, index) => ({ ...item, id: index + 1 }))
        }
    })
}

function handleSplitPlaylist() {
    const hour_in_ms = 1000 * 60 * 60
    let time = 0
    let updatedPlaylists: any[] = []
    let buffer: Track[] = []
    selectedPlaylist.value?.tracks.items.forEach((track) => {
        if (time < hour_in_ms) {
            time += track.track["duration_ms"]
            buffer.push(track)
        } else {
            updatedPlaylists.push(createNewPlaylist(updatedPlaylists.length, `Playlist ${updatedPlaylists.length + 1}`, [...buffer]))
            time = track.track["duration_ms"]
            buffer = [track]
        }
    });

    if (buffer.length != 0) {
        updatedPlaylists.push(createNewPlaylist(updatedPlaylists.length, `Playlist ${updatedPlaylists.length + 1}`, [...buffer]))
    }

    newPlaylists.value = updatedPlaylists
    activeWindow.value++
}

function handleSelectPlaylist(id: string) {
    activeWindow.value++
    selectedID.value = id
    getPlaylist(id)
}

function handleBackPress() {
    activeWindow.value--
}

function handleNewPlaylist(id: number) {
    selectedPlaylist.value = newPlaylists.value[id]
}

async function handleSavePlaylist() {
    if (!selectedPlaylist.value) {
        return;
    }

    showModal.value = true
    modalMessage.value = "One moment while we save your playlist"
    try {
        const [user, accessToken] = await Promise.all([fetchUser(), fetchAccessToken()])
        const resp = await fetch(`https://api.spotify.com/v1/users/${user.id}/playlists`, {
            method: "POST",
            headers: { "Content-Type": "application/json", "Authorization": `Bearer ${accessToken}` },
            body: JSON.stringify({
                "name": selectedPlaylist.value.name,
                "description": selectedPlaylist.value.description,
                "public": "false"
            })
        })

        if (!resp.ok) {
            console.log("unable to create playlist")
            showModal.value = true
            modalMessage.value = "An Error occured saving your playlist"
            return;
        }

        const { id } = await resp.json()
        const uris = selectedPlaylist.value.tracks.items.map(track => track.track.uri)
        const resp2 = await fetch(`https://api.spotify.com/v1/playlists/${id}/tracks`, {
            method: "POST",
            headers: { "Content-Type": "application/json", "Authorization": `Bearer ${accessToken}` },
            body: JSON.stringify({
                "uris": uris
            })
        })

        if (!resp2.ok) {
            console.log("unable to add tracks to playlist with id: ", id)
            showModal.value = true
            modalMessage.value = "An Error occured saving your playlist"
            return
        }

        console.log("tracks added successfully")
        showModal.value = true
        modalMessage.value = "Playlist saved successfully!"
    } catch (err) {
        console.log(err)
    }
}
</script>

<template>
    <div class="container">
        <section class="content-area" :class="{ 'activeWindow': activeWindow > -1 }">
            <span class="playlist-title">Your Playlists</span>
            <div class="scrollable">
                <PlaylistList :playlists="playlists" :selected="selectedID" :loading="loading"
                    @on-select-playlist="handleSelectPlaylist" />
            </div>
        </section>

        <section class="content-area grow" :class="{ 'activeWindow': activeWindow > 0 }">
            <div class="button-container">
                <button @click="handleBackPress" class="back-button">&larr;</button>
            </div>
            <PlaylistView :playlist="selectedPlaylist" :loading="loading" @on-split-playlist="handleSplitPlaylist"
                @on-save-playlist="handleSavePlaylist" />
            <v-snackbar v-model="showModal" timeout="2000">{{ modalMessage }}</v-snackbar>
        </section>

        <section class="content-area" v-if="newPlaylists.length > 0" :class="{ 'activeWindow': activeWindow > 1 }">
            <div class="button-container">
                <button @click="handleBackPress" class="back-button">&larr;</button>
            </div>
            <span class="playlist-title">Generated Playlists</span>
            <div class="scrollable">
                <PlaylistList :playlists="newPlaylists" selected="0" @on-select-playlist="handleNewPlaylist" />
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
    width: 20vw;
    display: flex;
    flex-direction: column;
    padding: 16px;
}

.grow {
    flex: 1
}

.playlist-title {
    font-size: large;
    font-weight: 500;
}

.scrollable {
    overflow-y: scroll;
}

.button-container {
    display: none;
}

@media (max-width:960px) {
    .content-area {
        height: 100%;
        /* 100% Full-height */
        width: 0;
        /* 0 width - change this with JavaScript */
        position: fixed;
        /* Stay in place */
        z-index: 1;
        /* Stay on top */
        top: 0;
        right: 0;
        background-color: var(--color-background-soft);
        /* Black*/
        overflow-x: hidden;
        /* Disable horizontal scroll */
        /* padding-top: 60px; */
        padding: 0px;
        /* Place content 60px from the top */
        transition: 0.5s;
        /* 0.5 second transition effect to slide in the sidebar */
        border-radius: unset;
    }

    .activeWindow {
        width: 100vw;
        padding: 16px;
    }

    .button-container {
        width: 100%;
        font-size: 1.5rem;
        display: block;
    }

}
</style>
