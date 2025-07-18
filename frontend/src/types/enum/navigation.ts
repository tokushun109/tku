import { AccountCircle, Description, Email, ShoppingCart, Tag, Web } from '@mui/icons-material'

export const NavigationType = {
    Home: '/',
    About: '/about',
    Product: '/product',
    Contact: '/contact',
    AdminSeo: '/admin/seo',
    AdminProduct: '/admin/product',
    AdminClassification: '/admin/classification',
    AdminSite: '/admin/site',
    AdminContact: '/admin/contact',
    AdminCsv: '/admin/csv',
} as const
export type NavigationType = (typeof NavigationType)[keyof typeof NavigationType]

export const NavigationTitleEnum = {
    [NavigationType.Home]: 'Home',
    [NavigationType.About]: 'About',
    [NavigationType.Product]: 'Product',
    [NavigationType.Contact]: 'Contact',
    [NavigationType.AdminSeo]: 'SEO',
    [NavigationType.AdminProduct]: '商品',
    [NavigationType.AdminClassification]: '分類',
    [NavigationType.AdminSite]: 'サイト',
    [NavigationType.AdminContact]: 'お問い合わせ',
    [NavigationType.AdminCsv]: 'CSV',
} as const

export interface NavigationItem {
    icon: React.ComponentType
    link: NavigationType
    name: string
}

export const NavigationItems: NavigationItem[] = [
    { icon: AccountCircle, link: NavigationType.AdminSeo, name: NavigationTitleEnum[NavigationType.AdminSeo] },
    { icon: ShoppingCart, link: NavigationType.AdminProduct, name: NavigationTitleEnum[NavigationType.AdminProduct] },
    { icon: Tag, link: NavigationType.AdminClassification, name: NavigationTitleEnum[NavigationType.AdminClassification] },
    { icon: Web, link: NavigationType.AdminSite, name: NavigationTitleEnum[NavigationType.AdminSite] },
    { icon: Email, link: NavigationType.AdminContact, name: NavigationTitleEnum[NavigationType.AdminContact] },
    { icon: Description, link: NavigationType.AdminCsv, name: NavigationTitleEnum[NavigationType.AdminCsv] },
]
