from rest_framework import serializers
from .models import Product, AccessoryCategory, MaterialCategory


class ProductSerializer(serializers.ModelSerializer):
    class Meta:
        model = Product
        fields = ('uuid', 'name', 'description', 'accessory_category')


class AccessoryCategorySerializer(serializers.ModelSerializer):
    class Meta:
        model = AccessoryCategory
        fields = ('id', 'name')

class MaterialCategorySerializer(serializers.ModelSerializer):
    class Meta:
        model = MaterialCategory
        fields = ('id', 'name')
