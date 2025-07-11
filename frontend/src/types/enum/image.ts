// 登録画像の種類
export const ImageEnum = {
    Registered: 'registered',
    Preview: 'preview',
} as const

export type ImageType = (typeof ImageEnum)[keyof typeof ImageEnum]
