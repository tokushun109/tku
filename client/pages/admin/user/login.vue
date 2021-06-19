<template>
    <c-page class="page-login" title="ログイン">
        <c-form bordered slim>
            <c-input-label label="メールアドレス" required>
                <c-input :model.sync="form.email" />
            </c-input-label>
            <c-input-label label="パスワード" required>
                <c-input :model.sync="form.password" password />
            </c-input-label>
            <div class="form-actions">
                <c-button label="ログイン" primary @c-click="onSubmit" />
            </div>
        </c-form>
    </c-page>
</template>

<script lang="ts">
import { Component, Vue } from 'nuxt-property-decorator'

interface IloginForm {
    email: string
    password: string
}

interface ISession {
    uuid: string
}

@Component({
    head: {
        titleTemplate: 'ログイン | admin',
    },
})
export default class PageAdminUserLogin extends Vue {
    form: IloginForm = {
        email: '',
        password: '',
    }

    async onSubmit() {
        try {
            const session: ISession = await this.$axios.$post(`/user/login`, this.form).catch(() => {})
            await this.$cookies.set('__sess__', session.uuid, {
                path: '/',
            })
            this.$router.replace('/admin')
        } catch {}
    }
}
</script>

<style lang="stylus"></style>
