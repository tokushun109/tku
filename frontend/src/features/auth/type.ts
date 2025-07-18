import z from 'zod'

import { LoginSchema } from './schema'

/** ログインフォームのデータ型 */
export type ILoginForm = z.infer<typeof LoginSchema>
export interface ISession {
    uuid: string
}
