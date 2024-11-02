import React from 'react';
import { useJobContext } from './JobContext';

const JobView = () => {
  const { selectedJob } = useJobContext();

  if (!selectedJob) {
    return <div className='w-full h-full flex items-center justify-center bg-cyan-200 m-3 p-4 rounded'>
        Please select a job to view details.
    </div>;
  }

  return (
    <div className='w-full h-full bg-cyan-200 p-4 m-3 rounded'>
      <h2 className='text-2xl text-zinc-700 font-bold mb-4'>#{selectedJob.job_id} - {selectedJob.job_role} </h2>
      <p><strong>Location:</strong> {selectedJob.city}, {selectedJob.country}</p>
      <p><strong>Work Site:</strong> {selectedJob.work_site}</p>
      <p><strong>Salary:</strong> ${selectedJob.min_salary} - ${selectedJob.max_salary}</p>
      <p><strong>Employment Type:</strong> {selectedJob.employment_type}</p>
      <p><strong>Discipline:</strong> {selectedJob.discipline}</p>
      <p><strong>Profession:</strong> {selectedJob.profession}</p>
      <p></p><br /><br />
      <p><strong>Description:</strong></p>
      <p>{selectedJob.job_description}</p>
    </div>
  );
};

export default JobView;
