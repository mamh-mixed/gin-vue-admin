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
            {{ info.Name }}
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
            {{ info.OperatingSystem }}
          </div>

          <div class="detail-footer">
            {{ info.KernelVersion }}
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
            {{ info.ServerVersion }}
          </div>

          <div class="detail-footer">
            {{ version.Platform.Name }}
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
            {{ version.GoVersion }}
          </div>

          <div class="detail-footer">
            {{ version.ApiVersion }}
          </div>
        </div>
      </el-col>
    </el-row>
  </div>
</template>

<script>

import {mapGetters} from "vuex";

export default {
  name: 'DetailInfo',
  computed: {
    ...mapGetters('docker', ['info', 'version']),

  }
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
