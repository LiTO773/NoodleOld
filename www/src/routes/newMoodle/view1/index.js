import React, { useState } from 'react'
import PropTypes from 'prop-types'
import { useHistory } from 'react-router-dom'
import { bindActionCreators } from 'redux'
import { connect } from 'react-redux'
import './style.css'

import Button from '../../../components/button'
import { updateHeader } from '../../../actions/headerActions'
import { goForward } from '../../../actions/historyActions'
import { editView1 } from '../../../actions/newMoodleActions'

const NewMoodle = props => {
  // Check if the state has already been declared
  const stateToUse = props.infos.url !== '' ? { ...props.infos } : { protocol: 'http://', url: '', username: '', password: '' }
  const [formState, setFormState] = useState(stateToUse)
  const history = useHistory()

  props.updateHeader(true, 'New Moodle (1/3)')

  const handleProtocolChange = event => {
    setFormState({ ...formState, protocol: event.target.value })
  }

  const handleURLChange = event => {
    setFormState({ ...formState, url: event.target.value })
  }

  const handleUsernameChange = event => {
    setFormState({ ...formState, username: event.target.value })
  }

  const handlePasswordChange = event => {
    setFormState({ ...formState, password: event.target.value })
  }

  const handleForward = () => {
    if (formState.url !== '' && formState.username !== '' && formState.password !== '') {
      props.editView1(formState.protocol, formState.url, formState.username, formState.password)
      props.goForward('/newMoodle2')
      history.push('/newMoodle2', { ...formState })
    }
  }

  return (
    <div>
      <p className='view-title'>Insert the Moodle's information in the fields below:</p>
      <div className='form'>
        <label className='label'>URL</label>
        <div className='max' style={{ display: 'flex' }}>
          <select className='field' onChange={handleProtocolChange} value={formState.protocol}>
            <option value='http://'>http://</option>
            <option value='https://'>https://</option>
          </select>
          <input onChange={handleURLChange} value={formState.url} className='field' style={{ flexGrow: '2' }} type='text' />
        </div>
        <label className='label'>Username</label>
        <input onChange={handleUsernameChange} value={formState.username} className='field max' type='text' />
        <label className='label'>Password</label>
        <input onChange={handlePasswordChange} value={formState.password} className='field max' type='password' />
      </div>
      <div onClick={handleForward} style={{ position: 'fixed', bottom: 20, right: 20 }}>
        <Button
          text='Connect'
          color='btn-default'
          disabled={formState.url === '' || formState.username === '' || formState.password === ''}
        />
      </div>
    </div>
  )
}

NewMoodle.propTypes = {
  updateHeader: PropTypes.func,
  goForward: PropTypes.func,
  editView1: PropTypes.func,
  infos: PropTypes.object
}

const mapStateToProps = (state, props) => ({
  props,
  infos: state.newMoodle
})

const mapActionsToProps = (dispatch, props) => (
  bindActionCreators({
    updateHeader,
    goForward,
    editView1
  }, dispatch)
)

export default connect(mapStateToProps, mapActionsToProps)(NewMoodle)
