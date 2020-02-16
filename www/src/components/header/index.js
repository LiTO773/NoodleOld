import React from 'react'
import { Link } from 'react-router-dom'
import { bindActionCreators } from 'redux'
import { connect } from 'react-redux'
import './style.css'

import { goBack } from '../../actions/historyActions'

// Icons
import iconChevron from '../../assets/fontawesome/chevron-left-solid.svg'

const Header = props => {
  return (
    <header className='header'>
      {
        props.header.back &&
          <Link onClick={props.handleBack} to={props.history[props.history.length - 2]}>
            <img height='25' src={iconChevron} alt='Go back' />
          </Link>
      }
      <h1>
        {props.header.title}
      </h1>
    </header>
  )
}

const mapStateToProps = (state, props) => ({
  props,
  header: state.header,
  history: state.history
})

const mapActionsToProps = (dispatch, props) => (
  bindActionCreators({
    handleBack: goBack
  }, dispatch)
)

export default connect(mapStateToProps, mapActionsToProps)(Header)
