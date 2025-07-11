export const slideAnimation = {
    initial: { translateY: '-100vh' },
    animate: { translateY: 0 },
    exit: { translateY: '-100vh' },
    transition: { duration: 0.3, ease: 'easeInOut' },
} as const
