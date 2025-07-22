import z from 'zod'

/** 分類フォームのバリデーションスキーマ */
export const ClassificationSchema = z.object({
    name: z.string().min(1, '名前は必須項目です').max(20, '20文字以内で入力してください'),
})
