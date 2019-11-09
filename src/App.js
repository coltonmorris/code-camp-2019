import React, { useState } from 'react';
import Drawer from '@material-ui/core/Drawer';
import Button from '@material-ui/core/Button';
import MenuIcon from '@material-ui/icons/Menu';
import UnfoldMoreIcon from '@material-ui/icons/UnfoldMore';
import Profile from './Profile';
import LoginPage from './LoginPage';
import JobsContainer from './JobsContainer';
import GenericContainer from './GenericContainer';
import { createMuiTheme, ThemeProvider } from '@material-ui/core/styles';
import grey from '@material-ui/core/colors/grey';
import './App.css';

const outerTheme = createMuiTheme({
  palette: {
    secondary: {
      main: grey[50]
    },
  },
});

function App() {
  const [loggedIn, setLoggedIn] = useState(true);
  const [drawerOpen, setDrawerOpen] = useState(false);

  const displayFooter = () => {
    return (
      <div
        className="App-footer"
        onClick={() => setDrawerOpen(o => !o)}
        style={{ cursor: "pointer" }}
      >
        <UnfoldMoreIcon fontSize="inherit" />
        Jobs List
        <UnfoldMoreIcon fontSize="inherit"/>
      </div>
    );
  }

  return (
    <ThemeProvider theme={outerTheme}>
      <div className="App">
        <header className="App-header">
          <span className="BasicPad">
            <Button
              variant="outlined"
              color="secondary"
              style={{
                minWidth: "30px",
                padding: "6px"
              }}
            >
              <MenuIcon />
            </Button>
          </span>
          <span>SyncList.tech</span>
          <span className="rest">
          </span>
        </header>
        <div className="GenericContainer">
          <div className="AppContainer">
            <div className="AppBody">
              {
                loggedIn ?
                  <Profile openDrawer={() => setDrawerOpen(true)} /> :
                  <LoginPage login={setLoggedIn} />
              }
            </div>
          </div>
        </div>

        <Drawer
          anchor="bottom"
          variant="temporary"
          open={drawerOpen}
          onClose={() => setDrawerOpen(false)}
        >
          { displayFooter() }
          <GenericContainer>
            <JobsContainer />
          </GenericContainer>
        </Drawer>

        { drawerOpen ? "" : displayFooter() }
      </div>
    </ThemeProvider>
  );
}

export default App;
