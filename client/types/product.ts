import { IClassification, ISiteDetail } from 'types'

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
    siteDetails: Array<ISiteDetail>
    isActive: boolean
}
