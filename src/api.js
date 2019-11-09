import axios from 'axios';

const promiseWrapper = (url) => {
  return new Promise((resolve, reject) => {
    axios
      .get(url)
      .then(res => resolve(res))
      .catch(err => reject(err))
  })
}

export const authenticateApp = (url) => promiseWrapper(url)
