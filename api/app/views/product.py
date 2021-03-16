from rest_framework import generics

from app.models.product import Product
from app.serializers.product import ProductSerializer


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
