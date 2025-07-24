import z from 'zod'

/** サイトフォームのバリデーションスキーマ */
export const SiteSchema = z.object({
    name: z.string().min(1, '名前は必須項目です').max(20, '20文字以内で入力してください'),
    url: z.string().min(1, 'URLは必須項目です').url('正しいURL形式で入力してください'),
})
