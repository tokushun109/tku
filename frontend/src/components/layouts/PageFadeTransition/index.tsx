'use client'

import { AnimatePresence, motion } from 'framer-motion'
import { usePathname } from 'next/navigation'

import { NavigationType } from '@/types'

import { FrozenRouter } from './components'

interface Props {
    children: React.ReactNode
}

export const PageFadeTransition = ({ children }: Props) => {
    const pathname = usePathname() as NavigationType

    return (
        <div>
            <AnimatePresence mode="wait">
                <motion.div
                    animate={{ opacity: 1 }}
                    exit={{ opacity: 0 }}
                    initial={{ opacity: 0 }}
                    key={pathname}
                    transition={{ duration: 0.3, ease: 'easeInOut' }}
                >
                    <FrozenRouter>{children}</FrozenRouter>
                </motion.div>
            </AnimatePresence>
        </div>
    )
}
