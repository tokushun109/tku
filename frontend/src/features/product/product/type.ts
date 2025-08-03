import z from 'zod'

import { ProductSchema } from './schema'

/** 商品フォームのデータ型 */
export type IProductForm = z.infer<typeof ProductSchema>
