<template>
  <div>
    <div class="gva-search-box">
      <el-form
          ref="elSearchFormRef"
          :inline="true"
          :model="searchInfo"
          class="demo-form-inline"
          :rules="searchRule"
          @keyup.enter="onSubmit"
      >
        <el-form-item
            label="创建日期"
            prop="createdAt"
        >
          <template #label>
            <span>
              创建日期
              <el-tooltip content="搜索范围是开始日期（包含）至结束日期（不包含）">
                <el-icon><QuestionFilled /></el-icon>
              </el-tooltip>
            </span>
          </template>
          <el-date-picker
              v-model="searchInfo.startCreatedAt"
              type="datetime"
              placeholder="开始日期"
              :disabled-date="time=> searchInfo.endCreatedAt ? time.getTime() > searchInfo.endCreatedAt.getTime() : false"
          />
          —
          <el-date-picker
              v-model="searchInfo.endCreatedAt"
              type="datetime"
              placeholder="结束日期"
              :disabled-date="time=> searchInfo.startCreatedAt ? time.getTime() < searchInfo.startCreatedAt.getTime() : false"
          />
        </el-form-item>
        <el-form-item
            label="租户名"
            prop="username"
        >
          <el-input
              v-model="searchInfo.username"
              placeholder="搜索条件"
          />

        </el-form-item>
        <el-form-item
            label="租户昵称"
            prop="nickname"
        >
          <el-input
              v-model="searchInfo.nickname"
              placeholder="搜索条件"
          />

        </el-form-item>
        <el-form-item
            label="有效期"
            prop="endTime"
        >

          <template #label>
            <span>
              有效期
              <el-tooltip content="搜索范围是开始日期（包含）至结束日期（不包含）">
                <el-icon><QuestionFilled /></el-icon>
              </el-tooltip>
            </span>
          </template>
          <el-date-picker
              v-model="searchInfo.startEndTime"
              type="datetime"
              placeholder="开始日期"
              :disabled-date="time=> searchInfo.endEndTime ? time.getTime() > searchInfo.endEndTime.getTime() : false"
          />
          —
          <el-date-picker
              v-model="searchInfo.endEndTime"
              type="datetime"
              placeholder="结束日期"
              :disabled-date="time=> searchInfo.startEndTime ? time.getTime() < searchInfo.startEndTime.getTime() : false"
          />

        </el-form-item>
        <el-form-item
            label="分润模式"
            prop="allocation"
        >
          <el-select
              v-model="searchInfo.allocation"
              clearable
              placeholder="请选择"
              @clear="()=>{searchInfo.allocation=undefined}"
          >
            <el-option
                v-for="(item,key) in allocationOptions"
                :key="key"
                :label="item.label"
                :value="item.value"
            />
          </el-select>
        </el-form-item>
        <el-form-item>
          <el-button
              type="primary"
              icon="search"
              @click="onSubmit"
          >查询</el-button>
          <el-button
              icon="refresh"
              @click="onReset"
          >重置</el-button>
        </el-form-item>
      </el-form>
    </div>
    <div class="gva-table-box">
      <div class="gva-btn-list">
        <el-button
            type="primary"
            icon="plus"
            @click="openDialog"
        >新增</el-button>
        <el-popover
            v-model:visible="deleteVisible"
            placement="top"
            width="160"
        >
          <p>确定要删除吗？</p>
          <div style="text-align: right; margin-top: 8px;">
            <el-button
                type="primary"
                link
                @click="deleteVisible = false"
            >取消</el-button>
            <el-button
                type="primary"
                @click="onDelete"
            >确定</el-button>
          </div>
          <template #reference>
            <el-button
                icon="delete"
                style="margin-left: 10px;"
                :disabled="!multipleSelection.length"
                @click="deleteVisible = true"
            >删除</el-button>
          </template>
        </el-popover>
      </div>
      <el-table
          ref="multipleTable"
          style="width: 100%"
          tooltip-effect="dark"
          :data="tableData"
          row-key="ID"
          @selection-change="handleSelectionChange"
      >
        <el-table-column
            type="selection"
            width="55"
        />
        <el-table-column
            align="left"
            label="日期"
            width="180"
        >
          <template #default="scope">{{ formatDate(scope.row.CreatedAt) }}</template>
        </el-table-column>
        <el-table-column
            align="left"
            label="租户名"
            prop="username"
            width="120"
        />

        <el-table-column
            align="left"
            label="租户昵称"
            prop="nickname"
            width="120"
        />
        <el-table-column
            label="租户logo"
            width="200"
        >
          <template #default="scope">
            <el-image
                style="width: 100px; height: 100px"
                :src="getUrl(scope.row.logo)"
                fit="cover"
            />
          </template>
        </el-table-column>
        <el-table-column
            align="left"
            label="有效期"
            width="180"
        >
          <template #default="scope">{{ formatDate(scope.row.endTime) }}</template>
        </el-table-column>
        <el-table-column
            align="left"
            label="分润模式"
            prop="allocation"
            width="120"
        >
          <template #default="scope">
            {{ filterDict(scope.row.allocation,allocationOptions) }}
          </template>
        </el-table-column>
        <el-table-column
            align="left"
            label="分润比例"
            prop="allocationProportion"
            width="120"
        />
        <el-table-column align="left" label="登录连接" width="180">
          <template #default="scope">{{ getLoginUrl(scope.row)  }}</template>
        </el-table-column>
        <el-table-column
            align="left"
            label="操作"
        >
          <template #default="scope">
            <el-button
                type="primary"
                link
                icon="edit"
                class="table-button"
                @click="opdendrawer(scope.row)"
            >分配权限</el-button>
            <el-button type="primary" link icon="edit" class="table-button" @click="createAdmin(scope.row.ID)">创建管理员</el-button>
            <el-button
                type="primary"
                link
                icon="edit"
                class="table-button"
                @click="updateCsTenantFunc(scope.row)"
            >变更</el-button>
            <el-button
                type="primary"
                link
                icon="delete"
                @click="deleteRow(scope.row)"
            >删除</el-button>
          </template>
        </el-table-column>
      </el-table>
      <div class="gva-pagination">
        <el-pagination
            layout="total, sizes, prev, pager, next, jumper"
            :current-page="page"
            :page-size="pageSize"
            :page-sizes="[10, 30, 50, 100]"
            :total="total"
            @current-change="handleCurrentChange"
            @size-change="handleSizeChange"
        />
      </div>
    </div>
    <el-dialog
        v-model="dialogFormVisible"
        :before-close="closeDialog"
        :title="type==='create'?'添加':'修改'"
        destroy-on-close
    >
      <el-scrollbar height="500px">
        <el-form
            ref="elFormRef"
            :model="formData"
            label-position="right"
            :rules="rule"
            label-width="80px"
        >
          <el-form-item
              label="租户名:"
              prop="username"
          >
            <el-input
                v-model="formData.username"
                :clearable="true"
                placeholder="请输入"
            />
          </el-form-item>

          <el-form-item
              label="租户昵称:"
              prop="nickname"
          >
            <el-input
                v-model="formData.nickname"
                :clearable="true"
                placeholder="请输入"
            />
          </el-form-item>
          <el-form-item
              label="租户logo:"
              prop="logo"
          >
            <SelectImage v-model="formData.logo" />
          </el-form-item>
          <el-form-item
              label="有效期:"
              prop="endTime"
          >
            <el-date-picker
                v-model="formData.endTime"
                type="date"
                style="width:100%"
                placeholder="选择日期"
                :clearable="true"
            />
          </el-form-item>
          <el-form-item
              label="分润模式:"
              prop="allocation"
          >
            <el-select
                v-model="formData.allocation"
                placeholder="请选择"
                style="width:100%"
                :clearable="true"
            >
              <el-option
                  v-for="(item,key) in allocationOptions"
                  :key="key"
                  :label="item.label"
                  :value="item.value"
              />
            </el-select>
          </el-form-item>
          <el-form-item
              label="分润比例:"
              prop="allocationProportion"
          >
            <el-input-number
                v-model="formData.allocationProportion"
                style="width:100%"
                :precision="2"
                :clearable="true"
            />
          </el-form-item>
        </el-form>
      </el-scrollbar>
      <template #footer>
        <div class="dialog-footer">
          <el-button @click="closeDialog">取 消</el-button>
          <el-button
              type="primary"
              @click="enterDialog"
          >确 定</el-button>
        </div>
      </template>
    </el-dialog>

    <el-drawer
        v-if="drawer"
        v-model="drawer"
        custom-class="auth-drawer"
        :with-header="false"
        size="40%"
        title="角色配置"
    >
      <el-tabs
          type="border-card"
      >
        <el-tab-pane label="角色菜单">
          <Menus
              ref="menus"
              :row="activeRow"
              @changeRow="changeRow"
          />
        </el-tab-pane>
        <el-tab-pane label="角色api">
          <Apis
              ref="apis"
              :row="activeRow"
              @changeRow="changeRow"
          />
        </el-tab-pane>
      </el-tabs>
    </el-drawer>
  </div>
