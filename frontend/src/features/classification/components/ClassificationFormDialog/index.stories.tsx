import { ClassificationType } from '@/types'

import { ClassificationFormDialog } from '.'

import type { Meta, StoryObj } from '@storybook/nextjs-vite'

const meta: Meta<typeof ClassificationFormDialog> = {
    component: ClassificationFormDialog,
    args: {
        isOpen: true,
        isSubmitting: false,
        onClose: () => {
            console.log('ダイアログを閉じます')
        },
        onSubmit: async (data) => {
            console.log('フォーム送信:', data)
            await new Promise((resolve) => setTimeout(resolve, 1000))
        },
        submitError: null,
        classificationType: ClassificationType.Category,
    },
    argTypes: {
        isOpen: {
            control: { type: 'boolean' },
        },
        isSubmitting: {
            control: { type: 'boolean' },
        },
        classificationType: {
            control: { type: 'select' },
            options: Object.values(ClassificationType),
        },
    },
}

export default meta
type Story = StoryObj<typeof ClassificationFormDialog>

export const Category: Story = {
    args: {
        classificationType: ClassificationType.Category,
    },
}

export const Tag: Story = {
    args: {
        classificationType: ClassificationType.Tag,
    },
}

export const Target: Story = {
    args: {
        classificationType: ClassificationType.Target,
    },
}

export const WithError: Story = {
    args: {
        classificationType: ClassificationType.Category,
        submitError: '送信中にエラーが発生しました。もう一度お試しください。',
    },
}

export const Submitting: Story = {
    args: {
        classificationType: ClassificationType.Category,
        isSubmitting: true,
    },
}

export const Closed: Story = {
    args: {
        classificationType: ClassificationType.Category,
        isOpen: false,
    },
}
