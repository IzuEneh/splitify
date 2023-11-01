<script setup lang="ts">
import { ref } from 'vue';
import { useRoute } from 'vue-router';
import mockPlaylistData from "@/views/mockPlaylist.json";

const clientId = "a69b0d8ab6b74d1598fd9e08c9741d7d"; // Replace with your client id
const route = useRoute()
const { code } = route.query
const userData = ref({
    id: "",
    email: "",
    uri: "",
    externalUri: "",
    link: "",
    imgUrl: "",
    displayName: ""
})
console.log(mockPlaylistData)
const mockProfile = {
    "display_name": "izueneh21",
    "external_urls": {
        "spotify": "https://open.spotify.com/user/izueneh21"
    },
    "href": "https://api.spotify.com/v1/users/izueneh21",
    "id": "izueneh21",
    "images": [],
    "type": "user",
    "uri": "spotify:user:izueneh21",
    "followers": {
        "href": null,
        "total": 7
    },
    "country": "CA",
    "product": "premium",
    "explicit_content": {
        "filter_enabled": false,
        "filter_locked": false
    },
    "email": "izueneh21@gmail.com"
}

getAccessToken(clientId, code as string)
    .then(accessToken => {
        getPlaylists("izueneh21", accessToken).then(playlists => console.log(playlists))
        return fetchProfile(accessToken)
    })
    .then(profile => {
        if (profile.error) {
            populateUI(mockProfile)
        } else {
            populateUI(profile)
        }
    })


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

async function fetchProfile(token: string): Promise<any> {
    const result = await fetch("https://api.spotify.com/v1/me", {
        method: "GET", headers: { Authorization: `Bearer ${token}` }
    });

    return await result.json();
}

function populateUI(profile: any) {
    userData.value.displayName = profile.display_name;
    // if (profile.images[0]) {
    //   const profileImage = new Image(200, 200);
    //   profileImage.src = profile.images[0].url;
    //   document.getElementById("avatar")!.appendChild(profileImage);
    // }

    userData.value.id = profile.id;
    userData.value.email = profile.email;
    userData.value.uri = profile.uri;
    userData.value.externalUri = profile.external_urls.spotify
    userData.value.link = profile.href
    userData.value.imgUrl = profile.images[0]?.url ?? '(no profile image)';
}
</script>

<template>
    <div class="container">
        <header>
            <h1>Display Spotify Profile</h1>
        </header>

        <main>
            <p>This is the spotify app</p>
            <section id="profile">
                <h2>Logged in as {{ userData.displayName }}</h2>
                <span id="avatar"></span>
                <ul>
                    <li>User ID: {{ userData.id }}</li>
                    <li>Email: {{ userData.email }}</li>
                    <li>Spotify URI: <a id="uri" :href="userData.externalUri">{{ userData.uri }}</a></li>
                    <li>Link: <a id="url" :href="userData.link">{{ userData.link }}</a></li>
                    <li>Profile Image: {{ userData.imgUrl }}</li>
                </ul>
            </section>
        </main>
    </div>
</template>

<style scoped>
.container {
    display: grid;
    grid-template-columns: 1fr;
    grid-template-rows: min-content auto;
    grid-template-areas:
        "header"
        "main";
    row-gap: 1rem;
    height: 100vh;
}

header {
    grid-area: header;
}

main {
    grid-area: main
}
</style>
