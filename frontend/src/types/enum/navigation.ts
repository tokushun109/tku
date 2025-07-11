export const NavigationEnum = {
    Home: '/',
    About: '/about',
    Product: '/product',
    Contact: '/contact',
} as const
export type NavigationType = (typeof NavigationEnum)[keyof typeof NavigationEnum]

export const NavigationTitleEnum = {
    [NavigationEnum.Home]: 'Home',
    [NavigationEnum.About]: 'About',
    [NavigationEnum.Product]: 'Product',
    [NavigationEnum.Contact]: 'Contact',
} as const
