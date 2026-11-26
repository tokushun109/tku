import { afterEach, describe, expect, it, vi } from 'vitest'

import { SOCIAL_LINKS, WORK_LINKS } from '@/app/link/constants'
import LinkPage, { generateMetadata } from '@/app/link/page'
import LinkTemplate from '@/app/link/template'

import { render, screen, within } from '../helpers'

describe('Link Page Integration Test', () => {
    afterEach(() => {
        vi.unstubAllEnvs()
    })

    it('作品の写真と、4件の作品リンクとInstagramを表示する', () => {
        render(<LinkPage />)

        expect(screen.getByRole('heading', { level: 1, name: 'とこりり' })).toBeInTheDocument()
        expect(screen.getByRole('img', { name: '白い石をあしらった手編みのブレスレットを着けた、二人の手元' })).toHaveAttribute(
            'src',
            '/image/link/bracelets.png',
        )

        const workNavigation = screen.getByRole('navigation', { name: '作品を見る' })
        const workLinks = within(workNavigation).getAllByRole('link')
        expect(workLinks.map((link) => [link.textContent, link.getAttribute('href'), link.getAttribute('target')])).toEqual([
            ['Creema', 'https://www.creema.jp/c/tku_accessory/item/onsale', '_blank'],
            ['minne', 'https://minne.com/@tku001', '_blank'],
            ['メルカリShops', 'https://mercari-shops.com/shops/jzQWRyiKK67Q4FR5q75PMZ', '_blank'],
            ['とこりり公式サイト', '/product', null],
        ])

        const socialNavigation = screen.getByRole('navigation', { name: 'SNS' })
        const instagramLink = within(socialNavigation).getByRole('link', { name: 'Instagram' })
        expect(instagramLink).toHaveAttribute('href', 'https://www.instagram.com/tocoriri/')
        expect(within(instagramLink).getByTestId('InstagramIcon')).toBeInTheDocument()
        expect(within(socialNavigation).queryByRole('heading')).not.toBeInTheDocument()
    })

    it('SNSが2件以上ならSNS見出しと各リンクを表示する', () => {
        render(<LinkTemplate socialLinks={[...SOCIAL_LINKS, { href: 'https://example.com/social', name: '追加のSNS' }]} workLinks={WORK_LINKS} />)

        const socialNavigation = screen.getByRole('navigation', { name: 'SNS' })
        expect(within(socialNavigation).getByRole('heading', { name: 'SNS' })).toBeInTheDocument()
        expect(within(socialNavigation).getAllByRole('link')).toHaveLength(2)
        expect(within(socialNavigation).getByRole('link', { name: '追加のSNS' })).toHaveAttribute('href', 'https://example.com/social')
    })

    it.each(['http://localhost:3000', 'https://tocoriri.com'])('設定されたドメイン %s と作品画像で公開メタデータを作る', (domain) => {
        vi.stubEnv('DOMAIN_URL', domain)

        const metadata = generateMetadata()
        expect(metadata.metadataBase?.origin).toBe(domain)
        expect(metadata.alternates?.canonical).toBe('/link')
        expect(metadata.openGraph).toMatchObject({ images: [{ url: '/image/link/bracelets.png' }] })
    })
})
