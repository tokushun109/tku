export interface ISite {
    icon: string
    name: string
    url: string
    uuid?: string
}

export interface ISiteDetail {
    detailUrl: string
    salesSite: ISite
    uuid: string
}
