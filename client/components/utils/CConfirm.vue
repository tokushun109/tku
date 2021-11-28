<template>
    <c-dialog :visible.sync="$store.getters.confirmVisible" title="確認メッセージ" @confirm="confirmHandler" @close="cancelHandler">
        <template #content>
            <v-container> {{ $store.getters.confirmMessage }} </v-container>
        </template>
    </c-dialog>
</template>

<script lang="ts">
import { Component, Vue } from 'nuxt-property-decorator'
@Component
export default class CConfirm extends Vue {
    // 確定イベント
    confirmHandler() {
        if (this.$store.getters.confirmAction) {
            this.$store.getters.confirmAction()
        }
        this.$store.dispatch('closeConfirm')
    }

    // キャンセルイベント
    cancelHandler() {
        if (this.$store.getters.cancelAction) {
            this.$store.getters.cancelAction()
        }
        this.$store.dispatch('closeConfirm')
    }
}
</script>

<style lang="stylus" scoped></style>
