# Generated by Django 3.1.5 on 2021-01-30 00:23

from django.db import migrations, models
import django.db.models.deletion
import django.utils.timezone


class Migration(migrations.Migration):

    dependencies = [
        ('app', '0001_ProductInit'),
    ]

    operations = [
        migrations.CreateModel(
            name='AccessoryCategory',
            fields=[
                ('id', models.AutoField(auto_created=True, primary_key=True, serialize=False, verbose_name='ID')),
                ('name', models.CharField(max_length=120, verbose_name='アクセサリーカテゴリー名')),
                ('created_at', models.DateTimeField(auto_now_add=True, verbose_name='作成日')),
                ('updated_at', models.DateTimeField(auto_now=True, verbose_name='更新日')),
            ],
            options={
                'verbose_name_plural': 'アクセサリーカテゴリー',
            },
        ),
        migrations.CreateModel(
            name='MaterialCategory',
            fields=[
                ('id', models.AutoField(auto_created=True, primary_key=True, serialize=False, verbose_name='ID')),
                ('name', models.CharField(max_length=120, verbose_name='材料カテゴリー名')),
                ('created_at', models.DateTimeField(auto_now_add=True, verbose_name='作成日')),
                ('updated_at', models.DateTimeField(auto_now=True, verbose_name='更新日')),
            ],
            options={
                'verbose_name_plural': '材料カテゴリー',
            },
        ),
        migrations.AlterModelOptions(
            name='product',
            options={'ordering': ['created_at'], 'verbose_name_plural': '商品'},
        ),
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
            field=models.TextField(verbose_name='商品説明'),
        ),
        migrations.AlterField(
            model_name='product',
            name='name',
            field=models.CharField(max_length=120, verbose_name='商品名'),
        ),
        migrations.AddField(
            model_name='product',
            name='accessory_category',
            field=models.ForeignKey(null=True, on_delete=django.db.models.deletion.SET_NULL, to='app.accessorycategory', verbose_name='アクセサリーカテゴリー'),
        ),
        migrations.AddField(
            model_name='product',
            name='material_category',
            field=models.ManyToManyField(null=True, to='app.MaterialCategory', verbose_name='材料カテゴリー'),
        ),
    ]