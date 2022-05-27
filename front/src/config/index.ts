// 是否开发环境
const isDev = import.meta.env.DEV

const Config = {
  // 图片上传
  uploadUrl: isDev ? 'http://localhost:8081/api/upload/' : 'http://beegobackstage.start6.cn/upload/',
}

export default Config
