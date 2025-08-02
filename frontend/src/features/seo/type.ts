import z from 'zod'

import { SeoEditSchema } from './schema'

/** SEO編集フォームのデータ型 */
export type ISeoEdit = z.infer<typeof SeoEditSchema>
