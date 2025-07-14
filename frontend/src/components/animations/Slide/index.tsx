'use client'

import { motion, useInView } from 'framer-motion'
import { useRef } from 'react'

export const SlideDirectionEnum = {
    Down: 'down',
    Up: 'up',
} as const
export type SlideDirectionType = (typeof SlideDirectionEnum)[keyof typeof SlideDirectionEnum]

type Props = {
    children: React.ReactNode
}

export const Slide = ({ children }: Props) => {
    const ref = useRef<HTMLDivElement>(null)
    const isInView = useInView(ref, { once: true })

    return (
        <motion.div
            animate={{ opacity: isInView ? 1 : 0, y: isInView ? 0 : 30 }}
            initial={{ opacity: 0, y: 20 }}
            ref={ref}
            transition={{ duration: 0.6, ease: 'easeOut' }}
        >
            {children}
        </motion.div>
    )
}
