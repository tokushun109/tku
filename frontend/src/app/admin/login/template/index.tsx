'use client'

import { zodResolver } from '@hookform/resolvers/zod'
import { useState } from 'react'
import { useForm } from 'react-hook-form'
import { z } from 'zod'

import { Button } from '@/components/bases/Button'
import { Input } from '@/components/bases/Input'
import { ColorType } from '@/types'

import styles from './styles.module.scss'

const LoginSchema = z.object({
    email: z.string().min(1, 'emailを入力してください'),
    password: z.string().min(1, 'パスワードを入力してください'),
})

type LoginForm = z.infer<typeof LoginSchema>

export const AdminLoginTemplate = () => {
    const [isSubmitting, setIsSubmitting] = useState<boolean>(false)

    const {
        register,
        handleSubmit,
        formState: { errors, isValid },
    } = useForm<LoginForm>({
        resolver: zodResolver(LoginSchema),
        mode: 'onChange',
    })

    const onSubmit = async (_data: LoginForm) => {
        setIsSubmitting(true)
        try {
            // TODO: APIの呼び出し実装
            // 本来はここでログイン処理を実装
            // await loginApi(data)
            // router.push('/admin/product')
        } catch (error) {
            console.error('Login error:', error)
        } finally {
            setIsSubmitting(false)
        }
    }

    return (
        <div className={styles['page-admin-login']}>
            <div className={styles['admin-login-area']}>
                <h1 className={styles['login-title']}>管理者ログイン</h1>
                <form onSubmit={handleSubmit(onSubmit)}>
                    <div className={styles['form-field']}>
                        <Input {...register('email')} error={errors.email?.message} label="email(必須)" required type="email" />
                    </div>
                    <div className={styles['form-field']}>
                        <Input {...register('password')} error={errors.password?.message} label="パスワード(必須)" required type="password" />
                    </div>
                    <div className={styles['submit-button']}>
                        <Button colorType={ColorType.Primary} disabled={!isValid || isSubmitting} type="submit">
                            {isSubmitting ? '送信中...' : '確定'}
                        </Button>
                    </div>
                </form>
            </div>
        </div>
    )
}
