import { Button } from '@/components/bases/Button'
import { Input } from '@/components/bases/Input'
import { TextArea } from '@/components/bases/TextArea'
import { ColorType } from '@/types/enum/color'

import { Form } from '.'

import type { Meta, StoryObj } from '@storybook/nextjs-vite'

const meta: Meta<typeof Form> = {
    component: Form,
    args: {
        noValidate: true,
        onSubmit: (e) => {
            e.preventDefault()
            console.log('フォーム送信')
        },
        children: (
            <>
                <Input label="名前" placeholder="名前を入力してください" required />
                <Input label="メールアドレス" placeholder="メールアドレスを入力してください" type="email" />
                <TextArea label="メッセージ" placeholder="メッセージを入力してください" required />
                <Button colorType={ColorType.Primary} type="submit">
                    送信
                </Button>
            </>
        ),
    },
    argTypes: {
        noValidate: {
            control: { type: 'boolean' },
        },
    },
}

export default meta
type Story = StoryObj<typeof Form>

export const Default: Story = {}

export const WithError: Story = {
    args: {
        children: (
            <>
                <Input error="名前は必須項目です" label="名前" placeholder="名前を入力してください" required />
                <Input label="メールアドレス" placeholder="メールアドレスを入力してください" type="email" />
                <TextArea label="メッセージ" placeholder="メッセージを入力してください" required />
                <Button colorType={ColorType.Primary} type="submit">
                    送信
                </Button>
            </>
        ),
    },
}

export const LoginForm: Story = {
    args: {
        children: (
            <>
                <Input label="メールアドレス" placeholder="example@example.com" required type="email" />
                <Input label="パスワード" required type="password" />
                <Button colorType={ColorType.Primary} type="submit">
                    ログイン
                </Button>
            </>
        ),
    },
}

export const ContactForm: Story = {
    args: {
        children: (
            <>
                <Input label="お名前" placeholder="山田太郎" required />
                <Input label="会社名" placeholder="株式会社サンプル" />
                <Input label="電話番号" placeholder="09012345678" type="tel" />
                <Input label="メールアドレス" placeholder="example@example.com" required type="email" />
                <TextArea label="お問い合わせ内容" placeholder="商品についてのご質問をお聞かせください" required rows={5} />
                <Button colorType={ColorType.Primary} type="submit">
                    送信する
                </Button>
            </>
        ),
    },
}
