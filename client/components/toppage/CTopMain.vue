<template>
    <div class="c-top-main">
        <div class="c-top-main-contents">
            <img v-if="creator" class="creator-logo top-logo" :src="creator.logo" alt="creatorLogo" />
            <img v-else class="home-logo top-logo" src="/icon/home.png" alt="homeLogo" />
        </div>
    </div>
</template>

<script lang="ts">
import { Component, Vue, Prop } from 'nuxt-property-decorator'
import { ICreator } from '~/types'

@Component({})
export default class CTopMain extends Vue {
    // 製作者情報
    @Prop(Object) creator!: ICreator
}
</script>

<style lang="stylus">
.c-top-main
    position relative
    height 'calc(100vh - %s)' % $page-padding[0]
    text-align center
    .c-top-main-contents
        position absolute
        top 0
        right 0
        bottom 0
        left 0
        margin auto
        height 275px
        .top-logo
            width 250px
            height 250px
            // TODO mainのカラーを決定する
            border double 1px #000000
            object-fit cover
            &.producer-logo
                border-radius 50%
</style>
