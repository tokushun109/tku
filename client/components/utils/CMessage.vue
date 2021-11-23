<template>
    <v-sheet :color="color" rounded class="title pa-4 text-center">
        <slot />
    </v-sheet>
</template>

<script lang="ts">
import { Component, Prop, Vue } from 'nuxt-property-decorator'
@Component
export default class CMessage extends Vue {
    @Prop({ type: String, default: 'warning' }) color!: string
}
</script>

<style lang="stylus" scoped></style>
