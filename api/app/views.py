from rest_framework import generics

from .models import Product, AccessoryCategory, MaterialCategory
from .serializers import ProductSerializer, AccessoryCategorySerializer, MaterialCategorySerializer


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
    アクセサリーカテゴリーモデルの取得(一覧)APIクラス
    '''
    queryset = MaterialCategory.objects.all()
    serializer_class = MaterialCategorySerializer
