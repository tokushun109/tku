import { Metadata } from 'next'

import { getSalesSiteList } from '@/apis/salesSite'
import { getSnsList } from '@/apis/sns'

import { SiteTemplate } from './template'

export const dynamic = 'force-dynamic'

export const metadata: Metadata = {
    title: 'サイト管理 | admin',
    robots: {
        index: false,
        follow: false,
    },
}

const SitePage = async () => {
    try {
        const [snsList, salesSiteList] = await Promise.all([getSnsList(), getSalesSiteList()])

        return <SiteTemplate salesSiteList={salesSiteList} snsList={snsList} />
    } catch (error) {
        console.error('データの取得に失敗しました:', error)
        return <SiteTemplate salesSiteList={[]} snsList={[]} />
    }
}

export default SitePage
