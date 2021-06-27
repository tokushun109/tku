import { ICategory, ISite } from 'types'

export interface IProductImage {
    uuid: string
    name: string
    apiPath: string
}
export interface IProduct {
    uuid: string
    name: string
    description: string
    accessoryCategory: ICategory | null
    materialCategories: Array<ICategory>
    productImages: Array<IProductImage>
    salesSites: Array<ISite>
}
