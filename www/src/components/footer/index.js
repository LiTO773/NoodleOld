import React from 'react'
import { Link } from 'react-router-dom'
import { bindActionCreators } from 'redux'
import { connect } from 'react-redux'
import './style.css'

import { goForward } from '../../actions/historyActions'

// Icons
import iconSyncAlt from '../../assets/fontawesome/sync-alt-solid.svg'
import iconPlus from '../../assets/fontawesome/plus-solid.svg'
import iconCog from '../../assets/fontawesome/cog-solid.svg'

const Footer = props => {
  return (
    <div className='footer'>
      <img className='footer-icon' src={iconSyncAlt} alt='Refresh' />
      <div className='icon-container-middle clickable'>
        {/* Add highlight o image when selected */}
        <Link onClick={props.handleForward.bind(this, '/addMoodle')} to='/addMoodle'>
          <img className='footer-icon' src={iconPlus} alt='Add Moodle' />
        </Link>
      </div>
      <img className='footer-icon' src={iconCog} alt='Settings' />
    </div>
  )
}

const mapStateToProps = (state, props) => props
const mapActionsToProps = (dispatch, props) => (
  bindActionCreators({
    handleForward: goForward
  }, dispatch)
)

export default connect(mapStateToProps, mapActionsToProps)(Footer)
