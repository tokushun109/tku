from django.db import models


class Product(models.Model):
    '''
    商品のモデル
    '''
    # 名前
    name = models.CharField(max_length=120)
    # 商品説明
    description = models.TextField(max_length=3000)

    def __str__(self):
        return self.name
