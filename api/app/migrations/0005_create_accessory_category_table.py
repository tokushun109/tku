# Generated by Django 3.1.5 on 2021-01-21 00:00

from django.db import migrations, models
import uuid


class Migration(migrations.Migration):

    dependencies = [
        ('app', '0004_add_table_verbose_name_plural_Product_table'),
    ]

    operations = [
        migrations.CreateModel(
            name='AccessoryCategory',
            fields=[
                ('id', models.AutoField(auto_created=True, primary_key=True, serialize=False, verbose_name='ID')),
                ('uuid', models.UUIDField(default=uuid.uuid4, editable=False)),
                ('name', models.CharField(max_length=120, verbose_name='アクセサリーカテゴリー名')),
                ('created_at', models.DateTimeField(auto_now_add=True, verbose_name='作成日')),
                ('updated_at', models.DateTimeField(auto_now=True, verbose_name='更新日')),
            ],
            options={
                'verbose_name_plural': 'アクセサリーカテゴリー',
            },
        ),
    ]
