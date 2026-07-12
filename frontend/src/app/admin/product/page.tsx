import { Metadata } from 'next'

import { getCategories } from '@/apis/category'
import { ADMIN_PRODUCT_PAGE_LIMIT, getProducts } from '@/apis/product'
import { getSalesSiteList } from '@/apis/salesSite'
import { getTags } from '@/apis/tag'
import { getTargets } from '@/apis/target'
import { IProductList } from '@/features/product/type'

import { AdminProductTemplate } from './template'

export const dynamic = 'force-dynamic'

export const metadata: Metadata = {
    title: '商品一覧 | admin',
    robots: {
        index: false,
        follow: false,
    },
}

const AdminProductPage = async () => {
    const emptyProductList: IProductList = {
        pageInfo: {
            limit: ADMIN_PRODUCT_PAGE_LIMIT,
            page: 1,
            total: 0,
            totalPages: 0,
        },
        products: [],
    }

    try {
        const [productList, categories, targets, tags, salesSites] = await Promise.all([
            getProducts({
                mode: 'all',
                category: 'all',
                limit: ADMIN_PRODUCT_PAGE_LIMIT,
                page: 1,
                target: 'all',
            }),
            getCategories({ mode: 'all' }),
            getTargets({ mode: 'all' }),
            getTags(),
            getSalesSiteList(),
        ])

        return <AdminProductTemplate categories={categories} initialProductList={productList} salesSites={salesSites} tags={tags} targets={targets} />
    } catch (error) {
        console.error('データの取得に失敗しました:', error)
        return <AdminProductTemplate categories={[]} initialProductList={emptyProductList} salesSites={[]} tags={[]} targets={[]} />
    }
}

export default AdminProductPage
