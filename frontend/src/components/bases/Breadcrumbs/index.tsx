'use client'

import classNames from 'classnames'
import { useRouter } from 'next/navigation'

import styles from './styles.module.scss'

export interface IBreadcrumb {
    label: string
    link?: string
}

type Props = {
    breadcrumbs: IBreadcrumb[]
}

export const Breadcrumbs = ({ breadcrumbs }: Props) => {
    const router = useRouter()

    const onClickLabel = (breadcrumb: IBreadcrumb) => {
        if (breadcrumb.link === undefined) return

        router.push(breadcrumb.link)
    }

    return (
        <div className={styles['container']}>
            {breadcrumbs.map((v, index) => (
                <div className={classNames(styles['breadcrumb'], styles[breadcrumbs.length - 1 === index ? 'current-page' : ''])} key={index}>
                    <span
                        className={classNames(styles['breadcrumb__label'], styles[v.link !== undefined ? 'link' : ''])}
                        onClick={() => onClickLabel(v)}
                    >
                        {v.label}
                    </span>
                    <small className={styles['breadcrumb__indicator']}>&gt;</small>
                </div>
            ))}
        </div>
    )
}
