export interface ITableHeader {
    text: string
    align?: string
    sortable?: boolean
    value: string
}

interface IEnabledBreadCrumb {
    text: string
    href: string
}

interface IDisabledBreadCrumb {
    text: string
    disabled: true
}

export type IBreadCrumb = IEnabledBreadCrumb | IDisabledBreadCrumb
