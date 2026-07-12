import React from 'react'

import { ColorCode, ColorType } from '@/types'

import styles from './styles.module.scss'

interface Props extends React.ButtonHTMLAttributes<HTMLButtonElement> {
    children: React.ReactNode
    colorType?: ColorType
    contrast?: boolean
    fullWidth?: boolean
    noBoxShadow?: boolean
    outlined?: boolean
    pill?: boolean
}

export const Button = ({
    children,
    colorType = ColorType.Primary,
    contrast = false,
    fullWidth = false,
    noBoxShadow = false,
    outlined = false,
    pill = false,
    disabled,
    className,
    ...props
}: Props) => {
    const getButtonStyles = () => {
        const baseColor = ColorCode[colorType]
        const shapeStyle = {
            ...(pill ? { borderRadius: '999px' } : {}),
            ...(fullWidth ? { width: '100%' } : {}),
        }

        if (!contrast) {
            return {
                backgroundColor: baseColor,
                color: ColorCode[ColorType.White],
                border: outlined ? `2px solid ${ColorCode[ColorType.White]}` : 'none',
                ...shapeStyle,
            }
        } else {
            return {
                backgroundColor: ColorCode[ColorType.White],
                color: baseColor,
                border: outlined ? `2px solid ${baseColor}` : 'none',
                ...shapeStyle,
            }
        }
    }

    return (
        <button
            className={`${styles['container']} ${className || ''} ${noBoxShadow ? styles['no-box-shadow'] : ''}`}
            disabled={disabled}
            style={getButtonStyles()}
            {...props}
        >
            <span>{children}</span>
        </button>
    )
}
