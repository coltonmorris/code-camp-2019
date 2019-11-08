import React from 'react';
import Grid from '@material-ui/core/Grid';
import Profile from './Profile';
import './App.css';

function App() {
  return (
    <div className="App">
      <header className="App-header">
        SyncList.tech
      </header>
      <div className="AppOuter">
        <div className="AppContainer">
          <Profile />
        </div>
      </div>
    </div>
  );
}

export default App;
