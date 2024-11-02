// JobContext.js
import React, { createContext, useState, useEffect, useContext } from 'react';
import backend from '../config';

const JobContext = createContext();

export const JobProvider = ({ children }) => {
  const [filterList, setFilterList] = useState([]);
  const [jobs, setJobs] = useState([]);
  const [offset, setOffset] = useState(0);
  const [limit, setLimit] = useState(5);
  const [selectedJob, setSelectedJob] = useState(null);

  const fetchJobs = async () => {
    try {
      const response = await fetch(`${backend}/jobs`, {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify({
          ...filterList,
          limit: limit,
          offset: offset,
        }),
      });

      const data = await response.json();
      setJobs(data["jobs"]);
    } catch (error) {
      console.error('Error fetching jobs:', error);
    }
  };

  useEffect(() => {
    if (Object.keys(filterList).length > 0) {
      fetchJobs();
    }
  }, [filterList, offset, limit]);

  const selectJob = (job) => {
    setSelectedJob(job);
  };

  const nextPage = () => {
    if (jobs.length>=5) {setOffset((prev) => prev + limit); } 
  };

  const prevPage = () => {
    setOffset((prev) => Math.max(prev - limit, 0));
  };

  return (
    <JobContext.Provider value={{ jobs, fetchJobs, nextPage, prevPage, offset, setOffset, selectedJob, selectJob, filterList, setFilterList }}>
      {children}
    </JobContext.Provider>
  );
};

export const useJobContext = () => useContext(JobContext);
