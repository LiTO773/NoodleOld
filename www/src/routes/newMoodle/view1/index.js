import React, { useState } from 'react'
import PropTypes from 'prop-types'
import { Link } from 'react-router-dom'
import { bindActionCreators } from 'redux'
import { connect } from 'react-redux'
import './style.css'

import Button from '../../../components/button'
import { updateHeader } from '../../../actions/headerActions'
import { goForward } from '../../../actions/historyActions'

const NewMoodle = props => {
  const [formState, setFormState] = useState({ url: '', username: '', password: '' })

  props.updateHeader(true, 'New Moodle (1/5)')

  const handleURLChange = event => {
    setFormState({ ...formState, url: event.target.value })
  }

  const handleUsernameChange = event => {
    setFormState({ ...formState, username: event.target.value })
  }

  const handlePasswordChange = event => {
    setFormState({ ...formState, password: event.target.value })
  }

  return (
    <div>
      <p className='view-title'>Insert the Moodle's information in the fields below:</p>
      <div className='form'>
        <label className='label'>URL</label>
        <input onChange={handleURLChange} value={formState.url} className='field' type='text' />
        <label className='label'>Username</label>
        <input onChange={handleUsernameChange} value={formState.username} className='field' type='text' />
        <label className='label'>Password</label>
        <input onChange={handlePasswordChange} value={formState.password} className='field' type='password' />
      </div>
      <div style={{ position: 'fixed', bottom: 20, right: 20 }}>
        <Link to={{
          pathname: '/newMoodle2',
          state: {
            ...formState
          }
        }}
        >
          <Button text='Connect' color='btn-default' />
        </Link>
      </div>
    </div>
  )
}

NewMoodle.propTypes = {
  updateHeader: PropTypes.func
}

const mapStateToProps = (state, props) => props
const mapActionsToProps = (dispatch, props) => (
  bindActionCreators({
    updateHeader,
    goForward
  }, dispatch)
)

export default connect(mapStateToProps, mapActionsToProps)(NewMoodle)