</template>

<script>
export default {
  name: 'CsTenant'
}
</script>

<script setup>
import {
  createCsTenant,
  deleteCsTenant,
  deleteCsTenantByIds,
  updateCsTenant,
  findCsTenant,
  getCsTenantList,
  createCsTenantAdmin
} from '@/api/csTenant'
import { getUrl } from '@/utils/image'
// 图片选择组件
import SelectImage from '@/components/selectImage/selectImage.vue'

// 全量引入格式化工具 请按需保留
import { getDictFunc, formatDate, formatBoolean, filterDict, ReturnArrImg, onDownloadFile } from '@/utils/format'
import { ElMessage, ElMessageBox } from 'element-plus'
import { ref, reactive } from 'vue'
import Menus from '@/view/csTenant/components/menus.vue'
import Apis from '@/view/csTenant/components/apis.vue'

const activeRow = ref({})
const changeRow = (key, value) => {
  activeRow.value[key] = value
}

const getLoginUrl = (row) => {
  return `${window.location.origin}/#/login?tenant=${row.onlyKey}`
}

const createAdmin = async (tenantID) =>{
  const res = await createCsTenantAdmin(tenantID)
  if (res.code === 0) {
    ElMessage({
      type: 'success',
      message: '创建成功'
    })
  }
}

