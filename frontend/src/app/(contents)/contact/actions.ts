'use server'

import { redirect } from 'next/navigation'

interface ContactFormData {
    company?: string
    content: string
    email: string
    name: string
    phoneNumber?: string
}

type ContactFormErrors = Partial<Record<keyof ContactFormData, string>>

export interface FormState {
    errors?: ContactFormErrors
    message?: string
    success?: boolean
}

export async function submitContact(prevState: FormState, formData: FormData): Promise<FormState> {
    try {
        // フォームデータの取得
        const data: ContactFormData = {
            name: formData.get('name') as string,
            company: (formData.get('company') as string) || '',
            phoneNumber: (formData.get('phoneNumber') as string) || '',
            email: formData.get('email') as string,
            content: formData.get('content') as string,
        }

        // バリデーション
        const errors: ContactFormErrors = {}

        if (!data.name || data.name.trim().length === 0) {
            errors.name = 'お名前は必須項目です'
        } else if (data.name.length > 20) {
            errors.name = 'お名前は20文字以内で入力してください'
        }

        if (data.company && data.company.length > 20) {
            errors.company = '会社名は20文字以内で入力してください'
        }

        if (data.phoneNumber && !/^\d+$/.test(data.phoneNumber)) {
            errors.phoneNumber = '電話番号は数字のみで入力してください'
        }

        if (!data.email || data.email.trim().length === 0) {
            errors.email = 'メールアドレスは必須項目です'
        } else if (data.email.length > 50) {
            errors.email = 'メールアドレスは50文字以内で入力してください'
        } else if (!/^[^\s@]+@[^\s@]+\.[^\s@]+$/.test(data.email)) {
            errors.email = '正しいメールアドレスを入力してください'
        }

        if (!data.content || data.content.trim().length === 0) {
            errors.content = 'お問い合わせ内容は必須項目です'
        }

        if (Object.keys(errors).length > 0) {
            return {
                errors,
                message: '入力内容を確認してください',
            }
        }

        // APIへの送信（ここではバックエンドAPIを想定）
        const response = await fetch(`${process.env.NEXT_PUBLIC_API_URL || 'http://localhost:8080'}/contact`, {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify(data),
        })

        if (!response.ok) {
            throw new Error('送信に失敗しました')
        }

        // 成功時は3秒後にホームページへリダイレクト
        setTimeout(() => {
            redirect('/')
        }, 3000)

        return {
            success: true,
            message: 'お問い合わせを送信しました',
        }
    } catch (error) {
        console.error('Contact form submission error:', error)
        return {
            message: '送信中にエラーが発生しました。もう一度お試しください。',
        }
    }
}
