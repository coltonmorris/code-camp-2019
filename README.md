[![Build Status](https://travis-ci.com/coltonmorris/code-camp-2019.svg?branch=master)](https://travis-ci.com/coltonmorris/code-camp-2019)
# PAWG CHAMPS



## Youtube Setup
1. Get the `api_key` and the `client_secret.json` file from `https://console.developers.google.com/apis/api/youtube.googleapis.com/credentials?project=coce-camp-2019`
2. Give heroku the env var (note: The `YOUTUBE_CLIENT_SECRET_JSON` env var will just be stringified json`):
```
  heroku config:set YOUTUBE_API_KEY=api_key_here -a code-camp-2019 && \
  heroku config:set YOUTUBE_CLIENT_SECRET_JSON=client_secret_json_file_contents_here -a code-camp-2019
```
3. Ensure Redirect URIs are inputed in the dev console


## Spotify Setup
1. Get the `client_id` and `client_secret` from `https://developer.spotify.com/dashboard/`.
2. Give heroku the env vars:
```
  heroku config:set SPOTIFY_ID=client_id_here -a code-camp-2019 && \
  heroku config:set SPOTIFY_SECRET=client_secret_here -a code-camp-2019
```
3. Ensure Redirect URIs are inputed in the dev console
