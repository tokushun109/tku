<template>
    <div class="c-open-menu">
        <ul class="c-open-menu-wrapper">
            <li @click="closeMenu()"><nuxt-link to="/">ホーム</nuxt-link></li>
            <li @click="closeMenu()"><nuxt-link to="/product">商品</nuxt-link></li>
            <li @click="closeMenu()"><nuxt-link to="/producer_profile">製作者</nuxt-link></li>
            <li @click="closeMenu()"><nuxt-link to="/order">仕事の依頼</nuxt-link></li>
        </ul>
    </div>
</template>

<script lang="ts">
import { Component, Vue, Emit } from 'nuxt-property-decorator'

@Component({})
export default class COpenMenu extends Vue {
    @Emit('close-menu')
    closeMenu() {}
}
</script>

<style lang="stylus">
.c-open-menu
    position fixed
    z-index 998
    background-color white
    top 0
    left 0
    width 100vw
    height 100vh
    .c-open-menu-wrapper
        position absolute
        top 50%
        left 50%
        transform translate(-50%, -50%)
        z-index 1000
</style>
