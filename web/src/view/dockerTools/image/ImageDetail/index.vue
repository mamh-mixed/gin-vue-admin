<template>
  <div>
    <el-tabs v-model="activeName" @tab-click="handleClick">
      <el-tab-pane name="first">
        <template #label>
          <span><i class="fa fa-headphones"></i> 基础信息</span>
        </template>
        <BaseInfo/>
      </el-tab-pane>

      <el-tab-pane name="second">
        <template #label>
          <span><i class="fa fa-superpowers"></i> 构建信息</span>
        </template>
        <BuildInfo/>
      </el-tab-pane>

      <el-tab-pane name="third">
        <template #label>
          <span><i class="fa fa-gear"></i> 配置信息</span>
        </template>
        <ConfigInfo/>
      </el-tab-pane>

      <el-tab-pane name="fourth">
        <template #label>
          <span><i class="fa fa-history"></i> 历史信息</span>
        </template>
        <HistoryInfo/>
      </el-tab-pane>

      <el-tab-pane name="five">
        <template #label>
          <span><i class="fa fa-text-width"></i> Json</span>
        </template>
        <JsonInfo/>
      </el-tab-pane>
    </el-tabs>
  </div>
</template>

<script>
import BaseInfo from './BaseInfo'
import BuildInfo from './BuildInfo'
import ConfigInfo from './ConfigInfo'
import HistoryInfo from './HistoryInfo'
import JsonInfo from './JsonInfo'

export default {
  name: 'ImageDetail',
  components: {
    BaseInfo,
    BuildInfo,
    ConfigInfo,
    HistoryInfo,
    JsonInfo
  },

  data() {
    return {
      activeName: 'first'
    }
  },
  methods: {
    handleClick(tab, event) {

    }
  }
}
</script>

<style scoped lang="scss">

</style>
