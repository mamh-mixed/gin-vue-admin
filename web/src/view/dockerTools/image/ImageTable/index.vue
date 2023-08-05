<template>
  <div>

    <el-table :data="tableData" border row-key="ID" stripe style="width: 100%">
      <el-table-column label="序号" width="width" type="index" />
      <el-table-column label="IMAGE ID" width="200" prop="Id" />
      <el-table-column label="REPOSITORY" width="width" prop="Repository" />
      <el-table-column label="TAG" width="200" prop="Tag" />
      <el-table-column label="CREATED" width="100" prop="Created" />
      <el-table-column label="SIZE" width="100" prop="Size" />
      <el-table-column prop="prop" label="操作" width="300">
        <template #default="scope">
          <el-button
              size="small"
              type="warning"
              icon="el-icon-edit"
              @click="handleEdit(scope.$index, scope.row)"
          >查看
          </el-button>
          <el-button
              size="small"
              type="danger"
              icon="el-icon-delete"
              @click="handleDelete(scope.$index, scope.row)"
          >运行
          </el-button>
          <el-button
              size="small"
              type="danger"
              icon="el-icon-delete"
              @click="handleDelete(scope.$index, scope.row)"
          >删除
          </el-button>
          <el-button
              size="small"
              type="danger"
              icon="el-icon-delete"
              @click="handleDelete(scope.$index, scope.row)"
          >标记
          </el-button>
        </template>

      </el-table-column>
    </el-table>

  </div>
</template>

<script>
import {getImages} from '@/api/docker'
import moment from "moment";

export default {
  name: 'ImageTable',

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
      const images = res.data.images

      for(let image of images){
        for(let tag of image.RepoTags){
          let tagArray = tag.split(':');
          let idArray = image.Id.split(':');
          let daytime = moment(parseInt(image.Created)*1000).format('YYYY-MM-DD HH:mm:ss')

          this.tableData.push({
            Id: idArray[1].slice(0, 12),
            Repository: tagArray[0],
            Tag: tagArray[1],
            Created: daytime,
            Size: `${(image.Size/1000/1000).toFixed(0)} MB`,
            Labels:  "",
          })

        }
      }



    },
  }
}
</script>

<style scoped lang="scss">

</style>
