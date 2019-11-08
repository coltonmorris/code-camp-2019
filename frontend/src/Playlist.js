import React from 'react';
import Grid from '@material-ui/core/Grid';
import Button from '@material-ui/core/Button';
import './App.css';

function Playlist(props) {
  return (
    <Grid className="Playlist">
      <span className="PlaylistPart">
        {`${props.name}`}
      </span>
      <span className="PlaylistPart">
        {`${props.count} songs`}
      </span>
      <span className="PlaylistLinks">
        <Button variant="contained">Spotify</Button>
        <Button variant="contained">Apple</Button>
        <Button variant="contained">Google</Button>
      </span>
    </Grid>
  )
}

export default Playlist;
