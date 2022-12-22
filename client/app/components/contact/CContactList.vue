<template>
    <div class="c-contact-list">
        <v-data-table :headers="headers" :items="contactList" :items-per-page="5" class="elevation-1"></v-data-table>
    </div>
</template>

<script lang="ts">
import { Component, Prop, Vue } from 'nuxt-property-decorator'
import { ITableHeader, IContact } from '~/types'

@Component({})
export default class CContactList extends Vue {
    @Prop({ type: Array, default: () => [] }) contactList!: Array<IContact>

    headers: Array<ITableHeader> = [
        { text: 'ID', align: 'start', value: 'id' },
        { text: '氏名', value: 'name' },
        { text: '会社', value: 'company' },
        { text: '電話番号', value: 'phoneNumber' },
        { text: 'メールアドレス', value: 'email' },
        { text: '内容', value: 'content' },
        { text: 'お問合わせ日時', value: 'createdAt' },
    ]

    // ダイアログの表示
    dialogVisible: boolean = false

    // 通知の表示
    notificationVisible: boolean = false
}
</script>

<style lang="stylus" scoped></style>
