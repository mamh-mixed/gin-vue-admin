import service from '@/utils/request'


export const getVersion = () => {
  return service({
    url: '/docker/getVersion',
    method: 'get'
  })
}

export const getInfo = () => {
    return service({
        url: '/docker/getInfo',
        method: 'get'
    })
}

export const getUsages = () => {
  return service({
    url: '/docker/getUsages',
    method: 'get'
  })
}
