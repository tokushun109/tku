from django.db import models
from django.utils import timezone

import uuid


class AccessoryCategory(models.Model):
    '''
    アクセサリーカテゴリーモデル
    '''
    class Meta:
        verbose_name_plural = 'アクセサリーカテゴリー'

    # アクセサリーカテゴリー名
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

    # 材料カテゴリー名
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

    # 販売サイト名
    name = models.CharField('販売サイト名', max_length=120)
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


class ProducerProfile(models.Model):
    '''
    製作者プロフィールのモデル
    '''
    class Meta:
        verbose_name_plural = '製作者'
        ordering = ['created_at']

    # 製作者名
    name = models.CharField('製作者名', max_length=120, primary_key=True)
    # 紹介
    introduction = models.TextField('製作者の紹介')
    # 製作者のロゴ
    logo = models.ImageField(verbose_name='ロゴ', upload_to='images/')
    # 作成日
    created_at = models.DateTimeField(
        '作成日', auto_now_add=True)
    # 更新日
    updated_at = models.DateTimeField(
        '更新日', auto_now=True)

    def __str__(self):
        return self.name


class SkillMarket(models.Model):
    '''
    スキルマーケットモデル
    '''
    class Meta:
        verbose_name_plural = 'スキルマーケット'

    # スキルマーケット名
    name = models.CharField('スキルマーケット名', max_length=120)
    # url
    url = models.URLField('スキルマーケットURL', max_length=200, null=False)
    # 作成日
    created_at = models.DateTimeField(
        '作成日', auto_now_add=True)
    # 更新日
    updated_at = models.DateTimeField(
        '更新日', auto_now=True)

    def __str__(self):
        return self.name


class Sns(models.Model):
    '''
    SNSモデル
    '''
    class Meta:
        verbose_name_plural = 'SNS'

    # SNS名
    name = models.CharField('SNS名', max_length=120)
    # url
    url = models.URLField('SNSのURL', max_length=200, null=False)
    # 作成日
    created_at = models.DateTimeField(
        '作成日', auto_now_add=True)
    # 更新日
    updated_at = models.DateTimeField(
        '更新日', auto_now=True)

    def __str__(self):
        return self.name
