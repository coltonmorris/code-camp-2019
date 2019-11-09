import React from 'react';
import Button from '@material-ui/core/Button';
import Grid from '@material-ui/core/Grid';

function LoginPage(props) {
  return (
    <Grid>
      Please Log In to Continue
      <Button
        variant="contained"
        color="primary"
        onClick={props.login}
      >
        Login
      </Button>
    </Grid>
  )
}

export default LoginPage;
