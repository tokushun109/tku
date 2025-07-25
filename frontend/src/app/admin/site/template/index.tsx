'use client'

import { Link, Store } from '@mui/icons-material'
import { useState } from 'react'

import { Tab, TabItem } from '@/components/bases/Tab'
import { ISite } from '@/features/site/type'
import { SiteType, SiteLabel } from '@/types'

import styles from './styles.module.scss'

interface Props {
    salesSiteList: ISite[]
    snsList: ISite[]
}

export const SiteTemplate = ({ salesSiteList, snsList }: Props) => {
    const [activeTab, setActiveTab] = useState<SiteType>(SiteType.Sns)

    const tabItems: TabItem<SiteType>[] = [
        {
            key: SiteType.Sns,
            label: SiteLabel[SiteType.Sns],
            icon: <Link />,
        },
        {
            key: SiteType.SalesSite,
            label: SiteLabel[SiteType.SalesSite],
            icon: <Store />,
        },
    ]

    const renderSiteList = (sites: ISite[], siteType: SiteType) => (
        <div className={styles['site-list-container']}>
            <div className={styles['site-list-header']}>
                <h2 className={styles['site-list-title']}>{SiteLabel[siteType]}一覧</h2>
                <div className={styles['site-count']}>
                    <span>総件数: {sites.length}件</span>
                </div>
            </div>
            <div className={styles['site-list']}>
                {sites.length === 0 ? (
                    <div className={styles['empty-state']}>
                        <p>データがありません</p>
                    </div>
                ) : (
                    sites.map((site) => (
                        <div className={styles['site-item']} key={site.uuid}>
                            <div className={styles['site-info']}>
                                <div className={styles['site-header']}>
                                    <h3 className={styles['site-name']}>{site.name}</h3>
                                </div>
                                <div className={styles['site-details']}>
                                    <div className={styles['site-url']}>
                                        <strong>URL:</strong>
                                        <a className={styles['site-link']} href={site.url} rel="noopener noreferrer" target="_blank">
                                            {site.url}
                                        </a>
                                    </div>
                                </div>
                            </div>
                        </div>
                    ))
                )}
            </div>
        </div>
    )

    return (
        <div className={styles['site-container']}>
            <Tab activeKey={activeTab} items={tabItems} onTabChange={setActiveTab} />
            <div className={styles['tab-content']}>
                {(() => {
                    switch (activeTab) {
                        case SiteType.Sns:
                            return renderSiteList(snsList, SiteType.Sns)
                        case SiteType.SalesSite:
                            return renderSiteList(salesSiteList, SiteType.SalesSite)
                        default:
                            return null
                    }
                })()}
            </div>
        </div>
    )
}
