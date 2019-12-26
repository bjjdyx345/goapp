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
        v-model="listQuery.order_username"
        placeholder="姓名"
        clearable
        prefix-icon="el-icon-search"
        style="width: 90px;"
        class="filter-item"
        @keyup.enter.native="handleFilter"
        @clear="handleFilter"
      />
      <el-input
        v-model="listQuery.order_status"
        placeholder="订单状态"
        clearable
        prefix-icon="el-icon-search"
        style="width: 200px;"
        class="filter-item"
        @keyup.enter.native="handleFilter"
        @clear="handleFilter"
      />
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
      <el-table-column label="订单ID" align="center">
        <template slot-scope="scope">
          <span>{{ scope.row.id }}</span>
        </template>
      </el-table-column>
      <el-table-column label="姓名" align="center">
        <template slot-scope="scope">
          <span>{{ scope.row.order_username }}</span>
        </template>
      </el-table-column>
      <el-table-column label="手机号" align="center">
        <template slot-scope="scope">
          <span>{{ scope.row.order_phone_number }}</span>
        </template>
      </el-table-column>
      <el-table-column label="卡号" align="center">
        <template slot-scope="scope">
          <span>{{ scope.row.order_card_id }}</span>
        </template>
      </el-table-column>
      <el-table-column label="订单价格" align="center">
        <template slot-scope="scope">
          <span>{{ scope.row.order_price }}</span>
        </template>
      </el-table-column>
      <el-table-column label="订单地址ID" align="center">
        <template slot-scope="scope">
          <span>{{ scope.row.order_address }}</span>
        </template>
      </el-table-column>
      <el-table-column label="订单状态" align="center">
        <template slot-scope="scope">
          <span>{{ scope.row.order_status }}</span>
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
        <el-form-item label="姓名" prop="order_username">
          <el-input v-model="temp.order_username" />
        </el-form-item>
        <el-form-item label="手机号" prop="order_phone_number">
          <el-input v-model="temp.order_phone_number" type="tel" />
        </el-form-item>
        <el-form-item label="卡号" prop="order_card_id">
          <el-input v-model="temp.order_card_id" />
        </el-form-item>
        <el-form-item label="价格" prop="order_price">
          <el-input v-model.number="temp.order_price" type="number" />
        </el-form-item>

        <el-form-item v-if="dialogStatus==='detail'" label="订单状态" prop="order_status">
          <el-input v-model="temp.order_status" :disabled="true" />
        </el-form-item>

        <el-form-item v-if="dialogStatus!=='detail'" label="订单状态" prop="order_status">
          <el-select
            v-model="temp.order_status"
            type="number"
            placeholder="订单状态"
          >
            <el-option
              v-for="item in orderstatusTypeOptions"
              :key="item.key"
              :label="item.display_name"
              :value="item.display_name"
            />
          </el-select>
        </el-form-item>

        <el-form-item v-if="dialogStatus!=='detail'" label="绑定卡号" prop="order_card_id">
          <el-select
            v-model="temp.order_card_id"
            type="number"
            placeholder="绑定卡号"
            @change="getBindAddress"
          >
            <el-option
              v-for="item in cardlist"
              :key="item"
              :label="item"
              :value="item"
            />
          </el-select>
        </el-form-item>

        <el-form-item v-if="dialogStatus!=='detail'" label="绑定地址" prop="address_id">
          <el-select
            v-model="myselectvalue"
            placeholder="绑定卡号"
            value-key="address_id"
            clearable
            style="width: 180px"
            class="filter-item"
            @change="getAddressId"
          >
            <el-option
              v-for="item in addresslist"
              :key="item.address_id"
              :label="item.id_and_address"
              :value="item.address_id"
            />
          </el-select>
        </el-form-item>

        <el-form-item label="具体地址" prop="address_id">
          <el-input
            v-model="temp.address_content"
            rows="5"
            :type="textarea"
          />
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
import { requestList, requestDetail, requestUpdate, requestCreate, requestAll, requestDelete } from '@/api/app/sys/book'
import { requestMenuButton } from '@/api/app/sys/menu'
import { requestFirstCityAll, requestSecondCity, requestThirdCity } from '@/api/app/sys/city'
import { requestAllVillageAddress } from '@/api/app/sys/village'

