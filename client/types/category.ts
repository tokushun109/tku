export interface ICategory {
    uuid: string
    name: string
}

export interface ICategoryModelValidation {
    name: boolean
}

export enum CategoryType {
    Accessory = 'accessoryCategory',
    Material = 'materialCategory',
}
