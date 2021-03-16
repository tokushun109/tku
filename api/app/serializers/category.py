from rest_framework import serializers
from app.models.category import AccessoryCategory, MaterialCategory


class AccessoryCategorySerializer(serializers.ModelSerializer):
    class Meta:
        model = AccessoryCategory
        fields = ('id', 'name')


class MaterialCategorySerializer(serializers.ModelSerializer):
    class Meta:
        model = MaterialCategory
        fields = ('id', 'name')
