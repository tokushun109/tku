import { IAccessoryCategory, IMaterialCategory, ISalesSite } from 'types'

export interface IProductImage {
    uuid: string
    name: string
    path: string
}
export interface IProduct {
    uuid: string
    name: string
    description: string
    accessoryCategory: IAccessoryCategory | null
    materialCategories: Array<IMaterialCategory>
    productImages: Array<IProductImage>
    salesSites: Array<ISalesSite>
}
