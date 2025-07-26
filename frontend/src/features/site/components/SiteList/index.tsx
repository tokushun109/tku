'use client'

import { Add, Delete } from '@mui/icons-material'
import { Virtuoso } from 'react-virtuoso'

import { Button } from '@/components/bases/Button'
import { Dialog } from '@/components/bases/Dialog'
import { ExternalLink } from '@/components/bases/ExternalLink'
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
        <div className={styles['site-list-container']}>
            <div className={styles['site-list-header']}>
                <h2 className={styles['site-list-title']}>{siteName}一覧</h2>
                <div className={styles['site-count']}>
                    <span>総件数: {items.length}件</span>
                </div>
            </div>
            <div className={styles['site-list']}>
                {items.length === 0 ? (
                    <div className={styles['empty-state']}>
                        <p>登録されていません</p>
                    </div>
                ) : (
                    <Virtuoso
                        computeItemKey={(_index, item) => item.uuid || _index}
                        data={items}
                        itemContent={(_index, item) => (
                            <div
                                className={styles['site-item']}
                                onClick={() => {
                                    handleOpenDialog(item)
                                }}
                            >
                                <div className={styles['site-info']}>
                                    <div className={styles['site-header']}>
                                        <h3 className={styles['site-name']}>{item.name}</h3>
                                        <div className={styles['item-actions']}>
                                            <Delete
                                                className={styles['icon-button']}
                                                onClick={(e) => {
                                                    e.stopPropagation()
                                                    handleOpenDeleteDialog(item)
                                                }}
                                            />
                                        </div>
                                    </div>
                                    <div className={styles['site-details']}>
                                        <div className={styles['site-url']}>
                                            {item.url ? (
                                                <ExternalLink className={styles['site-link']} href={item.url}>
                                                    {item.url}
                                                </ExternalLink>
                                            ) : (
                                                <span className={styles['site-url-empty']}>URLが設定されていません</span>
                                            )}
                                        </div>
                                    </div>
                                </div>
                            </div>
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
                        <Add className={styles['add-icon']} />
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
