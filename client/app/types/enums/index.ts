import { mdiCart, mdiClose, mdiDelete, mdiMenu, mdiNotePlus, mdiPencil, mdiPlus } from '@mdi/js'
export * from '~/types/enums/classification'
export * from '~/types/enums/site'

// 操作の種類
export interface IExecutionType {
    [key: string]: string
}
export const ExecutionType: IExecutionType = {
    Create: '作成',
    Edit: '編集',
    Delete: '削除',
} as const
export type TExecutionType = typeof ExecutionType[keyof typeof ExecutionType]

// アイコンの種類
export interface IIconType {
    [key: string]: { name: string; icon: string }
}
export const IconType: IIconType = {
    New: { name: 'new', icon: mdiNotePlus },
    Edit: { name: 'edit', icon: mdiPencil },
    Delete: { name: 'delete', icon: mdiDelete },
    Close: { name: 'close', icon: mdiClose },
    Cart: { name: 'cart', icon: mdiCart },
    Plus: { name: 'plus', icon: mdiPlus },
    Menu: { name: 'menu', icon: mdiMenu },
} as const
export const ColorType = {
    Red: 'red',
    Orange: 'orange',
    Yellow: 'yellow',
    Green: 'green',
    LightGreen: 'light-green',
    Blue: 'blue',
    Cyan: 'cyan',
    Lime: 'lime',
    Purple: 'purple',
    Black: 'black',
    Grey: 'grey',
    White: 'white',
    Transparent: 'transparent',
} as const
export type TColorType = typeof ColorType[keyof typeof ColorType]

// 登録画像の種類
export const ImageType = {
    Registered: 'registered',
    Preview: 'preview',
} as const

export type TImageType = typeof ImageType[keyof typeof ImageType]
