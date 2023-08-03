<template>
  <div>
    <div class="card-header">
      <span>{{ title }}</span>
      <el-tag :type="tagType" size="mini">实时</el-tag>
    </div>

    <div class="card-content">
      {{ content }}
    </div>

    <div class="card-footer">
      {{ footer }}
    </div>


  </div>
</template>

<script>

export default {
  name: 'DetailInfo',
  props: ['title', 'content', 'footer', 'tagType']

}
</script>

<style lang="scss" scoped>
.card-header {
  display: flex;
  justify-content: space-between;
  font-family: "Helvetica Neue", Helvetica, "PingFang SC", "Hiragino Sans GB", "Microsoft YaHei", "微软雅黑", Arial, sans-serif;
  color: #000;
  padding-right: 10px;
}

.card-content {
  font-size: 30px;
  padding: 10px 0px;
}

.card-footer {
  border-top: 1px solid #eee;
  padding: 10px 0px;

}
</style>
