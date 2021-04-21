export const getSessionKey = (key, defaultValue) => {
  const item = window.sessionStorage.getItem(key)
  if (item === undefined && defaultValue !== undefined) {
    return defaultValue
  }
  return item
}

export const getLocalStore = (key) => {
  return window.localStorage.getItem(key)
}

export const setLocalStore = (key, value) => {
  if (key && value) {
    window.localStorage.setItem(key, value)
  }
}

export const getLocal = () => {
  if (window.localStorage.getItem('locale')) {
    if (window.localStorage.getItem('locale') === 'zh') {
      return 'zh_CN'
    } else if (window.localStorage.getItem('locale') === 'en') {
      return 'en_US'
    }
  } else {
    return 'zh_CN'
  }
}
