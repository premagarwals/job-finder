import React from 'react';
import { useJobContext } from './JobContext';
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';
import { faLocationDot } from '@fortawesome/free-solid-svg-icons';

const Jobs = () => {
  const { jobs, nextPage, prevPage, offset, selectJob } = useJobContext();

  return (
    <div className='w-1/2 h-full bg-cyan-200 p-4 m-3 rounded'>
      <h2 className='text-2xl text-zinc-700 font-bold mb-4'>Jobs </h2>
      {jobs.length > 0 ? (
        <ul>
          {jobs.map((job, index) => (
            <li 
              key={index} 
              className='m-2 bg-cyan-100 p-3 rounded cursor-pointer' 
              onClick={() => selectJob(job)}
            >
              <h3 className='font-semibold text-zinc-600 text-lg'>{job.job_role}</h3>
              <p><FontAwesomeIcon icon={faLocationDot}/> {job.city}, {job.country}</p>
              <p>{job.work_site}</p>
              <p>${job.min_salary}-{job.max_salary}</p>
            </li>
          ))}
        </ul>
      ) : (
        <p>No jobs found.</p>
      )}
      <div className="pagination">
        <button className='bg-cyan-300 p-4 m-2 rounded text-zinc-700 font-semibold' onClick={prevPage} disabled={offset === 0}>Previous</button>
        <button className='bg-cyan-300 p-4 m-2 rounded text-zinc-700 font-semibold' onClick={nextPage}>Next</button>
      </div>
    </div>
  );
};

export default Jobs;
