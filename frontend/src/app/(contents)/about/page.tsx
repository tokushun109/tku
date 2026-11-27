import { Metadata } from 'next'

import { getCreator } from '@/apis/creator'
import { getSalesSiteList } from '@/apis/salesSite'
import { getSnsList } from '@/apis/sns'
import AboutTemplate from '@/app/(contents)/about/template'
import { MERCARI_SHOP_URL } from '@/features/site/constants'

export const dynamic = 'force-dynamic'

export async function generateMetadata(): Promise<Metadata> {
    const title = 'とこりりについて | とこりり'
    const creator = await getCreator()
    const description =
        creator && creator.introduction ? creator.introduction : 'とこりりはハンドメイドのマクラメ編みアクセサリーを制作・販売しているお店です。'
    const image = creator && creator.apiPath ? creator.apiPath : '/image/about/story.jpg'
    return {
        metadataBase: new URL(process.env.DOMAIN_URL || ''),
        title,
        description,
        openGraph: {
            title,
            description,
            type: 'website',
            images: [image],
        },
        twitter: {
            title,
            description,
            images: [image],
        },
    }
}

const About = async () => {
    const snsList = await getSnsList()
    const salesSiteList = await getSalesSiteList()
    const hasMercariShop = salesSiteList.some(({ url }) => url?.replace(/[?#].*$/, '').replace(/\/+$/, '') === MERCARI_SHOP_URL)
    const shopLinks = hasMercariShop ? salesSiteList : [...salesSiteList, { name: 'Mercari Shops', url: MERCARI_SHOP_URL }]

    return <AboutTemplate salesSiteList={shopLinks} snsList={snsList} />
}

export default About
