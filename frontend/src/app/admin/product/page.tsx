import { Metadata } from 'next'

import { getCategories } from '@/apis/category'
import { getProducts } from '@/apis/product'
import { getSalesSiteList } from '@/apis/salesSite'
import { getTags } from '@/apis/tag'
import { getTargets } from '@/apis/target'

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
    try {
        const [products, categories, targets, tags, salesSites] = await Promise.all([
            getProducts({
                mode: 'all',
                category: 'all',
                target: 'all',
            }),
            getCategories({ mode: 'all' }),
            getTargets({ mode: 'all' }),
            getTags(),
            getSalesSiteList(),
        ])

        return <AdminProductTemplate categories={categories} initialProducts={products} salesSites={salesSites} tags={tags} targets={targets} />
    } catch (error) {
        console.error('データの取得に失敗しました:', error)
        return <AdminProductTemplate categories={[]} initialProducts={[]} salesSites={[]} tags={[]} targets={[]} />
    }
}

export default AdminProductPage
