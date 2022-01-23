export interface ISite {
    uuid: string
    name: string
    url?: string
}

export interface ISiteDetail {
    uuid: string
    url: string
    salesSite: ISite
}
