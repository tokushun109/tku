'use client'

import { Menu } from '@mui/icons-material'
import classNames from 'classnames'
import { AnimatePresence, motion } from 'framer-motion'
import Image from 'next/image'
import { useRouter } from 'next/navigation'
import { useState } from 'react'

import { Icon } from '@/components/bases/Icon'
import { ColorEnum } from '@/types'

import { slideAnimation } from './animations'
import { MenuScreen } from './components/MenuScreen'
import styles from './styles.module.scss'

export const Header = () => {
    const [isVisibleMenu, setIsVisibleMenu] = useState<boolean>(false)
    const router = useRouter()

    return (
        <div className={styles['container']}>
            <header className={classNames(styles['header'], styles['sm'])}>
                <h1
                    className={styles['logo-area']}
                    onClick={() => {
                        router.push('/')
                    }}
                >
                    <Image
                        alt="とこりり"
                        height={48}
                        priority
                        src="/logo/tocoriri_logo_white.png"
                        style={{
                            objectFit: 'cover',
                        }}
                        width={96}
                    />
                </h1>
            </header>
            {!isVisibleMenu && (
                <div className={classNames(styles['menu-icon'], styles['default'])}>
                    <Icon
                        color={ColorEnum.Primary}
                        onClick={() => {
                            setIsVisibleMenu(true)
                        }}
                        size={72}
                    >
                        <Menu fontSize="large" />
                    </Icon>
                </div>
            )}
            <AnimatePresence>
                {isVisibleMenu && (
                    <motion.div className={classNames(styles['menu-screen'], styles['default'])} {...slideAnimation}>
                        <MenuScreen
                            onCloseClick={() => {
                                setIsVisibleMenu(false)
                            }}
                        />
                    </motion.div>
                )}
            </AnimatePresence>
        </div>
    )
}
