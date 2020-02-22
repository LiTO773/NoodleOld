import React from 'react'
import PropTypes from 'prop-types'
import './style.css'

const Button = props => {
  return (
    <div disabled={props.disabled} className={'btn ' + (props.disabled ? '' : props.color)}>
      {props.text}
    </div>
  )
}

Button.propTypes = {
  color: PropTypes.string,
  text: PropTypes.string,
  disabled: PropTypes.bool
}

export default Button
