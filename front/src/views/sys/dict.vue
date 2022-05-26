<script lang="ts">
import {
  defineComponent, nextTick, onMounted, reactive, ref, toRaw,
} from 'vue'
import { Form, message } from 'ant-design-vue'
import type { Method } from 'axios'
import DictDrawer from './dict-drawer.vue'
import type { TableDataType, TablePaginType } from '@/types/type'
import type { SysDict } from '@/types/sys/dict'
import http from '@/utils/request'
import type { DrawerProps } from '@/components/Drawer/Drawer.vue'
import CommonDrawer from '@/components/Drawer/Drawer.vue'
import CommonButton from '@/components/Button/Button.vue'
import { searchParam } from '@/utils/search_param'
import log from '@/utils/log'

const SysDictView = defineComponent({
  name: 'SysDict',
  components: {
    CommonButton,
    CommonDrawer,
    DictDrawer,
  },
  setup() {
    const useForm = Form.useForm
    const formRef = ref()
    const formData = reactive<SysDict>({
      name: '',
      serialNumber: '',
      describe: '',
    })
    const RefDictDrawer = ref()
    const editId = reactive<{ id: number | undefined }>({ id: 0 })
    const commonDrawerData = reactive<DrawerProps>({
      title: '添加',
      loading: false,
      visible: false,
    })
    const rules = ref({
      name: [
        {
          required: true,
          message: '请输入名称',
        },
      ],
    })

    const { validate, resetFields, validateInfos } = useForm(formData, rules)
    const tableData = reactive<TableDataType<SysDict>>({
      data: [],
      pageNum: 1,
      pageSize: 10,
      total: 0,
      loading: false,
      columns: [
        {
          title: '字典名称',
          key: 'name',
          dataIndex: 'name',
        },
        {
          title: '字典编号',
          key: 'serialNumber',
          dataIndex: 'serialNumber',
        },
        {
          title: '描述',
          key: 'describe',
          dataIndex: 'describe',
        },
        {
          title: '操作',
          key: 'action',
        },
      ],
    })

    function getList() {
      tableData.loading = true
      http<SysDict>({
        url: `dict${searchParam({
          pageNum: tableData.pageNum,
          pageSize: tableData.pageSize,
        })}`,
        method: 'GET',
      }).then((res) => {
        tableData.loading = false
        tableData.pageNum = res.list?.pageNum ?? 1
        tableData.pageSize = res.list?.pageSize ?? 10
        tableData.total = res.list?.total ?? 0
        tableData.data = res.list?.list || []
        log.i(res.list?.list, 'res.list?.list')
      })
    }
    onMounted(() => {
      getList()
    })

    function Submit() {
      validate().then(() => {
        const url = 'dict'
        let method: Method = 'POST'
        const body = toRaw(formData)
        commonDrawerData.loading = true
        if (editId.id) {
          body.id = editId.id
          method = 'PUT'
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
      commonDrawerData.visible = true
      editId.id = 0
      nextTick(() => {
        resetFields()
      })
    }

    function Editor(record: SysDict) {
      if (record.id) {
        editId.id = record.id
        commonDrawerData.title = '修改'
        commonDrawerData.visible = true
        formData.name = record.name
        formData.serialNumber = record.serialNumber
        formData.describe = record.describe
      }
    }
    function Del(record: SysDict) {
      http({ url: `dict/${record.id}`, method: 'delete' }).then(() => {
        message.success('删除成功')
        getList()
      })
    }
    function Change(pagin: TablePaginType) {
      tableData.pageNum = pagin.current
      getList()
    }

    const Setting = (record: SysDict) => {
      RefDictDrawer.value.IsShow(record)
    }

    function PowerAllocation() {
      //
    }

    function formatStatus(text: string) {
      return text
    }

    return {
      // data
      tableData,
      commonDrawerData,
      RefDictDrawer,
      // methods
      ChangAdd,
      Submit,
      Editor,
      Del,
      Change,
      Setting,
      PowerAllocation,
      formatStatus,

      // form
      formData,
      formRef,
      rules,
      validateInfos,
    }
  },
})

export default SysDictView
</script>

<template>
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
    <template #bodyCell="{ column, text, record }">
      <template v-if="column.key === 'status'">
        {{ formatStatus(text) }}
      </template>
      <template v-if="column.key === 'action'">
        <a-button
          v-bt-auth:power
          type="primary"
          style="margin-right: 15px"
          @click="PowerAllocation()"
        />

        <a-button v-bt-auth:edit type="primary" style="margin-right: 15px" @click="Editor(record)" />
        <a-button
          v-bt-auth:setting
          type="primary"
          style="margin-right: 15px"
          @click="Setting(record)"
        />
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
    <a-form ref="formRef" :label-col="{ span: 4 }" :wrapper-col="{ span: 20 }">
      <a-form-item v-bind="validateInfos.name" label="字典名称">
        <a-input v-model:value="formData.name" />
      </a-form-item>
      <a-form-item v-bind="validateInfos.serialNumber" label="字典编号">
        <a-input v-model:value="formData.serialNumber" />
      </a-form-item>
      <a-form-item v-bind="validateInfos.describe" label="描述">
        <a-input v-model:value="formData.describe" />
      </a-form-item>
    </a-form>
  </CommonDrawer>
  <DictDrawer ref="RefDictDrawer" />
</template>
