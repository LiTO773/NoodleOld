/**
 * Receives an erro code and return if it is critical
 * @return {boolean} true if critical
 */
export default errorCode => {
  switch (errorCode) {
    case 2:
    case 3:
    case 5:
    case 6:
    case 7:
      return true
  }

  return false
}
