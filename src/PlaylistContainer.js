import React from 'react';
import Playlist from './Playlist';
import Paper from '@material-ui/core/Paper';
import List from '@material-ui/core/List';
import Tabs from '@material-ui/core/Tabs';
import Tab from '@material-ui/core/Tab';
import './App.css';

const mockLists = () => {
  let results = []
  for(let i=0;i<40;i++){
    results.push({
      id: i,
      name: Math.random().toString(36).substring(2, 15) + Math.random().toString(36).substring(2, 15),
      count: Math.round(Math.random() * 100) * Math.round(Math.random() *  100)
    })
  }
  return results;
}

const data = { // TODO get data
  playlists: [
    mockLists(),
    mockLists(),
    mockLists(),
  ]
}

function PlaylistContainer(props) {
  const [tab, setTab] = React.useState(0);

  const handleChange = (event: React.ChangeEvent<{}>, newValue: number) => {
    setTab(newValue);
  };

  const displayPlaylists = (tab) => {
    return data.playlists[tab].map((list) => {
      return <Playlist
        key={list.id}
        openDrawer={props.openDrawer}
        currentTab={props.services[tab]}
        {...list}
      />
    })
  }

  return (
    <div>
      <Paper>
        <Tabs
          value={tab}
          onChange={handleChange}
          indicatorColor="primary"
          textColor="primary"
          variant="fullWidth"
          centered
        >
          { props.services.map(s => {
            return <Tab label={s} key={s} />
          })}
        </Tabs>
      </Paper>
      <List className="scrollable">
        {displayPlaylists(tab)}
      </List>
    </div>
  );
}

export default PlaylistContainer;
