<template>
  <div>
    <el-row :gutter="10">
      <el-col :span="2">CPU使用率:</el-col>
      <el-col :span="12">
        <el-progress
            type="line"
            :percentage="cpuUsages"
            :color="colors"
        />
      </el-col>
    </el-row>

    <el-row :gutter="10">
      <el-col :span="2">内存占用率:</el-col>
      <el-col :span="12">
        <el-progress
            type="line"
            :percentage="+memUsages"
            :color="colors"
        />
      </el-col>
    </el-row>

    <el-row :gutter="10">
      <el-col :span="2">磁盘占用率:</el-col>
      <el-col :span="12">
        <el-progress
            type="line"
            :percentage="+diskUsages"
            :color="colors"
        />
      </el-col>
    </el-row>
  </div>
</template>

<script>

import {mapGetters} from 'vuex'

export default {
  name: 'ResourceUsage',
  data() {
    return {
      timer: null,
      diskUsages: 10,
      cpuUsages: 30,
      memUsages: 89,
      colors: [
        {color: '#5cb87a', percentage: 20},
        {color: '#e6a23c', percentage: 40},
        {color: '#f56c6c', percentage: 80}
      ]
    }
  },


  computed: {
    ...mapGetters('docker', ['info', 'version', 'usages']),
    diskUsages() {
      return (this.usages.disk.used * 100 / this.usages.disk.total).toFixed(2)
    },
    cpuUsages() {
      return (+this.usages.cpu[0]).toFixed(2)
    },
    memUsages() {
      return (this.usages.memory.used * 100 / this.usages.memory.total).toFixed(2)
    }
  },
  methods: {}
}
</script>

<style lang="scss" scoped>
</style>
