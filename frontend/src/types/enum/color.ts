export const ColorEnum = {
    Primary: 'primary',
    Secondary: 'secondary',
    Accent: 'accent',
    White: 'white',
} as const
export type ColorType = (typeof ColorEnum)[keyof typeof ColorEnum]

export const ColorCodeEnum = {
    [ColorEnum.Primary]: '#7b675b',
    [ColorEnum.Secondary]: '#bcaaa4',
    [ColorEnum.Accent]: '#ffb74D',
    [ColorEnum.White]: '#ffffff',
} as const
