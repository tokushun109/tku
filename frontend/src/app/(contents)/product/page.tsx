import { Metadata } from 'next'

import { getCategories } from '@/apis/category'
import { getProductsByCategory } from '@/apis/product'
import { getTargets } from '@/apis/target'
import ProductTemplate from '@/app/(contents)/product/template'
import { INITIAL_CATEGORY_PRODUCT_LIMIT } from '@/features/product/constants'

export const dynamic = 'force-dynamic'

export async function generateMetadata(): Promise<Metadata> {
    const title = '商品一覧 | とこりり'
    const description = 'とこりりの商品一覧ページです。'
    const image = '/image/about/story.jpg'
    return {
        metadataBase: new URL(process.env.DOMAIN_URL || ''),
        title,
        description,
        openGraph: {
            title,
            description,
            type: 'article',
            images: [image],
        },
        twitter: {
            title,
            description,
            images: [image],
        },
    }
}

const Product = async () => {
    const productsByCategory = await getProductsByCategory({
        category: 'all',
        limit: INITIAL_CATEGORY_PRODUCT_LIMIT,
        target: 'all',
    })
    const categories = await getCategories({ mode: 'used' })
    const targets = await getTargets({ mode: 'used' })
    return <ProductTemplate categories={categories} productsByCategory={productsByCategory} targets={targets} />
}

export default Product
