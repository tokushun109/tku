# Generated by Django 3.1.5 on 2021-01-17 12:20

from django.db import migrations, models
import django.utils.timezone


class Migration(migrations.Migration):

    dependencies = [
        ('app', '0002_add_uuid_Product_table'),
    ]

    operations = [
        migrations.AddField(
            model_name='product',
            name='created_at',
            field=models.DateTimeField(auto_now_add=True, default=django.utils.timezone.now, verbose_name='作成日'),
            preserve_default=False,
        ),
        migrations.AddField(
            model_name='product',
            name='updated_at',
            field=models.DateTimeField(auto_now=True, verbose_name='更新日'),
        ),
        migrations.AlterField(
            model_name='product',
            name='description',
            field=models.TextField(max_length=3000, verbose_name='商品説明'),
        ),
        migrations.AlterField(
            model_name='product',
            name='name',
            field=models.CharField(max_length=120, verbose_name='商品名'),
        ),
    ]
