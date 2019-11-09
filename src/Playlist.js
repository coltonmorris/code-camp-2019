import React from 'react';
import LinkAccountButton from './LinkAccountButton';
import { sync } from './api';
import './App.css';

const data = { // TODO get data
  profile: {
    links: {
      spotify: true,
      apple: false,
      youtube: true,
    }
  }
}

const tabs = ['spotify', 'apple', 'youtube'];
function Playlist(props) {
  // name  playlist from  to
  const doSync = async (destination) => {
    console.log(props.user, props.name, tabs[props.tab], destination);
    try {
      const res = await sync(props.user, props.name, tabs[props.tab], destination);
      props.openDrawer()
      props.addJob({
        name: props.name,
        from: tabs[props.tab],
        to: destination,
        recordsSuccess: 0,
        recordsFailed: 0,
        totalRecords: 100,
      })
    } catch (err) {
      console.error('Bad Sync.', err);
    }
  }


  const displayLinkedAccounts = () => {
    return Object.keys(data.profile.links).map(key => {
      return <LinkAccountButton
        name={key}
        linked={data.profile.links[key]}
        key={key}
        big={false}
        onClick={() => doSync(key)}
      />
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
