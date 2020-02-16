import React from 'react'
// import { connect } from 'preact-redux'
import './style.css'

const Header = props => {
  return (
    <header className='header'>
      <h1>My Moodles</h1>
    </header>
  )
}

export default Header

// const mapStateToProps = (state, props) => ({
//   props,
//   header: state.header
// })

// export default connect(mapStateToProps)(Header)
