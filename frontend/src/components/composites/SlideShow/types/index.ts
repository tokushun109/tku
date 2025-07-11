export type SwipePosition = {
    end: number | undefined
    start: number | undefined
}

export const ImageIndexEnum = {
    Previous: 'previous',
    Display: 'display',
    Next: 'next',
} as const

export type ImageIndexType = (typeof ImageIndexEnum)[keyof typeof ImageIndexEnum]

export type ImageIndex = {
    [_ in ImageIndexType]: number
}

export type SwipeDirection = 'right' | 'left'
