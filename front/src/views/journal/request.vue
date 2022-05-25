<script lang="ts">
import { defineComponent, onMounted, ref } from 'vue'
import log from '@/utils/log'
import http from '@/utils/request'
import type { ILogLogin } from '@/types/log'

export default defineComponent({
  name: 'Request',
  setup() {
    const page = ref({
      pageNum: 1,
      pageSize: 10,
      total: 0,
    })
    const dataSource = ref<ILogLogin[]>([])
    const columns = ref([
      {
        title: '用户ID',
        dataIndex: 'id',
      },
      {
        title: '昵称',
        dataIndex: 'account',
      },
      {
        title: '请求地址',
        dataIndex: 'requestAddress',
      },
      {
        title: '参数',
        dataIndex: 'params',
      },
      {
        title: 'ip地址',
        dataIndex: 'ip',
      },
      {
        title: '创建时间',
        dataIndex: 'createTime',
      }])

    function change(page, pageSize) {
      log.i(page, 'page')
      log.i(pageSize, 'pageSize')
    }

    function getList() {
      http<ILogLogin>({
        url: 'loginLog',
        method: 'GET',
        params: {
          pageNum: page.value.pageNum,
          pageSize: page.value.pageSize,
        },
      }).then((res) => {
        log.i(res.list, 'res.list')
        dataSource.value = res.list?.list ?? []
        page.value.pageSize = res.list?.pageSize ?? 10
        page.value.pageNum = res.list?.pageNum ?? 1
        page.value.total = res.list?.total ?? 0
      })
    }
    onMounted(() => {
      getList()
    })
    return {
      dataSource,
      columns,
      change,
      page,
    }
  },
})
</script>

<template>
  <a-table
    :data-source="dataSource" :columns="columns" :pagination="{
      change,
      total: page.total,
    }"
  />
</template>
