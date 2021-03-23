<template>
    <header>
        <div class="header-wrapper">
            <div v-if="menuOpenFrag" class="open-menu-wrapper">
                <div class="close-icon">
                    <img src="/icon/close.png" alt="close" @click="toggle" />
                </div>
                <c-open-menu @close-menu="toggle" />
            </div>
            <div v-else class="open-icon">
                <img src="/icon/menu.png" alt="open" @click="toggle" />
            </div>
        </div>
    </header>
</template>

<script lang="ts">
import { Component, Vue } from 'nuxt-property-decorator'
@Component({})
export default class Header extends Vue {
    menuOpenFrag: boolean = false
    toggle() {
        this.menuOpenFrag = !this.menuOpenFrag
        return this.menuOpenFrag
    }
}
</script>

<style lang="stylus">
header
    position fixed
    z-index 999
    padding 80px
    width 100vw
    .header-wrapper
        .open-menu-wrapper
            .close-icon
                position fixed
                right 60px
                z-index 999
        .open-icon
            position fixed
            right 60px
            z-index 999
</style>
