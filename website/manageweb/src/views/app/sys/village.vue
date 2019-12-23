<template>
  <div class="app-container">
    <div class="filter-container">
      <el-input
        v-model="listQuery.id"
        placeholder="ID"
        clearable
        prefix-icon="el-icon-search"
        style="width: 90px;"
        class="filter-item"
        @keyup.enter.native="handleFilter"
        @clear="handleFilter"
      />
      <el-input
        v-model="listQuery.village_name"
        placeholder="小区名称"
        clearable
        prefix-icon="el-icon-search"
        style="width: 90px;"
        class="filter-item"
        @keyup.enter.native="handleFilter"
        @clear="handleFilter"
      />
      <el-input
        v-model="listQuery.village_at_province"
        placeholder="所在省"
        clearable
        prefix-icon="el-icon-search"
        style="width: 200px;"
        class="filter-item"
        @keyup.enter.native="handleFilter"
        @clear="handleFilter"
      />
      <el-input
        v-model="listQuery.village_at_city"
        placeholder="所在市"
        clearable
        prefix-icon="el-icon-search"
        style="width: 200px;"
        class="filter-item"
        @keyup.enter.native="handleFilter"
        @clear="handleFilter"
      />
      <el-input
        v-model="listQuery.village_at_district"
        placeholder="所在区"
        clearable
        prefix-icon="el-icon-search"
        style="width: 200px;"
        class="filter-item"
        @keyup.enter.native="handleFilter"
        @clear="handleFilter"
      />
      <el-select
        v-model="listQuery.village_id"
        placeholder="卡类型"
        clearable
        style="width: 90px"
        class="filter-item"
        @change="handleFilter"
      >
        <el-option
          v-for="item in menuTypeOptions"
          :key="item.key"
          :label="item.display_name"
          :value="item.key"
        />
      </el-select>

      <el-select
        v-model="listQuery.sort"
        style="width: 140px"
        class="filter-item"
        @change="handleFilter"
      >
        <el-option
          v-for="item in sortOptions"
          :key="item.key"
          :label="item.label"
          :value="item.key"
        />
      </el-select>
      <el-button
        v-waves
        class="filter-item"
        type="primary"
        icon="el-icon-search"
        @click="handleFilter"
      >
        {{ "搜索" }}
      </el-button>
      <el-button
        v-if="CardList.add"
        class="filter-item"
        style="margin-left: 10px;"
        type="primary"
        icon="el-icon-edit"
        @click="handleCreate"
      >
        {{ "添加" }}
      </el-button>
      <el-button
        v-if="CardList.del"
        class="filter-item"
        type="danger"
        icon="el-icon-delete"
        @click="handleBatchDel"
      >
        {{ "删除" }}
      </el-button>
    </div>

    <el-table
      :key="tableKey"
      v-loading="listLoading"
      stripe
      :data="list"
      border
      fit
      highlight-current-row
      style="width: 100%;"
      @sort-change="sortChange"
      @selection-change="handleSelectionChange"
    >
      <el-table-column type="selection" width="55" />
      <el-table-column label="ID" align="center">
        <template slot-scope="scope">
          <span>{{ scope.row.id }}</span>
        </template>
      </el-table-column>
      <el-table-column label="小区名称" align="center">
        <template slot-scope="scope">
          <span>{{ scope.row.village_name }}</span>
        </template>
      </el-table-column>
      <el-table-column label="单元号" align="center">
        <template slot-scope="scope">
          <span>{{ scope.row.village_id }}</span>
        </template>
      </el-table-column>
      <el-table-column label="楼号" align="center">
        <template slot-scope="scope">
          <span>{{ scope.row.village_apartment_id }}</span>
        </template>
      </el-table-column>
      <el-table-column label="所在区" align="center">
        <template slot-scope="scope">
          <span>{{ scope.row.village_at_district }}</span>
        </template>
      </el-table-column>
      <el-table-column label="所在市" align="center">
        <template slot-scope="scope">
          <span>{{ scope.row.village_at_city }}</span>
        </template>
      </el-table-column>
      <el-table-column label="所在省" align="center">
        <template slot-scope="scope">
          <span>{{ scope.row.village_at_province }}</span>
        </template>
      </el-table-column>
      <el-table-column
        label="操作"
        align="center"
        width="230"
        class-name="small-padding fixed-width"
      >
        <template slot-scope="{ row }">
          <el-button
            v-if="CardList.view"
            size="mini"
            type="success"
            @click="handleDetail(row.id)"
          >
            {{ "查看" }}
          </el-button>
          <el-button
            v-if="CardList.update"
            type="primary"
            size="mini"
            @click="handleUpdate(row.id)"
          >
            {{ "编辑" }}
          </el-button>
          <el-button
            v-if="CardList.del"
            size="mini"
            type="danger"
            @click="handleDelete(row)"
          >
            {{ "删除" }}
          </el-button>
        </template>
      </el-table-column>
    </el-table>

    <pagination
      v-show="total > 0"
      :total="total"
      :page.sync="listQuery.page"
      :limit.sync="listQuery.limit"
      @pagination="getList"
    />
    <!--接下来是小窗口，添加界面-->
    <el-dialog :title="textMap[dialogStatus]" :visible.sync="dialogFormVisible">
      <el-form
        ref="dataForm"
        v-loading="loading"
        element-loading-text="正在执行"
        element-loading-background="rgba(255,255,255,0.7)"
        :rules="rules"
        :model="temp"
        label-position="left"
        label-width="80px"
        style="width: 400px; margin-left:50px;"
      >
        <el-form-item label="小区名称" prop="village_name">
          <el-input v-model="temp.village_name" />
        </el-form-item>
        <el-form-item label="楼号" prop="village_id">
          <el-input v-model="temp.village_id" />
        </el-form-item>

        <el-form-item label="单元号" prop="village_apartment_id">
          <el-input v-model="temp.village_apartment_id" />
        </el-form-item>

        <el-form-item label="所在省" prop="village_at_province">
          <el-input v-model="temp.village_at_province" />
        </el-form-item>

        <el-form-item label="所在市" prop="village_at_city">
          <el-input v-model="temp.village_at_city" />
        </el-form-item>
        <el-form-item label="所在区" prop="village_at_district">
          <el-input v-model="temp.village_at_district" />
        </el-form-item>
        <el-form-item label="具体地址" prop="village_at_address">
          <el-input v-model="temp.village_at_address" />
        </el-form-item>
      </el-form>
      <div
        v-if="
          dialogStatus !== 'detail' ? (loading === true ? false : true) : false
        "
        slot="footer"
        class="dialog-footer"
      >
        <el-button @click="dialogFormVisible = false">
          {{ "取消" }}
        </el-button>
        <el-button
          type="primary"
          @click="dialogStatus === 'create' ? createData() : updateData()"
        >
          {{ "确定" }}
        </el-button>
      </div>
    </el-dialog>
  </div>
