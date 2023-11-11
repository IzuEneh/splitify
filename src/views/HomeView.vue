<script setup lang="ts">

const clientId = "a69b0d8ab6b74d1598fd9e08c9741d7d";

async function redirectToAuthCodeFlow() {
    const verifier = generateCodeVerifier(128);
    const challenge = await generateCodeChallenge(verifier);

    localStorage.setItem("verifier", verifier);

    const params = new URLSearchParams();
    params.append("client_id", clientId);
    params.append("response_type", "code");
    params.append("redirect_uri", "http://localhost:5173/callback");
    params.append("scope", "user-read-private user-read-email playlist-read-private playlist-read-collaborative");
    params.append("code_challenge_method", "S256");
    params.append("code_challenge", challenge);

    document.location = `https://accounts.spotify.com/authorize?${params.toString()}`;
}

function generateCodeVerifier(length: number) {
    let text = '';
    let possible = 'ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789';

    for (let i = 0; i < length; i++) {
        text += possible.charAt(Math.floor(Math.random() * possible.length));
    }
    return text;
}

async function generateCodeChallenge(codeVerifier: string) {
    const data = new TextEncoder().encode(codeVerifier);
    const digest = await window.crypto.subtle.digest('SHA-256', data);
    return btoa(String.fromCharCode.apply(null, [...new Uint8Array(digest)]))
        .replace(/\+/g, '-')
        .replace(/\//g, '_')
        .replace(/=+$/, '');
}

</script>

<template>
    <div class="container">
        <header>
            <h1>Playlist Splitter</h1>
            <span>split your spotify playlists!</span>
        </header>

        <main>
            <button @click="redirectToAuthCodeFlow">Authenticate</button>
        </main>
    </div>
</template>

<style scoped>
.container {
    display: grid;
    grid-template-columns: 1fr;
    grid-template-rows: 1fr 2fr;
    grid-template-areas:
        "header"
        "main";
    height: 100vh;
    width: 100vw;
    justify-items: center;
}

header {
    grid-area: header;
    display: flex;
    flex-direction: column;
    justify-content: center;
    align-items: center;
}

h1 {
    font-size: 3rem;
}

span {
    font-size: 1.5rem;
}

main {
    grid-area: main
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

button:hover {
    transform: scale(1.05);
}

@media (min-width:961px) {
    .container {
        grid-template-columns: 1fr 1fr;
        grid-template-rows: 1fr;
        grid-template-areas:
            "header main";
        align-items: center;
    }
}
</style>
