from django.db import models
from django.utils import timezone
from app.models import AccessoryCategory, MaterialCategory, SalesSite

import uuid


class ProductImage(models.Model):
    '''
    商品画像モデル
    '''
    class Meta:
        verbose_name_plural = '商品画像'
        ordering = ['created_at']

    # uuid
    uuid = models.UUIDField(
        default=uuid.uuid4, editable=False, primary_key=True)
    # 商品画像名
    name = models.CharField('商品画像名', max_length=120)
    # 商品画像
    image = models.ImageField(verbose_name='商品画像', upload_to='images/')
    # 作成日
    created_at = models.DateTimeField(
        '作成日', auto_now_add=True)
    # 更新日
    updated_at = models.DateTimeField(
        '更新日', auto_now=True)

    def __str__(self):
        return self.name


class Product(models.Model):
    '''
    商品のモデル
    '''
    class Meta:
        verbose_name_plural = '商品'
        ordering = ['created_at']

    # uuid
    uuid = models.UUIDField(
        default=uuid.uuid4, editable=False, primary_key=True)
    # 商品名
    name = models.CharField('商品名', max_length=120)
    # 商品説明
    description = models.TextField('商品説明')
    # アクセサリーカテゴリー
    accessory_category = models.ForeignKey(
        AccessoryCategory, verbose_name='アクセサリーカテゴリー', on_delete=models.SET_NULL, null=True)
    # 材料カテゴリー
    material_category = models.ManyToManyField(
        MaterialCategory, verbose_name='材料カテゴリー', null=True)
    # 商品画像
    product_image = models.ForeignKey(
        ProductImage, verbose_name='商品画像', on_delete=models.SET_NULL, null=True)
    # 販売サイト
    sales_site = models.ManyToManyField(
        SalesSite, verbose_name='販売サイト', null=True)
    # 作成日
    created_at = models.DateTimeField(
        '作成日', auto_now_add=True)
    # 更新日
    updated_at = models.DateTimeField(
        '更新日', auto_now=True)

    def __str__(self):
        return self.name
