import React from 'react'
import PropTypes from 'prop-types'
import { bindActionCreators } from 'redux'
import { connect } from 'react-redux'

import { updateHeader } from '../../../actions/headerActions'
// import { goForward } from '../../../actions/historyActions'

const NewMoodle2 = props => {
  props.updateHeader(false, 'My Moodles')
  return (
    <div>
      <h1 className='view-title'>New Moodle 2</h1>
      <p>Nuff said</p>
    </div>
  )
}

NewMoodle2.propTypes = {
  updateHeader: PropTypes.func
}

const mapStateToProps = (state, props) => props
const mapActionsToProps = (dispatch, props) => (
  bindActionCreators({
    updateHeader
  }, dispatch)
)

export default connect(mapStateToProps, mapActionsToProps)(NewMoodle2)
