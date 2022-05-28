<template>
  <CommonDrawer
    title="权限分配"
    ok-text="确定"
    cancel-text="取消"
    :visible="visible"
    :loading="loading"
    @on-close="DClose()"
    @on-ok="DSubmit()"
  >
    <a-spin :spinning="treeData.spinningLoading">
      <a-tree
        v-model:checkedKeys="treeData.checkedKeys"
        checkable
        check-strictly
        default-expand-all
        :selected-keys="treeData.selectedKeys"
        :tree-data="treeData.data"
        @select="onSelect"
      />
    </a-spin>
  </CommonDrawer>
</template>

<script lang="ts">
import {
  defineComponent, onMounted, reactive, toRaw, watch,
} from 'vue'
import CommonDrawer from '@/components/Drawer/Drawer.vue'
import http from '@/utils/request'
import type { MenuType } from '@/types/sys'
import { ListObjCompare, ListToTree } from '@/utils'
import { searchParam } from '@/utils/search_param'
import log from '@/utils/log'

interface TreeDataType {
  spinningLoading: boolean
  checkedKeys: string[]
  expandedKeys?: string[]
  autoExpandParent: boolean
  selectedKeys?: string[]
  data: MenuType[]
}
const CommonTree = defineComponent({
  name: 'CommonTree',
  components: {
    CommonDrawer,
  },
  props: {
    visible: {
      type: Boolean,
      default: false,
    },
    loading: {
      type: Boolean,
      default: false,
    },
    data: {
      type: Array,
    },
  },
  emits: ['on-close', 'on-submit'],
  setup(props, { emit }) {
    const treeData = reactive<TreeDataType>({
      spinningLoading: false,
      checkedKeys: [],
      autoExpandParent: true,
      data: [],
    })
    watch(
      () => props.data,
      (newValue: any) => {
        treeData.checkedKeys = toRaw<string[]>(newValue)
      },
      {
        deep: false,
      },
    )
    function getList() {
      treeData.spinningLoading = true
      http<MenuType>({
        url: `/menu${searchParam({
          pageNum: 1,
          pageSize: 1000,
        })}`,
        method: 'GET',
      }).then((res) => {
        res.list?.list.forEach((item: MenuType) => {
          item.key = item.id
        })
        const list = res.list?.list.sort(ListObjCompare('orderNum'))
        log.i(list, '排序的列表')
        treeData.spinningLoading = false
        treeData.data = ListToTree(list || [])
        console.log(treeData.data, 'data')
      })
    }
    onMounted(() => {
      getList()
    })
    function onSelect(selectedKeys: string[], info) {
      console.log(selectedKeys, 'selec')
      console.log(info, 'info')
    }
    function DClose() {
      emit('on-close')
    }
    function DSubmit() {
      emit('on-submit', toRaw(treeData.checkedKeys))
    }

    return {
      // data
      treeData,

      // methods
      onSelect,
      DClose,
      DSubmit,
    }
  },
})
export default CommonTree
</script>
