import { ICreator } from '@/features/creator/type'

import { SeoEdit } from '.'

import type { Meta, StoryObj } from '@storybook/nextjs-vite'

const meta: Meta<typeof SeoEdit> = {
    component: SeoEdit,
    args: {
        creator: {
            apiPath: '/api/images/sample-logo.jpg',
            introduction: 'ハンドメイドアクセサリーの制作を行っています。\n一つ一つ丁寧に手作りしています。',
            logo: 'sample-logo.jpg',
            name: 'tocoriri',
        } as ICreator,
        onUpdate: () => {},
    },
    parameters: {
        layout: 'padded',
    },
}

export default meta
type Story = StoryObj<typeof SeoEdit>

export const Default: Story = {}

export const LongIntroduction: Story = {
    args: {
        creator: {
            apiPath: '/api/images/sample-logo.jpg',
            introduction:
                'ハンドメイドアクセサリーの制作を行っています。\n一つ一つ丁寧に手作りしており、素材にもこだわりを持って選んでいます。\n\nお客様の想いを形にするために、オーダーメイドも承っております。\nぜひお気軽にお問い合わせください。\n\n※配送は全国対応しております。',
            logo: 'sample-logo.jpg',
            name: 'tocoriri ハンドメイドアクセサリー',
        } as ICreator,
    },
}

export const NoImage: Story = {
    args: {
        creator: {
            apiPath: '',
            introduction: 'ハンドメイドアクセサリーの制作を行っています。',
            logo: '',
            name: 'tocoriri',
        } as ICreator,
    },
}