import waves from '@/directive/waves'
import Pagination from '@/components/Pagination' // Secondary package based on el-pagination
// import SelectTree from '@/components/TreeSelect'
import { checkAuthAdd, checkAuthDel, checkAuthView, checkAuthUpdate } from '@/utils/permission'

const orderstatusTypeOptions = [
  { key: 1, display_name: 'start' },
  { key: 2, display_name: 'Unconfirmed' },
  { key: 3, display_name: 'Confirmed' },
  { key: 4, display_name: 'Delivering' },
  { key: 5, display_name: 'Done' }
]

export default {
  name: 'Order',
  components: { Pagination }, // SelectTree
  directives: { waves },
  data() {
    return {
      valueIdSelectTree: 0,
      propsSelectlist: [],
      operationList: [],
      CardList: {
        add: false,
        del: false,
        view: false,
        update: false
      },
      myselectvalue: '',
      orderstatusTypeOptions,
      tableKey: 0,
      addresslist: [],
      firstcitylist: [],
      secondcitylist: [],
      cardlist: [],
      list: [],
      total: 0,
      listLoading: true,
      loading: true,
      listQuery: {
        page: 1,
        limit: 20,
        id: undefined,
        key: undefined,
        order_card_id: undefined,
        order_username: undefined,
        order_phone_number: undefined,
        order_address: undefined,
        order_price: undefined,
        order_status: undefined,
        sort: '-id'
      },
      sortOptions: [
        { label: 'ID Ascending', key: '+id' },
        { label: 'ID Descending', key: '-id' },
        { label: 'card_id Ascending', key: '+card_id' },
        { label: 'card_id Descending', key: '-card_id' }
      ],
      temp: {
        id: 0,
        order_card_id: '',
        order_username: '',
        order_phone_number: '',
        order_address: '',
        order_price: 0,
        order_status: '',
        address_id: 0,
        address_content: '',
        address_city: ''
      },
      dialogFormVisible: false,
      dialogStatus: '',
      textMap: {
        update: '编辑',
        create: '添加',
        detail: '详情'
      },
      rules: {
        type: [{ required: true, message: '请输入', trigger: 'change' }],
        order_name: [{ required: true, message: '请输入姓名', trigger: 'blur' }],
        order_phone_number: [{ required: true, message: '请输入手机号', trigger: 'blur' }],
        order_address: [{ required: true, message: '请输入地址', trigger: 'blur' }]
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
      requestMenuButton('Book').then(response => {
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
    getAllFirstCity() {
      requestFirstCityAll().then(response => {
        this.firstcitylist = response.data
      })
    },
    getAllSecondCity(cityname) {
      this.secondcitylist = []
      if (cityname !== '') {
        requestSecondCity(cityname).then(response => {
          this.secondcitylist = response.data
          // alert('has request leve2 city')
        })
      }
    },
    getBindAddress(cityname) {
      this.thirdcitylist = []
      if (cityname !== '') {
        requestThirdCity(cityname).then(response => {
          this.thirdcitylist = response.data
          // alert('has request leve3 city')
        })
      }
    },
    getAddressId: function(val) {
      const obj = this.addresslist.find(function(item) {
        return item.address_id === val
      })
      // console.log( obj )
      this.temp.address_id = obj.address_id
      this.temp.address_content = obj.address_content
    },
    handleFilter(val) {
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
    //  this.valueIdSelectTree2 = 0
      this.temp = {
        id: 0,
        order_card_id: '',
        order_username: '',
        order_phone_number: '',
        order_address: '',
        order_price: 0,
        order_status: ''
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
      requestAllVillageAddress().then(response => {
        console.log(response)
        this.addresslist = response.data
        // console.log(this.addresslist)
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
    handleUpdate: function(id) {
      this.loading = true
      requestDetail(id).then(response => {
        this.loading = false
        this.temp = response.data
      })
      requestFirstCityAll().then(response => {
        this.firstcitylist = response.data
      })
      requestAllVillageAddress().then(response => {
        console.log(response)
        this.addresslist = response.data
        // console.log(this.addresslist)
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
