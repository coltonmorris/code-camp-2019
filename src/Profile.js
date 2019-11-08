import React from 'react';
import Grid from '@material-ui/core/Grid';
import JobsContainer from './JobsContainer';
import PlaylistContainer from './PlaylistContainer';
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

function Profile() {
  const displayLinkedAccounts = () => {
    return Object.keys(data.profile.links).map(key => {
      return <LinkAccountButton name={key} linked={data.profile.links[key]} key={key} />
    })
  }

  return (
    <Grid className="ProfileContainer">
      <Grid>
        <span>{`${data.profile.name}'s Profile`}</span>
      </Grid>
      <Grid className="LinkedAccountContainer">
        {displayLinkedAccounts()}
      </Grid>
      <PlaylistContainer />
      <JobsContainer />
    </Grid>
  );
}

export default Profile;
