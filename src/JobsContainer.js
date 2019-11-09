import React from 'react';
import ProgressBar from './ProgressBar';
import './App.css';

function JobsContainer(props) {
  const displayJobs = () => {
    return props.jobs.map((job) => {
      return <ProgressBar {...job} key={job.id} />
    })
  }

  return (
    <div>
      <br />
      {displayJobs()}
      <br />
    </div>
  );
}

export default JobsContainer;
