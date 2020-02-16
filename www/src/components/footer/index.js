import React from 'react'
import { Link } from 'react-router-dom'
import './style.css'

// Icons
import iconSyncAlt from '../../assets/fontawesome/sync-alt-solid.svg'
import iconPlus from '../../assets/fontawesome/plus-solid.svg'
import iconCog from '../../assets/fontawesome/cog-solid.svg'

const Footer = () => (
  <div className='footer'>
    <img className='footer-icon' src={iconSyncAlt} alt='Refresh' />
    <div className='icon-container-middle clickable'>
      {/* Add highlight o image when selected */}
      <Link to='/addMoodle'>
        <img className='footer-icon' src={iconPlus} alt='Add Moodle' />
      </Link>
    </div>
    <img className='footer-icon' src={iconCog} alt='Settings' />
  </div>
)

export default Footer
