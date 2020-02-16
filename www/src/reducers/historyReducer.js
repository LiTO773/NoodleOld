import { GO_FORWARD, GO_BACK } from '../actions/historyActions'

export default function historyReducer (state = ['/'], { type, payload }) {
  switch (type) {
    case GO_FORWARD:
      return [...state, payload]
    case GO_BACK:
      return state.slice(0, -1)
    default:
      return state
  }
}
