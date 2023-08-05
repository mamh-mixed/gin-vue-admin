<template>
  <div>
    <el-card>
      <template #header>
        <span>基础信息</span>
      </template>

      <el-table :data="tableData" border :show-header="false">
        <el-table-column width="200" prop="infoKey"/>
        <el-table-column width="width" prop="infoValue"/>
      </el-table>

    </el-card>

  </div>
</template>

<script>

export default {
  name: 'BaseInfo',

  data() {
    return {
      tableData: [
        {infoKey: "NAME", infoValue: "nginx:1.25.1"},
        {infoKey: "ID", infoValue: "sha256:021283c8eb95be02b23db0de7f609d603553c6714785e7a673c6594a624ffbda"},
        {infoKey: "RepoTags", infoValue: "nginx:1.25.1,harbor.magesfc.com/library/nginx:1.25.1"},
        {infoKey: "RepoDigests", infoValue: "nginx@sha256:08bc36ad52474e528cc1ea3426b5e3f4bad8a130318e3140d6cfe29c8892c7ef,harbor.magesfc.com/library/nginx@sha256:1bb5c4b86cb7c1e9f0209611dc2135d8a2c1c3a6436163970c99193787d067ea"},
        {infoKey: "Parent", infoValue: ""},
        {infoKey: "Comment", infoValue: ""},
        {infoKey: "CreatAt", infoValue: "2023-07-05 01:24:26"},
        {infoKey: "DockerVersion", infoValue: "20.10.23"},
        {infoKey: "Author", infoValue: ""},
        {infoKey: "Os", infoValue: "linux amd64"},
        {infoKey: "Size", infoValue: "178.20MB"},
        {infoKey: "LastTagTime", infoValue: "2023-07-26 16:46:09"},
      ]
    }
  },
  methods: {}
}
</script>

<style scoped lang="scss">

</style>
