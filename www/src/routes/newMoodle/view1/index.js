import React from 'react'
import { Link } from 'react-router-dom'
import { bindActionCreators } from 'redux'
import { connect } from 'react-redux'
import './style.css'

import Button from '../../../components/button'
import { updateHeader } from '../../../actions/headerActions'
import { goForward } from '../../../actions/historyActions'

const NewMoodle = props => {
  props.updateHeader(true, 'New Moodle (1/5)')
  return (
    <div>
      <p className='view-title'>Insert the Moodle's information in the fields below:</p>
      <div className='form'>
        <label className='label'>URL</label>
        <input className='field' type='text' />
        <label className='label'>Username</label>
        <input className='field' type='text' />
        <label className='label'>Password</label>
        <input className='field' type='password' />
      </div>
      <div style={{ position: 'fixed', bottom: 20, right: 20 }}>
        <Link to='/newMoodle2'>
          <Button text='Connect' color='btn-default' />
        </Link>
      </div>
    </div>
  )
}

const mapStateToProps = (state, props) => props
const mapActionsToProps = (dispatch, props) => (
  bindActionCreators({
    updateHeader,
    goForward
  }, dispatch)
)

export default connect(mapStateToProps, mapActionsToProps)(NewMoodle)
