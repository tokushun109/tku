'use client'

import { useEffect } from 'react'

interface Props {
    error: Error & { digest?: string }
    reset: () => void
}

const Error = ({ error, reset }: Props) => {
    useEffect(() => {
        console.error(error)
    }, [error])

    return (
        <div
            style={{
                padding: '48px 16px',
                textAlign: 'center',
                minHeight: '50vh',
                display: 'flex',
                flexDirection: 'column',
                justifyContent: 'center',
                alignItems: 'center',
            }}
        >
            <h2 style={{ marginBottom: '16px' }}>商品が見つかりませんでした</h2>
            <p style={{ marginBottom: '24px', color: '#666' }}>お探しの商品は存在しないか、削除された可能性があります。</p>
            <button
                onClick={reset}
                style={{
                    padding: '12px 24px',
                    backgroundColor: '#007bff',
                    color: 'white',
                    border: 'none',
                    borderRadius: '4px',
                    cursor: 'pointer',
                }}
            >
                もう一度試す
            </button>
        </div>
    )
}

export default Error
