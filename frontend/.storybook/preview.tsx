import 'ress'

import React from 'react'

import { mainFontFace } from '@/utils/font'

import type { Preview } from '@storybook/nextjs-vite'
import '@/styles/globals.scss'

const preview: Preview = {
    parameters: {
        controls: {
            matchers: {
                color: /(background|color)$/i,
                date: /Date$/i,
            },
        },
        a11y: {
            // 'todo' - show a11y violations in the test UI only
            // 'error' - fail CI on a11y violations
            // 'off' - skip a11y checks entirely
            test: 'todo',
        },
    },
    decorators: [
        (Story) => (
            <div className={mainFontFace.className}>
                <Story />
            </div>
        ),
    ],
}

export default preview
