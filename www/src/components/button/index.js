import React from 'react'
import './style.css'

const Button = props => {
  return (
    <div className={'btn ' + props.color}>
      {props.text}
    </div>
  )
}

export default Button
