type Props = {
    children: React.ReactNode
    className?: string
    href: string
}

export const ExternalLink = ({ children, className, href }: Props) => {
    return (
        <a className={className} href={href} rel="noopener noreferrer" target="_blank">
            {children}
        </a>
    )
}
