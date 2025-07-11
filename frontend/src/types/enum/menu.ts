import { Email, Diamond, Face3 } from '@mui/icons-material'

import { capitalize } from '@/utils/string'

export const MaterialIconEnum = {
    Face: Face3,
    Diamond: Diamond,
    Email: Email,
} as const
export type MaterialIconType = (typeof MaterialIconEnum)[keyof typeof MaterialIconEnum]

export const MenuEnum = {
    About: 'about',
    Product: 'product',
    Contact: 'contact',
} as const
export type MenuType = (typeof MenuEnum)[keyof typeof MenuEnum]

type MenuObject = {
    icon: MaterialIconType
    label: string
    link: string
}

export const MenuList: { [_ in MenuType]: MenuObject } = {
    [MenuEnum.About]: {
        icon: MaterialIconEnum.Face,
        label: capitalize(MenuEnum.About),
        link: '/' + MenuEnum.About,
    },
    [MenuEnum.Product]: {
        icon: MaterialIconEnum.Diamond,
        label: capitalize(MenuEnum.Product),
        link: '/' + MenuEnum.Product,
    },
    [MenuEnum.Contact]: {
        icon: MaterialIconEnum.Email,
        label: capitalize(MenuEnum.Contact),
        link: '/' + MenuEnum.Contact,
    },
} as const
