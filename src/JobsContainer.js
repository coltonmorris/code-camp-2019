import React from 'react';
import Grid from '@material-ui/core/Grid';
import ProgressBar from './ProgressBar';
import './App.css';

const data = { // TODO get data
  jobs: [{
    id: 1,
    name: 'Basic Playlist',
    recordsSuccess: 34,
    recordsFailed: 12,
    totalRecords: 70,
    jobFinished: false,
  }, {
    id: 2,
    name: 'Baller Playlist',
    recordsSuccess: 66,
    recordsFailed: 0,
    totalRecords: 120,
    jobFinished: false,
  }, {
    id: 3,
    name: 'Giga Playlist',
    recordsSuccess: 55,
    recordsFailed: 2,
    totalRecords: 57,
    jobFinished: true,
  }]
}

function JobsContainer() {
  const displayJobs = () => {
    return data.jobs.map((job) => {
      return <ProgressBar {...job} key={job.id} />
    })
  }

  return (
    <Grid className="JobsOuter">
      <Grid className="JobsContainer">
        <div className="JobsTitle">Jobs</div>
        {displayJobs()}
      </Grid>
    </Grid>
  );
}

export default JobsContainer;
