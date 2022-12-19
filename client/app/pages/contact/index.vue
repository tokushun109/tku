<template>
    <v-container class="page-contact">
        <v-container class="page-title-container">
            <h2 class="page-title text-sm-h3 text-h4">CONTACT</h2>
        </v-container>
        <div class="contact-introduction">
            <p>お問い合わせ・<br class="sm" />ご意見・ご相談はこちらから</p>
        </div>
        <v-sheet>
            <v-container>
                <c-contact-form />
            </v-container>
        </v-sheet>
    </v-container>
</template>

<script lang="ts">
import { Component, Vue } from 'nuxt-property-decorator'
@Component({})
export default class PageContactIndex extends Vue {
    head() {
        return {
            meta: [
                {
                    hid: 'robots',
                    name: 'robots',
                    content: 'noindex',
                },
            ],
        }
    }
}
</script>

<style lang="stylus" scoped>
.page-title-container
    text-align center
    +sm()
        display none
    .page-title
        margin-bottom 20px
        color $site-title-text-color
        text-align center

.contact-introduction
    text-align center
    .sm
        display none
        +sm()
            display block
</style>
