import React from 'react';
import Button from '@material-ui/core/Button';
import './App.css';

// big if true
function LinkAccountButton(props) {
  const getIconImage = () => {
    switch(props.name){
      case 'apple': return '/images/apple.png';
      case 'spotify': return '/images/spotify.png';
      case 'youtube': return '/images/youtube.png';
      default: return ""
    }
  }

  return (
    <Button
      variant="contained"
      color={props.linked ? "primary" : "secondary"}
      className={props.big ? "LinkAccountButton" : "SyncPlaylistButton"}
      style={{
        backgroundImage: `url(${getIconImage()})`,
        minWidth: "40px"
      }}
    >
      {""}
    </Button>
  );
}

export default LinkAccountButton;
