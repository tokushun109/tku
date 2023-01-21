import { IClassification, ISiteDetail, TImageType } from '~/types'

export interface IProductImage {
    uuid: string
    name: string
    apiPath: string
    order: number
}

export interface IImagePathOrder {
    path: string
    order: number | null
    type: TImageType
}

export interface IProduct {
    uuid: string
    name: string
    description: string
    price: number
    category: IClassification
    target: IClassification
    tags: Array<IClassification>
    productImages: Array<IProductImage>
    siteDetails: Array<ISiteDetail>
    isActive: boolean
}

export interface IGetProductsParams {
    mode: 'all' | 'active'
    category: 'all' | string
}

export interface ICarouselItem {
    product: IProduct
    apiPath: string
}
