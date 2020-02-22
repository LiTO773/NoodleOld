import React from 'react'
import ReactDOM from 'react-dom'
import App from './App'
import * as serviceWorker from './serviceWorker'

import { Provider } from 'react-redux'
import { createStore, combineReducers } from 'redux'
import historyReducer from './reducers/historyReducer'
import headerReducer from './reducers/headerReducer'
import newMoodleReducer from './reducers/newMoodleReducer'

const allReducers = combineReducers({
  header: headerReducer,
  history: historyReducer,
  newMoodle: newMoodleReducer
})

const store = createStore(
  allReducers,
  {
    header: {
      back: false,
      title: 'My Moodles'
      // TODO implement custom right side actions
    },
    history: ['/'],
    newMoodle: {
      protocol: 'http://',
      url: '',
      username: '',
      password: ''
    }
  },
  window.__REDUX_DEVTOOLS_EXTENSION__ && window.__REDUX_DEVTOOLS_EXTENSION__()
)

ReactDOM.render(
  <Provider store={store}>
    <App />
  </Provider>,
  document.getElementById('root')
)

// If you want your app to work offline and load faster, you can change
// unregister() to register() below. Note this comes with some pitfalls.
// Learn more about service workers: https://bit.ly/CRA-PWA
serviceWorker.unregister()
