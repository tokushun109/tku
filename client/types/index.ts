export * from '~/types/error'
export * from '~/types/admin'
export * from '~/types/product'
export * from '~/types/classification'
export * from '~/types/creator'
export * from '~/types/site'
export * from '~/types/initializer'
export * from '~/types/user'

export interface IExecutionType {
    [key: string]: string
}
export const ExecutionType: IExecutionType = {
    Create: '作成',
    Edit: '編集',
    Delete: '削除',
} as const
export type TExecutionType = typeof ExecutionType[keyof typeof ExecutionType]

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
} as const

export type TColorType = typeof ColorType[keyof typeof ColorType]

export const ImageType = {
    Registered: 'registered',
    Preview: 'preview',
} as const

export type TImageType = typeof ImageType[keyof typeof ImageType]
