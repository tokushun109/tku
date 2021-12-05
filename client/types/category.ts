export interface ICategory {
    uuid: string
    name: string
}

export interface ICategoryType {
    [key: string]: { name: string; value: string }
}

export const CategoryType: ICategoryType = {
    Category: { name: 'category', value: 'カテゴリー' },
    Tag: { name: 'tag', value: 'タグ' },
} as const
