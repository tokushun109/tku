'use client'

import { zodResolver } from '@hookform/resolvers/zod'
import { useRouter } from 'next/navigation'
import { useState } from 'react'
import React from 'react'
import { useForm, SubmitHandler } from 'react-hook-form'

import { postContact } from '@/apis/contact'
import { Button } from '@/components/bases/Button'
import { Form } from '@/components/bases/Form'
import { Input } from '@/components/bases/Input'
import { TextArea } from '@/components/bases/TextArea'
import { ContactSchema } from '@/features/contact/schema'
import { IContact } from '@/features/contact/type'

import styles from './styles.module.scss'

const ContactPage = () => {
    const [success, setSuccess] = useState<boolean>(false)
    const [isSubmitting, setIsSubmitting] = useState<boolean>(false)
    const [submitError, setSubmitError] = useState<string | null>(null)
    const router = useRouter()

    const {
        register,
        handleSubmit,
        formState: { errors, isValid },
    } = useForm<IContact>({
        mode: 'onChange',
        resolver: zodResolver(ContactSchema),
    })

    const onSubmit: SubmitHandler<IContact> = async (data) => {
        try {
            setIsSubmitting(true)
            setSubmitError(null)
            await postContact({ form: data })
            setSuccess(true)
            setTimeout(() => {
                router.push('/')
            }, 3000)
        } catch {
            setSubmitError('送信中にエラーが発生しました。もう一度お試しください。')
        } finally {
            setIsSubmitting(false)
        }
    }

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
                            {submitError && (
                                <div aria-live="polite" className={styles['error-message']}>
                                    {submitError}
                                </div>
                            )}

                            <Form noValidate onSubmit={handleSubmit(onSubmit)}>
                                {/* お名前 */}
                                <Input
                                    {...register('name')}
                                    error={errors.name?.message}
                                    id="name"
                                    label="お名前"
                                    placeholder="山田太郎"
                                    required
                                    type="text"
                                />

                                {/* 会社名 */}
                                <Input
                                    {...register('company')}
                                    error={errors.company?.message}
                                    id="company"
                                    label="会社名"
                                    placeholder="株式会社サンプル"
                                    type="text"
                                />

                                {/* 電話番号 */}
                                <Input
                                    {...register('phoneNumber')}
                                    error={errors.phoneNumber?.message}
                                    id="phoneNumber"
                                    label="電話番号(-を入れずに入力)"
                                    placeholder="09012345678"
                                    type="tel"
                                />

                                {/* メールアドレス */}
                                <Input
                                    {...register('email')}
                                    error={errors.email?.message}
                                    id="email"
                                    label="メールアドレス"
                                    placeholder="example@example.com"
                                    required
                                    type="email"
                                />

                                {/* お問い合わせ内容 */}
                                <TextArea
                                    {...register('content')}
                                    error={errors.content?.message}
                                    id="content"
                                    label="お問い合わせ内容"
                                    placeholder="商品についてのご質問、ご要望などをお聞かせください"
                                    required
                                    rows={5}
                                />

                                <div className={styles['submit-container']}>
                                    <Button disabled={isSubmitting || !isValid} type="submit">
                                        {isSubmitting ? '送信中...' : '送信する'}
                                    </Button>
                                </div>
                            </Form>
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
