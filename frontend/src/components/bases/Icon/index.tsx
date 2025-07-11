'use client'

import classNames from 'classnames'

import { ColorCode, ColorType } from '@/types'

import styles from './styles.module.scss'
import { ColorObject } from './types'

type Props = {
    children: React.ReactNode
    color: ColorType
    contrast?: boolean
    onClick?: () => void
    shadow?: boolean
    size: number
}

export const Icon = ({ color, size, children, onClick = () => {}, contrast = false, shadow = true }: Props) => {
    const colorObject = ((): ColorObject => {
        if (!contrast)
            return {
                backGround: color,
                text: ColorType.White,
            }
        else
            return {
                backGround: ColorType.White,
                text: color,
            }
    })()

    return (
        <div
            className={classNames(styles['container'], !shadow && styles['no-shadow'])}
            onClick={onClick}
            style={{
                width: `${size}px`,
                height: `${size}px`,
                background: ColorCode[colorObject.backGround],
                color: ColorCode[colorObject.text],
            }}
        >
            <div className={styles['content']}>{children}</div>
        </div>
    )
}
