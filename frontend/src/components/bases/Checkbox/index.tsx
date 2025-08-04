import React from 'react'

import styles from './styles.module.scss'

interface Props extends Omit<React.InputHTMLAttributes<HTMLInputElement>, 'type'> {
    label?: string
}

export const Checkbox = React.forwardRef<HTMLInputElement, Props>(({ label, className, ...props }, ref) => {
    const inputId = props.id || `checkbox-${Math.random().toString(36).substr(2, 9)}`

    return (
        <div className={`${styles.checkbox} ${className || ''}`}>
            <input className={styles['checkbox-input']} id={inputId} ref={ref} type="checkbox" {...props} />
            {label && (
                <label className={styles['checkbox-label']} htmlFor={inputId}>
                    {label}
                </label>
            )}
        </div>
    )
})

Checkbox.displayName = 'Checkbox'
