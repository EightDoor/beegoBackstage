import type { UserType } from '@/types/sys'

export interface ILogLogin {
  id: number
  createdAt: string
  updatedAt: string
  deletedAt: string
  equipment: string
  user: UserType
  ip: string
}
