from django.db import models
from django.utils import timezone


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
