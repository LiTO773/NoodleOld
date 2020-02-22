/* global newUser */
import React, { useState } from 'react'
import PropTypes from 'prop-types'
// import { useHistory } from 'react-router-dom'
import { bindActionCreators } from 'redux'
import { connect } from 'react-redux'

import { updateHeader } from '../../../actions/headerActions'
import { removeEntries } from '../../../actions/historyActions'
import Loader from '../../../components/loaderView'
import isCriticalError from '../../../helpers/isCriticalError'
import errorToString from '../../../helpers/errorToString'

const NewMoodle2 = props => {
  const [statusState, setStatusState] = useState({ status: 'loading', msg: 'Connecting to Moodle...' })

  if (typeof newUser === 'function' && statusState.status === 'loading') {
    props.updateHeader(false, 'Connecting to Moodle...')
    newUser(props.location.state.protocol + props.location.state.url, props.location.state.username, props.location.state.password, '')
      .then(result => {
        // It ran flawlessly
        if (result === 0) {
          // Go to next view
        }

        // Something happened
        props.updateHeader(true, 'An error occured...')
        // If it is critical, the chevron should navigate back to the home screen
        if (isCriticalError(result)) {
          props.removeEntries(1)
        }

        setStatusState({
          status: isCriticalError(result) ? 'critical' : 'error',
          msg: errorToString(result)
        })
      })
  }

  return (
    <div>
      <Loader status={statusState.status} msg={statusState.msg} />
    </div>
  )
}

NewMoodle2.propTypes = {
  updateHeader: PropTypes.func,
  removeEntries: PropTypes.func,
  location: PropTypes.object
}

const mapStateToProps = (state, props) => props
const mapActionsToProps = (dispatch, props) => (
  bindActionCreators({
    updateHeader,
    removeEntries
  }, dispatch)
)

export default connect(mapStateToProps, mapActionsToProps)(NewMoodle2)
