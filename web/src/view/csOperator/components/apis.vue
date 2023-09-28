<template>
  <div>
    <div class="sticky top-0.5 z-10 bg-white">
      <el-input
        v-model="filterText"
        class="w-3/5"
        placeholder="筛选"
      />
      <el-button
        class="float-right"
        type="primary"
        @click="authApiEnter"
      >确 定</el-button>
    </div>
    <div class="tree-content">
      <el-tree
        ref="apiTree"
        :data="apiTreeData"
        :default-checked-keys="apiTreeIds"
        :props="apiDefaultProps"
        default-expand-all
        highlight-current
        node-key="ID"
        show-checkbox
        :filter-node-method="filterNode"
        @check="nodeChange"
      />
    </div>
  </div>
</template>

<script setup>
import { getAllApis } from '@/api/api'
import { getApisByOperatorID,setOperatorApis } from '@/api/csOperator'
import { ref, watch } from 'vue'
import { ElMessage } from 'element-plus'

defineOptions({
  name: 'Apis',
})

const props = defineProps({
  row: {
    default: function() {
      return {}
    },
    type: Object
  }
})

const apiDefaultProps = ref({
  children: 'children',
  label: 'description'
})
const filterText = ref('')
const apiTreeData = ref([])
const apiTreeIds = ref([])
const activeOperator = ref('')
const init = async() => {
  const res2 = await getAllApis()
  const apis = res2.data.apis
  apiTreeData.value = buildApiTree(apis)
  const res = await getApisByOperatorID({ operatorID: props.row.ID })
  activeOperator.value = props.row.ID
  apiTreeIds.value = []
  res.data && res.data.forEach(item => {
    apiTreeIds.value.push(item.apiId)
  })
}

init()

const needConfirm = ref(false)
const nodeChange = () => {
  needConfirm.value = true
}
// 暴露给外层使用的切换拦截统一方法
const enterAndNext = () => {
  authApiEnter()
}

// 创建api树方法
const buildApiTree = (apis) => {
  const apiObj = {}
  apis &&
        apis.forEach(item => {
          if (Object.prototype.hasOwnProperty.call(apiObj, item.apiGroup)) {
            apiObj[item.apiGroup].push(item)
          } else {
            Object.assign(apiObj, { [item.apiGroup]: [item] })
          }
        })
  const apiTree = []
  for (const key in apiObj) {
    const treeNode = {
      ID: apiObj[key].ID,
      description: key + '组',
      children: apiObj[key]
    }
    apiTree.push(treeNode)
  }
  return apiTree
}

// 关联关系确定
const apiTree = ref(null)
const authApiEnter = async() => {
  const checkArr = apiTree.value.getCheckedNodes(true)
  var apiIds = []

  checkArr.forEach(item => {
    if (item.ID) {
      apiIds.push(item.ID)
    }
  })
  const res = await setOperatorApis({
    operatorID: props.row.ID,
    apiIds
  })
  if (res.code === 0) {
    ElMessage({ type: 'success', message: 'api设置成功' })
  }
}

defineExpose({
  needConfirm,
  enterAndNext
})

const filterNode = (value, data) => {
  if (!value) return true
  return data.description.indexOf(value) !== -1
}
watch(filterText, (val) => {
  apiTree.value.filter(val)
})

</script>
