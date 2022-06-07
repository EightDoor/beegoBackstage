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
    :scroll="{ x: 500 }"
    :columns="tableCont.columns"
    row-key="id"
    :data-source="tableCont.data"
    :pagination="{
      total: tableCont.total,
      pageSize: 10000,
      hideOnSinglePage: true,
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
    @onClose="onClose"
    @onOk="onSubmit"
  >
    <a-form
      ref="formRef"
      :label-col="{ span: 4 }"
      :wrapper-col="{ span: 20 }"
    >
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
      <a-form-item label="名称" v-bind="validateInfos.title">
        <a-input v-model:value="formData.title" />
      </a-form-item>
      <a-form-item label="菜单类型" v-bind="validateInfos.type">
        <a-select v-model:value="formData.type">
          <a-select-option
            v-for="(item, index) in optionsType"
            :key="index"
            :value="item.value"
          >
            {{ item.label }}
          </a-select-option>
        </a-select>
      </a-form-item>
      <a-form-item v-if="formData.type === 3" label="权限标识" v-bind="validateInfos.perms">
        <a-input v-model:value="formData.perms" />
      </a-form-item>
      <template v-else>
        <a-form-item label="菜单name" v-bind="validateInfos.name">
          <a-input v-model:value="formData.name" />
        </a-form-item>
        <a-form-item label="是否首页" v-bind="validateInfos.isHome">
          <a-radio-group v-model:value="formData.isHome">
            <a-radio :value="1">
              是
            </a-radio>
            <a-radio :value="0">
              否
            </a-radio>
          </a-radio-group>
        </a-form-item>
        <a-form-item label="图标" v-bind="validateInfos.icon">
          <a-input-group compact>
            <a-input v-model:value="formData.icon" style="width: calc(100% - 200px)" />
            <a-button type="primary" @click="selectIcon">
              选择图标
            </a-button>
          </a-input-group>
        </a-form-item>
        <a-form-item label="重定向地址" v-bind="validateInfos.redirect">
          <a-input v-model:value="formData.redirect" />
        </a-form-item>
        <a-form-item label="是否隐藏" v-bind="validateInfos.hidden">
          <a-radio-group v-model:value="formData.hidden" name="radioGroup">
            <a-radio :value="0">
              否
            </a-radio>
            <a-radio :value="1">
              是
            </a-radio>
          </a-radio-group>
        </a-form-item>
      </template>
      <a-form-item label="排序" v-bind="validateInfos.orderNum">
        <a-input-number v-model:value="formData.orderNum" />
      </a-form-item>
    </a-form>
  </CommonDrawer>

  <CommonDrawer
    :footer="false"
    title="选择图标"
    cancel-text="取消"
    :visible="iconSelectVisible"
    @onClose="onIconClose"
  >
    <ul class="drawer_icon">
      <li v-for="(item, index) of iconList" :key="index" @click="clickIcon(item)">
        <SvgIcon :name="item" />
      </li>
    </ul>
  </CommonDrawer>
</template>

<script lang="ts">
import {
  defineComponent, nextTick, onMounted, reactive, ref, toRaw, unref,
} from 'vue'

import type { FormInstance } from 'ant-design-vue'
import { Form, message } from 'ant-design-vue'
import type { Method } from 'axios'
import { cloneDeep } from 'lodash-es'
import type { Rule } from 'ant-design-vue/lib/form'
import SvgIcon from '@/components/SvgIcon/index.vue'
import CommonButton from '@/components/Button/Button.vue'
import type { DrawerProps } from '@/components/Drawer/Drawer.vue'
import CommonDrawer from '@/components/Drawer/Drawer.vue'
import http from '@/utils/request'
import type { MenuType } from '@/types/sys'
import type { TableDataType, TablePaginType } from '@/types/type'
import { ListObjCompare, ListToTree } from '@/utils'
import { searchParam } from '@/utils/search_param'
import log from '@/utils/log'

