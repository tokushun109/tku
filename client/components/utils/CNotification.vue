<template>
    <v-dialog v-model="notificationVisible" hide-overlay max-width="400" height="100">
        <v-card>
            <v-card-title class="justify-center"><slot /></v-card-title>
        </v-card>
    </v-dialog>
</template>

<script lang="ts">
import { Component, Vue, Watch, PropSync } from 'nuxt-property-decorator'

@Component
export default class CNotification extends Vue {
    // 表示フラグ
    @PropSync('visible') notificationVisible!: boolean

    @Watch('notificationVisible')
    autoClosing() {
        setTimeout(() => {
            this.notificationVisible = false
        }, 2000)
    }
}
</script>

<style lang="stylus" scoped>
.v-dialog__content
    justify-content flex-end
    align-items flex-start

.v-card__title
    padding 16px 24px !important
</style>
