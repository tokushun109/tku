'use client'

import classNames from 'classnames'
import { usePathname, useRouter } from 'next/navigation'

import { MenuType, MenuList } from '@/types'
import { NavigationType } from '@/types/enum/navigation'
import { labelFontFace } from '@/utils/font'

import styles from './styles.module.scss'

export const MobileMenu = () => {
    const router = useRouter()
    const pathname = usePathname() as NavigationType

    return (
        <ul className={styles['container']}>
            {Object.values(MenuType).map((menu) => (
                <li
                    className={classNames(styles['menu-list'], pathname === MenuList[menu].link && styles['active'])}
                    key={MenuList[menu].label}
                    onClick={() => {
                        router.push(MenuList[menu].link)
                    }}
                >
                    <div className={styles['icon']}>
                        {(() => {
                            const Icon = MenuList[menu].icon
                            return <Icon />
                        })()}
                    </div>
                    <div className={classNames(styles['label'], labelFontFace.className)}>{MenuList[menu].label}</div>
                </li>
            ))}
        </ul>
    )
}