const SysMenu = defineComponent({
  name: 'SysMenu',
  components: {
    CommonButton,
    CommonDrawer,
    SvgIcon,
  },
  setup() {
    const useForm = Form.useForm
    const optionsType = ref([
      {
        value: 1,
        label: '目录',
      },
      {
        value: 2,
        label: '菜单',
      },
      {
        value: 3,
        label: '按钮',
      },
    ])
    const width = 150
    const tableCont = reactive<TableDataType<MenuType>>({
      pageNum: 1,
      pageSize: 1000,
      total: 0,
      loading: false,
      columns: [
        {
          title: '名称',
          dataIndex: 'title',
          width: width * 2,
          fixed: true,
        },
        {
          title: '类型',
          dataIndex: 'type',
          width: width / 2,
        },
        {
          title: '图标',
          dataIndex: 'icon',
          width,
        },
        {
          title: '重定向地址',
          dataIndex: 'redirect',
          width,
        },
        {
          title: '父级节点',
          dataIndex: 'parentId',
          width,
        },
        {
          title: '权限标识',
          dataIndex: 'perms',
          width: width / 2,
        },
        {
          title: '菜单标识',
          dataIndex: 'name',
          width: width / 2,
        },
        {
          title: '排序',
          dataIndex: 'orderNum',
          width: width / 2,
        },
        {
          title: '操作',
          key: 'action',
          fixed: 'right',
          width: 200,
        },
      ],
      data: [],
    })
    const treeOptions = reactive<{ options: MenuType[] }>({ options: [] })
    const drawerData = reactive<DrawerProps>({
      title: '添加',
      visible: false,
      loading: false,
    })
    const formRef = ref<FormInstance>()
    let formData = reactive<MenuType>({
      parentId: 0,
      path: '',
      component: '',
      redirect: '',
      icon: '',
      title: '',
      perms: '',
      name: '',
      type: 1,
      orderNum: 0,
      id: 0,
      hidden: 0,
      isHome: 0,
    })
    const rules = ref<Record<string, Rule[]>>({
      parentId: [
        {
          required: true,
          message: '请选择父级',
        },
      ],
      title: [
        {
          required: true,
          message: '请输入名称',
        },
      ],
      type: [
        {
          required: true,
          message: '请选择类型',
        },
      ],
      orderNum: [
        {
          required: true,
          message: '请输入排序',
        },
      ],
    })
    const { validate, resetFields, validateInfos } = useForm(formData, rules)
    function getList() {
      tableCont.loading = true
      http<MenuType>({
        url: `/menu${searchParam({
          pageNum: tableCont.pageNum,
          pageSize: tableCont.pageSize,
        })}`,
        method: 'GET',
      }).then((res) => {
        const list = cloneDeep(res.list?.list.sort(ListObjCompare('orderNum'))) || []
        tableCont.loading = false
        tableCont.data = ListToTree(list)
        tableCont.total = res.list?.total ?? 0
        const treeMenus = cloneDeep(res.list?.list || [])
        treeMenus.forEach((item) => {
          item.value = String(item.id)
        })
        treeOptions.options = ListToTree(treeMenus)
      })
    }
    onMounted(() => {
      console.log('执行了')
      getList()
    })
    const onSubmit = () => {
      validate()
        .then(() => {
          drawerData.loading = true
          const data = cloneDeep(toRaw(formData))
          if (!data.name)
            data.name = data.perms
          delete data.id
          let method: Method = 'POST'
          if (formData.id) {
            method = 'PUT'
            data.id = formData.id
          }
          data.parentId = Number(data.parentId)
          http<MenuType>({
            url: 'menu',
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

    function Editor(record: MenuType) {
      drawerData.title = '修改'
      drawerData.visible = true
      log.i(record, '修改值')
      formData = Object.assign(formData, record)
    }
    function Del(record: MenuType) {
      const data = toRaw(record)
      log.i(data)
      http({ url: `menu/${data.id}`, method: 'delete' }).then(() => {
        message.success('删除成功')
        getList()
      })
    }
    function Change(pagination: TablePaginType) {
      tableCont.pageNum = pagination.current
      getList()
    }
    function onClose() {
      drawerData.visible = false
    }
    // 选择图标
    const iconSelectVisible = ref(false)
    const iconList = ref(['dept', 'dict', 'menu', 'role', 'user'])
    function selectIcon() {
      iconSelectVisible.value = true
    }
    function onIconClose() {
      iconSelectVisible.value = false
    }
    function clickIcon(val: string) {
      formData.icon = val
      iconSelectVisible.value = false
    }

    return {
      tableCont,
      ChangeClick,
      drawerData,
      treeOptions,
      Change,
      optionsType,
      selectIcon,
      iconSelectVisible,
      onIconClose,
      iconList,
      clickIcon,

      // table
      Editor,
      Del,

      // form
      formData,
      formRef,
      rules,
      onSubmit,
      onClose,
      validateInfos,
    }
  },
})
export default SysMenu
</script>

<style scoped lang="less">
@import "./depart.less";
.drawer_icon {
  display: flex;
  flex-direction: row;
  li {
    list-style: none;
    margin-right: 15px;
    &:hover {
      cursor: pointer;
    }
  }
}
</style>
