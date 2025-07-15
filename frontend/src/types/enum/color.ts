export const ColorType = {
    Primary: 'primary',
    Secondary: 'secondary',
    Accent: 'accent',
    White: 'white',
    Danger: 'danger',
} as const
export type ColorType = (typeof ColorType)[keyof typeof ColorType]

export const ColorCode = {
    [ColorType.Primary]: '#7b675b',
    [ColorType.Secondary]: '#bcaaa4',
    [ColorType.Accent]: '#ffb74D',
    [ColorType.White]: '#ffffff',
    [ColorType.Danger]: '#f8d7da',
} as const
