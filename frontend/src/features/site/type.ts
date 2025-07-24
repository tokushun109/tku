import z from 'zod'

import { SiteSchema } from './schema'

export interface ISite {
    icon?: string
    name: string
    url?: string
    uuid?: string
}

export interface ISiteDetail {
    detailUrl: string
    salesSite: ISite
    uuid: string
}

/** サイトフォームのデータ型 */
export type ISiteForm = z.infer<typeof SiteSchema>

export interface IGetSiteParams {
    mode: 'all' | 'used'
}
