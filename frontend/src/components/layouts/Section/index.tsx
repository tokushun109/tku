import classNames from 'classnames'

import { Slide } from '@/components/animations/Slide'
import { AnimatedButton } from '@/components/bases/AnimatedButton'
import { ColorCode, ColorType } from '@/types'
import { labelFontFace } from '@/utils/font'

import styles from './styles.module.scss'

// 色も指定する

type BaseProps = {
    children: React.ReactNode
    color: ColorType
    contrast?: boolean
    title: string
}

type Props =
    | (BaseProps & {
          button: false
      })
    | (BaseProps & {
          button: true
          buttonLabel: string
          onButtonClick: () => void
      })

const Section = (props: Props) => {
    const { button } = props
    if (button) {
        // ボタンがある時
        const { title, contrast, children, buttonLabel, color, onButtonClick } = props
        return (
            <div className={classNames(styles['container'], contrast && styles['contrast'])} style={contrast ? { background: ColorCode[color] } : {}}>
                <Slide>
                    <div
                        className={classNames(styles['title'], labelFontFace.className)}
                        style={{ color: ColorCode[!contrast ? color : ColorType.White] }}
                    >
                        {title}
                    </div>
                </Slide>
                <Slide>
                    <div className={styles['sentence']} style={{ color: !contrast ? '#757575' : ColorCode[ColorType.White] }}>
                        {children}
                    </div>
                </Slide>
                <Slide>
                    <div className={styles['button']}>
                        <AnimatedButton onClick={onButtonClick}>{buttonLabel}</AnimatedButton>
                    </div>
                </Slide>
            </div>
        )
    } else {
        // ボタンがない時
        const { title, contrast, children, color } = props
        return (
            <div className={classNames(styles['container'], contrast && styles['contrast'])}>
                <Slide>
                    <div className={classNames(styles['title'], labelFontFace.className)} style={{ color: ColorCode[color] }}>
                        {title}
                    </div>
                </Slide>
                <Slide>
                    <div className={styles['sentence']}>{children}</div>
                </Slide>
            </div>
        )
    }
}

export default Section
