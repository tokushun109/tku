export interface IClassification {
    uuid: string
    name: string
}

export interface IGetCategoriesParams {
    mode: 'all' | 'used'
}
