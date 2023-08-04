import service from '@/utils/request'


export const getVersion = () => {
  return service({
    url: '/docker/getVersion',
    method: 'get',
    donNotShowLoading: true
  })
}

export const getInfo = () => {
  return service({
    url: '/docker/getInfo',
    method: 'get',
    donNotShowLoading: true
  })
}

export const getUsages = () => {
  return service({
    url: '/docker/getUsages',
    method: 'get',
    donNotShowLoading: true
  })
}
