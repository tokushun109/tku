import * as z from 'zod'

/** SEO編集フォームのバリデーションスキーマ */
export const SeoEditSchema = z.object({
    introduction: z.string().min(1, 'サイト説明を入力してください'),
    name: z.string().min(1, 'サイト名を入力してください'),
})
