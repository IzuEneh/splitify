<script setup lang="ts">
import { ref, watch } from 'vue';
import { useRouter } from 'vue-router';
import PlaylistList from '@/components/PlaylistList.vue';
import PlaylistView from '@/components/PlaylistView.vue';
import type { PlaylistResponse, Track } from '@/types';
import usePlaylists from '@/composables/usePlaylists';
import fetchAccessToken from '@/api/fetch';

const selectedID = ref("")
const selectedPlaylist = ref<PlaylistResponse | null>(null)
const newPlaylists = ref<any[]>([])
const router = useRouter()
const activeWindow = ref(0)
const { playlists, loading, error } = usePlaylists()

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

function createNewPlaylist(id: string, name: string, tracks: Track[]) {
    return ({
        images: [{ url: "https://source.unsplash.com/random/60×60/?music", height: 60, with: 60 }],
        name: name,
        owner: {
            display_name: "current user"
        },
        tracks: {
            items: tracks
        }
    })
}

function handleSplitPlaylist() {
    const hour_in_ms = 1000 * 60 * 60
    let time = 0
    let upodatedPlaylists: any[] = []
    let buffer: Track[] = []
    selectedPlaylist.value?.tracks.items.forEach(track => {
        if (time < hour_in_ms) {
            time += track.track["duration_ms"]
            buffer.push(track)
        } else {
            upodatedPlaylists.push(createNewPlaylist(`${upodatedPlaylists.length}`, `Playlist ${upodatedPlaylists.length + 1}`, [...buffer]))
            time = track.track["duration_ms"]
            buffer = [track]
        }
    });

    if (buffer.length != 0) {
        upodatedPlaylists.push(createNewPlaylist(`${upodatedPlaylists.length}`, `Playlist ${upodatedPlaylists.length + 1}`, [...buffer]))
    }

    newPlaylists.value = upodatedPlaylists
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
            <PlaylistView :playlist="selectedPlaylist" :loading="loading" @on-split-playlist="handleSplitPlaylist" />
        </section>

        <section class="content-area" v-if="newPlaylists.length > 0" :class="{ 'activeWindow': activeWindow > 1 }">
            <div class="button-container">
                <button @click="handleBackPress" class="back-button">&larr;</button>
            </div>
            <span class="playlist-title">Generated Playlists</span>
            <div class="scrollable">
                <PlaylistList :playlists="newPlaylists" selected="0" />
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
