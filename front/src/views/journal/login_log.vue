<script lang="ts">
import { defineComponent, onMounted, ref } from 'vue'
import { dateUtil } from 'zhoukai_utils'
import http from '@/utils/request'
import log from '@/utils/log'
import type { ILogLogin } from '@/types/log'

export default defineComponent({
  name: 'LoginLog',
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
        dataIndex: ['user', 'id'],
      },
      {
        title: '账号',
        dataIndex: ['user', 'account'],
      },
      {
        title: 'ip地址',
        dataIndex: 'ip',
      },
      {
        title: '登录设备',
        dataIndex: 'equipment',
      }, {
        title: '创建时间',
        dataIndex: 'createdAt',
      }])

    function getList() {
      http<ILogLogin>({
        url: 'logLog',
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

    function change(pageNum, pageSize) {
      log.i(pageNum, 'page')
      log.i(pageSize, 'pageSize')
      page.value.pageNum = pageNum.current
      getList()
    }
    function formatTime(val) {
      return dateUtil.formatTime(val, 'YYYY-MM-DD HH:mm:ss')
    }
    return {
      dataSource,
      columns,
      change,
      page,
      formatTime,
    }
  },
})
</script>

<template>
  <a-table
    :data-source="dataSource"
    :columns="columns"
    :pagination="{
      total: page.total,
    }"
    @change="change"
  >
    <template #bodyCell="{ column, record }">
      <template v-if="column.dataIndex === 'createdAt'">
        {{ formatTime(record.createdAt) }}
      </template>
    </template>
  </a-table>
</template>
