import React from 'react'
import { Link } from 'react-router-dom'
import { bindActionCreators } from 'redux'
import { connect } from 'react-redux'

import { updateHeader } from '../../../actions/headerActions'

const NewMoodle = props => {
  props.updateHeader(true, 'New Moodle (1/5)')
  return (
    <div>
      <h1>Add Moodle</h1>
      <Link to='/'><p>Go back</p></Link>
    </div>
  )
}

const mapStateToProps = (state, props) => props
const mapActionsToProps = (dispatch, props) => (
  bindActionCreators({
    updateHeader
  }, dispatch)
)

export default connect(mapStateToProps, mapActionsToProps)(NewMoodle)
