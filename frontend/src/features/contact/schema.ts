import * as z from 'zod'

/** お問い合わせフォームのバリデーションスキーマ */
export const ContactSchema = z.object({
    name: z.string().min(1, 'お名前は必須項目です').max(20, 'お名前は20文字以内で入力してください'),
    company: z.string().max(20, '会社名は20文字以内で入力してください').optional(),
    phoneNumber: z.string().regex(/^\d*$/, '電話番号は数字のみで入力してください').optional(),
    email: z
        .string()
        .min(1, 'メールアドレスは必須項目です')
        .max(50, 'メールアドレスは50文字以内で入力してください')
        .pipe(z.email('正しいメールアドレスを入力してください')),
    content: z.string().min(1, 'お問い合わせ内容は必須項目です'),
})
