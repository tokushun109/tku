<template>
    <v-breadcrumbs :items="items" divider=">">
        <template #item="{ item }">
            <v-breadcrumbs-item nuxt :href="item.href" :disabled="item.disabled">
                {{ item.text }}
            </v-breadcrumbs-item>
        </template>
    </v-breadcrumbs>
</template>

<script lang="ts">
import { Component, Vue, Prop } from 'nuxt-property-decorator'
import { IBreadCrumb } from '~/types'

@Component({})
export default class CBreadcrumbs extends Vue {
    @Prop({ type: Array, default: false }) items!: Array<IBreadCrumb>
}
</script>

<style lang="stylus" scoped></style>
