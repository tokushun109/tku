import 'ress'

import React from 'react'

import { mainFontFace } from '@/utils/font'

import type { Preview } from '@storybook/react'
import '@/styles/globals.scss'

const preview: Preview = {
    parameters: {
        actions: { argTypesRegex: '^on[A-Z].*' },
        controls: {
            matchers: {
                color: /(background|color)$/i,
                date: /Date$/i,
            },
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
