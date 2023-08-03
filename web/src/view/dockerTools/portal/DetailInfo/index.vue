<template>
  <div>
    <el-row>
      <el-col :span="12">
        <div class="detail-item">
          <div class="detail-header">
            <span>节点名称</span>
            <el-tag type="success" size="mini">实时</el-tag>
          </div>

          <div class="detail-content">
            mamh-left
          </div>

          <div class="detail-footer">

          </div>
        </div>
      </el-col>
      <el-col :span="12">
        <div class="detail-item">

          <div class="detail-header">
            <span>操作系统</span>
            <el-tag type="success" size="mini">实时</el-tag>
          </div>

          <div class="detail-content">
            Ubuntu 18.04.6 LTS
          </div>

          <div class="detail-footer">
            5.4.0-150-generic
          </div>
        </div>
      </el-col>
    </el-row>

    <el-row>
      <el-col :span="12">
        <div class="detail-item">

          <div class="detail-header">
            <span>Docker版本</span>
            <el-tag type="success" size="mini">实时</el-tag>
          </div>

          <div class="detail-content">
            24.0.2
          </div>

          <div class="detail-footer">
            Docker Engine - Community
          </div>
        </div>
      </el-col>
      <el-col :span="12">
        <div class="detail-item">
          <div class="detail-header">
            <span>Build版本</span>
            <el-tag type="success" size="mini">实时</el-tag>
          </div>

          <div class="detail-content">
            go1.20.4
          </div>

          <div class="detail-footer">
            ApiVersion: 1.43
          </div>
        </div>
      </el-col>
    </el-row>
  </div>
</template>

<script>

export default {
  name: 'DetailInfo',

}
</script>

<style lang="scss" scoped>
.detail-item{
  margin: 5px 5px;
}
.detail-header {
  display: flex;
  justify-content: space-between;
  font-family: "Helvetica Neue", Helvetica, "PingFang SC", "Hiragino Sans GB", "Microsoft YaHei", "微软雅黑", Arial, sans-serif;
  color: #000;
  padding-right: 10px;
}

.detail-content {
  font-size: 30px;
  padding: 10px 0px;
}

.detail-footer {
  border-top: 1px solid #eee;
  padding: 10px 0px;

}
</style>
