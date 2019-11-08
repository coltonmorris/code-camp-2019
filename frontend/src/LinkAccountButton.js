import React from 'react';
import Button from '@material-ui/core/Button';
import './App.css';

function LinkAccountButton(props) {
  return (
    <Button
      variant="outlined"
      color={props.linked ? "primary" : "secondary"}
    >
      {props.name}
    </Button>
  );
}

export default LinkAccountButton;
