import { GoogleAnalytics } from '@next/third-parties/google'

import { DESCRIPTION, HERO_IMAGE, SOCIAL_LINKS, WORK_LINKS } from './constants'
import LinkTemplate from './template'

import type { Metadata } from 'next'

export function generateMetadata(): Metadata {
    const title = 'リンク集 | とこりり'
    const description = `${DESCRIPTION}「とこりり」の作品一覧、ショップ、SNSへのリンク集です。`
    const image = {
        alt: HERO_IMAGE.alt,
        height: HERO_IMAGE.height,
        url: HERO_IMAGE.src,
        width: HERO_IMAGE.width,
    }

    return {
        alternates: { canonical: '/link' },
        description,
        metadataBase: new URL(process.env.DOMAIN_URL || ''),
        openGraph: { description, images: [image], title, type: 'website', url: '/link' },
        title,
        twitter: { card: 'summary_large_image', description, images: [image], title },
    }
}

const LinkPage = () => {
    const googleTag = process.env.GOOGLE_TAG

    return (
        <>
            <LinkTemplate socialLinks={SOCIAL_LINKS} workLinks={WORK_LINKS} />
            {googleTag ? <GoogleAnalytics gaId={googleTag} /> : null}
        </>
    )
}

export default LinkPage
