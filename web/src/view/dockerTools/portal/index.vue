<template>
  <div>

    <el-row>
      <el-col :span="18">
        <el-row :gutter="15" class="docker_portal">
          <el-col :span="12">
            <el-card>

              <template #header>
                <h1>节点信息</h1>
              </template>

              <el-row>
                <el-col :span="12">
                  <Detail title="节点名称" content="mamh-left"/>
                </el-col>
                <el-col :span="12">
                  <Detail title="操作系统" content="Ubuntu 18.04.6" footer="5.4.0-150-generic "/>
                </el-col>
              </el-row>

              <el-row>
                <el-col :span="12">
                  <Detail title="Docker版本" content="24.0.2" footer="Docker Engine - Community "/>
                </el-col>
                <el-col :span="12">
                  <Detail title="Build版本" content="go1.20.4" footer="ApiVersion: 1.43 "/>
                </el-col>
              </el-row>

            </el-card>
          </el-col>
          <el-col :span="12">
            <el-card class="card_item">
              <template #header>
                <h1>快捷入口</h1>
              </template>

              <el-row>
                <el-col :span="6">镜像管理
                  <div class="layui-col-md3 layuimini-qiuck-module">
                    <a href="#" onClick="javascript:triggerNavMenuClick('ALL', 'images')"
                       title="镜像管理" data-icon="fa fa-window-maximize">
                      <i class="fa fa-window-maximize"></i>
                      <cite>镜像管理</cite>
                    </a>
                  </div>
                </el-col>
                <el-col :span="6">容器管理
                </el-col>
                <el-col :span="6">数据卷管理
                </el-col>
                <el-col :span="6">仓库管理
                </el-col>
              </el-row>

              <el-row>
                <el-col :span="6">
                  集群信息
                </el-col>
                <el-col :span="6">服务管理
                </el-col>
                <el-col :span="6">配置管理
                </el-col>
                <el-col :span="6">网络信息
                </el-col>
              </el-row>

            </el-card>
          </el-col>
        </el-row>

        <el-row :gutter="15" class="docker_portal">
          <el-col :span="12">
            <el-card class="card_item">
              <h1 style="background-color: #00a38e">333</h1>
            </el-card>
          </el-col>
          <el-col :span="12">
            <el-card class="card_item">
              <h1 style="background-color: #00a38e">444</h1>
            </el-card>
          </el-col>
        </el-row>

      </el-col>

      <el-col :span="6">
        <el-row :gutter="15" class="docker_portal">
          <el-col :span="24">
            <el-card class="right_item">

              <el-table
                  :data="tableData"
                  height="600"
                  style="width: 100%">
                <el-table-column
                    prop="date"
                    label="日期"
                    width="180">
                </el-table-column>
                <el-table-column
                    prop="name"
                    label="姓名"
                    width="180">
                </el-table-column>


              </el-table>
            </el-card>
          </el-col>
        </el-row>
      </el-col>
    </el-row>


  </div>
</template>

<script>
import Detail from './Detail'

export default {
  name: 'Portal',

  components: {
    Detail,
  },


  data() {
    return {

    }
  },
  created() {

  },
  methods: {}
}
</script>

<style scoped lang="scss">
.docker_portal {
  padding: 5px;
}

.card_item {
  height: 300px;
}

.right_item {
  height: 600px;

}
</style>
