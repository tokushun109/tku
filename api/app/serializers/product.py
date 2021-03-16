from rest_framework import serializers
from app.models.product import Product, ProductImage
from app.serializers import AccessoryCategorySerializer, MaterialCategorySerializer, SalesSiteSerializer

import uuid


class ProductImageSerializer(serializers.ModelSerializer):
    class Meta:
        model = ProductImage
        fields = ('uuid', 'name', 'image')


class ProductSerializer(serializers.ModelSerializer):
    accessory_category = AccessoryCategorySerializer()
    material_category = MaterialCategorySerializer(many=True)
    sales_site = SalesSiteSerializer(many=True)
    product_image = ProductImageSerializer()

    class Meta:
        model = Product
        fields = ('uuid', 'name', 'description', 'accessory_category',
                  'material_category', 'product_image', 'sales_site')
