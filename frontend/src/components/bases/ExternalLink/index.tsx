type Props = {
    children: React.ReactNode
    className?: string
    href: string
}

export const ExternalLink = ({ children, className, href }: Props) => {
    const handleClick = (e: React.MouseEvent) => {
        e.stopPropagation()
    }

    return (
        <a className={className} href={href} onClick={handleClick} rel="noopener noreferrer" target="_blank">
            {children}
        </a>
    )
}
