from django.db import models
from django.utils import timezone


class ProducerProfile(models.Model):
    '''
    製作者プロフィールのモデル
    '''
    class Meta:
        verbose_name_plural = '製作者'
        ordering = ['created_at']

    # 製作者名
    name = models.CharField('製作者名', max_length=120)
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
