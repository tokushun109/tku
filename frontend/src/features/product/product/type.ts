import z from 'zod'

import { CreemaDuplicateSchema, ProductSchema } from './schema'

/** 商品フォームのデータ型 */
export type IProductForm = z.infer<typeof ProductSchema>

/** Creema複製フォームのデータ型 */
export type ICreemaDuplicateForm = z.infer<typeof CreemaDuplicateSchema>
