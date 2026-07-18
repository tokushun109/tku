export const handler = async () => {
  try {
    const res = await fetch('https://api.tocoriri.com/api/health_check', {
      headers: {
        'Content-Type': 'application/json',
      },
      method: 'GET',
      signal: AbortSignal.timeout(15000),
    })
    const result = (await res.json()) as { success?: unknown }
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
  }
}

const lineNotification = async () => {
  const token = process.env.LINE_HEALTH_CHECK_TOKEN
  const userId = process.env.LINE_HEALTH_CHECK_USER_ID

  if (!token || !userId) {
    console.error('LINE_HEALTH_CHECK_TOKEN または LINE_HEALTH_CHECK_USER_ID が設定されていません。')
    return
  }

  const headers = {
    'Content-Type': 'application/json',
    Authorization: `Bearer ${token}`,
  }

  const payload = {
    to: userId,
    messages: [
      {
        type: 'text',
        text: 'とこりりのヘルスチェックに失敗しました。サーバーが落ちている可能性があります。',
      },
    ],
  }

  try {
    const response = await fetch('https://api.line.me/v2/bot/message/push', {
      method: 'POST',
      body: JSON.stringify(payload),
      headers,
    })
    if (!response.ok) {
      throw new Error(`LINE API returned ${response.status}`)
    }
    console.log('送信しました')
  } catch (e) {
    console.error('LINE通知の送信に失敗しました。', e)
  }
}
