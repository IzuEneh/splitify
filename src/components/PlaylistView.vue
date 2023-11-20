<script setup lang="ts">
import type { GeneratedPlaylist, PlaylistResponse, Track } from '@/types';
import format from 'date-fns/format'

const { playlist } = defineProps<{
    playlist: PlaylistResponse | GeneratedPlaylist | null,
    loading?: boolean
}>()
const isDesktop = window.matchMedia('(min-width:961px)');
let table_headers: any[];
if (isDesktop.matches) {
    table_headers = [
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
} else {
    table_headers = [
        { title: "Title", value: "track.name" },
    ]
}
</script>

<template>
    <div class="playlist-container">
        <div class="header">
            <template v-if="loading">
                <v-skeleton-loader type="image" color="grey-darken-3" class="loading-avatar"></v-skeleton-loader>
                <v-skeleton-loader type="text" color="grey-darken-3" class="w-25"></v-skeleton-loader>
            </template>
            <template v-else>
                <img :src="playlist?.images[0].url" />
                <div>
                    <h1>{{ playlist?.name }}</h1>
                    <span><span v-html="playlist?.description"></span></span>
                </div>
            </template>
        </div>

        <div class="main">
            <!-- section containing action buttons [split, save all?, share?] -->
            <section>
                <button @click="$emit('onSplitPlaylist')">Split</button>
            </section>
            <div>
                <template v-if="loading">
                    <v-skeleton-loader type="table" color="grey-darken-3" class=""></v-skeleton-loader>
                </template>
                <template v-else>
                    <v-data-table-virtual :headers="table_headers" :items="playlist?.tracks.items" class="table">
                        <template v-slot:headers="{ columns }">
                            <tr class="mobile-header">
                                <template v-for="column in columns" :key="column.key">
                                    <td>
                                        <span>{{ column.title }}</span>
                                    </td>
                                </template>
                            </tr>
                        </template>
                    </v-data-table-virtual>
                </template>
            </div>
        </div>
    </div>
</template>

<style scoped>
.loading-avatar {
    height: 200px;
    width: 200px;
}

.playlist-container {
    display: flex;
    flex-direction: column;
    width: 100%;
    overflow-y: scroll;
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
}

.main {
    display: flex;
    flex-direction: column;
    gap: 16px;
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

.table {
    color: white;
    background-color: transparent;
}

.table th span {
    color: white
}

@media (max-width:960px) {
    .header {
        flex-direction: column;
        text-align: center;
    }

    .main {
        align-items: center;
    }

    .mobile-header {
        display: none;
    }
}
</style>