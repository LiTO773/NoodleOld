/* global newUser */
import React from 'react'
// import { Redirect } from 'react-router'
import PropTypes from 'prop-types'
import { bindActionCreators } from 'redux'
import { connect } from 'react-redux'

import { updateHeader } from '../../../actions/headerActions'
// import { goForward } from '../../../actions/historyActions'

const NewMoodle2 = props => {
  console.log(props)

  const handleClick = () => {
    if (typeof newUser === 'function') {
      newUser(props.location.state.protocol + props.location.state.url, props.location.state.username, props.location.state.password, '')
        .then(result => console.log(result))
    }
  }

  props.updateHeader(false, 'My Moodles')
  return (
    <div>
      <h1 className='view-title'>New Moodle 2</h1>
      <p onClick={handleClick}>Nuff said</p>
    </div>
  )
}

NewMoodle2.propTypes = {
  updateHeader: PropTypes.func,
  location: PropTypes.object
}

const mapStateToProps = (state, props) => props
const mapActionsToProps = (dispatch, props) => (
  bindActionCreators({
    updateHeader
  }, dispatch)
)

export default connect(mapStateToProps, mapActionsToProps)(NewMoodle2)
