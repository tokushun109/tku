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
    New: { name: 'new', icon: 'mdi-note-plus' },
    Edit: { name: 'edit', icon: 'mdi-pencil' },
    Delete: { name: 'delete', icon: 'mdi-delete' },
    Close: { name: 'close', icon: 'mdi-close' },
    Cart: { name: 'cart', icon: 'mdi-cart' },
    Plus: { name: 'plus', icon: 'mdi-plus' },
    Menu: { name: 'menu', icon: 'mdi-menu' },
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
