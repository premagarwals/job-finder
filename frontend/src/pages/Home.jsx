import React from 'react'
import SearchArea from '../comps/SearchArea'
import Navbar from '../comps/Navbar'
import { JobProvider } from '../comps/JobContext'
import Jobs from '../comps/Jobs'
import JobView from '../comps/JobView'

const Home = () => {
  return (
    <div className='h-screen w-screen flex flex-col'>
      <Navbar />
      <JobProvider>
        <div className='w-screen bg-cyan-50 shadow-inner p-4'>
          <SearchArea />
        </div>
        <div className='flex h-full'>
          <Jobs />
          <JobView />
        </div>
      </JobProvider>

      </div>
  )
}

export default Home
