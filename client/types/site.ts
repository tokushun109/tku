export interface ISite {
    uuid: string
    name: string
    url: string
    icon: string
}

export interface ISiteDetail {
    uuid: string
    detailUrl: string
    salesSite: ISite
}
