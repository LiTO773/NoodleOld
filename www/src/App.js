import React from 'react'
import './index.css'
import Header from './components/header'
import Footer from './components/footer'
import Home from './routes/home'
import Profile from './routes/profile'

import { BrowserRouter as Router, Switch, Route } from 'react-router-dom'

const App = () => {
  return (
    <Router>
      <Header />
      <div className='route'>
        <Switch>
          <Route path='/' exact component={Home} />
          <Route path='/addMoodle' component={Profile} />
        </Switch>
      </div>
      <Footer />
    </Router>
  )
}

export default App
