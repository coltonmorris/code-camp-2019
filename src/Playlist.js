import React from 'react';
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
      if (key !== props.currentTab) {
        return <LinkAccountButton
          name={key}
          linked={data.profile.links[key]}
          key={key}
          big={false}
          onClick={props.openDrawer}
        />
      } else { return null; }
    })
  }

  return (
    <div className="Playlist">
      <span className="rest">
        {`${props.name}`}
      </span>
      <span className="SongCountRight">
          {`${props.count}`}
      </span>
      <span className="SongCount">
          songs
      </span>
      <span className="SyncPlaylistContainer">
        {displayLinkedAccounts()}
      </span>
    </div>
  )
}

export default Playlist;
