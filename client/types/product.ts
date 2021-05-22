import { IAccessoryCategory, IMaterialCategory, ISalesSite } from 'types'

export interface IProductImage {
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
