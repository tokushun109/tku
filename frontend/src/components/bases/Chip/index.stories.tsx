import { ColorType, FontSizeType } from '@/types'

import { Chip } from '.'

import type { Meta, StoryObj } from '@storybook/nextjs'

const meta: Meta<typeof Chip> = {
    component: Chip,
    args: {
        children: 'Sample Chip',
        color: ColorType.Primary,
        fontColor: '#ffffff',
        fontSize: FontSizeType.Medium,
    },
}

export default meta
type Story = StoryObj<typeof Chip>

export const Default: Story = {}

export const Primary: Story = {
    args: {
        color: ColorType.Primary,
    },
}

export const Secondary: Story = {
    args: {
        color: ColorType.Secondary,
    },
}

export const Accent: Story = {
    args: {
        color: ColorType.Accent,
    },
}

export const SmallSize: Story = {
    args: {
        fontSize: FontSizeType.Small,
        children: 'Small Chip',
    },
}

export const LargeSize: Story = {
    args: {
        fontSize: FontSizeType.Large,
        children: 'Large Chip',
    },
}

export const CustomFontColor: Story = {
    args: {
        fontColor: '#000000',
        color: ColorType.Secondary,
        children: 'Dark Text',
    },
}

export const CategoryChip: Story = {
    args: {
        children: 'アクセサリー',
        color: ColorType.Secondary,
        fontSize: FontSizeType.SmMd,
    },
}

export const TagChip: Story = {
    args: {
        children: 'ハンドメイド',
        color: ColorType.Primary,
        fontSize: FontSizeType.Medium,
    },
}
