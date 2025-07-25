import { ISite } from '@/features/site/type'
import { SiteType } from '@/types'

import { SiteList } from '.'

import type { Meta, StoryObj } from '@storybook/nextjs-vite'

const meta: Meta<typeof SiteList> = {
    component: SiteList,
    parameters: {
        layout: 'padded',
    },
    args: {
        siteType: SiteType.Sns,
        sites: [],
    },
    argTypes: {
        siteType: {
            control: { type: 'select' },
            options: [SiteType.Sns, SiteType.SalesSite],
        },
    },
}

export default meta
type Story = StoryObj<typeof SiteList>

// モックデータ
const mockSnsSites: ISite[] = [
    {
        uuid: '1',
        name: 'Instagram',
        url: 'https://www.instagram.com/tku_handmade',
        icon: 'https://via.placeholder.com/24x24/E4405F/ffffff?text=IG',
    },
    {
        uuid: '2',
        name: 'Twitter',
        url: 'https://twitter.com/tku_handmade',
        icon: 'https://via.placeholder.com/24x24/1DA1F2/ffffff?text=TW',
    },
    {
        uuid: '3',
        name: 'YouTube',
        url: 'https://www.youtube.com/channel/example',
        icon: 'https://via.placeholder.com/24x24/FF0000/ffffff?text=YT',
    },
]

const mockSalesSites: ISite[] = [
    {
        uuid: '4',
        name: 'Creema',
        url: 'https://www.creema.jp/creator/1234567',
        icon: 'https://via.placeholder.com/24x24/78C2AD/ffffff?text=CR',
    },
    {
        uuid: '5',
        name: 'minne',
        url: 'https://minne.com/@tku-handmade',
        icon: 'https://via.placeholder.com/24x24/F5A623/ffffff?text=MI',
    },
    {
        uuid: '6',
        name: 'BASE',
        url: 'https://tku-handmade.thebase.in',
        icon: 'https://via.placeholder.com/24x24/5D5D5D/ffffff?text=BA',
    },
]

const mockSiteWithoutUrl: ISite[] = [
    {
        uuid: '7',
        name: 'Coming Soon Site',
        icon: 'https://via.placeholder.com/24x24/CCCCCC/666666?text=CS',
    },
]

export const Default: Story = {}

export const EmptyState: Story = {
    args: {
        sites: [],
        siteType: SiteType.Sns,
    },
}

export const SnsList: Story = {
    args: {
        sites: mockSnsSites,
        siteType: SiteType.Sns,
    },
}

export const SalesSiteList: Story = {
    args: {
        sites: mockSalesSites,
        siteType: SiteType.SalesSite,
    },
}

export const SingleSite: Story = {
    args: {
        sites: [mockSnsSites[0]],
        siteType: SiteType.Sns,
    },
}

export const ManySites: Story = {
    args: {
        sites: [
            ...mockSnsSites,
            ...mockSalesSites,
            {
                uuid: '8',
                name: 'Pinterest',
                url: 'https://www.pinterest.jp/tku_handmade',
                icon: 'https://via.placeholder.com/24x24/BD081C/ffffff?text=PI',
            },
            {
                uuid: '9',
                name: 'Facebook',
                url: 'https://www.facebook.com/tku.handmade',
                icon: 'https://via.placeholder.com/24x24/1877F2/ffffff?text=FB',
            },
        ],
        siteType: SiteType.Sns,
    },
}

export const SiteWithoutUrl: Story = {
    args: {
        sites: mockSiteWithoutUrl,
        siteType: SiteType.Sns,
    },
}

export const MixedSites: Story = {
    args: {
        sites: [...mockSnsSites.slice(0, 2), ...mockSiteWithoutUrl],
        siteType: SiteType.Sns,
    },
}

export const LongUrl: Story = {
    args: {
        sites: [
            {
                uuid: '10',
                name: 'Very Long URL Site',
                url: 'https://www.example-very-long-domain-name.com/path/to/very/long/url/with/many/parameters?param1=value1&param2=value2&param3=value3',
                icon: 'https://via.placeholder.com/24x24/666666/ffffff?text=LU',
            },
        ],
        siteType: SiteType.SalesSite,
    },
}
