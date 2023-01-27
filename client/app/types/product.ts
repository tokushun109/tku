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
    isRecommend: boolean
    isActive: boolean
}

export interface ICategoryProducts {
    category: IClassification
    products: Array<IProduct>
}

export interface IGetProductsParams {
    mode: 'all' | 'active'
    category: 'all' | string
    target: 'all' | string
}

export interface ICarouselItem {
    product: IProduct
    apiPath: string
}
