import { ExternalLink } from '.'

import type { Meta, StoryObj } from '@storybook/nextjs-vite'

const meta: Meta<typeof ExternalLink> = {
    component: ExternalLink,
    args: {
        href: 'https://example.com',
        children: 'Visit Example.com',
    },
}

export default meta
type Story = StoryObj<typeof ExternalLink>

export const Default: Story = {}

export const SalesLink: Story = {
    args: {
        href: 'https://creema.jp',
        children: 'Creema で購入',
    },
}

export const WithCustomClass: Story = {
    args: {
        href: 'https://example.com',
        children: 'Styled Link',
        className: 'custom-link-style',
    },
}

export const LongText: Story = {
    args: {
        href: 'https://example.com',
        children: 'This is a longer link text that might wrap to multiple lines',
    },
}

export const JapaneseText: Story = {
    args: {
        href: 'https://example.jp',
        children: 'こちらから購入できます',
    },
}
