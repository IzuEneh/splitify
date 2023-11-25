# Project Setup

1. Go to https://developer.spotify.com/documentation/web-api and create a new app.
2. Copy the Client ID into the `.env` file. Assign it to the variable `SPOTIFY_CLIENT_ID`.
3. Confirm ID has been set by running 'make env' from the root directory. Your id should show up in the console output like so CLIENT_ID=**your id here**
3. In the root of the project run "make" to build the docker image and run the container.
4. See Makefile for other commands. Comment are provided for each command with a brief description of what it does.
5. If you have any questions or issues please contact me at me@victorsmith.dev.