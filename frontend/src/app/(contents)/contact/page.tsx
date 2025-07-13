'use client'

import { useActionState } from 'react'
import React from 'react'

import { Input } from '@/components/bases/Input'
import { TextArea } from '@/components/bases/TextArea'

import { submitContact, FormState } from './actions'
import styles from './styles.module.scss'

const initialState: FormState = {
    message: '',
    errors: {},
}

const ContactPage: React.FC = () => {
    const [state, formAction, pending] = useActionState(submitContact, initialState)

    const handleSubmit = (event: React.FormEvent<HTMLFormElement>) => {
        event.preventDefault()
        const formData = new FormData(event.currentTarget)
        formAction(formData)
    }

    return (
        <div className={styles['page-contact']}>
            {!state?.success && (
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
                    {!state?.success ? (
                        <div className={styles['form-container']}>
                            {state?.message && (
                                <div aria-live="polite" className={styles['error-message']}>
                                    {state.message}
                                </div>
                            )}

                            <form noValidate onSubmit={handleSubmit}>
                                {/* お名前 */}
                                <Input error={state?.errors?.name} id="name" label="お名前" maxLength={20} name="name" required type="text" />

                                {/* 会社名 */}
                                <Input error={state?.errors?.company} id="company" label="会社名" maxLength={20} name="company" type="text" />

                                {/* 電話番号 */}
                                <Input
                                    error={state?.errors?.phoneNumber}
                                    id="phoneNumber"
                                    label="電話番号(-を入れずに入力)"
                                    name="phoneNumber"
                                    type="tel"
                                />

                                {/* メールアドレス */}
                                <Input
                                    error={state?.errors?.email}
                                    id="email"
                                    label="メールアドレス"
                                    maxLength={50}
                                    name="email"
                                    required
                                    type="email"
                                />

                                {/* お問い合わせ内容 */}
                                <TextArea error={state?.errors?.content} id="content" label="お問い合わせ内容" name="content" required rows={5} />

                                <div className={styles['submit-container']}>
                                    <button className={styles['submit-button']} disabled={pending} type="submit">
                                        {pending ? '送信中...' : '送信する'}
                                    </button>
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
