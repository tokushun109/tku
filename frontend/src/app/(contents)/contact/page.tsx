'use client'

import { useRouter } from 'next/navigation'
import { useActionState, useEffect } from 'react'
import React from 'react'

import { Button } from '@/components/bases/Button'
import { Input } from '@/components/bases/Input'
import { TextArea } from '@/components/bases/TextArea'

import { submitContact } from './actions'
import styles from './styles.module.scss'

const ContactPage: React.FC = () => {
    const [{ message, success, formData, errors }, formAction, pending] = useActionState(submitContact, {})
    const router = useRouter()

    useEffect(() => {
        // 送信が成功した場合、3秒後にトップページへリダイレクト
        if (success) {
            const timer = setTimeout(() => {
                router.push('/')
            }, 3000)

            return () => clearTimeout(timer)
        }
    }, [success, router])

    return (
        <div className={styles['page-contact']}>
            {!success && (
                <div className={styles['contact-introduction']}>
                    <p>
                        お問い合わせ・
                        <br className={styles['sm']} />
                        ご意見・ご相談はこちらから
                    </p>
                </div>
            )}

            <div className={styles['content-wrapper']}>
                <div className={styles['content-form-wrapper']}>
                    {!success ? (
                        <div className={styles['form-container']}>
                            {message && (
                                <div aria-live="polite" className={styles['error-message']}>
                                    {message}
                                </div>
                            )}

                            <form action={formAction} noValidate>
                                {/* お名前 */}
                                <Input
                                    defaultValue={formData?.name || ''}
                                    error={errors?.name}
                                    id="name"
                                    label="お名前"
                                    maxLength={20}
                                    name="name"
                                    placeholder="山田太郎"
                                    required
                                    type="text"
                                />

                                {/* 会社名 */}
                                <Input
                                    defaultValue={formData?.company || ''}
                                    error={errors?.company}
                                    id="company"
                                    label="会社名"
                                    maxLength={20}
                                    name="company"
                                    placeholder="株式会社サンプル"
                                    type="text"
                                />

                                {/* 電話番号 */}
                                <Input
                                    defaultValue={formData?.phoneNumber || ''}
                                    error={errors?.phoneNumber}
                                    id="phoneNumber"
                                    label="電話番号(-を入れずに入力)"
                                    name="phoneNumber"
                                    placeholder="09012345678"
                                    type="tel"
                                />

                                {/* メールアドレス */}
                                <Input
                                    defaultValue={formData?.email || ''}
                                    error={errors?.email}
                                    id="email"
                                    label="メールアドレス"
                                    maxLength={50}
                                    name="email"
                                    placeholder="example@example.com"
                                    required
                                    type="email"
                                />

                                {/* お問い合わせ内容 */}
                                <TextArea
                                    defaultValue={formData?.content || ''}
                                    error={errors?.content}
                                    id="content"
                                    label="お問い合わせ内容"
                                    name="content"
                                    placeholder="商品についてのご質問、ご要望などをお聞かせください"
                                    required
                                    rows={5}
                                />

                                <div className={styles['submit-container']}>
                                    <Button disabled={pending} formAction={formAction} type="submit">
                                        {pending ? '送信中...' : '送信する'}
                                    </Button>
                                </div>
                            </form>
                        </div>
                    ) : (
                        <div className={styles['content-message-wrapper']}>
                            <strong>お問い合わせを送信しました</strong>
                        </div>
                    )}
                </div>
            </div>
        </div>
    )
}

export default ContactPage
