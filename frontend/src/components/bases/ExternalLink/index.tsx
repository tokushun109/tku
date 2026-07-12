import styles from './styles.module.scss'

type Props = {
    children: React.ReactNode
    href: string
}

export const ExternalLink = ({ children, href }: Props) => {
    const handleClick = (e: React.MouseEvent) => {
        e.stopPropagation()
    }

    return (
        <a className={styles['external-link']} href={href} onClick={handleClick} rel="noopener noreferrer" target="_blank">
            {children}
        </a>
    )
}
