import React from 'react'
import { connect } from 'react-redux'
import './style.css'

import iconChevron from '../../assets/fontawesome/chevron-left-solid.svg'

const Header = props => {
  return (
    <header className='header'>
      {
        props.header.back && <img height='25' src={iconChevron} alt='' />
      }
      <h1>
        {props.header.title}
      </h1>
    </header>
  )
}

const mapStateToProps = (state, props) => ({
  props,
  header: state.header
})

export default connect(mapStateToProps)(Header)
