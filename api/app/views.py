from django.shortcuts import get_object_or_404
from rest_framework import generics, status, views
from rest_framework.response import Response

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


class ProducerProfileAPIView(views.APIView):
    '''
    製作者モデルの取得APIクラス
    '''

    def get(self, request, *args, **kwargs):
        # pk=1の製作者のプロフィールを取得する
        producer_profile = get_object_or_404(ProducerProfile, pk=1)
        serializer = ProducerProfileSerializer(instance=producer_profile)
        return Response(serializer.data, status.HTTP_200_OK)


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
