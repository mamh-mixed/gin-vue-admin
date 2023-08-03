<template>
  <div>

    <el-row>
      <el-col :span="18">
        <el-row :gutter="15" class="docker_portal">
          <el-col :span="12">
            <el-card>

              <template #header>
                <span>节点信息</span>
                <el-switch
                    v-model="value"
                    active-color="#13ce66"
                    inactive-color="#ff4949">
                </el-switch>
              </template>
              <DetailInfo/>
            </el-card>
          </el-col>
          <el-col :span="12">
            <el-card>
              <template #header>
                <h1>快捷入口</h1>
              </template>
              <QuickEntry/>
            </el-card>
          </el-col>
        </el-row>

        <el-row :gutter="15" class="docker_portal">
          <el-col :span="24">
            <el-card>
              <CardInfo/>
            </el-card>
          </el-col>
        </el-row>

        <el-row :gutter="15" class="docker_portal">
          <el-col :span="24">
            <el-card class="card_item">
              <template #header>
                <div >
                  <span class="card-head-icon"><i class="fa fa-table icon"></i></span>
                  <span>资源使用率</span>
                </div>
              </template>
            </el-card>
          </el-col>
        </el-row>
      </el-col>

      <el-col :span="6">
        <el-row :gutter="15" class="docker_portal">
          <el-col :span="24">
            <el-card class="right_item">
              <template #header>
                <h1>概览</h1>
              </template>
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
import DetailInfo from './DetailInfo'
import QuickEntry from './QuickEntry'
import CardInfo from "./CardInfo"

export default {
  name: 'Portal',

  components: {
    CardInfo,
    DetailInfo,
    QuickEntry
  },

  data() {
    return {}
  },
  mounted() {
    this.reload()
  },
  methods: {
    reload() {
      this.$store.dispatch('docker/getInfo')
      this.$store.dispatch('docker/getVersion')
    },
  }
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
