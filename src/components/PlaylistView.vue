<script setup lang="ts">
import type { PlaylistResponse, Track } from '@/types';
import format from 'date-fns/format'

const { playlist } = defineProps<{ playlist: PlaylistResponse | null }>()
const table_headers = [
    { title: "#", value: "id" },
    { title: "Title", value: "track.name" },
    { title: "Album", value: "track.album.name" },
    {
        title: "Date added",
        key: "date",
        value: (item: Track) => format(new Date(item.added_at), "MMM d, yyy")
    },
    {
        title: "Length",
        key: "length",
        value: (item: Track) => format(new Date(item.track.duration_ms), "m:ss")
    }
]

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
    // newPlaylists.value = upodatedPlaylists
}
</script>

<template>
    <div class="playlist-container">
        <div class="header">
            <img :src="playlist?.images[0].url" />
            <div>
                <h1>{{ playlist?.name }}</h1>
                <span>{{ playlist?.description }}</span>
            </div>
        </div>

        <div class="main scroll-section">
            <!-- section containing action buttons [split, save all?, share?] -->
            <section>
                <button @click="splitTracks">Split</button>
            </section>
            <div class="scroll-section">
                <!-- initially contain original list after split display new playlists -->
                <v-data-table-virtual :headers="table_headers" :items="playlist?.tracks.items" class="table">
                    <template v-slot:headers="{ columns }">
                        <tr>
                            <template v-for="column in columns" :key="column.key">
                                <td>
                                    <span>{{ column.title }}</span>
                                </td>
                            </template>
                        </tr>
                    </template>
                </v-data-table-virtual>
            </div>
        </div>
    </div>
</template>

<style scoped>
.playlist-container {
    display: flex;
    flex-direction: column;
    width: 100%;
    padding: 0px 16px;
    overflow: hidden;
}

.header {
    display: flex;
    flex-direction: row;
    gap: 16px;
    align-items: center;
    justify-content: flex-start;
    padding: 16px 0px;
}

.header img {
    width: 200px;
    height: auto;
}

.header div {
    display: flex;
    flex-direction: column;
    /* align-items: flex-start; */
    /* justify-content: ; */
}

.main {
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

.scroll-section {
    overflow-y: scroll;
}

.table {
    color: white;
    background-color: transparent;
}

.table th span {
    color: white
}
</style>