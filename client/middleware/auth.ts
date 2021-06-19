import { Context } from '@nuxt/types'

export default async function ({ app, redirect }: Context) {
    // sessionのcookieを取得
    const credential = app.$cookies.get('__sess__')
    if (credential) {
        // ログインしているユーザの情報を取得
        const user = await app.$axios.$get(`/user/login/${credential}`)
        if (!user) {
            return redirect('/admin/user/login')
        }
    } else {
        return redirect('/admin/user/login')
    }
}
