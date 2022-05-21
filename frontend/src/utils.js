export const serverURL = process.env.NODE_ENV === 'development' ? 'http://localhost:3000' : '/api'
export const socketURL = process.env.NODE_ENV === 'development' ? 'http://localhost:3000' : '/'

// Errors enum
export const ERRORS = Object.freeze({
  NOT_SIGNED_IN: "not-signed-in",
  USER_DOES_NOT_EXIST: "user-does-not-exist",
})

/* 
  Date utils 
*/
export const getDateString = (date) => {
  date = new Date(date)
  return `${date.getMonth()+1}/${date.getDate()}`
}

export const getDateRangeString = (date1, date2) => {
  date1 = new Date(date1)
  date2 = new Date(date2)
  return  getDateString(date1) + ' - ' + getDateString(date2)
}

export const getDateWithTime = (date, timeString) => {
  /* Returns a new date object with the given date (e.g. 5/2/2022) and the specified time (e.g. 11:30) */
  date = new Date(date)

  const { hours, minutes } = splitTime(timeString)
  return new Date(date.getFullYear(), date.getMonth(), date.getDate(), hours, minutes)
}

export const getDateWithTimeInt = (date, timeInt) => {
  /* Returns a new date object with the given date (e.g. 5/2/2022) and the specified timeInt (e.g. 11.5) */
  date = new Date(date)

  const hours = parseInt(timeInt)
  const minutes = (timeInt - hours) * 60
  return new Date(date.getFullYear(), date.getMonth(), date.getDate(), hours, minutes)
}

export const splitTime = (timeString) => {
  /* Takes a time string (e.g. 13:30) and splits it into hours and minutes, returning an object of the form { hours, minutes } */
  const [hours, minutes] = timeString.split(':')
  return { hours: parseInt(hours), minutes: parseInt(minutes) }
}

export const getDateDayOffset = (date, offset) => {
  /* Returns the specified date offset by the given number of days (can be positive or negative) */
  date = new Date(date)
  return new Date(date.getTime() + offset * 24*60*60*1000)
}

export const timeIntToTimeText = (timeInt) => {
  /* Converts a timeInt (e.g. 13) to a timeText (e.g. "1 pm") */
  if (timeInt == 0) return '12 am'
  else if (timeInt <= 11) return `${timeInt} am`
  else if (timeInt == 12) return '12 pm'
  return `${timeInt - 12} pm`
}

export const dateToTimeInt = (date) => {
  /* Converts a date to a timeInt (e.g. 9.5) */
  date = new Date(date)
  return date.getHours() + date.getMinutes() / 60
}

export const clampDateToTimeInt = (date, timeInt, type) => {
  /* Clamps the date to the given time, type can either be "upper" or "lower" */
  const diff = dateToTimeInt(date) - timeInt
  if (type === 'upper' && diff < 0) {
    return getDateWithTimeInt(date, timeInt)     
  } else if (type === 'lower' && diff > 0) {
    return getDateWithTimeInt(date, timeInt)
  }
  
  // Return original date
  return date
}

export const dateCompare = (date1, date2) => {
  /* Returns negative if date1 < date2, positive if date2 > date1, and 0 if date1 == date2 */
  date1 = new Date(date1)
  date2 = new Date(date2)
  return date1.getTime() - date2.getTime()
}

export const compareDateDay = (a, b) => {
  // returns -1 if a is before b, 1 if a is after b, 0 otherwise
  a = new Date(a)
  b = new Date(b)
  if (a.getFullYear() !== b.getFullYear()) {
    return a.getFullYear() - b.getFullYear()
  } else if (a.getMonth() !== b.getMonth()) {
    return a.getMonth() - b.getMonth()
  } else {
    return a.getDate() - b.getDate()
  }
}

/* 
  Fetch utils
*/
export const get = (route) => {
  return fetchMethod('GET', route)
}

export const post = (route, body={}) => {
  return fetchMethod('POST', route, body)
}

export const patch = (route, body={}) => {
  return fetchMethod('PATCH', route, body)
}

export const _delete = (route, body={}) => {
  return fetchMethod('DELETE', route, body)
}

export const fetchMethod = (method, route, body={}) => {
  /* Calls the given route with the give method and body */
  const params = {
    method,
    credentials: 'include',
  }

  if (method !== 'GET') {
    // Add params specific to POST/PATCH/DELETE
    params.headers = {
      'Content-Type': 'application/json'
    }
    params.body = JSON.stringify(body)
  }

  return fetch(serverURL + route, params).then(async res => {
    // Parse data if it is json, otherwise just return an empty object
    const text = await res.text()
    try {
      return JSON.parse(text)
    } catch (err) {
      return {}
    }
  }).then(data => {
    // Throw an error if one occurred
    if (data.error)
      throw data.error
    
    return data
  })
}

/*
  Other
*/

export const signInGoogle = (state=null) => {
  const clientId = '523323684219-jfakov2bgsleeb6den4ktpohq4lcnae2.apps.googleusercontent.com'
  const redirectUri = 'http://localhost:8080/auth'
  const scope = encodeURIComponent('openid email profile https://www.googleapis.com/auth/calendar.calendarlist.readonly https://www.googleapis.com/auth/calendar.events.readonly')
  
  let stateString = ''
  if (state !== null) {
    state = encodeURIComponent(JSON.stringify(state))
    stateString = `&state=${state}`
  }
  
  window.location.href = `https://accounts.google.com/o/oauth2/v2/auth?client_id=${clientId}&redirect_uri=${redirectUri}&response_type=code&scope=${scope}&access_type=offline${stateString}`
}

export const onLongPress = (element, callback, capture=false) => {
  var timeoutId;

  element.addEventListener('touchstart', function(e) {
      timeoutId = setTimeout(function() {
          timeoutId = null;
          e.stopPropagation();
          callback(e.target);
      }, 500);
  }, capture);

  element.addEventListener('contextmenu', function(e) {
      e.preventDefault();
  }, capture);

  element.addEventListener('touchend', function () {
      if (timeoutId) clearTimeout(timeoutId);
  }, capture);

  element.addEventListener('touchmove', function () {
      if (timeoutId) clearTimeout(timeoutId);
  }, capture);
}

export const isBetween = (value, lower, upper, inclusive=true) => {
  if (inclusive) {
    return value >= lower && value <= upper
  } else {
    return value > lower && value < upper
  }
}

export const clamp = (value, lower, upper) => {
  if (value < lower) return lower
  if (value > upper) return upper
  return value
}