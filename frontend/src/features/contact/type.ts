import z from 'zod'

import { contactSchema } from './schema'

/** お問い合わせフォームのデータ型 */
export type IContact = z.infer<typeof contactSchema>
