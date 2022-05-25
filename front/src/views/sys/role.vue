<script lang="ts">
import {
  defineComponent, onMounted, reactive, ref, toRaw,
} from 'vue'
import { Form, message } from 'ant-design-vue'
import type { Method } from 'axios'
import type { UnwrapRef } from 'vue'
import type {
  CommonTreeSelectKeys,
  TableDataType,
  TablePaginType,
} from '@/types/type'
import type { MenuType, RoleType } from '@/types/sys'
import http from '@/utils/request'
import type { DrawerProps } from '@/components/Drawer/Drawer.vue'
import CommonDrawer from '@/components/Drawer/Drawer.vue'
import CommonButton from '@/components/Button/Button.vue'
import CommonTree from '@/views/common/tree.vue'
import { searchParam } from '@/utils/search_param'
import log from '@/utils/log'

export interface AllocateType {
  visible: boolean
  loading: boolean
  data?: number[]
  allocateId: string
}
const SysRole = defineComponent({
  name: 'SysRole',
  isRouter: true,
  components: {
    CommonButton,
    CommonDrawer,
    CommonTree,
  },
  setup() {
    const useForm = Form.useForm
    const formData: UnwrapRef<RoleType> = reactive({
      remark: '',
      roleName: '',
    })
    const formRef = ref()
    const editId = reactive<{ id: number | undefined }>({ id: 0 })
    const commonDrawerData = reactive<DrawerProps>({
      title: '添加',
      loading: false,
      visible: false,
    })
    const allocationTree = reactive<AllocateType>({
      visible: false,
      loading: false,
      data: [],
      allocateId: '',
    })
    const rules = reactive({
      roleName: [
        {
          required: true,
          message: '请输入角色名称',
        },
      ],
    })
    const { resetFields, validate, validateInfos } = useForm(formData, rules)

    const tableData = reactive<TableDataType<RoleType>>({
      data: [],
      pageNum: 1,
      pageSize: 10,
      total: 0,
      loading: false,
      columns: [
        {
          title: '角色名称',
          key: 'roleName',
          dataIndex: 'roleName',
        },
        {
          title: '描述',
          key: 'remark',
          dataIndex: 'remark',
        },
        {
          title: '操作',
          key: 'action',
        },
      ],
    })

    function getList() {
      tableData.loading = true
      http<RoleType>({
        url: `role${searchParam({
          pageNum: tableData.pageNum,
          pageSize: tableData.pageSize,
        })}`,
        method: 'GET',
      }).then((res) => {
        tableData.loading = false
        tableData.pageNum = res.list?.pageNum ?? 1
        tableData.total = res.list?.total ?? 0
        tableData.data = res.list?.list || []
        console.log(tableData, 'data')
      })
    }
    onMounted(() => {
      getList()
    })

    function Submit() {
      validate().then(() => {
        const url = 'role'
        let method: Method = 'POST'
        const body = toRaw(formData)
        commonDrawerData.loading = true
        if (editId.id) {
          method = 'PUT'
          body.id = editId.id
        }
        http({
          url,
          method,
          body,
        }).then(() => {
          message.success(`${commonDrawerData.title}成功`)
          commonDrawerData.loading = false
          commonDrawerData.visible = false
          getList()
        })
      })
    }

    function ChangAdd() {
      resetFields()
      commonDrawerData.visible = true
      editId.id = 0
    }

    function Editor(record: RoleType) {
      if (record.id) {
        editId.id = record.id
        commonDrawerData.title = '修改'
        commonDrawerData.visible = true
        formData.remark = record.remark
        formData.roleName = record.roleName
      }
    }
    function Del(record: RoleType) {
      http({ url: `role/${record.id}`, method: 'delete' }).then(() => {
        message.success('删除成功')
        getList()
      })
    }
    function Change(pagin: TablePaginType) {
      tableData.pageNum = pagin.current
      getList()
    }
    function PowerAllocation(record: RoleType) {
      allocationTree.loading = true
      allocationTree.visible = true
      if (record.id != null)
        allocationTree.allocateId = String(record.id)

      http<MenuType>({
        url: `/roleMenuRelationList/${record.id}`,
        method: 'get',
      }).then((res) => {
        if (res && res.data) {
          log.i(res.data, 'res-roleMenuRelationList')
          const ids: number[] = []
          res.data.forEach((item) => {
            ids.push(item.id)
          })
          allocationTree.data = ids
        }
        else {
          allocationTree.data = []
        }
        allocationTree.loading = false
      })
    }
    function Close() {
      allocationTree.visible = false
    }
    function SubmitOk(val: CommonTreeSelectKeys) {
      const data = {
        roleId: Number(allocationTree.allocateId),
        menuId: val.checked,
      }
      allocationTree.loading = true
      http<MenuType>({
        url: 'roleMenuRelation',
        method: 'post',
        body: data,
      })
        .then(() => {
          message.success('更新成功')
          allocationTree.loading = false
          allocationTree.visible = false
        })
        .catch(() => {
          allocationTree.loading = false
        })
    }
    return {
      // data
      tableData,
      commonDrawerData,
      allocationTree,

      // methods
      ChangAdd,
      Submit,
      Editor,
      Del,
      Change,
      PowerAllocation,
      Close,
      SubmitOk,

      // form
      formData,
      formRef,
      rules,
      validateInfos,
    }
  },
})

export default SysRole
</script>

<template>
  <div>
    <CommonButton v-bt-auth:add icon-name="add" @change="ChangAdd" />
    <a-table
      style="margin-top: 15px"
      :columns="tableData.columns"
      :data-source="tableData.data"
      :loading="tableData.loading"
      row-key="id"
      :pagination="{
        total: tableData.total,
      }"
      @change="Change"
    >
      <template #bodyCell="{ column, record }">
        <template v-if="column.key === 'action'">
          <a-button
            v-bt-auth:power
            type="primary"
            style="margin-right: 15px"
            @click="PowerAllocation(record)"
          />

          <a-button v-bt-auth:edit type="primary" style="margin-right: 15px" @click="Editor(record)" />
          <a-popconfirm title="确定删除吗?" ok-text="删除" cancel-text="取消" @confirm="Del(record)">
            <a-button v-bt-auth:del danger />
          </a-popconfirm>
        </template>
      </template>
    </a-table>

    <CommonDrawer
      :title="commonDrawerData.title"
      :visible="commonDrawerData.visible"
      :loading="commonDrawerData.loading"
      ok-text="确定"
      cancel-text="取消"
      @on-ok="Submit()"
      @on-close="commonDrawerData.visible = false"
    >
      <a-form :label-col="{ span: 4 }" :wrapper-col="{ span: 20 }">
        <a-form-item label="角色名称" v-bind="validateInfos.roleName">
          <a-input v-model:value="formData.roleName" />
        </a-form-item>
        <a-form-item label="描述" v-bind="validateInfos.remark">
          <a-input v-model:value="formData.remark" />
        </a-form-item>
      </a-form>
    </CommonDrawer>
    <CommonTree
      :visible="allocationTree.visible"
      :data="allocationTree.data"
      :loading="allocationTree.loading"
      @on-close="Close"
      @on-submit="SubmitOk"
    />
  </div>
</template>
