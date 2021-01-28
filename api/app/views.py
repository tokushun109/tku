from rest_framework import generics

from .models import Product, AccessoryCategory
from .serializers import ProductSerializer, AccessoryCategorySerializer


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


class AccessoryListAPIView(generics.ListAPIView):
    '''
    アクセサリーカテゴリーモデルの取得(一覧)APIクラス
    '''

    queryset = AccessoryCategory.objects.all()
    serializer_class = AccessoryCategorySerializer
