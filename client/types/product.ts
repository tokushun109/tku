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
    category: IClassification | null
    tags: Array<IClassification>
    productImages: Array<IProductImage>
    salesSites: Array<ISite>
}
