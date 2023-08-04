<template>
  <div>
    <el-row :gutter="10">
      <el-col :span="2">CPU使用率:</el-col>
      <el-col :span="12">
        <el-progress
            type="line"
            :percentage="+cpuUsages.toFixed(0)"
            :color="colors"
        />
      </el-col>
    </el-row>

    <el-row :gutter="10">
      <el-col :span="2">内存占用率:</el-col>
      <el-col :span="12">
        <el-progress
            type="line"
            :percentage="+memUsages.toFixed(0)"
            :color="colors"
        />
      </el-col>
    </el-row>

    <el-row :gutter="10">
      <el-col :span="2">磁盘占用率:</el-col>
      <el-col :span="12">
        <el-progress
            type="line"
            :percentage="+diskUsages.toFixed(0)"
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
    ...mapGetters('docker', ['info', 'version']),

  }
}
</script>

<style lang="scss" scoped>
</style>
