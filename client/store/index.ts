import { GetterTree, ActionTree, MutationTree } from 'vuex'

interface RootState {}

interface State {
    confirmVisible?: boolean
    confirmMessage: string
    confirmAction: () => void
    cancelAction: () => void
}

export interface ConfirmState {
    confirmVisible?: boolean
    confirmMessage: string
    confirmAction: () => void
    cancelAction: () => void
}
function newConfirmState(): ConfirmState {
    return {
        confirmVisible: false,
        confirmMessage: '',
        confirmAction: () => {},
        cancelAction: () => {},
    }
}

export const state = (): State => Object.assign(newConfirmState())

export const mutations: MutationTree<State> = {
    setConfirm(state, value: ConfirmState) {
        state.confirmVisible = value.confirmVisible
        state.confirmMessage = value.confirmMessage
        state.confirmAction = value.confirmAction
        state.cancelAction = value.cancelAction
    },
}

export const actions: ActionTree<State, RootState> = {
    confirm({ commit }, value: ConfirmState) {
        commit('setConfirm', {
            confirmVisible: true,
            confirmMessage: value.confirmMessage,
            confirmAction: value.confirmAction,
            cancelAction: value.cancelAction,
        })
    },
    closeConfirm({ commit }) {
        commit('setConfirm', newConfirmState())
    },
}

export const getters: GetterTree<State, RootState> = {
    // 確認ダイアログ表示
    confirmVisible(state: ConfirmState) {
        return state.confirmVisible
    },
    // 確認ダイアログメッセージ
    confirmMessage(state) {
        return state.confirmMessage
    },
    // 確認ダイアログキャンセル処理
    cancelAction(state) {
        return state.cancelAction
    },
    // 確認ダイアログ確定処理
    confirmAction(state) {
        return state.confirmAction
    },
}
