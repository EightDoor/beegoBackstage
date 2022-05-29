// 是否开发环境
const isDev = import.meta.env.DEV

const Config = {
  // 图片上传
  uploadUrl: isDev ? 'http://localhost:8081/api/upload/' : 'http://localhost/api/upload/',
}

export default Config
