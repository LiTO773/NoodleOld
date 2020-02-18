import React from 'react'
import PropTypes from 'prop-types'
import './style.css'

const Button = props => {
  return (
    <div className={'btn ' + props.color}>
      {props.text}
    </div>
  )
}

Button.propTypes = {
  color: PropTypes.string,
  text: PropTypes.string
}

export default Button
