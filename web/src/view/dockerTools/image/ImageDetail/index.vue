<template>
  <div>
    <el-tabs v-model="activeName" @tab-click="handleClick">
      <el-tab-pane name="first">
        <template #label>
          <span><i class="fa fa-headphones"></i> 基础信息</span>
        </template>
        镜像信息
      </el-tab-pane>

      <el-tab-pane name="second">
        <template #label>
          <span><i class="fa fa-superpowers"></i> 构建信息</span>
        </template>
        构建信息
      </el-tab-pane>
      <el-tab-pane name="third">
        <template #label>
          <span><i class="fa fa-gear"></i> 配置信息</span>
        </template>
        配置信息
      </el-tab-pane>
      <el-tab-pane name="fourth">
        <template #label>
          <span><i class="fa fa-history"></i> 历史信息</span>
        </template>
        历史信息
      </el-tab-pane>
      <el-tab-pane name="five">
        <template #label>
          <span><i class="fa fa-text-width"></i> Json</span>
        </template>

      </el-tab-pane>
    </el-tabs>
  </div>
</template>

<script>

export default {
  name: 'ImageDetail',

  data() {
    return {
      activeName: 'first'
    }
  },
  methods: {
    handleClick(tab, event) {
      console.log(tab, event);
    }
  }
}
</script>

<style scoped lang="scss">

</style>
