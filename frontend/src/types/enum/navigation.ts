export const NavigationType = {
    Home: '/',
    About: '/about',
    Product: '/product',
    Contact: '/contact',
} as const
export type NavigationType = (typeof NavigationType)[keyof typeof NavigationType]

export const NavigationTitleEnum = {
    [NavigationType.Home]: 'Home',
    [NavigationType.About]: 'About',
    [NavigationType.Product]: 'Product',
    [NavigationType.Contact]: 'Contact',
} as const
