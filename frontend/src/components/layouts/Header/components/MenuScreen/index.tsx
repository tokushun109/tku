'use client'

import { Close } from '@mui/icons-material'
import Image from 'next/image'
import { usePathname, useRouter } from 'next/navigation'

import { Icon } from '@/components/bases/Icon'
import { IconCard } from '@/components/composites/IconCard'
import { MenuEnum, MenuList, ColorEnum, MenuType } from '@/types'
import { NavigationType } from '@/types/enum/navigation'

import styles from './styles.module.scss'

type Props = {
    onCloseClick: () => void
}

export const MenuScreen = ({ onCloseClick }: Props) => {
    const router = useRouter()
    const pathname = usePathname() as NavigationType

    const onClickIcon = (menu: MenuType | '/') => {
        onCloseClick()

        setTimeout(() => {
            if (menu === '/') {
                router.push(menu)
            } else {
                router.push(MenuList[menu].link)
            }
        }, 200)
    }

    return (
        <div className={styles['container']}>
            <div className={styles['close-icon']}>
                <Icon color={ColorEnum.Primary} contrast onClick={onCloseClick} size={72}>
                    <Close fontSize="large" />
                </Icon>
            </div>
            <div
                className={styles['logo-area']}
                onClick={() => {
                    onClickIcon('/')
                }}
            >
                <h1>
                    <Image
                        alt="とこりり"
                        height={200}
                        priority
                        src="/logo/tocoriri_logo_white.png"
                        style={{
                            objectFit: 'cover',
                        }}
                        width={400}
                    />
                </h1>
            </div>
            <div className={styles['icon-card-area']}>
                <div className={styles['icon-card-area-wrapper']}>
                    {Object.values(MenuEnum).map((menu) => (
                        <div
                            className={styles['icon-card']}
                            key={MenuList[menu].label}
                            onClick={() => {
                                onClickIcon(menu)
                            }}
                        >
                            <IconCard Icon={MenuList[menu].icon} isSelected={pathname === MenuList[menu].link} label={MenuList[menu].label} />
                        </div>
                    ))}
                </div>
            </div>
        </div>
    )
}
