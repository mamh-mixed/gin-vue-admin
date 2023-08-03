import { getInfo, getVersion } from '@/api/docker'

export const docker = {
  namespaced: true,
  state: {
    info: {},
    version: {
      Platform:{
        Name:'null'
      }
    }
  },
  mutations: {
    setInfo(state, info) {
      state.info = info
    },
    setVersion(state, version) {
      state.version = version
    },
  },
  actions: {
    async getInfo({ commit }) {
      const res = await getInfo()

      if (res.code === 0) {
        commit('setInfo', res.data.info)
      }
      return res
    },

    async getVersion({ commit }) {
      const res = await getVersion()
      if (res.code === 0) {
        commit('setVersion', res.data.version)
      }
      return res
    }
  },
  getters: {
    info(state) {
      return state.info
    },
    version(state) {
      return state.version
    },
  }
}
