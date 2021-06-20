<template>
    <header>
        <div class="header-wrapper">
            <div v-if="!isAdmin" class="user-menu">
                <div v-if="menuOpenFlag" class="open-menu-wrapper">
                    <div class="close-icon">
                        <img src="/icon/close.png" alt="close" @click="toggle" />
                    </div>
                    <c-open-menu @close-menu="toggle" />
                </div>
                <div v-else class="open-icon">
                    <img src="/icon/menu.png" alt="open" @click="toggle" />
                </div>
            </div>
            <!-- TODO ログインしている時だけ表示する -->
            <div v-else class="admin-menu">
                <a href="/admin/user/login" @click="logoutHandler">ログアウト</a>
            </div>
        </div>
    </header>
</template>

<script lang="ts">
import { Component, Vue } from 'nuxt-property-decorator'
@Component({})
export default class Header extends Vue {
    menuOpenFlag: boolean = false
    toggle() {
        this.menuOpenFlag = !this.menuOpenFlag
        return this.menuOpenFlag
    }

    // urlにadminが含まれているかを確認
    get isAdmin() {
        return this.$route.path.includes('admin')
    }

    // sessin用のcookieを削除して、ログアウトする
    logoutHandler() {
        this.$cookies.remove('__sess__')
    }
}
</script>

<style lang="stylus">
header
    position fixed
    z-index 999
    width 100vw
    .header-wrapper
        .user-menu
            .open-menu-wrapper
                .close-icon
                    position fixed
                    top 60px
                    right 60px
                    z-index 999
            .open-icon
                position fixed
                top 60px
                right 60px
                z-index 999
        .admin-menu
            position fixed
            top 60px
            right 60px
            z-index 999
</style>
