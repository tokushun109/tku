<template>
    <div class="c-top-main">
        <div class="c-top-main-logo">
            <img v-if="producer" class="producer-logo top-logo" :src="producer.logo" alt="producerLogo" />
            <img v-else class="home-logo top-logo" src="/icon/home.png" alt="homeLogo" />
        </div>
        <div class="c-top-main-message">
            <h1>ここにメッセージが入ります</h1>
        </div>
    </div>
</template>

<script lang="ts">
import { Component, Vue, Prop } from 'nuxt-property-decorator'
import { IProducerProfile } from '~/types'

@Component({})
export default class CTopMain extends Vue {
    // 製作者情報
    @Prop(Object) producer!: IProducerProfile
}
</script>

<style lang="stylus">
.c-top-main
    height 100vh
    text-align center
    .c-top-main-logo
        .top-logo
            width 250px
            height 250px
            // TODO mainのカラーを決定する
            border double 1px #000000
            object-fit cover
            &.producer-logo
                border-radius 50%
</style>
