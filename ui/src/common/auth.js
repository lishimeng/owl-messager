export default {
  getUid: function() {
    var uid = window.localStorage.getItem('userId')
    return uid
  },

  getSid: function() {
    var sid = window.localStorage.getItem('sessionId')
    return sid
  },

  login(userId, token) {
    window.localStorage.setItem('userId', userId)
    window.localStorage.setItem('sessionId', token)
  },
  logout(cb) {
    window.localStorage.removeItem('userId')
    window.localStorage.removeItem('sessionId')
    window.localStorage.removeItem('userName')
    window.localStorage.removeItem('roles')
    if (cb) {
      cb()
    }
  },
  loggedIn() {
    return !!window.localStorage.getItem('userId')
  }

}
