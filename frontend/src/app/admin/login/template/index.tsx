'use client'

import { useRouter } from 'next/navigation'
import { useActionState, useState, startTransition, useEffect } from 'react'

import { Button } from '@/components/bases/Button'
import { Input } from '@/components/bases/Input'
import { Message, MessageType } from '@/components/bases/Message'
import { loginAction } from '@/features/auth/action'
import { ColorType } from '@/types'
import { NavigationType } from '@/types/enum/navigation'

import styles from './styles.module.scss'

export const AdminLoginTemplate = () => {
    const router = useRouter()
    const [emailError, setEmailError] = useState<string>('')
    const [passwordError, setPasswordError] = useState<string>('')

    const [state, formAction, pending] = useActionState(loginAction, { success: false })

    // Server Actionの結果に基づいてリダイレクト処理
    useEffect(() => {
        if (state.success) {
            router.push(NavigationType.AdminProduct)
        }
    }, [state.success, router])

    // フォーム送信時の入力検証
    const handleFormSubmit = (event: React.FormEvent<HTMLFormElement>) => {
        event.preventDefault()

        const formData = new FormData(event.currentTarget)
        const email = formData.get('email') as string
        const password = formData.get('password') as string

        // クライアント側バリデーション
        let hasError = false

        if (!email || email.trim() === '') {
            setEmailError('emailを入力してください')
            hasError = true
        } else {
            setEmailError('')
        }

        if (!password || password.trim() === '') {
            setPasswordError('パスワードを入力してください')
            hasError = true
        } else {
            setPasswordError('')
        }

        if (hasError) {
            return
        }

        // Server Actionを実行
        startTransition(() => {
            formAction(formData)
        })
    }

    return (
        <div className={styles['page-admin-login']}>
            <div className={styles['admin-login-area']}>
                <h1 className={styles['login-title']}>ログイン</h1>
                {state.error && <Message type={MessageType.Error}>{state.error}</Message>}
                <form noValidate onSubmit={handleFormSubmit}>
                    <div className={styles['form-field']}>
                        <Input error={emailError} label="email(必須)" name="email" onChange={() => setEmailError('')} required type="email" />
                    </div>
                    <div className={styles['form-field']}>
                        <Input
                            error={passwordError}
                            label="パスワード(必須)"
                            name="password"
                            onChange={() => setPasswordError('')}
                            required
                            type="password"
                        />
                    </div>
                    <div className={styles['submit-button']}>
                        <Button colorType={ColorType.Primary} disabled={pending} type="submit">
                            {pending ? '送信中...' : '確定'}
                        </Button>
                    </div>
                </form>
            </div>
        </div>
    )
}
