import classNames from 'classnames'

import styles from './styles.module.scss'

type Props = {
    children: React.ReactNode
    enabledAnimation?: boolean
    onClick?: () => void
}

export const AnimatedButton = ({ children, onClick = () => {}, enabledAnimation = true }: Props) => {
    return (
        <div className={classNames(styles['container'], !enabledAnimation && styles['no-animation'])} onClick={onClick}>
            <span>{children}</span>
        </div>
    )
}