const drawer = ref(false)

const opdendrawer = (row) => {
  drawer.value = true
  activeRow.value = row
}

// 自动化生成的字典（可能为空）以及字段
const allocationOptions = ref([])
const formData = ref({
  username: '',
  nickname: '',
  logo: '',
  endTime: new Date(),
  allocation: undefined,
  allocationProportion: 0,
})

// 验证规则
const rule = reactive({
  username: [{
    required: true,
    message: '',
    trigger: ['input', 'blur'],
  },
    {
      whitespace: true,
      message: '不能只输入空格',
      trigger: ['input', 'blur'],
    }
  ],

})

const searchRule = reactive({
  createdAt: [
    { validator: (rule, value, callback) => {
        if (searchInfo.value.startCreatedAt && !searchInfo.value.endCreatedAt) {
          callback(new Error('请填写结束日期'))
        } else if (!searchInfo.value.startCreatedAt && searchInfo.value.endCreatedAt) {
          callback(new Error('请填写开始日期'))
        } else if (searchInfo.value.startCreatedAt && searchInfo.value.endCreatedAt && (searchInfo.value.startCreatedAt.getTime() === searchInfo.value.endCreatedAt.getTime() || searchInfo.value.startCreatedAt.getTime() > searchInfo.value.endCreatedAt.getTime())) {
          callback(new Error('开始日期应当早于结束日期'))
        } else {
          callback()
        }
      }, trigger: 'change' }
  ],
  endTime: [{ validator: (rule, value, callback) => {
      if (searchInfo.value.startEndTime && !searchInfo.value.endEndTime) {
        callback(new Error('请填写结束日期'))
      } else if (!searchInfo.value.startEndTime && searchInfo.value.endEndTime) {
        callback(new Error('请填写开始日期'))
      } else if (searchInfo.value.startEndTime && searchInfo.value.endEndTime && (searchInfo.value.startEndTime.getTime() === searchInfo.value.endEndTime.getTime() || searchInfo.value.startEndTime.getTime() > searchInfo.value.endEndTime.getTime())) {
        callback(new Error('开始日期应当早于结束日期'))
      } else {
        callback()
      }
    }, trigger: 'change' }],
})

const elFormRef = ref()
const elSearchFormRef = ref()

// =========== 表格控制部分 ===========
const page = ref(1)
const total = ref(0)
const pageSize = ref(10)
const tableData = ref([])
const searchInfo = ref({})

