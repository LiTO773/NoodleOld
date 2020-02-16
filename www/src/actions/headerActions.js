export const UPDATE_HEADER = 'header:updateHeader'

/**
 * Allows each view to specify the content being displayed in the header
 * @param {boolean} back displays the back chevron in the header if true
 * @param {string} title title to be displayed
 */
export function updateHeader (back, title) {
  return {
    type: UPDATE_HEADER,
    payload: {
      back,
      title
    }
  }
}
