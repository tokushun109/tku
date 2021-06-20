import { GetterTree, ActionTree, MutationTree } from 'vuex'
import { IUser, ILoginForm, ISession } from '~/types'

export interface State {
    isGuest: boolean
    user: IUser | null
}

export const state = (): State => ({
    isGuest: true,
    user: null,
})

export interface RootState {}

export const mutations: MutationTree<State> = {
    setUser(state, user: IUser | null) {
        if (user !== null) {
            state.isGuest = false
            state.user = user
        } else {
            state.isGuest = true
            state.user = null
        }
    },
}

export const actions: ActionTree<State, RootState> = {
    async loginUser(context, loginForm: ILoginForm) {
        try {
            const session: ISession = await this.$axios.$post(`/user/login`, loginForm)
            // cookie保存
            this.$cookies.set('__sess__', session.uuid, {
                path: '/',
            })
            const user = await this.$axios.$get(`/user/login/${session.uuid}`)
            // store保存
            context.commit('setUser', user)
        } catch {}
    },
    setUser(context, user) {
        // store保存
        context.commit('setUser', user)
    },
    async logoutUser(context) {
        try {
            const credential = this.$cookies.get('__sess__')
            await this.$axios.$post(`/user/logout/${credential}`)
            this.$cookies.remove('__sess__', {
                path: '/',
            })
            location.reload()
            // store保存
            context.commit('setUser', null)
        } catch {}
    },
}

export const getters: GetterTree<State, RootState> = {
    isGuest(state) {
        return state.isGuest
    },
    user(state): IUser | null {
        return state.user
    },
}
