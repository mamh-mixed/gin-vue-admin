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
              <h1 style="background-color: #00a38e"> 资源使用率 </h1>
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
