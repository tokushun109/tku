'use client'

import { useActionState } from 'react'
import React from 'react'

import { submitContact } from './actions'
import styles from './styles.module.scss'

interface FormState {
    errors?: Record<string, string>
    message?: string
    success?: boolean
}

const initialState: FormState = {
    message: '',
    errors: {},
}

const ContactPage: React.FC = () => {
    const [state, formAction, pending] = useActionState(submitContact, initialState)

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

                            <form action={formAction}>
                                {/* お名前 */}
                                <div className={styles['require-form']}>
                                    <label htmlFor="name">お名前</label>
                                    <div className={styles['chip-container']}>
                                        <span className={styles['require-chip']}>必須</span>
                                    </div>
                                    <input
                                        className={state?.errors?.name ? styles['error'] : ''}
                                        id="name"
                                        maxLength={20}
                                        name="name"
                                        required
                                        type="text"
                                    />
                                    {state?.errors?.name && <span className={styles['field-error']}>{state.errors.name}</span>}
                                </div>

                                {/* 会社名 */}
                                <div className={styles['form-field']}>
                                    <label htmlFor="company">会社名</label>
                                    <input id="company" maxLength={20} name="company" type="text" />
                                    {state?.errors?.company && <span className={styles['field-error']}>{state.errors.company}</span>}
                                </div>

                                {/* 電話番号 */}
                                <div className={styles['form-field']}>
                                    <label htmlFor="phoneNumber">電話番号(-を入れずに入力)</label>
                                    <input id="phoneNumber" name="phoneNumber" type="tel" />
                                    {state?.errors?.phoneNumber && <span className={styles['field-error']}>{state.errors.phoneNumber}</span>}
                                </div>

                                {/* メールアドレス */}
                                <div className={styles['require-form']}>
                                    <label htmlFor="email">メールアドレス</label>
                                    <div className={styles['chip-container']}>
                                        <span className={styles['require-chip']}>必須</span>
                                    </div>
                                    <input
                                        className={state?.errors?.email ? styles['error'] : ''}
                                        id="email"
                                        maxLength={50}
                                        name="email"
                                        required
                                        type="email"
                                    />
                                    {state?.errors?.email && <span className={styles['field-error']}>{state.errors.email}</span>}
                                </div>

                                {/* お問い合わせ内容 */}
                                <div className={styles['require-form']}>
                                    <label htmlFor="content">お問い合わせ内容</label>
                                    <div className={styles['chip-container']}>
                                        <span className={styles['require-chip']}>必須</span>
                                    </div>
                                    <textarea
                                        className={state?.errors?.content ? styles['error'] : ''}
                                        id="content"
                                        name="content"
                                        required
                                        rows={5}
                                    />
                                    {state?.errors?.content && <span className={styles['field-error']}>{state.errors.content}</span>}
                                </div>

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
