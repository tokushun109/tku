import { Metadata } from 'next'
import { notFound } from 'next/navigation'

import { getProduct } from '@/apis/product'
import ProductDetailTemplate from '@/app/(contents)/product/[uuid]/template'

interface Props {
    params: Promise<{ uuid: string }>
}

export async function generateMetadata({ params }: Props): Promise<Metadata> {
    try {
        const { uuid } = await params
        const product = await getProduct(uuid)
        const title = `${product.name} | とこりり`
        const description = product.description.replace(/\r?\n/g, '')
        const image = product.productImages.length ? product.productImages[0].apiPath : '/image/about/story.jpg'

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
    } catch {
        return {
            title: '商品詳細 | とこりり',
            description: 'とこりりの商品詳細ページです。',
        }
    }
}

const ProductDetail = async ({ params }: Props) => {
    const { uuid } = await params

    try {
        const product = await getProduct(uuid)
        if (!product) {
            notFound()
        }
        return <ProductDetailTemplate product={product} />
    } catch {
        notFound()
    }
}

export default ProductDetail
