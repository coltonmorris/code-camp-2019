import React from 'react';
import Grid from '@material-ui/core/Grid';
import Paper from '@material-ui/core/Paper';
import Button from '@material-ui/core/Button';

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
