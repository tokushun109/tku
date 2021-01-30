from django.db import models
from django.utils import timezone

import uuid


class AccessoryCategory(models.Model):
    '''
    アクセサリーカテゴリーモデル
    '''
    class Meta:
        verbose_name_plural = 'アクセサリーカテゴリー'

    # 名前
    name = models.CharField('アクセサリーカテゴリー名', max_length=120)
    # 作成日
    created_at = models.DateTimeField(
        '作成日', auto_now_add=True)
    # 更新日
    updated_at = models.DateTimeField(
        '更新日', auto_now=True)

    def __str__(self):
        return self.name

class MaterialCategory(models.Model):
    '''
    材料カテゴリーモデル
    '''
    class Meta:
        verbose_name_plural = '材料カテゴリー'

    # 名前
    name = models.CharField('材料カテゴリー名', max_length=120)
    # 作成日
    created_at = models.DateTimeField(
        '作成日', auto_now_add=True)
    # 更新日
    updated_at = models.DateTimeField(
        '更新日', auto_now=True)

    def __str__(self):
        return self.name

class SalesSite(models.Model):
    '''
    販売サイトモデル
    '''
    class Meta:
        verbose_name_plural = '販売サイト'

    # 名前
    name = models.CharField('材料カテゴリー名', max_length=120)
    # url
    url = models.URLField('販売サイトURL', max_length=200, null=False)
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
    uuid = models.UUIDField(default=uuid.uuid4, editable=False, primary_key=True)
    # 名前
    name = models.CharField('商品名', max_length=120)
    # 商品説明
    description = models.TextField('商品説明')
    # アクセサリーカテゴリー
    accessory_category = models.ForeignKey(
        AccessoryCategory, verbose_name='アクセサリーカテゴリー', on_delete=models.SET_NULL, null=True)
    # 材料カテゴリー
    material_category = models.ManyToManyField(MaterialCategory, verbose_name='材料カテゴリー', null=True)
    # 販売サイト
    sales_site = models.ManyToManyField(SalesSite, verbose_name='販売サイト', null=True)
    # 作成日
    created_at = models.DateTimeField(
        '作成日', auto_now_add=True)
    # 更新日
    updated_at = models.DateTimeField(
        '更新日', auto_now=True)

    def __str__(self):
        return self.name
