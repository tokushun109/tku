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
    accessoryCategory: IAccessoryCategory
    // materialCategory: Array<IMaterialCategory>
    productImage: IProductImage
    // salesSite: Array<ISalesSite>
}
