import { ImageType } from '@/types'

import { IClassification } from '../classification/type'
import { ISiteDetail } from '../site/type'

export interface IProductImage {
    apiPath: string
    name: string
    order: number
    uuid: string
}

export interface IImagePathOrder {
    order: number | null
    path: string
    type: ImageType
}

export interface IProduct {
    category: IClassification
    description: string
    isActive: boolean
    isRecommend: boolean
    name: string
    price: number
    productImages: IProductImage[]
    siteDetails: ISiteDetail[]
    tags: IClassification[]
    target: IClassification
    uuid: string
}

export interface IProductsByCategory {
    category: IClassification
    products: IProduct[]
}

export interface IThumbnail {
    apiPath: string
    product: IProduct
}

// フォーム用の型を再エクスポート
export type { IProductForm } from './product/type'
