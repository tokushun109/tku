from rest_framework import generics

from .models import Product, AccessoryCategory, MaterialCategory, SalesSite, ProducerProfile, SkillMarket, Sns
from .serializers import ProductSerializer, AccessoryCategorySerializer, MaterialCategorySerializer, SalesSiteSerializer, ProducerProfileSerializer, SkillMarketSerializer, SnsSerializer


class ProductListAPIView(generics.ListAPIView):
    '''
    商品モデルの取得(一覧)APIクラス
    '''
    queryset = Product.objects.all()
    serializer_class = ProductSerializer


class ProductRetrieveAPIView(generics.RetrieveAPIView):
    '''
    商品モデルの取得(詳細)APIクラス
    '''
    queryset = Product.objects.all()
    serializer_class = ProductSerializer


class AccessoryCategoryListAPIView(generics.ListAPIView):
    '''
    アクセサリーカテゴリーモデルの取得(一覧)APIクラス
    '''
    queryset = AccessoryCategory.objects.all()
    serializer_class = AccessoryCategorySerializer


class MaterialCategoryListAPIView(generics.ListAPIView):
    '''
    材料カテゴリーモデルの取得(一覧)APIクラス
    '''
    queryset = MaterialCategory.objects.all()
    serializer_class = MaterialCategorySerializer


class SalesSiteListAPIView(generics.ListAPIView):
    '''
    販売サイトモデルの取得(一覧)APIクラス
    '''
    queryset = SalesSite.objects.all()
    serializer_class = SalesSiteSerializer


class ProducerProfileListAPIView(generics.ListAPIView):
    '''
    製作者モデルの取得(一覧)APIクラス
    '''
    queryset = ProducerProfile.objects.all()
    serializer_class = ProducerProfileSerializer


class SkillMarketListAPIView(generics.ListAPIView):
    '''
    販売サイトモデルの取得(一覧)APIクラス
    '''
    queryset = SkillMarket.objects.all()
    serializer_class = SkillMarketSerializer


class SnsListAPIView(generics.ListAPIView):
    '''
    販売サイトモデルの取得(一覧)APIクラス
    '''
    queryset = Sns.objects.all()
    serializer_class = SnsSerializer
