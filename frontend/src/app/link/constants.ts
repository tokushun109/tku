import Instagram from '@mui/icons-material/Instagram'

export const DESCRIPTION = '天然石と手編みのアクセサリー'

export const HERO_IMAGE = {
    alt: '白い石をあしらった手編みのブレスレットを着けた、二人の手元',
    height: 583,
    src: '/image/link/bracelets.png',
    width: 583,
}

export type WorkLink = {
    external: boolean
    href: string
    name: string
}

export type SocialLink = {
    href: string
    icon?: typeof Instagram
    name: string
}

export const WORK_LINKS: readonly WorkLink[] = [
    { external: true, href: 'https://www.creema.jp/c/tku_accessory/item/onsale', name: 'Creema' },
    { external: true, href: 'https://minne.com/@tku001', name: 'minne' },
    { external: true, href: 'https://mercari-shops.com/shops/jzQWRyiKK67Q4FR5q75PMZ', name: 'メルカリShops' },
    { external: false, href: '/product', name: 'とこりり公式サイト' },
]

export const SOCIAL_LINKS: readonly SocialLink[] = [{ href: 'https://www.instagram.com/tocoriri/', icon: Instagram, name: 'Instagram' }]
