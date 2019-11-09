import axios from 'axios';

const promiseWrapper = (url) => {
  return new Promise((resolve, reject) => {
    axios
      .get(url)
      .then(res => resolve(res))
      .catch(err => reject(err))
  })
}

export const login = (name) => promiseWrapper(`http://${window.location.host}/login/${name}`);

export const register = (name, service) => promiseWrapper(`http://${window.location.host}/register/${name}/${service}`);
