export * from '~/types/error'
export * from '~/types/admin'
export * from '~/types/product'
export * from '~/types/category'
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
