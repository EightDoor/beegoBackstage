<template>
  <div class="space-margin-bottom">
    <CommonButton
      v-bt-auth:add="{ title: true }"
      title="添加"
      icon-name="add"
      @change="ChangeClick()"
    />
  </div>
  <a-table
    :columns="tableCont.columns"
    row-key="id"
    :data-source="tableCont.data"
    :pagination="{
      total: tableCont.total,
    }"
    :loading="tableCont.loading"
    @change="Change"
  >
    <template #bodyCell="{ column, record }">
      <template v-if="column.key === 'action'">
        <a-button v-bt-auth:edit type="primary" style="margin-right: 15px" @click="Editor(record)" />
        <a-popconfirm title="确定删除吗?" ok-text="删除" cancel-text="取消" @confirm="Del(record)">
          <a-button v-bt-auth:del danger />
        </a-popconfirm>
      </template>
    </template>
  </a-table>
  <CommonDrawer
    :title="drawerData.title"
    :visible="drawerData.visible"
    cancel-text="取消"
    ok-text="确定"
    :loading="drawerData.loading"
    @onClose="drawerData.visible = false"
    @onOk="onSubmit"
  >
    <a-form ref="formRef" :label-col="{ span: 4 }" :wrapper-col="{ span: 20 }">
      <a-form-item label="父级id" v-bind="validateInfos.parentId">
        <a-tree-select
          v-model:value="formData.parentId"
          style="width: 100%"
          :dropdown-style="{ maxHeight: '400px', overflow: 'auto' }"
          :tree-data="treeOptions.options"
          placeholder="请选择"
          tree-default-expand-all
        />
      </a-form-item>
      <a-form-item label="名称" v-bind="validateInfos.name">
        <a-input v-model:value="formData.name" />
      </a-form-item>
      <a-form-item label="排序" v-bind="validateInfos.orderNum">
        <a-input-number v-model:value="formData.orderNum" />
      </a-form-item>
    </a-form>
  </CommonDrawer>
</template>

<script lang="ts">
import {
  defineComponent, nextTick, onMounted, reactive, ref, toRaw,
} from 'vue'
import { Form, message } from 'ant-design-vue'
import type { Method } from 'axios'
import { cloneDeep } from 'lodash-es'
import CommonButton from '@/components/Button/Button.vue'
import type { DrawerProps } from '@/components/Drawer/Drawer.vue'
import CommonDrawer from '@/components/Drawer/Drawer.vue'
import http from '@/utils/request'
import type { DepartType } from '@/types/sys'
import type { TableDataType, TablePaginType } from '@/types/type'
import { ListObjCompare, ListToTree } from '@/utils'
import { searchParam } from '@/utils/search_param'

const SysDepart = defineComponent({
  name: 'SysDepart',
  components: { CommonButton, CommonDrawer },
  setup() {
    const useForm = Form.useForm
    const tableCont = reactive<TableDataType<DepartType>>({
      pageNum: 1,
      pageSize: 10,
      total: 0,
      loading: false,
      columns: [
        {
          title: '名称',
          dataIndex: 'name',
        },
        {
          title: '父级节点',
          dataIndex: 'parentId',
        },
        {
          title: '排序',
          dataIndex: 'orderNum',
        },
        {
          title: '操作',
          key: 'action',
        },
      ],
      data: [],
    })
    const treeOptions = reactive<{ options: DepartType[] }>({ options: [] })
    const drawerData = reactive<DrawerProps>({
      title: '添加',
      visible: false,
      loading: false,
    })
    const formRef = ref()
    const formData = reactive<DepartType>({
      parentId: '0',
      name: '',
      orderNum: 0,
      id: '',
    })
    const rules = ref({
      parent_id: [
        {
          required: true,
          message: '请选择父级',
        },
      ],
      name: [
        {
          required: true,
          message: '请输入名称',
        },
      ],
      order_num: [
        {
          required: true,
          message: '请输入排序',
        },
      ],
    })
    const { validate, resetFields, validateInfos } = useForm(formData, rules)
    function getList() {
      tableCont.loading = true
      http<DepartType>({
        url: `/dept${searchParam({
          page: tableCont.pageNum,
          limit: tableCont.pageSize,
        })}`,
        method: 'GET',
      }).then((res) => {
        let list = res.list?.list.sort(ListObjCompare('order_num')) || []
        tableCont.loading = false
        tableCont.data = ListToTree(list)
        tableCont.total = res.list?.total ?? 0
        list = list.map((item) => {
          item.title = item.name
          item.value = item.id
          return item
        })
        treeOptions.options = ListToTree(list)
      })
    }
    onMounted(() => {
      getList()
    })
    const onSubmit = () => {
      validate()
        .then(() => {
          drawerData.loading = true
          const data = cloneDeep(toRaw(formData))
          delete data.id
          let method: Method = 'POST'
          if (formData.id) {
            method = 'PUT'
            data.id = formData.id
          }
          data.parentId = Number(data.parentId)
          http<DepartType>({
            url: 'dept',
            method,
            body: data,
          }).then(() => {
            message.success(`${drawerData.title}成功`)
            drawerData.loading = false
            drawerData.visible = false
            getList()
          })
        })
        .catch((err) => {
          console.log('error', err)
        })
    }

    function ChangeClick() {
      drawerData.title = '添加'
      drawerData.visible = true
      nextTick(() => {
        resetFields()
      })
    }

    function Editor(record: DepartType) {
      const data = toRaw(record)
      drawerData.title = '修改'
      drawerData.visible = true
      formData.name = data.name
      formData.orderNum = data.orderNum
      formData.parentId = data.parentId
      formData.id = data.id
    }
    function Del(record: DepartType) {
      const data = toRaw(record)
      http({ url: `dept/${data.id}`, method: 'delete' }).then(() => {
        message.success('删除成功')
        getList()
      })
    }
    function Change(pagination: TablePaginType) {
      tableCont.pageNum = pagination.current
      getList()
    }

    return {
      tableCont,
      ChangeClick,
      drawerData,
      treeOptions,
      Change,

      // table
      Editor,
      Del,

      // form
      resetFields,
      rules,
      formData,
      formRef,
      onSubmit,
      validateInfos,
    }
  },
})
export default SysDepart
</script>

<style scoped lang="less">
@import "./depart.less";
</style>
