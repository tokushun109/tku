import classNames from 'classnames'
import { useState, MouseEvent } from 'react'

import { doRippleAnimation, RippleColorEnum } from '@/utils/animation'

import styles from './styles.module.scss'

export type SelectOption<T = string> = {
    label: string
    value: T
}

type Props<T = string> = {
    initialSelectedOption?: SelectOption<T>
    isSelectedAll?: boolean
    onSelect: (_option: SelectOption<T> | undefined) => void
    options: SelectOption<T>[]
    suffix?: React.ReactNode
    title: string
}

export const Select = <T,>({ title, options, initialSelectedOption, isSelectedAll = true, suffix, onSelect }: Props<T>) => {
    // プルダウンが開いているか
    const [isOpen, setIsOpen] = useState<boolean>(false)

    // 選択されているオプションのインデックス
    const [selectedOption, setSelectedOption] = useState<SelectOption<T> | undefined>(initialSelectedOption)

    // 表示されるタイトル
    const displayTitle = ((): string => {
        if (selectedOption === undefined) {
            return isSelectedAll ? `${title} - All` : title
        } else {
            return `${title} - ${selectedOption.label}`
        }
    })()

    const onClickTitle = (e: MouseEvent) => {
        doRippleAnimation(e, RippleColorEnum.Black)
        setIsOpen(!isOpen)
    }

    const onClickOption = (e: MouseEvent, selectedOption: SelectOption<T> | undefined) => {
        doRippleAnimation(e, RippleColorEnum.Orange)
        setSelectedOption(selectedOption)
        onSelect(selectedOption)

        // 2秒後にオプションを閉じる
        setTimeout(() => {
            setIsOpen(false)
        }, 200)
    }

    return (
        <div className={styles['container']}>
            <div className={styles['container__title']} onClick={onClickTitle}>
                {suffix && <span className={classNames(styles['container__title__suffix'], styles['active'])}>{suffix}</span>}
                {displayTitle}
            </div>
            <div className={classNames(styles['container__options'], isOpen ? styles['visible'] : '')}>
                <ul className={styles['container__options__inner']}>
                    {isSelectedAll && (
                        <li
                            className={classNames(
                                styles['container__option'],
                                suffix !== undefined && styles[`suffix-padding`],
                                isSelectedAll && selectedOption === undefined && styles['active'],
                            )}
                            onClick={(e) => {
                                onClickOption(e, undefined)
                            }}
                        >
                            All
                        </li>
                    )}
                    {options.map((v, i) => (
                        <li
                            className={classNames(
                                styles['container__option'],
                                suffix !== undefined && styles[`suffix-padding`],
                                selectedOption === v && styles['active'],
                            )}
                            key={`${i}-${v.label}`}
                            onClick={(e) => {
                                onClickOption(e, v)
                            }}
                        >
                            {v.label}
                        </li>
                    ))}
                </ul>
            </div>
        </div>
    )
}
