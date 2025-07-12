import { PageFadeTransition } from '.'

import type { Meta, StoryObj } from '@storybook/nextjs'

const meta: Meta<typeof PageFadeTransition> = {
    component: PageFadeTransition,
    args: {
        children: (
            <div style={{ padding: '40px', textAlign: 'center' }}>
                <h1>Sample Page Content</h1>
                <p>This content will fade in when the page loads.</p>
                <p>Lorem ipsum dolor sit amet, consectetur adipiscing elit.</p>
            </div>
        ),
    },
}

export default meta
type Story = StoryObj<typeof PageFadeTransition>

export const Default: Story = {}

export const WithComplexContent: Story = {
    args: {
        children: (
            <div style={{ padding: '40px' }}>
                <h1>Complex Page Content</h1>
                <section>
                    <h2>Section 1</h2>
                    <p>This is a more complex page with multiple sections.</p>
                </section>
                <section>
                    <h2>Section 2</h2>
                    <ul>
                        <li>Item 1</li>
                        <li>Item 2</li>
                        <li>Item 3</li>
                    </ul>
                </section>
                <section>
                    <h2>Section 3</h2>
                    <p>All content will fade in together.</p>
                </section>
            </div>
        ),
    },
}
