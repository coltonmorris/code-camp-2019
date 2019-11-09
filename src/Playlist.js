import React from 'react';
// import ListItem from '@material-ui/core/ListItem';
import Paper from '@material-ui/core/Paper';
import LinkAccountButton from './LinkAccountButton';
import './App.css';

const data = { // TODO get data
  profile: {
    name: "Brady",
    links: {
      spotify: true,
      apple: false,
      youtube: true,
    }
  }
}

function Playlist(props) {

  const displayLinkedAccounts = () => {
    return Object.keys(data.profile.links).map(key => {
      return <LinkAccountButton
        name={key}
        linked={data.profile.links[key]}
        key={key}
        big={false}
      />
    })
  }

  return (
    <Paper className="Playlist">
      <span>
        {`${props.name}`}
      </span>
      <span>
        {`${props.count} songs`}
      </span>
      <span className="SyncPlaylistContainer">
        {displayLinkedAccounts()}
      </span>
    </Paper>
  )
}

export default Playlist;
