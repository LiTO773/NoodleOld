export const GO_FORWARD = 'history:push'
export const GO_BACK = 'history:pop'

/**
 * Every time it goes forward, it stores the path
 * @param {string} path corresponds to the link of the next view that will be loaded
 */
export function goForward (path) {
  return {
    type: GO_FORWARD,
    payload: path
  }
}

/**
 * Every time it goes back, it will remove the path of the current view
 */
export function goBack () {
  return {
    type: GO_BACK,
    payload: '' // doesn't matter
  }
}