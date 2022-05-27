import type { CommonTableList } from '@/types/type'

export interface FileItem {
  uid?: string
  name?: string
  status?: 'uploading' | 'done' | 'error' | 'removed'
  response?: any
  url: string
  thumbUrl?: string
  xhr?: string // XMLHttpRequest Header
  linkProps?: string // 下载链接额外的 HTML 属性
  data?: FileBusiness
}

export interface FileBusiness extends CommonTableList {
  fileName: string
  url: string
  fileSize: number
  storagePath: string
}
