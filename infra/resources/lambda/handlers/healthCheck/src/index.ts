import fetch from 'node-fetch'

export const handler = async () => {
    const timeout = setTimeout(async () => {
        await lineNotification()
        throw new Error('タイムアウトしました')
    }, 15000)
    try {
        const res = await fetch('https://api.tocoriri.com/api/health_check', {
            headers: {
                'Content-Type': 'application/json',
            },
            method: 'GET',
        })
        const result = (await res.json()) as { success?: boolean }
        if (result.success === undefined) {
            throw new Error('api取得失敗')
        }
        console.log('api取得OK')
        return {
            statusCode: 200,
            body: {
                message: 'api取得OK',
                result,
            },
        }
    } catch (e) {
        await lineNotification()
        return {
            statusCode: 500,
            body: {
                message: 'apiの取得に失敗しました',
                result: null,
            },
        }
    } finally {
        clearTimeout(timeout)
    }
}

const lineNotification = async () => {
    const headers = {
        'Content-Type': 'application/json',
        Authorization: `Bearer ${process.env.LINE_HEALTH_CHECK_TOKEN}`,
    }

    const payload = {
        to: process.env.LINE_HEALTH_CHECK_USER_ID,
        messages: [
            {
                type: 'text',
                text: 'とこりりのヘルスチェックに失敗しました。サーバーが落ちている可能性があります。',
            },
        ],
    }

    const options = {
        method: 'POST',
        body: JSON.stringify(payload),
        headers: headers,
    }

    await fetch('https://api.line.me/v2/bot/message/push', options)
    console.log('送信しました')
}
