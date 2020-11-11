import api from '.'

export function findPrime (number) {
  return api.request('get', `stably/prime/highest/` + number)
}