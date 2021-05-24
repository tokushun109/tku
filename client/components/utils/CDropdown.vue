<template>
    <span class="c-dropdown" :class="disabled">
        <select
            v-model="syncModel"
            :name="selectName"
            class="c-dropdown-input"
            :class="{ multiple: multiple }"
            :multiple="multiple"
            :disabled="disabled"
            @change="checkedHandler"
        >
            <option v-if="!multiple" value=""></option>
            <option v-for="(item, index) in items" :key="index" :value="item" :selected="index === 0">
                {{ item[property] }}
            </option>
            <slot />
        </select>
    </span>
</template>

<script lang="ts">
import { Component, Vue, Prop, PropSync, Emit } from 'nuxt-property-decorator'
@Component
export default class CDropdown extends Vue {
    // selectタグのnameに使用
    @Prop(String) name!: string
    // 無効状態か否か
    @Prop(Boolean) disabled?: boolean
    // 複数選択
    @Prop(Boolean) multiple?: boolean
    // モデル
    @PropSync('model', { default: undefined }) syncModel!: any
    // optionに渡されるリスト
    @Prop({ type: Array, default: null }) items!: Array<any>
    // itemのどのプロパティを使用するか
    @Prop({ type: String, default: 'name' }) property!: string

    // selectのname属性を取得
    get selectName() {
        return this.name + 'dropdown'
    }

    // 変更通知
    @Emit('c-change')
    private checkedHandler() {
        return this.syncModel
    }
}
</script>

<style lang="stylus">
.c-dropdown
    display block
    select
        padding 8px 32px 8px 8px
        width 100%
        outline none
        border 1px solid #ccc
        border-radius 4px
        background #fff url('/img/arrow/c-arrow-down.svg') no-repeat right 10px center
        background-size 12px auto
        color inherit
        vertical-align middle
        text-indent 0.01px
        text-overflow ''
        font-size inherit
        appearance none
        &.multiple
            padding-right 8px
            background-image none
        &:disabled
            background-color $super-light-color
            cursor not-allowed
    select::-ms-expand
        display none
</style>
