import React from 'react';
import Grid from '@material-ui/core/Grid';
import PlaylistContainer from './PlaylistContainer';
import LinkAccountButton from './LinkAccountButton';
import GenericContainer from './GenericContainer';
import './App.css';

import { ToastContainer, toast } from 'react-toastify';
import 'react-toastify/dist/ReactToastify.css';

import { authenticateApp } from './api';

const data = { // TODO get data
  profile: {
    name: "Brady",
    links: [{
      name: 'spotify',
      requestRedirect: "https://jsonplaceholder.typicode.com/users",
      connected: true,
    }, {
      name: 'apple',
      requestRedirect: "https://jsonplaceholder.typicode.com/users",
      connected: false,
    }, {
      name: 'youtube',
      // requestRedirect: "https://jsonplaceholder.typicode.com/users",
      requestRedirect: "https://google.com",
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
        onClick={() => authenticate(service.requestRedirect)}
      />
    })
  }

  const authenticate = async (redirect) => {
    try {
      // get the redirect url
      const url = await authenticateApp(redirect);
      const fakeUrl = "http://www.w3schools.com" // TODO
      // navigate to the redirect url to finish logging  in
      window.location = fakeUrl;
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
