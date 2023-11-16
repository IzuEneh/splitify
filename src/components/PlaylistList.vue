<script setup lang="ts">
import type { Playlist } from '@/types';

const { playlists, selected } = defineProps<{
    playlists: Playlist[],
    selected: string
}>()

const getURL = (imageArr: Array<{
    url: string
    height: number
    width: number
}>) => {
    if (imageArr.length < 1) {
        return 'https://source.unsplash.com/random/60×60/?music'
    }

    return imageArr[0].url
}
</script>

<template>
    <ul class="playlist-list">
        <li v-for="playlist in playlists" :key="playlist.id" @click="$emit('onSelectPlaylist', playlist.id)">
            <div class="playlist-item" :class="{ 'selected': playlist.id === selected }">
                <img :src="getURL(playlist.images)" />
                <div>
                    <span>{{ playlist.name }}</span>
                    <span>{{ playlist.owner.display_name }}</span>
                </div>
            </div>
        </li>
    </ul>
</template>

<style scoped>
.playlist-list {
    display: flex;
    flex-direction: column;
    gap: 16px;
    list-style-type: none;
    padding: 0px;
    width: 100%;
    height: 100%;
}

.playlist-item {
    display: flex;
    flex-direction: row;
    align-items: center;
    gap: 16px;
    padding: 8px;
    border-radius: 5px;
}

.playlist-item:hover {
    cursor: pointer;
    background-color: rgb(43, 43, 43);
}

.playlist-item img {
    width: 60px;
    height: auto;
    border-radius: 5px;
}

.playlist-item div {
    display: flex;
    flex-direction: column;
}

.selected {
    background-color: rgb(43, 43, 43);
}
</style>
