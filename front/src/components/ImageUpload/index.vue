<template>
  <div>
    <a-upload
      v-model:file-list="fileList"
      :action="Config.uploadUrl"
      :multiple="true"
      :max-count="props.limit"
      :data="extendedParam"
      :headers="headers"
      list-type="picture"
      @change="handleChange"
    >
      <a-button>
        <UploadOutlined />
        上传
      </a-button>
    </a-upload>
  </div>
</template>

<script lang="ts">
import { UploadOutlined } from '@ant-design/icons-vue'
import type { PropType } from 'vue'
import {
  defineComponent,
  markRaw,
  onMounted,
  reactive,
  ref, toRaw,
  watch,
} from 'vue'
import { message } from 'ant-design-vue'
import { isArray, isObject } from 'lodash-es'
import Config from '@/config'
import log from '@/utils/log'
import http from '@/utils/request'
import type { FileBusiness, FileItem } from '@/types'
import { TOKEN } from '@/utils/constant'

interface FileInfo {
  file: FileItem
  fileList: FileItem[]
  event: any
}

export default defineComponent({
  name: 'CommImageUpload',
  components: {
    UploadOutlined,
  },
  props: ['list', 'limit'],
  emits: ['update:list'],
  setup(props, { emit }) {
    const headers = ref()
    const fileList = ref<FileItem[]>([])
    const isFirst = ref(false)
    const extendedParam = reactive({
    })

    function delFile(key?: number) {
      http({
        url: `upload/${key}`,
        method: 'DELETE',
        body: key,
      }).then(() => {
        message.success('删除成功')
      })
    }

    const handleChange = (info: FileInfo) => {
      log.d(info, '图片文件改变')
      let resFileList = [...info.fileList]

      resFileList = resFileList.map((file) => {
        if (file.status === 'done') {
          log.d(file.response, 'file.response')
          if (file.response) {
            file.data = file.response.data
            file.url = file.response.data.url
          }
        }
        return file
      })

      const { file } = info
      if (file.status === 'removed') {
        if (file.data) {
          log.d(file, 'file')
          delFile(file.data.id)
        }
      }

      fileList.value = markRaw(resFileList)
      const result = resFileList.filter(item => item.data)
      const list: FileBusiness[] = []
      result.forEach((item) => {
        list.push(item.data as FileBusiness)
      })
      if (props.limit === 1 && list.length > 0)
        emit('update:list', list[0])
      else
        emit('update:list', list)
    }

    onMounted(() => {
      headers.value = {
        Authorization: `Bearer ${localStorage.getItem(TOKEN)}`,
      }
    })

    watch(
      () => props.list,
      (newVal: FileBusiness | FileBusiness[] | null) => {
        log.d(newVal, 'watch - 文件')
        if (newVal) {
          if (!isFirst.value) {
            const newData = toRaw(newVal)
            if (isObject(newData))
              fileList.value = [newData as FileBusiness]
            else if (isArray(newData))
              fileList.value = newData as FileBusiness[]
            isFirst.value = true
          }
        }
        else {
          fileList.value = []
        }
      },
      {
        deep: true,
        immediate: true,
      },
    )

    return {
      handleChange,

      Config,
      extendedParam,
      fileList,
      headers,
      props,
    }
  },
})
</script>

<style scoped lang="less"></style>
