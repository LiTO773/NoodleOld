/**
 * Converts an error code number to a string to be displayed
 * @returns {string} The message to display
 */
export default errorCode => {
  switch (errorCode) {
    case 1:
      return 'Couldn\'t connect to the site. Please check your connection an try again.'
    case 2:
      return 'The data couldn\'t be saved locally!'
    case 3:
      return 'The data couldn\'t be saved locally!'
    case 4:
      return 'Moodle returned unexpected information.'
    case 5:
      return 'Oh oh! Due to Moodle Web Services being disabled in the specified on your website we\'re unable to connect :/ Please talk with the Moodle\'s admin in your to activate it.'
    case 6:
      return 'Seems like that information isn\'t available in your computer. Are you sure you\'ve registered this Moodle before?'
    case 7:
      return 'The data couldn\'t be saved locally!'
    case 8:
      return 'Seems like the session currently stored in your computer has expired.' // The user will likelly never see this one
    default:
      return ''
  }
}
