import { KeyboardArrowDown } from '@mui/icons-material'

import { Select } from '.'

import type { Meta, StoryObj } from '@storybook/nextjs-vite'

const meta: Meta<typeof Select> = {
    component: Select,
    args: {
        title: 'title',
        options: [1, 2, 3].map((i) => ({ value: `value${i}`, label: `label${i}` })),
        initialSelectedOption: { value: 'value1', label: 'label1' },
        suffix: <KeyboardArrowDown />,
        onSelect: (option) => {
            console.log(option === undefined ? '選択を解除しました' : `${option?.value}を選択しました`)
        },
    },
}

export default meta
type Story = StoryObj<typeof Select>

export const Default: Story = {}
