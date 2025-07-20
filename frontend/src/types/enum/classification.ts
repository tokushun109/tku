export const ClassificationType = {
    Category: 'category',
    Target: 'target',
    Tag: 'tag',
} as const
export type ClassificationType = (typeof ClassificationType)[keyof typeof ClassificationType]

export const ClassificationLabel = {
    [ClassificationType.Category]: 'カテゴリー',
    [ClassificationType.Target]: 'ターゲット',
    [ClassificationType.Tag]: 'タグ',
} as const
