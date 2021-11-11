<template>
    <v-sheet color="warning" rounded class="title pa-4 text-center">
        <slot />
    </v-sheet>
</template>

<script lang="ts">
import { Component, Vue } from 'nuxt-property-decorator'
@Component
export default class CMessage extends Vue {}
</script>

<style lang="stylus" scoped></style>
