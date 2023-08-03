import service from '@/utils/request'


export const getDockerVersion = () => {
  return service({
    url: '/docker/getVersion',
    method: 'get'
  })
}
