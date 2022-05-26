export interface PanesType {
  id: number
  title: string
  path: string
  parentId: number
  closable?: boolean
  timestamp?: number
}
export interface CrumbsStoreType {
  panes: PanesType[]
}

export default {
  namespace: true,
  state: {
    panes: [],
  },
}
