export const EDIT_VIEW1 = 'newMoodle:editView1'

/**
 * Saves the data inserted in view1
 * @param {string} protocol protocol chosen
 * @param {string} url url chosen
 * @param {string} username username chosen
 * @param {string} password password chosen
 */
export function editView1 (protocol, url, username, password) {
  return {
    type: EDIT_VIEW1,
    payload: {
      protocol,
      url,
      username,
      password
    }
  }
}
