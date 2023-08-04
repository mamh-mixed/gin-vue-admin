import {getInfo, getUsages, getVersion} from '@/api/docker'

export const docker = {
  namespaced: true,
  state: {
    info: {},
    usages: {
      cpu:[0],
      disk:{},
      memory:{}
    },
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
    setUsages(state, usages) {
      state.usages = usages
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
    },

    async getUsages({ commit }) {
      const res = await getUsages()
      if (res.code === 0) {
        commit('setUsages', res.data.usages)
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

    usages(state) {
      return state.usages
    },
  }
}
