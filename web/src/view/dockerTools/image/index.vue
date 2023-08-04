<template>
  <div>

    <el-table :data="tableData" border row-key="ID" stripe>
      <el-table-column label="序号" width="100" type="index" />
      <el-table-column label="IMAGE ID" width="width" prop="Id" />
      <el-table-column label="REPOSITORY" width="width" prop="name" />
      <el-table-column label="TAG" width="width" prop="RepoTags" />
      <el-table-column label="CREATED" width="width" prop="Created" />
      <el-table-column label="SIZE" width="width" prop="Size" />
      <el-table-column label="LABELS" width="width" prop="Labels" />
      <el-table-column prop="prop" label="操作" width="width">
        <template #default="scope">
          <el-button
              size="small"
              type="warning"
              icon="el-icon-edit"
              @click="handleEdit(scope.$index, scope.row)"
          >编辑
          </el-button>
          <el-button
              size="small"
              type="danger"
              icon="el-icon-delete"
              @click="handleDelete(scope.$index, scope.row)"
          >删除
          </el-button>
        </template>

      </el-table-column>
    </el-table>

  </div>
</template>

<script>
import {getImages} from '@/api/docker'

export default {
  name: 'ImageView',

  data() {
    return {
      tableData: [],
    }
  },
  created() {
    this.getTableData()
  },

  methods: {
    handleEdit(){

    },
    handleDelete(){

    },
    async getTableData() {
      const res = await  getImages()
      this.tableData = res.data.images
      console.log(res.data.images)
    },
  }
}
</script>

<style scoped lang="scss">

</style>
