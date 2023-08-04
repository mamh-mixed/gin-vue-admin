<template>
  <div>
    <table>
      <tbody>
      <tr>
        <td>节点名称</td>
        <td>{{ info.Name }}</td>
      </tr>
      <tr>
        <td>产品</td>
        <td>{{ version.Platform.Name }}</td>
      </tr>
      <tr>
        <td>版本</td>
        <td>{{ version.Version }}</td>
      </tr>
      <tr>
        <td>操作系统</td>
        <td>{{ info.OperatingSystem }}</td>
      </tr>
      <tr>
        <td>系统内核</td>
        <td>{{ info.KernelVersion }}</td>
      </tr>
      <tr>
        <td>Arch</td>
        <td>{{ info.OSType }} {{ info.Architecture }}</td>
      </tr>
      <tr>
        <td>Build</td>
        <td>{{ version.GoVersion }}</td>
      </tr>
      <tr>
        <td>API</td>
        <td>{{ version.ApiVersion }}</td>
      </tr>
      <tr>
        <td>CPU</td>
        <td>{{ info.NCPU }}</td>
      </tr>
      <tr>
        <td>Memory</td>
        <td>{{ info.MemTotal }}</td>
      </tr>
      <tr>
        <td>RootDir</td>
        <td>{{ info.DockerRootDir }}</td>
      </tr>
      <tr>
        <td>仓库镜像</td>
        <td>{{ info.Mirrors }}</td>
      </tr>
      <!--      <tr>-->
      <!--        <td>Swarm</td>-->
      <!--        <td>{{ info.swarm.status }}</td>-->
      <!--      </tr>-->
      <!--      <tr>-->
      <!--        <td>ClusterID</td>-->
      <!--        <td>{{ info.swarm.custerid }}</td>-->
      <!--      </tr>-->
      <!--      <tr>-->
      <!--        <td>NodeID</td>-->
      <!--        <td>{{ info.swarm.nodeid }}</td>-->
      <!--      </tr>-->
      <!--      <tr>-->
      <!--        <td>NODEIP</td>-->
      <!--        <td>{{ info.Swarm.IP }}</td>-->
      <!--      </tr>-->

      <tr>
        <td>DOCKERUI</td>
        <td>
          {{ info.DOCKERUI_MODE === 1 ? '宿主机模式' : '独立运行模式' }}
        </td>
      </tr>

      <tr>
        <td>Gitee</td>
        <td>
          <el-row :gutter="20">
            <el-col :span="12">
              <a href="https://gitee.com/mamh-mixed/awesome-docker"
                 target="_blank">
                <img src="https://gitee.com/mamh-mixed/awesome-docker/badge/star.svg?theme=dark" alt="star">
              </a>
            </el-col>
            <el-col :span="12">
              <a href="https://gitee.com/mamh-mixed/awesome-docker"
                 target="_blank">
                <img src="https://gitee.com/mamh-mixed/awesome-docker/badge/fork.svg?theme=dark" alt="fork">
              </a>
            </el-col>
          </el-row>


        </td>
      </tr>

      <tr>
        <td>Github</td>
        <td>
          <el-row :gutter="20">
            <el-col :span="8">
              <a href="https://github.com/mamh-mixed/gin-vue-admin">
                <img
                    class="dom-center"
                    src="https://img.shields.io/github/watchers/mamh-mixed/gin-vue-admin.svg?label=Watch"
                    alt=""
                >
              </a>
            </el-col>
            <el-col :span="8">
              <a href="https://github.com/mamh-mixed/gin-vue-admin">
                <img
                    class="dom-center"
                    src="https://img.shields.io/github/stars/mamh-mixed/gin-vue-admin.svg?style=social"
                    alt=""
                >
              </a>
            </el-col>
            <el-col :span="8">
              <a href="https://github.com/mamh-mixed/gin-vue-admin">
                <img
                    class="dom-center"
                    src="https://img.shields.io/github/forks/mamh-mixed/gin-vue-admin.svg?label=Fork"
                    alt=""
                >
              </a>
            </el-col>
          </el-row>
        </td>
      </tr>
      </tbody>
    </table>
  </div>
</template>

<script>

import {mapGetters} from "vuex";

export default {
  name: 'OverView',
  computed: {
    ...mapGetters('docker', ['info', 'version']),

  }
}
</script>

<style lang="scss" scoped>
.card-body {
  padding: 5px;
}

.layui-table {
  margin-top: 0 !important;
}

.layui-table {
  margin-top: 0;
  border-left: 5px solid #e2e2e2 !important;
}

.layui-table td, .layui-table th {
  font-size: 12px;
}
</style>
