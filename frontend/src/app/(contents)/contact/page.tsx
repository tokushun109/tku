import { Metadata } from 'next'

import ContactTemplate from '@/app/(contents)/contact/template'

export async function generateMetadata(): Promise<Metadata> {
    const title = 'お問い合わせ | とこりり'
    const description = 'とこりりへのお問い合わせ・ご意見・ご相談はこちらから。ハンドメイドアクセサリーに関するご質問やご要望をお聞かせください。'
    const image = '/logo/tocoriri_logo.png'
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

const Contact = () => {
    return <ContactTemplate />
}

export default Contact
