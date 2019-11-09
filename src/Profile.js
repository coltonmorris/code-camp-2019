import React from 'react';
import Grid from '@material-ui/core/Grid';
import PlaylistContainer from './PlaylistContainer';
import LinkAccountButton from './LinkAccountButton';
import GenericContainer from './GenericContainer';
import './App.css';

import { ToastContainer, toast } from 'react-toastify';
import 'react-toastify/dist/ReactToastify.css';

import { register } from './api';

const data = { // TODO get data
  profile: {
    name: "Brady",
    links: [{
      name: 'spotify',
      connected: true,
    }, {
      name: 'apple',
      connected: false,
    }, {
      name: 'youtube',
      connected: true,
    }]
  }
}

function Profile(props) {
  const displayLinkedAccounts = () => {
    return data.profile.links.map(service => {
      return <LinkAccountButton
        name={service.name}
        linked={service.connected}
        key={service.name}
        big
        onClick={() => authenticate(props.name, service.name)}
      />
    })
  }

  const authenticate = async (name, service) => {
    try {
      // get the redirect url
      const res = await register(name, service);
      // navigate to the redirect url to finish logging  in
      window.location = res.data;
      // the page will navigate back after authentication
    } catch (err) {
      console.error('Expected to receive a redirect URL.', err);
      toast.error(`Expected to receive a redirectl URL. ${err}`);
    }
  }

  return (
    <Grid className="ProfileContainer">
      <div className="threehundo">
        <Grid>
          <span className="ProfileTitle">Linked Accounts</span>
        </Grid>
        <br />
        <Grid className="LinkedAccountContainer">
          {displayLinkedAccounts()}
        </Grid>
      </div>
      <GenericContainer>
        <PlaylistContainer
          openDrawer={props.openDrawer}
          services={data.profile.links.map((svc) => {
            return svc.name;
          })}
        />
      </GenericContainer>
      <ToastContainer />
    </Grid>
  );
}

export default Profile;
