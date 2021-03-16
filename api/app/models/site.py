from django.db import models
from django.utils import timezone


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
