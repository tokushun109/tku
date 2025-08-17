'use client'

import { Facebook, Reply, X } from '@mui/icons-material'

import { Icon } from '@/components/bases/Icon'
import { ColorType } from '@/types'

import styles from './styles.module.scss'

interface Props {
    url?: string
}

export const ShareButtons = ({ url }: Props) => {
    const getShareUrl = (): string => {
        if (url) {
            return url
        }
        // クライアントサイドでのみ実行
        if (typeof window !== 'undefined') {
            return window.location.href
        }
        return ''
    }

    const shareUrl = getShareUrl()

    const handleXShare = () => {
        const xShareUrl = `https://twitter.com/share?url=${encodeURIComponent(shareUrl)}`
        window.open(xShareUrl, '_blank', 'nofollow')
    }

    const handleFacebookShare = () => {
        const facebookShareUrl = `https://www.facebook.com/share.php?u=${encodeURIComponent(shareUrl)}`
        window.open(facebookShareUrl, '_blank', 'nofollow,noopener')
    }

    return (
        <div className={styles['container']}>
            <div className={styles['message']}>
                <div className={styles['reply']}>
                    <Reply />
                </div>
                <div>Share This Page♪</div>
            </div>
            <div className={styles['icon-area']}>
                <div>
                    <button aria-label="X(Twitter)でシェア" className={styles['share-button']} onClick={handleXShare} type="button">
                        <Icon color={ColorType.Primary} contrast shadow={false} size={40}>
                            <X />
                        </Icon>
                    </button>
                </div>
                <div>
                    <button aria-label="Facebookでシェア" className={styles['share-button']} onClick={handleFacebookShare} type="button">
                        <Icon color={ColorType.Primary} contrast shadow={false} size={40}>
                            <Facebook />
                        </Icon>
                    </button>
                </div>
            </div>
        </div>
    )
}
