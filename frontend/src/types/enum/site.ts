export const SiteType = {
    Sns: 'sns',
    SalesSite: 'salesSite',
} as const
export type SiteType = (typeof SiteType)[keyof typeof SiteType]

export const SiteLabel = {
    [SiteType.Sns]: 'SNS',
    [SiteType.SalesSite]: '販売サイト',
} as const
