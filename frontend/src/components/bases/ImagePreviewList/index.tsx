import { Close, DragIndicator } from '@mui/icons-material'
import React from 'react'

import { ColorType, FontSizeType } from '@/types'

import { Chip, ChipSize } from '../Chip'
import { Image } from '../Image'
import styles from './styles.module.scss'

export interface ImageItem {
    id: string
    order?: number
    src: string
    type: 'existing' | 'new'
}

interface Props {
    images: ImageItem[]
    onDelete: (_id: string) => void
    onReorder: (_dragId: string, _hoverId: string) => void
    title?: string
}

export const ImagePreviewList = ({ images, onDelete, onReorder, title = '画像プレビュー' }: Props) => {
    const handleDeleteClick = (id: string) => {
        onDelete(id)
    }

    const handleDragStart = (e: React.DragEvent, dragId: string) => {
        e.dataTransfer.setData('text/plain', dragId)
    }

    const handleDragOver = (e: React.DragEvent) => {
        e.preventDefault()
    }

    const handleDrop = (e: React.DragEvent, hoverId: string) => {
        e.preventDefault()
        const dragId = e.dataTransfer.getData('text/plain')
        if (dragId !== hoverId) {
            onReorder(dragId, hoverId)
        }
    }

    if (images.length === 0) {
        return null
    }

    return (
        <div className={styles.container}>
            <h4 className={styles.title}>{title}</h4>
            <div className={styles['image-grid']}>
                {images.map((image) => (
                    <div
                        className={styles['image-item']}
                        draggable
                        key={image.id}
                        onDragOver={handleDragOver}
                        onDragStart={(e) => handleDragStart(e, image.id)}
                        onDrop={(e) => handleDrop(e, image.id)}
                    >
                        <div className={styles['image-wrapper']}>
                            <Image alt="プレビュー" src={image.src} />
                            <div className={styles['image-overlay']}>
                                <button className={styles['delete-button']} onClick={() => handleDeleteClick(image.id)} type="button">
                                    <Close fontSize="small" />
                                </button>
                            </div>
                        </div>
                    </div>
                ))}
            </div>
        </div>
    )
}
