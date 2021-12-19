<template>
    <header>
        <div class="header-wrapper">
            <div v-if="!isAdmin" class="user-menu">
                <div v-if="menuOpenFlag" class="open-menu-wrapper">
                    <div class="close-icon">
                        <img src="/icon/close.png" alt="close" @click="toggle" />
                    </div>
                    <c-click-menu @close-menu="toggle" />
                </div>
                <div v-else class="open-icon">
                    <img src="/icon/menu.png" alt="open" @click="toggle" />
                </div>
            </div>
            <div v-else-if="!$store.getters['user/isGuest']" class="admin-menu">
                <c-button class="logout-button" label="ログアウト" @c-click="visibleSync = true" />
            </div>
        </div>
        <c-dialog :visible.sync="visibleSync" height="150px" width="400px" :is-header="false" @confirm="logoutHandler" @close="visibleSync = false">
            <h4 class="logout-title">ログアウトします。よろしいですか？</h4>
        </c-dialog>
    </header>
</template>

<script lang="ts">
import { Component, Vue } from 'nuxt-property-decorator'
@Component({})
export default class Header extends Vue {
    visibleSync: boolean = false
    menuOpenFlag: boolean = false
    toggle() {
        this.menuOpenFlag = !this.menuOpenFlag
        return this.menuOpenFlag
    }

    // urlにadminが含まれているかを確認
    get isAdmin() {
        return this.$route.path.includes('admin')
    }

    async logoutHandler() {
        await this.$store.dispatch('user/logoutUser')
        this.$router.replace('/admin/user/login')
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
