export interface IClassification {
    uuid: string
    name: string
}

export interface IGetClassificationParams {
    mode: 'all' | 'used'
}
