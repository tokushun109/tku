import { Metadata } from 'next'

import { getCreator } from '@/apis/creator'
import { getCarouselImages } from '@/apis/product'
import HomeTemplate from '@/app/(home)/template'

export async function generateMetadata(): Promise<Metadata> {
    const title = 'アクセサリーショップ とこりり'
    const creator = await getCreator()
    const description =
        creator && creator.introduction ? creator.introduction : 'とこりりはハンドメイドのマクラメ編みアクセサリーを制作・販売しているお店です。'
    const image = creator && creator.apiPath ? creator.apiPath : '/logo/tocoriri_logo.png'
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

const Home = async () => {
    const carouselImages = await getCarouselImages()
    return <HomeTemplate carouselImages={carouselImages} />
}

export default Home
