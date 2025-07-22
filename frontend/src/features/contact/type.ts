import z from 'zod'

import { ContactSchema } from './schema'

/** お問い合わせフォームのデータ型 */
export type IContact = z.infer<typeof ContactSchema>
