import React from 'react';
import Grid from '@material-ui/core/Grid';
import './App.css';

function GenericContainer(props) {
  return (
    <Grid className="GenericContainer">
      <Grid className="GenericContainerBody">
        {props.children}
      </Grid>
    </Grid>
  );
}

export default GenericContainer;
