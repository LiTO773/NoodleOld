import React from 'react'
import PropTypes from 'prop-types'
import './style.css'

// Icons
import iconLoading from '../../assets/fontawesome/sync-alt-solid.svg'
import iconSuccess from '../../assets/fontawesome/check-circle-solid.svg'
import iconError from '../../assets/fontawesome/exclamation-circle-solid.svg'

const Loader = props => {
  var icon = iconLoading
  var colorClass = ''
  if (props.status === 'success') {
    icon = iconSuccess
    colorClass = 'green'
  } else if (props.status === 'error' || props.status === 'critical') {
    icon = iconError
    colorClass = props.status === 'error' ? 'yellow' : 'red'
  }

  return (
    <div style={{ textAlign: 'center' }}>
      <img className={colorClass} height='190' src={icon} alt='Action icon' />
      <p>{props.msg}</p>
    </div>
  )
}

Loader.propTypes = {
  status: PropTypes.string, // success, error, critical or anything else for loading
  msg: PropTypes.string
}

export default Loader
