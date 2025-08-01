'use client'

import { useState } from 'react'

import { Button } from '@/components/bases/Button'
import { Card } from '@/components/bases/Card'
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
            const response = await fetch('/api/csv/product', {
                credentials: 'include',
            })

            if (!response.ok) {
                throw new Error('CSVダウンロードに失敗しました')
            }

            const csvText = await response.text()
            const blob = new Blob([csvText], { type: 'text/csv' })
            const link = document.createElement('a')
            link.href = URL.createObjectURL(blob)
            link.download = '商品レコード.csv'
            link.click()
        } catch (error) {
            console.error('CSV download error:', error)
            setErrorMessage('CSVダウンロードに失敗しました')
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
            const formData = new FormData()
            formData.append('csv', uploadFile)

            const response = await fetch('/api/csv/product', {
                method: 'POST',
                body: formData,
                credentials: 'include',
            })

            if (!response.ok) {
                throw new Error('CSVアップロードに失敗しました')
            }

            setSuccessMessage('アップロードを完了しました')
            setIsDialogVisible(false)
            setUploadFile(null)
            setErrorMessage('')
        } catch (error) {
            console.error('CSV upload error:', error)
            setErrorMessage('CSVアップロードに失敗しました')
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
            <Card>
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
            </Card>

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
