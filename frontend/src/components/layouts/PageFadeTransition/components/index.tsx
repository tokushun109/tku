import { LayoutRouterContext } from 'next/dist/shared/lib/app-router-context.shared-runtime'
import { useContext, useRef } from 'react'

interface Props {
    children: React.ReactNode
}

export const FrozenRouter = ({ children }: Props) => {
    const context = useContext(LayoutRouterContext ?? {})
    const frozen = useRef(context).current

    return <LayoutRouterContext.Provider value={frozen}>{children}</LayoutRouterContext.Provider>
}
