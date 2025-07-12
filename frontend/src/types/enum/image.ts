// 登録画像の種類
export const ImageType = {
    Registered: 'registered',
    Preview: 'preview',
} as const

export type ImageType = (typeof ImageType)[keyof typeof ImageType]
