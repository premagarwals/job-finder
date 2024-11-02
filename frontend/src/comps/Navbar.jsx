import React, { useContext } from 'react';
import { AuthContext } from './AuthProvider';
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';
import { faCircleCheck, faEye, faPencil, faRightToBracket } from '@fortawesome/free-solid-svg-icons';

const Navbar = () => {
    const { isLoggedIn, setIsLoggedIn } = useContext(AuthContext);
    return (
        <div className='w-full h-16 bg-cyan-200 flex justify-between items-center font-semibold'>
            <div className='bg-cyan-100 px-6 py-3 m-2 flex items-center justify-center text-zinc-600 rounded font-semibold'>
                <p> <FontAwesomeIcon icon={isLoggedIn ? faCircleCheck : faEye} className='mr-1' />{isLoggedIn ? `Admin` : `Viewer`}</p>
            </div>
            <h3 className='text-xl text-zinc-700'>JOB FINDER</h3>
            <div className='bg-cyan-100 px-6 py-3 m-2 flex items-center justify-center text-zinc-600 rounded font-semibold cursor-pointer'>
                <p> <FontAwesomeIcon icon={isLoggedIn ? faRightToBracket: faPencil} className='mr-2' />{isLoggedIn ? `Exit Edit Mode` : `Enter Edit Mode`}</p>
            </div>
        </div>
    )
}

export default Navbar