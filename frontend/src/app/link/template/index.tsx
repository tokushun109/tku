import Image from 'next/image'
import Link from 'next/link'

import { DESCRIPTION, HERO_IMAGE } from '../constants'
import styles from './styles.module.scss'

import type { SocialLink, WorkLink } from '../constants'

type Props = {
    socialLinks: readonly SocialLink[]
    workLinks: readonly WorkLink[]
}

const LinkTemplate = ({ socialLinks, workLinks }: Props) => (
    <main className={styles['page']}>
        <div className={styles['content']}>
            <Image
                alt={HERO_IMAGE.alt}
                className={styles['hero-image']}
                height={HERO_IMAGE.height}
                priority
                sizes="(max-width: 600px) 100vw, (max-width: 1040px) 46vw, 470px"
                src={HERO_IMAGE.src}
                width={HERO_IMAGE.width}
            />
            <div className={styles['details']}>
                <div className={styles['identity']}>
                    <h1 className={styles['brand-name']}>とこりり</h1>
                    <p className={styles['roman-name']} lang="en">
                        tocoriri
                    </p>
                    <p className={styles['description']}>{DESCRIPTION}</p>
                </div>
                <nav aria-labelledby="work-links-title">
                    <h2 className={styles['section-title']} id="work-links-title">
                        作品を見る
                    </h2>
                    <ul className={styles['work-links']}>
                        {workLinks.map(({ external, href, name }) => {
                            const linkProps = { 'aria-label': `${name}の作品一覧`, className: styles['work-link'], href }

                            return (
                                <li key={href}>
                                    {external ? (
                                        <a {...linkProps} rel="noopener" target="_blank">
                                            {name}
                                        </a>
                                    ) : (
                                        <Link {...linkProps} prefetch={false}>
                                            {name}
                                        </Link>
                                    )}
                                </li>
                            )
                        })}
                    </ul>
                </nav>
                {socialLinks.length > 0 ? (
                    <nav aria-label="SNS" className={styles['social']}>
                        {socialLinks.length > 1 ? <h2 className={styles['section-title']}>SNS</h2> : null}
                        <ul className={styles['social-links']}>
                            {socialLinks.map(({ href, icon: Icon, name }) => (
                                <li key={href}>
                                    <a className={styles['social-link']} href={href} rel="noopener" target="_blank">
                                        {Icon ? (
                                            <span aria-hidden="true" className={styles['social-icon']}>
                                                <Icon fontSize="inherit" />
                                            </span>
                                        ) : null}
                                        <span className={styles['social-label']}>{name}</span>
                                    </a>
                                </li>
                            ))}
                        </ul>
                    </nav>
                ) : null}
            </div>
        </div>
    </main>
)

export default LinkTemplate
