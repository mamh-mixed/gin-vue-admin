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
                  inactive-color="#ff4949"
                />
              </template>
              <DetailInfo />
            </el-card>
          </el-col>
          <el-col :span="12">
            <el-card>
              <template #header>
                <h1>快捷入口</h1>
              </template>
              <QuickEntry />
            </el-card>
          </el-col>
        </el-row>

        <el-row :gutter="15" class="docker_portal">
          <el-col :span="24">
            <el-card>
              <CardInfo />
            </el-card>
          </el-col>
        </el-row>

        <el-row :gutter="15" class="docker_portal">
          <el-col :span="24">
            <el-card class="card_item">
              <template #header>
                <div>
                  <span class="card-head-icon"><i class="fa fa-table icon" /></span>
                  <span>资源使用率</span>
                </div>
              </template>

              <ResourceUsage/>

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

              <OverView />

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
import CardInfo from './CardInfo'
import OverView from './OverView'
import ResourceUsage from './ResourceUsage'

export default {
  name: 'Portal',

  components: {
    CardInfo,
    DetailInfo,
    QuickEntry,
    OverView,
    ResourceUsage
  },

  data() {
    return {
      timer: null,
    }
  },
  created() {
    this.reload()
    this.timer = setInterval(() => {
      this.reload()
    }, 1000 * 30)
  },
  beforeDestroy() {
    clearInterval(this.timer)
    this.timer = null
  },
  methods: {
    reload() {
      this.$store.dispatch('docker/getInfo')
      this.$store.dispatch('docker/getVersion')
      this.$store.dispatch('docker/getUsages')
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
