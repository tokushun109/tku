import { Context } from '@nuxt/types'

export default async function ({ app, route, store, redirect }: Context) {
    // /creatorを/aboutにリダイレクトさせる
    if (route.path === '/creator') {
        return redirect('/about')
    }
    // SSRの時は認証処理を行わない
    if (!process.server) {
        return
    }

    // ログイン画面の時は認証処理を行わない
    if (!route.path.includes('/admin') || route.path === '/admin/user/login') {
        return
    }
    // sessionのcookieを取得
    const credential = app.$cookies.get('__sess__')
    if (credential) {
        // ログインしているユーザの情報を取得
        const user = await app.$axios.$get(`/user/login/${credential}`)
        store.dispatch('user/setUser', user)
        if (!user) {
            return redirect('/admin/user/login')
        }
    } else {
        return redirect('/admin/user/login')
    }
}
