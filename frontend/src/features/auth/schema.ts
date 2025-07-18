import z from 'zod'

/** ログインフォームのバリデーションスキーマ */
export const LoginSchema = z.object({
    email: z
        .string()
        .min(1, 'メールアドレスは必須項目です')
        .max(50, 'メールアドレスは50文字以内で入力してください')
        .pipe(z.email('正しいメールアドレスを入力してください')),
    password: z.string().min(1, 'パスワードを入力してください'),
})
