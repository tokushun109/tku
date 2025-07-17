'use client'

import { useRouter } from 'next/navigation'
import { useEffect } from 'react'

import { healthCheck } from '@/apis/healthCheck'
import ErrorPage from '@/app/error'

const MaintenancePage: React.FC = () => {
    const router = useRouter()

    useEffect(() => {
        const checkHealthStatus = async () => {
            try {
                await healthCheck()
                // ヘルスチェックが成功した場合（メンテナンス中ではない）、トップページにリダイレクト
                router.push('/')
            } catch {
                // ヘルスチェックが失敗した場合はメンテナンスページを表示
            }
        }

        checkHealthStatus()
    }, [router])

    return (
        <ErrorPage
            errorMessage={
                <>
                    <p>ただいまメンテナンス中です</p>
                    <p>しばらく経ってからお試しください</p>
                </>
            }
            showHomeButton={false}
        />
    )
}

export default MaintenancePage
