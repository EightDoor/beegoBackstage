// 请求接口分页
export interface PaginType {
  pageNum: number
  pageSize: number
  total: number
}

interface CommReqList<T> extends PaginType {
  list: T[]
}
export interface CommonResponse<T> {
  list: CommReqList<T> | null
  data: any
}

// 表格默认的分页
export interface TablePaginType {
  current: number
  pageSize: number
  total: number
}

export interface CommonResponseSing<T> extends PaginType {
  list: T
}

interface AntColumnSlot {
  bodyCell?: {
    customRender: string
  }
  title: string
  dataIndex?: string
  key?: string
  fixed?: string | boolean
  width?: number
}

interface AntColumn extends AntColumnSlot {}

// 表格
export interface TableDataType<T> extends PaginType {
  data: T[]
  columns: AntColumn[]
  loading: boolean
}

// 通用的业务类型字段
export interface CommonTableList {
  id?: number
  createdAt?: string
  updatedAt?: string
  deletedAt?: string
}

// 树形tree-select
export interface CommonTreeSelect {
  title?: string
  value?: string
  key?: string
}
// 树形结构选中selectKey
export interface CommonTreeSelectKeys {
  checked: string[]
  halfChecked: string[]
}
