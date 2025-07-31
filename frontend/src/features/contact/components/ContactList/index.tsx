'use client'

import { Email, Person } from '@mui/icons-material'
import { useState } from 'react'
import { Virtuoso } from 'react-virtuoso'

import { Dialog } from '@/components/bases/Dialog'
import { ListItem } from '@/components/bases/ListItem'
import { IContactListItem } from '@/features/contact/type'

import styles from './styles.module.scss'

interface Props {
    contacts: IContactListItem[]
}

export const ContactList = ({ contacts }: Props) => {
    const [selectedContact, setSelectedContact] = useState<IContactListItem | null>(null)
    const [isDetailDialogOpen, setIsDetailDialogOpen] = useState<boolean>(false)

    const handleContactClick = (contact: IContactListItem) => {
        setSelectedContact(contact)
        setIsDetailDialogOpen(true)
    }

    const handleCloseDetailDialog = () => {
        setIsDetailDialogOpen(false)
        setSelectedContact(null)
    }

    const formatDate = (dateString: string) => {
        return new Date(dateString).toLocaleString('ja-JP', {
            year: 'numeric',
            month: '2-digit',
            day: '2-digit',
            hour: '2-digit',
            minute: '2-digit',
        })
    }

    return (
        <div className={styles['contact-list']}>
            <div className={styles['list-content']}>
                {contacts.length === 0 ? (
                    <div className={styles['empty-message']}>お問い合わせがありません</div>
                ) : (
                    <Virtuoso
                        computeItemKey={(_index, contact) => contact.id}
                        data={contacts}
                        itemContent={(_index, contact) => (
                            <ListItem onClick={() => handleContactClick(contact)}>
                                <div className={styles['contact-item']}>
                                    <div className={styles['contact-header']}>
                                        <div className={styles['contact-name']}>
                                            <Person className={styles['icon']} fontSize="small" />
                                            {contact.name}
                                            {contact.company && <span className={styles['company']}>({contact.company})</span>}
                                        </div>
                                        <div className={styles['contact-date']}>{formatDate(contact.createdAt)}</div>
                                    </div>
                                    <div className={styles['contact-email']}>
                                        <Email className={styles['icon']} fontSize="small" />
                                        {contact.email}
                                    </div>
                                    <div className={styles['contact-preview']}>
                                        {contact.content.length > 100 ? `${contact.content.substring(0, 100)}...` : contact.content}
                                    </div>
                                </div>
                            </ListItem>
                        )}
                    />
                )}
            </div>

            <Dialog
                confirmOption={{
                    label: '閉じる',
                    onClick: handleCloseDetailDialog,
                }}
                isOpen={isDetailDialogOpen}
                onClose={handleCloseDetailDialog}
                title="お問い合わせ詳細"
            >
                {selectedContact && (
                    <div className={styles['contact-detail']}>
                        <div className={styles['detail-row']}>
                            <label className={styles['detail-label']}>お名前</label>
                            <div className={styles['detail-value']}>{selectedContact.name}</div>
                        </div>
                        {selectedContact.company && (
                            <div className={styles['detail-row']}>
                                <label className={styles['detail-label']}>会社名</label>
                                <div className={styles['detail-value']}>{selectedContact.company}</div>
                            </div>
                        )}
                        <div className={styles['detail-row']}>
                            <label className={styles['detail-label']}>メールアドレス</label>
                            <div className={styles['detail-value']}>{selectedContact.email}</div>
                        </div>
                        {selectedContact.phoneNumber && (
                            <div className={styles['detail-row']}>
                                <label className={styles['detail-label']}>電話番号</label>
                                <div className={styles['detail-value']}>{selectedContact.phoneNumber}</div>
                            </div>
                        )}
                        <div className={styles['detail-row']}>
                            <label className={styles['detail-label']}>受信日時</label>
                            <div className={styles['detail-value']}>{formatDate(selectedContact.createdAt)}</div>
                        </div>
                        <div className={styles['detail-row']}>
                            <label className={styles['detail-label']}>お問い合わせ内容</label>
                            <div className={styles['detail-content']}>{selectedContact.content}</div>
                        </div>
                    </div>
                )}
            </Dialog>
        </div>
    )
}
