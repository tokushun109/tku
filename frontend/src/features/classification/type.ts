export interface IClassification {
    name: string
    uuid: string
}

export interface IGetClassificationParams {
    mode: 'all' | 'used'
}

// フォーム用の型を再エクスポート
export type { IClassificationForm } from './classification/type'
