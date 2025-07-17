'use client'

import { zodResolver } from '@hookform/resolvers/zod'
import React, { useState } from 'react'
import { useForm } from 'react-hook-form'
import { z } from 'zod'

import { Button } from '@/components/bases/Button'
import { Input } from '@/components/bases/Input'
import { ColorType } from '@/types'

import styles from './styles.module.scss'

const LoginSchema = z.object({
    email: z.string().min(1, 'メールアドレスを入力してください').email('正しいメールアドレスを入力してください'),
    password: z.string().min(1, 'パスワードを入力してください'),
})

type LoginForm = z.infer<typeof LoginSchema>

export const LoginTemplate = () => {
    const [errors, setErrors] = useState<string[]>([])

    const {
        register,
        handleSubmit,
        formState: { errors: formErrors, isValid },
    } = useForm<LoginForm>({
        resolver: zodResolver(LoginSchema),
        mode: 'onChange',
    })

    const onSubmit = async (data: LoginForm) => {
        try {
            setErrors([])
            // TODO: API呼び出し実装
            // eslint-disable-next-line no-console
            console.log('Login data:', data)
            // 成功時のリダイレクト処理
            // router.push('/admin/product')
        } catch (error) {
            setErrors([error instanceof Error ? error.message : '予期しないエラーが発生しました'])
        }
    }

    return (
        <div className={styles['page-admin-login']}>
            <div className={styles['admin-login-area']}>
                {errors.length > 0 && (
                    <div className={styles['error-container']}>
                        {errors.map((error, index) => (
                            <div className={styles['error-message']} key={index}>
                                {error}
                            </div>
                        ))}
                    </div>
                )}

                <h3 className={styles['login-title']}>ログイン</h3>

                <form className={styles['login-form']} onSubmit={handleSubmit(onSubmit)}>
                    <Input
                        {...register('email')}
                        error={formErrors.email?.message}
                        label="email"
                        placeholder="メールアドレスを入力"
                        required
                        type="email"
                    />

                    <Input
                        {...register('password')}
                        error={formErrors.password?.message}
                        label="パスワード"
                        placeholder="パスワードを入力"
                        required
                        type="password"
                    />

                    <Button className={styles['submit-button']} colorType={ColorType.Primary} disabled={!isValid} type="submit">
                        確定
                    </Button>
                </form>
            </div>
        </div>
    )
}
