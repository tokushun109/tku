import z from 'zod'

/** 商品フォームのバリデーションスキーマ */
export const ProductSchema = z.object({
    name: z.string().min(1, '商品名は必須項目です').max(50, '50文字以内で入力してください'),
    description: z.string().optional(),
    price: z.number().min(1, '価格は1円以上で入力してください').max(1000000, '価格は100万円以下で入力してください'),
    isActive: z.boolean(),
    isRecommend: z.boolean(),
    categoryUuid: z.string().optional(),
    targetUuid: z.string().optional(),
    tagUuids: z.array(z.string()).optional(),
    siteDetails: z
        .array(
            z.object({
                salesSiteUuid: z.string(),
                detailUrl: z.string().url('正しいURLを入力してください'),
            }),
        )
        .optional(),
    uploadImages: z.array(z.instanceof(File)).optional(),
    imageItems: z.array(z.any()).optional(), // ImageItem型 - zodでは複雑な型の検証を簡略化
    isImageOrderChanged: z.boolean().optional(),
})
