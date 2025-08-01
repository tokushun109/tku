'use client'

import { useState } from 'react'

import { downloadProductCsv, uploadProductCsv } from '@/apis/csv'
import { Button } from '@/components/bases/Button'
import { Dialog } from '@/components/bases/Dialog'
import { FileInput } from '@/components/bases/FileInput'
import { Message, MessageType } from '@/components/bases/Message'

import styles from './styles.module.scss'

export const AdminCsvTemplate = () => {
    const [uploadFile, setUploadFile] = useState<File | null>(null)
    const [isDialogVisible, setIsDialogVisible] = useState<boolean>(false)
    const [successMessage, setSuccessMessage] = useState<string>('')
    const [errorMessage, setErrorMessage] = useState<string>('')

    const handleDownload = async () => {
        try {
            await downloadProductCsv()
        } catch (error) {
            setErrorMessage(error instanceof Error ? error.message : 'CSVダウンロードに失敗しました')
        }
    }

    const handleUploadDialog = () => {
        setIsDialogVisible(true)
        setErrorMessage('')
    }

    const handleConfirm = async () => {
        if (!uploadFile) {
            setErrorMessage('CSVファイルが添付されていません')
            return
        }

        try {
            await uploadProductCsv(uploadFile)

            setSuccessMessage('アップロードを完了しました')
            setIsDialogVisible(false)
            setUploadFile(null)
            setErrorMessage('')
        } catch (error) {
            setErrorMessage(error instanceof Error ? error.message : 'CSVアップロードに失敗しました')
        }
    }

    const handleClose = () => {
        setIsDialogVisible(false)
        setUploadFile(null)
        setErrorMessage('')
    }

    return (
        <div className={styles['csv-container']}>
            <div className={styles['page-header']}>
                <h1 className={styles['page-title']}>商品レコード</h1>
            </div>
            <div className={styles['divider']} />
            <div className={styles['csv-content']}>
                <div className={styles['csv-buttons']}>
                    <Button className={styles['button']} onClick={handleDownload}>
                        ダウンロード
                    </Button>
                    <Button className={styles['button']} onClick={handleUploadDialog}>
                        アップロード
                    </Button>
                </div>
            </div>

            <Dialog
                cancelOption={{
                    label: 'キャンセル',
                    onClick: handleClose,
                }}
                confirmOption={{
                    label: 'アップロード',
                    onClick: handleConfirm,
                }}
                isOpen={isDialogVisible}
                onClose={handleClose}
                title="CSVのアップロード"
            >
                <div className={styles['upload-content']}>
                    {errorMessage && (
                        <Message className={styles['error-message']} type={MessageType.Error}>
                            {errorMessage}
                        </Message>
                    )}
                    <FileInput accept=".csv" label="CSVファイル" onChange={setUploadFile} required value={uploadFile} />
                </div>
            </Dialog>

            {successMessage && (
                <Message className={styles['success-message']} type={MessageType.Success}>
                    {successMessage}
                </Message>
            )}
        </div>
    )
}
