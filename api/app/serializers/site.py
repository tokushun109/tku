from rest_framework import serializers
from app.models.site import SalesSite, SkillMarket, Sns


class SalesSiteSerializer(serializers.ModelSerializer):
    class Meta:
        model = SalesSite
        fields = ('id', 'name', 'url')


class SkillMarketSerializer(serializers.ModelSerializer):
    class Meta:
        model = SkillMarket
        fields = ('id', 'name', 'url')


class SnsSerializer(serializers.ModelSerializer):
    class Meta:
        model = Sns
        fields = ('id', 'name', 'url')
