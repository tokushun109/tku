'use client'

import { AnimatePresence, motion } from 'framer-motion'
import { usePathname } from 'next/navigation'

import { NavigationType } from '@/types'

import { fadeAnimation } from './animations'
import { FrozenRouter } from './components'

interface Props {
    children: React.ReactNode
}

export const PageFadeTransition = ({ children }: Props) => {
    const pathname = usePathname() as NavigationType

    return (
        <div>
            <AnimatePresence mode="wait">
                <motion.div key={pathname} {...fadeAnimation}>
                    <FrozenRouter>{children}</FrozenRouter>
                </motion.div>
            </AnimatePresence>
        </div>
    )
}
