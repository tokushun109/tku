export interface ISite {
    uuid: string
    name: string
    url: string
}

export interface ISiteType {
    [key: string]: { name: string; value: string }
}

export const SiteType: ISiteType = {
    Sns: { name: 'sns', value: 'SNS' },
    SalesSite: { name: 'salesSite', value: '販売サイト' },
    SkillMarket: { name: 'skillMarket', value: 'スキルマーケット' },
} as const
