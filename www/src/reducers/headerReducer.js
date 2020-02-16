import { UPDATE_HEADER } from '../actions/headerActions'

export default function headerReducer (state = {}, { type, payload }) {
  switch (type) {
    case UPDATE_HEADER:
      return {
        back: payload.back,
        title: payload.title
      }
    default:
      return state
  }
}
