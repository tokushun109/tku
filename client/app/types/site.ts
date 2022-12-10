export interface ISite {
    name: string
    url: string
    icon: string
    uuid?: string
}

export interface ISiteDetail {
    uuid: string
    detailUrl: string
    salesSite: ISite
}
