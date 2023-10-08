import fetch from 'node-fetch'

export const handler = async () => {
    const timeout = setTimeout(async () => {
        await lineNotification()
        throw new Error('タイムアウトしました')
    }, 15000);
    try {

        const res = await fetch('https://api.tocoriri.com/api/health_check', {
            headers: {
                'Content-Type': 'application/json',
            },
            method: 'GET',
        })
        const result = await res.json() as { success?: boolean }
        if (result.success === undefined) {
            throw new Error('api取得失敗')
        }
        console.log('api取得OK')
        return {
            statusCode: 200,
            body: {
                message: 'api取得OK',
                result
            },
        }
    } catch (e) {
        await lineNotification()
        return {
            statusCode: 500,
            body: {
                message: 'apiの取得に失敗しました',
                result: null
            },
        }
    } finally { clearTimeout(timeout); }
}

const lineNotification = async () => {
    const qs = require('querystring')
    await fetch('https://notify-api.line.me/api/notify', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/x-www-form-urlencoded',
            Authorization: `Bearer ${process.env.LINE_HEALTH_CHECK_TOKEN}`,
        },
        body: qs.stringify({
            message: 'とこりりのヘルスチェックに失敗しました。サーバーが落ちている可能性があります。',
        }),
    })
}

