import React from 'react'

import { ColorCode, ColorType } from '@/types'

import styles from './styles.module.scss'

interface Props extends React.ButtonHTMLAttributes<HTMLButtonElement> {
    children: React.ReactNode
    colorType?: ColorType
}

export const Button = ({ children, colorType = ColorType.Primary, disabled, className, ...props }: Props) => {
    return (
        <button
            className={`${styles['container']} ${className || ''}`}
            disabled={disabled}
            style={{ backgroundColor: ColorCode[colorType] }}
            {...props}
        >
            <span>{children}</span>
        </button>
    )
}
