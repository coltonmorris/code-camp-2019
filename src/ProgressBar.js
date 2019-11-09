import React from 'react';
import Grid from '@material-ui/core/Grid';
import LinearProgress from '@material-ui/core/LinearProgress';
import HourglassEmptyIcon from '@material-ui/icons/HourglassEmpty';
import DoneOutlineIcon from '@material-ui/icons/DoneOutline';

function ProgressBar(props) {
  const calculateProgress = () => {
    return parseFloat(props.recordsSuccess + props.recordsFailed) / props.totalRecords * 100;
  }

  const progress = calculateProgress();
  return (
    <Grid>
      <div
        className="Progress"
        style={{ fontSize: '1.75vmin' }}
      >
        <span className="ProgressBarTitle">
            {props.name}
        </span>
        <span style={{ width: '300px', textAlign: 'center' }}>
          {`${props.from} ⟶   ${props.to}`}
        </span>
        <span style={{ width: '100px', textAlign: 'right' }}>
          {`${props.recordsSuccess+props.recordsFailed} / ${props.totalRecords}`}
        </span>
        <span style={{ width: '50px', textAlign: 'center' }}>
          { progress === 100 ?
            <DoneOutlineIcon fontSize="inherit" /> :
            <HourglassEmptyIcon fontSize="inherit" className="Spinner" />
          }
        </span>
      </div>
      <LinearProgress
        style={{ height: "20px" }}
        className="ProgressBar"
        variant="determinate"
        value={progress}
      />
    </Grid>
  )
}

export default ProgressBar;
