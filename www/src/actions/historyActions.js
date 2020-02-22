export const GO_FORWARD = 'history:push'
export const GO_BACK = 'history:pop'
export const REMOVE_ENTRIES = 'history:popAmount'

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

/**
 * Removes the x latest entries in the history
 * Note: The header component will always remove one when the chevron is clicked,
 * which should be taken into consideration when using this function
 * @param {number} x num of entries
 */
export function removeEntries (x) {
  return {
    type: REMOVE_ENTRIES,
    payload: x
  }
}
