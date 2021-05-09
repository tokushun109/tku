<template>
    <div class="c-open-menu">
        <ul class="c-open-menu-wrapper">
            <c-column column="4">
                <li @click="closeMenu()">
                    <nuxt-link to="/">ホーム</nuxt-link>
                </li>
                <li @click="closeMenu()">
                    <nuxt-link to="/product">商品</nuxt-link>
                </li>
                <li @click="closeMenu()">
                    <nuxt-link to="/creator">製作者</nuxt-link>
                </li>
                <li @click="closeMenu()">
                    <nuxt-link to="/order">仕事の依頼</nuxt-link>
                </li>
            </c-column>
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
    top 0
    left 0
    z-index 998
    width 100vw
    height 100vh
    background-color white
    .c-open-menu-wrapper
        position relative
        top 50%
        margin 0 auto
        width 50%
        text-align center
</style>
