import React from 'react'
import './index.css'
import 'normalize.css'
import Header from './components/header'
import Home from './routes/home'
import NewMoodle from './routes/newMoodle/view1'
import NewMoodle2 from './routes/newMoodle/view2'
import NewMoodle3 from './routes/newMoodle/view3'

import { BrowserRouter as Router, Switch, Route } from 'react-router-dom'

const App = () => {
  return (
    <Router>
      <Header />
      <div className='route'>
        <Switch>
          <Route path='/' exact component={Home} />
          <Route path='/newMoodle' component={NewMoodle} />
          <Route path='/newMoodle2' component={NewMoodle2} />
          <Route path='/newMoodle3' component={NewMoodle3} />
        </Switch>
      </div>
    </Router>
  )
}

export default App
