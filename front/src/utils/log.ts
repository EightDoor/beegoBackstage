import { LogUtil } from 'zhoukai_utils'

/** *
 * 日志
 */

const isProd = import.meta.env.PROD
const log = new LogUtil({
  logStatus: !isProd,
})

export default log
