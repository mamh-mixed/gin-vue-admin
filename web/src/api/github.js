import axios from 'axios'

const service = axios.create()

export function Commits(page) {
  return service({
    url: 'https://api.github.com/repos/mamh-mixed/gin-vue-admin/commits?page=' + page,
    method: 'get'
  })
}

export function Members() {
  return service({
    url: 'https://api.github.com/orgs/mamh-mixed/members',
    method: 'get'
  })
}
