from rest_framework import serializers
from .models import Product, AccessoryCategory, MaterialCategory, SalesSite, ProductImage, ProducerProfile, SkillMarket, Sns


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

class ProducerProfileSerializer(serializers.ModelSerializer):
    class Meta:
        model = ProducerProfile
        fields = ('name', 'introduction', 'logo')

class SkillMarketSerializer(serializers.ModelSerializer):
    class Meta:
        model = SkillMarket
        fields = ('id', 'name', 'url')

class SnsSerializer(serializers.ModelSerializer):
    class Meta:
        model = Sns
        fields = ('id', 'name', 'url')
