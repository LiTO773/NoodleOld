import React from 'react'
import './index.css'
import Header from './components/header'
import Home from './routes/home'
import NewMoodle1 from './routes/newMoodle/view1'

import { BrowserRouter as Router, Switch, Route } from 'react-router-dom'

const App = () => {
  return (
    <Router>
      <Header />
      <div className='route'>
        <Switch>
          <Route path='/' exact component={Home} />
          <Route path='/addMoodle' component={NewMoodle1} />
        </Switch>
      </div>
    </Router>
  )
}

export default App
