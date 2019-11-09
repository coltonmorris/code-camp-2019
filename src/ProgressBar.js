import React from 'react';
// import Button from '@material-ui/core/Button';
import Grid from '@material-ui/core/Grid';
import LinearProgress from '@material-ui/core/LinearProgress';

function ProgressBar(props) {
  const calculateProgress = () => {
    return parseFloat(props.recordsSuccess + props.recordsFailed) / props.totalRecords * 100;
  }

  return (
    <Grid>
      <div className="ProgressBarTitle">{props.name}</div>
      <LinearProgress
        style={{ height: "20px" }}
        className="ProgressBar"
        variant="determinate"
        value={calculateProgress()}
      />
    </Grid>
  )
}

export default ProgressBar;
