<template>
  <div >
    <el-row>
      <el-col :span="6" >
        <div class="layuimini-qiuck-module">
          <a href="#" title="镜像管理">
            <i class="fa fa-window-maximize"></i>
            <span>镜像管理</span>
          </a>
        </div>
      </el-col>
      <el-col :span="6">
        <div class="layuimini-qiuck-module">
          <a href="#" title="容器管理">
            <i class="fa fa-gears"></i>
            <span>容器管理</span>
          </a>
        </div>
      </el-col>
      <el-col :span="6">
        <div class="layuimini-qiuck-module">
          <a href="#" title="数据卷管理">
            <i class="fa fa-file-text"></i>
            <span>数据卷管理</span>
          </a>
        </div>
      </el-col>
      <el-col :span="6">
        <div class="layuimini-qiuck-module">
          <a href="#" title="仓库管理">
            <i class="fa fa-dot-circle-o"></i>
            <span>仓库管理</span>
          </a>
        </div>
      </el-col>
    </el-row>

    <el-row>
      <el-col :span="6">
        <div class="layuimini-qiuck-module">
          <a href="#" title="集群信息">
            <i class="fa fa-calendar"></i>
            <span>集群信息</span>
          </a>
        </div>
      </el-col>
      <el-col :span="6">
        <div class="layuimini-qiuck-module">
          <a href="#" title="服务管理">
            <i class="fa fa-hourglass-end"></i>
            <span>服务管理</span>
          </a>
        </div>
      </el-col>
      <el-col :span="6">
        <div class="layuimini-qiuck-module">
          <a href="#" title="配置管理">
            <i class="fa fa-search"></i>
            <span>配置管理</span>
          </a>
        </div>
      </el-col>
      <el-col :span="6">
        <div class="layuimini-qiuck-module">
          <a href="#" title="网络信息">
            <i class="fa fa-usb"></i>
            <span>网络信息</span>
          </a>
        </div>
      </el-col>
    </el-row>
  </div>
</template>

<script>

export default {
  name: 'QuickEntry',
}
</script>

<style lang="scss" scoped>
.layuimini-qiuck-module a {
  color: #333;
  text-decoration: none;
}

.layuimini-qiuck-module {
  text-align: center;
  margin: 5px 5px;
}

.layuimini-qiuck-module a i {
  display: inline-block;
  width: 100%;
  height: 90px !important;
  line-height: 80px !important;
  text-align: center;
  border-radius: 2px;
  font-size: 40px !important;
  background-color: #F8F8F8;
  color: #333;
  transition: all .3s;
  -webkit-transition: all .3s;
}

.layuimini-qiuck-module a span {
  position: relative;
  top: 8px;
  display: block;
  color: #666;
  text-overflow: ellipsis;
  overflow: hidden;
  white-space: nowrap;
  font-size: 14px;
}

</style>
