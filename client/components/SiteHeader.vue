<template>
    <header>
        <div class="header-wrapper">
            <nuxt-link class="site-title" to="/">
                <h1>tku</h1>
            </nuxt-link>
            <div id="nav-drawer">
                <ul>
                    <li><nuxt-link to="/product">商品</nuxt-link></li>
                    <li><nuxt-link to="/producer_profile">製作者</nuxt-link></li>
                    <li><nuxt-link to="/order">仕事の依頼</nuxt-link></li>
                </ul>
            </div>
        </div>
    </header>
</template>

<script lang="ts">
import { Component, Vue } from 'nuxt-property-decorator'
@Component({})
export default class Header extends Vue {}
</script>

<style lang="stylus"></style>
