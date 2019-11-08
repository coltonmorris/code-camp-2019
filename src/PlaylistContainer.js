import React from 'react';
import Grid from '@material-ui/core/Grid';
import List from '@material-ui/core/List';
import Playlist from './Playlist';
import './App.css';

const data = { // TODO get data
  playlists: [{
    id: 1,
    name: 'Basic Playlist',
    count: 20,
  }, {
    id: 2,
    name: 'Baller Playlist',
    count: 35,
  }, {
    id: 3,
    name: 'Giga Playlist',
    count: 748,
  }]
}

function PlaylistContainer() {
  const displayPlaylists = () => {
    return data.playlists.map((list) => {
      return <Playlist
        key={list.id}
        {...list}
      />
    })
  }

  return (
    <Grid className="PlaylistContainer">
      <div className="PlaylistList">
        <div className="PlaylistTitle">Playlists</div>
        <List>
          {displayPlaylists()}
        </List>
      </div>
    </Grid>
  );
}

export default PlaylistContainer;
