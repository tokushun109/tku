export interface IClassificationType {
    [key: string]: { name: string; value: string }
}

export const ClassificationType: IClassificationType = {
    Category: { name: 'category', value: 'カテゴリー' },
    Target: { name: 'target', value: 'ターゲット' },
    Tag: { name: 'tag', value: 'タグ' },
} as const
