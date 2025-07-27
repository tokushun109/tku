'use client'

import { Add, Delete } from '@mui/icons-material'
import { Virtuoso } from 'react-virtuoso'

import { Button } from '@/components/bases/Button'
import { Dialog } from '@/components/bases/Dialog'
import { ExternalLink } from '@/components/bases/ExternalLink'
import { ListItem } from '@/components/bases/ListItem'
import { ISite } from '@/features/site/type'
import { SiteType, SiteLabel } from '@/types'

import { SiteFormDialog } from '../SiteFormDialog'
import { useSiteList } from './hooks'
import styles from './styles.module.scss'

interface Props {
    sites: ISite[]
    siteType: SiteType
}

export const SiteList = ({ sites, siteType }: Props) => {
    const {
        items,
        isOpen,
        isSubmitting,
        submitError,
        targetItem,
        isDeleteDialogOpen,
        handleOpenDialog,
        handleCloseDialog,
        handleFormSubmit,
        handleOpenDeleteDialog,
        handleCloseDeleteDialog,
        handleConfirmDelete,
    } = useSiteList({
        initialItems: sites,
        siteType,
    })

    const siteName = SiteLabel[siteType]

    return (
        <div className={styles['site-list']}>
            <div className={styles['list-content']}>
                {items.length === 0 ? (
                    <div className={styles['empty-message']}>登録されていません</div>
                ) : (
                    <Virtuoso
                        computeItemKey={(_index, item) => item.uuid || _index}
                        data={items}
                        itemContent={(_index, item) => (
                            <ListItem
                                actions={
                                    <Delete
                                        className={styles['icon-button']}
                                        fontSize="medium"
                                        onClick={(e) => {
                                            e.stopPropagation()
                                            handleOpenDeleteDialog(item)
                                        }}
                                    />
                                }
                                onClick={() => {
                                    handleOpenDialog(item)
                                }}
                                subItem={
                                    item.url ? (
                                        <ExternalLink className={styles['site-link']} href={item.url}>
                                            {item.url}
                                        </ExternalLink>
                                    ) : (
                                        <span className={styles['site-url-empty']}>URLが設定されていません</span>
                                    )
                                }
                            >
                                <span className={styles['site-name']}>{item.name}</span>
                            </ListItem>
                        )}
                    />
                )}
            </div>
            <div className={styles['add-button-container']}>
                <Button
                    onClick={() => {
                        handleOpenDialog(null)
                    }}
                >
                    <div className={styles['add-button-content']}>
                        <Add className={styles['add-icon']} fontSize="large" />
                        追加
                    </div>
                </Button>
            </div>
            <SiteFormDialog
                isOpen={isOpen}
                isSubmitting={isSubmitting}
                onClose={handleCloseDialog}
                onSubmit={handleFormSubmit}
                siteType={siteType}
                submitError={submitError}
                updateItem={targetItem}
            />
            <Dialog
                cancelOption={{
                    label: 'キャンセル',
                    onClick: handleCloseDeleteDialog,
                }}
                confirmOption={{
                    label: '削除',
                    onClick: handleConfirmDelete,
                }}
                isOpen={isDeleteDialogOpen}
                onClose={handleCloseDeleteDialog}
                title="削除確認"
            >
                {targetItem && (
                    <>
                        <p>
                            {siteName}「{targetItem.name}」を削除しますか？
                        </p>
                        <p>この操作は取り消せません。</p>
                    </>
                )}
            </Dialog>
        </div>
    )
}
