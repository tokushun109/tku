import { Diamond, Email, Face3 } from '@mui/icons-material'

import { capitalize } from '@/utils/string'

export const MaterialIconType = {
    Face: Face3,
    Diamond: Diamond,
    Email: Email,
} as const
export type MaterialIconType = (typeof MaterialIconType)[keyof typeof MaterialIconType]

export const MenuType = {
    About: 'about',
    Product: 'product',
    Contact: 'contact',
} as const
export type MenuType = (typeof MenuType)[keyof typeof MenuType]

type MenuObject = {
    icon: MaterialIconType
    label: string
    link: string
}

export const MenuList: { [_ in MenuType]: MenuObject } = {
    [MenuType.About]: {
        icon: MaterialIconType.Face,
        label: capitalize(MenuType.About),
        link: '/' + MenuType.About,
    },
    [MenuType.Product]: {
        icon: MaterialIconType.Diamond,
        label: capitalize(MenuType.Product),
        link: '/' + MenuType.Product,
    },
    [MenuType.Contact]: {
        icon: MaterialIconType.Email,
        label: capitalize(MenuType.Contact),
        link: '/' + MenuType.Contact,
    },
} as const
