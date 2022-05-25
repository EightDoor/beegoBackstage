import { createLogger, createStore } from 'vuex'

import type { SysStoreType } from './sys/sys-store'
import sys from './sys/sys-store'
import type { CrumbsStoreType } from './sys/sys-crumbs'
import crumbs from './sys/sys-crumbs'

const debug = process.env.NODE_ENV !== 'production'

interface RootState {
  sys: SysStoreType
  crumbs: CrumbsStoreType
}
export default createStore<RootState>({
  modules: {
    sys,
    crumbs,
  },
  strict: debug,
  plugins: debug ? [createLogger()] : [],
})
