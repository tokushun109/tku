<template>
    <header>
        <div class="header-wrapper">
            <div v-if="menuOpenFrag" class="open-menu">
                <ul>
                    <li><nuxt-link to="/product">商品</nuxt-link></li>
                    <li><nuxt-link to="/producer_profile">製作者</nuxt-link></li>
                    <li><nuxt-link to="/order">仕事の依頼</nuxt-link></li>
                </ul>
                <img class="close-icon" src="/icon/close.png" alt="close" @click="toggle" />
            </div>
            <img v-else class="menu-icon" src="/icon/menu.png" alt="menu" @click="toggle" />
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

<style lang="stylus"></style>
