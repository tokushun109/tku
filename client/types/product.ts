import { IClassification, ISite } from 'types'

export interface IProductImage {
    uuid: string
    name: string
    apiPath: string
}
export interface IProduct {
    uuid: string
    name: string
    description: string
    price: number
    category: IClassification
    tags: Array<IClassification>
    productImages: Array<IProductImage>
    salesSites: Array<ISite>
    isActive: boolean
}
