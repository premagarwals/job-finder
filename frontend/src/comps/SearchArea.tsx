import React, { useContext, useState, useEffect } from 'react';
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';
import { faLocationDot, faSearch, faXmark } from '@fortawesome/free-solid-svg-icons';
import disciplinesData from './disciplines.json';
import { useJobContext } from './JobContext';

interface Profession {
  name: string;
  job_roles: string[];
}

interface Discipline {
  name: string;
  professions: Profession[];
}

const SearchArea: React.FC = () => {

  const [disciplines, setDisciplines] = useState<Discipline[]>([]);
  const [selectedDiscipline, setSelectedDiscipline] = useState<string>('');
  const [professions, setProfessions] = useState<Profession[]>([]);
  const [selectedProfession, setSelectedProfession] = useState<string>('');
  const [jobRoles, setJobRoles] = useState<string[]>([]);
  const [selectedJobRole, setSelectedJobRole] = useState<string>('');
  const [salary, setSalary] = useState<string | ''>('');
  const [location, setLocation] = useState<string>('');
  const [employmentType, setEmploymentType] = useState<string>('');
  const [workSite, setWorkSite] = useState<string>('');
  const [experience, setExperience] = useState<string>('');

  useEffect(() => {
    setDisciplines(disciplinesData.disciplines as Discipline[]);
  }, []);

  const handleDisciplineChange = (e: React.ChangeEvent<HTMLSelectElement>) => {
    const disciplineName = e.target.value;
    setSelectedDiscipline(disciplineName);

    const discipline = disciplines.find((d) => d.name === disciplineName);
    setProfessions(discipline ? discipline.professions : []);
    setSelectedProfession(''); 
    setJobRoles([]);
  };

  const handleProfessionChange = (e: React.ChangeEvent<HTMLSelectElement>) => {
    const professionName = e.target.value;
    setSelectedProfession(professionName);

    const profession = professions.find((p) => p.name === professionName);
    setJobRoles(profession ? profession.job_roles : []);
  };

  const { fetchJobs, setOffset, setFilterList } = useJobContext();

  const handleFindJobs = () => {
    const filterData = {
      job_role: selectedJobRole,
      profession: selectedProfession,
      discipline: selectedDiscipline,
      experience: experience,
      work_site: workSite, 
      city: location,   
      employment_type: employmentType, 
      salary: parseInt(salary),
    };

    setFilterList(filterData);
    setOffset(0);
  };

  return (
    <div className="w-11/12 m-auto mt-4 bg-cyan-100 flex flex-col p-4">

      {/* Inputs and Dropdowns */}
      <div className='flex w-full'>
        <div className='w-2/5 bg-cyan-50 mx-2 p-4 flex justify-between items-center'>
          <FontAwesomeIcon icon={faSearch} className='text-cyan-400' />
          <input type="number" value={salary} onChange={(e) => setSalary(e.target.value)} className='bg-transparent w-auto mx-2 w-5/6 focus:outline-none' placeholder='Expected salary...' />
          <FontAwesomeIcon icon={faXmark} />
        </div>
        <div className='w-2/5 bg-cyan-50 mx-2 p-4 flex justify-start items-center'>
          <FontAwesomeIcon icon={faLocationDot} className='text-cyan-400' />
          <input value={location} onChange={(e) => setLocation(e.target.value)} className='bg-transparent w-auto mx-2 w-5/6 focus:outline-none' placeholder='City, Country' />
        </div>
        <div className='w-1/5 bg-cyan-500 mx-2 p-4 font-bold text-center text-cyan-50 text-lg cursor-pointer' onClick={handleFindJobs}>
          Find Jobs
        </div>
      </div>

      {/* Dropdowns for additional filters */}
      <div className='mt-4 border-t-2 w-full border-cyan-200 p-4 flex gap-4'>
        <div className='p-3 rounded border-2 border-gray-500 inline-block text-gray-600'>
          <select className='bg-transparent' name="employmentType" id="employmentType" onChange={(e) => setEmploymentType(e.target.value)}>
            <option value="">Select Employment Type</option>
            <option value="Internship">Internship</option>
            <option value="Part Time">Part Time</option>
            <option value="Full Time">Full Time</option>
          </select>
        </div>

        <div className='p-3 rounded border-2 border-gray-500 inline-block text-gray-600'>
          <select className='bg-transparent' name="workSite" id="workSite" onChange={(e) => setWorkSite(e.target.value)}>
            <option value="">Select Worksite</option>
            <option value="remote">Remote</option>
            <option value="hybrid">Hybrid</option>
            <option value="on-site">On site</option>
          </select>
        </div>

        <div className='p-3 rounded border-2 border-gray-500 inline-block text-gray-600'>
          <select className='bg-transparent' name="discipline" id="discipline" onChange={handleDisciplineChange}>
            <option value="">Select Discipline</option>
            {disciplines.map((discipline, index) => (
              <option key={index} value={discipline.name}>
                {discipline.name}
              </option>
            ))}
          </select>
        </div>

        <div className='p-3 rounded border-2 border-gray-500 inline-block text-gray-600'>
          <select className='bg-transparent' name="profession" id="profession" onChange={handleProfessionChange} disabled={!selectedDiscipline}>
            <option value="">Select Profession</option>
            {professions.map((profession, index) => (
              <option key={index} value={profession.name}>
                {profession.name}
              </option>
            ))}
          </select>
        </div>
        
        <div className='p-3 rounded border-2 border-gray-500 inline-block text-gray-600'>
          <select className='bg-transparent' name="jobRole" id="jobRole" onChange={(e) => setSelectedJobRole(e.target.value)} disabled={!selectedProfession}>
            <option value="">Select Job Role</option>
            {jobRoles.map((role, index) => (
              <option key={index} value={role}>
                {role}
              </option>
            ))}
          </select>
        </div>

        <div className='p-3 rounded border-2 border-gray-500 inline-block text-gray-600'>
          <select className='bg-transparent' name="experience" id="experience" onChange={(e) => setExperience(e.target.value)}>
            <option value="">Select Experience</option>
            <option value="early">Early</option>
            <option value="intern">Intern</option>
            <option value="intermediate">Intermediate</option>
            <option value="expert">Expert</option>
            <option value="director">Director</option>
          </select>
        </div>
        
      </div>
    </div>
  );
};

export default SearchArea;
