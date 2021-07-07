export interface ICategory {
    uuid: string
    name: string
}

export interface ICategoryModelValidation {
    name: boolean
}

export enum CategoryType {
    Accessory = 'accessory_category',
    Material = 'material_category',
}
