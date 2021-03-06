# Generated by Django 3.1.5 on 2021-01-30 12:04

from django.db import migrations, models


class Migration(migrations.Migration):

    dependencies = [
        ('app', '0002_create_some__category_table'),
    ]

    operations = [
        migrations.CreateModel(
            name='SalesSite',
            fields=[
                ('id', models.AutoField(auto_created=True,
                                        primary_key=True, serialize=False, verbose_name='ID')),
                ('name', models.CharField(max_length=120, verbose_name='材料カテゴリー名')),
                ('url', models.URLField(verbose_name='販売サイトURL')),
                ('created_at', models.DateTimeField(
                    auto_now_add=True, verbose_name='作成日')),
                ('updated_at', models.DateTimeField(
                    auto_now=True, verbose_name='更新日')),
            ],
            options={
                'verbose_name_plural': '販売サイト',
            },
        ),
        migrations.AddField(
            model_name='product',
            name='sales_site',
            field=models.ManyToManyField(
                null=True, to='app.SalesSite', verbose_name='販売サイト'),
        ),
    ]
