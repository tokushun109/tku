'use client'

import { Link, Store } from '@mui/icons-material'
import { useState } from 'react'

import { Tab, TabItem } from '@/components/bases/Tab'
import { ISite } from '@/features/site/type'
import { SiteType, SiteLabel } from '@/types'

import { SiteList } from './SiteList'
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

    return (
        <div className={styles['site-container']}>
            <Tab activeKey={activeTab} items={tabItems} onTabChange={setActiveTab} />
            <div className={styles['tab-content']}>
                {(() => {
                    switch (activeTab) {
                        case SiteType.Sns:
                            return <SiteList siteType={SiteType.Sns} sites={snsList} />
                        case SiteType.SalesSite:
                            return <SiteList siteType={SiteType.SalesSite} sites={salesSiteList} />
                        default:
                            return null
                    }
                })()}
            </div>
        </div>
    )
}
