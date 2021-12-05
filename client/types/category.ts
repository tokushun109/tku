export interface ICategory {
    uuid: string
    name: string
}

export interface ICategoryType {
    [key: string]: { name: string; value: string }
}

export const CategoryType: ICategoryType = {
    Accessory: { name: 'accessory', value: 'アクセサリーカテゴリー' },
    Material: { name: 'tag', value: 'タグ' },
} as const
