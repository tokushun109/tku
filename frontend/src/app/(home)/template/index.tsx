'use client'

import classNames from 'classnames'
import Image from 'next/image'
import { useRouter } from 'next/navigation'

import { Slide } from '@/components/animations/Slide'
import { Indicator } from '@/components/bases/Indicator'
import { Carousel } from '@/components/composites/Carousel'
import SlideShow from '@/components/composites/SlideShow'
import Section from '@/components/layouts/Section'
import { IThumbnail } from '@/features/product/type'
import { ColorType } from '@/types'
import { NavigationType } from '@/types/enum/navigation'

import styles from './styles.module.scss'

type Props = {
    carouselImages: IThumbnail[]
}

const HomeTemplate = ({ carouselImages }: Props) => {
    const router = useRouter()

    return (
        <div className={styles['container']}>
            <div className={classNames(styles['logo-area'], styles['default'])}>
                <h1>
                    <Image
                        alt="とこりり"
                        height={200}
                        priority
                        src="/logo/tocoriri_logo.png"
                        style={{
                            objectFit: 'cover',
                        }}
                        width={400}
                    />
                </h1>
            </div>
            <Slide>
                <div className={classNames(styles['carousel-area'], styles['default'])}>
                    <Carousel items={carouselImages} />
                </div>
            </Slide>
            <Slide>
                <div className={classNames(styles['slide-show-area'], styles['sm'])}>
                    <SlideShow items={carouselImages} size="90vw" />
                </div>
            </Slide>
            <Section
                button
                buttonLabel="詳しくはこちら"
                color={ColorType.Primary}
                onButtonClick={() => {
                    router.push(NavigationType.About)
                }}
                title="About"
            >
                <p>仕事や出産、育児、家事...</p>
                <p>頑張る女性の味方になりたい、</p>
                <p>
                    そんな想いで
                    <br className={styles['sm']} />
                    マクラメ編みのアクセサリーを作っています。
                </p>
            </Section>
            <Section
                button
                buttonLabel="お問い合わせフォーム"
                color={ColorType.Primary}
                contrast
                onButtonClick={() => {
                    router.push(NavigationType.Contact)
                }}
                title="Contact"
            >
                <p>お問い合わせ・ご意見・ご相談はこちらから</p>
            </Section>
            <div className={styles['border']} />
            <span className={styles['indicator']}>
                <Indicator>
                    Scroll<span className={styles['arrow']}>→</span>
                </Indicator>
            </span>
        </div>
    )
}

export default HomeTemplate
