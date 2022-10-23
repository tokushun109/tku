export interface IClassificationType {
    [key: string]: { name: string; value: string }
}

export const CategoryType: IClassificationType = {
    Category: { name: 'category', value: 'カテゴリー' },
    Tag: { name: 'tag', value: 'タグ' },
} as const
