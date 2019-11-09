import React from 'react';
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
    from: 'spotify',
    to: 'youtube',
  }, {
    id: 1,
    name: 'Basic Playlist',
    recordsSuccess: 34,
    recordsFailed: 12,
    totalRecords: 70,
    jobFinished: false,
    from: 'spotify',
    to: 'youtube',
  }, {
    id: 1,
    name: 'Basic Playlist',
    recordsSuccess: 34,
    recordsFailed: 12,
    totalRecords: 70,
    jobFinished: false,
    from: 'spotify',
    to: 'youtube',
  }, {
    id: 2,
    name: 'Baller Playlist',
    recordsSuccess: 66,
    recordsFailed: 0,
    totalRecords: 120,
    jobFinished: false,
    from: 'apple',
    to: 'spotify',
  }, {
    id: 2,
    name: 'Baller Playlist',
    recordsSuccess: 66,
    recordsFailed: 0,
    totalRecords: 120,
    jobFinished: false,
    from: 'spotify',
    to: 'youtube',
  }, {
    id: 2,
    name: 'Baller Playlist',
    recordsSuccess: 66,
    recordsFailed: 0,
    totalRecords: 120,
    jobFinished: false,
    from: 'spotify',
    to: 'youtube',
  }, {
    id: 2,
    name: 'Baller Playlist',
    recordsSuccess: 66,
    recordsFailed: 0,
    totalRecords: 120,
    jobFinished: false,
    from: 'spotify',
    to: 'youtube',
  }, {
    id: 2,
    name: 'Baller Playlist',
    recordsSuccess: 66,
    recordsFailed: 0,
    totalRecords: 120,
    jobFinished: false,
    from: 'spotify',
    to: 'youtube',
  }, {
    id: 2,
    name: 'Baller Playlist',
    recordsSuccess: 66,
    recordsFailed: 0,
    totalRecords: 120,
    jobFinished: false,
    from: 'apple',
    to: 'spotify',
  }, {
    id: 2,
    name: 'Baller Playlist',
    recordsSuccess: 66,
    recordsFailed: 0,
    totalRecords: 120,
    jobFinished: false,
    from: 'spotify',
    to: 'youtube',
  }, {
    id: 3,
    name: 'Giga Playlist',
    recordsSuccess: 55,
    recordsFailed: 2,
    totalRecords: 57,
    jobFinished: true,
    from: 'youtube',
    to: 'apple',
  }, {
    id: 4,
    name: 'Giga Playlist',
    recordsSuccess: 55,
    recordsFailed: 2,
    totalRecords: 57,
    jobFinished: true,
    from: 'apple',
    to: 'youtube',
  }, {
    id: 5,
    name: 'Giga Playlist',
    recordsSuccess: 55,
    recordsFailed: 2,
    totalRecords: 57,
    jobFinished: true,
    from: 'spotify',
    to: 'youtube',
  }]
}

function JobsContainer() {
  const displayJobs = () => {
    return data.jobs.map((job) => {
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
