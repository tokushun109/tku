import z from 'zod'

import { ContactSchema } from './schema'

/** お問い合わせフォームのデータ型 */
export type IContact = z.infer<typeof ContactSchema>

/** お問い合わせ一覧用のデータ型 */
export interface IContactListItem {
    company?: string
    content: string
    createdAt: string
    email: string
    id: number
    name: string
    phoneNumber?: string
}
