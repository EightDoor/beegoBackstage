import type { FileBusiness, FileItem } from '@/types'

const BusinessUtils = {
  /**
   * ImageUpload 图片提交数据格式化
   * @param data
   */
  formatUploadImg(data: FileBusiness[] | FileBusiness) {
    if (data as FileBusiness) {
      return (data as FileBusiness).id
    }
    else if (data as FileBusiness[]) {
      const ids: number[] = [];
      (data as FileBusiness[]).forEach((item) => {
        ids.push(item.id ?? 0)
      })
      if (ids.length > 0)
        return ids.join(',')
    }
    return ''
  },
  /**
   * ImageUpload 图片回显列表
   * @param img
   */
  formatUploadShow(img: string): FileItem[] {
    if (img) {
      return [{
        thumbUrl: img,
        url: img,
        uid: String(Date.now()),
      }]
    }
    return []
  },
}

export default BusinessUtils
