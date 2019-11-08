import React from 'react';
import Button from '@material-ui/core/Button';
import './App.css';

function SyncPlaylistButton(props) {
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
      className="SyncPlaylistButton"
      style={{
        backgroundImage: `url(${getIconImage()})`,
        backgroundSize: "140px 140px",
      }}
    >
      {""}
    </Button>
  );
}

export default SyncPlaylistButton;
