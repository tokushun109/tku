export const FontSizeType = {
    Xxlarge: 'xxlarge',
    Xlarge: 'xlarge',
    Large: 'large',
    Normal: 'normal',
    Medium: 'medium',
    SmMd: 'smMd',
    Small: 'small',
    Tiny: 'tiny',
} as const
export type FontSizeType = (typeof FontSizeType)[keyof typeof FontSizeType]

export const FontSizeValue = {
    [FontSizeType.Xxlarge]: 24,
    [FontSizeType.Xlarge]: 20,
    [FontSizeType.Large]: 18,
    [FontSizeType.Normal]: 16,
    [FontSizeType.Medium]: 14,
    [FontSizeType.SmMd]: 12,
    [FontSizeType.Small]: 10,
    [FontSizeType.Tiny]: 8,
} as const
