import React from 'react'
import { bindActionCreators } from 'redux'
import { connect } from 'react-redux'

import Footer from '../../components/footer'
import { updateHeader } from '../../actions/headerActions'

const Home = props => {
  props.updateHeader(false, 'My Moodles')
  return (
    <div>
      <h1>All Moodles are displayed here</h1>
      <p>Nuff said</p>
      <Footer />
    </div>
  )
}

const mapStateToProps = (state, props) => props
const mapActionsToProps = (dispatch, props) => (
  bindActionCreators({
    updateHeader
  }, dispatch)
)

export default connect(mapStateToProps, mapActionsToProps)(Home)