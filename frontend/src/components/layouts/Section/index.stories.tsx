import { ColorType } from '@/types'

import Section from '.'

import type { Meta, StoryObj } from '@storybook/nextjs-vite'

const meta: Meta<typeof Section> = {
    component: Section,
    args: {
        title: 'Section Title',
        color: ColorType.White,
        contrast: false,
        button: false,
        children: (
            <div>
                <p>This is the content of the section.</p>
                <p>It can contain any React components or elements.</p>
            </div>
        ),
    },
}

export default meta
type Story = StoryObj<typeof Section>

export const Default: Story = {}

export const WithButton: Story = {
    args: {
        button: true,
        buttonLabel: 'View More',
        onButtonClick: () => {
            console.log('Button clicked')
        },
    },
}

export const PrimaryBackground: Story = {
    args: {
        color: ColorType.Primary,
        contrast: true,
    },
}

export const SecondaryBackground: Story = {
    args: {
        color: ColorType.Secondary,
        contrast: true,
    },
}

export const WithButtonAndPrimary: Story = {
    args: {
        color: ColorType.Primary,
        contrast: true,
        button: true,
        buttonLabel: 'もっと見る',
        onButtonClick: () => {
            console.log('Japanese button clicked')
        },
    },
}

export const LongContent: Story = {
    args: {
        title: 'Long Content Section',
        children: (
            <div>
                <p>This section contains a lot of content to demonstrate how the component handles longer text.</p>
                <p>Lorem ipsum dolor sit amet, consectetur adipiscing elit. Sed do eiusmod tempor incididunt ut labore et dolore magna aliqua.</p>
                <p>Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat.</p>
                <p>Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur.</p>
            </div>
        ),
        button: true,
        buttonLabel: 'Read More',
        onButtonClick: () => {
            console.log('Read more clicked')
        },
    },
}
