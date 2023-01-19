import { GetterTree, ActionTree, MutationTree } from 'vuex/types'
import { IUser, ILoginForm, ISession, errorCustomize } from '~/types'

interface RootState {}
interface State {
    isGuest: boolean
    user: IUser | null
}

export const state = (): State => ({
    isGuest: true,
    user: null,
})

function newState(): State {
    return {
        isGuest: true,
        user: null,
    }
}
export const mutations: MutationTree<State> = {
    setUser(state, user: IUser | null) {
        if (user !== null) {
            state.isGuest = false
            state.user = user
        } else {
            newState()
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
            const user = await this.$axios.$get(`/user/login`, { withCredentials: true })
            // store保存
            context.commit('setUser', user)
            return Promise.resolve(user)
        } catch (e: any) {
            return Promise.reject(errorCustomize(e.response, 'メールアドレスかパスワードが間違っています'))
        }
    },
    setUser(context, user) {
        // store保存
        context.commit('setUser', user)
    },
    async logoutUser(context) {
        try {
            await this.$axios.$post(`/user/logout/`, null, { withCredentials: true })
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
    isGuest(state: State) {
        return state.isGuest
    },
    user(state: State): IUser | null {
        return state.user
    },
}
