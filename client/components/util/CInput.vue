<template>
    <div class="c-text-input">
        <input
            v-if="password"
            v-model="syncModel"
            :disabled="disabled"
            :readonly="readonly"
            type="password"
            class="c-text-input-input"
            :placeholder="placeholder"
            @keyup.enter="keyupEnter"
        />
        <input
            v-else-if="number"
            v-model.number="syncModel"
            :disabled="disabled"
            :readonly="readonly"
            type="number"
            :step="step"
            class="c-text-input-input"
            :placeholder="placeholder"
            @keyup.enter="keyupEnter"
            @change="changeHandler"
        />
        <textarea
            v-else-if="multiline"
            v-model="syncModel"
            :disabled="disabled"
            :readonly="readonly"
            class="c-text-input-input"
            :style="{ height: height }"
            :placeholder="placeholder"
        ></textarea>
        <input
            v-else
            v-model="syncModel"
            :disabled="disabled"
            :readonly="readonly"
            type="text"
            class="c-text-input-input"
            :placeholder="placeholder"
            @input="inputHandler"
            @focus="focusHandler"
            @keyup.enter="keyupEnter"
        />
    </div>
</template>

<script lang="ts">
import { Component, Vue, Prop, PropSync } from 'nuxt-property-decorator'

@Component
export default class CInput extends Vue {
    // 無効状態か否か
    @Prop(Boolean) disabled?: boolean
    // 高さ
    @Prop({ type: String, default: '100px' }) height?: string
    // パスワード入力
    @Prop(Boolean) password?: boolean
    // テキストエリア
    @Prop(Boolean) multiline?: boolean
    // 読み取り専用
    @Prop(Boolean) readonly?: boolean
    // 数値入力
    @Prop(Boolean) number?: boolean
    // textのみ表示の場合は、itemsではなくtextで
    @Prop(String) text?: String
    // placeholder
    @Prop() placeholder?: string | number
    // number型の場合stepを指定
    @Prop({ type: String, default: '1' }) step?: String
    // どんなデータで欲しいのかよくわからないからとりあえず文字列で返す
    @PropSync('model', { default: undefined }) syncModel!: string | number | boolean | null | undefined

    inputHandler() {
        this.$emit('c-input', this.text)
    }

    focusHandler() {
        this.$emit('c-focus')
    }

    keyupEnter() {
        this.$emit('c-enter')
    }

    changeHandler() {
        this.$emit('c-change')
    }
}
</script>

<style lang="stylus">
.c-text-input
    input
    textarea
        display block
        margin 2px
        padding 8px
        width calc(100% - 22px)
        border 1px solid $light-dark-color
        border-radius 3px
        background-color #fff
        // width: 100%
        font-size $form-item-font-size
        resize none
        &:disabled
            background-color #dadada
    ::placeholder
        color $light-dark-color
    input[type='number']
        text-align right
</style>