</template>

<script>
import { requestList, requestDetail, requestUpdate, requestCreate, requestAll, requestDelete, requestMenuButton } from '@/api/app/sys/village'
import waves from '@/directive/waves'
import Pagination from '@/components/Pagination' // Secondary package based on el-pagination
// import SelectTree from '@/components/TreeSelect'
import { checkAuthAdd, checkAuthDel, checkAuthView, checkAuthUpdate } from '@/utils/permission'

const menuTypeOptions = [
  { key: 1, display_name: 'ISO14443A' },
  { key: 2, display_name: 'ISO14443B' },
  { key: 3, display_name: 'LOWFREQ_125K' }
]

export default {
  name: 'Village',
  components: { Pagination }, // SelectTree
  directives: { waves },
  data() {
    return {
      valueIdSelectTree: 0,
      valueIdSelectTree2: 0,
      propsSelectlist: [],
      operationList: [],
      CardList: {
        add: false,
        del: false,
        view: false,
        update: false
      },
      tableKey: 0,
      list: [],
      total: 0,
      listLoading: true,
      loading: true,
      listQuery: {
        page: 1,
        limit: 20,
        id: undefined,
        key: undefined,
        village_name: undefined,
        village_at_province: undefined,
        village_at_city: undefined,
        village_at_district: undefined,
        village_id: undefined,
        village_apartment_id: undefined,
        village_at_address: undefined,
        sort: '-id'
      },
      menuTypeOptions,
      sortOptions: [
        { label: 'ID Ascending', key: '+id' },
        { label: 'ID Descending', key: '-id' },
        { label: 'card_id Ascending', key: '+card_id' },
        { label: 'card_id Descending', key: '-card_id' }
      ],
      temp: {
        id: 0,
        memo: '',
        village_name: '',
        village_at_province: '',
        village_at_city: '',
        village_at_district: '',
        village_id: '',
        village_apartment_id: '',
        village_at_address: ''
      },
      dialogFormVisible: false,
      dialogStatus: '',
      textMap: {
        update: '编辑',
        create: '添加',
        detail: '详情'
      },
      rules: {
        type: [{ required: true, message: '请选择卡片类型', trigger: 'change' }],
        name: [{ required: true, message: '请输入卡片id', trigger: 'blur' }],
        code: [{ required: true, message: '请输入卡片文件名', trigger: 'blur' }],
        sequence: [{ required: true, message: '请输入排序值', trigger: 'blur' }]
      },
      multipleSelection: []
    }
  },

  created() {
    this.getMenuButton()
    this.getList()
    this.getAll()
  },
  methods: {
    checkPermission() {
      this.CardList.add = checkAuthAdd(this.operationList)
      this.CardList.del = checkAuthDel(this.operationList)
      this.CardList.view = checkAuthView(this.operationList)
      this.CardList.update = checkAuthUpdate(this.operationList)
    },
    getMenuButton() {
      requestMenuButton('Village').then(response => {
        this.operationList = response.data
      }).then(() => {
        this.checkPermission()
      })
    },
    getList() {
      this.listLoading = true
      requestList(this.listQuery).then(response => {
        this.list = response.data.items
        this.total = response.data.total
        alert(this.list[0].village_name)
        this.listLoading = false
      })
    },
    getAll() {
      requestAll().then(response => {
        if (response.data) {
          this.propsSelectlist = response.data
        } else {
          this.propsSelectlist = this.propsSelectlist2
        }
      })
    },
    handleFilter(val) {
      // this.listQuery.parent_id = this.valueIdSelectTree
      // this.listQuery.page = 1
      // alert(val)
      this.temp.card_type = val
      this.getList()
    },
    sortChange(data) {
      const { prop, order } = data
      if (order === 'ascending') {
        this.listQuery.sort = '+' + prop
      } else if (order === 'descending') {
        this.listQuery.sort = '-' + prop
      } else {
        this.listQuery.sort = undefined
      }
      this.handleFilter()
    },
    resetTemp() {
      this.valueIdSelectTree2 = 0
      this.temp = {
        id: 0,
        memo: '',
        name: '',
        url: '',
        code: '',
        icon: 'list',
        operate_type: 'none',
        parent_id: 0,
        menu_type: 2,
        status: 1,
        sequence: 10
      }
    },
    handleCreate() {
      this.resetTemp()
      this.dialogStatus = 'create'
      this.dialogFormVisible = true
      this.loading = false
      this.$nextTick(() => {
        this.$refs['dataForm'].clearValidate()
      })
    },
    createData() {
      this.$refs['dataForm'].validate((valid) => {
        if (valid) {
          this.loading = true
          requestCreate(this.temp).then(response => {
            this.temp.id = response.data.id
            this.list.unshift(this.temp)
            this.dialogFormVisible = false
            this.$notify({
              title: '成功',
              message: '创建成功',
              type: 'success',
              duration: 2000
            })
            this.total = this.total + 1
            this.getAll()
          }).catch(() => {
            this.loading = false
          })
        }
      })
    },
    handleDetail(id) {
      this.loading = true
      requestDetail(id).then(response => {
        this.loading = false
        this.temp = response.data
      })
      this.dialogStatus = 'detail'
      this.dialogFormVisible = true
      this.$nextTick(() => {
        this.$refs['dataForm'].clearValidate()
      })
    },
    handleUpdate(id) {
      this.loading = true
      requestDetail(id).then(response => {
        this.loading = false
        this.temp = response.data
        this.valueIdSelectTree2 = this.temp.parent_id
      })
      this.dialogStatus = 'update'
      this.dialogFormVisible = true
      this.$nextTick(() => {
        this.$refs['dataForm'].clearValidate()
      })
    },
    updateData() {
      this.$refs['dataForm'].validate((valid) => {
        if (valid) {
          this.loading = true
          const tempData = Object.assign({}, this.temp)
          requestUpdate(tempData).then(() => {
            for (const v of this.list) {
              if (v.id === this.temp.id) {
                const index = this.list.indexOf(v)
                this.list.splice(index, 1, this.temp)
                break
              }
            }
            this.dialogFormVisible = false
            this.$notify({
              title: '成功',
              message: '更新成功',
              type: 'success',
              duration: 2000
            })
            this.getAll()
          }).catch(() => {
            this.loading = false
          })
        }
      })
    },
    handleDelete(row) {
      var ids = []
      ids.push(row.id)
      this.$confirm('是否确定删除?', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
      }).then(() => {
        requestDelete(ids).then(() => {
          this.$message({
            message: '删除成功',
            type: 'success'
          })
          this.total = this.total - 1
          const index = this.list.indexOf(row)
          this.list.splice(index, 1)
          this.getAll()
        })
      }).catch(() => {
        this.$message({
          type: 'info',
          message: '已取消删除'
        })
      })
    },
    getSelectTreeValue(value, type) {
      if (type === 1) {
        this.valueIdSelectTree = value
        this.handleFilter()
      } else {
        this.valueIdSelectTree2 = value
      }
    },
    handleSelectionChange(val) {
      this.multipleSelection = val
    },
    handleBatchDel() {
      if (this.multipleSelection.length === 0) {
        this.$message({
          message: '未选中任何行',
          type: 'warning',
          duration: 2000
        })
        return
      }
      var ids = []
      for (const v of this.multipleSelection) {
        ids.push(v.id)
      }
      alert(ids)
      this.$confirm('是否确定删除?', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
      }).then(() => {
        requestDelete(ids).then(() => {
          this.$message({
            message: '删除成功',
            type: 'success'
          })
          for (const row of this.multipleSelection) {
            this.total = this.total - 1
            const index = this.list.indexOf(row)
            this.list.splice(index, 1)
          }
          this.getAll()
        })
      }).catch(() => {
        this.$message({
          type: 'info',
          message: '已取消删除'
        })
      })
    }
  }
}
</script>
