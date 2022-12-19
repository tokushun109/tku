<template>
    <v-container class="c-contact-form">
        <v-sheet class="content-form-wrapper">
            <template v-if="!isSent">
                <v-form ref="form" v-model="valid" lazy-validation>
                    <v-text-field v-model="contact.name" :rules="nameRules" label="お名前(必須)" outlined counter="20" />
                    <v-text-field v-model="contact.company" :rules="companyRules" label="会社名" outlined counter="20" />
                    <v-text-field
                        v-model="contact.phoneNumber"
                        :rules="phoneRules"
                        label="電話番号(-を入れずに入力)"
                        validate-on-blur
                        outlined
                        counter="10"
                    />
                    <v-text-field
                        v-model="contact.mailAddress"
                        :rules="mailAddressRules"
                        label="メールアドレス(必須)"
                        outlined
                        validate-on-blur
                        counter="50"
                    />
                    <v-textarea v-model="contact.content" :rules="contentRules" label="お問い合わせ内容(必須)" outlined />
                    <div class="text-center">
                        <v-btn color="primary" :disabled="!valid" @click="confirmHandler">送信する</v-btn>
                    </div>
                </v-form>
            </template>
            <div v-else class="content-message-wrapper">送信しました</div>
        </v-sheet>
    </v-container>
</template>

<script lang="ts">
import { Component, Vue } from 'nuxt-property-decorator'
import { IError, IContact } from '~/types'
import { min20, min50, newContact, required, validMailAddress, validPhoneNumber } from '~/methods'
@Component({})
export default class CContactForm extends Vue {
    contact: IContact = newContact()

    isSent: boolean = false

    valid: boolean = false

    errors: Array<IError> = []

    nameRules = [required, min20]

    companyRules = [min20]

    phoneRules = [validPhoneNumber]

    mailAddressRules = [required, min50, validMailAddress]

    contentRules = [required]

    async confirmHandler() {
        const refs: any = this.$refs.form
        await refs.validate()
        try {
            if (!this.valid) {
                return
            }
            await this.$axios.$post(`/contact`, this.contact)
            this.isSent = true
        } catch (e) {
            this.errors.push(e)
        }
    }
}
</script>

<style lang="stylus" scoped></style>
