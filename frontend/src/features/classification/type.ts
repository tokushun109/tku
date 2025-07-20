export interface IClassification {
    name: string
    uuid: string
}

export interface IGetClassificationParams {
    mode: 'all' | 'used'
}
