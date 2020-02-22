import { EDIT_VIEW1 } from '../actions/newMoodleActions'

export default function newMoodleReducer (state = {}, { type, payload }) {
  switch (type) {
    case EDIT_VIEW1:
      return {
        protocol: payload.protocol,
        url: payload.url,
        username: payload.username,
        password: payload.password
      }
    default:
      return state
  }
}
