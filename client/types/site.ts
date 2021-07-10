export interface ISite {
    uuid: string
    name: string
    url: string
}

export interface ISiteModelValidation {
    name: boolean
    url: boolean
}

export enum SiteType {
    Sns = 'sns',
    SkillMarket = 'skillMarket',
    SalesSite = 'salesSite',
}
