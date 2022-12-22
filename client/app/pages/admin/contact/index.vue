<template>
    <v-sheet>
        <c-contact-list :contact-list="contactList" />
    </v-sheet>
</template>

<script lang="ts">
import { Context } from '@nuxt/types'
import { Component, Vue } from 'nuxt-property-decorator'
import { IContact } from '~/types'
@Component({
    head: {
        title: 'お問い合わせ',
    },
})
export default class PageAdminContactIndex extends Vue {
    // お問合わせリスト
    contactList: Array<IContact> = []

    async asyncData({ app }: Context) {
        try {
            const contactList = await app.$axios.$get(`/contact`)
            return { contactList }
        } catch (e) {
            return { contactList: [] }
        }
    }
}
</script>

<style lang="stylus"></style>
