from rest_framework import serializers
from .models import Product, AccessoryCategory, MaterialCategory, SalesSite, ProductImage


class ProductSerializer(serializers.ModelSerializer):
    class Meta:
        model = Product
        fields = ('uuid', 'name', 'description', 'accessory_category', 'material_category', 'product_image', 'sales_site')


class AccessoryCategorySerializer(serializers.ModelSerializer):
    class Meta:
        model = AccessoryCategory
        fields = ('id', 'name')

class MaterialCategorySerializer(serializers.ModelSerializer):
    class Meta:
        model = MaterialCategory
        fields = ('id', 'name')

class SalesSiteSerializer(serializers.ModelSerializer):
    class Meta:
        model = SalesSite
        fields = ('id', 'name', 'url')

class ProductImageSerializer(serializers.ModelSerializer):
    class Meta:
        model = ProductImage
        fields = ('uuid', 'name', 'image')