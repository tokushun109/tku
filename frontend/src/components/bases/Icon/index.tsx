'use client'

import classNames from 'classnames'

import { ColorCodeEnum, ColorEnum, ColorType } from '@/types'

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
                text: ColorEnum.White,
            }
        else
            return {
                backGround: ColorEnum.White,
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
                background: ColorCodeEnum[colorObject.backGround],
                color: ColorCodeEnum[colorObject.text],
            }}
        >
            <div className={styles['content']}>{children}</div>
        </div>
    )
}
