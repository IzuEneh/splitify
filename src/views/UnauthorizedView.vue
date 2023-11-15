<script setup lang="ts">
import { useRouter } from 'vue-router';
const router = useRouter()
</script>

<template>
    <div class="container">
        <h1>Unable to Authenticate with Spotify</h1>
        <p>please try again.</p>
        <button @click="router.push({ name: 'home' })">Try again</button>
    </div>
</template>

<style scoped>
.container {
    display: flex;
    flex-direction: column;
    justify-content: center;
    align-items: center;
    width: 100vw;
    height: 100vh;
    gap: 16px;
}

h1,
p {
    padding: 0px;
    margin: 0px;
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
</style>