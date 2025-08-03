'use client'

import { ContactList } from '@/features/contact/components/ContactList'
import { IContactListItem } from '@/features/contact/type'

import styles from './styles.module.scss'

interface Props {
    contacts: IContactListItem[]
}

export const AdminContactTemplate = ({ contacts }: Props) => {
    return (
        <div className={styles['contact-container']}>
            <div className={styles['page-header']}>
                <h1 className={styles['page-title']}>お問い合わせ一覧</h1>
                <div className={styles['contact-count']}>{contacts.length}件のお問い合わせ</div>
            </div>
            <div className={styles['divider']} />
            <div className={styles['contact-content']}>
                <ContactList contacts={contacts} />
            </div>
        </div>
    )
}