// 重置
const onReset = () => {
  searchInfo.value = {}
  getTableData()
}

// 搜索
const onSubmit = () => {
  elSearchFormRef.value?.validate(async(valid) => {
    if (!valid) return
    page.value = 1
    pageSize.value = 10
    getTableData()
  })
}

// 分页
const handleSizeChange = (val) => {
  pageSize.value = val
  getTableData()
}

// 修改页面容量
const handleCurrentChange = (val) => {
  page.value = val
  getTableData()
}

// 查询
const getTableData = async() => {
  const table = await getCsTenantList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
  if (table.code === 0) {
    tableData.value = table.data.list
    total.value = table.data.total
    page.value = table.data.page
    pageSize.value = table.data.pageSize
  }
}

getTableData()

// ============== 表格控制部分结束 ===============

// 获取需要的字典 可能为空 按需保留
const setOptions = async() => {
  allocationOptions.value = await getDictFunc('allocation')
}

// 获取需要的字典 可能为空 按需保留
setOptions()

// 多选数据
const multipleSelection = ref([])
// 多选
const handleSelectionChange = (val) => {
  multipleSelection.value = val
}

// 删除行
const deleteRow = (row) => {
  ElMessageBox.confirm('确定要删除吗?', '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning'
  }).then(() => {
    deleteCsTenantFunc(row)
  })
}

// 批量删除控制标记
const deleteVisible = ref(false)

// 多选删除
const onDelete = async() => {
  const ids = []
  if (multipleSelection.value.length === 0) {
    ElMessage({
      type: 'warning',
      message: '请选择要删除的数据'
    })
    return
  }
  multipleSelection.value &&
  multipleSelection.value.map(item => {
    ids.push(item.ID)
  })
  const res = await deleteCsTenantByIds({ ids })
  if (res.code === 0) {
    ElMessage({
      type: 'success',
      message: '删除成功'
    })
    if (tableData.value.length === ids.length && page.value > 1) {
      page.value--
    }
    deleteVisible.value = false
    getTableData()
  }
}

// 行为控制标记（弹窗内部需要增还是改）
const type = ref('')

// 更新行
const updateCsTenantFunc = async(row) => {
  const res = await findCsTenant({ ID: row.ID })
  type.value = 'update'
  if (res.code === 0) {
    formData.value = res.data.recsTenant
    dialogFormVisible.value = true
  }
}

// 删除行
const deleteCsTenantFunc = async(row) => {
  const res = await deleteCsTenant({ ID: row.ID })
  if (res.code === 0) {
    ElMessage({
      type: 'success',
      message: '删除成功'
    })
    if (tableData.value.length === 1 && page.value > 1) {
      page.value--
    }
    getTableData()
  }
}

// 弹窗控制标记
const dialogFormVisible = ref(false)

// 查看详情控制标记
const detailShow = ref(false)

// 打开详情弹窗
const openDetailShow = () => {
  detailShow.value = true
}

// 打开详情
const getDetails = async(row) => {
  // 打开弹窗
  const res = await findCsTenant({ ID: row.ID })
  if (res.code === 0) {
    formData.value = res.data.recsTenant
    openDetailShow()
  }
}

// 关闭详情弹窗
const closeDetailShow = () => {
  detailShow.value = false
  formData.value = {
    username: '',
    nickname: '',
    endTime: new Date(),
    allocation: undefined,
    allocationProportion: 0,
  }
}

// 打开弹窗
const openDialog = () => {
  type.value = 'create'
  dialogFormVisible.value = true
}

// 关闭弹窗
const closeDialog = () => {
  dialogFormVisible.value = false
  formData.value = {
    username: '',
    nickname: '',
    endTime: new Date(),
    allocation: undefined,
    allocationProportion: 0,
  }
}
// 弹窗确定
const enterDialog = async() => {
  elFormRef.value?.validate(async(valid) => {
    if (!valid) return
    let res
    switch (type.value) {
      case 'create':
        res = await createCsTenant(formData.value)
        break
      case 'update':
        res = await updateCsTenant(formData.value)
        break
      default:
        res = await createCsTenant(formData.value)
        break
    }
    if (res.code === 0) {
      ElMessage({
        type: 'success',
        message: '创建/更改成功'
      })
      closeDialog()
      getTableData()
    }
  })
}

</script>

<style>

</style>
