[![Build Status](https://travis-ci.com/coltonmorris/code-camp-2019.svg?branch=master)](https://travis-ci.com/coltonmorris/code-camp-2019)
# PAWG CHAMPS

'{\"web\":{\"client_id\":\"612720349908-4qpltha782sgpooira264m51smkpqmr1.apps.googleusercontent.com\",\"project_id\":\"synclist-blake\",\"auth_uri\":\"https://accounts.google.com/o/oauth2/auth\",\"token_uri\":\"https://oauth2.googleapis.com/token\",\"auth_provider_x509_cert_url\":\"https://www.googleapis.com/oauth2/v1/certs\",\"client_secret\":\"EezHkcB0gFrycMJe0Vl7tn_4\",\"redirect_uris\":[\"http://synclist.tech/callback/youtube\",\"http://localhost:8080/callback/youtube\"],\"javascript_origins\":[\"http://synclist.tech\",\"http://localhost:8080\"]}}'


## Youtube Setup
1. Get the `api_key` and the `client_secret.json` file from `https://console.developers.google.com/apis/api/youtube.googleapis.com/credentials?project=coce-camp-2019`
2. Give heroku the env var (note: The `YOUTUBE_CLIENT_SECRET_JSON` env var will just be stringified json`):
```
  heroku config:set YOUTUBE_API_KEY=api_key_here -a code-camp-2019 && \
  heroku config:set YOUTUBE_CLIENT_SECRET_JSON=client_secret_json_file_contents_here -a code-camp-2019 && \
  heroku config:set YOUTUBE_REDIRECT="http://synclist.tech/callback/youtube" -a code-camp-2019
```
3. Ensure Redirect URIs are inputed in the dev console


## Spotify Setup
1. Get the `client_id` and `client_secret` from `https://developer.spotify.com/dashboard/`.
2. Give heroku the env vars:
```
  heroku config:set SPOTIFY_ID=client_id_here -a code-camp-2019 && \
  heroku config:set SPOTIFY_SECRET=client_secret_here -a code-camp-2019 && \
  heroku config:set SPOTIFY_REDIRECT="http://synclist.tech/callback/spotify" -a code-camp-2019
```
3. Ensure Redirect URIs are inputed in the dev console
