import z from 'zod'

import { ClassificationSchema } from './schema'

/** 分類フォームのデータ型 */
export type IClassificationForm = z.infer<typeof ClassificationSchema>
