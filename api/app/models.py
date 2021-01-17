from django.db import models
from django.utils import timezone

import uuid


class Product(models.Model):
    '''
    商品のモデル
    '''
    # uuid
    uuid = models.UUIDField(default=uuid.uuid4, editable=False)
    # 名前
    name = models.CharField('商品名', max_length=120)
    # 商品説明
    description = models.TextField('商品説明', max_length=3000)
    # 作成日
    created_at = models.DateTimeField(
        '作成日', auto_now_add=True)
    # 更新日
    updated_at = models.DateTimeField(
        '更新日', auto_now=True)

    def __str__(self):
        return self.name
