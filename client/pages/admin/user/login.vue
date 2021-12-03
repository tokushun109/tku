<template>
    <v-main class="grey lighten-4">
        <v-container>
            <v-sheet class="pa-4 lighten-4 text-center">
                <c-error :errors.sync="errors" />
                <h3 class="title mb-4 green--text text--darken-3">ログイン</h3>
                <v-form ref="form" v-model="valid" lazy-validation>
                    <v-text-field v-model="form.email" :rules="rules" label="email(必須)" outlined />
                    <v-text-field v-model="form.password" :rules="rules" label="パスワード(必須)" outlined />
                    <v-btn color="primary" :disabled="!valid" @click="onSubmit">確定</v-btn>
                </v-form>
            </v-sheet>
        </v-container>
    </v-main>
</template>

<script lang="ts">
import { Component, Vue } from 'nuxt-property-decorator'
import { required } from '~/methods'
import { IError, ILoginForm } from '~/types'

@Component({
    head: {
        titleTemplate: 'ログイン | admin',
    },
})
export default class PageAdminUserLogin extends Vue {
    form: ILoginForm = {
        email: '',
        password: '',
    }

    valid: boolean = true

    rules = [required]

    errors: Array<IError> = []

    async onSubmit() {
        try {
            this.errors = []
            await this.$store.dispatch('user/loginUser', this.form)
            this.$router.replace('/admin/product')
        } catch (e) {
            this.errors.push(e)
        }
    }
}
</script>

<style lang="stylus"></style>
