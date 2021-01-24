from rest_framework import serializers
from .models import Product, AccessoryCategory


class ProductSerializer(serializers.ModelSerializer):
    class Meta:
        model = Product
        fields = ('uuid', 'name', 'description')


class AccessoryCategorySerializer(serializers.ModelSerializer):
    class Meta:
        model = AccessoryCategory
        fields = ('uuid', 'name')
